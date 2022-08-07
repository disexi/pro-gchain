package callback

import (
	"context"
	"fmt"
	"testing"
)

func TestManager_RegisterCallback(t *testing.T) {
	// Create a new Manager
	manager := NewManager()

	// Register some callbacks
	manager.RegisterCallback("event1", callback1)
	manager.RegisterCallback("event2", callback2)
	manager.RegisterCallback("event1", callback3)

	// Register duplicate callback for event1
	manager.RegisterCallback("event1", callback1)

	// Inspect the registered callbacks
	registeredCallbacks := manager.inspect()

	// Assert the number of registered callbacks
	if len(registeredCallbacks) != 2 {
		t.Errorf("Expected 2 registered events, got %d", len(registeredCallbacks))
	}

	// Assert the registered callbacks for a specific event
	// this will check if there's duplicate or not
	event1Callbacks := registeredCallbacks["event1"]
	if len(event1Callbacks) != 2 {
		t.Errorf("Expected 2 registered callbacks for event1, got %d", len(event1Callbacks))
	}

	event2Callbacks := registeredCallbacks["event2"]
	if len(event2Callbacks) != 1 {
		t.Errorf("Expected 1 registered callback for event2, got %d", len(event2Callbacks))
	}
}

func TestManager_TriggerEvent(t *testing.T) {
	// Create a new Manager
	manager := NewManager()

	// Variables to track callback invocations
	var callb