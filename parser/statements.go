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
	case token.STRUCT:
		return p.parseStructStatement()
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
		// Contextual keyword: 'type' at statement start is a type alias
		// but only if the next token is an identifier (not '(' which means function call).
		if p.curTokenIs(token.IDENT) && p.curToken.Literal == "type" &&
			p.peekTokenIs(token.IDENT) {
			return p.parseTypeAliasStatement()
		}
		return p.parseExpressionStatement()
	}
}

func (p *Parser) parseLetStatement() ast.Statement {
	tok := p.curToken

	if !p.expectPeek(token.IDENT) {
		p.skipToNextStatement()
		return nil
	}

	name := &ast.Identifier{Token: p.curToken, Value: p.curToken.Literal}

	// Check for type annotation: let x: type
	if p.peekTokenIs(token.COLON) {
		p.nextToken() // consume ':'
		p.nextToken()

		// For simple zero-value declarations (let x: int), avoid parseType()
		// since it would advance past the type and potentially consume the
		// next statement. For compound types or typed values, use parseType.
		if p.curTokenIs(token.IDENT) || p.curTokenIs(token.ERROR) {
			// Simple type annotation: check if '=' follows (value) or not (zero value)
			// Need to peek past the type to check for '='
			// Save and restore: advance to peek, check, then restore
			// Actually, simpler: just check if the type keyword is followed by '='
			// For basic types: int, string, bool, float64, etc.
			typeName := p.curToken.Literal
			typeTok := p.curToken

			// Check what follows this identifier — if it's '=', we have a value
			if p.peekTokenIs(token.ASSIGN) {
				// let x: int = 5 — curToken is 'int', parseType reads it and advances to '='
				declaredType := p.parseType()
				if declaredType == nil {
					p.skipToNextStatement()
					return nil
				}
				if p.curTokenIs(token.ASSIGN) {
					p.nextToken()
					value := p.parseExpression(LOWEST)
					if p.peekTokenIs(token.SEMICOLON) {
						p.nextToken()
					}
					return &ast.TypedLetStatement{
						Token: tok,
						Name:  name,
						TypeAnnotation: &ast.TypeAnnotation{
							Token: tok,
							Type:  declaredType,
						},
						Value: value,
					}
				}
			}

			// Zero value: let x: int
			declaredType := &ast.BasicType{Token: typeTok, Name: typeName}
			if p.peekTokenIs(token.SEMICOLON) {
				p.nextToken()
			}
			return &ast.TypedLetStatement{
				Token: tok,
				Name:  name,
				TypeAnnotation: &ast.TypeAnnotation{
					Token: tok,
					Type:  declaredType,
				},
			}
		}

		// Compound type or unknown: use parseType (may have advancing issues
		// with zero-value declarations but handles complex type syntax)
		declaredType := p.parseType()
		if declaredType == nil {
			p.skipToNextStatement()
			return nil
		}
		if p.curTokenIs(token.ASSIGN) {
			p.nextToken()
			value := p.parseExpression(LOWEST)
			if p.peekTokenIs(token.SEMICOLON) {
				p.nextToken()
			}
			return &ast.TypedLetStatement{
				Token: tok,
				Name:  name,
				TypeAnnotation: &ast.TypeAnnotation{
					Token: tok,
					Type:  declaredType,
				},
				Value: value,
			}
		}
		if p.peekTokenIs(token.SEMICOLON) {
			p.nextToken()
		}
		return &ast.TypedLetStatement{
			Token: tok,
			Name:  name,
			TypeAnnotation: &ast.TypeAnnotation{
				Token: tok,
				Type:  declaredType,
			},
		}
	}

	// No type annotation: let x = value (inferred) or error
	if p.peekTokenIs(token.ASSIGN) {
		p.nextToken()
		p.nextToken()
		stmt := &ast.LetStatement{Token: tok, Name: name}
		stmt.Value = p.parseExpression(LOWEST)
		if p.peekTokenIs(token.SEMICOLON) {
			p.nextToken()
		}
		return stmt
	}

	// let x with no type, no value — error in strict mode
	p.addError(p.l.GetFilename() + ":" + itoa(p.curToken.Line) +
		": variable '" + name.Value + "' must have either a type annotation or an initializer")
	p.skipToNextStatement()
	return nil
}

