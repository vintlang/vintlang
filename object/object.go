package object

import (
	"fmt"
	"strings"
)

// ObjectType represents various types of objects
type ObjectType string

// Constants for object types
const (
	INTEGER_OBJ       = "INTEGER"
	FLOAT_OBJ         = "FLOAT"
	BOOLEAN_OBJ       = "BOOLEAN"
	NULL_OBJ          = "NULL"
	RETURN_VALUE_OBJ  = "RETURN_VALUE"
	ERROR_OBJ         = "ERROR"
	FUNCTION_OBJ      = "FUNCTION"
	STRING_OBJ        = "STRING"
	BUILTIN_OBJ       = "BUILTIN"
	ARRAY_OBJ         = "ARRAY"
	DICT_OBJ          = "DICT"
	CONTINUE_OBJ      = "CONTINUE"
	BREAK_OBJ         = "BREAK"
	FILE_OBJ          = "FILE"
	TIME_OBJ          = "TIME"
	DURATION_OBJ      = "DURATION"
	JSON_OBJ          = "JSON"
	MODULE_OBJ        = "MODULE"
	BYTE_OBJ          = "BYTE"
	PACKAGE_OBJ       = "PACKAGE"
	INSTANCE          = "INSTANCE"
	NATIVE_OBJ        = "NATIVE_OBJ"
	POINTER_OBJ       = "POINTER"
	AT                = "@"
	DEFERRED_CALL_OBJ = "DEFERRED_CALL"

	// Async/Concurrency Objects
	PROMISE_OBJ    = "PROMISE"
	CHANNEL_OBJ    = "CHANNEL"
	ASYNC_FUNC_OBJ = "ASYNC_FUNCTION"

	// Custom Error Objects
	CUSTOM_ERROR_OBJ = "CUSTOM_ERROR"
	ERROR_TYPE_OBJ   = "ERROR_TYPE"

	// HTTP Objects
	HTTP_APP_OBJ      = "HTTP_APP"
	HTTP_REQUEST_OBJ  = "HTTP_REQUEST"
	HTTP_RESPONSE_OBJ = "HTTP_RESPONSE"
	UPLOADED_FILE_OBJ = "UPLOADED_FILE"
)

// VintObject interface represents any object in the system
type VintObject interface {
	Type() ObjectType
	Inspect() string
}

// HashKey is used for hashable objects like strings and integers
type HashKey struct {
	Type  ObjectType
	Value uint64
}

// Hashable interface is for objects that can be used as keys in a hash map
type Hashable interface {
	HashKey() HashKey
}

// Iterable interface supports iteration for collections like dicts, strings, and arrays
type Iterable interface {
	Next() (VintObject, VintObject)
	Reset()
}

// newError creates a formatted error object with red text
func newError(format string, a ...interface{}) *Error {
	redFormat := fmt.Sprintf("\x1b[%dm%s\x1b[0m", 31, format) // Red-colored error message
	return &Error{Message: fmt.Sprintf(redFormat, a...)}
}

// DeferredCall represents a function call that has been deferred

type DeferredCall struct {
	Fn   VintObject
	Args []VintObject
}

func (dc *DeferredCall) Type() ObjectType { return DEFERRED_CALL_OBJ }
func (dc *DeferredCall) Inspect() string {
	return "deferred call"
}

// ErrorType represents a custom error type definition
type ErrorType struct {
	Name       string
	Parameters []string
}

func (et *ErrorType) Type() ObjectType { return ERROR_TYPE_OBJ }
func (et *ErrorType) Inspect() string {
	return fmt.Sprintf("error type: %s(%s)", et.Name, strings.Join(et.Parameters, ", "))
}

// CustomError represents an instance of a custom error type
type CustomError struct {
	ErrorType *ErrorType
	Arguments []VintObject
}

func (ce *CustomError) Type() ObjectType { return CUSTOM_ERROR_OBJ }
func (ce *CustomError) Inspect() string {
	args := []string{}
	for _, arg := range ce.Arguments {
		args = append(args, arg.Inspect())
	}
	return fmt.Sprintf("\x1b[1;31m%s:\x1b[0m %s(%s)",
		ce.ErrorType.Name, ce.ErrorType.Name, strings.Join(args, ", "))
}
