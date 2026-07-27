package parser

import (
	"github.com/vintlang/vintlang/ast"
)

func (p *Parser) parseIdentifier() ast.Expression {
	ident := &ast.Identifier{Token: p.curToken, Value: p.curToken.Literal}
	p.nextToken()
	return ident
}
