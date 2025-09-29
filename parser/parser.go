package parser

import (
	"fmt"

	"github.com/vintlang/vintlang/ast"
	"github.com/vintlang/vintlang/lexer"
	"github.com/vintlang/vintlang/token"
)

const (
	_ int = iota
	LOWEST
	ASSIGN      // =
	COND        // OR or AND
	EQUALS      // ==
	LESSGREATER // > OR <
	RANGE       // ..
	SUM         // +
	PRODUCT     // *
	POWER       // ** we got the power XD
	MODULUS     // %
	PREFIX      //  -X OR !X
	CALL        // myFunction(X)
	INDEX       // Arrays
	DOT         // For methods
)

var precedences = map[token.TokenType]int{
	token.AND:             COND,
	token.OR:              COND,
	token.NULL_COALESCE:   COND,
	token.IN:              COND,
	token.ASSIGN:          ASSIGN,
	token.EQ:              EQUALS,
	token.NOT_EQ:          EQUALS,
	token.LT:              LESSGREATER,
	token.LTE:             LESSGREATER,
	token.GT:              LESSGREATER,
	token.GTE:             LESSGREATER,
	token.RANGE:           RANGE,
	token.PLUS:            SUM,
	token.PLUS_ASSIGN:     SUM,
	token.MINUS:           SUM,
	token.MINUS_ASSIGN:    SUM,
	token.SLASH:           PRODUCT,
	token.SLASH_ASSIGN:    PRODUCT,
	token.ASTERISK:        PRODUCT,
	token.ASTERISK_ASSIGN: PRODUCT,
	token.POW:             POWER,
	token.MODULUS:         MODULUS,
	token.MODULUS_ASSIGN:  MODULUS,
	// token.BANG:     PREFIX,
	token.LPAREN:   CALL,
	token.LBRACKET: INDEX,
	token.DOT:      DOT, // Highest priority
}

type (
	prefixParseFn  func() ast.Expression
	infixParseFn   func(ast.Expression) ast.Expression
	postfixParseFn func() ast.Expression
)

type Parser struct {
	l *lexer.Lexer

	curToken  token.Token
	peekToken token.Token
	prevToken token.Token

	errors []string

	prefixParseFns  map[token.TokenType]prefixParseFn
	infixParseFns   map[token.TokenType]infixParseFn
	postfixParseFns map[token.TokenType]postfixParseFn
}

// adds error
func (p *Parser) addError(msg string) {
	p.errors = append(p.errors, msg)
}

func (p *Parser) registerPrefix(tokenType token.TokenType, fn prefixParseFn) {
	p.prefixParseFns[tokenType] = fn
}

func (p *Parser) registerInfix(tokenType token.TokenType, fn infixParseFn) {
	p.infixParseFns[tokenType] = fn
}

func (p *Parser) registerPostfix(tokenType token.TokenType, fn postfixParseFn) {
	p.postfixParseFns[tokenType] = fn
}

