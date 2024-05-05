
[![Go Reference](https://pkg.go.dev/badge/github.com/disexi/pro-gchain.svg)](https://pkg.go.dev/github.com/disexi/pro-gchain)
![Build workflow](https://github.com/disexi/pro-gchain/actions/workflows/go.yml/badge.svg)
[![Integration test](https://github.com/disexi/pro-gchain/actions/workflows/integration.yml/badge.svg)](https://github.com/disexi/pro-gchain/actions/workflows/integration.yml)


# 🤔 What is this?
Inspired by [langchain](https://github.com/hwchase17/langchain) and designed to provide enhanced composability when building Large Language Model Applications. The mission of Pro-Gchain is to bring the langchain concept to Go in an idiomatic way.

This Library is useful for many usecases, such as :

**❓ Question Answering over specific documents**
- [Building chatbot with vector databased backed knowledge base](https://wejick.wordpress.com/2023/06/18/building-llm-based-chatbot-with-a-knowledge-base-in-go/)

**💬 Chatbots**
- [Streaming Chatbot with Go and WebSocket](https://wejick.wordpress.com/2023/06/24/making-an-llm-based-streaming-chatbot-with-go-and-websocket/)

**📄 Document Summarization**


## Installation and Importing

```bash
$ go get github.com/disexi/pro-gchain
```

```golang
import "github.com/disexi/pro-gchain"
```


## Example
```golang
llmModel = _openai.NewOpenAIModel(authToken, "", "text-davinci-003",callback.NewManager(), true)
chain, err := llm_chain.NewLLMChain(llmModel, nil)
if err != nil {
    //handle error
}
outputMap, err := chain.Run(context.Background(), map[string]string{"input": "Indonesia Capital is Jakarta\nJakarta is the capital of "})
fmt.Println(outputMap["output"])
```
Explore more in the [example](./example/) folder

Since our documentation is not yet complete, please check out the examples and integration test for reference.

## Notice
1. Don't use it if you have a better option.
2. Pro-GChain's priority is going idiomatic. Therefore, despite borrowing many aspects from langchain, don't expect exactly the same behavior, as this is not an exact reimplementation.