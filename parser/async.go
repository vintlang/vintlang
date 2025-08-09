package parser

import (
	"github.com/vintlang/vintlang/ast"
	"github.com/vintlang/vintlang/token"
)

// parseAsyncFunctionLiteral parses async function literals
func (p *Parser) parseAsyncFunctionLiteral() ast.Expression {
	lit := &ast.AsyncFunctionLiteral{Token: p.curToken}

	if !p.expectPeek(token.FUNCTION) {
		return nil
	}

	if !p.expectPeek(token.LPAREN) {
		return nil
	}

	if !p.parseAsyncFunctionParameters(lit) {
		return nil
	}

	if !p.expectPeek(token.LBRACE) {
		return nil
	}

	lit.Body = p.parseBlockStatement()

	return lit
}

// parseAsyncFunctionParameters parses parameters for async functions
func (p *Parser) parseAsyncFunctionParameters(lit *ast.AsyncFunctionLiteral) bool {
	lit.Defaults = make(map[string]ast.Expression)
	hasDefaults := false // Track if any default parameter has been encountered

	for !p.peekTokenIs(token.RPAREN) {
		p.nextToken()

		if p.curTokenIs(token.COMMA) {
			continue
		}

		ident := &ast.Identifier{Token: p.curToken, Value: p.curToken.Literal}
		lit.Parameters = append(lit.Parameters, ident)

		if p.peekTokenIs(token.ASSIGN) {
			p.nextToken() // Consume '='
			p.nextToken() // Parse default expression
			lit.Defaults[ident.Value] = p.parseExpression(LOWEST)
			hasDefaults = true
		} else {
			if hasDefaults {
				p.addError("Non-default parameter cannot appear after a default parameter")
				return false
			}
		}

		if !(p.peekTokenIs(token.COMMA) || p.peekTokenIs(token.RPAREN)) {
			return false
		}
	}

	return p.expectPeek(token.RPAREN)
}

// parseAwaitExpression parses await expressions
func (p *Parser) parseAwaitExpression() ast.Expression {
	expr := &ast.AwaitExpression{Token: p.curToken}

	p.nextToken()
	expr.Value = p.parseExpression(PREFIX)

	return expr
}

// parseGoStatement parses go statements
func (p *Parser) parseGoStatement() ast.Statement {
	stmt := &ast.GoStatement{Token: p.curToken}

	p.nextToken()
	stmt.Expression = p.parseExpression(LOWEST)

	if p.peekTokenIs(token.SEMICOLON) {
		p.nextToken()
	}

	return stmt
}

// parseChannelExpression parses channel creation expressions
func (p *Parser) parseChannelExpression() ast.Expression {
	expr := &ast.ChannelExpression{Token: p.curToken}

	// Check if there's a buffer size specified
	if p.peekTokenIs(token.LPAREN) {
		p.nextToken() // consume 'chan'
		p.nextToken() // consume '('
		
		expr.Buffer = p.parseExpression(LOWEST)
		
		if !p.expectPeek(token.RPAREN) {
			return nil
		}
	}

	return expr
}