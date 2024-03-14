package textsplitter

import (
	"strings"

	"github.com/pkoukk/tiktoken-go"
	"github.com/wejick/gchain/document"
)

type TikTokenSplitter struct {
	tkm *tiktoken.Tiktoken
}

// NewTikTokenSplitter create new TikTokenSplitter instance
// if modelName empty, the default one is gpt-3.5-turbo-0301
func NewTikTokenSplit