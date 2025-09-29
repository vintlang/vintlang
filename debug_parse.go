package main

import (
	"fmt"

	"github.com/vintlang/vintlang/ast"
	"github.com/vintlang/vintlang/lexer"
	"github.com/vintlang/vintlang/parser"
)

func main() {
	input := `import time`
	
	l := lexer.New(input)
	p := parser.New(l)
	program := p.ParseProgram()
	
	fmt.Printf("Program has %d statements\n", len(program.Statements))
	
	if len(program.Statements) > 0 {
		stmt := program.Statements[0]
		fmt.Printf("First statement type: %T\n", stmt)
		
		if exprStmt, ok := stmt.(*ast.ExpressionStatement); ok {
			fmt.Printf("Expression type: %T\n", exprStmt.Expression)
			
			if importExpr, ok := exprStmt.Expression.(*ast.Import); ok {
				fmt.Printf("Import found with %d identifiers:\n", len(importExpr.Identifiers))
				for alias, ident := range importExpr.Identifiers {
					fmt.Printf("  alias: '%s' -> identifier: '%s'\n", alias, ident.Value)
				}
			}
		}
	}
	
	if len(p.Errors()) > 0 {
		fmt.Println("Parser errors:")
		for _, err := range p.Errors() {
			fmt.Println("  ", err)
		}
	}
}