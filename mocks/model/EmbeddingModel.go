
// Code generated by mockery v2.34.2. DO NOT EDIT.

package mocks

import mock "github.com/stretchr/testify/mock"

// EmbeddingModel is an autogenerated mock type for the EmbeddingModel type
type EmbeddingModel struct {
	mock.Mock
}

// EmbedDocuments provides a mock function with given fields: documents
func (_m *EmbeddingModel) EmbedDocuments(documents []string) ([][]float32, error) {
	ret := _m.Called(documents)

	var r0 [][]float32
	var r1 error
	if rf, ok := ret.Get(0).(func([]string) ([][]float32, error)); ok {
		return rf(documents)
	}
	if rf, ok := ret.Get(0).(func([]string) [][]float32); ok {
		r0 = rf(documents)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([][]float32)
		}
	}

	if rf, ok := ret.Get(1).(func([]string) error); ok {
		r1 = rf(documents)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// EmbedQuery provides a mock function with given fields: input
func (_m *EmbeddingModel) EmbedQuery(input string) ([]float32, error) {
	ret := _m.Called(input)

	var r0 []float32
	var r1 error
	if rf, ok := ret.Get(0).(func(string) ([]float32, error)); ok {
		return rf(input)
	}
	if rf, ok := ret.Get(0).(func(string) []float32); ok {
		r0 = rf(input)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]float32)
		}
	}

	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(input)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewEmbeddingModel creates a new instance of EmbeddingModel. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewEmbeddingModel(t interface {
	mock.TestingT
	Cleanup(func())
}) *EmbeddingModel {
	mock := &EmbeddingModel{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}