package greeting

import (
	"context"
	"encoding/json"
	"errors"

	"github.com/wejick/gchain/model"
	"github.com/wejick/gchain/tools"
)

type GreetingTool struct {
	functionDefinition model.FunctionDefinition
}