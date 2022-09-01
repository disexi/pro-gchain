package conversation

import (
	"context"
	"errors"

	"github.com/wejick/gchain/callback"
	basechain "github.com/wejick/gchain/chain"
	"github.com/wejick/gchain/model"
)

// ConversationChain that carries on a conversation from a prompt plus history.
type ConversationChain struct {
	chatModel       model.ChatModel // only allow using ChatModel
	memory          []model.ChatMessage
	callbackManager *callback.Manager
}

// NewConversationChain create new conversation chain
// firstSystemPrompt will be the first chat in chat memory, with role as System
func NewConversationChain(chatModel model.ChatModel,
	memory []model.ChatMessage,
	callbackManager *callback.Manager,
	firstSystemPrompt string,
	verbose bool) (chain *ConversationChain) {

	if verbose {
		callbackManager.RegisterCallback(basechain.CallbackChainEnd, callback.VerboseCallback)
	}
	memory = append(memory, model.ChatMessage{Role: model.ChatMessageRoleSystem, Content: firstSystemPrompt})
	return &ConversationChain{
		chatModel:       chatModel,
		callbackManager: callbackManager,
		memory:          memory,
	}
}

// AppendMemory to add conversation to the memory
func (C *ConversationChain) AppendToMemory(message model.ChatMessage) {
	C.memory = append(C.memory, message)
}

// Run expect chat["input"] as input, and put the result to output["output"]
func (C *ConversationChain) Run(ctx context.Context, chat map[string]string, options ...func(*model.Option)) (output map[string]string, err error) {
	if _, ok := chat["input"]; !ok {
		return output, errors.New("input[\"input\"] is not specified")
	}
	output = make(map[string]string)

	//trigger callback
	C.callbackManager.TriggerEvent(ctx, basechain.Ca