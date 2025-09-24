package lexer

import (
	"testing"

	"github.com/vintlang/vintlang/token"
)

func TestNextToken(t *testing.T) {
	input := `let five = 5;
let ten = 10;

let add = func(x, y) {
	x + y;
};

let result = add(five, ten);
!-/+5;
5 < 10 > 5;

if (5 < 10) {
	return true;
} else {
	return false;
}

10 == 10;
10 != 9;
"foobar"
"foo bar"
[1, 2];
{"foo": "bar"}
3.14;
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
		{token.IDENT, "add"},
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
		{token.IDENT, "result"},
		{token.ASSIGN, "="},
		{token.IDENT, "add"},
		{token.LPAREN, "("},
		{token.IDENT, "five"},
		{token.COMMA, ","},
		{token.IDENT, "ten"},
		{token.RPAREN, ")"},
		{token.SEMICOLON, ";"},
		{token.BANG, "!"},
		{token.MINUS, "-"},
		{token.SLASH, "/"},
		{token.PLUS, "+"},
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
		{token.ELSE, "else"},
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
		{token.STRING, "foobar"},
		{token.STRING, "foo bar"},
		{token.LBRACKET, "["},
		{token.INT, "1"},
		{token.COMMA, ","},
		{token.INT, "2"},
		{token.RBRACKET, "]"},
		{token.SEMICOLON, ";"},
		{token.LBRACE, "{"},
		{token.STRING, "foo"},
		{token.COLON, ":"},
		{token.STRING, "bar"},
		{token.RBRACE, "}"},
		{token.FLOAT, "3.14"},
		{token.SEMICOLON, ";"},
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

func TestVintLangSpecificTokens(t *testing.T) {
	input := `
	// VintLang specific keywords and operators
	const PI = 3.14;
	while (true) {
		break;
		continue;
	}
	for x in [1, 2, 3] {
		print(x);
	}
	match x {
		case 1 => "one";
		default => "other";
	}
	import time;
	error MyError(msg);
	throw new MyError("test");
	defer cleanup();
	todo "implement this";
	warn "deprecated";
	1..10
	x += 5;
	y++;
	`

	tests := []struct {
		expectedType    token.TokenType
		expectedLiteral string
	}{
		{token.CONST, "const"},
		{token.IDENT, "PI"},
		{token.ASSIGN, "="},
		{token.FLOAT, "3.14"},
		{token.SEMICOLON, ";"},
		{token.WHILE, "while"},
		{token.LPAREN, "("},
		{token.TRUE, "true"},
		{token.RPAREN, ")"},
		{token.LBRACE, "{"},
		{token.BREAK, "break"},
		{token.SEMICOLON, ";"},
		{token.CONTINUE, "continue"},
		{token.SEMICOLON, ";"},
		{token.RBRACE, "}"},
		{token.FOR, "for"},
		{token.IDENT, "x"},
		{token.IN, "in"},
		{token.LBRACKET, "["},
		{token.INT, "1"},
		{token.COMMA, ","},
		{token.INT, "2"},
		{token.COMMA, ","},
		{token.INT, "3"},
		{token.RBRACKET, "]"},
		{token.LBRACE, "{"},
		{token.IDENT, "print"},
		{token.LPAREN, "("},
		{token.IDENT, "x"},
		{token.RPAREN, ")"},
		{token.SEMICOLON, ";"},
		{token.RBRACE, "}"},
		{token.MATCH, "match"},
		{token.IDENT, "x"},
		{token.LBRACE, "{"},
		{token.CASE, "case"},
		{token.INT, "1"},
		{token.ARROW, "=>"},
		{token.STRING, "one"},
		{token.SEMICOLON, ";"},
		{token.DEFAULT, "default"},
		{token.ARROW, "=>"},
		{token.STRING, "other"},
		{token.SEMICOLON, ";"},
		{token.RBRACE, "}"},
		{token.IMPORT, "import"},
		{token.IDENT, "time"},
		{token.SEMICOLON, ";"},
		{token.ERROR, "error"},
		{token.IDENT, "MyError"},
		{token.LPAREN, "("},
		{token.IDENT, "msg"},
		{token.RPAREN, ")"},
		{token.SEMICOLON, ";"},
		{token.THROW, "throw"},
		{token.IDENT, "new"},
		{token.IDENT, "MyError"},
		{token.LPAREN, "("},
		{token.STRING, "test"},
		{token.RPAREN, ")"},
		{token.SEMICOLON, ";"},
		{token.DEFER, "defer"},
		{token.IDENT, "cleanup"},
		{token.LPAREN, "("},
		{token.RPAREN, ")"},
		{token.SEMICOLON, ";"},
		{token.TODO, "todo"},
		{token.STRING, "implement this"},
		{token.SEMICOLON, ";"},
		{token.WARN, "warn"},
		{token.STRING, "deprecated"},
		{token.SEMICOLON, ";"},
		{token.INT, "1"},
		{token.RANGE, ".."},
		{token.INT, "10"},
		{token.IDENT, "x"},
		{token.PLUS_ASSIGN, "+="},
		{token.INT, "5"},
		{token.SEMICOLON, ";"},
		{token.IDENT, "y"},
		{token.PLUS_PLUS, "++"},
		{token.SEMICOLON, ";"},
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

func TestStringTokens(t *testing.T) {
	tests := []struct {
		input           string
		expectedLiteral string
	}{
		{`"hello"`, "hello"},
		{`"hello world"`, "hello world"},
		{`"hello\nworld"`, "hello\nworld"},
		{`"hello\"world"`, "hello\"world"},
		{`""`, ""},
	}

	for _, tt := range tests {
		l := New(tt.input)
		tok := l.NextToken()

		if tok.Type != token.STRING {
			t.Fatalf("token type wrong. expected=%q, got=%q", token.STRING, tok.Type)
		}

		if tok.Literal != tt.expectedLiteral {
			t.Fatalf("literal wrong. expected=%q, got=%q", tt.expectedLiteral, tok.Literal)
		}
	}
}

func TestNumberTokens(t *testing.T) {
	tests := []struct {
		input        string
		expectedType token.TokenType
		expectedLiteral string
	}{
		{"42", token.INT, "42"},
		{"0", token.INT, "0"},
		{"123456789", token.INT, "123456789"},
		{"3.14", token.FLOAT, "3.14"},
		{"0.5", token.FLOAT, "0.5"},
		{"123.456", token.FLOAT, "123.456"},
	}

	for _, tt := range tests {
		l := New(tt.input)
		tok := l.NextToken()

		if tok.Type != tt.expectedType {
			t.Fatalf("token type wrong. expected=%q, got=%q", tt.expectedType, tok.Type)
		}

		if tok.Literal != tt.expectedLiteral {
			t.Fatalf("literal wrong. expected=%q, got=%q", tt.expectedLiteral, tok.Literal)
		}
	}
}

func TestLineNumbers(t *testing.T) {
	input := `let x = 5;
let y = 10;
let z = x + y;`

	l := New(input)

	expectedLines := []int{1, 1, 1, 1, 1, 2, 2, 2, 2, 2, 3, 3, 3, 3, 3, 3, 3}
	
	for i, expectedLine := range expectedLines {
		tok := l.NextToken()
		if tok.Line != expectedLine {
			t.Fatalf("tests[%d] - line number wrong. expected=%d, got=%d",
				i, expectedLine, tok.Line)
		}
	}
}

func TestComments(t *testing.T) {
	input := `// This is a comment
let x = 5; // Another comment
/* Multi-line
   comment */
let y = 10;`

	l := New(input)

	expectedTokens := []token.TokenType{
		token.LET, token.IDENT, token.ASSIGN, token.INT, token.SEMICOLON,
		token.LET, token.IDENT, token.ASSIGN, token.INT, token.SEMICOLON,
		token.EOF,
	}

	for i, expectedType := range expectedTokens {
		tok := l.NextToken()
		if tok.Type != expectedType {
			t.Fatalf("tests[%d] - tokentype wrong. expected=%q, got=%q",
				i, expectedType, tok.Type)
		}
	}
}