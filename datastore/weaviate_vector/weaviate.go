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
func NewWeaviateVectorStore(host string, scheme string, apiKey string, embeddingModel model.EmbeddingModel, hea