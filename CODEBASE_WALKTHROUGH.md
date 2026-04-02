# VintLang Codebase Walkthrough

A comprehensive guide to the VintLang interpreter internals — how source code becomes running programs.

---

## Table of Contents

1. [Repository Structure](#1-repository-structure)
2. [Execution Pipeline Overview](#2-execution-pipeline-overview)
3. [Entry Point — `main.go`](#3-entry-point--maingo)
4. [Token System — `token/`](#4-token-system--token)
5. [Lexer — `lexer/`](#5-lexer--lexer)
6. [AST (Abstract Syntax Tree) — `ast/`](#6-ast-abstract-syntax-tree--ast)
7. [Parser — `parser/`](#7-parser--parser)
8. [Object System — `object/`](#8-object-system--object)
9. [Environment and Scoping](#9-environment-and-scoping)
10. [Evaluator — `evaluator/`](#10-evaluator--evaluator)
11. [Built-in Functions — `evaluator/builtins/`](#11-built-in-functions--evaluatorbuiltins)
12. [Module System — `module/`](#12-module-system--module)
13. [REPL — `repl/`](#13-repl--repl)
14. [Error Handling — `vintErrors/`](#14-error-handling--vinterrors)
15. [End-to-End Execution Example](#15-end-to-end-execution-example)
16. [Additional Components](#16-additional-components)

---

## 1. Repository Structure

```
vintlang/
├── main.go                 # CLI entry point and command dispatcher
├── config/                 # Version and configuration constants
│
│  ── Core Interpreter Pipeline ──
├── token/                  # Token type definitions and keyword lookup
├── lexer/                  # Tokenizer — converts source text to tokens
├── ast/                    # Abstract Syntax Tree node definitions
├── parser/                 # Parser — converts tokens to AST (Pratt parsing)
├── object/                 # Runtime value types and environment/scope
├── evaluator/              # Tree-walking interpreter
│   └── builtins/           # Built-in function implementations
│
│  ── Standard Library ──
├── module/                 # Built-in module registry (os, json, http, etc.)
│
│  ── Interactive Tools ──
├── repl/                   # Read-Eval-Print Loop (interactive shell)
├── toolkit/                # CLI utilities (init, get, test, docs)
│
│  ── Error System ──
├── vintErrors/             # Structured error codes and formatting
├── errors/                 # Additional error utilities
│
│  ── Packaging & Tooling ──
├── bundler/                # Binary packaging for Vint programs
├── compiler/               # Compilation support (experimental)
├── vm/                     # Virtual machine (experimental)
│
│  ── Resources ──
├── examples/               # Example Vint programs
├── docs/                   # Module documentation (markdown)
├── styles/                 # Terminal UI styling (lipgloss)
├── extensions/             # Editor extensions
├── website/                # Project website
│
├── go.mod / go.sum         # Go module dependencies
├── Makefile                # Build targets
├── Dockerfile              # Container build
└── .goreleaser.yml         # Release automation
```

The core interpreter follows a classic four-stage pipeline: **Lexer → Parser → AST → Evaluator**. Each stage has a dedicated package with clean boundaries between them.

---

## 2. Execution Pipeline Overview

When VintLang runs a program, the source code flows through four sequential stages:

```
┌─────────────────┐
│   Source Code    │   "let x = 5 + 3"
│   (text file)   │
└────────┬────────┘
         │
         ▼
┌─────────────────┐
│     LEXER       │   Scans characters, produces tokens
│  (lexer.go)     │   → [LET, IDENT("x"), ASSIGN, INT(5), PLUS, INT(3)]
└────────┬────────┘
         │
         ▼
┌─────────────────┐
│     PARSER      │   Reads tokens, builds tree using Pratt parsing
│  (parser.go)    │   → LetStatement{ Name: "x", Value: InfixExpr{5, "+", 3} }
└────────┬────────┘
         │
         ▼
┌─────────────────┐
│   EVALUATOR     │   Walks the AST, executes each node
│ (evaluator.go)  │   → Environment: { "x": Integer(8) }
└────────┬────────┘
         │
         ▼
┌─────────────────┐
│     OUTPUT      │   Runtime objects, side effects, printed output
│   (objects)     │
└─────────────────┘
```

Each stage communicates through well-defined interfaces:
- The **Lexer** produces `token.Token` values.
- The **Parser** consumes tokens and produces `ast.Node` trees.
- The **Evaluator** walks `ast.Node` trees and produces `object.VintObject` values.

---

## 3. Entry Point — `main.go`

The `main.go` file (342 lines) is the CLI dispatcher. It determines what to do based on command-line arguments.

### Command Routing

```go
func main() {
    args := os.Args[1:]
    if len(args) == 0 {
        repl.Start()    // No arguments → interactive REPL
        return
    }
    switch args[0] {
    case "help":      // Show help text
    case "play":      repl.Playground()       // TUI playground
    case "docs":      toolkit.RunDocs()        // Documentation browser
    case "bundler":   bundler.Bundle(args[1])  // Package into binary
    case "get":       toolkit.GetPackage(...)  // Install packages
    case "init":      toolkit.InitProject()    // Create new project
    case "test":      toolkit.Test()           // Run test files
    case "fmt":       formatFile(args[1])      // Format source code
    case "version":   // Print version
    case "--trace":   runWithTrace(args[1], ...)  // Debug pipeline
    default:          run(args[0])             // Execute .vint file
    }
}
```

### File Execution — `run()`

The `run()` function is the primary entry point for running `.vint` files:

```go
func run(file string) {
    contents, err := os.ReadFile(file)
    // ... error handling ...

    // Add the script's directory to the import search path
    dir := filepath.Dir(absPath)
    evaluator.AddSearchPath(dir)

    // Execute via the REPL reader (same pipeline for files and REPL)
    repl.ReadWithFilename(string(contents), file)
}
```

### Trace Mode — `runWithTrace()`

The `--trace` flag reveals all four pipeline stages for debugging:

1. **Stage 1 — Source Code**: Displays the raw source.
2. **Stage 2 — Lexer Output**: Shows every token with its type, literal, line, and column.
3. **Stage 3 — Parser Output**: Displays the AST as a formatted string.
4. **Stage 4 — Evaluator Output**: Shows the final result of evaluation.

This is an excellent tool for understanding how the interpreter processes code.

---

## 4. Token System — `token/`

**File: `token/token.go`**

Tokens are the atomic units produced by the lexer. Each token represents a meaningful piece of syntax.

### Token Structure

```go
type TokenType string

type Token struct {
    Type    TokenType  // What kind of token (e.g., INT, PLUS, LET)
    Literal string     // The actual source text (e.g., "42", "+", "let")
    Line    int        // Source line number (1-based)
    Column  int        // Source column number (1-based)
}
```

### Token Categories

**Literals and Identifiers:**

| Token | Example | Description |
|-------|---------|-------------|
| `IDENT` | `myVar` | Variable/function name |
| `INT` | `42` | Integer literal |
| `FLOAT` | `3.14` | Floating-point literal |
| `STRING` | `"hello"` | String literal |

**Operators:**

| Token | Symbol | Token | Symbol |
|-------|--------|-------|--------|
| `PLUS` | `+` | `MINUS` | `-` |
| `ASTERISK` | `*` | `SLASH` | `/` |
| `MODULUS` | `%` | `POW` | `**` |
| `EQ` | `==` | `NOT_EQ` | `!=` |
| `LT` | `<` | `GT` | `>` |
| `LTE` | `<=` | `GTE` | `>=` |
| `AND` | `&&` | `OR` | `\|\|` |
| `ASSIGN` | `=` | `BANG` | `!` |
| `PLUS_ASSIGN` | `+=` | `MINUS_ASSIGN` | `-=` |
| `PLUS_PLUS` | `++` | `MINUS_MINUS` | `--` |
| `NULL_COALESCE` | `??` | `RANGE` | `..` |

**Keywords (mapped via `LookupIdent()`):**

| Keyword | Token | Keyword | Token |
|---------|-------|---------|-------|
| `let` | `LET` | `const` | `CONST` |
| `func` | `FUNCTION` | `return` | `RETURN` |
| `if` | `IF` | `else` | `ELSE` |
| `for` | `FOR` | `while` | `WHILE` |
| `switch` | `SWITCH` | `case` | `CASE` |
| `match` | `MATCH` | `break` | `BREAK` |
| `continue` | `CONTINUE` | `import` | `IMPORT` |
| `true` | `TRUE` | `false` | `FALSE` |
| `null` | `NULL` | `in` | `IN` |
| `struct` | `STRUCT` | `enum` | `ENUM` |
| `defer` | `DEFER` | `go` | `GO` |
| `async` | `ASYNC` | `await` | `AWAIT` |

### Keyword Lookup

The `LookupIdent()` function maps identifier strings to their token types:

```go
func LookupIdent(ident string) TokenType {
    if tok, ok := keywords[ident]; ok {
        return tok    // It's a keyword (e.g., "let" → LET)
    }
    return IDENT      // It's a user identifier
}
```

This is how the lexer distinguishes keywords from variable names — every identifier is looked up in the keyword table.

---

## 5. Lexer — `lexer/`

**File: `lexer/lexer.go`** (637 lines)

The lexer (also called tokenizer or scanner) reads raw source text character by character and produces a stream of tokens.

### Lexer Structure

```go
type Lexer struct {
    input        []rune   // Source code as Unicode runes
    position     int      // Index of current character
    readPosition int      // Index of next character (look-ahead)
    ch           rune     // Current character being examined
    line         int      // Current line number
    column       int      // Current column number
    filename     string   // Source file name (for error messages)
    errors       []string // Accumulated lexer errors
}
```

The input is stored as `[]rune` (not `[]byte`) for full Unicode support — VintLang can handle identifiers and strings with any Unicode characters.

### Initialization

```go
func New(input string) *Lexer              // Default filename "main.vint"
func NewWithFilename(input, filename string) *Lexer  // Custom filename
```

Both constructors convert the input string to runes and call `readChar()` to load the first character.

### Character Reading

```go
func (l *Lexer) readChar() {
    if l.readPosition >= len(l.input) {
        l.ch = 0    // EOF signaled as null character
    } else {
        l.ch = l.input[l.readPosition]
    }
    l.position = l.readPosition
    l.readPosition++
    l.column++
    // Track newlines for line counting
}
```

The `peekChar()` method looks ahead one character without advancing the position — essential for recognizing multi-character tokens like `==`, `!=`, `&&`, `++`.

### The Core — `NextToken()`

`NextToken()` is the heart of the lexer. It's called repeatedly by the parser to get the next token. The method:

1. **Skips whitespace** (spaces, tabs, carriage returns, newlines — tracking line numbers on newlines).

2. **Skips comments:**
   - Single-line: `//` until end of line
   - Multi-line: `/* ... */`
   - Shebang: `#!` on the first line

3. **Matches the current character** against token patterns:

```go
func (l *Lexer) NextToken() token.Token {
    l.skipWhitespace()
    // Skip comments...

    switch l.ch {
    case '=':
        if l.peekChar() == '=' {
            // Two-character token: ==
            tok = token.Token{Type: token.EQ, Literal: "=="}
            l.readChar()  // consume second '='
        } else if l.peekChar() == '>' {
            // Arrow: =>
            tok = token.Token{Type: token.ARROW, Literal: "=>"}
            l.readChar()
        } else {
            tok = token.Token{Type: token.ASSIGN, Literal: "="}
        }
    case '+':
        if l.peekChar() == '+' {
            tok = token.Token{Type: token.PLUS_PLUS, Literal: "++"}
            l.readChar()
        } else if l.peekChar() == '=' {
            tok = token.Token{Type: token.PLUS_ASSIGN, Literal: "+="}
            l.readChar()
        } else {
            tok = token.Token{Type: token.PLUS, Literal: "+"}
        }
    // ... similar patterns for -, *, /, <, >, !, &, |, etc.

    case '"':
        tok.Type = token.STRING
        tok.Literal = l.readString()    // Read until closing "

    case '\'':
        tok.Type = token.STRING
        tok.Literal = l.readSingleQuoteString()

    case 0:
        tok.Type = token.EOF

    default:
        if isLetter(l.ch) {
            literal := l.readIdentifier()
            tok.Type = token.LookupIdent(literal)  // keyword or IDENT
            tok.Literal = literal
            return tok
        } else if isDigit(l.ch) {
            return l.readDecimal()  // INT or FLOAT token
        } else {
            tok = l.createIllegalToken()  // Unrecognized character
        }
    }
    l.readChar()
    return tok
}
```

### String Parsing

Double-quoted strings support escape sequences:

```go
func (l *Lexer) readString() string {
    // Read characters until closing "
    // Handle escape sequences:
    //   \n → newline    \t → tab        \r → carriage return
    //   \\ → backslash  \" → quote      \0 → null
    //   \xHH → hex byte  \uHHHH → Unicode codepoint
}
```

Single-quoted strings (`'...'`) have simpler escape handling.

### Number Parsing

```go
func (l *Lexer) readDecimal() token.Token {
    integer := l.readNumber()  // Read digit sequence
    if l.ch == '.' && isDigit(l.peekChar()) {
        l.readChar()  // consume '.'
        fraction := l.readNumber()
        return token.Token{Type: token.FLOAT, Literal: integer + "." + fraction}
    }
    return token.Token{Type: token.INT, Literal: integer}
}
```

### Error Handling

When the lexer encounters an unrecognized character, it creates an `ILLEGAL` token and adds a detailed error message with the source line and a caret pointing to the problem:

```
Error at line 5, column 12 in main.vint:
    let x = @#$
               ^
    Illegal character: '#'
```

---

## 6. AST (Abstract Syntax Tree) — `ast/`

**Files: `ast/ast.go`, `ast/expressions.go`, `ast/statements.go`** (7 files total)

The AST is the structured representation of the source code that the evaluator walks to execute the program.

### Core Interfaces

```go
type Node interface {
    TokenLiteral() string  // Returns the token's literal text
    String() string        // Returns a printable representation of the node
}

type Statement interface {
    Node
    statementNode()  // Marker method — distinguishes from Expression
}

type Expression interface {
    Node
    expressionNode()  // Marker method — distinguishes from Statement
}
```

Everything in the AST is either a `Statement` (which does something) or an `Expression` (which produces a value). The distinction matters: expressions can appear inside other expressions, while statements are top-level constructs.

### Program Root

```go
type Program struct {
    Statements []Statement  // A program is a list of statements
}
```

### Key Statement Nodes

```go
// let x = 5
type LetStatement struct {
    Token token.Token     // The LET token
    Name  *Identifier     // Variable name
    Value Expression      // Right-hand side expression
}

// const PI = 3.14
type ConstStatement struct {
    Token token.Token
    Name  *Identifier
    Value Expression
}

// return x + y
type ReturnStatement struct {
    Token       token.Token
    ReturnValue Expression
}

// { stmt1; stmt2; stmt3 }
type BlockStatement struct {
    Token      token.Token
    Statements []Statement
}

// import os
type ImportStatement struct {
    Token token.Token
    Value Expression
}

// defer cleanup()
type DeferStatement struct {
    Token token.Token
    Call  Expression
}

// struct Point { x; y }
type StructStatement struct {
    Token   token.Token
    Name    string
    Fields  []*StructField
    Methods map[string]*FunctionLiteral
}

// enum Color { RED; GREEN; BLUE }
type EnumStatement struct {
    Token   token.Token
    Name    string
    Members map[string]Expression
}
```

### Key Expression Nodes

```go
// x + y, a == b, c && d
type InfixExpression struct {
    Token    token.Token
    Left     Expression   // Left operand
    Operator string       // "+", "-", "==", etc.
    Right    Expression   // Right operand
}

// -x, !x, ++x
type PrefixExpression struct {
    Token    token.Token
    Operator string
    Right    Expression
}

// x++, y--
type PostfixExpression struct {
    Token    token.Token
    Operator string
    Left     Expression
}

// if (condition) { consequence } else { alternative }
type IfExpression struct {
    Token       token.Token
    Condition   Expression
    Consequence *BlockStatement
    Alternative *BlockStatement
}

// func(x, y) { return x + y }
type FunctionLiteral struct {
    Token      token.Token
    Name       string
    Parameters []*Identifier
    Defaults   map[string]Expression
    Body       *BlockStatement
}

// add(3, 4)
type CallExpression struct {
    Token     token.Token
    Function  Expression     // The function being called
    Arguments []Expression   // Arguments passed
}

// obj.method(args)
type MethodExpression struct {
    Token     token.Token
    Object    Expression
    Method    Expression
    Arguments []Expression
    Defaults  map[string]Expression
}

// arr[0], dict["key"]
type IndexExpression struct {
    Token token.Token
    Left  Expression   // The array/dict
    Index Expression   // The index/key
}

// obj.property
type DotExpression struct {
    Token token.Token
    Left  Expression
    Right Expression
}

// for item in items { ... }
type ForIn struct {
    Token    token.Token
    Key      string
    Value    string
    Iterable Expression
    Body     *BlockStatement
}

// while (condition) { ... }
type WhileExpression struct {
    Token     token.Token
    Condition Expression
    Body      *BlockStatement
}

// match x { "a" => 1, "b" => 2, _ => 0 }
type MatchExpression struct {
    Token       token.Token
    Subject     Expression
    Cases       []*MatchCase
    Default     *BlockStatement
}

// switch (x) { case 1 { ... } default { ... } }
type SwitchExpression struct {
    Token   token.Token
    Subject Expression
    Cases   []*SwitchCase
    Default *BlockStatement
}

// [1, 2, 3]
type ArrayLiteral struct {
    Token    token.Token
    Elements []Expression
}

// {"key": value}
type DictLiteral struct {
    Token token.Token
    Pairs map[Expression]Expression
}
```

### The `String()` Method

Every AST node implements `String()` which reconstructs the source code from the tree. This powers the `vint fmt` code formatter — the source is parsed into an AST and then printed back via `String()`.

---

## 7. Parser — `parser/`

**Files: `parser/parser.go` + 27 additional files** (~3000+ lines)

The parser reads the token stream and builds an AST. VintLang uses a **Pratt parser** (top-down operator precedence parser), which is elegant for handling expressions with different precedence levels.

### Parser Structure

```go
type Parser struct {
    l         *lexer.Lexer
    curToken  token.Token     // Current token being examined
    peekToken token.Token     // Next token (one-token lookahead)
    prevToken token.Token     // Previous token (for postfix operators)
    errors    []string

    // Function registries for Pratt parsing:
    prefixParseFns  map[token.TokenType]prefixParseFn
    infixParseFns   map[token.TokenType]infixParseFn
    postfixParseFns map[token.TokenType]postfixParseFn
}

type prefixParseFn  func() ast.Expression
type infixParseFn   func(ast.Expression) ast.Expression
type postfixParseFn func(ast.Expression) ast.Expression
```

### Operator Precedence

The parser defines 14 precedence levels (lowest to highest):

```go
const (
    _ int = iota
    LOWEST        // Starting precedence
    ASSIGN        // =
    COND          // || && ??
    EQUALS        // == !=
    LESSGREATER   // < > <= >=
    RANGE         // ..
    SUM           // + -
    PRODUCT       // * /
    POWER         // **
    MODULUS       // %
    PREFIX        // -x !x
    CALL          // function()
    INDEX         // array[i]
    DOT           // obj.prop  (highest)
)
```

Higher precedence means tighter binding. For example, `2 + 3 * 4` is parsed as `2 + (3 * 4)` because `PRODUCT > SUM`.

### How Pratt Parsing Works

The key insight of Pratt parsing is that each token type can have associated parsing functions:

- **Prefix parse functions**: Handle tokens that appear at the start of an expression (identifiers, literals, unary operators, `if`, `func`, etc.)
- **Infix parse functions**: Handle tokens that appear between two expressions (binary operators, `(` for calls, `[` for indexing, `.` for property access)

The algorithm works by recursive descent with precedence:

```go
func (p *Parser) parseExpression(precedence int) ast.Expression {
    // 1. Get the prefix parser for the current token
    prefix := p.prefixParseFns[p.curToken.Type]
    if prefix == nil {
        p.noPrefixParseFnError(p.curToken.Type)
        return nil
    }
    leftExp := prefix()  // Parse the left side

    // 2. While the next token has higher precedence, parse as infix
    for !p.peekTokenIs(token.SEMICOLON) && precedence < p.peekPrecedence() {
        infix := p.infixParseFns[p.peekToken.Type]
        if infix == nil {
            return leftExp
        }
        p.nextToken()
        leftExp = infix(leftExp)  // The left side becomes part of a larger expression
    }

    return leftExp
}
```

**Example: Parsing `2 + 3 * 4`**

1. `parseExpression(LOWEST)` starts
2. Prefix: `2` → `IntegerLiteral(2)`, leftExp = `2`
3. Peek is `+` (precedence SUM). LOWEST < SUM, so enter loop:
   - Infix `+`: calls `parseInfixExpression(IntegerLiteral(2))`
   - Inside, calls `parseExpression(SUM)` for the right side
   - Prefix: `3` → `IntegerLiteral(3)`, leftExp = `3`
   - Peek is `*` (precedence PRODUCT). SUM < PRODUCT, so enter loop:
     - Infix `*`: calls `parseInfixExpression(IntegerLiteral(3))`
     - Inside, calls `parseExpression(PRODUCT)` for right side
     - Prefix: `4` → `IntegerLiteral(4)`
     - Peek has lower precedence, return `4`
     - Result: `InfixExpression(3, *, 4)`
   - Back in `+` context: right side is `InfixExpression(3, *, 4)`
   - Result: `InfixExpression(2, +, InfixExpression(3, *, 4))`

This naturally produces the correct tree: `2 + (3 * 4)`.

### Parse Function Registration

The parser constructor registers 50+ parse functions:

```go
func New(l *lexer.Lexer) *Parser {
    p := &Parser{l: l}

    // Prefix parsers (handle start of expression)
    p.registerPrefix(token.IDENT, p.parseIdentifier)
    p.registerPrefix(token.INT, p.parseIntegerLiteral)
    p.registerPrefix(token.FLOAT, p.parseFloatLiteral)
    p.registerPrefix(token.STRING, p.parseStringLiteral)
    p.registerPrefix(token.TRUE, p.parseBoolean)
    p.registerPrefix(token.FALSE, p.parseBoolean)
    p.registerPrefix(token.NULL, p.parseNull)
    p.registerPrefix(token.BANG, p.parsePrefixExpression)
    p.registerPrefix(token.MINUS, p.parsePrefixExpression)
    p.registerPrefix(token.IF, p.parseIfExpression)
    p.registerPrefix(token.FUNCTION, p.parseFunctionLiteral)
    p.registerPrefix(token.LBRACKET, p.parseArrayLiteral)
    p.registerPrefix(token.LBRACE, p.parseDictLiteral)
    p.registerPrefix(token.SWITCH, p.parseSwitchStatement)
    p.registerPrefix(token.MATCH, p.parseMatchExpression)
    p.registerPrefix(token.IMPORT, p.parseImportExpression)
    // ... more

    // Infix parsers (handle middle of expression)
    p.registerInfix(token.PLUS, p.parseInfixExpression)
    p.registerInfix(token.MINUS, p.parseInfixExpression)
    p.registerInfix(token.ASTERISK, p.parseInfixExpression)
    p.registerInfix(token.SLASH, p.parseInfixExpression)
    p.registerInfix(token.EQ, p.parseInfixExpression)
    p.registerInfix(token.NOT_EQ, p.parseInfixExpression)
    p.registerInfix(token.LPAREN, p.parseCallExpression)    // f(x)
    p.registerInfix(token.LBRACKET, p.parseIndexExpression)  // a[i]
    p.registerInfix(token.DOT, p.parseDotExpression)         // o.p
    p.registerInfix(token.ASSIGN, p.parseAssignExpression)   // x = v
    // ... more

    // Postfix parsers
    p.registerPostfix(token.PLUS_PLUS, p.parsePostfixExpression)   // x++
    p.registerPostfix(token.MINUS_MINUS, p.parsePostfixExpression) // x--

    return p
}
```

### Statement Parsing

The `ParseProgram()` method processes statements in sequence:

```go
func (p *Parser) ParseProgram() *ast.Program {
    program := &ast.Program{}
    for !p.curTokenIs(token.EOF) {
        stmt := p.parseStatement()
        if stmt != nil {
            program.Statements = append(program.Statements, stmt)
        }
        p.nextToken()
    }
    return program
}
```

`parseStatement()` dispatches to the appropriate parser based on the current token:

```go
func (p *Parser) parseStatement() ast.Statement {
    switch p.curToken.Type {
    case token.LET:       return p.parseLetStatement()
    case token.CONST:     return p.parseConstStatement()
    case token.RETURN:    return p.parseReturnStatement()
    case token.BREAK:     return p.parseBreakStatement()
    case token.CONTINUE:  return p.parseContinueStatement()
    case token.IMPORT:    return p.parseImportStatement()
    case token.PACKAGE:   return p.parsePackageStatement()
    case token.DEFER:     return p.parseDeferStatement()
    case token.ENUM:      return p.parseEnumStatement()
    case token.STRUCT:    return p.parseStructStatement()
    case token.GO:        return p.parseGoStatement()
    // ... logging keywords (todo, warn, error, etc.)
    default:              return p.parseExpressionStatement()
    }
}
```

### Parser File Organization

Each major language feature has its own parser file:

| File | Parses |
|------|--------|
| `parser.go` | Core structure, precedences, Pratt algorithm |
| `statements.go` | `let`, `const`, `return`, `break`, `continue` |
| `expressions.go` | Prefix, infix, postfix expressions |
| `function.go` | `func` literals and parameters |
| `if.go` | `if`/`else` expressions |
| `while.go` | `while` loops |
| `for.go` | `for...in` loops |
| `match.go` | `match` pattern matching |
| `switch.go` | `switch`/`case` statements |
| `arrays.go` | Array literals `[...]` |
| `dict.go` | Dictionary literals `{...}` |
| `index.go` | Index expressions `a[i]` and slicing `a[i:j]` |
| `dot.go` | Dot expressions `obj.prop` and method calls |
| `import.go` | Import expressions |
| `package.go` | Package declarations |
| `string.go`, `integer.go`, `float.go`, `boolean.go` | Literal parsing |
| `enum.go` | Enum definitions |
| `struct.go` | Struct definitions |
| `async.go` | `async`/`await` expressions |

### Error Recovery

When the parser encounters a syntax error, it tries to recover by skipping to the next statement boundary:

```go
func (p *Parser) Synchronize() {
    for !p.curTokenIs(token.EOF) {
        if p.curTokenIs(token.SEMICOLON) { return }
        switch p.peekToken.Type {
        case token.LET, token.RETURN, token.FUNCTION, token.IF, token.FOR:
            return  // Found a statement boundary
        }
        p.nextToken()
    }
}
```

---

## 8. Object System — `object/`

**Files: `object/object.go` + 20 additional files** (21 files total)

Every runtime value in VintLang is represented as a `VintObject`. This is the type system that the evaluator produces and manipulates.

### Core Interface

```go
type VintObject interface {
    Type() VintObjectType   // Returns the type tag (e.g., "INTEGER")
    Inspect() string        // Returns a human-readable string representation
}
```

### Primitive Types

```go
// object/integer.go
type Integer struct {
    Value int64
}
// Supports methods: abs(), pow(), sqrt(), etc.

// object/float.go
type Float struct {
    Value float64
}

// object/strings.go
type String struct {
    Value string
}
// Supports methods: len(), upper(), lower(), contains(), split(),
// replace(), trim(), startsWith(), endsWith(), etc.

// object/boolean.go
type Boolean struct {
    Value bool
}

// object/object.go
type Null struct{}   // The null value (singleton)
```

### Collection Types

```go
// object/array.go
type Array struct {
    Elements []VintObject
}
// Supports methods: push(), pop(), shift(), length(), join(),
// contains(), indexOf(), reverse(), sort(), unique(), first(), last(), etc.
// Implements Iterable for for-in loops.

// object/dict.go
type Dict struct {
    Pairs  map[HashKey]DictPair
    Keys   []HashKey            // Maintains insertion order
}
type DictPair struct {
    Key   VintObject
    Value VintObject
}
// Supports methods: keys(), values(), has_key(), merge(), toJSON(), etc.
// Implements Iterable for for-in loops.
```

### Function Types

```go
// object/function.go
type Function struct {
    Name       string
    Parameters []*ast.Identifier
    Defaults   map[string]ast.Expression  // Default parameter values
    Body       *ast.BlockStatement
    Env        *Environment               // Captured closure environment
    IsAsync    bool
    IsStreaming bool
}

// Built-in functions (implemented in Go)
type Builtin struct {
    Fn func(args ...VintObject) VintObject
}
```

### The Hashable Interface

Objects that can be used as dictionary keys implement `Hashable`:

```go
type Hashable interface {
    HashKey() HashKey
}

type HashKey struct {
    Type  VintObjectType
    Value uint64
}
```

Integers, strings, and booleans are hashable. Arrays and dicts are not (like most languages).

### The Iterable Interface

Objects that support `for...in` loops implement `Iterable`:

```go
type Iterable interface {
    Next() (VintObject, VintObject)  // Returns (key, value) pair
    Reset()                           // Resets iteration to the beginning
}
```

Arrays iterate as `(index, element)`, dicts as `(key, value)`, strings as `(index, character)`.

### Struct and Enum Types

```go
// object/struct.go
type Struct struct {
    Name    string
    Fields  []StructField           // Field definitions with optional defaults
    Methods map[string]*Function    // Associated methods
    Env     *Environment            // Definition scope
}

type StructInstance struct {
    Struct *Struct
    Fields *Environment             // Instance field values
}

// object/object.go
type Enum struct {
    Name    string
    Members map[string]VintObject   // Named members with values
}
```

### Special Types

```go
// Control flow signals (not errors — used internally)
type ReturnValue struct { Value VintObject }
type Break struct{}
type Continue struct{}

// Error type
type Error struct { Message string }

// Module type (for import)
type Module struct {
    Name      string
    Functions map[string]*Builtin
}

// Deferred function call
type DeferredCall struct {
    Function VintObject
    Args     []VintObject
}

// Package type (user-defined module)
type Package struct {
    Name string
    Env  *Environment
}
```

---

## 9. Environment and Scoping

**File: `object/environment.go`**

The `Environment` is VintLang's scope system. It maps variable names to their values and implements lexical scoping through a chain of nested environments.

### Environment Structure

```go
type Environment struct {
    store       map[string]VintObject   // Regular variables
    funcs       map[string][]*Function  // Function overloads (multiple functions, same name)
    constants   map[string]bool         // Tracks which names are constants
    outer       *Environment            // Parent scope (lexical scoping)
    isFuncScope bool                    // Marks function boundaries
    deferredCalls []*DeferredCall       // Deferred calls in this scope
    deferMu     sync.Mutex              // Thread safety for defers
}
```

### Scope Chain

When looking up a variable, the environment walks up the chain:

```go
func (e *Environment) Get(name string) (VintObject, bool) {
    obj, ok := e.store[name]
    if !ok && e.outer != nil {
        return e.outer.Get(name)  // Walk up to parent scope
    }
    return obj, ok
}
```

This implements **lexical scoping**: inner scopes can access variables from outer scopes.

### Creating Scopes

```go
// New global scope (no parent)
func NewEnvironment() *Environment

// New child scope (has parent — for blocks, functions, loops)
func NewEnclosedEnvironment(outer *Environment) *Environment
```

### Variable Operations

```go
// Define a new variable in current scope
func (e *Environment) Define(name string, val VintObject) VintObject

// Define a constant (cannot be reassigned)
func (e *Environment) DefineConst(name string, val VintObject) VintObject

// Reassign an existing variable (walks scope chain)
func (e *Environment) Assign(name string, val VintObject) (VintObject, bool)
// Returns false if variable doesn't exist in any scope
// Returns error if variable is a constant

// Set in current scope only (for struct fields, etc.)
func (e *Environment) SetScoped(name string, val VintObject) VintObject
```

### Function Overloading

VintLang supports function overloading — multiple functions with the same name but different parameter counts:

```go
func (e *Environment) Define(name string, val VintObject) VintObject {
    if fn, ok := val.(*Function); ok {
        e.funcs[name] = append(e.funcs[name], fn)  // Add to overload list
    }
    e.store[name] = val
    return val
}

func (e *Environment) GetAllFunctions(name string) []*Function {
    if fns, ok := e.funcs[name]; ok {
        return fns
    }
    if e.outer != nil {
        return e.outer.GetAllFunctions(name)
    }
    return nil
}
```

### Deferred Calls

Functions can register calls to be executed when the function returns:

```go
func (e *Environment) AddDefer(dc *DeferredCall) {
    e.deferMu.Lock()
    defer e.deferMu.Unlock()
    e.deferredCalls = append(e.deferredCalls, dc)
}

func (e *Environment) PopDefers() []*DeferredCall {
    e.deferMu.Lock()
    defer e.deferMu.Unlock()
    defers := e.deferredCalls
    e.deferredCalls = nil
    return defers
}
```

### Scope Example

```
Global Environment (store: {print: Builtin, len: Builtin, ...})
  └── Function "greet" Environment (store: {name: "Alice"})
       └── If Block Environment (store: {greeting: "Hello, Alice"})
```

Each nested scope has access to its parent's variables but not vice versa.

---

## 10. Evaluator — `evaluator/`

**Files: `evaluator/evaluator.go` + 38 additional files** (39 files total, ~3500 lines)

The evaluator is the core of the interpreter. It walks the AST tree and executes each node, producing runtime values.

### Singleton Values

To avoid allocating new objects for common values, the evaluator pre-creates singletons:

```go
var (
    NULL     = &object.Null{}
    TRUE     = &object.Boolean{Value: true}
    FALSE    = &object.Boolean{Value: false}
    BREAK    = &object.Break{}
    CONTINUE = &object.Continue{}
)
```

### The Main Dispatcher — `Eval()`

The `Eval()` function is a massive type switch that dispatches to specific evaluation functions based on the AST node type:

```go
func Eval(node ast.Node, env *object.Environment) object.VintObject {
    switch node := node.(type) {

    // ── Literals ──
    case *ast.IntegerLiteral:
        return &object.Integer{Value: node.Value}
    case *ast.FloatLiteral:
        return &object.Float{Value: node.Value}
    case *ast.StringLiteral:
        return &object.String{Value: node.Value}
    case *ast.Boolean:
        return nativeBoolToBooleanObject(node.Value)
    case *ast.Null:
        return NULL
    case *ast.ArrayLiteral:
        elements := evalExpressions(node.Elements, env)
        return &object.Array{Elements: elements}
    case *ast.DictLiteral:
        return evalDictLiteral(node, env)

    // ── Statements ──
    case *ast.Program:
        return evalProgram(node, env)
    case *ast.LetStatement:
        val := Eval(node.Value, env)
        env.Define(node.Name.Value, val)
        return val
    case *ast.ConstStatement:
        val := Eval(node.Value, env)
        env.DefineConst(node.Name.Value, val)
        return val
    case *ast.ReturnStatement:
        val := Eval(node.ReturnValue, env)
        return &object.ReturnValue{Value: val}
    case *ast.BlockStatement:
        return evalBlockStatement(node, env)
    case *ast.ExpressionStatement:
        return Eval(node.Expression, env)

    // ── Expressions ──
    case *ast.PrefixExpression:
        right := Eval(node.Right, env)
        return evalPrefixExpression(node.Operator, right, node.Token.Line)
    case *ast.InfixExpression:
        left := Eval(node.Left, env)
        right := Eval(node.Right, env)
        return evalInfixExpression(node.Operator, left, right, node.Token.Line)
    case *ast.PostfixExpression:
        return evalPostfixExpression(node, env)
    case *ast.IfExpression:
        return evalIfExpression(node, env)

    // ── Functions ──
    case *ast.FunctionLiteral:
        return evalFunctionLiteral(node, env)
    case *ast.CallExpression:
        return evalCall(node, env)
    case *ast.MethodExpression:
        return evalMethodExpression(node, env)

    // ── Control Flow ──
    case *ast.ForIn:
        return evalForInExpression(node, env, node.Token.Line)
    case *ast.WhileExpression:
        return evalWhileExpression(node, env)
    case *ast.SwitchExpression:
        return evalSwitchExpression(node, env)
    case *ast.MatchExpression:
        return evalMatchExpression(node, env)
    case *ast.BreakStatement:
        return BREAK
    case *ast.ContinueStatement:
        return CONTINUE

    // ── Advanced Features ──
    case *ast.Identifier:
        return evalIdentifier(node, env)
    case *ast.Assign:
        return evalAssign(node, env)
    case *ast.IndexExpression:
        return evalIndexExpression(node, env)
    case *ast.DotExpression:
        return evalPropertyExpression(node, env)
    case *ast.Import:
        return evalImport(node, env)
    case *ast.ImportStatement:
        return evalImportStatement(node, env)
    case *ast.EnumStatement:
        return evalEnumStatement(node, env)
    case *ast.StructStatement:
        return evalStructStatement(node, env)
    case *ast.DeferStatement:
        return evalDeferStatement(node, env)
    // ... more cases
    }
}
```

### Program Evaluation

```go
func evalProgram(program *ast.Program, env *object.Environment) object.VintObject {
    var result object.VintObject
    for _, stmt := range program.Statements {
        result = Eval(stmt, env)

        switch result := result.(type) {
        case *object.ReturnValue:
            return result.Value    // Unwrap return at top level
        case *object.Error:
            return result          // Stop on error
        }
    }

    // If a main() function was defined, call it automatically
    if mainFn, ok := env.Get("main"); ok {
        if fn, ok := mainFn.(*object.Function); ok {
            result = applyFunction(fn, []object.VintObject{}, 0)
        }
    }

    return result
}
```

### Key Evaluation Patterns

#### Identifier Lookup (`evaluator/identifier.go`)

```go
func evalIdentifier(node *ast.Identifier, env *object.Environment) object.VintObject {
    // 1. Check the environment scope chain
    if val, ok := env.Get(node.Value); ok {
        return val
    }
    // 2. Check built-in functions
    if builtin, ok := builtins.GetBuiltin(node.Value); ok {
        return builtin
    }
    // 3. Not found — error
    return newError("Line %d: identifier not found: %s", node.Token.Line, node.Value)
}
```

#### If Expression (`evaluator/if.go`)

```go
func evalIfExpression(ie *ast.IfExpression, env *object.Environment) object.VintObject {
    condition := Eval(ie.Condition, env)
    if isTruthy(condition) {
        return Eval(ie.Consequence, env)
    } else if ie.Alternative != nil {
        return Eval(ie.Alternative, env)
    }
    return NULL
}
```

#### Infix (Binary) Operations (`evaluator/infix.go`)

The infix evaluator handles arithmetic, comparison, logical, and special operations:

```go
func evalInfixExpression(op string, left, right object.VintObject, line int) object.VintObject {
    switch {
    // Integer + Integer
    case left.Type() == object.INTEGER_OBJ && right.Type() == object.INTEGER_OBJ:
        return evalIntegerInfixExpression(op, left, right, line)
    // Float + Float
    case left.Type() == object.FLOAT_OBJ && right.Type() == object.FLOAT_OBJ:
        return evalFloatInfixExpression(op, left, right, line)
    // Mixed Integer/Float (auto-promote to float)
    case left.Type() == object.FLOAT_OBJ && right.Type() == object.INTEGER_OBJ:
        return evalFloatIntegerInfixExpression(op, left, right, line)
    // String operations
    case left.Type() == object.STRING_OBJ && right.Type() == object.STRING_OBJ:
        return evalStringInfixExpression(op, left, right, line)
    // String * Integer (repetition)
    case left.Type() == object.STRING_OBJ && right.Type() == object.INTEGER_OBJ:
        return &object.String{Value: strings.Repeat(leftStr, int(rightInt))}
    // Array + Array (concatenation)
    case left.Type() == object.ARRAY_OBJ && right.Type() == object.ARRAY_OBJ:
        // Concatenate element slices
    // Null coalescing (??)
    case op == "??":
        if left.Type() == object.NULL_OBJ { return right }
        return left
    // Boolean logic
    case op == "&&":
        return nativeBoolToBooleanObject(isTruthy(left) && isTruthy(right))
    case op == "||":
        return nativeBoolToBooleanObject(isTruthy(left) || isTruthy(right))
    // Equality
    case op == "==":
        return nativeBoolToBooleanObject(left == right)  // Pointer comparison for singletons
    case op == "!=":
        return nativeBoolToBooleanObject(left != right)
    }
}
```

**Special arithmetic behaviors:**
- Division of integers returns an `Integer` if the result is exact, otherwise a `Float`
- Power (`**`) always returns a `Float`
- Division by zero returns an error

#### For-In Loops (`evaluator/forin.go`)

```go
func evalForInExpression(fie *ast.ForIn, env *object.Environment, line int) object.VintObject {
    iterable := Eval(fie.Iterable, env)

    // Create isolated iterator (snapshot) to prevent nested loop issues
    iter := createIsolatedIterator(iterable.(object.Iterable))
    iter.Reset()

    return loopIterable(iter.Next, env, fie, line)
}

func loopIterable(next func() (object.VintObject, object.VintObject),
                  env *object.Environment, fi *ast.ForIn, line int) object.VintObject {
    for {
        key, value := next()
        if key == nil { break }  // Iterator exhausted

        // Create new scope for this iteration
        loopEnv := object.NewEnclosedEnvironment(env)
        loopEnv.Define(fi.Key, key)
        if fi.Value != "" {
            loopEnv.Define(fi.Value, value)
        }

        // Evaluate loop body
        result := Eval(fi.Body, loopEnv)

        // Handle control flow
        switch result.(type) {
        case *object.Break:
            return NULL      // Exit loop
        case *object.Continue:
            continue         // Next iteration
        case *object.ReturnValue:
            return result    // Propagate return up
        case *object.Error:
            return result    // Propagate error up
        }
    }
    return NULL
}
```

#### Function Calls (`evaluator/call.go`)

Function calls involve overload resolution, argument evaluation, and application:

```go
func evalCall(node *ast.CallExpression, env *object.Environment) object.VintObject {
    // 1. Evaluate the function expression
    function := Eval(node.Function, env)

    // 2. Check for function overloads
    if ident, ok := node.Function.(*ast.Identifier); ok {
        overloads := env.GetAllFunctions(ident.Value)
        if len(overloads) > 1 {
            // Find overload matching argument count
            for _, fn := range overloads {
                if len(node.Arguments) == len(fn.Parameters) {
                    function = fn
                    break
                }
            }
        }
    }

    // 3. Evaluate arguments
    args := evalExpressions(node.Arguments, env)

    // 4. Apply function
    return applyFunction(function, args, node.Token.Line)
}
```

#### Function Application

```go
func applyFunction(fn object.VintObject, args []object.VintObject, line int) object.VintObject {
    switch fn := fn.(type) {
    case *object.Function:
        // Create closure scope
        extendedEnv := extendedFunctionEnv(fn, args)
        extendedEnv.SetFuncScope(true)  // Mark for defer tracking

        // Evaluate body
        result := Eval(fn.Body, extendedEnv)

        // Execute deferred calls (LIFO order)
        defers := extendedEnv.PopDefers()
        for i := len(defers) - 1; i >= 0; i-- {
            applyFunction(defers[i].Function, defers[i].Args, line)
        }

        return unwrapReturnValue(result)

    case *object.Builtin:
        return fn.Fn(args...)  // Call Go function directly

    case *object.Struct:
        return instantiateStruct(fn, args, line)

    // ... more cases
    }
}

func extendedFunctionEnv(fn *object.Function, args []object.VintObject) *object.Environment {
    env := object.NewEnclosedEnvironment(fn.Env)  // Closure!
    for i, param := range fn.Parameters {
        if i < len(args) {
            env.Define(param.Value, args[i])
        } else if defaultExpr, ok := fn.Defaults[param.Value]; ok {
            env.Define(param.Value, Eval(defaultExpr, fn.Env))  // Use default
        }
    }
    return env
}
```

Note how closures work: the new environment's `outer` is set to `fn.Env` (the environment where the function was *defined*, not where it was *called*). This is what makes closures capture their defining scope.

#### Method Calls (`evaluator/method.go`)

Method calls dispatch based on the object type:

```go
func applyMethod(obj object.VintObject, method ast.Expression, args []object.VintObject, ...) object.VintObject {
    methodName := method.(*ast.Identifier).Value

    switch obj := obj.(type) {
    case *object.String:
        return obj.Method(methodName, args)     // e.g., "hello".upper()
    case *object.Array:
        // Special handling for map, filter, sortBy (need eval access)
        if methodName == "map" { return maap(obj, args) }
        if methodName == "filter" { return filter(obj, args) }
        return obj.Method(methodName, args)     // e.g., [1,2].push(3)
    case *object.Dict:
        // Special handling for functional methods
        return obj.Method(methodName, args)
    case *object.Module:
        fn := obj.Functions[methodName]         // e.g., os.readFile()
        return fn.Fn(args...)
    case *object.StructInstance:
        return callStructMethod(obj, methodName, args, ...)
    // ... more types
    }
}
```

#### While Loops (`evaluator/while.go`)

```go
func evalWhileExpression(we *ast.WhileExpression, env *object.Environment) object.VintObject {
    for {
        condition := Eval(we.Condition, env)
        if !isTruthy(condition) {
            break  // Condition is false — exit loop
        }

        result := Eval(we.Body, env)

        switch result.(type) {
        case *object.Break:
            return NULL
        case *object.Continue:
            continue
        case *object.ReturnValue:
            return result
        case *object.Error:
            return result
        }
    }
    return NULL
}
```

#### Block Evaluation (`evaluator/block.go`)

```go
func evalBlockStatement(block *ast.BlockStatement, env *object.Environment) object.VintObject {
    var result object.VintObject
    for _, stmt := range block.Statements {
        result = Eval(stmt, env)

        if result != nil {
            rt := result.Type()
            // Stop early for control flow signals
            if rt == object.RETURN_VALUE_OBJ || rt == object.ERROR_OBJ ||
               rt == object.CONTINUE_OBJ || rt == object.BREAK_OBJ {
                return result
            }
        }
    }
    return result
}
```

---

## 11. Built-in Functions — `evaluator/builtins/`

**Files: 11 files**

Built-in functions are implemented in Go and registered in a central registry.

### Registry (`registry.go`)

```go
var BuiltinRegistry = make(map[string]*object.Builtin)

func RegisterBuiltin(name string, builtin *object.Builtin) {
    BuiltinRegistry[name] = builtin
}

func GetBuiltin(name string) (*object.Builtin, bool) {
    b, ok := BuiltinRegistry[name]
    return b, ok
}
```

### Core Built-ins (`core.go`)

| Function | Description |
|----------|-------------|
| `print(args...)` | Print values to stdout (space-separated) |
| `println(args...)` | Print values with trailing newline |
| `input(prompt)` | Read a line from stdin |
| `type(obj)` | Returns the type name as a string |
| `string(obj)` / `str(obj)` | Convert to string |

### Array Built-ins (`arrays.go`)

| Function | Description |
|----------|-------------|
| `range(n)` | Generate array `[0, 1, 2, ..., n-1]` |
| `append(arr, val)` | Add element to array |
| `pop(arr)` | Remove and return last element |
| `shift(arr)` | Remove and return first element |
| `len(obj)` | Length of array, string, or dict |
| `indexOf(arr, val)` | Find index of element |
| `last(arr)` | Get last element |

### Type Conversion (`type_conversion.go`)

| Function | Description |
|----------|-------------|
| `int(val)` | Convert to integer |
| `float(val)` | Convert to float |
| `str(val)` | Convert to string |
| `bool(val)` | Convert to boolean |

### I/O Built-ins (`io.go`)

Functions for file operations, usually delegated to the `os` module.

### Channel Built-ins (`channels.go`)

Functions for goroutine-based concurrency support.

---

## 12. Module System — `module/`

**File: `module/module.go`** + 50+ module implementation files

Modules provide VintLang's standard library. Each module is a collection of built-in functions grouped by functionality.

### Module Registry

```go
var Mapper = map[string]*object.Module{}

func init() {
    Mapper["os"]        = &object.Module{Name: "os",        Functions: OsFunctions}
    Mapper["time"]      = &object.Module{Name: "time",      Functions: TimeFunctions}
    Mapper["json"]      = &object.Module{Name: "json",      Functions: JSONFunctions}
    Mapper["math"]      = &object.Module{Name: "math",      Functions: MathFunctions}
    Mapper["http"]      = &object.Module{Name: "http",      Functions: HTTPFunctions}
    Mapper["string"]    = &object.Module{Name: "string",    Functions: StringFunctions}
    Mapper["regex"]     = &object.Module{Name: "regex",     Functions: RegexFunctions}
    Mapper["crypto"]    = &object.Module{Name: "crypto",    Functions: CryptoFunctions}
    Mapper["sqlite"]    = &object.Module{Name: "sqlite",    Functions: SQLiteFunctions}
    Mapper["mysql"]     = &object.Module{Name: "mysql",     Functions: MySQLFunctions}
    Mapper["postgres"]  = &object.Module{Name: "postgres",  Functions: PostgresFunctions}
    Mapper["redis"]     = &object.Module{Name: "redis",     Functions: RedisFunctions}
    Mapper["csv"]       = &object.Module{Name: "csv",       Functions: CSVFunctions}
    Mapper["yaml"]      = &object.Module{Name: "yaml",      Functions: YAMLFunctions}
    Mapper["xml"]       = &object.Module{Name: "xml",       Functions: XMLFunctions}
    Mapper["uuid"]      = &object.Module{Name: "uuid",      Functions: UUIDFunctions}
    Mapper["dotenv"]    = &object.Module{Name: "dotenv",    Functions: DotenvFunctions}
    Mapper["net"]       = &object.Module{Name: "net",       Functions: NetFunctions}
    Mapper["email"]     = &object.Module{Name: "email",     Functions: EmailFunctions}
    Mapper["jwt"]       = &object.Module{Name: "jwt",       Functions: JWTFunctions}
    Mapper["clipboard"] = &object.Module{Name: "clipboard", Functions: ClipboardFunctions}
    Mapper["excel"]     = &object.Module{Name: "excel",     Functions: ExcelFunctions}
    Mapper["llm"]       = &object.Module{Name: "llm",       Functions: LLMFunctions}
    // ... 40+ total modules
}
```

### How Imports Work (`evaluator/import.go`)

```go
func evalImport(node *ast.Import, env *object.Environment) object.VintObject {
    name := node.Value  // e.g., "os", "json", "./mylib"

    // 1. Check built-in modules first
    if mod, ok := module.Mapper[name]; ok {
        env.Define(name, mod)  // Register in scope
        return mod
    }

    // 2. Search for .vint file on disk
    return evalImportFile(name, env)
}
```

File imports search these locations:
1. Bundled files (for packaged executables)
2. Current directory
3. Directories in the search path (added via `AddSearchPath()`)

Each module file has the standard `.vint` extension and is evaluated in its own scope. The resulting environment becomes the module object.

### Module Error Helper

Modules use a consistent error message format:

```go
func ErrorMessage(moduleName, funcName string, expectedArgs int, receivedArgs int, usage string) string {
    return fmt.Sprintf(
        "%s.%s: expected %d argument(s), got %d\nUsage: %s",
        moduleName, funcName, expectedArgs, receivedArgs, usage,
    )
}
```

---

## 13. REPL — `repl/`

**Files: `repl/repl.go`, `repl/playground.go`, `repl/docs.go`**

The REPL (Read-Eval-Print Loop) provides an interactive VintLang session.

### Starting the REPL

```go
func Start() {
    // Create persistent environment
    env := object.NewEnvironment()

    // Set up prompt with go-prompt library
    d := &dummy{env: env}
    p := prompt.New(
        d.executor,
        completer,
        prompt.OptionPrefix(">>> "),
        prompt.OptionTitle("VintLang REPL"),
    )
    p.Run()
}
```

### REPL Execution Loop

```go
func (d *dummy) executor(input string) {
    // Handle exit
    if strings.TrimSpace(input) == "exit()" {
        os.Exit(0)
    }

    // Run through the full pipeline
    l := lexer.New(input)
    p := parser.New(l)
    program := p.ParseProgram()

    if len(p.Errors()) > 0 {
        printParserErrors(p.Errors())
        return
    }

    result := evaluator.Eval(program, d.env)

    // Print result if non-null
    if result != nil && result.Type() != object.NULL_OBJ {
        fmt.Println(result.Inspect())
    }
}
```

The environment persists across inputs, so variables defined in one line are available in the next.

### File Execution via REPL

```go
func ReadWithFilename(contents, filename string) {
    env := object.NewEnvironment()
    l := lexer.NewWithFilename(contents, filename)
    p := parser.New(l)
    program := p.ParseProgram()

    if len(p.Errors()) > 0 {
        for _, e := range p.Errors() {
            fmt.Println(e)
        }
        return
    }

    result := evaluator.Eval(program, env)
    if result != nil && result.Type() == object.ERROR_OBJ {
        fmt.Println(result.Inspect())
    }
}
```

### Playground Mode

The `Playground()` function launches a full TUI (Terminal User Interface) using the Bubble Tea framework, providing a richer interactive experience with syntax highlighting and styled output.

---

## 14. Error Handling — `vintErrors/`

**File: `vintErrors/errors.go`**

VintLang has a structured error system with error codes, severity levels, and contextual information.

### Error Codes

Errors are categorized by their source in the pipeline:

| Range | Source | Examples |
|-------|--------|----------|
| `E1xx` | Lexer | E100: Illegal character, E101: Unterminated string, E102: Invalid escape |
| `E2xx` | Parser | E200: Unexpected token, E201: Missing token, E202: Invalid syntax |
| `E3xx` | Semantic | E300: Undeclared variable, E301: Type mismatch, E302: Invalid operation |
| `E4xx` | Runtime | E400: Index out of bounds, E401: Wrong arguments, E402: Null reference |

### Error Structure

```go
type VintError struct {
    Code       ErrorCode   // E100, E200, etc.
    Severity   Severity    // ERROR, WARNING, INFO
    Message    string      // Human-readable description
    Line       int         // Source line number
    Column     int         // Source column
    Source     string      // Source file name
    Context    string      // Surrounding source code
    Suggestion string      // Helpful fix suggestion
}
```

### Error Formatting

```go
func (e *VintError) Error() string {
    // Format: [SEVERITY CODE] Line:Col: Message
    // Includes source context with caret pointing to error location
    // Includes suggestion if available
}
```

Example error output:

```
[ERROR E300] Line 5, Col 12:
    let result = unknownVar + 1
                 ^^^^^^^^^^
    Undeclared variable: unknownVar
    Suggestion: Did you mean to declare it with 'let unknownVar = ...'?
```

### Builder Pattern

```go
err := vintErrors.NewError(vintErrors.E300, vintErrors.ERROR, "Undeclared variable: x", 5, 12, "main.vint")
err = err.WithContext("let result = x + 1")
err = err.WithSuggestion("Did you mean to declare 'x' first?")
```

### Error Propagation in the Evaluator

The evaluator uses a simple convention: if any evaluation returns an `*object.Error`, it's immediately propagated up the call stack:

```go
func isError(obj object.VintObject) bool {
    return obj != nil && obj.Type() == object.ERROR_OBJ
}

// Used throughout the evaluator:
result := Eval(node, env)
if isError(result) {
    return result  // Stop and propagate
}
```

---

## 15. End-to-End Execution Example

Let's trace the complete execution of this VintLang program:

```vint
let greet = func(name) {
    return "Hello, " + name + "!"
}
print(greet("World"))
```

### Stage 1: Lexing

The lexer produces this token stream:

```
Token{ Type: LET,       Literal: "let",     Line: 1, Col: 1  }
Token{ Type: IDENT,     Literal: "greet",   Line: 1, Col: 5  }
Token{ Type: ASSIGN,    Literal: "=",       Line: 1, Col: 11 }
Token{ Type: FUNCTION,  Literal: "func",    Line: 1, Col: 13 }
Token{ Type: LPAREN,    Literal: "(",       Line: 1, Col: 17 }
Token{ Type: IDENT,     Literal: "name",    Line: 1, Col: 18 }
Token{ Type: RPAREN,    Literal: ")",       Line: 1, Col: 22 }
Token{ Type: LBRACE,    Literal: "{",       Line: 1, Col: 24 }
Token{ Type: RETURN,    Literal: "return",  Line: 2, Col: 5  }
Token{ Type: STRING,    Literal: "Hello, ", Line: 2, Col: 12 }
Token{ Type: PLUS,      Literal: "+",       Line: 2, Col: 22 }
Token{ Type: IDENT,     Literal: "name",    Line: 2, Col: 24 }
Token{ Type: PLUS,      Literal: "+",       Line: 2, Col: 29 }
Token{ Type: STRING,    Literal: "!",       Line: 2, Col: 31 }
Token{ Type: RBRACE,    Literal: "}",       Line: 3, Col: 1  }
Token{ Type: IDENT,     Literal: "print",   Line: 4, Col: 1  }
Token{ Type: LPAREN,    Literal: "(",       Line: 4, Col: 6  }
Token{ Type: IDENT,     Literal: "greet",   Line: 4, Col: 7  }
Token{ Type: LPAREN,    Literal: "(",       Line: 4, Col: 12 }
Token{ Type: STRING,    Literal: "World",   Line: 4, Col: 13 }
Token{ Type: RPAREN,    Literal: ")",       Line: 4, Col: 20 }
Token{ Type: RPAREN,    Literal: ")",       Line: 4, Col: 21 }
Token{ Type: EOF,       Literal: "",        Line: 4, Col: 22 }
```

### Stage 2: Parsing

The parser builds this AST:

```
Program
├── LetStatement
│   ├── Name: Identifier("greet")
│   └── Value: FunctionLiteral
│       ├── Parameters: [Identifier("name")]
│       └── Body: BlockStatement
│           └── ReturnStatement
│               └── ReturnValue: InfixExpression
│                   ├── Left: InfixExpression
│                   │   ├── Left: StringLiteral("Hello, ")
│                   │   ├── Operator: "+"
│                   │   └── Right: Identifier("name")
│                   ├── Operator: "+"
│                   └── Right: StringLiteral("!")
│
└── ExpressionStatement
    └── CallExpression
        ├── Function: Identifier("print")
        └── Arguments:
            └── CallExpression
                ├── Function: Identifier("greet")
                └── Arguments:
                    └── StringLiteral("World")
```

### Stage 3: Evaluation

The evaluator walks the tree:

1. **`Program`** → Evaluate each statement in order.

2. **`LetStatement("greet")`**:
   - Evaluate value: `FunctionLiteral` → Creates `object.Function{Name: "greet", Parameters: ["name"], Body: ..., Env: globalEnv}`
   - Define in environment: `globalEnv.Define("greet", Function{...})`

3. **`ExpressionStatement`** → Evaluate the `CallExpression`:

4. **`CallExpression(print, [CallExpression(greet, ["World"])])`**:
   - Evaluate `print` → `Builtin{Fn: printFn}` (from built-in registry)
   - Evaluate arguments:
     - **`CallExpression(greet, ["World"])`**:
       - Evaluate `greet` → `Function{...}` (from environment)
       - Evaluate arguments: `"World"` → `String{Value: "World"}`
       - **Apply function**:
         - Create enclosed environment: `funcEnv = {outer: globalEnv}`
         - Bind parameter: `funcEnv.Define("name", String{"World"})`
         - Evaluate body:
           - **`ReturnStatement`**:
             - Evaluate `InfixExpression(InfixExpression("Hello, ", "+", name), "+", "!")`
               - Inner infix: `"Hello, " + "World"` → `String{"Hello, World"}`
               - Outer infix: `"Hello, World" + "!"` → `String{"Hello, World!"}`
             - Return: `ReturnValue{Value: String{"Hello, World!"}}`
         - Unwrap return: `String{"Hello, World!"}`
   - Call `print(String{"Hello, World!"})` → Prints to stdout

### Output

```
Hello, World!
```

---

## 16. Additional Components

### Bundler (`bundler/`)

The bundler packages a VintLang program into a standalone binary. It embeds the source code and a minimal VintLang runtime into a single executable.

### Code Formatter (`vint fmt`)

The formatter round-trips source code through the pipeline: `source → lexer → parser → AST → AST.String()`. The `String()` methods on AST nodes produce consistently formatted output.

### Testing Infrastructure

Tests are located alongside the code they test:

- `lexer/lexer_test.go` — Tests token generation for all token types
- `parser/parser_test.go` — Tests AST construction for all syntax forms
- `evaluator/evaluator_test.go` — Tests evaluation results for expressions and statements
- `evaluator/forin_test.go` — Tests for-in loop behavior
- `evaluator/enum_test.go` — Tests enum definitions and access
- `evaluator/struct_test.go` — Tests struct instantiation and methods
- `object/strings_test.go` — Tests string method implementations

Run tests with:
```bash
go test ./...
```

### Configuration (`config/`)

Currently contains just the version constant:
```go
const VINT_VERSION = "0.2.5"
```

### CLI Toolkit (`toolkit/`)

Utilities for project management:
- `init` — Scaffolds a new VintLang project
- `get` — Installs VintLang packages
- `test` — Discovers and runs `.vint` test files
- `docs` — Interactive documentation browser

### Dependencies

VintLang is built with Go 1.21+ and uses these key dependencies:

| Category | Package | Purpose |
|----------|---------|---------|
| CLI/REPL | `github.com/c-bata/go-prompt` | Interactive REPL with autocomplete |
| TUI | `github.com/charmbracelet/bubbletea` | Terminal UI framework |
| Styling | `github.com/charmbracelet/lipgloss` | Terminal styling |
| Database | `github.com/mattn/go-sqlite3` | SQLite driver |
| Database | `github.com/go-sql-driver/mysql` | MySQL driver |
| Database | `github.com/lib/pq` | PostgreSQL driver |
| Cache | `github.com/redis/go-redis` | Redis client |
| Auth | `github.com/golang-jwt/jwt` | JWT tokens |
| Data | `gopkg.in/yaml.v3` | YAML parsing |
| Data | `github.com/xuri/excelize` | Excel files |
| UUID | `github.com/google/uuid` | UUID generation |
| System | `github.com/shirou/gopsutil` | System information |
| Net | `github.com/gorilla/websocket` | WebSocket support |

---

## Summary

VintLang is a **tree-walking interpreter** implemented in Go. The architecture follows the classic interpreter pattern:

1. **Lexer** scans source text into tokens with full Unicode support
2. **Parser** uses Pratt parsing to build an AST with correct operator precedence
3. **Evaluator** walks the AST recursively, maintaining scope through chained environments
4. **Object System** represents all runtime values uniformly through the `VintObject` interface

Key design decisions:
- **Pratt parsing** allows clean, modular handling of complex expression syntax
- **Environment chains** implement lexical scoping and closures naturally
- **VintObject interface** provides a uniform type system for all values
- **Singleton objects** (`NULL`, `TRUE`, `FALSE`) reduce memory allocation
- **Function overloading** via the `funcs` map in Environment
- **Module system** with 40+ built-in modules provides a rich standard library
- **Error propagation** through return values (no exceptions/panics in normal flow)

The codebase is well-organized with each language feature in its own file, making it straightforward to understand and extend.
