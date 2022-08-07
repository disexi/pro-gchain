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
	if len(event1Call