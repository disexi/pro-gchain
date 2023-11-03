
// Code generated by mockery v2.34.2. DO NOT EDIT.

package mocks

import (
	mock "github.com/stretchr/testify/mock"
	model "github.com/wejick/gchain/model"
)

// Memory is an autogenerated mock type for the Memory type
type Memory struct {
	mock.Mock
}

// AddMessage provides a mock function with given fields: message
func (_m *Memory) AddMessage(message model.ChatMessage) {
	_m.Called(message)
}

// GetAll provides a mock function with given fields:
func (_m *Memory) GetAll() []model.ChatMessage {
	ret := _m.Called()

	var r0 []model.ChatMessage
	if rf, ok := ret.Get(0).(func() []model.ChatMessage); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]model.ChatMessage)
		}
	}

	return r0
}

// String provides a mock function with given fields:
func (_m *Memory) String() string {
	ret := _m.Called()

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// NewMemory creates a new instance of Memory. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMemory(t interface {
	mock.TestingT
	Cleanup(func())
}) *Memory {
	mock := &Memory{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}