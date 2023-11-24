// Code generated by moq; DO NOT EDIT.
// github.com/matryer/moq

package model

import (
	"context"
	"sync"
)

// Ensure, that LLMModelMock does implement LLMModel.
// If this is not the case, regenerate this file with moq.
var _ LLMModel = &LLMModelMock{}

// LLMModelMock is a mock implementation of LLMModel.
//
//	func TestSomethingThatUsesLLMModel(t *testing.T) {
//
//		// make and configure a mocked LLMModel
//		mockedLLMModel := &LLMModelMock{
//			CallFunc: func(ctx context.Context, prompt string, options ...func(*Option)) (string, error) {
//				panic("mock out the Call method")
//			},
//		}
//
//		// use mockedLLMModel in code that requires LLMModel
//		// and then make assertions.
//
//	}
type LLMModelMock struct {
	// CallFunc mocks the Call method.
	CallFunc func(ctx context.Context, prompt string, options ...func(*Option)) (string, error)

	// calls tracks calls to the methods.
	calls struct {
		// Call