func New(l *lexer.Lexer) *Parser {
	p := &Parser{l: l, errors: []string{}}

	p.nextToken()
	p.nextToken()

	p.prefixParseFns = make(map[token.TokenType]prefixParseFn)
	p.registerPrefix(token.STRING, p.parseStringLiteral)
	p.registerPrefix(token.IDENT, p.parseIdentifier)
	p.registerPrefix(token.INT, p.parseIntegerLiteral)
	p.registerPrefix(token.FLOAT, p.parseFloatLiteral)
	p.registerPrefix(token.BANG, p.parsePrefixExpression)
	p.registerPrefix(token.MINUS, p.parsePrefixExpression)
	p.registerPrefix(token.PLUS, p.parsePrefixExpression)
	p.registerPrefix(token.AMPERSAND, p.parsePrefixExpression)
	p.registerPrefix(token.ASTERISK, p.parsePrefixExpression)
	p.registerPrefix(token.TRUE, p.parseBoolean)
	p.registerPrefix(token.FALSE, p.parseBoolean)
	p.registerPrefix(token.LPAREN, p.parseGroupedExpression)
	p.registerPrefix(token.IF, p.parseIfExpression)
	p.registerPrefix(token.FUNCTION, p.parseFunctionLiteral)
	p.registerPrefix(token.LBRACKET, p.parseArrayLiteral)
	p.registerPrefix(token.LBRACE, p.parseDictLiteral)
	p.registerPrefix(token.WHILE, p.parseWhileExpression)
	p.registerPrefix(token.NULL, p.parseNull)
	p.registerPrefix(token.FOR, p.parseForExpression)
	p.registerPrefix(token.SWITCH, p.parseSwitchStatement)
	p.registerPrefix(token.MATCH, p.parseMatchExpression)
	p.registerPrefix(token.IMPORT, p.parseImport)
	p.registerPrefix(token.PACKAGE, p.parsePackage)
	p.registerPrefix(token.TODO, p.parseTodoStatement)
	p.registerPrefix(token.WARN, p.parseWarnStatement)
	p.registerPrefix(token.ERROR, p.parseErrorStatement)
	p.registerPrefix(token.DEFER, p.parseDeferStatement)
	p.registerPrefix(token.AT, p.parseAt)
	p.registerPrefix(token.INFO, p.parseInfoStatement)
	p.registerPrefix(token.DEBUG, p.parseDebugStatement)
	p.registerPrefix(token.NOTE, p.parseNoteStatement)
	p.registerPrefix(token.SUCCESS, p.parseSuccessStatement)
	p.registerPrefix(token.TRACE, p.parseTraceStatement)
	p.registerPrefix(token.FATAL, p.parseFatalStatement)
	p.registerPrefix(token.CRITICAL, p.parseCriticalStatement)
	p.registerPrefix(token.LOG, p.parseLogStatement)
	p.registerPrefix(token.REPEAT, p.parseRepeatStatement)

	// Capitalized declaratives
	p.registerPrefix(token.INFO_CAP, p.parseInfoStatement)
	p.registerPrefix(token.DEBUG_CAP, p.parseDebugStatement)
	p.registerPrefix(token.NOTE_CAP, p.parseNoteStatement)
	p.registerPrefix(token.TODO_CAP, p.parseTodoStatement)
	p.registerPrefix(token.WARN_CAP, p.parseWarnStatement)
	p.registerPrefix(token.SUCCESS_CAP, p.parseSuccessStatement)
	p.registerPrefix(token.ERROR_CAP, p.parseErrorStatement)
	p.registerPrefix(token.TRACE_CAP, p.parseTraceStatement)
	p.registerPrefix(token.FATAL_CAP, p.parseFatalStatement)
	p.registerPrefix(token.CRITICAL_CAP, p.parseCriticalStatement)
	p.registerPrefix(token.LOG_CAP, p.parseLogStatement)

	// Async/Concurrency prefix parsers
	p.registerPrefix(token.ASYNC, p.parseAsyncFunctionLiteral)
	p.registerPrefix(token.AWAIT, p.parseAwaitExpression)
	p.registerPrefix(token.CHAN, p.parseChannelExpression)

	// Error handling prefix parsers
	p.registerPrefix(token.THROW, p.parseThrowStatement)

	p.infixParseFns = make(map[token.TokenType]infixParseFn)
	p.registerInfix(token.AND, p.parseInfixExpression)
	p.registerInfix(token.OR, p.parseInfixExpression)
	p.registerInfix(token.NULL_COALESCE, p.parseInfixExpression)
	p.registerInfix(token.PLUS, p.parseInfixExpression)
	p.registerInfix(token.PLUS_ASSIGN, p.parseAssignEqualExpression)
	p.registerInfix(token.MINUS, p.parseInfixExpression)
	p.registerInfix(token.MINUS_ASSIGN, p.parseAssignEqualExpression)
	p.registerInfix(token.SLASH, p.parseInfixExpression)
	p.registerInfix(token.SLASH_ASSIGN, p.parseAssignEqualExpression)
	p.registerInfix(token.ASTERISK, p.parseInfixExpression)
	p.registerInfix(token.ASTERISK_ASSIGN, p.parseAssignEqualExpression)
	p.registerInfix(token.POW, p.parseInfixExpression)
	p.registerInfix(token.MODULUS, p.parseInfixExpression)
	p.registerInfix(token.MODULUS_ASSIGN, p.parseAssignmentExpression)
	p.registerInfix(token.EQ, p.parseInfixExpression)
	p.registerInfix(token.NOT_EQ, p.parseInfixExpression)
	p.registerInfix(token.LT, p.parseInfixExpression)
	p.registerInfix(token.LTE, p.parseInfixExpression)
	p.registerInfix(token.GT, p.parseInfixExpression)
	p.registerInfix(token.GTE, p.parseInfixExpression)
	p.registerInfix(token.LPAREN, p.parseCallExpression)
	p.registerInfix(token.LBRACKET, p.parseIndexExpression)
	p.registerInfix(token.ASSIGN, p.parseAssignmentExpression)
	p.registerInfix(token.IN, p.parseInfixExpression)
	p.registerInfix(token.DOT, p.parseMethod)
	p.registerInfix(token.RANGE, p.parseRangeExpression)

	p.postfixParseFns = make(map[token.TokenType]postfixParseFn)
	p.registerPostfix(token.PLUS_PLUS, p.parsePostfixExpression)
	p.registerPostfix(token.MINUS_MINUS, p.parsePostfixExpression)

	return p
}

