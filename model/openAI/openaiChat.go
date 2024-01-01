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
	} else {
		output = responds.Content
	}

	return
}

// Chat call chat completion
func (O *OpenAIChatModel) Chat(ctx context.Context, messages []model.ChatMessage, options ...func(*model.Option)) (output model.ChatMessage, err error) {
	opts := model.Option{}
	for _, opt := range options {
		opt(&opts)
	}

	// Trigger start callback
	flattenMessages := model.FlattenChatMessages(messages)
	O.callbackManager.TriggerEvent(ctx, model.CallbackModelStart, callback.CallbackData{
		EventName:    model.CallbackModelStart,
		FunctionName: "OpenAIChatModel.Chat",
		Input:        map[string]string{"input": flattenMessages},
		Output:       map[string]string{"output": output.String()},
	})

	// call chatStreaming if it's streaming chat
	if opts.IsStreaming && opts.StreamingChannel != nil {
		output, err = O.chatStreaming(ctx, messages, options...)
		return
	}

	RequestFunctions := []goopenai.FunctionDefin