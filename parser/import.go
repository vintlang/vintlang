package parser

import (
	"github.com/vintlang/vintlang/ast"
	"github.com/vintlang/vintlang/token"
)

func (p *Parser) parseImport() ast.Expression {
	exp := &ast.Import{Token: p.curToken}
	exp.Identifiers = make(map[string]*ast.Identifier)
	
	// Parse the first identifier after "import"
	if !p.expectPeek(token.IDENT) {
		return nil
	}
	
	identifier := &ast.Identifier{Value: p.curToken.Literal}
	exp.Identifiers[p.curToken.Literal] = identifier
	
	// Handle comma-separated imports like "import time, math, string"
	for p.peekTokenIs(token.COMMA) {
		p.nextToken() // consume comma
		if !p.expectPeek(token.IDENT) {
			return nil
		}
		identifier := &ast.Identifier{Value: p.curToken.Literal}
		exp.Identifiers[p.curToken.Literal] = identifier
	}

	return exp
}