func (p *Parser) ParseProgram() *ast.Program {
	program := &ast.Program{}
	program.Statements = []ast.Statement{}

	for !p.curTokenIs(token.EOF) {
		stmt := p.parseStatement()
		program.Statements = append(program.Statements, stmt)

		p.nextToken()
	}
	return program
}

// manage token literals:

func (p *Parser) nextToken() {
	p.prevToken = p.curToken
	p.curToken = p.peekToken
	p.peekToken = p.l.NextToken()
}

func (p *Parser) curTokenIs(t token.TokenType) bool {
	return p.curToken.Type == t
}

func (p *Parser) peekTokenIs(t token.TokenType) bool {
	return p.peekToken.Type == t
}

func (p *Parser) expectPeek(t token.TokenType) bool {
	if p.peekTokenIs(t) {
		p.nextToken()
		return true
	} else {
		p.peekError(t)
		return false
	}
}

func (p *Parser) peekPrecedence() int {
	if p, ok := precedences[p.peekToken.Type]; ok {
		return p
	}
	return LOWEST
}

func (p *Parser) curPrecedence() int {
	if p, ok := precedences[p.curToken.Type]; ok {
		return p
	}

	return LOWEST
}

// error messages

func (p *Parser) Errors() []string {
	return p.errors
}

func (p *Parser) peekError(t token.TokenType) {
	msg := fmt.Sprintf("Line %d: Expected next token to be %s, got %s instead", p.peekToken.Line, t, p.peekToken.Type)
	p.errors = append(p.errors, msg)
}

func (p *Parser) addErrorWithContext(msg string, line int) {
	contextMsg := fmt.Sprintf("Line %d: %s", line, msg)
	p.errors = append(p.errors, contextMsg)
}

// Synchronize parser after error - skip to next statement boundary
func (p *Parser) synchronize() {
	p.nextToken()

	for p.curToken.Type != token.EOF {
		if p.curToken.Type == token.SEMICOLON {
			p.nextToken()
			return
		}

		switch p.curToken.Type {
		case token.LET, token.CONST, token.FUNCTION, token.IF, token.WHILE,
			token.FOR, token.RETURN, token.MATCH, token.SWITCH:
			return
		}

		p.nextToken()
	}
}

// parse expressions

func (p *Parser) parseExpressionStatement() *ast.ExpressionStatement {
	stmt := &ast.ExpressionStatement{Token: p.curToken}

	stmt.Expression = p.parseExpression(LOWEST)

	if p.peekTokenIs(token.SEMICOLON) {
		p.nextToken()
	}

	return stmt
}

func (p *Parser) parseExpression(precedence int) ast.Expression {
	postfix := p.postfixParseFns[p.curToken.Type]
	if postfix != nil {
		return (postfix())
	}
	prefix := p.prefixParseFns[p.curToken.Type]
	if prefix == nil {
		p.noPrefixParseFnError(p.curToken.Type)
		return nil
	}
	leftExp := prefix()

	for !p.peekTokenIs(token.SEMICOLON) && precedence < p.peekPrecedence() {
		infix := p.infixParseFns[p.peekToken.Type]
		if infix == nil {
			p.noInfixParseFnError(p.peekToken.Type)
			return nil
		}

		p.nextToken()
		leftExp = infix(leftExp)
	}
	return leftExp

}

// prefix expressions

func (p *Parser) parsePrefixExpression() ast.Expression {
	expression := &ast.PrefixExpression{
		Token:    p.curToken,
		Operator: p.curToken.Literal,
	}

	p.nextToken()

	expression.Right = p.parseExpression(PREFIX)

	return expression
}

