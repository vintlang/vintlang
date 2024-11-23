package parser

import (
	"github.com/ekilie/vint-lang/ast"
)

func (p *Parser) parseNull() ast.Expression {
	return &ast.Null{Token: p.curToken}
}
