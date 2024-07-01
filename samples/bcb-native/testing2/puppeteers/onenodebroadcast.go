package puppeteers

import (
	"context"
	"fmt"
	"time"

	"github.com/filecoin-project/mir/fuzzer2/nodeinstance"
	bcbevents "github.com/filecoin-project/mir/samples/bcb-native/events"
	"github.com/filecoin-project/mir/stdtypes"
)

type OneNodeBroadcast struct {
	node stdtypes.NodeID
}

func NewOneNodeBroadcast(nodeId stdtypes.NodeID) *OneNodeBroadcast {
	return &OneNodeBroadcast{
		nodeId,
	}
}

func (onb *OneNodeBroadcast) Run(nodeInstances map[stdtypes.NodeID]nodeinstance.NodeInstance) error {
	ctx := context.Background()
	time.Sleep(time.Second)
	nodeInstance := nodeInstances[onb.node]
	node := nodeInstance.GetNode()
	msg := fmt.Sprintf("node %s injecting", node.ID)
	return node.InjectEvents(ctx, stdtypes.ListOf(bcbevents.NewBroadcastRequest("bcb", []byte(msg))))
}