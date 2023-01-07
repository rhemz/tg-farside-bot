package handlers

const (
	JsonIndent = "  "
)

type (
	Response struct {
		Message string `json:"message,omitempty"`
		Data    any    `json:"data,omitempty"`
	}

	handler struct{}
)
