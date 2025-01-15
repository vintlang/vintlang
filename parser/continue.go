package parser

import (
	"github.com/vintlang/vintlang/ast"
	"github.com/vintlang/vintlang/token"
)

func (p *Parser) parseContinue() *ast.Continue {
	stmt := &ast.Continue{Token: p.curToken}
	for p.curTokenIs(token.SEMICOLON) {
		p.nextToken()
	}
	return stmt
}
