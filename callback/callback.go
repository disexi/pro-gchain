package callback

import (
	"context"
	"reflect"

	"github.com/k0kubun/pp"
)

type Event string

type Callback func(context context.Context, data CallbackData)
type CallbackIdentifier struct {
	EventName   Event
	FuncPointer uintptr
}

type CallbackData struct