package parser

import (
	"github.com/vintlang/vintlang/ast"
	"github.com/vintlang/vintlang/token"
)

func (p *Parser) parseIndexExpression(left ast.Expression) ast.Expression {
	startToken := p.curToken

	p.nextToken()
	
	// Check if we start with a colon (e.g., [:3] or [:])
	if p.curToken.Type == token.COLON {
		// This is a slice expression starting with colon
		sliceExp := &ast.SliceExpression{Token: startToken, Left: left}
		sliceExp.Start = nil
		
		// Check what's after the colon
		if p.peekToken.Type == token.RBRACKET {
			// No end expression either (e.g., [:])
			sliceExp.End = nil
		} else {
			// Move past the colon and parse end expression
			p.nextToken()
			sliceExp.End = p.parseExpression(LOWEST)
		}
		
		if !p.expectPeek(token.RBRACKET) {
			return nil
		}
		
		return sliceExp
	}
	
	// Parse the first expression (could be start index or just a regular index)
	firstExpr := p.parseExpression(LOWEST)
	
	// Now check what's next
	if p.peekToken.Type == token.COLON {
		// This is a slice expression
		sliceExp := &ast.SliceExpression{Token: startToken, Left: left}
		sliceExp.Start = firstExpr
		
		// Move to the colon
		p.nextToken() // Now at colon
		
		// Check what's after the colon
		if p.peekToken.Type == token.RBRACKET {
			// No end expression (e.g., [2:])
			sliceExp.End = nil
			// We're already positioned correctly, expectPeek will advance to RBRACKET
		} else {
			// There's an end expression
			p.nextToken() // Move past colon
			sliceExp.End = p.parseExpression(LOWEST)
		}
		
		if !p.expectPeek(token.RBRACKET) {
			return nil
		}
		
		return sliceExp
	} else {
		// This is a regular index expression
		exp := &ast.IndexExpression{Token: startToken, Left: left}
		exp.Index = firstExpr
		
		if !p.expectPeek(token.RBRACKET) {
			return nil
		}
		
		return exp
	}
}
