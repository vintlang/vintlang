package vm

import (
	"testing"

	"github.com/vintlang/vintlang/ast"
	"github.com/vintlang/vintlang/lexer"
	"github.com/vintlang/vintlang/parser"
)

type vmTestCase struct {
	input    string
	expected interface{}
}

func TestIntegerArithmetic(t *testing.T) {
	// We will add test cases here later
}

func parse(input string) *ast.Program {
	l := lexer.New(input)
	p := parser.New(l)
	return p.ParseProgram()
}
