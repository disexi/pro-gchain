package weaviateVS

import (
	"context"
	"errors"

	"github.com/weaviate/weaviate-go-client/v4/weaviate"
	"github.com/weaviate/weaviate-go-client/v4/weaviate/auth"
	"github.com/weaviate/weaviate-go-client/v4/weaviate/graphql"
	"github.com/weaviate/weaviate/entities/models"
	"github.com/wejick/gchain/datastore"
	"github.com/wejick/gchain/document"
	"github.com/wejick/gchain/model"
)

var _ datastore.VectorStore = &WeaviateVectorStore{}

// WeaviateVectorStore provide access to weaviate vector db
type WeaviateVectorStore struct {
	client         *weaviate.Client
	embeddingModel model.EmbeddingModel

	existClass map[string]bool
}

// NewWeaviateVectorStore return new Weaviate Vector Store instance
// headers is optional, if you want to add additional headers to the request
func NewWeaviateVectorStore(host string, scheme string, apiKey string, embeddingModel model.EmbeddingModel, headers map[string]string) (WVS *WeaviateVectorStore, err error) {
	WVS = &WeaviateVectorStore{
		existClass:     map[string]bool{},
		embeddingModel: embeddingModel,
	}
	cfg := weaviate.Config{
		Host:       host,
		Scheme:     scheme,
		Headers:    headers,
		AuthConfig: auth.ApiKey{Value: apiKey},
	}
	WVS.client, err = weaviate.NewClient(cfg)

	return
}

// SearchVector query weaviate using vector
// for weaviate support to return additional field / metadata is not yet implemented,
func (W *WeaviateVectorStore) SearchVector(ctx context.Context, className string, vector []float32, options ...func(*datastore.Option)) (output []document.Document, err error) {
	opts := datastore.Option{}
	for _, opt := range options {
		opt(&opts)
	}

	if opts.Similarity == 0 {
		opts.Similarity = 0.8
	}

	query := W.client.GraphQL().NearVectorArgBuilder().WithVector(vector).WithCertainty(opts.Similarity)
	fields := []graphql.Field{
		{Name: "text"},
	}
	// add additional fields
	for _, fieldName := range opts.AdditionalFields {
		fields = append(fields, graphql.Field{
			Name: fieldName,
		})
	}
	resp, err := W.client.GraphQL().Get().WithClassName(className).WithNearVector(query).WithFields(fields...).WithLimit(5).Do(ctx)
	if err != nil {
		return
	}

	output, err = objectsToDocument(className, resp.Data["Get"], opts.AdditionalFields)

	return
}

// Search query weaviate db, the query parameter will be translated into embedding
// the underlying query is the same with SearchVector
func (W *WeaviateVectorStore) Search(ctx context.Context, className string, query string, options ...func(*datastore.Option)) (output []document.Document, err error) {
	vectorQuery, err := W.embeddingModel.EmbedQuery(query)
	if err != nil {
		return
	}

	output, err = W.SearchVector(ctx, className, vectorQuery)

	return
}

// AddText add single string document
func (W *WeaviateVectorStore) AddText(ctx context.Context, className string, input string) (err error) {
	_, err = W.AddDocuments(ctx, className, []document.Document{{Text: input}})
	return
}

// AddDocuments add multiple string documents
func (W *WeaviateVectorStore) AddDocuments(ctx context.Context, className string, documents []document.Document) (batchErr []error, err error) {
	err = W.createClassIfNotExist(ctx, className)
	if err != nil {
		return
	}

	objVectors, err := W.embeddingModel.EmbedDocuments(document.DocumentsToStrings(documents))
	if err != nil {
		return
	}
	objs := documentsToObject(className, documents, objVectors)
	batchResp, err := W.client.Batch().ObjectsBatcher().WithObjects(objs...).Do(ctx)
	if err != nil {
		return
	}
	for _, res := range batchResp {
		if res.Result.Errors != nil {
			batchErr = append(batchErr, errors