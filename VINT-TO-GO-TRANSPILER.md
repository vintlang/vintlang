# VintLang → Go Transpiler: Vision & Implementation Plan

**Status:** Planning phase  
**Date:** July 2026  
**Author:** Tachera W

---

## 1. Vision

VintLang is a Go-embedded scripting language. Every parser node, every runtime value, every builtin function, and every one of its 44+ modules (http, json, os, time, crypto, etc.) is already Go code. A transpiler converts VintLang source to Go source, then compiles it with `go build` to produce a native binary.

**Why this is uniquely practical:**

| Concern | Reality |
|---------|---------|
| Standard library | 44+ modules already written in Go — nothing to port |
| Runtime | The `object/` package IS the runtime — reuse directly |
| Concurrency | Go's goroutines + channels map 1:1 to Vint's `go`/`chan` |
| Garbage collection | Go's GC handles everything |
| Error model | `(T, error)` multi-return is native Go |
| Cross-compilation | `GOOS=linux GOARCH=arm64 go build output.go` |
| Performance | Type-annotated code can emit native `int64`, `string`, etc. |

**What it achieves:**

- Native binary performance (no interpretation overhead)
- No Go runtime dependency in the output (unlike the bundler approach)
- Full native binary via `go build`
- Preserves all existing VintLang libraries unchanged

---

## 2. Pipeline

```
source.vint
    │
    ▼
lexer.Lexer ──► token.Token stream
    │
    ▼
parser.Parser ──► *ast.Program (full AST with types)
    │
    ▼
transpiler.Transpile(program) ──► output.go
    │
    ▼
go build output.go ──► native binary
```

