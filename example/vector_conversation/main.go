
package main

import (
	"bufio"
	"context"
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/sashabaranov/go-openai"
	"github.com/wejick/gchain/callback"
	"github.com/wejick/gchain/chain/conversational_retrieval"
	weaviateVS "github.com/wejick/gchain/datastore/weaviate_vector"
	"github.com/wejick/gchain/document"
	"github.com/wejick/gchain/model"
	_openai "github.com/wejick/gchain/model/openAI"
	"github.com/wejick/gchain/textsplitter"
)

var OAIauthToken = os.Getenv("OPENAI_API_KEY")
var chatModel *_openai.OpenAIChatModel
var embeddingModel *_openai.OpenAIEmbedModel
var wvClient *weaviateVS.WeaviateVectorStore
var textplitter *textsplitter.TikTokenSplitter

const (
	wvhost   = "question-testing-twjfnqyp.weaviate.network"
	wvscheme = "https"
	wvApiKey = ""
)

type source struct {
	filename string
	url      string
	doc      string
}

func Init() (err error) {
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	textplitter, err = textsplitter.NewTikTokenSplitter(_openai.GPT3Dot5Turbo0301)
	if err != nil {
		log.Println(err)