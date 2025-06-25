package compiler

import (
	"testing"

	"github.com/vintlang/vintlang/ast"
	"github.com/vintlang/vintlang/code"
	"github.com/vintlang/vintlang/lexer"
	"github.com/vintlang/vintlang/parser"
)

type compilerTestCase struct {
	input                string
	expectedConstants    []interface{}
	expectedInstructions []code.Instructions
}

func TestIntegerArithmetic(t *testing.T) {
	// We will add test cases here later
}

func parse(input string) *ast.Program {
	l := lexer.New(input)
	p := parser.New(l)
	return p.ParseProgram()
}
