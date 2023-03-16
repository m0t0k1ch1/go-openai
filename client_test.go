package openai

import (
	"net/http"
	"os"
)

var (
	testClient = NewClient(Config{
		APIKey: os.Getenv("OPENAI_API_KEY"),
	}, http.DefaultClient)
)
