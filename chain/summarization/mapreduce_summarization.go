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
"{{.text}}""
CONCISE SUMMARY:
	`
)

type MapReduceSummarizationChain struct {
	mapReduceCombineDocument *combine_document.MapReduceCombineDocument
}

var _ chain.BaseChain = &MapReduceSummarizationChain{}

// NewMapReduceSummarizationChain create new map reduce summarization chain instance
// put empty "" string to use default prompt
// put 0 to use default maxToken
func NewMapReduceSummarizationChain(llmChain *llm_chain.LLMChain, mapPromptString string, reducePromptString string,
	promptTemplateKey string,
	splitter textsplitter.TextSplitter, maxToken int) (m *MapReduceSummarizationChain, err error) {

	var promptTemplateMap, promptTemplateReduce *prompt.PromptTemplate

	if mapPromptString == "" {
		promptTemplateMap, err = prompt.NewPromptTemplate("map", promptSummarizeMapReduce)
		if err != nil {
			return
		}
		promptTemplateKey = "text"
	}

	if reducePromptString == "" {
		promptTemplateReduce, err = prom