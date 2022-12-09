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

// WeaviateVe