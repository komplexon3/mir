package vcinterceptor

import (
	"github.com/filecoin-project/mir/pkg/pb/eventpb"
	"github.com/filecoin-project/mir/pkg/vcinterceptor/vectorclock"
	"github.com/filecoin-project/mir/stdtypes"
	es "github.com/go-errors/errors"
)

func ExtractVCFromPbEvent(pEvt *eventpb.Event) (*vectorclock.VectorClock, error) {
	vc := vectorclock.NewVectorClock()
	vcRaw, err := pEvt.GetMetadata("vc")
	if err != nil {
		return nil, err
	}

	// TODO: it being under V is ugly anyways - why isn't this JUST a map again?
	// HERE - this is failing

	// TODO: this is ugly af make it better
	values, ok := vcRaw.(map[string]interface{})["V"].(map[string]interface{})
	if !ok {
		return nil, es.Errorf("Failed to extract (V) from raw vector clock")
	}

	for n, c := range values {
		vc.V[stdtypes.NodeID(n)] = uint32(c.(float64))
	}

	return vc, nil
}