func (p *Parser) parseConstStatement() ast.Statement {
	tok := p.curToken

	if !p.expectPeek(token.IDENT) {
		p.skipToNextStatement()
		return nil
	}

	name := &ast.Identifier{Token: p.curToken, Value: p.curToken.Literal}

	// Check for type annotation: const x: type
	if p.peekTokenIs(token.COLON) {
		p.nextToken() // consume ':'
		p.nextToken()

		// For simple types, avoid parseType advancing issue
		if p.curTokenIs(token.IDENT) || p.curTokenIs(token.ERROR) {
			typeName := p.curToken.Literal
			typeTok := p.curToken

			if p.peekTokenIs(token.ASSIGN) {
				declaredType := p.parseType()
				if declaredType == nil {
					p.skipToNextStatement()
					return nil
				}
				if p.curTokenIs(token.ASSIGN) {
					p.nextToken()
					value := p.parseExpression(LOWEST)
					if p.peekTokenIs(token.SEMICOLON) {
						p.nextToken()
					}
					return &ast.TypedLetStatement{
						Token: tok,
						Name:  name,
						TypeAnnotation: &ast.TypeAnnotation{
							Token: tok,
							Type:  declaredType,
						},
						Value: value,
					}
				}
			}

			declaredType := &ast.BasicType{Token: typeTok, Name: typeName}
			if p.peekTokenIs(token.SEMICOLON) {
				p.nextToken()
			}
			return &ast.TypedLetStatement{
				Token: tok,
				Name:  name,
				TypeAnnotation: &ast.TypeAnnotation{
					Token: tok,
					Type:  declaredType,
				},
			}
		}

		declaredType := p.parseType()
		if declaredType == nil {
			p.skipToNextStatement()
			return nil
		}
		if p.curTokenIs(token.ASSIGN) {
			p.nextToken()
			value := p.parseExpression(LOWEST)
			if p.peekTokenIs(token.SEMICOLON) {
				p.nextToken()
			}
			return &ast.TypedLetStatement{
				Token: tok,
				Name:  name,
				TypeAnnotation: &ast.TypeAnnotation{
					Token: tok,
					Type:  declaredType,
				},
				Value: value,
			}
		}
		if p.peekTokenIs(token.SEMICOLON) {
			p.nextToken()
		}
		return &ast.TypedLetStatement{
			Token: tok,
			Name:  name,
			TypeAnnotation: &ast.TypeAnnotation{
				Token: tok,
				Type:  declaredType,
			},
		}
	}

	// No type annotation: const x = value (inferred) or error
	if p.peekTokenIs(token.ASSIGN) {
		p.nextToken()
		p.nextToken()
		stmt := &ast.ConstStatement{Token: tok, Name: name}
		stmt.Value = p.parseExpression(LOWEST)
		if p.peekTokenIs(token.SEMICOLON) {
			p.nextToken()
		}
		return stmt
	}

	p.addError(p.l.GetFilename() + ":" + itoa(p.curToken.Line) +
		": constant '" + name.Value + "' must have either a type annotation or an initializer")
	p.skipToNextStatement()
	return nil
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

// parseStructStatement parses struct declarations
// Syntax: struct Name { field1: default1, field2: default2, func method() { ... } }
func (p *Parser) parseStructStatement() *ast.StructStatement {
	stmt := &ast.StructStatement{Token: p.curToken}

	// Expect struct name
	if !p.expectPeek(token.IDENT) {
		return nil
	}

	stmt.Name = &ast.Identifier{Token: p.curToken, Value: p.curToken.Literal}

	// Expect opening brace
	if !p.expectPeek(token.LBRACE) {
		return nil
	}

	p.nextToken() // Move past {

	// Parse struct members (fields and methods)
	for !p.curTokenIs(token.RBRACE) && !p.curTokenIs(token.EOF) {
		// Check if this is a method (starts with 'func')
		if p.curTokenIs(token.FUNCTION) {
			method := p.parseStructMethod()
			if method == nil {
				return nil
			}
			stmt.Methods = append(stmt.Methods, *method)

			// Skip comma if present after method
			if p.peekTokenIs(token.COMMA) {
				p.nextToken()
			}
		} else if p.curTokenIs(token.IDENT) {
			// It's a field
			field := ast.StructField{}
			field.Name = &ast.Identifier{Token: p.curToken, Value: p.curToken.Literal}

			// Check for ':' — could be type annotation or default value
			if p.peekTokenIs(token.COLON) {
				p.nextToken() // Move to ':'
				p.nextToken() // Move to next token

				// Disambiguate: type keyword or compound type start → type annotation
				// Otherwise → default value expression (back-compat)
				if p.isTypeStart() {
					field.Type = p.parseType()
					// After parseType, curToken might be '=' for default value
					if p.curTokenIs(token.ASSIGN) {
						p.nextToken()
						field.Default = p.parseExpression(LOWEST)
					}
				} else {
					field.Default = p.parseExpression(LOWEST)
				}
			}

			stmt.Fields = append(stmt.Fields, field)

			// Skip comma if present
			if p.peekTokenIs(token.COMMA) {
				p.nextToken()
			}
		} else {
			p.errors = append(p.errors,
				fmt.Sprintf("Line %d: Expected field name or 'func' in struct, got %s",
					p.curToken.Line, p.curToken.Type))
			return nil
		}

		p.nextToken()
	}

	// Optional semicolon after closing brace
	if p.peekTokenIs(token.SEMICOLON) {
		p.nextToken()
	}

	return stmt
}

// parseStructMethod parses a method inside a struct declaration
func (p *Parser) parseStructMethod() *ast.StructMethod {
	method := &ast.StructMethod{}
	method.Defaults = make(map[string]ast.Expression)

	// Expect method name
	if !p.expectPeek(token.IDENT) {
		return nil
	}

	method.Name = &ast.Identifier{Token: p.curToken, Value: p.curToken.Literal}

	// Expect opening paren
	if !p.expectPeek(token.LPAREN) {
		return nil
	}

	// Parse parameters using the typed function parameters logic
	typedParams, _ := p.parseTypedFunctionParameters()
	if typedParams == nil {
		return nil
	}

	// Convert TypedParameters to simple identifiers + types
	for _, tp := range typedParams {
		method.Parameters = append(method.Parameters, tp.Identifier)
		method.ParamTypes = append(method.ParamTypes, tp.Type)
		if tp.Default != nil {
			method.Defaults[tp.Identifier.Value] = tp.Default
		}
	}

	// Check for return type: func method(): returnType { ... }
	var returnType ast.Type
	if p.peekTokenIs(token.COLON) {
		p.nextToken() // consume ':'
		p.nextToken() // move to return type
		returnType = p.parseType()
	}
	method.ReturnType = returnType

	// Expect opening brace for body
	if p.curTokenIs(token.LBRACE) {
		// Already at '{' (e.g. after return type was parsed)
	} else if !p.expectPeek(token.LBRACE) {
		return nil
	}

	method.Body = p.parseBlockStatement()

	return method
}
