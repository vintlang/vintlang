package parser

import (
	"github.com/vintlang/vintlang/ast"
	"github.com/vintlang/vintlang/token"
	"strconv"
)

// typeKeywords is the set of identifiers that are built-in type names.
var typeKeywords = map[string]bool{
	"bool":    true,
	"string":  true,
	"int":     true,
	"int8":    true,
	"int16":   true,
	"int32":   true,
	"int64":   true,
	"uint":    true,
	"uint8":   true,
	"uint16":  true,
	"uint32":  true,
	"uint64":  true,
	"byte":    true,
	"float32": true,
	"float64": true,
	"any":     true,
	"error":   true,
	"nil":     true,
}

func isTypeKeyword(name string) bool {
	return typeKeywords[name]
}

// isTypeStart returns true if the current token can start a type annotation.
func (p *Parser) isTypeStart() bool {
	switch p.curToken.Type {
	case token.IDENT:
		if isTypeKeyword(p.curToken.Literal) {
			return true
		}
		// Check if it's a registered alias
		_, ok := p.aliases[p.curToken.Literal]
		return ok
	case token.ERROR:
		return true
	case token.LBRACKET, token.LBRACE, token.ASTERISK, token.FUNCTION, token.CHAN, token.LPAREN:
		return true
	default:
		return false
	}
}

// parseType parses a type annotation starting at the current token.
// It dispatches based on the current token to produce the right ast.Type node.
func (p *Parser) parseType() ast.Type {
	switch p.curToken.Type {
	case token.IDENT:
		// Check if this identifier is a registered type alias
		if aliasTarget, ok := p.aliases[p.curToken.Literal]; ok {
			p.nextToken()
			return aliasTarget
		}
		return p.parseBasicType()
	case token.ERROR:
		return p.parseBasicType()
	case token.LBRACKET:
		return p.parseArrayOrFixedArrayType()
	case token.LBRACE:
		return p.parseDictType()
	case token.ASTERISK:
		return p.parsePointerType()
	case token.FUNCTION:
		return p.parseFunctionType()
	case token.CHAN:
		return p.parseChannelType()
	case token.LPAREN:
		return p.parseMultiReturnType()
	default:
		p.addError(p.l.GetFilename() + ":" + itoa(p.curToken.Line) + ": expected a type, got " + p.curToken.Literal)
		return nil
	}
}

func (p *Parser) parseBasicType() ast.Type {
	if !p.curTokenIs(token.IDENT) && !p.curTokenIs(token.ERROR) {
		p.addError(p.l.GetFilename() + ":" + itoa(p.curToken.Line) + ": expected type name, got " + p.curToken.Literal)
		return nil
	}
	name := p.curToken.Literal
	tok := p.curToken
	p.nextToken()
	return &ast.BasicType{Token: tok, Name: name}
}

func (p *Parser) parseArrayOrFixedArrayType() ast.Type {
	tok := p.curToken
	p.nextToken() // advance past '['

	// Check if this is a fixed-size array: [n]T
	if p.curTokenIs(token.INT) {
		size, err := strconv.ParseInt(p.curToken.Literal, 10, 64)
		if err != nil {
			p.addError(p.l.GetFilename() + ":" + itoa(p.curToken.Line) + ": invalid array size: " + p.curToken.Literal)
			return nil
		}
		p.nextToken() // advance past size
		if !p.curTokenIs(token.RBRACKET) {
			p.addError(p.l.GetFilename() + ":" + itoa(p.curToken.Line) + ": expected ']' in fixed array type, got " + p.curToken.Literal)
			return nil
		}
		p.nextToken() // advance past ']'
		elemType := p.parseType()
		if elemType == nil {
			return nil
		}
		return &ast.FixedArrayType{Token: tok, Size: size, ElementType: elemType}
	}

	// Slice type: []T — curToken should be ']'
	if !p.curTokenIs(token.RBRACKET) {
		p.addError(p.l.GetFilename() + ":" + itoa(p.curToken.Line) + ": expected ']' in slice type, got " + p.curToken.Literal)
		return nil
	}
	p.nextToken() // advance past ']'
	elemType := p.parseType()
	if elemType == nil {
		return nil
	}
	return &ast.ArrayType{Token: tok, ElementType: elemType}
}

func (p *Parser) parseDictType() ast.Type {
	tok := p.curToken
	p.nextToken() // advance past '{'
	keyType := p.parseType()
	if keyType == nil {
		return nil
	}
	// After parseType, curToken should be ':'
	if !p.curTokenIs(token.COLON) {
		p.addError(p.l.GetFilename() + ":" + itoa(p.curToken.Line) +
			": expected ':' in dict type, got " + p.curToken.Literal)
		return nil
	}
	p.nextToken() // advance past ':'
	valueType := p.parseType()
	if valueType == nil {
		return nil
	}
	// After parseType, curToken should be '}'
	if !p.curTokenIs(token.RBRACE) {
		p.addError(p.l.GetFilename() + ":" + itoa(p.curToken.Line) +
			": expected '}' in dict type, got " + p.curToken.Literal)
		return nil
	}
	p.nextToken() // advance past '}'
	return &ast.DictType{Token: tok, KeyType: keyType, ValueType: valueType}
}

func (p *Parser) parsePointerType() ast.Type {
	tok := p.curToken
	p.nextToken()
	baseType := p.parseType()
	if baseType == nil {
		return nil
	}
	return &ast.PointerType{Token: tok, BaseType: baseType}
}

