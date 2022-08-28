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
		callbackManager.Reg