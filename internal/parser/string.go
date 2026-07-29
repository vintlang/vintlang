package parser

import (
	"github.com/vintlang/vintlang/internal/ast"
)

func (p *Parser) parseStringLiteral() ast.Expression {
	return &ast.StringLiteral{Token: p.curToken, Value: p.curToken.Literal}
}
