package object

import (
	"testing"
)

func TestStringMethods(t *testing.T) {
	tests := []struct {
		method   string
		str      string
		args     []VintObject
		expected string
		isError  bool
	}{
		// startsWith tests
		{"startsWith", "hello world", []VintObject{&String{Value: "hello"}}, "true", false},
		{"startsWith", "hello world", []VintObject{&String{Value: "world"}}, "false", false},
		
		// endsWith tests
		{"endsWith", "hello world", []VintObject{&String{Value: "world"}}, "true", false},
		{"endsWith", "hello world", []VintObject{&String{Value: "hello"}}, "false", false},
		
		// includes tests
		{"includes", "hello world", []VintObject{&String{Value: "lo wo"}}, "true", false},
		{"includes", "hello world", []VintObject{&String{Value: "xyz"}}, "false", false},
		
		// repeat tests
		{"repeat", "ha", []VintObject{&Integer{Value: 3}}, "hahaha", false},
		{"repeat", "x", []VintObject{&Integer{Value: 0}}, "", false},
		
		// capitalize tests
		{"capitalize", "hello", []VintObject{}, "Hello", false},
		{"capitalize", "", []VintObject{}, "", false},
		
		// isNumeric tests
		{"isNumeric", "123", []VintObject{}, "true", false},
		{"isNumeric", "123.45", []VintObject{}, "true", false},
		{"isNumeric", "hello", []VintObject{}, "false", false},
		{"isNumeric", "", []VintObject{}, "false", false},
		
		// isAlpha tests
		{"isAlpha", "hello", []VintObject{}, "true", false},
		{"isAlpha", "Hello", []VintObject{}, "true", false},
		{"isAlpha", "hello123", []VintObject{}, "false", false},
		{"isAlpha", "", []VintObject{}, "false", false},
		
		// compareIgnoreCase tests
		{"compareIgnoreCase", "Hello", []VintObject{&String{Value: "hello"}}, "0", false},
		{"compareIgnoreCase", "apple", []VintObject{&String{Value: "banana"}}, "-1", false},
		
		// format tests
		{"format", "Hello {0}, you are {1} years old", []VintObject{&String{Value: "John"}, &Integer{Value: 25}}, "Hello John, you are 25 years old", false},
		
		// removeAccents tests
		{"removeAccents", "café", []VintObject{}, "cafe", false},
		{"removeAccents", "naïve", []VintObject{}, "naive", false},
	}

	for _, test := range tests {
		str := &String{Value: test.str}
		result := str.Method(test.method, test.args)
		
		if test.isError {
			if _, ok := result.(*Error); !ok {
				t.Errorf("Expected error for %s('%s'), got %s", test.method, test.str, result.Inspect())
			}
		} else {
			if result.Inspect() != test.expected {
				t.Errorf("Expected %s('%s') = %s, got %s", test.method, test.str, test.expected, result.Inspect())
			}
		}
	}
}

func TestStringToInt(t *testing.T) {
	tests := []struct {
		str      string
		expected string
		isError  bool
	}{
		{"123", "123", false},
		{"0", "0", false},
		{"-456", "-456", false},
		{"abc", "", true},
		{"", "", true},
	}

	for _, test := range tests {
		str := &String{Value: test.str}
		result := str.toInt([]VintObject{})
		
		if test.isError {
			if _, ok := result.(*Error); !ok {
				t.Errorf("Expected error for toInt('%s'), got %s", test.str, result.Inspect())
			}
		} else {
			if result.Inspect() != test.expected {
				t.Errorf("Expected toInt('%s') = %s, got %s", test.str, test.expected, result.Inspect())
			}
		}
	}
}

func TestStringRepeatEdgeCases(t *testing.T) {
	str := &String{Value: "test"}
	
	// Test negative count
	result := str.repeat([]VintObject{&Integer{Value: -1}})
	if _, ok := result.(*Error); !ok {
		t.Errorf("Expected error for negative repeat count")
	}
	
	// Test very large count
	result = str.repeat([]VintObject{&Integer{Value: 2000000}})
	if _, ok := result.(*Error); !ok {
		t.Errorf("Expected error for too large repeat count")
	}
}