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
			templateName:   "tplt",
			templateString: "{{.string}} {{.stringfloat}} {{.stringinteger}}",
			args: args{
				Data: map[string]string{"string": "string", "stringfloat": "0.1", "stringinteger": "1"},
			},
			wantErr:           false,
			wantOutput_prompt: "string 0.1 1",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			P, err := NewPromptTemplate(tt.templateName, tt.templateString)
			if err != nil {
				t.Error