// Code generated by mockery v2.34.2. DO NOT EDIT.

package mocks

import (
	context "context"

	mock "github.com/stretchr/testify/mock"
	model "github.com/wejick/gchain/model"
)

// BaseChain is an autogenerated mock type for the BaseChain type
type BaseChain struct {
	mock.Mock
}

// Run provides a mock function with given fields: ctx, prompt, options
func (_m *BaseChain) Run(ctx context.Context, prompt map[string]string, options ...func(*model.Option)) (map[string]string, error) {
	_va := make([]interface{}, len(options))
	for _i := range options {
		_va[_i] = options[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx,