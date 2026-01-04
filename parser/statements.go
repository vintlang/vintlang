package parser

import (
	"fmt"

	"github.com/vintlang/vintlang/ast"
	"github.com/vintlang/vintlang/token"
)

func (p *Parser) parseStatement() ast.Statement {
	switch p.curToken.Type {
	case token.LET:
		return p.parseLetStatement()
	case token.CONST:
		return p.parseConstStatement()
	case token.ENUM:
		return p.parseEnumStatement()
	case token.RETURN:
		return p.parseReturnStatement()
	case token.BREAK:
		return p.parseBreak()
	case token.CONTINUE:
		return p.parseContinue()
	case token.INCLUDE:
		return p.parseIncludeStatement()
	case token.GO:
		return p.parseGoStatement()
	default:
		return p.parseExpressionStatement()
	}
}

func (p *Parser) parseLetStatement() *ast.LetStatement {
	stmt := &ast.LetStatement{Token: p.curToken}

	if !p.expectPeek(token.IDENT) {
		return nil
	}

	stmt.Name = &ast.Identifier{Token: p.curToken, Value: p.curToken.Literal}

	if !p.expectPeek(token.ASSIGN) {
		return nil
	}

	p.nextToken()

	stmt.Value = p.parseExpression(LOWEST)

	if p.peekTokenIs(token.SEMICOLON) {
		p.nextToken()
	}

	return stmt
}

func (p *Parser) parseConstStatement() *ast.ConstStatement {
	stmt := &ast.ConstStatement{Token: p.curToken}

	if !p.expectPeek(token.IDENT) {
		return nil
	}

	stmt.Name = &ast.Identifier{Token: p.curToken, Value: p.curToken.Literal}

	if !p.expectPeek(token.ASSIGN) {
		return nil
	}

	p.nextToken()

	stmt.Value = p.parseExpression(LOWEST)

	if p.peekTokenIs(token.SEMICOLON) {
		p.nextToken()
	}

	return stmt
}

func (p *Parser) parseReturnStatement() *ast.ReturnStatement {
	stmt := &ast.ReturnStatement{Token: p.curToken}
	p.nextToken()

	stmt.ReturnValue = p.parseExpression(LOWEST)

	if p.peekTokenIs(token.SEMICOLON) {
		p.nextToken()
	}

	return stmt
}

func (p *Parser) parseBlockStatement() *ast.BlockStatement {
	block := &ast.BlockStatement{Token: p.curToken}
	block.Statements = []ast.Statement{}

	p.nextToken()

	for !p.curTokenIs(token.RBRACE) {
		if p.curTokenIs(token.EOF) {
			msg := fmt.Sprintf("Line %d: You did not close the '}' bracket", p.curToken.Line)
			p.errors = append(p.errors, msg)
			return nil
		}
		stmt := p.parseStatement()
		block.Statements = append(block.Statements, stmt)
		p.nextToken()
	}

	return block
}

func (p *Parser) parseIncludeStatement() *ast.IncludeStatement {
	stmt := &ast.IncludeStatement{Token: p.curToken}

	if !p.expectPeek(token.STRING) {
		return nil
	}

	stmt.Path = &ast.StringLiteral{Token: p.curToken, Value: p.curToken.Literal}

	if p.peekTokenIs(token.SEMICOLON) {
		p.nextToken()
	}

	return stmt
}

// parseEnumStatement parses enum declarations
// Syntax: enum Name { MEMBER1 = value1, MEMBER2 = value2 }
func (p *Parser) parseEnumStatement() *ast.EnumStatement {
	stmt := &ast.EnumStatement{Token: p.curToken}
	stmt.Values = make(map[string]ast.Expression)

	// Expect enum name
	if !p.expectPeek(token.IDENT) {
		return nil
	}

	stmt.Name = &ast.Identifier{Token: p.curToken, Value: p.curToken.Literal}

	// Expect opening brace
	if !p.expectPeek(token.LBRACE) {
		return nil
	}

	p.nextToken() // Move past {

	// Parse enum members
	for !p.curTokenIs(token.RBRACE) && !p.curTokenIs(token.EOF) {
		// Get member name
		if !p.curTokenIs(token.IDENT) {
			p.errors = append(p.errors,
				fmt.Sprintf("Line %d: Expected identifier in enum, got %s",
					p.curToken.Line, p.curToken.Type))
			return nil
		}

		memberName := p.curToken.Literal

		// Expect '='
		if !p.expectPeek(token.ASSIGN) {
			return nil
		}

		p.nextToken() // Move to value

		// Parse the value expression
		value := p.parseExpression(LOWEST)
		if value == nil {
			return nil
		}

		stmt.Values[memberName] = value

		// Check for comma or closing brace
		if p.peekTokenIs(token.COMMA) {
			p.nextToken() // Move to comma
			p.nextToken() // Move past comma
		} else if p.peekTokenIs(token.RBRACE) {
			p.nextToken() // Move to closing brace
			break
		} else {
			p.errors = append(p.errors,
				fmt.Sprintf("Line %d: Expected ',' or '}' in enum, got %s",
					p.peekToken.Line, p.peekToken.Type))
			return nil
		}
	}

	// Optional semicolon after closing brace
	if p.peekTokenIs(token.SEMICOLON) {
		p.nextToken()
	}

	return stmt
}
