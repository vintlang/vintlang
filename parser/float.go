package parser

import (
	"fmt"
	"strconv"

	"github.com/vintlang/vintlang/ast"
)

func (p *Parser) parseFloatLiteral() ast.Expression {
	fl := &ast.FloatLiteral{Token: p.curToken}
	value, err := strconv.ParseFloat(p.curToken.Literal, 64)
	if err != nil {
		msg := fmt.Sprintf("Line %d: We cannot parse %q as a float", p.curToken.Line, p.curToken.Literal)
		p.errors = append(p.errors, msg)
		return nil
	}
	fl.Value = value
	return fl
}