The lexer, parser, and type system already exist (PR #277). Only the `transpiler/` package is new.

---

## 3. Architecture

### 3.1 The `transpiler/` package

```
transpiler/
  transpiler.go       Entry point: *ast.Program → Go source string
  statements.go       let, const, return, block, if, for, while, repeat, forin
  expressions.go      infix, prefix, call, index, property, slice, method
  literals.go         int, float, string, bool, null, array, dict, range
  functions.go        function literals (typed & untyped), closures
  types.go            type annotations → Go zero values + cast/check
  structs.go          struct definitions, instantiation, method calls
  modules.go          imports, includes, packages
  helpers.go          indent tracking, import dedup, string escaping
  main_template.go    The `main()` wrapper and variable registration
```

### 3.2 Execution model

The transpiler generates a Go program that:

1. Creates an `object.NewEnvironment()` at startup
2. Defines all builtins from `builtin.BuiltinRegistry`
3. Defines all modules from `module.Mapper`
4. Executes top-level statements sequentially
5. If a `main` function exists, calls it at the end
6. Returns the final result (or prints errors)

The generated Go file imports two packages:
- `"github.com/vintlang/vintlang/object"` — all runtime types
- `"github.com/vintlang/vintlang/evaluator/builtins"` — builtin functions

For modules, it imports and registers from `"github.com/vintlang/vintlang/module"`.

### 3.3 Value mapping

Every VintLang value is an `object.VintObject`. The transpiler generates Go code that constructs these objects:

| VintLang expression | Generated Go |
|---|---|
| `5` | `&object.Integer{Value: 5}` |
| `3.14` | `&object.Float{Value: 3.14}` |
| `"hello"` | `&object.String{Value: "hello"}` |
| `true` / `false` | `evaluator.TRUE` / `evaluator.FALSE` |
| `nil` | `evaluator.NULL` |
| `[1, 2, 3]` | `&object.Array{Elements: []object.VintObject{...}}` |
| `{"a": 1}` | `&object.Dict{Pairs: map[object.HashKey]object.DictPair{...}}` |
| `x + y` | `evaluator.Eval(&ast.InfixExpression{...}, env)` |
| `fn(a, b)` | `evaluator.applyFunction(fnObj, args, line)` |
| `obj.method(args)` | `evaluator.applyMethod(obj, methodIdent, args, defs, line)` |
| `x as int` | `evaluator.evalTypeCast(&ast.TypeCastExpression{...}, env)` |
| `x is int` | `evaluator.evalTypeCheck(&ast.TypeCheckExpression{...}, env)` |
| `let x: int = 5` | `env.DefineTyped("x", &object.Integer{Value: 5}, &ast.BasicType{Name: "int"})` |
| `let x = 5` | `env.Define("x", &object.Integer{Value: 5})` |
| `func(a: int): int { ... }` | `&object.Function{ParamTypes: ..., ReturnType: ..., Body: ..., Env: env}` |
| `::println("hi")` | `builtins.BuiltinRegistry["println"].Fn(/* args */)` |

**Key principle:** The transpiler reuses existing evaluator functions where possible. Complex operations like `+`, `-`, `==`, array indexing, method dispatch, and type checking call into `evaluator.xxx()` at runtime. Simple operations like constant literals and variable references are inlined directly.

---

## 4. Implementation Plan

### Phase 1: Scaffold + Literals (Days 1-2)

**Goal:** `transpile 42 → "42"` works.

- Create `transpiler/` package with entry point
- `Transpile(program *ast.Program) (string, error)` — walks `program.Statements`, concatenates Go output
- Generate `package main` header + `func main()` wrapper
- Handle literals: `IntegerLiteral`, `FloatLiteral`, `StringLiteral`, `Boolean`, `Null`
- Handle `ExpressionStatement` (statement wrapping an expression)
- Generate `env := object.NewEnvironment()` and builtin registration
- Write test: transpile `42` to `.go`, `go build`, run, verify output

**Files:** `transpiler.go`, `literals.go`, `helpers.go`, `main_template.go`

### Phase 2: Variables + Identifiers (Day 3)

**Goal:** `let x: int = 5` transpiles correctly.

- `Identifier` → `env.Get("x")` call
- `LetStatement` (untyped) → `env.Define("x", value)`
- `TypedLetStatement` → `env.DefineTyped("x", value, typeNode)`
- `ConstStatement` → `env.DefineConst("x", value)`
- `Assign` (x = expr) → `env.Assign("x", value)`
- `AssignEqual` (x = expr in call args) → same as Assign
- Handle `BuiltinExpression` (`::println`) → `builtins.BuiltinRegistry["println"]`

**Files:** `statements.go`, `expressions.go`

### Phase 3: Control Flow (Days 4-5)

**Goal:** `if`, `for`, `while`, `return`, `block` transpile.

- `BlockStatement` → Go `{ ... }` block with scoped env
- `IfExpression` → Go `if` / `else if` / `else`
- `WhileExpression` → Go `for condition { ... }`
- `RepeatStatement` → Go `for i := 0; i < N; i++ { ... }`
- `ForIn` → Go `for key, value := range iterable { ... }` (via helper)
- `ReturnStatement` → `return &object.ReturnValue{Value: val}`
- `SwitchExpression` → Go `switch { ... }`
- `MatchExpression` → Go `switch` with case patterns
- `Break` / `Continue` → Go `break` / `continue`

**Files:** `statements.go`, `expressions.go`

### Phase 4: Functions + Closures (Days 6-7)

**Goal:** Functions with typed params/returns, recursion, closures.

- `FunctionLiteral` → `&object.Function{...}` struct literal
- `TypedFunctionLiteral` → same with `ParamTypes` + `ReturnType`
- `CallExpression` → `evaluator.applyFunction(fn, args, line)`
- Handle function name binding (recursion via `fn.Env.Define(fn.Name, fn)`)
- Handle default parameter values
- Handle overloaded functions (multiple `Define` to `funcs` map)
- `AsyncFunctionLiteral` → `&object.AsyncFunction{...}`
- `AwaitExpression` → `await` logic (promise.Wait())

**Files:** `functions.go`, `expressions.go`

### Phase 5: Types + Casts (Day 8)

**Goal:** Type annotations, zero values, `as`/`is` operators.

- Generate `ast.BasicType{Name: "int"}` for type annotations
- Zero-value declarations: `let x: int` → `zeroValueFromType(typeAnnotation.Type)`
- `TypeCastExpression` → `evaluator.evalTypeCast(node, env)`
- `TypeCheckExpression` → `evaluator.evalTypeCheck(node, env)`
- `TypeAliasStatement` → (handled during parsing, no runtime effect)
- Handle `MultiReturnType` in function return type annotations

**Files:** `types.go`

### Phase 6: Structs + Enums + Objects (Days 9-10)

**Goal:** Struct definitions, instantiation, field access, methods.

- `StructStatement` → `evaluator.evalStructStatement(node, env)`
- `StructLiteral` (User{name: "A"}) → `evaluator.evalStructLiteral(node, env)`
- `Struct` call syntax (User("A", 30)) → `evaluator.evalStructCall(node, structDef, env)`
- `PropertyExpression` (obj.field) → `evaluator.evalPropertyExpression(node, env)`
- `PropertyAssignment` (obj.field = val) → `evaluator.evalPropertyAssignment(node, val, env)`
- `MethodExpression` (obj.method(args)) → `evaluator.applyMethod(obj, method, args, defs, line)`
- `EnumStatement` → `evaluator.evalEnumStatement(node, env)`
- `ErrorDeclaration` + `ThrowStatement` → existing evaluator calls

**Files:** `structs.go`

### Phase 7: Arrays, Dicts, Indexing (Day 11)

**Goal:** All collection operations.

- `ArrayLiteral` → `&object.Array{Elements: [...]}`
- `DictLiteral` → `&object.Dict{Pairs: map[...]...}`
- `IndexExpression` (arr[i], dict["k"]) → `evaluator.evalIndexExpression(obj, index, line)`
- `SliceExpression` (arr[1:3]) → `evaluator.evalSliceExpression(obj, start, end, line)`
- `RangeExpression` (1..10) → `&object.Range{Start: 1, End: 10, Current: 1}`
- Builtin methods: `.len()`, `.map()`, `.filter()`, etc. dispatched via `applyMethod`

**Files:** `expressions.go`, `literals.go`

### Phase 8: Imports + Packages (Day 12)

**Goal:** `import "http"`, `import "file.vint"`, `package`.

- Module imports (`import "http"`) → register from `module.Mapper`
- File imports (`import "helper"`) → transpile the imported file and inline
- `Package` statement → generate package wrapper with scope
- `IncludeStatement` → read, parse, transpile included file

**Files:** `modules.go`

### Phase 9: Concurrency (Day 13)

**Goal:** `go fn()`, `chan`, `defer`.

- `GoStatement` → `go func() { Eval(expr, env) }()`
- `ChannelExpression` → `object.NewChannel()` or `object.NewBufferedChannel(n)`
- `DeferStatement` → `defer func() { ... }()` or `env.AddDefer(...)` logic
- `DebouncedFunction` → existing `object.DebouncedFunction`

**Files:** `expressions.go`, `statements.go`

### Phase 10: Optimization Pass (Day 14+)

**Goal:** Speed up by eliminating runtime dispatch where types are known.

- **Constant folding:** `1 + 2` → `&object.Integer{Value: 3}` (no runtime eval)
- **Type-specialized variables:** `let x: int = 5` → `var x int64 = 5` (skip boxing)
- **Type-specialized arithmetic:** `(x: int) + (y: int)` → native `x + y` (skip `evalInfixExpression`)
- **Go-native math:** Replace `&object.Integer{Value: 5}` with `int64(5)` in typed scopes

This phase is optional — the transpiler works correctly without it. Each optimization is independent and incremental.

---

## 5. Output Structure

The transpiler produces a single `.go` file:

```go
package main

import (
    "github.com/vintlang/vintlang/object"
    "github.com/vintlang/vintlang/evaluator"
    "github.com/vintlang/vintlang/evaluator/builtins"
    "github.com/vintlang/vintlang/module"
)

func main() {
    env := object.NewEnvironment()

    // Register builtins
    for name, fn := range builtins.BuiltinRegistry {
        env.Define(name, fn)
    }

    // Register modules
    for name, mod := range module.Mapper {
        env.Define(name, mod)
    }

    // Transpiled program
    var x object.VintObject
    x = &object.Integer{Value: 5}
    // ... rest of transpiled code ...
}
```

For modules that register into `module.Mapper` via `init()`, those are automatically available in Go — the transpiler just generates a reference.

---

## 6. Type-Specialized Mode (Phase 10)

When a variable has an explicit type annotation and all operations on it are monomorphic, the transpiler can emit native Go types instead of `VintObject`:

| VintLang | Normal transpile | Optimized transpile |
|---|---|---|
| `let x: int = 5` | `var x object.VintObject = &object.Integer{Value: 5}` | `var x int64 = 5` |
| `let s: string = "hi"` | `var s object.VintObject = &object.String{Value: "hi"}` | `var s string = "hi"` |
| `x + 1` | `evaluator.evalInfixExpression("+", x, y, 0)` | `x + int64(1)` |
| `"hello" + s` | `evaluator.evalInfixExpression("+", x, y, 0)` | `"hello" + s` |
| `5 as string` | `evaluator.evalTypeCast(...)` | `strconv.FormatInt(5, 10)` |

This is tracked as Phase 10 because it depends on complete, correct type information flowing through the AST. The PR #277 type system provides this foundation.

---

## 7. Testing Strategy

### Integration tests

For each example in `examples/typed/`:
1. Transpile `file.vint` → `file.go`
2. Run `go run file.go`
3. Compare output with `vint run file.vint`

### Snapshot tests

Transpile known VintLang snippets and compare the generated Go against expected output.

### Runtime equivalence tests

Every evaluator test in `evaluator/evaluator_test.go` has a transpiler twin: same VintLang input, same expected output, but via the transpiler pipeline.

---

## 8. CLI Integration

Add a `build` command to `main.go`:

```
vint build file.vint        → produces file.go + runs go build → binary
vint build file.vint -o app → produces app binary
vint build file.vint -S     → outputs .go source only (no go build)
```

Implementation in `main.go`:

```go
case "build", "-build", "--build":
    if len(args) < 3 {
        fmt.Println("Error: Please specify a Vint file to build")
        os.Exit(1)
    }
    transpiler.Build(args[2], outputPath)
```

---

## 9. Key Design Decisions

### Decision 1: Reuse evaluator functions vs. emit native Go

**Chosen:** Reuse evaluator functions for complex ops; inline only for trivials.

For `x + y`, emit:
```go
evaluator.evalInfixExpression("+", x, y, lineNumber)
```

This avoids reimplementing the entire evaluator. Phase 10 (optimization) can replace specific calls with native Go when types are known.

### Decision 2: One big file vs. multiple files

**Chosen:** One flat `.go` file per VintLang module.

Simple and debuggable. The user can inspect the generated `.go` file. For large projects, the transpiler can be extended to split output.

### Decision 3: Environment at runtime vs. Go scopes

**Chosen:** Use `object.Environment` at runtime.

VintLang's closure semantics (function env chains, overloading, `GetDeclaredType`) match the existing `Environment` implementation. Translating to Go scopes would lose closure semantics.

### Decision 4: Error handling

**Chosen:** Use `object.Error` return values (same as evaluator).

VintLang errors are values, not panics. Every operation checks `isError(result)` and returns immediately. This matches the evaluator's pattern exactly.

---

## 10. File Map

| File | Responsibility | Est. lines |
|------|---------------|------------|
| `transpiler/transpiler.go` | Entry, Program → Go, main wrapper | 80 |
| `transpiler/literals.go` | Integer, Float, String, Bool, Null, Array, Dict literals | 120 |
| `transpiler/statements.go` | Let, Const, TypedLet, Assign, Return, Block, If, While, For, ForIn, Repeat, Break, Continue, Switch, Match | 250 |
| `transpiler/expressions.go` | Infix, Prefix, Postfix, Call, Index, Slice, Property, Method, Identifier, Builtin, Range | 200 |
| `transpiler/functions.go` | FunctionLiteral, TypedFunctionLiteral, AsyncFunctionLiteral, closures, defaults | 120 |
| `transpiler/types.go` | Type annotations, TypeCast/Check, zero values, type aliases | 80 |
| `transpiler/structs.go` | StructStatement, StructLiteral, struct call, enum, error decl, throw | 100 |
| `transpiler/modules.go` | Import (module + file), Include, Package | 80 |
| `transpiler/helpers.go` | Indent tracker, import dedup, string escaping, line numbers | 60 |
| `transpiler/transpiler_test.go` | Tests for each node type | 300+ |

**Total: ~1400 lines**

---

## 11. Risks + Mitigations

| Risk | Impact | Mitigation |
|------|--------|------------|
| Closure mutation across goroutines | Data races | Go's race detector. Closures capture `*Environment` which is not thread-safe for mutations. Use `sync.Mutex` in generated code if needed. |
| `init()` function ordering | Wrong module init order | Modules register in `module.Mapper` via Go's `init()`. The transpiler just references `module.Mapper`, which respects Go's init ordering. |
| VintLang package member privacy | Encapsulation broken | The transpiler uses `object.Package.GetPublic()` just like the evaluator. Privacy is maintained. |
| Circular imports | Stack overflow | Detected at parse time (existing `importedModules` map). Transpiler can detect and error at transpile time. |
| Go compilation errors | Broken output | Every generated line is a known pattern. Test suite catches regressions. |

---

## 12. Future Work

- **Phase 10 optimizations** — type-specialized native Go emission
- **Separate compilation** — one `.go` file per VintLang module/package
- **LTO** — inline small functions at transpile time
- **WASM target** — transpile + `GOOS=js GOARCH=wasm go build`
- **Plugin support** — `vint build plugin.vint` → `.so` shared library
- **Source maps** — line number mapping from `.vint` to `.go` for debugging
