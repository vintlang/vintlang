package lexer

import (
	"fmt"
	"strings"

	"github.com/vintlang/vintlang/token"
)

type Lexer struct {
	input        []rune
	position     int
	readPosition int
	ch           rune
	line         int
	column       int
	filename     string
	errors       []string
}

func New(input string) *Lexer {
	l := &Lexer{
		input:    []rune(input),
		line:     1,
		column:   1,
		filename: "main.vint", // default filename
	}
	l.readChar()
	return l
}

func NewWithFilename(input string, filename string) *Lexer {
	l := &Lexer{
		input:    []rune(input),
		line:     1,
		column:   1,
		filename: filename,
	}
	l.readChar()
	return l
}

func (l *Lexer) readChar() {
	if l.readPosition >= len(l.input) {
		l.ch = 0
	} else {
		l.ch = l.input[l.readPosition]
	}

	l.position = l.readPosition
	l.readPosition += 1

	if l.ch == '\n' {
		l.line += 1
		l.column = 0
	} else {
		l.column += 1
	}
}

func (l *Lexer) NextToken() token.Token {
	var tok token.Token
	l.skipWhitespace()
	if l.ch == rune('/') && l.peekChar() == rune('/') {
		l.skipSingleLineComment()
		return l.NextToken()
	}
	if l.ch == rune('/') && l.peekChar() == rune('*') {
		l.skipMultiLineComment()
		return l.NextToken()
	}

	switch l.ch {
	case rune('='):
		if l.peekChar() == rune('=') {
			ch := l.ch
			l.readChar()
			tok = token.Token{Type: token.EQ, Literal: string(ch) + string(l.ch), Line: l.line}
		} else if l.peekChar() == rune('>') {
			ch := l.ch
			l.readChar()
			tok = token.Token{Type: token.ARROW, Literal: string(ch) + string(l.ch), Line: l.line}
		} else {
			tok = newToken(token.ASSIGN, l.line, l.ch)
		}
	case rune(';'):
		tok = newToken(token.SEMICOLON, l.line, l.ch)
	case rune('('):
		tok = newToken(token.LPAREN, l.line, l.ch)
	case rune(')'):
		tok = newToken(token.RPAREN, l.line, l.ch)
	case rune('{'):
		tok = newToken(token.LBRACE, l.line, l.ch)
	case rune('}'):
		tok = newToken(token.RBRACE, l.line, l.ch)
	case rune(','):
		tok = newToken(token.COMMA, l.line, l.ch)
	case rune('+'):
		if l.peekChar() == rune('=') {
			ch := l.ch
			l.readChar()
			tok = token.Token{Type: token.PLUS_ASSIGN, Line: l.line, Literal: string(ch) + string(l.ch)}
		} else if l.peekChar() == rune('+') {
			ch := l.ch
			l.readChar()
			tok = token.Token{Type: token.PLUS_PLUS, Literal: string(ch) + string(l.ch), Line: l.line}
		} else {
			tok = newToken(token.PLUS, l.line, l.ch)
		}
	case rune('-'):
		if l.peekChar() == rune('=') {
			ch := l.ch
			l.readChar()
			tok = token.Token{Type: token.MINUS_ASSIGN, Line: l.line, Literal: string(ch) + string(l.ch)}
		} else if l.peekChar() == rune('-') {
			ch := l.ch
			l.readChar()
			tok = token.Token{Type: token.MINUS_MINUS, Literal: string(ch) + string(l.ch), Line: l.line}
		} else {
			tok = newToken(token.MINUS, l.line, l.ch)
		}
	case rune('!'):
		if l.peekChar() == rune('=') {
			ch := l.ch
			l.readChar()
			tok = token.Token{Type: token.NOT_EQ, Literal: string(ch) + string(l.ch), Line: l.line}
		} else {
			tok = newToken(token.BANG, l.line, l.ch)
		}
	case rune('/'):
		if l.peekChar() == rune('=') {
			ch := l.ch
			l.readChar()
			tok = token.Token{Type: token.SLASH_ASSIGN, Line: l.line, Literal: string(ch) + string(l.ch)}
		} else {
			tok = newToken(token.SLASH, l.line, l.ch)
		}
	case rune('*'):
		if l.peekChar() == rune('=') {
			ch := l.ch
			l.readChar()
			tok = token.Token{Type: token.ASTERISK_ASSIGN, Line: l.line, Literal: string(ch) + string(l.ch)}
		} else if l.peekChar() == rune('*') {
			ch := l.ch
			l.readChar()
			tok = token.Token{Type: token.POW, Literal: string(ch) + string(l.ch), Line: l.line}
		} else {
			tok = newToken(token.ASTERISK, l.line, l.ch)
		}
	case rune('<'):
		if l.peekChar() == rune('=') {
			ch := l.ch
			l.readChar()
			tok = token.Token{Type: token.LTE, Literal: string(ch) + string(l.ch), Line: l.line}
		} else {
			tok = newToken(token.LT, l.line, l.ch)
		}
	case rune('>'):
		if l.peekChar() == rune('=') {
			ch := l.ch
			l.readChar()
			tok = token.Token{Type: token.GTE, Literal: string(ch) + string(l.ch), Line: l.line}
		} else {
			tok = newToken(token.GT, l.line, l.ch)
		}

	case rune('"'):
		tok.Type = token.STRING
		tok.Literal = l.readString()
		tok.Line = l.line
	case rune('\''):
		tok = token.Token{Type: token.STRING, Literal: l.readSingleQuoteString(), Line: l.line}
	case rune('['):
		tok = newToken(token.LBRACKET, l.line, l.ch)
	case rune(']'):
		tok = newToken(token.RBRACKET, l.line, l.ch)
	case rune(':'):
		tok = newToken(token.COLON, l.line, l.ch)
	case rune('@'):
		tok = newToken(token.AT, l.line, l.ch)
	case rune('.'):
		if l.peekChar() == rune('.') {
			ch := l.ch
			l.readChar()
			tok = token.Token{Type: token.RANGE, Literal: string(ch) + string(l.ch), Line: l.line}
		} else {
			tok = newToken(token.DOT, l.line, l.ch)
		}
	case rune('&'):
		if l.peekChar() == rune('&') {
			ch := l.ch
			l.readChar()
			tok = token.Token{Type: token.AND, Literal: string(ch) + string(l.ch), Line: l.line}
		} else {
			tok = newToken(token.AMPERSAND, l.line, l.ch)
		}
	case rune('|'):
		if l.peekChar() == rune('|') {
			ch := l.ch
			l.readChar()
			tok = token.Token{Type: token.OR, Literal: string(ch) + string(l.ch), Line: l.line}
		}
	case rune('%'):
		if l.peekChar() == rune('=') {
			ch := l.ch
			l.readChar()
			tok = token.Token{Type: token.MODULUS_ASSIGN, Line: l.line, Literal: string(ch) + string(l.ch)}
		} else {
			tok = newToken(token.MODULUS, l.line, l.ch)
		}
	case rune('?'):
		if l.peekChar() == rune('?') {
			ch := l.ch
			l.readChar()
			tok = token.Token{Type: token.NULL_COALESCE, Literal: string(ch) + string(l.ch), Line: l.line}
		} else {
			tok = l.createIllegalToken(l.ch, "- single '?' is not a valid operator, did you mean '??'?")
		}
	case rune('#'):
		if l.peekChar() == rune('!') && l.line == 1 {
			l.skipSingleLineComment()
			return l.NextToken()
		}
	case 0:
		tok.Literal = ""
		tok.Type = token.EOF
		tok.Line = l.line
	default:
		if isLetter(l.ch) {
			tok.Literal = l.readIdentifier()
			// Special handling for import() - check if followed by (
			if tok.Literal == "import" {
				// We look ahead to see if this is import( function call
				if l.peekNextNonWhitespace() == '(' {
					tok.Type = token.IDENT // We Treat as identifier for function call
				} else {
					tok.Type = token.IMPORT // We Treat as import statement keyword
				}
			} else {
				tok.Type = token.LookupIdent(tok.Literal)
			}
			tok.Line = l.line
			return tok
		} else if isDigit(l.ch) && isLetter(l.peekChar()) {
			tok.Literal = l.readIdentifier()
			tok.Type = token.LookupIdent(tok.Literal)
			tok.Line = l.line
			return tok
		} else if isDigit(l.ch) {
			tok = l.readDecimal()
			return tok
		} else {
			tok = l.createIllegalToken(l.ch, "- unexpected character, not a valid token")
		}
	}

	l.readChar()
	return tok
}

