package parser

import (
	"github.com/vintlang/vintlang/internal/ast"
	"github.com/vintlang/vintlang/internal/token"
)

func (p *Parser) parseContinue() *ast.Continue {
	stmt := &ast.Continue{Token: p.curToken}
	for p.curTokenIs(token.SEMICOLON) {
		p.nextToken()
	}
	return stmt
}
