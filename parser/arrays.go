package parser

import (
	"github.com/vintlang/vintlang/ast"
	"github.com/vintlang/vintlang/token"
)

func (p *Parser) parseArrayLiteral() ast.Expression {
	array := &ast.ArrayLiteral{Token: p.curToken}
	array.Elements = p.parseExpressionList(token.RBRACKET)

	return array
}

func (p *Parser) parseArrayPattern() ast.Expression {
	pattern := &ast.ArrayPattern{Token: p.curToken}

	if p.peekTokenIs(token.RBRACKET) {
		p.nextToken()
		return pattern // Empty array pattern []
	}

	p.nextToken()

	// Parse array pattern elements
	for !p.curTokenIs(token.RBRACKET) {
		if p.curTokenIs(token.ELLIPSIS) {
			// Handle spread pattern ...rest
			if pattern.Rest != nil {
				p.addError("Only one spread element allowed in array pattern")
				return nil
			}

			p.nextToken() // move to identifier after ...
			if !p.curTokenIs(token.IDENT) {
				p.addError("Expected identifier after ... in array pattern")
				return nil
			}

			pattern.Rest = &ast.Identifier{
				Token: p.curToken,
				Value: p.curToken.Literal,
			}

			p.nextToken()
			// Spread must be last element
			if p.curTokenIs(token.COMMA) {
				p.addError("Spread element must be last in array pattern")
				return nil
			}
			break
		} else {
			// Regular pattern element
			pattern.Elements = append(pattern.Elements, p.parseExpression(LOWEST))

			if p.peekTokenIs(token.COMMA) {
				p.nextToken() // consume comma
				p.nextToken() // move to next element
			} else if p.peekTokenIs(token.RBRACKET) {
				break
			} else {
				p.addError("Expected ',' or ']' in array pattern")
				return nil
			}
		}
	}

	if !p.expectPeek(token.RBRACKET) {
		return nil
	}

	return pattern
}

func (p *Parser) parseExpressionList(end token.TokenType) []ast.Expression {
	list := []ast.Expression{}

	if p.peekTokenIs(end) {
		p.nextToken()
		return list
	}

	p.nextToken()
	list = append(list, p.parseExpression(LOWEST))

	for p.peekTokenIs(token.COMMA) {
		p.nextToken()
		p.nextToken()
		list = append(list, p.parseExpression(LOWEST))
	}

	if !p.expectPeek(end) {
		return nil
	}
	return list
}

func (p *Parser) parseSpreadPattern() ast.Expression {
	spread := &ast.SpreadPattern{Token: p.curToken}

	if p.peekTokenIs(token.IDENT) {
		p.nextToken()
		spread.Name = &ast.Identifier{
			Token: p.curToken,
			Value: p.curToken.Literal,
		}
	}

	return spread
}
