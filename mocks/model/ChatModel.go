// Code generated by mockery v2.34.2. DO NOT EDIT.

package mocks

import (
	context "context"

	mock "github.com/stretchr/testify/mock"
	model "github.com/wejick/gchain/model"
)

// ChatModel is an autogenerated mock type for the ChatModel type
type ChatModel struct {
	mock.Mock
}

// Call provides a mock function with given fields: ctx, prompt, options
func (_m *ChatModel) Call(ctx context.Context, prompt string, options ...func(*model.Option)) (string, error) {
	_va := make([]interface{}, len(options))
	for _i := range options {
		_va[_i] = options[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, prompt)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 string
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string, ...func(*model.Option)) (string, error)); ok {
		return rf(ctx, prompt, options...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string, ...func(*model.Option)) string); ok {
		r0 = rf(ctx, prompt, options...