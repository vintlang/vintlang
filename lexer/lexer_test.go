package lexer

import (
	"testing"

	"github.com/ekilie/vint-lang/token"
)

func TestNextToken(t *testing.T) {
	input := `
	
	let five = 5;
	let ten = 10;

	let sum = func(x, y){
	x + y;
	};

	let ans = sum(five, ten);

	!-/5;
	5 < 10 > 5;

	if (5 < 10) {
		return true;
	} else {
		return false;
	}

	10 == 10;
	10 != 9; // This is a comment
	// Comment 

	/*
	multiline comment
	*/

	/* multiline comment number twooooooooooo */
	5
	`

	tests := []struct {
		expectedType    token.TokenType
		expectedLiteral string
	}{
		{token.LET, "let"},
		{token.IDENT, "five"},
		{token.ASSIGN, "="},
		{token.INT, "5"},
		{token.SEMICOLON, ";"},
		{token.LET, "let"},
		{token.IDENT, "ten"},
		{token.ASSIGN, "="},
		{token.INT, "10"},
		{token.SEMICOLON, ";"},
		{token.LET, "let"},
		{token.IDENT, "sum"},
		{token.ASSIGN, "="},
		{token.FUNCTION, "func"},
		{token.LPAREN, "("},
		{token.IDENT, "x"},
		{token.COMMA, ","},
		{token.IDENT, "y"},
		{token.RPAREN, ")"},
		{token.LBRACE, "{"},
		{token.IDENT, "x"},
		{token.PLUS, "+"},
		{token.IDENT, "y"},
		{token.SEMICOLON, ";"},
		{token.RBRACE, "}"},
		{token.SEMICOLON, ";"},
		{token.LET, "let"},
		{token.IDENT, "ans"},
		{token.ASSIGN, "="},
		{token.IDENT, "sum"},
		{token.LPAREN, "("},
		{token.IDENT, "five"},
		{token.COMMA, ","},
		{token.IDENT, "ten"},
		{token.RPAREN, ")"},
		{token.SEMICOLON, ";"},
		{token.BANG, "!"},
		{token.MINUS, "-"},
		{token.SLASH, "/"},
		{token.INT, "5"},
		{token.SEMICOLON, ";"},
		{token.INT, "5"},
		{token.LT, "<"},
		{token.INT, "10"},
		{token.GT, ">"},
		{token.INT, "5"},
		{token.SEMICOLON, ";"},
		{token.IF, "if"},
		{token.LPAREN, "("},
		{token.INT, "5"},
		{token.LT, "<"},
		{token.INT, "10"},
		{token.RPAREN, ")"},
		{token.LBRACE, "{"},
		{token.RETURN, "return"},
		{token.TRUE, "true"},
		{token.SEMICOLON, ";"},
		{token.RBRACE, "}"},
		{token.ELSE, "sivyo"},
		{token.LBRACE, "{"},
		{token.RETURN, "return"},
		{token.FALSE, "false"},
		{token.SEMICOLON, ";"},
		{token.RBRACE, "}"},
		{token.INT, "10"},
		{token.EQ, "=="},
		{token.INT, "10"},
		{token.SEMICOLON, ";"},
		{token.INT, "10"},
		{token.NOT_EQ, "!="},
		{token.INT, "9"},
		{token.SEMICOLON, ";"},
		{token.INT, "5"},
		{token.STRING, "bangi"},
		{token.STRING, "ba ngi"},
		{token.LBRACKET, "["},
		{token.INT, "1"},
		{token.COMMA, ","},
		{token.INT, "2"},
		{token.RBRACKET, "]"},
		{token.SEMICOLON, ";"},
		{token.LBRACE, "{"},
		{token.STRING, "mambo"},
		{token.COLON, ":"},
		{token.STRING, "vipi"},
		{token.RBRACE, "}"},
		{token.DOT, "."},
		{token.IMPORT, "tumia"},
		{token.IDENT, "muda"},
		{token.SWITCH, "switch"},
		{token.LPAREN, "("},
		{token.IDENT, "a"},
		{token.RPAREN, ")"},
		{token.LBRACE, "{"},
		{token.CASE, "ikiwa"},
		{token.INT, "2"},
		{token.LBRACE, "{"},
		{token.IDENT, "print"},
		{token.LPAREN, "("},
		{token.INT, "2"},
		{token.RPAREN, ")"},
		{token.RBRACE, "}"},
		{token.DEFAULT, "kawaida"},
		{token.LBRACE, "{"},
		{token.IDENT, "print"},
		{token.LPAREN, "("},
		{token.INT, "0"},
		{token.RPAREN, ")"},
		{token.RBRACE, "}"},
		{token.RBRACE, "}"},
		{token.NULL, "tupu"},
		{token.FOR, "kwa"},
		{token.IDENT, "i"},
		{token.COMMA, ","},
		{token.IDENT, "v"},
		{token.IN, "ktk"},
		{token.IDENT, "j"},
		{token.EOF, ""},
	}

	l := New(input)

	for i, tt := range tests {
		tok := l.NextToken()

		if tok.Type != tt.expectedType {
			t.Fatalf("tests[%d] - tokentype wrong. expected=%q, got=%q",
				i, tt.expectedType, tok.Type)
		}

		if tok.Literal != tt.expectedLiteral {
			t.Fatalf("tests[%d] - literal wrong. expected=%q, got=%q",
				i, tt.expectedLiteral, tok.Literal)
		}
	}
}
