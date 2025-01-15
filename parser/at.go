package parser

import "github.com/vintlang/vintlang/ast"

func (p *Parser) parseAt() ast.Expression {
	return &ast.At{Token: p.curToken}
}