func newToken(tokenType token.TokenType, line int, ch rune) token.Token {
	return token.Token{Type: tokenType, Literal: string(ch), Line: line}
}

func (l *Lexer) newTokenWithColumn(tokenType token.TokenType, literal string) token.Token {
	return token.Token{Type: tokenType, Literal: literal, Line: l.line, Column: l.column}
}

// Errors returns all lexer errors
func (l *Lexer) Errors() []string {
	return l.errors
}

// addError adds an error message to the lexer
func (l *Lexer) addError(msg string) {
	l.errors = append(l.errors, msg)
}

// getSourceLine returns the source line for error context
func (l *Lexer) getSourceLine(lineNum int) string {
	lines := strings.Split(string(l.input), "\n")
	if lineNum > 0 && lineNum <= len(lines) {
		return lines[lineNum-1]
	}
	return ""
}

// addErrorWithContext adds an error with source code context
func (l *Lexer) addErrorWithContext(msg string, line, column int) {
	sourceLine := l.getSourceLine(line)
	if sourceLine != "" {
		contextMsg := fmt.Sprintf("%s\n    %s\n    %s^", msg, sourceLine, strings.Repeat(" ", column-1))
		l.errors = append(l.errors, contextMsg)
	} else {
		l.errors = append(l.errors, msg)
	}
}

