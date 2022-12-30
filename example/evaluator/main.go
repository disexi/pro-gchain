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
	Result      b