package openai

type ErrorResponse struct {
	Error Error `json:"error"`
}

type CreateChatCompletionsResponse struct{}
