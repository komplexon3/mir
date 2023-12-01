/*
Copyright IBM Corp. All Rights Reserved.

SPDX-License-Identifier: Apache-2.0
*/

// TODO: Properly comment all the code in this file.
//       Also revise the existing comments and make sure they are consistent and understandable.

package eventlog

import (
	"os"
	"path/filepath"
	"strconv"
	"sync"
	"time"

	es "github.com/go-errors/errors"
	"github.com/pkg/errors"

	"github.com/filecoin-project/mir/pkg/events"
	"github.com/filecoin-project/mir/pkg/logging"
	"github.com/filecoin-project/mir/pkg/pb/eventpb"
	t "github.com/filecoin-project/mir/pkg/types"
)

type eventRecord struct {
	Events *events.EventList
	Time   int64
}

// Recorder is intended to be used as an implementation of the
// mir.EventInterceptor interface.  It receives state Events,
// serializes them, compresses them, and writes them to a stream.
type Recorder struct {
	nodeID         t.NodeID
	dest           EventWriter
	timeSource     func() int64
	newDests       func(list *events.EventList) []*events.EventList
	path           string
	filter         func(event *eventpb.Event) bool
	newEventWriter func(dest string, nodeID t.NodeID, logger logging.Logger) (EventWriter, error)
	syncWrite      bool

	logger     logging.Logger
	eventC     chan eventRecord
	doneC      chan struct{}
	exitC      chan struct{}
	fileCount  int
	writerLock sync.Mutex

	exitErr      error
	exitErrMutex sync.Mutex
}

func NewRecorder(
	nodeID t.NodeID,
	path string, // TODO: Remove the path argument. It is writer-specific and should be specified in the writer factory.
	logger logging.Logger,
	opts ...RecorderOpt,
) (*Recorder, error) {
	if logger == nil {
		logger = logging.ConsoleErrorLogger
	}

	startTime := time.Now()

	i := &Recorder{
		nodeID: nodeID,
		timeSource: func() int64 {
			return time.Since(startTime).Milliseconds()
		},
		eventC:    make(chan eventRecord, DefaultBufferSize),
		doneC:     make(chan struct{}),
		exitC:     make(chan struct{}),
		fileCount: 1,
		newDests:  OneFileLogger(),
		path:      path,
		filter: func(_ *eventpb.Event) bool {
			// Record all events by default.
			return true
		},
		newEventWriter: DefaultNewEventWriter,
		syncWrite:      false,
		logger:         logger,
	}

	for _, opt := range opts {
		switch v := opt.(type) {
		case timeSourceOpt:
			i.timeSource = v
		case bufferSizeOpt:
			i.eventC = make(chan eventRecord, v)
		case fileSplitterOpt:
			i.newDests = v
		case eventFilterOpt:
			// Apply the given filter on top of the existing one.
			oldFilter := i.filter
			i.filter = func(e *eventpb.Event) bool {
				return oldFilter(e) && v(e)
			}
		case eventWriterOpt:
			i.newEventWriter = v
		case syncWriteOpt:
			i.syncWrite = true
		}
	}

	if err := os.MkdirAll(path, 0700); err != nil {
		return nil, es.Errorf("error creating interceptor directory: %w", err)
	}

	writer, err := i.newEventWriter(filepath.Join(path, "eventlog0"), nodeID, logger)
	if err != nil {
		return nil, es.Errorf("error initializing event writer: %w", err)
	}
	i.dest = writer

	// Only start background writing goroutine if synchronous writing is not enabled.
	if !i.syncWrite {
		go func() {
			err := i.run()
			if err != nil {
				logger.Log(logging.LevelError, "Interceptor returned with error.", "err", err)
			} else {
				logger.Log(logging.LevelDebug, "Interceptor returned successfully.")
			}
		}()
	}

	return i, nil
}

