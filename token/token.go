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
	NULL_COALESCE   = "??"
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
	TRACE    = "TRACE"
	FATAL    = "FATAL"
	CRITICAL = "CRITICAL"
	LOG      = "LOG"
	
	// Async/Concurrency Keywords
	ASYNC    = "ASYNC"
	AWAIT    = "AWAIT"
	GO       = "GO"
	CHAN     = "CHAN"
	
	// Error Handling Keywords
	THROW    = "THROW"
	
	// Capitalized Declaratives
	INFO_CAP     = "INFO_CAP"
	DEBUG_CAP    = "DEBUG_CAP"
	NOTE_CAP     = "NOTE_CAP"
	TODO_CAP     = "TODO_CAP"
	WARN_CAP     = "WARN_CAP"
	SUCCESS_CAP  = "SUCCESS_CAP"
	ERROR_CAP    = "ERROR_CAP"
	TRACE_CAP    = "TRACE_CAP"
	FATAL_CAP    = "FATAL_CAP"
	CRITICAL_CAP = "CRITICAL_CAP"
	LOG_CAP      = "LOG_CAP"
	
	// Type system tokens
	AS    = "AS"    // type casting: x as int
	IS    = "IS"    // type checking: x is int
	PIPE  = "|"     // union types: int | string
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
	"trace":    TRACE,
	"fatal":    FATAL,
	"critical": CRITICAL,
	"log":      LOG,
	"async":    ASYNC,
	"await":    AWAIT,
	"go":       GO,
	"chan":     CHAN,
	"throw":    THROW,
	"Info":     INFO_CAP,
	"Debug":    DEBUG_CAP,
	"Note":     NOTE_CAP,
	"Todo":     TODO_CAP,
	"Warn":     WARN_CAP,
	"Success":  SUCCESS_CAP,
	"Error":    ERROR_CAP,
	"Trace":    TRACE_CAP,
	"Fatal":    FATAL_CAP,
	"Critical": CRITICAL_CAP,
	"Log":      LOG_CAP,
	"as":       AS,
	"is":       IS,
}

func LookupIdent(ident string) TokenType {
	if tok, ok := keywords[ident]; ok {
		return tok
	}
	return IDENT
}