// createIllegalToken creates an ILLEGAL token and adds a descriptive error message
func (l *Lexer) createIllegalToken(ch rune, context string) token.Token {
	var errorMsg string
	if ch == 0 {
		errorMsg = fmt.Sprintf("%s:%d:%d: Unexpected end of file", l.filename, l.line, l.column)
	} else if ch < 32 || ch > 126 {
		errorMsg = fmt.Sprintf("%s:%d:%d: Illegal character '\\x%02x' (non-printable) %s", l.filename, l.line, l.column, ch, context)
	} else {
		errorMsg = fmt.Sprintf("%s:%d:%d: Illegal character '%c' %s", l.filename, l.line, l.column, ch, context)
	}
	l.addErrorWithContext(errorMsg, l.line, l.column)
	return token.Token{Type: token.ILLEGAL, Literal: string(ch), Line: l.line, Column: l.column}
}

func (l *Lexer) readIdentifier() string {
	position := l.position

	for isLetter(l.ch) || isDigit(l.ch) {
		l.readChar()
	}
	return string(l.input[position:l.position])
}

func isLetter(ch rune) bool {
	return 'a' <= ch && ch <= 'z' || 'A' <= ch && ch <= 'Z' || ch == '_' || ch == '@'
}

func (l *Lexer) skipWhitespace() {
	for l.ch == ' ' || l.ch == '\t' || l.ch == '\n' || l.ch == '\r' {
		if l.ch == '\n' {
			l.line++
		}
		l.readChar()
	}
}

func (l *Lexer) peekNextNonWhitespace() rune {
	// Save current position
	savedPosition := l.position
	savedReadPosition := l.readPosition
	savedCh := l.ch
	savedLine := l.line
	savedColumn := l.column

	// Skip whitespace and look for next non-whitespace character
	for l.ch == ' ' || l.ch == '\t' || l.ch == '\n' || l.ch == '\r' {
		l.readChar()
	}

	result := l.ch

	// Restore position
	l.position = savedPosition
	l.readPosition = savedReadPosition
	l.ch = savedCh
	l.line = savedLine
	l.column = savedColumn

	return result
}

func isDigit(ch rune) bool {
	return '0' <= ch && ch <= '9'
}

func (l *Lexer) readNumber() string {
	position := l.position
	for isDigit(l.ch) {
		l.readChar()
	}
	return string(l.input[position:l.position])
}

func (l *Lexer) readDecimal() token.Token {
	integer := l.readNumber()
	if l.ch == '.' && isDigit(l.peekChar()) {
		l.readChar()
		fraction := l.readNumber()
		return token.Token{Type: token.FLOAT, Literal: integer + "." + fraction, Line: l.line}
	}
	return token.Token{Type: token.INT, Literal: integer, Line: l.line}
}

func (l *Lexer) peekChar() rune {
	if l.readPosition >= len(l.input) {
		return rune(0)
	} else {
		return l.input[l.readPosition]
	}
}

// func (l *Lexer) peekTwoChar() rune {
// 	if l.readPosition+1 >= len(l.input) {
// 		return rune(0)
// 	} else {
// 		return l.input[l.readPosition+1]
// 	}
// }

func (l *Lexer) skipSingleLineComment() {
	for l.ch != '\n' && l.ch != 0 {
		l.readChar()
	}
	l.skipWhitespace()
}

func (l *Lexer) skipMultiLineComment() {
	endFound := false

	for !endFound {
		if l.ch == 0 {
			endFound = true
		}

		if l.ch == '*' && l.peekChar() == '/' {
			endFound = true
			l.readChar()
		}

		l.readChar()
		l.skipWhitespace()
	}

}

