package model

import "github.com/zmzlois/LinkGoGo/pkg/utils"

var (
	// InvalidTokenError is returned when the session token is invalid
	InvalidTokenError = NewError(InvalidToken, "Invalid token")

	// InvalidBodyError is returned when we cannot decode the request body
	InvalidBodyError = NewError(InvalidBody, "Invalid body")

	// UnknownError is the error instance used as a fallback error
	UnknownError = NewError(Unknown, "Something went wrong, please try again later.")
)

// Error defines the error body
type Error struct {
	Code     ErrorCode   `json:"code"`
	Message  string      `json:"message"`
	Metadata interface{} `json:"metadata,omitempty"`
}

// NewError creates a new error
func NewError(code ErrorCode, message string, metadata ...interface{}) *Error {
	return &Error{
		Code:     code,
		Message:  message,
		Metadata: utils.FirstOf(metadata),
	}
}
