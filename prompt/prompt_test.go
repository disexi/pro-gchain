package prompt

import (
	"testing"
)

func TestPrompt_template_FormatPrompt(t *testing.T) {
	type args struct {
		Data map[string]string
	}
	tests := []struct {
		name              string
		templateName      string
		templateString    string
		promptTemplate    *PromptTemplate
		args              args
		wantOutput_prompt string
		wantErr           bool
	}{
		{
			name:           "empty",
			templateName:   "empty",
			templateString: "",
			wantErr:        false,
		},
		{
			name:           "{{.string}} {{.stringfloat}} {{.stringinteger}}",
			templateName:   "