func (p *Parser) noPrefixParseFnError(t token.TokenType) {
	msg := fmt.Sprintf("Line %d: Unexpected token '%s' - this token cannot start an expression. Check for missing operands or invalid syntax", p.curToken.Line, t)
	p.errors = append(p.errors, msg)
}

// infix expressions

func (p *Parser) parseInfixExpression(left ast.Expression) ast.Expression {
	expression := &ast.InfixExpression{
		Token:    p.curToken,
		Operator: p.curToken.Literal,
		Left:     left,
	}

	precedence := p.curPrecedence()
	p.nextToken()
	expression.Right = p.parseExpression(precedence)
	return expression
}

func (p *Parser) parseRangeExpression(left ast.Expression) ast.Expression {
	expression := &ast.RangeExpression{
		Token: p.curToken,
		Start: left,
	}

	precedence := p.curPrecedence()
	p.nextToken()
	expression.End = p.parseExpression(precedence)
	return expression
}

func (p *Parser) noInfixParseFnError(t token.TokenType) {
	msg := fmt.Sprintf("Line %d: Unexpected token '%s' - cannot be used in this context. Check syntax for operators, expressions, or missing semicolon", p.curToken.Line, t)
	p.errors = append(p.errors, msg)
}

func (p *Parser) parseGroupedExpression() ast.Expression {
	p.nextToken()

	exp := p.parseExpression(LOWEST)

	if !p.expectPeek(token.RPAREN) {
		return nil
	}

	return exp
}

// postfix expressions

func (p *Parser) parsePostfixExpression() ast.Expression {
	expression := &ast.PostfixExpression{
		Token:    p.prevToken,
		Operator: p.curToken.Literal,
	}
	return expression
}

func (p *Parser) parseTodoStatement() ast.Expression {
	stmt := &ast.TodoStatement{Token: p.curToken}
	p.nextToken()
	stmt.Value = p.parseExpression(LOWEST)
	if p.peekTokenIs(token.SEMICOLON) {
		p.nextToken()
	}
	return stmt
}

func (p *Parser) parseWarnStatement() ast.Expression {
	stmt := &ast.WarnStatement{Token: p.curToken}
	p.nextToken()
	stmt.Value = p.parseExpression(LOWEST)
	if p.peekTokenIs(token.SEMICOLON) {
		p.nextToken()
	}
	return stmt
}

func (p *Parser) parseErrorStatement() ast.Expression {
	errorToken := p.curToken
	p.nextToken()

	// Check if this is an error type declaration: error TypeName(param1, param2)
	if p.curTokenIs(token.IDENT) && p.peekTokenIs(token.LPAREN) {
		// This is an error type declaration
		return p.parseErrorDeclaration(errorToken)
	}

	// This is a regular declarative error statement: error "message"
	stmt := &ast.ErrorStatement{Token: errorToken}
	stmt.Value = p.parseExpression(LOWEST)
	if p.peekTokenIs(token.SEMICOLON) {
		p.nextToken()
	}
	return stmt
}

func (p *Parser) parseErrorDeclaration(errorToken token.Token) ast.Expression {
	decl := &ast.ErrorDeclaration{Token: errorToken}

	// Parse the error type name (current token should be IDENT)
	if !p.curTokenIs(token.IDENT) {
		p.addError(fmt.Sprintf("expected identifier for error type name, got %s", p.curToken.Type))
		return nil
	}

	decl.Name = &ast.Identifier{Token: p.curToken, Value: p.curToken.Literal}

	// Expect opening parenthesis
	if !p.expectPeek(token.LPAREN) {
		return nil
	}

	// Parse parameters
	decl.Parameters = []*ast.Identifier{}

	if !p.peekTokenIs(token.RPAREN) {
		p.nextToken() // advance to first parameter

		// Parse first parameter
		if !p.curTokenIs(token.IDENT) {
			p.addError(fmt.Sprintf("expected parameter name, got %s", p.curToken.Type))
			return nil
		}
		decl.Parameters = append(decl.Parameters, &ast.Identifier{Token: p.curToken, Value: p.curToken.Literal})

		// Parse additional parameters
		for p.peekTokenIs(token.COMMA) {
			p.nextToken() // advance to comma
			p.nextToken() // advance to next parameter
			if !p.curTokenIs(token.IDENT) {
				p.addError(fmt.Sprintf("expected parameter name, got %s", p.curToken.Type))
				return nil
			}
			decl.Parameters = append(decl.Parameters, &ast.Identifier{Token: p.curToken, Value: p.curToken.Literal})
		}
	}

	// Expect closing parenthesis
	if !p.expectPeek(token.RPAREN) {
		return nil
	}

	return decl
}

