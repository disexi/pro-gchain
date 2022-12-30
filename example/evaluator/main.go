package main

/*
accept command flag -key to capture string
accept command flag -o to capture string
*/

import (
	"encoding/csv"
	"flag"
	"fmt"
	"io"
	"os"

	"github.com/wejick/gchain/callback"
	"github.com/wejick/gchain/eval"
	"github.com/wejick/gchain/model"
	_openai "github.com/wejick/gchain/model/openAI"
)

var (
	o string
)

func init() {
	flag.StringVar(&o, "o", "", "output file")
	flag.Parse()
}

type Test struct {
	Name        string
	Evaluator   string
	Input       string
	Expectation string
	Reason      string
	Result      bool
}

func main() {
	// open csv file from o
	// read csv file
	// for each row
	// put to array of Test
	tests, err := readCSV(o)
	if err != nil {
		fmt.Println(err)
		return
	}

	var authToken = os.Getenv("OPENAI_API_KEY")
	chatModel := _openai.NewOpenAIChatModel(authToken, _openai.GPT3Dot5Turbo0301, callback.NewManager())

	testRunner(tests, chatModel)
}

func testRunner(test []Test, llmModel model.LLMModel) {
	jsonEvaluator := eval.NewValidJson()
	var testResult