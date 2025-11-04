package parser

import (
	"fmt"

	"github.com/vintlang/vintlang/ast"
	"github.com/vintlang/vintlang/token"
)

func (p *Parser) parseMatchExpression() ast.Expression {
	expression := &ast.MatchExpression{Token: p.curToken}

	p.nextToken()
	expression.Value = p.parseExpression(LOWEST)

	if expression.Value == nil {
		return nil
	}

	// Expect an opening brace.
	if !p.expectPeek(token.LBRACE) {
		return nil
	}
	p.nextToken()

	// Loop through the match cases.
	for !p.curTokenIs(token.RBRACE) {

		// Check if we encounter EOF before a closing brace.
		if p.curTokenIs(token.EOF) {
			msg := fmt.Sprintf("Line %d: The MATCH statement was not properly closed", p.curToken.Line)
			p.errors = append(p.errors, msg)
			return nil
		}

		matchCase := &ast.MatchCase{Token: p.curToken}

		// Parse the pattern (dict literal, array pattern, or wildcard _)
		matchCase.Pattern = p.parseMatchPattern()

		if matchCase.Pattern == nil {
			return nil
		}

		// Check for guard condition (pattern if condition)
		if p.peekTokenIs(token.IF) {
			p.nextToken() // move to IF
			p.nextToken() // move to condition
			matchCase.Guard = p.parseExpression(LOWEST)
		}

		// Expect an arrow token.
		if !p.expectPeek(token.ARROW) {
			return nil
		}

		// Parse the action - it can be either a block statement or a single expression
		if p.peekTokenIs(token.LBRACE) {
			// Block statement
			if !p.expectPeek(token.LBRACE) {
				return nil
			}
			matchCase.Block = p.parseBlockStatement()
		} else {
			// Single expression - create a block with one expression statement
			p.nextToken()
			expr := p.parseExpression(LOWEST)
			if expr == nil {
				return nil
			}

			exprStmt := &ast.ExpressionStatement{
				Token:      p.curToken,
				Expression: expr,
			}

			matchCase.Block = &ast.BlockStatement{
				Token:      p.curToken,
				Statements: []ast.Statement{exprStmt},
			}
		}
		p.nextToken()
		expression.Cases = append(expression.Cases, matchCase)
	}

	return expression
}

func (p *Parser) parseMatchPattern() ast.Expression {
	switch p.curToken.Type {
	case token.LBRACKET:
		// Parse array pattern in match context
		return p.parseArrayPattern()
	case token.LBRACE:
		// Parse dict pattern
		return p.parseDictLiteral()
	case token.IDENT:
		// Parse identifier (variable binding or wildcard)
		return &ast.Identifier{
			Token: p.curToken,
			Value: p.curToken.Literal,
		}
	default:
		// For other patterns, use regular expression parsing
		return p.parseExpression(LOWEST)
	}
}
