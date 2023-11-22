package model

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFunctionJsonSchema_String(t *testing.T) {
	parameter := FunctionJsonSchema{}
	parameter.Required = []string{"parameter1"}
	parameter.Properties = map[string]FunctionJsonSchema{
		"parame