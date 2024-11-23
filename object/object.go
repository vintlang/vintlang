package object

import (
	"fmt"
)

// ObjectType represents various types of objects
type ObjectType string

// Constants for object types
const (
	INTEGER_OBJ      = "INTEGER"
	FLOAT_OBJ        = "FLOAT"
	BOOLEAN_OBJ      = "BOOLEAN"
	NULL_OBJ         = "NULL"
	RETURN_VALUE_OBJ = "RETURN_VALUE"
	ERROR_OBJ        = "ERROR"
	FUNCTION_OBJ     = "FUNCTION"
	STRING_OBJ       = "STRING"
	BUILTIN_OBJ      = "BUILTIN"
	ARRAY_OBJ        = "ARRAY"
	DICT_OBJ         = "DICT"
	CONTINUE_OBJ     = "CONTINUE"
	BREAK_OBJ        = "BREAK"
	FILE_OBJ         = "FILE"
	TIME_OBJ         = "TIME"
	JSON_OBJ         = "JSON"
	MODULE_OBJ       = "MODULE"
	BYTE_OBJ         = "BYTE"
	PACKAGE_OBJ      = "PACKAGE"
	INSTANCE         = "INSTANCE"
	AT               = "@"
)

// Object interface represents any object in the system
type Object interface {
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
	Next() (Object, Object)
	Reset()
}

// newError creates a formatted error object with red text
func newError(format string, a ...interface{}) *Error {
	redFormat := fmt.Sprintf("\x1b[%dm%s\x1b[0m", 31, format) // Red-colored error message
	return &Error{Message: fmt.Sprintf(redFormat, a...)}
}
