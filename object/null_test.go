package object

import (
	"testing"
)

func TestNullMethods(t *testing.T) {
	null := &Null{}

	// Test isNull
	result := null.isNull([]VintObject{})
	if result.Inspect() != "true" {
		t.Errorf("Expected null.isNull() = true, got %s", result.Inspect())
	}

	// Test toString
	result = null.toString([]VintObject{})
	if result.Inspect() != "null" {
		t.Errorf("Expected null.toString() = 'null', got %s", result.Inspect())
	}

	// Test equals with null
	result = null.equals([]VintObject{&Null{}})
	if result.Inspect() != "true" {
		t.Errorf("Expected null.equals(null) = true, got %s", result.Inspect())
	}

	// Test equals with non-null
	result = null.equals([]VintObject{&String{Value: "test"}})
	if result.Inspect() != "false" {
		t.Errorf("Expected null.equals('test') = false, got %s", result.Inspect())
	}
}

func TestNullCoalesce(t *testing.T) {
	null := &Null{}

	// Test coalesce with non-null values
	result := null.coalesce([]VintObject{&String{Value: "first"}, &String{Value: "second"}})
	if result.Inspect() != "first" {
		t.Errorf("Expected coalesce to return first non-null value 'first', got %s", result.Inspect())
	}

	// Test coalesce with all null values
	result = null.coalesce([]VintObject{&Null{}, &Null{}})
	if result.Type() != NULL_OBJ {
		t.Errorf("Expected coalesce with all nulls to return null, got %s", result.Type())
	}

	// Test coalesce with mixed values
	result = null.coalesce([]VintObject{&Null{}, &Integer{Value: 42}, &String{Value: "test"}})
	if result.Inspect() != "42" {
		t.Errorf("Expected coalesce to return first non-null value '42', got %s", result.Inspect())
	}
}

func TestNullIfNull(t *testing.T) {
	null := &Null{}

	// Test ifNull - should always return the provided value since this is null
	result := null.ifNull([]VintObject{&String{Value: "default"}})
	if result.Inspect() != "default" {
		t.Errorf("Expected ifNull to return 'default', got %s", result.Inspect())
	}

	result = null.ifNull([]VintObject{&Integer{Value: 0}})
	if result.Inspect() != "0" {
		t.Errorf("Expected ifNull to return '0', got %s", result.Inspect())
	}
}

func TestNullMethodErrors(t *testing.T) {
	null := &Null{}

	// Test isNull with wrong number of arguments
	result := null.isNull([]VintObject{&String{Value: "test"}})
	if _, ok := result.(*Error); !ok {
		t.Errorf("Expected error for isNull with arguments")
	}

	// Test toString with wrong number of arguments
	result = null.toString([]VintObject{&String{Value: "test"}})
	if _, ok := result.(*Error); !ok {
		t.Errorf("Expected error for toString with arguments")
	}

	// Test equals with wrong number of arguments
	result = null.equals([]VintObject{})
	if _, ok := result.(*Error); !ok {
		t.Errorf("Expected error for equals with no arguments")
	}

	result = null.equals([]VintObject{&String{Value: "test"}, &String{Value: "extra"}})
	if _, ok := result.(*Error); !ok {
		t.Errorf("Expected error for equals with too many arguments")
	}

	// Test ifNull with wrong number of arguments
	result = null.ifNull([]VintObject{})
	if _, ok := result.(*Error); !ok {
		t.Errorf("Expected error for ifNull with no arguments")
	}

	result = null.ifNull([]VintObject{&String{Value: "test"}, &String{Value: "extra"}})
	if _, ok := result.(*Error); !ok {
		t.Errorf("Expected error for ifNull with too many arguments")
	}

	// Test unsupported method
	result = null.Method("unsupported", []VintObject{})
	if _, ok := result.(*Error); !ok {
		t.Errorf("Expected error for unsupported method")
	}
}
