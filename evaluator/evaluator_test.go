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
 