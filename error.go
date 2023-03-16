package openai

type ErrorResponse struct {
	Error Error `json:"error"`
}

type Error struct {
	Message string     `json:"message"`
	Type    string     `json:"type"`
	Param   NullString `json:"param"`
	Code    NullString `json:"code"`
}

func (e *Error) Error() string {
	return e.Message
}
