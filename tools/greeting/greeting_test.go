package greeting

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGreetingsTool(t *testing.T) {
	greetingTool := NewGreetingTool()
	assert.NotNil(t, greetingTool)

	assert.NotNil(t, greetingTool.GetFunctionDefinition())

	// test simple run
	output, err := greetingTool.SimpleRun(context