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

func NewGreetingTool() *GreetingTool {
	return &GreetingTool{
		functionDefinition: model.FunctionDefinition{
			Name:        "greeting_tool",
			Description: "This tool is used to greet user with hello",
			Parameters: model.FunctionJsonSchema{
				Type: model.FunctionDataTypeObject,
				Properties: map[string]model.FunctionJsonSchema{
					"user_name": {
						Type:        model.FunctionDataTypeString,
						Description: "User name"