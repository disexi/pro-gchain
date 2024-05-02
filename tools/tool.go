package tools

import (
	"context"
	"fmt"

	"github.com/wejick/gchain/model"
)

// BaseTool is the interface for tool
// the idea is to keep it compatible with chain interface, so chain be used as tool as well
type BaseTool interface {
	// Run expect map string of string, each tool may expect different data
	Run(ctx context.Context, pro