func (p *Parser) parseChannelType() ast.Type {
	tok := p.curToken
	p.nextToken()
	elemType := p.parseType()
	if elemType == nil {
		return nil
	}
	return &ast.ChannelType{Token: tok, ElementType: elemType}
}

func (p *Parser) parseFunctionType() ast.Type {
	tok := p.curToken
	p.nextToken()

	// p.nextToken() advanced past 'func', so curToken should be '('
	if !p.curTokenIs(token.LPAREN) {
		p.addError(p.l.GetFilename() + ":" + itoa(p.curToken.Line) +
			": expected '(' after 'func' in function type, got " + p.curToken.Literal)
		return nil
	}

	params := []ast.Type{}
	p.nextToken() // advance past '('

	// Parse parameter types
	if !p.curTokenIs(token.RPAREN) {
		first := p.parseType()
		if first == nil {
			return nil
		}
		params = append(params, first)
		for p.curTokenIs(token.COMMA) {
			p.nextToken()
			t := p.parseType()
			if t == nil {
				return nil
			}
			params = append(params, t)
		}
	}

	// curToken should be ')' now
	if !p.curTokenIs(token.RPAREN) {
		p.addError(p.l.GetFilename() + ":" + itoa(p.curToken.Line) +
			": expected ')' in function type, got " + p.curToken.Literal)
		return nil
	}
	p.nextToken() // advance past ')'

	// Check for return type
	if p.curTokenIs(token.LBRACE) || p.curTokenIs(token.SEMICOLON) || p.curTokenIs(token.EOF) {
		return &ast.FunctionType{Token: tok, Parameters: params, ReturnType: nil}
	}

	returnType := p.parseType()
	return &ast.FunctionType{Token: tok, Parameters: params, ReturnType: returnType}
}

func (p *Parser) parseMultiReturnType() ast.Type {
	tok := p.curToken
	p.nextToken() // advance past '('

	types := []ast.Type{}
	if !p.curTokenIs(token.RPAREN) {
		first := p.parseType()
		if first == nil {
			return nil
		}
		types = append(types, first)
		for p.curTokenIs(token.COMMA) {
			p.nextToken()
			t := p.parseType()
			if t == nil {
				return nil
			}
			types = append(types, t)
		}
	}

	// curToken should be ')' now
	if !p.curTokenIs(token.RPAREN) {
		p.addError(p.l.GetFilename() + ":" + itoa(p.curToken.Line) +
			": expected ')' in multi-return type, got " + p.curToken.Literal)
		return nil
	}
	p.nextToken() // advance past ')'

	return &ast.MultiReturnType{Token: tok, Types: types}
}

// parseTypeAliasStatement parses: type Name = TargetType
func (p *Parser) parseTypeAliasStatement() ast.Statement {
	stmt := &ast.TypeAliasStatement{Token: p.curToken}

	if !p.expectPeek(token.IDENT) {
		p.skipToNextStatement()
		return nil
	}

	stmt.Name = &ast.Identifier{Token: p.curToken, Value: p.curToken.Literal}

	if !p.expectPeek(token.ASSIGN) {
		p.skipToNextStatement()
		return nil
	}

	p.nextToken()
	stmt.Target = p.parseType()
	if stmt.Target == nil {
		p.skipToNextStatement()
		return nil
	}

	// Register the alias so subsequent type references can resolve it
	p.aliases[stmt.Name.Value] = stmt.Target

	if p.peekTokenIs(token.SEMICOLON) {
		p.nextToken()
	}

	return stmt
}

// parseTypeCast parses: expression as Type
func (p *Parser) parseTypeCast(left ast.Expression) ast.Expression {
	tok := p.curToken
	p.nextToken()
	targetType := p.parseTypeNoAdvance()
	if targetType == nil {
		return nil
	}
	return &ast.TypeCastExpression{Token: tok, Expression: left, TargetType: targetType}
}

// parseTypeCheck parses: expression is Type
func (p *Parser) parseTypeCheck(left ast.Expression) ast.Expression {
	tok := p.curToken
	p.nextToken()
	checkType := p.parseTypeNoAdvance()
	if checkType == nil {
		return nil
	}
	return &ast.TypeCheckExpression{Token: tok, Expression: left, CheckType: checkType}
}

// parseTypeNoAdvance is like parseType but does NOT advance past the type token for basic types.
// This prevents consuming group terminators like ')' in as/is expressions.
func (p *Parser) parseTypeNoAdvance() ast.Type {
	switch p.curToken.Type {
	case token.IDENT:
		if aliasTarget, ok := p.aliases[p.curToken.Literal]; ok {
			return aliasTarget
		}
		name := p.curToken.Literal
		tok := p.curToken
		return &ast.BasicType{Token: tok, Name: name}
	case token.ERROR:
		name := p.curToken.Literal
		tok := p.curToken
		return &ast.BasicType{Token: tok, Name: name}
	case token.LBRACKET, token.LBRACE, token.ASTERISK, token.FUNCTION, token.CHAN, token.LPAREN:
		// For compound types, use parseType (may have advancing issues inside groups)
		return p.parseType()
	default:
		p.addError(p.l.GetFilename() + ":" + itoa(p.curToken.Line) + ": expected a type, got " + p.curToken.Literal)
		return nil
	}
}

func itoa(n int) string {
	return strconv.Itoa(n)
}
