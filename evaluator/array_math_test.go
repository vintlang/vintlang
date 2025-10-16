package evaluator

import (
	"testing"

	"github.com/vintlang/vintlang/lexer"
	"github.com/vintlang/vintlang/object"
	"github.com/vintlang/vintlang/parser"
)

func TestArrayMathMethods(t *testing.T) {
	tests := []struct {
		input    string
		expected any
	}{
		// Test sum
		{"[1, 2, 3, 4, 5].sum()", 15},
		{"[].sum()", 0},

		// Test average
		{"[1, 2, 3, 4, 5].average()", 3.0},
		{"[2, 4, 6].average()", 4.0},

		// Test min/max
		{"[5, 1, 9, 3].min()", 1},
		{"[5, 1, 9, 3].max()", 9},

		// Test product
		{"[2, 3, 4].product()", 24},
		{"[].product()", 1},
	}

	for _, tt := range tests {
		l := lexer.New(tt.input)
		p := parser.New(l)
		program := p.ParseProgram()
		env := object.NewEnvironment()
		evaluated := Eval(program, env)

		switch expected := tt.expected.(type) {
		case int:
			testIntegerObject(t, evaluated, int64(expected))
		case float64:
			testFloatObject(t, evaluated, expected)
		}
	}
}

func TestArraySortMethods(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		// Test sortAsc
		{"let arr = [3, 1, 4, 1, 5]; arr.sortAsc(); arr", "[1, 1, 3, 4, 5]"},

		// Test sortDesc
		{"let arr = [3, 1, 4, 1, 5]; arr.sortDesc(); arr", "[5, 4, 3, 1, 1]"},

		// Test with strings
		{"let arr = [\"banana\", \"apple\", \"cherry\"]; arr.sort(); arr", "[apple, banana, cherry]"},
		{"let arr = [\"banana\", \"apple\", \"cherry\"]; arr.sortDesc(); arr", "[cherry, banana, apple]"},

		// Test with floats
		{"let arr = [3.14, 1.41, 2.71]; arr.sort(); arr", "[1.41, 2.71, 3.14]"},
		{"let arr = [3.14, 1.41, 2.71]; arr.sortDesc(); arr", "[3.14, 2.71, 1.41]"},
	}

	for _, tt := range tests {
		l := lexer.New(tt.input)
		p := parser.New(l)
		program := p.ParseProgram()
		env := object.NewEnvironment()
		evaluated := Eval(program, env)

		if evaluated.Inspect() != tt.expected {
			t.Errorf("input: %q - expected=%s, got=%s", tt.input, tt.expected, evaluated.Inspect())
		}
	}
}

func TestArraySortBy(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		// Test sortBy with simple function
		{"let arr = [3, 1, 4, 1, 5]; arr.sortBy(func(x){ return x }); arr", "[1, 1, 3, 4, 5]"},

		// Test sortBy with reverse function
		{"let arr = [3, 1, 4, 1, 5]; arr.sortBy(func(x){ return -x }); arr", "[5, 4, 3, 1, 1]"},
	}

	for _, tt := range tests {
		l := lexer.New(tt.input)
		p := parser.New(l)
		program := p.ParseProgram()
		env := object.NewEnvironment()
		evaluated := Eval(program, env)

		if evaluated.Inspect() != tt.expected {
			t.Errorf("input: %q - expected=%s, got=%s", tt.input, tt.expected, evaluated.Inspect())
		}
	}
}

func TestArrayMode(t *testing.T) {
	input := "let arr = [1, 2, 2, 3, 2, 4]; arr.mode()"
	expected := "[2]"

	l := lexer.New(input)
	p := parser.New(l)
	program := p.ParseProgram()
	env := object.NewEnvironment()
	evaluated := Eval(program, env)

	if evaluated.Inspect() != expected {
		t.Errorf("input: %q - expected=%s, got=%s", input, expected, evaluated.Inspect())
	}
}

func testFloatObject(t *testing.T, obj object.VintObject, expected float64) bool {
	result, ok := obj.(*object.Float)
	if !ok {
		t.Errorf("object is not Float. got=%T (%+v)", obj, obj)
		return false
	}
	if result.Value != expected {
		t.Errorf("object has wrong value. got=%f, want=%f", result.Value, expected)
		return false
	}
	return true
}
