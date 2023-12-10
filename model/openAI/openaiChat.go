package _openai

import (
	"context"
	"errors"
	"io"
	"log"

	goopenai "github.com/sashabaranov/go-openai"
	"github.com/wejick/gchain/callback"
	model "github.com/wejick/gchain/model"
)

var _ model.LLMModel = &OpenAIChatModel{}

type OpenAIChatModel struct {
	c               *goopenai.Client
	modelName       string
	callbackManager *callback.Manager
}

// NewOpenAIChatModel return new openAI Model instance
func NewOpenAIChatModel(authToken string, modelName string, callbackManager *callback.Manager, options ...func(*OpenAIOption)) (llm *OpenAIChatModel) {
	opts := OpenAIOption{}
	for _, opt := range options {
		opt(&opts)
	}

	clientConfig := newOpenAIClientConfig(authToken, opts)
	client := goopenai.NewClientWithConfig(clientConfig)

	llm = &OpenAIChatModel{
		c:               client,
		modelName:       modelName,
		callbackManager: callbackManager,
	}

	if opts.Verbose {
		llm.callbackManager.RegisterCallback(model.CallbackModelEnd, callback.VerboseCallback)
	}

	return
}

// Call runs completion on chat model, the prompt will be put as user chat
func (O *OpenAIChatModel) Call(ctx context.Context, prompt string, options ...func(*model.Option)) (output string, err error) {
	messages := []model.ChatMessage{
		{Role: model.ChatMessageRoleUser, Content: prompt},
	}
	responds, err := O.Chat(ctx, messages, options...)
	if err != nil {
		return
	