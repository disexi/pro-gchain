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
	chatModel       model.ChatModel // only 