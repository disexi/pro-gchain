package model

import "strings"

type DataType string

const (
	FunctionDataTypeString DataType = "string"
	FunctionDataTypeObject DataType = "object"
)

type FunctionJsonSchema struct {
	Type        DataType                      `json:"type,omitempty"`
	Properties  map[string]FunctionJsonSchema `json:"properties,omitempty"`
	Required   