func (l *Lexer) readString() string {
	startLine := l.line
	var str strings.Builder
	for {
		l.readChar()
		if l.ch == '"' {
			break
		} else if l.ch == 0 {
			l.addError(fmt.Sprintf("Line %d: Unterminated string literal started on line %d", l.line, startLine))
			break
		} else if l.ch == '\\' {
			switch l.peekChar() {
			case 'n':
				l.readChar()
				str.WriteByte('\n')
			case 'r':
				l.readChar()
				str.WriteByte('\r')
			case 't':
				l.readChar()
				str.WriteByte('\t')
			case '"':
				l.readChar()
				str.WriteByte('"')
			case '\\':
				l.readChar()
				str.WriteByte('\\')
			case '0':
				l.readChar()
				str.WriteByte('\x00')
			case 'x':
				// Handle hex escape sequences \xHH
				l.readChar() // consume 'x'
				l.readChar() // get first hex digit
				h1 := l.ch
				l.readChar() // get second hex digit
				h2 := l.ch
				if isHexDigit(h1) && isHexDigit(h2) {
					value := hexValue(h1)*16 + hexValue(h2)
					str.WriteByte(byte(value))
				} else {
					// Invalid hex sequence, just include as-is
					str.WriteString("\\x")
					str.WriteRune(h1)
					str.WriteRune(h2)
				}
			case 'u':
				// Handle Unicode escape sequences \uHHHH
				l.readChar() // consume 'u'
				var hexDigits [4]rune
				for i := 0; i < 4; i++ {
					l.readChar()
					hexDigits[i] = l.ch
					if !isHexDigit(l.ch) {
						// Invalid Unicode sequence, include as-is
						str.WriteString("\\u")
						for j := 0; j <= i; j++ {
							str.WriteRune(hexDigits[j])
						}
						goto continueLoop
					}
				}
				// Convert to Unicode code point
				value := hexValue(hexDigits[0])*4096 + hexValue(hexDigits[1])*256 +
					hexValue(hexDigits[2])*16 + hexValue(hexDigits[3])
				str.WriteRune(rune(value))
			default:
				// Unknown escape sequence, keep the backslash
				str.WriteByte('\\')
				str.WriteRune(l.peekChar())
				l.readChar()
			}
		} else {
			str.WriteRune(l.ch)
		}
	continueLoop:
	}
	return str.String()
}

func (l *Lexer) readSingleQuoteString() string {
	startLine := l.line
	var str string
	for {
		l.readChar()
		if l.ch == '\'' {
			break
		} else if l.ch == 0 {
			l.addError(fmt.Sprintf("Line %d: Unterminated single-quoted string literal started on line %d", l.line, startLine))
			break
		} else if l.ch == '\\' {
			switch l.peekChar() {
			case 'n':
				l.readChar()
				l.ch = '\n'
			case 'r':
				l.readChar()
				l.ch = '\r'
			case 't':
				l.readChar()
				l.ch = '\t'
			case '"':
				l.readChar()
				l.ch = '"'
			case '\\':
				l.readChar()
				l.ch = '\\'
			}
		}
		str += string(l.ch)
	}
	return str
}

// Helper functions for hex digit processing
func isHexDigit(ch rune) bool {
	return isDigit(ch) || (ch >= 'a' && ch <= 'f') || (ch >= 'A' && ch <= 'F')
}

func hexValue(ch rune) int {
	if isDigit(ch) {
		return int(ch - '0')
	} else if ch >= 'a' && ch <= 'f' {
		return int(ch - 'a' + 10)
	} else if ch >= 'A' && ch <= 'F' {
		return int(ch - 'A' + 10)
	}
	return 0
}

// Smart suggestions for common typos and mistakes
var commonKeywords = []string{
	"let", "const", "if", "else", "for", "while", "func", "return",
	"true", "false", "null", "import", "switch", "case", "default",
	"break", "continue", "defer", "match", "in",
}

// levenshteinDistance calculates the edit distance between two strings
func levenshteinDistance(a, b string) int {
	if len(a) == 0 {
		return len(b)
	}
	if len(b) == 0 {
		return len(a)
	}

	matrix := make([][]int, len(a)+1)
	for i := range matrix {
		matrix[i] = make([]int, len(b)+1)
		matrix[i][0] = i
	}
	for j := 0; j <= len(b); j++ {
		matrix[0][j] = j
	}

	for i := 1; i <= len(a); i++ {
		for j := 1; j <= len(b); j++ {
			cost := 0
			if a[i-1] != b[j-1] {
				cost = 1
			}
			matrix[i][j] = min(
				matrix[i-1][j]+1,      // deletion
				matrix[i][j-1]+1,      // insertion
				matrix[i-1][j-1]+cost, // substitution
			)
		}
	}
	return matrix[len(a)][len(b)]
}

func min(a, b, c int) int {
	if a < b && a < c {
		return a
	}
	if b < c {
		return b
	}
	return c
}

// getSuggestion returns a "did you mean?" suggestion for unknown identifiers
func getSuggestion(input string) string {
	bestMatch := ""
	bestDistance := 3 // Only suggest if distance <= 3

	for _, keyword := range commonKeywords {
		distance := levenshteinDistance(input, keyword)
		if distance < bestDistance {
			bestDistance = distance
			bestMatch = keyword
		}
	}

	if bestMatch != "" {
		return fmt.Sprintf(" (did you mean '%s'?)", bestMatch)
	}
	return ""
}

func (l *Lexer) GetErrors() []string {
	return l.errors
}

func (l *Lexer) GetFilename() string {
	return l.filename
}
