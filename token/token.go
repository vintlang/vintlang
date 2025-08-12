package token

type TokenType string

type Token struct {
	Type    TokenType
	Literal string
	Line    int
}

const (
	ILLEGAL = "ILLEGAL"
	EOF     = "EOF"

	// Identifiers + literals
	IDENT  = "IDENT"
	INT    = "INT"
	STRING = "STRING"
	FLOAT  = "FLOAT"

	// Operators
	ASSIGN          = "="
	PLUS            = "+"
	MINUS           = "-"
	BANG            = "!"
	ASTERISK        = "*"
	AMPERSAND       = "&"
	POW             = "**"
	SLASH           = "/"
	MODULUS         = "%"
	LT              = "<"
	LTE             = "<="
	GT              = ">"
	GTE             = ">="
	EQ              = "=="
	NOT_EQ          = "!="
	AND             = "&&"
	OR              = "||"
	PLUS_ASSIGN     = "+="
	PLUS_PLUS       = "++"
	MINUS_ASSIGN    = "-="
	MINUS_MINUS     = "--"
	ASTERISK_ASSIGN = "*="
	SLASH_ASSIGN    = "/="
	MODULUS_ASSIGN  = "%="
	SHEBANG         = "#!"

	// Delimiters
	COMMA     = ","
	SEMICOLON = ";"
	LPAREN    = "("
	RPAREN    = ")"
	LBRACE    = "{"
	RBRACE    = "}"
	LBRACKET  = "["
	RBRACKET  = "]"
	COLON     = ":"
	DOT       = "."
	RANGE     = ".."
	AT        = "@"
	ARROW     = "=>"

	// Keywords
	FUNCTION = "FUNCTION"
	LET      = "LET"
	CONST    = "CONST"
	TRUE     = "TRUE"
	FALSE    = "FALSE"
	IF       = "IF"
	ELSE     = "ELSE"
	RETURN   = "RETURN"
	WHILE    = "WHILE"
	NULL     = "NULL"
	BREAK    = "BREAK"
	CONTINUE = "CONTINUE"
	IN       = "IN"
	FOR      = "FOR"
	SWITCH   = "SWITCH"
	CASE     = "CASE"
	DEFAULT  = "DEFAULT"
	MATCH    = "MATCH"
	IMPORT   = "IMPORT"
	PACKAGE  = "PACKAGE"
	INCLUDE  = "INCLUDE"
	TODO     = "TODO"
	WARN     = "WARN"
	ERROR    = "ERROR"
	DEFER    = "DEFER"
	REPEAT   = "REPEAT"
	INFO     = "INFO"
	DEBUG    = "DEBUG"
	NOTE     = "NOTE"
	SUCCESS  = "SUCCESS"
	
	// Async/Concurrency Keywords
	ASYNC    = "ASYNC"
	AWAIT    = "AWAIT"
	GO       = "GO"
	CHAN     = "CHAN"
	
	// Error Handling Keywords
	THROW    = "THROW"
)

var keywords = map[string]TokenType{
	"func":     FUNCTION,
	"let":      LET,
	"const":    CONST,
	"true":     TRUE,
	"false":    FALSE,
	"if":       IF,
	"else":     ELSE,
	"while":    WHILE,
	"return":   RETURN,
	"break":    BREAK,
	"continue": CONTINUE,
	"null":     NULL,
	"in":       IN,
	"for":      FOR,
	"switch":   SWITCH,
	"case":     CASE,
	"default":  DEFAULT,
	"match":    MATCH,
	"import":   IMPORT,
	"package":  PACKAGE,
	"include":  INCLUDE,
	"todo":     TODO,
	"warn":     WARN,
	"error":    ERROR,
	"@":        AT,
	"defer":    DEFER,
	"repeat":   REPEAT,
	"info":     INFO,
	"debug":    DEBUG,
	"note":     NOTE,
	"success":  SUCCESS,
	"async":    ASYNC,
	"await":    AWAIT,
	"go":       GO,
	"chan":     CHAN,
	"throw":    THROW,
}

func LookupIdent(ident string) TokenType {
	if tok, ok := keywords[ident]; ok {
		return tok
	}
	return IDENT
}
