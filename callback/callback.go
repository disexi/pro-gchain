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

type CallbackData struct {
	RunID        string // to be populated with data from context
	EventName    string
	FunctionName string
	Input        map[string]string
	Output       map[string]string
	Data         interface{}
}

type Manager struct {
	callbacks           map[Event][]Callback
	callbackIdentifiers []CallbackIdentifier
}

func NewManager() *Manager {
	return &Manager{
		callbacks: make(map[Event][]Callback),
	}
}

func (m *Manager) Regis