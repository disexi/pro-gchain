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
	manager.Register