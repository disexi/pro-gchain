package summarization

import (
	"context"
	"errors"

	"github.com/wejick/gchain/chain"
	"github.com/wejick/gchain/chain/combine_document"
	"github.com/wejick/gchain/chain/llm_chain"
	"github.com/wejick/gchain/model"
	"github.com/wejick/gchain/prompt"
)

const (
	promptSummarizeStuff = `Write a concise summary of the following:
"{{.text}}"
CONCISE SUMMARY:`
)

type StuffSummarizationChain struct {
	stuffCombineDocument *combine_document.StuffCombineDocument
}

var _ chain.BaseChain = &StuffSummarizationChain{}

func NewStuffSummarizationChain(llm_chain *llm_chain.LLMChain,
	promptTemplateString string, promptTemplateKey string) (s *StuffSummarizationChain, err 