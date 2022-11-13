package llm_chain

import (
	"context"
	"reflect"
	"testing"

	"github.com/wejick/gchain/callback"
	"github.com/wejick/gchain/model"
	"github.com/wejick/gchain/prompt"
)

func TestLLMChain_SimpleRun(t *testing.T) {
	type fields struct {
		llmModel        model.LLMModel
		callbackManager *callback.Manager
	}
	type args struct {
		ctx  