package llm_chain

import (
	"context"

	"github.com/wejick/gchain/callback"
	"github.com/wejick/gchain/chain"
	"github.com/wejick/gchain/model"
	"github.com/wejick/gchain/prompt"
)

var _ chain.BaseChain = &LLMChain{}

const defaultTemplate = `{{.input}}`

type LLMChain struct {
	llmModel        model.LLMModel
	callbackManager *callback.Manager
	promptTemplate  *prompt.PromptTemplate
}

// NewLLMChain create an LLMChain instanc