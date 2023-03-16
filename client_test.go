package openai

import (
	"context"
	"net/http"
	"os"
	"testing"

	"github.com/google/go-cmp/cmp"
)

var (
	testClient = NewClient(http.DefaultClient, Config{
		APIKey: os.Getenv("OPENAI_API_KEY"),
	})
)

func TestCreateChatCompletions(t *testing.T) {
	req := NewDefaultCreateChatCompletionsRequest()
	req.Model = "gpt-3.5-turbo-0301"
	req.Messages = []ChatCompletionMessage{{
		Role:    "user",
		Content: "Hello!",
	}}

	resp, err := testClient.CreateChatCompletion(context.Background(), req)
	if err != nil {
		t.Fatal(err)
	}
	if diff := cmp.Diff("chat.completion", resp.Object); len(diff) > 0 {
		t.Errorf("mismatch:\n%s", diff)
	}
}
