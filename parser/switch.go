package parser

import (
	"fmt"

	"github.com/ekilie/vint-lang/ast"
	"github.com/ekilie/vint-lang/token"
)

func (p *Parser) parseSwitchStatement() ast.Expression {
	expression := &ast.SwitchExpression{Token: p.curToken}

	// Expect an opening parenthesis.
	if !p.expectPeek(token.LPAREN) {
		return nil
	}

	p.nextToken()
	expression.Value = p.parseExpression(LOWEST)

	if expression.Value == nil {
		return nil
	}

	// Expect a closing parenthesis.
	if !p.expectPeek(token.RPAREN) {
		return nil
	}

	// Expect an opening brace.
	if !p.expectPeek(token.LBRACE) {
		return nil
	}
	p.nextToken()

	// Loop through the cases.
	for !p.curTokenIs(token.RBRACE) {

		// Check if we encounter EOF before a closing brace.
		if p.curTokenIs(token.EOF) {
			msg := fmt.Sprintf("Line %d: The SWITCH statement was not properly closed", p.curToken.Line)
			p.errors = append(p.errors, msg)
			return nil
		}
		tmp := &ast.CaseExpression{Token: p.curToken}

		// If it's a default case.
		if p.curTokenIs(token.DEFAULT) {
			tmp.Default = true
		} else if p.curTokenIs(token.CASE) {

			// Parse the CASE expression.
			p.nextToken()

			if p.curTokenIs(token.DEFAULT) {
				tmp.Default = true
			} else {
				tmp.Expr = append(tmp.Expr, p.parseExpression(LOWEST))
				// Handle multiple expressions in the case.
				for p.peekTokenIs(token.COMMA) {
					p.nextToken()
					p.nextToken()
					tmp.Expr = append(tmp.Expr, p.parseExpression(LOWEST))
				}
			}
		} else {
			msg := fmt.Sprintf("Line %d: Expected CASE or DEFAULT, but received: %s", p.curToken.Line, p.curToken.Type)
			p.errors = append(p.errors, msg)
			return nil
		}

		// Expect an opening brace for the case block.
		if !p.expectPeek(token.LBRACE) {
			return nil
		}

		// Parse the block statement inside the case.
		tmp.Block = p.parseBlockStatement()
		p.nextToken()
		expression.Choices = append(expression.Choices, tmp)
	}

	// Count how many default cases there are.
	count := 0
	for _, c := range expression.Choices {
		if c.Default {
			count++
		}
	}
	if count > 1 {
		msg := fmt.Sprintf("A SWITCH statement can only have one DEFAULT case! You have %d", count)
		p.errors = append(p.errors, msg)
		return nil
	}

	return expression
}
