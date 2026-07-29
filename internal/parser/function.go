package parser

import (
	"github.com/vintlang/vintlang/internal/ast"
	"github.com/vintlang/vintlang/internal/token"
)

func (p *Parser) parseFunctionLiteral() ast.Expression {
	tok := p.curToken

	name := ""
	if p.peekTokenIs(token.IDENT) {
		p.nextToken()
		name = p.curToken.Literal
	}

	if !p.expectPeek(token.LPAREN) {
		return nil
	}

	params, hasTypes := p.parseTypedFunctionParameters()
	if params == nil {
		return nil
	}

	var returnType ast.Type
	if p.peekTokenIs(token.COLON) {
		p.nextToken() // consume ':'
		p.nextToken() // move to return type
		returnType = p.parseType()
		hasTypes = true
	}

	if p.curTokenIs(token.LBRACE) {
		// Already at '{' (e.g. after return type was parsed)
	} else if !p.expectPeek(token.LBRACE) {
		return nil
	}

	body := p.parseBlockStatement()

	if !hasTypes {
		lit := &ast.FunctionLiteral{Token: tok, Name: name}
		lit.Defaults = make(map[string]ast.Expression)
		for _, tp := range params {
			lit.Parameters = append(lit.Parameters, tp.Identifier)
			if tp.Default != nil {
				lit.Defaults[tp.Identifier.Value] = tp.Default
			}
		}
		lit.Body = body
		return lit
	}

	flit := &ast.TypedFunctionLiteral{
		Token:      tok,
		Parameters: params,
		ReturnType: returnType,
		Body:       body,
	}
	if name != "" {
		flit.Name = name
	}
	return flit
}

// parseTypedFunctionParameters parses function parameters, supporting typed syntax.
// Returns the list of TypedParameters and whether any type annotations were found.
func (p *Parser) parseTypedFunctionParameters() ([]*ast.TypedParameter, bool) {
	params := []*ast.TypedParameter{}
	hasTypes := false
	hasDefaults := false

	// Advance past '('
	p.nextToken()

	// Handle empty parameter list
	if p.curTokenIs(token.RPAREN) {
		return params, hasTypes
	}

	for {
		// Skip commas between parameters
		if p.curTokenIs(token.COMMA) {
			p.nextToken()
			continue
		}

		// End of parameter list
		if p.curTokenIs(token.RPAREN) {
			break
		}

		if p.curToken.Type != token.IDENT {
			p.addError("expected parameter name, got " + p.curToken.Literal)
			return nil, false
		}

		ident := &ast.Identifier{Token: p.curToken, Value: p.curToken.Literal}
		tp := &ast.TypedParameter{Token: p.curToken, Identifier: ident}

		// Check for type annotation: name: type
		if p.peekTokenIs(token.COLON) {
			p.nextToken() // consume ':'
			p.nextToken() // move to type
			tp.Type = p.parseType()
			hasTypes = true
		}

		// Check for default value: name: type = value
		if p.peekTokenIs(token.ASSIGN) {
			p.nextToken() // '='
			p.nextToken() // value
			tp.Default = p.parseExpression(LOWEST)
			hasDefaults = true
		} else if hasDefaults {
			p.addError("non-default parameter cannot appear after a default parameter")
			return nil, false
		}

		params = append(params, tp)

		// After parsing a parameter, check what comes next.
		// parseType may have left us at ',' (typed param) or at the next token.
		if p.curTokenIs(token.RPAREN) {
			break
		}
		// If we're at a comma, the next iteration will skip it.
		// If we're at an IDENT, the loop will continue parsing it.
		// But we MUST have a comma before the next param.
		if !p.curTokenIs(token.COMMA) && !p.peekTokenIs(token.RPAREN) && !p.peekTokenIs(token.COMMA) {
			p.addError("expected ',' or ')' after parameter, got " + p.peekToken.Literal)
			return nil, false
		}

		p.nextToken()
	}

	return params, hasTypes
}

func (p *Parser) parseCallExpression(function ast.Expression) ast.Expression {
	exp := &ast.CallExpression{Token: p.curToken, Function: function}
	exp.Arguments = p.parseExpressionList(token.RPAREN)
	return exp
}
