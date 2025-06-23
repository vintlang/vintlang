package object

import (
	"fmt"
)

// Error represents a custom error type with an associated message.
type Error struct {
	Message string
}

func NewError(msg string) *Error {
	return &Error{
		msg,
	}
}

// Inspect returns a formatted string representation of the error.
func (e *Error) Inspect() string {
	return fmt.Sprintf("\x1b[1;31mError:\x1b[0m %s", e.Message)
}

// Type returns the object type of the error.
func (e *Error) Type() ObjectType {
	return ERROR_OBJ
}
