package openai

import (
	"net/http"
	"os"
)

var (
	testClient = NewClient(http.DefaultClient, Config{
		APIKey: os.Getenv("OPENAI_API_KEY"),
	})
)
