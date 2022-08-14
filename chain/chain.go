
package chain

import (
	"context"

	"github.com/wejick/gchain/model"
)

//go:generate mockery --name BaseChain
type BaseChain interface {
	// Run does prediction of input prompt of <string,string> and produce output of <string,string>
	// map of <string,string> of output and input prompt to accomodate many possible usecases
	Run(ctx context.Context, prompt map[string]string, options ...func(*model.Option)) (output map[string]string, err error)

	// SimpleRun does prediction of prompt of string and produce output of string
	// this is to accomodate simple prompt / output usage
	SimpleRun(ctx context.Context, prompt string, options ...func(*model.Option)) (output string, err error)
}

type DummyChain struct{}

func (D *DummyChain) Run(ctx context.Context, prompt map[string]string, options ...func(*model.Option)) (output map[string]string, err error) {
	return
}

func (D *DummyChain) SimpleRun(ctx context.Context, prompt string, options ...func(*model.Option)) (output string, err error) {
	return
}

const (
	CallbackChainStart = "chain_start"
	CallbackChainEnd   = "chain_end"
)