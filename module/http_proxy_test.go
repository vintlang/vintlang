package module

import (
	"testing"

	"github.com/vintlang/vintlang/object"
)

func TestConstantTimeCompare(t *testing.T) {
	tests := []struct {
		name     string
		a        string
		b        string
		expected bool
	}{
		{"equal strings", "secret123", "secret123", true},
		{"different strings", "secret123", "secret456", false},
		{"empty strings", "", "", true},
		{"one empty", "secret", "", false},
		{"different lengths", "short", "much longer string", false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			args := []object.VintObject{
				&object.String{Value: tt.a},
				&object.String{Value: tt.b},
			}

			result := constantTimeCompare(args, nil)
			boolResult, ok := result.(*object.Boolean)
			if !ok {
				t.Fatalf("Expected Boolean result, got %T: %s", result, result.Inspect())
			}
			if boolResult.Value != tt.expected {
				t.Errorf("constantTimeCompare(%q, %q) = %v, want %v",
					tt.a, tt.b, boolResult.Value, tt.expected)
			}
		})
	}
}

func TestConstantTimeCompareErrors(t *testing.T) {
	// Test wrong number of arguments
	result := constantTimeCompare([]object.VintObject{
		&object.String{Value: "only one"},
	}, nil)
	if _, ok := result.(*object.Error); !ok {
		t.Error("Expected error for wrong number of arguments")
	}

	// Test wrong argument types
	result = constantTimeCompare([]object.VintObject{
		&object.Integer{Value: 123},
		&object.String{Value: "string"},
	}, nil)
	if _, ok := result.(*object.Error); !ok {
		t.Error("Expected error for wrong argument types")
	}
}

func TestFetchRequestErrors(t *testing.T) {
	// Test missing URL
	result := fetchRequest(nil, map[string]object.VintObject{
		"method": &object.String{Value: "GET"},
	})
	if _, ok := result.(*object.Error); !ok {
		t.Error("Expected error when URL is missing")
	}

	// Test unknown parameter
	result = fetchRequest(nil, map[string]object.VintObject{
		"url":     &object.String{Value: "http://example.com"},
		"unknown": &object.String{Value: "value"},
	})
	if _, ok := result.(*object.Error); !ok {
		t.Error("Expected error for unknown parameter")
	}

	// Test wrong URL type
	result = fetchRequest(nil, map[string]object.VintObject{
		"url": &object.Integer{Value: 123},
	})
	if _, ok := result.(*object.Error); !ok {
		t.Error("Expected error for wrong URL type")
	}

	// Test wrong method type
	result = fetchRequest(nil, map[string]object.VintObject{
		"url":    &object.String{Value: "http://example.com"},
		"method": &object.Integer{Value: 123},
	})
	if _, ok := result.(*object.Error); !ok {
		t.Error("Expected error for wrong method type")
	}

	// Test wrong headers type
	result = fetchRequest(nil, map[string]object.VintObject{
		"url":     &object.String{Value: "http://example.com"},
		"headers": &object.String{Value: "not a dict"},
	})
	if _, ok := result.(*object.Error); !ok {
		t.Error("Expected error for wrong headers type")
	}
}
