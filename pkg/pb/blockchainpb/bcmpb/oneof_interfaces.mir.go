// Code generated by Mir codegen. DO NOT EDIT.

package bcmpb

type Event_Type = isEvent_Type

type Event_TypeWrapper[T any] interface {
	Event_Type
	Unwrap() *T
}

func (w *Event_NewBlock) Unwrap() *NewBlock {
	return w.NewBlock
}
