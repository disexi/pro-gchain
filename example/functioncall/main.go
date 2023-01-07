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
	functionList := map[string