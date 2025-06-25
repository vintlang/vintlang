package main

import (
	"fmt"

	"github.com/vintlang/vintlang/compiler"
	"github.com/vintlang/vintlang/lexer"
	"github.com/vintlang/vintlang/parser"
	"github.com/vintlang/vintlang/vm"
)

func main() {
	// --- Feel free to change the input string to test different expressions! ---
	input := `(50 / 2 * 2 + 10 - 5) > 20`

	l := lexer.New(input)
	p := parser.New(l)
	program := p.ParseProgram()

	comp := compiler.New()
	err := comp.Compile(program)
	if err != nil {
		fmt.Printf("compiler error: %s\n", err)
		return
	}

	machine := vm.New(comp.Bytecode())
	err = machine.Run()
	if err != nil {
		fmt.Printf("vm error: %s\n", err)
		return
	}

	// The result of the last expression is left on the stack.
	result := machine.LastPoppedStackElem()
	if result != nil {
		fmt.Printf("Result of '%s': %s\n", input, result.Inspect())
	}
}