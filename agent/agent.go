package agent

import (
	"context"
	"errors"

	"github.com/wejick/gchain/tools"
)

var (
	ErrMaxIteration = errors.New("max iteration reached")
)

// Action define what to execute and also containing the result
type Action struct {
	toolName      string // tool name to run by executor
	toolInputJson string // input for the tool in json
	toolOutput    string // output from tool
	message       string // message from agent llm prediction
	finalAction   bool   // is it final action or not
}

// BaseAgent
type BaseAgent interface {
	// Plan is determine what action to text next
	Plan(ctx context.Context, userPrompt string, actionTaken []Action) (plan Action, err error)
}

// BaseExecutor
type BaseExecutor interface {
	Run(ctx context.Context, input map[string]string) (output map[string]string, err error)
	RegisterTool(tool *tools.BaseTool)
}

// Executor
type Executor struct {
	agent        BaseAgent
	tools        map[string]tools.BaseTool
	maxIteration int
}

// NewExecutor create new executor
func NewExecutor(agent BaseAgent, maxIteration int) *Executor {
	if maxIteration == 0 {
		maxIteration = 10
	}
	return &Executor{
		agent:        agent,
		tools:        map[string]tools.BaseTool{},
		maxIteration: maxIteration,
	}
}

// RegisterTool put tool to executor
func (E *Executor) RegisterTool(tool tools.Bas