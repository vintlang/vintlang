package evaluator

import (
	"testing"

	"github.com/vintlang/vintlang/lexer"
	"github.com/vintlang/vintlang/object"
	"github.com/vintlang/vintlang/parser"
)

func TestIfExpressionAndStatement(t *testing.T) {
	cases := []struct {
		input    string
		expected object.Object
	}{
		// If as an expression
		{"let status = \"\"; status = if (true) { \"Online\" } else { \"Offline\" }; status", &object.String{Value: "Online"}},
		{"let status = \"\"; status = if (false) { \"Online\" } else { \"Offline\" }; status", &object.String{Value: "Offline"}},
		{"let status = \"\"; status = if (false) { \"Online\" }; status", &object.Null{}},
		// If as a statement (side effect)
		{"let x = 0; if (true) { x = 5 }; x", &object.Integer{Value: 5}},
		{"let x = 0; if (false) { x = 5 }; x", &object.Integer{Value: 0}},
	}

	for _, tt := range cases {
		l := lexer.New(tt.input)
		p := parser.New(l)
		program := p.ParseProgram()
		env := object.NewEnvironment()
		result := Eval(program, env)
		if result.Type() != tt.expected.Type() || result.Inspect() != tt.expected.Inspect() {
			t.Errorf("input: %q\nexpected: %s\ngot: %s", tt.input, tt.expected.Inspect(), result.Inspect())
		}
	}
}

func TestAsyncFunctions(t *testing.T) {
	tests := []struct {
		input    string
		expected interface{}
	}{
		{
			"let f = async func() { return 42 }; let p = f(); await p",
			42,
		},
		{
			"let f = async func(x) { return x + 1 }; let p = f(5); await p",
			6,
		},
		{
			"let f = async func() { return \"hello\" }; let p = f(); await p",
			"hello",
		},
	}

	for _, tt := range tests {
		l := lexer.New(tt.input)
		p := parser.New(l)
		program := p.ParseProgram()
		env := object.NewEnvironment()
		result := Eval(program, env)

		switch expected := tt.expected.(type) {
		case int:
			testIntegerObject(t, result, int64(expected))
		case string:
			testStringObject(t, result, expected)
		default:
			t.Errorf("unexpected expected type: %T", expected)
		}
	}
}

func TestChannels(t *testing.T) {
	tests := []struct {
		input    string
		expected interface{}
	}{
		{
			"let ch = chan; ch",
			"chan(unbuffered)",
		},
		{
			"let ch = chan(5); ch",
			"chan(buffered:5)",
		},
	}

	for _, tt := range tests {
		l := lexer.New(tt.input)
		p := parser.New(l)
		program := p.ParseProgram()
		env := object.NewEnvironment()
		result := Eval(program, env)

		if result.Inspect() != tt.expected {
			t.Errorf("wrong result. expected=%q, got=%q", tt.expected, result.Inspect())
		}
	}
}

func testIntegerObject(t *testing.T, obj object.Object, expected int64) bool {
	result, ok := obj.(*object.Integer)
	if !ok {
		t.Errorf("object is not Integer. got=%T (%+v)", obj, obj)
		return false
	}
	if result.Value != expected {
		t.Errorf("object has wrong value. got=%d, want=%d",
			result.Value, expected)
		return false
	}
	return true
}

func testStringObject(t *testing.T, obj object.Object, expected string) bool {
	result, ok := obj.(*object.String)
	if !ok {
		t.Errorf("object is not String. got=%T (%+v)", obj, obj)
		return false
	}
	if result.Value != expected {
		t.Errorf("object has wrong value. got=%q, want=%q",
			result.Value, expected)
		return false
	}
	return true
}
 