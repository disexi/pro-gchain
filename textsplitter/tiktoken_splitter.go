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
func NewTikTokenSplitter(modelName string) (*TikTokenSplitter, error) {
	if modelName == "" {
		modelName = "gpt-3.5-turbo-0301"
	}

	tkm, err := tiktoken.EncodingForModel(modelName)
	return &TikTokenSplitter{
		tkm: tkm,
	}, err
}

// SplitText creates chunks where length's doesn't exceed maxChunkSize.
func (T *TikTokenSplitter) SplitText(input string, maxChunkSize int, overlap int) []string {
	if input == "" {
		return []string{}
	}
	batches := []string{}

	words := strings.Fields(inpu