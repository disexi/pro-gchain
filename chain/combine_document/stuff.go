package combine_document

import (
	"context"
	"errors"

	"github.com/wejick/gchain/chain"
	"github.com/wejick/gchain/chain/llm_chain"
	"github.com/wejick/gchain/model"
	"github.com/wejick/gchain/prompt"
)

var _ CombinedDocument = &StuffCombineDocument{}
var _ chain.BaseChain = &StuffCombineDocument{}

// StuffCombineDocument chain to feed text document to LLM with specified prompt
type StuffCombineDocument struct {
	prompt            *prompt.PromptTemplate
	llmChain          *llm_chain.LLMChain
	promptTemplateKey string
}

// NewStuffCombineDocument creates new instance of StuffCombineDocument
func NewStuffCombineDocument(prompt *prompt.PromptTemplate,
	templateKey string, llmChain *llm_chain.LLMChain) *StuffCombineDocument {
	return &StuffCombineDocument{
		prompt:            prompt,
		llmChain:          llmChain,
		promptTemplateKey: templateKey,
	}
}

// Combine concatenate the given document and then feed to LLM
func (S *StuffCombineDocument) Combine(ctx context.Context, docs []string, options ...func(*model.Option)) (outp