package _openai

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewOpenAIClientConfig(t *testing.T) {
	authToken := "my-auth-token"
	opts := OpenAIOption{
		BaseURL:    "https://my-base-url.com",
		OrgID:      "my-org-id",
		APIVersion: "v2",
	}

	t.Run("DefaultConfig", func(t *testing.T) {
		clientConfig := newOpenAIClientConfig(authToken, OpenAIOption{})
		assert.Equal(t, "https://api.openai.com/v1", clientConfig.BaseURL)
