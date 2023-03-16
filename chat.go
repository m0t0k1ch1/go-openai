package openai

import (
	"context"
	"net/http"
)

type ChatCompletionMessage struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

type ChatCompletionChoice struct {
	Message      ChatCompletionMessage `json:"message"`
	FinishReason string                `json:"finish_reason"`
	Index        int                   `json:"index"`
}

type ChatCompletionsUsage struct {
	PromptTokens     int `json:"prompt_tokens"`
	CompletionTokens int `json:"completion_tokens"`
	TotalTokens      int `json:"total_tokens"`
}

type CreateChatCompletionsRequest struct {
	Model            string                  `json:"model"`
	Messages         []ChatCompletionMessage `json:"messages"`
	Temperature      float32                 `json:"temperature"`
	TopP             float32                 `json:"top_p"`
	N                int                     `json:"n"`
	Stream           bool                    `json:"stream"`
	Stop             []string                `json:"stop"`
	MaxTokens        int                     `json:"max_tokens,omitempty"`
	PresencePenalty  float32                 `json:"presence_penalty"`
	FrequencyPenalty float32                 `json:"frequency_penalty"`
	LogitBias        map[string]int          `json:"logit_bias,omitempty"`
	User             string                  `json:"user,omitempty"`
}

func NewDefaultCreateChatCompletionsRequest() CreateChatCompletionsRequest {
	return CreateChatCompletionsRequest{
		Temperature:      1,
		TopP:             1,
		N:                1,
		Stream:           false,
		Stop:             nil,
		PresencePenalty:  0,
		FrequencyPenalty: 0,
		LogitBias:        nil,
	}
}

type CreateChatCompletionsResponse struct {
	ID      string                 `json:"id"`
	Object  string                 `json:"object"`
	Created int64                  `json:"created"`
	Model   string                 `json:"model"`
	Usage   ChatCompletionsUsage   `json:"usage"`
	Choices []ChatCompletionChoice `json:"choices"`
}

func (c *Client) CreateChatCompletion(ctx context.Context, req CreateChatCompletionsRequest) (CreateChatCompletionsResponse, error) {
	var resp CreateChatCompletionsResponse

	if err := c.doAPI(ctx, http.MethodPost, "/v1/chat/completions", req, &resp); err != nil {
		return CreateChatCompletionsResponse{}, err
	}

	return resp, nil
}
