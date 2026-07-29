package parser

import (
	"github.com/vintlang/vintlang/internal/ast"
	"github.com/vintlang/vintlang/internal/token"
)

func (p *Parser) parseAt() ast.Expression {
	return &ast.At{Token: p.curToken}
}

// parseBuiltinExpression parses: ::identifier
// The :: prefix creates a reference to a builtin function.
// When followed by (args), it forms a builtin call: ::println("hello")
func (p *Parser) parseBuiltinExpression() ast.Expression {
	tok := p.curToken // the '::' token
	if !p.expectPeek(token.IDENT) {
		return nil
	}
	name := p.curToken.Literal
	return &ast.BuiltinExpression{Token: tok, Name: name}
}
