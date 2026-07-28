# VintLang Type System — Implementation Plan

**Version:** 1.0 (implemented)
**Date:** July 2026
**Status:** ✅ All v1 phases complete

---

## Table of Contents

1. [Decisions](#1-decisions)
2. [The Type Set (v1)](#2-the-type-set-v1)
3. [Architecture](#3-architecture)
4. [Phase 0: Lexer](#4-phase-0-lexer)
5. [Phase 1: Parser](#5-phase-1-parser)
6. [Phase 2: Object Layer](#6-phase-2-object-layer)
7. [Phase 3: Runtime Type Enforcement](#7-phase-3-runtime-type-enforcement)
8. [Phase 4: Type Aliases](#8-phase-4-type-aliases)
9. [Phase 5: Strict Mode Enforcement](#9-phase-5-strict-mode-enforcement)
10. [Phase 6: Tests + Update Examples](#10-phase-6-tests--update-examples)
11. [Deferred to v2](#11-deferred-to-v2)
12. [Type Compatibility Table](#12-type-compatibility-table)
13. [Error Messages Reference](#13-error-messages-reference)

---

## Implementation Status

| Phase | Status |
|-------|--------|
| Phase 0: Lexer | ✅ Complete — `type` keyword handled contextually |
| Phase 1: Parser | ✅ Complete — all type forms, as/is, type aliases, struct methods |
| Phase 2: Object Layer | ✅ Complete — types map, DefineTyped, GetDeclaredType |
| Phase 3: Runtime Type Enforcement | ✅ Complete — compatible() checks at all 8 sites |
| Phase 4: Type Aliases | ✅ Complete — resolved during parsing |
| Phase 5: Strict Mode | 🔲 Deferred |
| Phase 6: Tests + Examples | ✅ Complete — 33 parser tests, 15 examples |

**Bonus features added:**
- `::` prefix for builtin function calls
- Type signatures (`ParamTypes`/`ReturnType`) on builtins and module functions
- Improved error messages with source context and line numbers
- All `.vint` examples updated with `::` builtin prefix and type annotations

---

## 1. Decisions

### Locked for v1

| Decision | Choice |
|----------|--------|
| Mutability | `let` is mutable, `const` is immutable. **No `mut` keyword.** |
| Zero values | `let x: int` allowed (no `=` → zero value) |
| Strictness | Strict typing everywhere — every variable needs a type (explicit `: T` OR inferred from initializer) |
| Error model | Go-style `(T, error)` multi-return. No `!T` error union in v1. |
| Type inference | `:=` shorthand deferred to v2. But plain `let x = 5` still infers the type. |
| Optionals | `?T` and `?.` deferred to v2. Use `*T` + nil check, or `any`. |
| Type aliases | `type UserID = int` — included in v1. |
| Variadic | Deferred to v2. |
| Named returns | Deferred to v2. |
| Typed enums | Deferred to v2. |

### Three valid declaration forms (v1)

```vint
let x: int = 5         // explicit type + initializer
let y: int             // explicit type, zero value
let z = 5              // type inferred from initializer (=> int)
```

One invalid form:
```vint
let x                   // ERROR: cannot infer type
```

---

## 2. The Type Set (v1)

### Primitives

| Type | Go equivalent | Zero value | Notes |
|------|--------------|------------|-------|
| `bool` | `bool` | `false` | |
| `string` | `string` | `""` | UTF-8 |
| `int` | `int64` | `0` | platform-sized (64-bit) |
| `int8` | `int8` | `0` | |
| `int16` | `int16` | `0` | |
| `int32` | `int32` | `0` | |
| `int64` | `int64` | `0` | |
| `uint` | `uint64` | `0` | platform-sized unsigned |
| `uint8` | `uint8` | `0` | |
| `uint16` | `uint16` | `0` | |
| `uint32` | `uint32` | `0` | |
| `uint64` | `uint64` | `0` | |
| `byte` | `byte` | `0` | alias for `uint8` |
| `float32` | `float32` | `0.0` | |
| `float64` | `float64` | `0.0` | default float type |

### Special

| Type | Description | Zero value |
|------|-------------|------------|
| `any` | universal type (Go's `interface{}`) | `nil` |
| `error` | error interface (any type with `.message()`) | `nil` |
| `nil` | zero/undefined value | `nil` (self) |

### Compound

| Type | Syntax | Example | Zero value |
|------|--------|---------|------------|
| Slice | `[]T` | `[]int` | `nil` |
| Fixed array | `[n]T` | `[4]byte` | `[T{}, T{}, ...]` |
| Dict | `{K: V}` | `{string: int}` | `nil` |
| Pointer | `*T` | `*int` | `nil` |
| Channel | `chan T` | `chan string` | `nil` |
| Function | `func(params) ret` | `func(int) bool` | `nil` |

### User-defined

| Kind | Syntax | Example | Zero value |
|------|--------|---------|------------|
| Struct | `struct` | `struct Point { x: int; y: int }` | `struct{}` (all zero fields) |
| Enum | `enum` | `enum Status { ... }` | first member |
| Alias | `type Name = T` | `type UserID = int` | zero of `T` |

### Operators

| Operator | Usage | Description |
|----------|-------|-------------|
| `as` | `x as int` | Type cast (runtime checked) |
| `is` | `x is int` | Type check → bool |

### Error pattern

Go-style multi-return with `error` as the common error interface:

```vint
func divide(a: int, b: int): (int, error) {
    if b == 0 {
        return 0, error("division by zero")
    }
    return a / b, nil
}

let result, err := divide(10, 2)
if err != nil {
    println(err.message())
} else {
    println(result)
}
```

---

## 3. Architecture

```
Source ──► Lexer ──► Parser ──► AST (with Type nodes)
                                       │
                              ┌────────┴────────┐
                              ▼                  ▼
                        TypeChecker          Evaluator
                    (compile-time pass)   (runtime checks)
                        │                      │
                        ▼                      ▼
                   Error list              Result / Error
```

### Three layers stay decoupled

| Layer | What it does | Where |
|-------|-------------|-------|
| **Parsing types** | Produces `ast.Type` nodes | `parser/type.go` (new) |
| **Carrying types** | Attaches type info to runtime objects | `object/function.go`, `object/struct.go`, `object/environment.go` |
| **Checking types** | Validates at runtime boundaries | `evaluator/types.go` (new) |
| **Static checker** | (deferred to v2) Pre-evaluation pass | `typechecker/` (future) |

### Key insight: AST foundation already exists

The file `ast/types.go` already defines all the type AST nodes:
- `BasicType`, `ArrayType`, `FunctionType`, `OptionalType`, `DictType`, `UnionType`
- `TypedParameter`, `TypeAnnotation`
- `TypedLetStatement`, `TypedFunctionLiteral`
- `TypeCastExpression`, `TypeCheckExpression`

**Missing from ast/types.go (needs adding):**
- `PointerType`
- `ChannelType`
- `MultiReturnType` (for `(T, error)` syntax)
- `TypeAliasStatement`

---

## 4. Phase 0: Lexer

**Goal:** Add tokens needed for v1.

### Changes

| Token | Source | State | Action | File |
|-------|--------|-------|--------|------|
| `TYPE` | `type` word | not a keyword | Add `"type": TYPE` to `keywords` map | `token/token.go` |
| `PIPE` | `\|` | lexed but no parser fn | Already done by lexer. No action. | `token/token.go` |
| `AS` | `as` word | already a keyword, no parser fn | Already done. Ready for Phase 1.5. | `token/token.go:186` |
| `IS` | `is` word | already a keyword, no parser fn | Already done. Ready for Phase 1.5. | `token/token.go:187` |

### What is NOT needed in v1

- `QUESTION` token (`?`) — deferred to v2
- `WALRUS` token (`:=`) — deferred to v2
- `OPT_DOT` token (`?.`) — deferred to v2
- Backtick strings — deferred to v2
- Hex/binary/octal numbers — deferred to v2

### Files to modify

| File | Change |
|------|--------|
| `token/token.go` | Add `TYPE = "TYPE"` to const block. Add `"type": TYPE` to keywords map. |

---

## 5. Phase 1: Parser

**Goal:** Parse typed syntax into existing `ast.Type` nodes. Evaluator still works (ignores new nodes). Old code keeps running after this phase — it just produces `LetStatement` instead of `TypedLetStatement`.

### Total: ~400-500 lines across 6 sub-phases

---

### 1.1 `parseType()` Helper — NEW FILE `parser/type.go`

```go
func (p *Parser) parseType() ast.Type {
    // Dispatches on curToken to produce the right ast.Type node
}
```

**Parsing rules for each type form:**

| Input tokens | AST node produced | Example |
|--------------|-------------------|---------|
| `int`, `string`, `bool`, `float64`, etc. | `*ast.BasicType{Name: "int"}` | `int` |
| `[` then `]` then type | `*ast.ArrayType{ElementType: ...}` | `[]int` |
| `[` then INT then `]` then type | `*ast.FixedArrayType{Size: N, ElementType: ...}` (NEW node needed) | `[4]byte` |
| `{` then type then `:` then type then `}` | `*ast.DictType{KeyType: ..., ValueType: ...}` | `{string: int}` |
| `*` then type | `*ast.PointerType{BaseType: ...}` (NEW node needed) | `*int` |
| `chan` then type | `*ast.ChannelType{ElementType: ...}` (NEW node needed) | `chan string` |
| `func` then `(` then params then `)` then type | `*ast.FunctionType{Parameters: ..., ReturnType: ...}` | `func(int) bool` |
| `(` then types then `)` | `*ast.MultiReturnType{Types: ...}` (NEW node needed) | `(int, error)` |

**Disambiguation rule:** When the current token is an identifier that could be a type keyword (`int`, `string`, `bool`, `uint8`, etc.), check if it IS a type keyword. This is a new `TypeKeyword` category that can be reused.

**Files to create/modify:**

| File | Action |
|------|--------|
| `parser/type.go` | NEW — `parseType()`, `parseArrayType()`, `parseDictType()`, `parsePointerType()`, `parseChannelType()`, `parseFunctionType()`, `parseMultiReturnType()` |
| `ast/types.go` | Add `PointerType`, `ChannelType`, `FixedArrayType`, `MultiReturnType`, `TypeAliasStatement` nodes |

---

### 1.2 Wire let/const

**Current code** (`parser/statements.go:35-57`):

```go
func (p *Parser) parseLetStatement() *ast.LetStatement {
    stmt := &ast.LetStatement{Token: p.curToken}
    if !p.expectPeek(token.IDENT) { return nil }
    stmt.Name = &ast.Identifier{...}
    if !p.expectPeek(token.ASSIGN) { return nil }  // requires '='
    p.nextToken()
    stmt.Value = p.parseExpression(LOWEST)
    return stmt
}
```

**New logic:**

```
expectPeek(IDENT)
name = curToken

if peekTokenIs(COLON):
    nextToken()             // consume ':'
    declaredType = parseType()
    if peekTokenIs(ASSIGN):
        nextToken()
        value = parseExpression(LOWEST)
        return TypedLetStatement{Name: name, TypeAnnotation: declaredType, Value: value}
    else:
        return TypedLetStatement{Name: name, TypeAnnotation: declaredType}  // zero value

elif peekTokenIs(ASSIGN):
    nextToken()
    value = parseExpression(LOWEST)
    return LetStatement{Name: name, Value: value}  // inferred type at check time

else:
    error("missing type annotation or initializer for variable '%s'", name)
    return nil
```

Same logic for `parseConstStatement`.

**Files to modify:**

| File | Action |
|------|--------|
| `parser/statements.go` | Rewrite `parseLetStatement()`, `parseConstStatement()` |
| `ast/statements.go` | Add `Type` field to `LetStatement`/`ConstStatement` for inferred type storage (optional) |

---

### 1.3 Wire function params and returns

**Current code** (`parser/function.go:33-65`):

```go
func (p *Parser) parseFunctionParameters(lit *ast.FunctionLiteral) bool {
    for !p.peekTokenIs(token.RPAREN) {
        p.nextToken()
        ident := &ast.Identifier{...}
        lit.Parameters = append(lit.Parameters, ident)
        // ... defaults handling ...
    }
}
```

**New logic:**

```
for each param:
    ident := parseIdentifier()
    if peekTokenIs(COLON):
        nextToken()   // consume ':'
        paramType := parseType()
        if peekTokenIs(ASSIGN):
            // default value
        param = TypedParameter{Identifier: ident, Type: paramType, Default: ...}
    else:
        param = TypedParameter{Identifier: ident}  // type unknown (will error in strict mode)

After RPAREN:
    if peekTokenIs(COLON):
        nextToken()
        returnType = parseType()
```

**Return type forms:**

| Input | AST | Meaning |
|-------|-----|---------|
| `: int` | `BasicType` | Single return value |
| `: (int, error)` | `MultiReturnType` | Multiple return values |
| (nothing) | `nil` | No return value (void) |

If `TypedFunctionLiteral` is produced, `evalFunction` stores the typed params + return type on `*object.Function`.

**Files to modify:**

| File | Action |
|------|--------|
| `parser/function.go` | Rewrite `parseFunctionParameters()`, extend `parseFunctionLiteral()` |
| `ast/types.go` | `TypedParameter`, `TypedFunctionLiteral` already exist |
| `parser/struct.go` (implied: `parser/statements.go` `parseStructMethod`) | Same treatment for method params/returns |

---

### 1.4 Wire struct field types

**Current code** (`parser/statements.go:237-249`):

```go
field := ast.StructField{}
field.Name = &ast.Identifier{Token: p.curToken, Value: p.curToken.Literal}
if p.peekTokenIs(token.COLON) {
    p.nextToken()
    p.nextToken()
    field.Default = p.parseExpression(LOWEST)   // always treated as default value
}
```

**New logic:**

```
field := ast.StructField{}

// Check if next is COLON (might be type OR default)
if peekTokenIs(COLON):
    nextToken()      // consume ':'

    // Disambiguation:
    //   field: int          → type annotation
    //   field: "default"    → default value expression
    //   field: 0            → default value expression
    //   field: OtherStruct  → if OtherStruct is a known type, it's a type annotation
    //                         otherwise, it's a default value

    if peekTokenIsTypeKeyword() || peekTokenIs(token.LBRACKET) || peekTokenIs(token.LBRACE):
        // It's a type annotation
        field.Type = p.parseType()
        if peekTokenIs(ASSIGN):
            nextToken()
            field.Default = p.parseExpression(LOWEST)
    else:
        // It's a default value expression (back-compat)
        field.Default = p.parseExpression(LOWEST)
```

**Type keywords for disambiguation:** `bool`, `string`, `int`, `int8`, `int16`, `int32`, `int64`, `uint`, `uint8`, `uint16`, `uint32`, `uint64`, `byte`, `float32`, `float64`, `any`, `error`, `nil`, `func`, `chan`, `struct`, `enum`.

**Files to modify:**

| File | Action |
|------|--------|
| `parser/statements.go` | Rewrite struct field parsing in `parseStructStatement()` |
| `ast/declaratives.go` | Add `Type ast.Type` field to `StructField` struct (line 183) |
| `parser/type.go` | Add `isTypeKeyword()` helper |

---

### 1.5 Wire `as` and `is` operators

Tokens `AS` and `IS` already exist in token.go. AST nodes `TypeCastExpression` and `TypeCheckExpression` already exist in ast/types.go.

**Parser wiring** (`parser/parser.go` in `New()`):

```go
// Type cast: x as int
p.registerInfix(token.AS, p.parseTypeCast)
p.registerInfix(token.IS, p.parseTypeCheck)
```

**Parse functions:**

```go
func (p *Parser) parseTypeCast(left ast.Expression) ast.Expression {
    tok := p.curToken
    p.nextToken()
    targetType := p.parseType()
    return &ast.TypeCastExpression{Token: tok, Expression: left, TargetType: targetType}
}

func (p *Parser) parseTypeCheck(left ast.Expression) ast.Expression {
    tok := p.curToken
    p.nextToken()
    checkType := p.parseType()
    return &ast.TypeCheckExpression{Token: tok, Expression: left, CheckType: checkType}
}
```

**Precedence:** `AS` and `IS` should be LOWEST (or EQUALS-level) to parse `x + y as int` as `(x + y) as int`.

**Files to modify:**

| File | Action |
|------|--------|
| `parser/parser.go` | Add `registerInfix` calls in `New()`. Add precedence entries for `AS`/`IS`. |
| `parser/type.go` | NEW — add `parseTypeCast()`, `parseTypeCheck()` functions |

---

### 1.6 Wire type aliases

**New keyword:** `type` → token `TYPE`.

**Syntax:** `type UserID = int`

**New AST node** (add to `ast/types.go`):

```go
type TypeAliasStatement struct {
    Token    token.Token  // the 'type' token
    Name     *Identifier  // alias name
    Target   Type         // the type being aliased
}
func (tas *TypeAliasStatement) statementNode() {}
func (tas *TypeAliasStatement) TokenLiteral() string { return tas.Token.Literal }
func (tas *TypeAliasStatement) String() string {
    return "type " + tas.Name.String() + " = " + tas.Target.String()
}
```

**Parser function** (`parser/type.go`):

```go
func (p *Parser) parseTypeAliasStatement() *ast.TypeAliasStatement {
    stmt := &ast.TypeAliasStatement{Token: p.curToken}
    if !p.expectPeek(token.IDENT) { return nil }
    stmt.Name = &ast.Identifier{Token: p.curToken, Value: p.curToken.Literal}
    if !p.expectPeek(token.ASSIGN) { return nil }
    p.nextToken()
    stmt.Target = p.parseType()
    return stmt
}
```

**`TYPE` prefix registration** in `parser/parser.go`:

```go
// In parseStatement switch or in New() as a prefix for statement-level
p.registerPrefix(token.TYPE, p.parseTypeAliasStatement)
```

Or better — handle it in `parseStatement()`'s switch:

```go
case token.TYPE:
    return p.parseTypeAliasStatement()
```

**Files to modify:**

| File | Action |
|------|--------|
| `token/token.go` | Add `TYPE = "TYPE"` constant and `"type": TYPE` keyword |
| `ast/types.go` | Add `TypeAliasStatement` node |
| `parser/type.go` | Add `parseTypeAliasStatement()` |
| `parser/parser.go` | Register in `parseStatement()` switch |

---

## 6. Phase 2: Object Layer

**Goal:** Extend runtime objects to carry type metadata so the evaluator can check it.

### 2.1 Extend `object.Function`

**File:** `object/function.go`

**Current:**
```go
type Function struct {
    Name        string
    Parameters  []*ast.Identifier
    Defaults    map[string]ast.Expression
    Body        *ast.BlockStatement
    Env         *Environment
    // ...
}
```

**Extended:**
```go
type TypedParam struct {
    Name string
    Type ast.Type    // NEW
}

type Function struct {
    Name        string
    Parameters  []*ast.Identifier
    ParamTypes  []ast.Type         // NEW — parallel to Parameters
    ReturnTypes []ast.Type         // NEW — multi-return support
    Defaults    map[string]ast.Expression
    Body        *ast.BlockStatement
    Env         *Environment
    // ...
}
```

When `evalFunction` processes a `*ast.TypedFunctionLiteral`, it populates `ParamTypes` and `ReturnTypes`.

---

### 2.2 Extend `object.StructField` and `object.StructMethod`

**File:** `object/struct.go`

**Current:**
```go
type StructField struct {
    Name    string
    Default ast.Expression
}
```

**Extended:**
```go
type StructField struct {
    Name    string
    Type    ast.Type       // NEW — may be nil (untyped struct)
    Default ast.Expression
}
```

`StructMethod` gets `ParamTypes` and `ReturnTypes` (mirroring `Function` changes).

---

### 2.3 Extend `object.Environment`

**File:** `object/environment.go`

Add a `types` map parallel to `store`:

```go
type Environment struct {
    store     map[string]VintObject
    types     map[string]ast.Type   // NEW — declared types for each name
    funcs     map[string][]*Function
    constants map[string]bool
    outer     *Environment
    // ...
}
```

**New methods:**

```go
func (e *Environment) DefineTyped(name string, val VintObject, t ast.Type) VintObject {
    // Same as Define but also records the type
}

func (e *Environment) GetDeclaredType(name string) (ast.Type, bool) {
    if t, ok := e.types[name]; ok {
        return t, true
    }
    if e.outer != nil {
        return e.outer.GetDeclaredType(name)
    }
    return nil, false
}
```

**Types walk the closure chain** just like `store`.

---

## 7. Phase 3: Runtime Type Enforcement

**Goal:** When annotations are present, reject type-mismatched operations. Old code (no annotations) keeps running.

### 3.1 Compatibility Function: `compatible()`

**NEW FILE:** `evaluator/types.go`

```go
// compatible checks if a runtime value matches a declared type.
// Types:  declared = what the annotation says (e.g. "int")
//         actual   = the runtime object's Type() value (e.g. object.INTEGER_OBJ)
func compatible(declared ast.Type, actual object.VintObjectType) bool
```

**Rules:**

| Declared type | Accepts runtime type(s) |
|---------------|------------------------|
| `int` | `INTEGER_OBJ` |
| `int8`, `int16`, `int32`, `int64` | `INTEGER_OBJ` (with range check) |
| `uint`, `uint8`, `uint16`, `uint32`, `uint64` | `INTEGER_OBJ` (with non-negative + range check) |
| `byte` | `INTEGER_OBJ` (0-255), `BYTE_OBJ` |
| `float32`, `float64` | `FLOAT_OBJ`, `INTEGER_OBJ` (numeric widening) |
| `string` | `STRING_OBJ` |
| `bool` | `BOOLEAN_OBJ` |
| `any` | everything (escape hatch) |
| `error` | `ERROR_OBJ`, `CUSTOM_ERROR_OBJ` |
| `nil` | `NULL_OBJ` |
| `[]T` | `ARRAY_OBJ` (element type check deferred to v2) |
| `[n]T` | `ARRAY_OBJ` (with length check) |
| `{K: V}` | `DICT_OBJ` (key/value type check deferred to v2) |
| `*T` | `POINTER_OBJ` |
| `chan T` | `CHANNEL_OBJ` |
| `func(P) R` | `FUNCTION_OBJ` (signature match deferred to v2) |
| Struct name | `STRUCT_INSTANCE_OBJ` with matching `instance.Struct.Name` |
| Enum name | `ENUM_OBJ` with matching `enum.Name` |
| Type alias | Resolve alias → check target |

### 3.2 Check Sites

Insert `compatible()` checks at every assignment/declaration/call boundary.

#### Site A: let/const declarations

**File:** `evaluator/evaluator.go:79-94`

```go
case *ast.TypedLetStatement:
    if node.Value != nil {
        val := Eval(node.Value, env)
        if isError(val) { return val }
        if !compatible(node.TypeAnnotation.Type, val.Type()) {
            return newError("type mismatch: cannot assign %s to variable '%s' of type %s",
                val.Type(), node.Name.Value, node.TypeAnnotation.String())
        }
        return env.DefineTyped(node.Name.Value, val, node.TypeAnnotation.Type)
    } else {
        // Zero value
        val := zeroValue(node.TypeAnnotation.Type)
        return env.DefineTyped(node.Name.Value, val, node.TypeAnnotation.Type)
    }

case *ast.LetStatement:  // untyped — allow any (inferred type will be recorded after eval)
    val := Eval(node.Value, env)
    if isError(val) { return val }
    return env.DefineTyped(node.Name.Value, val, inferType(val))
```

Need `zeroValue(ast.Type) object.VintObject` helper that returns the zero object for each type.

#### Site B: Function argument binding

**File:** `evaluator/call.go` (evalArgsExpressions) and `evaluator/evaluator.go` (extendedFunctionEnv)

After resolving each argument to a value, check against `fn.ParamTypes[i]`:

```go
if i < len(fn.ParamTypes) && fn.ParamTypes[i] != nil {
    if !compatible(fn.ParamTypes[i], args[i].Type()) {
        return newError("type mismatch: parameter '%s' expects %s, got %s",
            fn.Parameters[i], fn.ParamTypes[i].String(), args[i].Type())
    }
}
```

#### Site C: Function return value

**File:** `evaluator/evaluator.go` near `ReturnStatement` (line 72-77)

Need to know "what function are we inside?" to get the declared return type.

Option: pass return type through the environment chain. Add `currentReturnType` to the environment or pass it via extended env when creating `extendedFunctionEnv`.

**Simpler approach:** When `applyFunction` runs `Eval(fn.Body, extendedEnv)`, the result is checked:

```go
result := Eval(fn.Body, extendedEnv)
if len(fn.ReturnTypes) > 0 {
    result = unwrapReturnValue(result)
    if !compatible(fn.ReturnTypes[0], result.Type()) {
        return newError("type mismatch: function returns %s, but body returned %s",
            fn.ReturnTypes[0].String(), result.Type())
    }
}
```

For multi-return `(T, error)`, `ReturnTypes` has multiple entries. The return statement must emit multiple values, and the caller must capture them.

#### Site D: Variable assignment

**File:** `evaluator/assign.go:8-19`

```go
func evalAssign(node *ast.Assign, env *object.Environment) object.VintObject {
    val := Eval(node.Value, env)
    if isError(val) { return val }

    // NEW: check declared type
    if declaredType, ok := env.GetDeclaredType(node.Name.Value); ok {
        if !compatible(declaredType, val.Type()) {
            return newError("type mismatch: cannot assign %s to variable '%s' of type %s",
                val.Type(), node.Name.Value, declaredType.String())
        }
    }

    return env.Assign(node.Name.Value, val)
}
```

#### Site E: Struct field set

**File:** `evaluator/property.go:84-90`

```go
// In evalPropertyAssignment for StructInstance:
if fieldType := si.Struct.GetFieldType(prop); fieldType != nil {
    if !compatible(fieldType, val.Type()) {
        return newError("type mismatch: field '%s' expects %s, got %s",
            prop, fieldType.String(), val.Type())
    }
}
return si.SetField(prop, val)
```

`GetFieldType(name string) ast.Type` is new on `*object.Struct`.

#### Site F: Struct instantiation

**File:** `evaluator/struct.go`

In `instantiateStruct()`, `evalStructCall()`, `evalStructLiteral()` — check each provided field value against `field.Type`.

#### Site G: Type cast `as`

NEW evaluator handler:

```go
func evalTypeCast(node *ast.TypeCastExpression, env *object.Environment) object.VintObject {
    val := Eval(node.Expression, env)
    if isError(val) { return val }

    targetType := node.TargetType.String() // e.g. "int", "string", "float64"

    switch targetType {
    case "int":
        return convertToInteger(val)
    case "float64", "float32":
        return convertToFloat(val)
    case "string":
        return convertToString(val)
    case "bool":
        return convertToBoolean(val)
    case "byte":
        return convertToByte(val)
    default:
        // For struct types, check if val is the right kind
        return newError("cannot cast %s to %s", val.Type(), targetType)
    }
}
```

#### Site H: Type check `is`

NEW evaluator handler:

```go
func evalTypeCheck(node *ast.TypeCheckExpression, env *object.Environment) object.VintObject {
    val := Eval(node.Expression, env)
    if isError(val) { return val }

    checkType := node.CheckType.String()
    actual := string(val.Type())

    // Match runtime type name to check type
    if checkType == actual {
        return TRUE
    }
    // Special case: struct instances match by struct name
    if si, ok := val.(*object.StructInstance); ok {
        if si.Struct.Name == checkType {
            return TRUE
        }
    }
    return FALSE
}
```

### 3.3 Multi-return `(T, error)`

This touches the most code.

**Parser changes:**
- `parseMultiReturnType()` — parse `(type1, type2, ...)` as a `MultiReturnType` AST node
- Return statement: support `return expr1, expr2` (comma-separated expressions)
- Let statement: support `let x, y := func()` (multi-variable capture)

**Evaluator changes:**
- `applyFunction` when `len(fn.ReturnTypes) > 1`: instead of returning `VintObject`, return `[]VintObject`
- OR: return a synthetic `Tuple` object from calls, then destructure in let/assign
- Simpler approach: wrap multiple returns in a `TupleObject` with `.first()`, `.second()` methods and destructure syntax
- Simplest approach for v1: multi-return is syntactic sugar that returns a single error object and the first value. But this doesn't match Go's pattern.

**Recommeneded approach for v1:** Use a `ReturnValue` object that can hold multiple values. This is the simplest path:

1. `ReturnStatement` evaluates a list of expressions → produces `*object.MultiReturnValue{Values: []VintObject}`
2. `applyFunction` detects `MultiReturnValue` and returns each value separately
3. The call site destructures: `let result, err := call()` is syntactic sugar for:
   ```
   let __call_result = call()
   let result = __call_result[0]
   let err = __call_result[1]
   ```

**OR for simplicity in v1**, we could use an alternative approach:
- `func divide(a, b: int): (int, error)` returns a single `Result` object (not multiple values)
- `let result = divide(10, 2)` gets the `Result` object
- Users check `.is_ok()`, `.unwrap()`, `.error()`
- This is more Rust/Zig-like than Go-like, but simpler to implement

**Decision needed:** True multi-return (complex) vs Result object (simpler).

### 3.4 Error Messages

Format for type errors:

```
TypeError: expected 'int' but got 'string'
  at file.vint:42

TypeError: cannot assign 'string' to variable 'count' of type 'int'
  at file.vint:15

TypeError: parameter 'b' expects 'int', got 'string'
  at file.vint:8

TypeError: function 'add' returns 'int', but body returned 'string'
  at file.vint:12

TypeError: field 'name' in struct 'User' expects 'string', got 'int'
  at file.vint:25

TypeError: cannot cast 'string' to 'int'
  at file.vint:30
```

---

## 8. Phase 4: Type Aliases

**Goal:** `type UserID = int` makes `UserID` usable as a type annotation anywhere.

### 8.1 Runtime representation

Store aliases in the environment alongside declared types:

```go
// In environment.go
type Environment struct {
    // ... existing fields ...
    aliases map[string]ast.Type   // NEW — type name aliases
}
```

### 8.2 Alias resolution

`compatible()` resolves aliases before checking:

```go
func compatible(declared ast.Type, actual object.VintObjectType) bool {
    // Resolve aliases first
    if alias, ok := declared.(*ast.TypeAlias); ok {
        return compatible(alias.Target, actual)
    }
    // ... rest of matching logic ...
}
```

Maybe simpler: store aliases in the `types` map of the environment, and resolve them during parsing or during checking. During parsing, `parseType()` could resolve `UserID` to `int` immediately if `UserID` is a registered alias.

**Cleanest approach:** During parsing, maintain a table of known aliases. When parsing a type and encountering an identifier that's an alias, replace it with the resolved type immediately. This way, `TypeAlias` nodes need not survive past parsing.

### 8.3 Scoping

Type aliases are scoped like variables — they live in the environment and walk the closure chain. An alias defined inside a function is only visible inside that function.

---

## 9. Phase 5: Strict Mode Enforcement

**Goal:** Every variable must have a type — either explicit (`: int`) or inferable from an initializer (`= 5`).

### 9.1 Parser-level enforcement

In `parseLetStatement()` / `parseConstStatement()`:

```go
if !hasColon && !hasEquals {
    p.addError(fmt.Sprintf(
        "%s:%d: variable '%s' must have either a type annotation or an initializer",
        p.l.GetFilename(), p.curToken.Line, name))
    p.skipToNextStatement()
    return nil
}
```

### 9.2 Evaluator-level enforcement

The evaluator already handles this via the `TypedLetStatement` vs `LetStatement` AST nodes:
- `TypedLetStatement` with no value → zero value
- `LetStatement` → inference from value
- Neither → parser error above (never reaches evaluator)

### 9.3 Function parameter type enforcement

In `parseFunctionParameters()`:

```go
if param.Type == nil {
    p.addError("parameter '%s' requires a type annotation in strict mode", name)
    return nil
}
```

### 9.4 Struct field type enforcement

In `parseStructStatement()`:

```go
if field.Type == nil && field.Default == nil {
    p.addError("struct field '%s' requires a type annotation in strict mode", name)
    return nil
}
```

---

## 10. Phase 6: Tests + Update Examples

### 10.1 Update typed example files

All 22 files in `examples/typed/` need revision to match locked decisions:

| Change needed | Files affected |
|---------------|---------------|
| Remove `mut` keyword | `01_basics.vint`, all files that used `mut` |
| Replace `:=` with `let x: T = ...` or `let x = ...` | `07_functions.vint`, `11_errors.vint`, `99_showcase.vint` |
| Replace `?T` with `*T` or nil checks | `13_optionals.vint` → remove entirely or rewrite without ?T |
| Replace `!T` with `(T, error)` | `14_error_union.vint` → remove or rewrite with multi-return |
| Rewrite to strict typing — every var needs a type | All files (many `let x = ...` need `: type`) |
| Remove backtick string interpolation | `03_strings.vint` → use + concatenation |

### 10.2 Add test files

**Parser tests** — add to `parser/parser_test.go` or `parser/type_test.go`:

| Test | What it validates |
|------|-------------------|
| `TestTypedLetStatement` | `let x: int = 5` parses as `TypedLetStatement` with `BasicType{Name: "int"}` and `Value: 5` |
| `TestTypedLetZeroValue` | `let x: int` parses as `TypedLetStatement` with nil Value |
| `TestUntypedLet` | `let x = 5` parses as `LetStatement` (inference) |
| `TestMissingTypeError` | `let x` produces parser error |
| `TestTypedFunctionLiteral` | `func add(a: int, b: int): int { return a + b }` produces `TypedFunctionLiteral` |
| `TestMultiReturnType` | `func divide(a, b: int): (int, error)` produces `MultiReturnType` |
| `TestTypedStructField` | `struct Point { x: int; y: int }` has `Type` on each field |
| `TestTypeAlias` | `type UserID = int` produces `TypeAliasStatement` |
| `TestAsExpression` | `x as int` produces `TypeCastExpression` |
| `TestIsExpression` | `x is int` produces `TypeCheckExpression` |
| `TestArrayType` | `[]int` parses as `ArrayType{ElementType: BasicType{Name: "int"}}` |
| `TestFixedArrayType` | `[4]byte` parses as `FixedArrayType{Size: 4, ElementType: BasicType{Name: "byte"}}` |
| `TestDictType` | `{string: int}` parses as `DictType{KeyType: ..., ValueType: ...}` |
| `TestPointerType` | `*int` parses as `PointerType{BaseType: BasicType{Name: "int"}}` |
| `TestChannelType` | `chan string` parses as `ChannelType{ElementType: BasicType{Name: "string"}}` |
| `TestFunctionType` | `func(int) bool` parses as `FunctionType` |

**Evaluator tests** — add to `evaluator/evaluator_test.go`:

| Test | What it validates |
|------|-------------------|
| `TestTypedLetCorrect` | `let x: int = 5` evaluates to 5, no error |
| `TestTypedLetMismatch` | `let x: int = "hello"` produces type error |
| `TestTypedLetZero` | `let x: int` evaluates to 0 |
| `TestTypedFunctionCorrect` | `func add(a: int, b: int): int { return a + b }` works |
| `TestTypedFunctionMismatch` | Calling with wrong type produces error |
| `TestTypedReturnCorrect` | Return value matches declared type |
| `TestTypedReturnMismatch` | Return value does not match declared type |
| `TestTypeCast` | `(5 as string)` → `"5"` |
| `TestTypeCheck` | `(5 is int)` → `true`, `(5 is string)` → `false` |
| `TestAssignmentTyped` | `let x: int = 5; x = 10` works, `x = "hi"` errors |
| `TestMultiReturn` | `let result, err := divide(10, 2)` destructures correctly |
| `TestTypeAlias` | `type ID = int; let x: ID = 5` works |
| `TestStructTypedField` | Struct with typed fields enforces on construction and mutation |
| `TestStrictModeParseError` | `let x` produces parser error |

---

## 11. Deferred to v2

| Feature | Reason |
|---------|--------|
| `:=` walrus operator | Nice-to-have shorthand, not needed for strict typing |
| `?T` optional types | `?` was ILLEGAL; needs full lex/parse/eval chain + `?.` chaining |
| `!T` error union | We chose `(T, error)` multi-return instead |
| `mut` keyword | Chose to keep `let` mutable |
| Variadic `...T` | ELLIPSIS exists; just needs parser wires |
| Named returns | Sugar on multi-return |
| Typed enum backing | `enum Status: int` — minor extension |
| Backtick strings | Needs string interpolation engine |
| Hex/binary/octal numbers | Extension to `readDecimal()` |
| Static type checker | Pre-evaluation pass; major work. Needs runtime checks working first. |
| Generics | Major redesign |
| Interfaces / traits | Major redesign |
| LSP | Depends on type checker being complete |

---

## 12. Type Compatibility Table

### Runtime type name → VintLang type mapping

The `object.VintObjectType` constants map to VintLang static type names:

| Runtime `Type()` string | VintLang static type | Notes |
|------------------------|---------------------|-------|
| `"INTEGER"` | `int`, `int8`-`int64`, `uint`-`uint64`, `byte` | Differentiated by range check |
| `"FLOAT"` | `float32`, `float64` | |
| `"STRING"` | `string` | |
| `"BOOLEAN"` | `bool` | |
| `"NULL"` | `nil`, `?T` (v2) | |
| `"BYTE"` | `byte` | |
| `"ARRAY"` | `[]T`, `[n]T` | |
| `"DICT"` | `{K: V}` | |
| `"FUNCTION"` | `func(P) R` | |
| `"BUILTIN"` | `func(P) R` | |
| `"STRUCT_INSTANCE"` | user-defined struct name | Nominal matching |
| `"ENUM"` | user-defined enum name | Nominal matching |
| `"ERROR"` | `error` | |
| `"CUSTOM_ERROR"` | `error` | |
| `"POINTER"` | `*T` | |
| `"CHANNEL"` | `chan T` | |
| `"MODULE"`, `"PACKAGE"`, `"INSTANCE"` | (special, not user-typed) | |
| `"HTTP_APP"`, etc. | (module-specific) | |

### Numeric type conversion matrix

| From ↓ / To → | `int` | `int8` | `int16` | `int32` | `int64` | `uint` | `uint8` | `uint16` | `uint32` | `uint64` | `float32` | `float64` |
|--------------|-------|--------|---------|---------|---------|--------|---------|----------|----------|----------|-----------|-----------|
| `int` | ✓ | ✓(1) | ✓(1) | ✓(1) | ✓(1) | ✓(1) | ✓(1,2) | ✓(1) | ✓(1) | ✓(1) | ✓ | ✓ |
| `int8` | ✓ | ✓ | ✓ | ✓ | ✓ | ✓(2) | ✓(2) | ✓ | ✓ | ✓ | ✓ | ✓ |
| `int16` | ✓ | — | ✓ | ✓ | ✓ | ✓(2) | — | ✓(2) | ✓ | ✓ | ✓ | ✓ |
| `int32` | ✓ | — | — | ✓ | ✓ | ✓(2) | — | — | ✓(2) | ✓ | ✓ | ✓ |
| `int64` | ✓ | — | — | — | ✓ | ✓(2) | — | — | — | ✓(2) | ✓ | ✓ |
| `uint` | ✓(2) | — | — | — | — | ✓ | — | ✓ | ✓ | ✓ | ✓ | ✓ |
| `uint8` | ✓ | ✓ | ✓ | ✓ | ✓ | ✓ | ✓ | ✓ | ✓ | ✓ | ✓ | ✓ |
| `uint16` | ✓ | — | ✓(2) | ✓ | ✓ | ✓ | — | ✓ | ✓ | ✓ | ✓ | ✓ |
| `uint32` | ✓ | — | — | ✓(2) | ✓ | ✓ | — | — | ✓ | ✓ | ✓ | ✓ |
| `uint64` | ✓ | — | — | — | ✓(2) | ✓ | — | — | — | ✓ | ✓ | ✓ |
| `float32` | ✓(3) | — | — | — | — | — | — | — | — | — | ✓ | ✓ |
| `float64` | ✓(3) | — | — | — | — | — | — | — | — | — | ✓ | ✓ |

**Legend:**
- `✓` = always compatible (implicit conversion or widening)
- `✓(1)` = compatible if value fits in target range (runtime check)
- `✓(2)` = compatible if value is non-negative + fits (runtime check)
- `✓(3)` = truncates decimal part at runtime
- `—` = incompatible (requires explicit `as` cast)

---

## 13. Error Messages Reference

### Parser errors

```
file.vint:1: variable 'x' must have either a type annotation or an initializer
file.vint:3: parameter 'a' requires a type annotation in strict mode
file.vint:5: struct field 'name' requires a type annotation or a default value
file.vint:10: unknown type 'Floatz' — did you mean 'float32' or 'float64'?
file.vint:12: 'type' keyword must be followed by 'identifier = type' — got 'if'
```

### Evaluator errors

```
TypeError: expected 'int' but got 'string'
  at file.vint:42

TypeError: cannot assign 'string' to variable 'count' of type 'int'
  at file.vint:15

TypeError: parameter 'b' expects 'int', got 'string'
  at file.vint:8

TypeError: function 'add' returns 'int', but body returned 'string'
  at file.vint:12

TypeError: field 'name' in struct 'User' expects 'string', got 'int'
  at file.vint:25

TypeError: cannot cast 'string' to 'int'
  at file.vint:30

TypeError: value overflow: 1000 does not fit in int8
  at file.vint:33

TypeError: cannot use nil for non-optional type 'int'
  at file.vint:40
```
