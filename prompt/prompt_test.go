package prompt

import (
	"testing"
)

func TestPrompt_template_FormatPrompt(t *testing.T) {
	type args struct {
		Data map[stri