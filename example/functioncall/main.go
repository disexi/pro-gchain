package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/wejick/gchain/callback"
	"github.com/wejick/gchain/model"
	_openai "github.com/wejick/gchain/model/openAI"
	"github.com/wejick/gchain/tools/greeting"
)

func main() {
	var authToken = os.Getenv("OPENAI_API_KEY")
	chatModel := _openai.NewOpenAIChatModel(authToken, _openai.GPT3Dot5Turbo0301, callback.NewManager())
	memory := []model.ChatMessage{}

	greeter := greeting.NewGreetingTool()

	// prepare a function register
	functionList := map[string]func(string) string{
		greeter.GetFunctionDefinition().Name: func(parameter string) string {
			greeting, err := greeter.SimpleRun(context.Background(), parameter)
			if err != nil {
				log.Println(err)
			}
			return greeting
		},
	}

	// The first call to the model, to see whether function call is needed
	memory = append(memory, model.ChatMessage{Role: model.ChatMessageRoleUser, Content: "Hi I'm Gio"})
	functionDef := greeter.Ge