package parser

import (
	"github.com/vintlang/vintlang/ast"
)

func (p *Parser) parseNull() ast.Expression {
	return &ast.Null{Token: p.curToken}
}
