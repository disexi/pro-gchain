// Code generated by mockery v2.34.2. DO NOT EDIT.

package mocks

import (
	context "context"

	datastore "github.com/wejick/gchain/datastore"
	document "github.com/wejick/gchain/document"

	mock "github.com/stretchr/testify/mock"
)

// VectorStore is an autogenerated mock type for the VectorStore type
type VectorStore struct {
	mock.Mock
}

// AddDocuments provides a mock function with given fields: ctx, indexName, documents
func (_m *VectorStore) AddDocuments(ctx context.Context, indexName string, documents []document.Document) ([]error, error) {
	ret := _m.Called(ctx, indexName, documents)

	var r0 []error
	var r1 error
	if 