package parser

import (
	"github.com/vintlang/vintlang/internal/ast"
)

func (p *Parser) parseNull() ast.Expression {
	return &ast.Null{Token: p.curToken}
}