func (p *Parser) parseThrowStatement() ast.Expression {
	stmt := &ast.ThrowStatement{Token: p.curToken}
	p.nextToken()

	stmt.ErrorExpr = p.parseExpression(LOWEST)

	if p.peekTokenIs(token.SEMICOLON) {
		p.nextToken()
	}

	return stmt
}

func (p *Parser) parseDeferStatement() ast.Expression {
	stmt := &ast.DeferStatement{Token: p.curToken}
	p.nextToken()
	stmt.Call = p.parseExpression(LOWEST)
	if p.peekTokenIs(token.SEMICOLON) {
		p.nextToken()
	}
	return stmt
}

func (p *Parser) parseInfoStatement() ast.Expression {
	stmt := &ast.InfoStatement{Token: p.curToken}
	p.nextToken()
	stmt.Value = p.parseExpression(LOWEST)
	if p.peekTokenIs(token.SEMICOLON) {
		p.nextToken()
	}
	return stmt
}

func (p *Parser) parseDebugStatement() ast.Expression {
	stmt := &ast.DebugStatement{Token: p.curToken}
	p.nextToken()
	stmt.Value = p.parseExpression(LOWEST)
	if p.peekTokenIs(token.SEMICOLON) {
		p.nextToken()
	}
	return stmt
}

func (p *Parser) parseNoteStatement() ast.Expression {
	stmt := &ast.NoteStatement{Token: p.curToken}
	p.nextToken()
	stmt.Value = p.parseExpression(LOWEST)
	if p.peekTokenIs(token.SEMICOLON) {
		p.nextToken()
	}
	return stmt
}

func (p *Parser) parseSuccessStatement() ast.Expression {
	stmt := &ast.SuccessStatement{Token: p.curToken}
	p.nextToken()
	stmt.Value = p.parseExpression(LOWEST)
	if p.peekTokenIs(token.SEMICOLON) {
		p.nextToken()
	}
	return stmt
}

func (p *Parser) parseTraceStatement() ast.Expression {
	stmt := &ast.TraceStatement{Token: p.curToken}
	p.nextToken()
	stmt.Value = p.parseExpression(LOWEST)
	if p.peekTokenIs(token.SEMICOLON) {
		p.nextToken()
	}
	return stmt
}

func (p *Parser) parseFatalStatement() ast.Expression {
	stmt := &ast.FatalStatement{Token: p.curToken}
	p.nextToken()
	stmt.Value = p.parseExpression(LOWEST)
	if p.peekTokenIs(token.SEMICOLON) {
		p.nextToken()
	}
	return stmt
}

func (p *Parser) parseCriticalStatement() ast.Expression {
	stmt := &ast.CriticalStatement{Token: p.curToken}
	p.nextToken()
	stmt.Value = p.parseExpression(LOWEST)
	if p.peekTokenIs(token.SEMICOLON) {
		p.nextToken()
	}
	return stmt
}

func (p *Parser) parseLogStatement() ast.Expression {
	stmt := &ast.LogStatement{Token: p.curToken}
	p.nextToken()
	stmt.Value = p.parseExpression(LOWEST)
	if p.peekTokenIs(token.SEMICOLON) {
		p.nextToken()
	}
	return stmt
}

func (p *Parser) parseRepeatStatement() ast.Expression {
	stmt := &ast.RepeatStatement{Token: p.curToken, VarName: "i"}
	// Check for optional (varname)
	if p.peekTokenIs(token.LPAREN) {
		p.nextToken() // consume '('
		p.nextToken() // move to identifier
		if p.curToken.Type == token.IDENT {
			stmt.VarName = p.curToken.Literal
		} else {
			return nil // syntax error
		}
		if !p.expectPeek(token.RPAREN) {
			return nil
		}
		// Do NOT call p.nextToken() here; the next token should be the count expression
	} else {
		p.nextToken()
	}
	stmt.Count = p.parseExpression(LOWEST)
	if !p.expectPeek(token.LBRACE) {
		return nil
	}
	stmt.Block = p.parseBlockStatement()
	return stmt
}
