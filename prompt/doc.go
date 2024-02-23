/*
	A “prompt” refers to the input to the model.

This input is rarely hard coded, but rather is often constructed from multiple components.
A PromptTemplate is responsible for the construction of this input.

PromptTemplate is powered by Go text/template, however we provided simplistic interface to interact with.
Everything is a string and data passed along to the template as map[string]string.

Example :

	template, _ := NewPromptT