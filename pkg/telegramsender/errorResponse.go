package telegramsender

import "fmt"

type ErrorResponse struct {
	Description string `json:"description"`
	ErrorCode   int    `json:"error_code"`
}

func (e *ErrorResponse) Error() string {
	return fmt.Sprintf("Telegram Error %d: %s", e.ErrorCode, e.Description)
}
