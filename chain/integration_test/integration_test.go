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
	chain, err := llm_chain.NewLLMChain(llmModel, callback.NewManager(), nil, false)
	assert.NoError(t, err, "NewLLMChain")
	outputMap, err := chain.Run(context.Background(), map[string]string{"input": "Indonesia Capital is Jakarta\nJakarta is the capital of "})
	assert.NoError(t, err, "error Run")
	assert.Contains(t, outputMap["output"], "Indonesia", "unexpected result")

	customPrompt, err := prompt.NewPromptTemplate("customPrompt", "{{.text}}")
	customPromptChain, err := llm_chain.NewLLMChain(llmModel, callback.NewManager(), customPrompt, false)
	assert.NoError(t, err, "NewLLMChain")

	customOutputMap, err := customPromptChain.Run(context.Background(), map[string]string{"text": "Indonesia Capital is Jakarta\nJakarta is the capital of "})
	assert.NoError(t, err, "error Run")
	assert.Contains(t, customOutputMap["output"], "Indonesia", "unexpected result")

}

func TestStuffSummarizationChain(t *testing.T) {
	llmchain, err := llm_chain.NewLLMChain(llmModel, callback.NewManager(), nil, false)
	assert.NoError(t, err, "NewLLMChain")

	chain, err := summarization.NewStuffSummarizationChain(llmchain, "", "text")
	assert.NoError(t, err, "error NewStuffSummarizationChain")
	output, err := chain.SimpleRun(context.Background(), `Modular audio and video hardware for retro machines like the Commodore 64. Designed to use 74 series TTL through hole ICs available back in the 1980s, something you can solder at home from parts or order ready assembled.
	One of the most recent videos shows a "Shadow of the Beast" demonstration, to show parallax scrolling with precisely timed raster effects. Please do consider subscribing to the YouTube channel if you want to see more updates to this project: 
	This project started when old retro arcade hardware was being discussed. In the back of my mind was the often fabled "Mega games" by Imagine Software whi