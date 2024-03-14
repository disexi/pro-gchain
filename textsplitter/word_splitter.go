package textsplitter

import (
	"strings"

	"github.com/wejick/gchain/document"
)

type WordSplitter struct {
}

// splitIntoBatches creates word batches where length's doesn't exceed maxChunkSize.
func (W *WordSplitter) Sp