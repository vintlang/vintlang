package parser

import (
	"github.com/vintlang/vintlang/internal/ast"
	"github.com/vintlang/vintlang/internal/token"
)

func (p *Parser) parseBreak() *ast.Break {
	stmt := &ast.Break{Token: p.curToken}
	for p.curTokenIs(token.SEMICOLON) {
		p.nextToken()
	}
	return stmt
}
