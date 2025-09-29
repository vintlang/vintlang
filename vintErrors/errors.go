package vinterrors
package vintErrors

import (
	"fmt"
	"strings"
)

// ErrorCode represents different types of errors
type ErrorCode string

const (
	// Lexer errors (E1xx)
	E100_ILLEGAL_CHAR     ErrorCode = "E100"
	E101_UNTERMINATED_STR ErrorCode = "E101" 
	E102_INVALID_ESCAPE   ErrorCode = "E102"
	
	// Parser errors (E2xx)  
	E200_UNEXPECTED_TOKEN ErrorCode = "E200"
	E201_MISSING_TOKEN    ErrorCode = "E201"
	E202_INVALID_SYNTAX   ErrorCode = "E202"
	
	// Semantic errors (E3xx)
	E300_UNDECLARED_VAR   ErrorCode = "E300"
	E301_TYPE_MISMATCH    ErrorCode = "E301" 
	E302_INVALID_OP       ErrorCode = "E302"
	
	// Runtime errors (E4xx)
	E400_INDEX_OUT_BOUNDS ErrorCode = "E400"
	E401_INVALID_ARG      ErrorCode = "E401"
	E402_NULL_REFERENCE   ErrorCode = "E402"
)

// Severity levels
type Severity string

const (
	ERROR   Severity = "ERROR"
	WARNING Severity = "WARNING" 
	INFO    Severity = "INFO"
)

// VintError represents a structured error with metadata
type VintError struct {
	Code     ErrorCode
	Severity Severity
	Message  string
	Line     int
	Column   int
	Source   string
	Context  string
	Suggestion string
}

// Error implements the error interface
func (e *VintError) Error() string {
	var builder strings.Builder
	
	// Write severity and code
	builder.WriteString(fmt.Sprintf("[%s %s] ", e.Severity, e.Code))
	
	// Write position
	if e.Line > 0 {
		if e.Column > 0 {
			builder.WriteString(fmt.Sprintf("Line %d:%d: ", e.Line, e.Column))
		} else {
			builder.WriteString(fmt.Sprintf("Line %d: ", e.Line))
		}
	}
	
	// Write message
	builder.WriteString(e.Message)
	
	// Add suggestion if available
	if e.Suggestion != "" {
		builder.WriteString(fmt.Sprintf(" %s", e.Suggestion))
	}
	
	// Add context if available
	if e.Context != "" {
		builder.WriteString(fmt.Sprintf("\n    %s", e.Context))
		if e.Column > 0 {
			builder.WriteString(fmt.Sprintf("\n    %s^", strings.Repeat(" ", e.Column-1)))
		}
	}
	
	return builder.String()
}

// NewError creates a new VintError
func NewError(code ErrorCode, severity Severity, message string, line, column int) *VintError {
	return &VintError{
		Code:     code,
		Severity: severity,
		Message:  message,
		Line:     line,
		Column:   column,
	}
}

// WithContext adds source code context to the error
func (e *VintError) WithContext(context string) *VintError {
	e.Context = context
	return e
}

// WithSuggestion adds a suggestion to the error
func (e *VintError) WithSuggestion(suggestion string) *VintError {
	e.Suggestion = suggestion
	return e
}