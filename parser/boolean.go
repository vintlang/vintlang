package parser

import (
	"github.com/ekilie/vint-lang/ast"
	"github.com/ekilie/vint-lang/token"
)

func (p *Parser) parseBoolean() ast.Expression {
	return &ast.Boolean{Token: p.curToken, Value: p.curTokenIs(token.TRUE)}
}
