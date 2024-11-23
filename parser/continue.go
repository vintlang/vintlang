package parser

import (
	"github.com/ekilie/vint-lang/ast"
	"github.com/ekilie/vint-lang/token"
)

func (p *Parser) parseContinue() *ast.Continue {
	stmt := &ast.Continue{Token: p.curToken}
	for p.curTokenIs(token.SEMICOLON) {
		p.nextToken()
	}
	return stmt
}
