
// Code generated by mockery v2.34.2. DO NOT EDIT.

package mocks

import mock "github.com/stretchr/testify/mock"

// Evaluator is an autogenerated mock type for the Evaluator type
type Evaluator struct {
	mock.Mock
}

// Evaluate provides a mock function with given fields: input
func (_m *Evaluator) Evaluate(input string) (bool, error) {
	ret := _m.Called(input)

	var r0 bool
	var r1 error
	if rf, ok := ret.Get(0).(func(string) (bool, error)); ok {
		return rf(input)
	}
	if rf, ok := ret.Get(0).(func(string) bool); ok {
		r0 = rf(input)
	} else {
		r0 = ret.Get(0).(bool)
	}

	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(input)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewEvaluator creates a new instance of Evaluator. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewEvaluator(t interface {
	mock.TestingT
	Cleanup(func())
}) *Evaluator {
	mock := &Evaluator{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}