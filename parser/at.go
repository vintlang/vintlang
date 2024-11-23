package parser

import "github.com/ekilie/vint-lang/ast"

func (p *Parser) parseAt() ast.Expression {
	return &ast.At{Token: p.curToken}
}
