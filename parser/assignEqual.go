package parser

import (
	"fmt"

	"github.com/ekilie/vint-lang/ast"
)

func (p *Parser) parseAssignEqualExpression(exp ast.Expression) ast.Expression {
	switch node := exp.(type) {
	case *ast.Identifier:
		e := &ast.AssignEqual{
			Token: p.curToken,
			Left:  exp.(*ast.Identifier),
		}
		precedence := p.curPrecedence()
		p.nextToken()
		e.Value = p.parseExpression(precedence)
		return e
	case *ast.IndexExpression:
		ae := &ast.AssignmentExpression{Token: p.curToken, Left: exp}

		p.nextToken()

		ae.Value = p.parseExpression(LOWEST)

		return ae
	default:
		if node != nil {
			msg := fmt.Sprintf("Line %d: Expected an identifier or array, but found: %s", p.curToken.Line, node.TokenLiteral())
			p.errors = append(p.errors, msg)
		} else {
			msg := fmt.Sprintf("Line %d: Unexpected syntax encountered during parsing.", p.curToken.Line)
			p.errors = append(p.errors, msg)
		}
		return nil
	}
}
