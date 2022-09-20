//go:build integration
// +build integration

package integration_test

import (
	"context"
	"fmt"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/wejick/gchain/callback"
	"github.com/wejick/gchain/chain/conversation"
	"github.com/wejick/gchain/chain/conversational_retrieval"
	"github.com/wejick/gchain/chain/llm_chain"
	"github.com/wejick/gchain/chain/summarization"
	wikipedia "github.com/wejick/gchain/datastore/wikipedia_retriever"
	"github.com/wejick/gchain/eval"
	"github.com/wejick/gchain/model"
	_openai "github.com/wejick/gchain/model/openAI"
	"github.com/wejick/gchain/prompt"
	"github.com/wejick/gchain/textsplitter"
)

var llmModel model.LLMModel
var chatModel model.ChatModel

func TestMain(m *testing.M) {
	fmt.Println("Running integration tests...")
	// Perform any setup or initialization here

	var authToken = os.Getenv("OPENAI_API_KEY")
	llmModel = _openai.NewOpenAIModel(authToken, "text-davinci-003", callback.NewManager())

	chatModel = _openai.NewOpenAIChatModel(authToken, _openai.GPT3Dot5Turbo0301, callback.NewManager())

	exitCode := m.Run()

	// Perform any cleanup or teardown here

	// Exit with the appropriate exit code
	// (0 for success, non-zero for failure)
	os.Exit(exitCode)
}

func TestLlmChain(t *testing.T) {
	chain, err := llm_chain.NewLLMChain(llmModel, callback.N