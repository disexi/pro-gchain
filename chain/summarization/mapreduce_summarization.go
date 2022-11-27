package summarization

import (
	"context"
	"errors"

	"github.com/wejick/gchain/chain"
	"github.com/wejick/gchain/chain/combine_document"
	"github.com/wejick/gchain/chain/llm_chain"
	"github.com/wejick/gchain/model"
	"github.com/wejick/gchain/prompt"
	"github.com/wejick/gchain/textsplitter"
)

const (
	promptSummarizeMapReduce = `Write a concise summary of the following:
"{{.text}