// Intercept takes an event and enqueues it into the event buffer.
// If there is no room in the buffer, it blocks.  If draining the buffer
// to the output stream has completed (successfully or otherwise), Intercept
// returns an error.
func (i *Recorder) Intercept(events *events.EventList) (*events.EventList, error) {

	// Avoid nil dereference if Intercept is called on a nil *Recorder and simply do nothing.
	// This can happen if a pointer type to *Recorder is assigned to a variable with the interface type Interceptor.
	// Mir would treat that variable as non-nil, thinking there is an interceptor, and call Intercept() on it.
	// For more explanation, see https://mangatmodi.medium.com/go-check-nil-interface-the-right-way-d142776edef1
	if i == nil {
		return events, nil
	}

	// If synchronous writing is enabled, write data and return immediately, without using any channels.
	// Event modification is only possible here (in synchronous mode).
	if i.syncWrite {
		i.writerLock.Lock()
		defer i.writerLock.Unlock()
		newEvents, err := i.writeEvents(events, i.timeSource())
		if err != nil {
			return nil, es.Errorf("error writing events: %w", err)
		}
		if err := i.dest.Flush(); err != nil {
			return nil, es.Errorf("error flushing written events: %w", err)
		}
		return newEvents, nil
	}

	// If writing is asynchronous, pass the record to the background writing goroutine.
	// Note that the recorder will not modify the event list in asynchronous mode, even if the used event writer
	// returns a modified event list.
	select {
	case i.eventC <- eventRecord{events, i.timeSource()}:
		return events, nil
	case <-i.exitC:
		i.exitErrMutex.Lock()
		defer i.exitErrMutex.Unlock()
		return events, i.exitErr
	}
}

// Stop must be invoked to release the resources associated with this
// Interceptor, and should only be invoked after the mir node has completely
// exited.  The returned error
func (i *Recorder) Stop() error {

	// In synchronous mode, the interceptor has nothing to stop.
	// Since we assume that the Mir node has already completely stopped,
	// we do not need to worry about concurrent or later invocations of Intercept().
	if i.syncWrite {
		return nil
	}

	// Send stop signal to writer thread and wait for it to stop.
	close(i.doneC)
	<-i.exitC

	// Close destination file.
	err := i.dest.Close()
	if err != nil {
		i.logger.Log(logging.LevelError, "Failed to close event writer.", "error", err)
	}

	// Return final error if it is different from the expected errStopped.
	i.exitErrMutex.Lock()
	defer i.exitErrMutex.Unlock()
	if errors.Is(i.exitErr, errStopped) {
		return nil
	}
	return i.exitErr
}

var errStopped = es.Errorf("interceptor stopped at caller request")

func (i *Recorder) run() (exitErr error) {
	cnt := 0 // Counts total number of Events written.

	defer func() {
		i.exitErrMutex.Lock()
		i.exitErr = exitErr
		i.exitErrMutex.Unlock()
		close(i.exitC)
		i.logger.Log(logging.LevelInfo, "Intercepted Events written to event log.", "numEvents", cnt, ", path", i.path)
	}()

	// Keep reading records of intercepted events and pass them to the event writer.
	// Note that the return value of writeEvents is ignored.
	// This is because, in asynchronous mode, the interceptor cannot modify the event stream.
	// All its modifications are thus ignored.
	for {
		select {
		case <-i.doneC:
			for {
				select {
				case record := <-i.eventC:
					var err error
					_, err = i.writeEvents(record.Events, record.Time)
					if err != nil {
						return es.Errorf("error serializing to stream: %w", err)
					}
				default:
					return errStopped
				}
			}
		case record := <-i.eventC:
			var err error
			_, err = i.writeEvents(record.Events, record.Time)
			if err != nil {
				return es.Errorf("error serializing to stream: %w", err)
			}
			cnt++
		}
	}
}

func (i *Recorder) writeEvents(evts *events.EventList, timestamp int64) (*events.EventList, error) {
	eventByDests := i.newDests(evts)
	count := 0
	eventsOut := events.EmptyList()
	for _, rec := range eventByDests {
		if count > 0 {
			// newDest required
			dest, err := i.newEventWriter(
				filepath.Join(i.path, "eventlog"+strconv.Itoa(i.fileCount)),
				i.nodeID,
				i.logger,
			)
			if err != nil {
				return nil, err
			}
			err = i.dest.Close()
			if err != nil {
				return nil, err
			}
			i.dest = dest
			i.fileCount++
		}

		if rec.Len() != 0 {
			var err error
			evtsOut, err := i.dest.Write(rec.Filter(i.filter), timestamp)
			if err != nil {
				return nil, err
			}
			eventsOut.PushBackList(evtsOut)
		}
		count++
	}
	return eventsOut, nil
}
