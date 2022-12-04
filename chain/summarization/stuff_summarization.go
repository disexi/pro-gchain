package summarization

import (
	"context"
	"errors"

	"github.com/wejick/gchain/chain"
	"github.com/wejick/gchain/chain/combine_document"
	"github.com/wejick/gchain/chain/llm_chain"
	"github.com/wejick/gc