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
	message       string // message 