// Code generated by Mir codegen. DO NOT EDIT.

package synchronizerpb

import (
	reflect "reflect"
)

func (*Event) ReflectTypeOptions() []reflect.Type {
	return []reflect.Type{
		reflect.TypeOf((*Event_SyncRequest)(nil)),
	}
}

func (*Message) ReflectTypeOptions() []reflect.Type {
	return []reflect.Type{
		reflect.TypeOf((*Message_ChainRequest)(nil)),
		reflect.TypeOf((*Message_ChainResponse)(nil)),
	}
}