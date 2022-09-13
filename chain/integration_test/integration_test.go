//go:build integration
// +build integration

package integration_test

import (
	"context"
	"fmt"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/wejick/gchain/callback"
	"github.com/wejick/gchain/chain/conversation"
	"github.com/wejick/gchain/chain/conversational_retrieval"
	"github.com/wejick/gchain/chain/llm_chain"
	"github.com/wejick/gchain/chain/summarization"
	wikipedia "gi