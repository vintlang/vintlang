# VintLang → Go Transpiler (`vint build`)

**Status:** Active design validated against source  
**Date:** July 2026  
**Author:** Tachera W

---

## 1. Vision what we are building

`vint build main.vint` produces a **single self-contained Go binary**:

```
vint build main.vint   ──►  output.go  ──►  go build  ──►  ./app
```

The transpiler converts VintLang source into Go source, then compiles it with `go build`.

**The goals (exactly):**

1. **A `vint build` command** next to the existing `bundle` command.
2. **Self-contained output.** Every `.vint` file the program depends on `import`ed files, `include`d files, `package`s is **transpiled and inlined into `output.go`**. The final binary reads **zero `.vint` files at runtime**. Nothing needs to be shipped beside the binary.
3. **Same pipeline as the bundler, different codegen.** The bundler already collects all dependencies and builds a binary we reuse its dependency analyzer and its `go build` orchestration. The only new thing is the code generator: instead of *embedding Vint source as a string*, we *emit Go statements*.
4. **Everything already available in Go stays as-is.** All 44+ modules and 57 builtins are plain Go functions reused unchanged via `module.Mapper` and `builtins.BuiltinRegistry`.

**What this honestly is (and isn't):**

- The output binary still contains the *runtime function machinery* (function bodies execute through Vint's `object.Function`/evaluator). This matches what the bundler's binary does today.
- The real wins over the bundler: no embedded source string, no lexing/parsing at startup, a smaller binary, and most importantly the foundation for a later phase that compiles function bodies and typed locals to **native Go** (Section 9).

**Why reuse the bundler instead of starting fresh:**

| Step | Bundler (`vint bundle`) | Transpiler (`vint build`) |
|------|------------------------|---------------------------|
| Collect all transitive `.vint` deps | `bundler.DependencyAnalyzer` | **reused as-is** |
| Resolve search paths / skip builtin modules / dedupe | analyzer internals | **reused as-is** |
| Emit Vint code into the binary | embeds source string + `repl.Read(...)` | **NEW transpiles each file to Go** |
| Temp dir + `go.mod` + `go build -trimpath` + move binary | `bundler.Bundle(...)` | **reused** (extracted helper) |

---

## 2. Pipeline

```
vint build main.vint
    │
    ▼
lexer ──► parser ──► *ast.Program            (existing)
    │
    ▼
bundler.DependencyAnalyzer.AnalyzeDependencies(main.vint)
    │   ──► *bundler.FileBundle {
    │         MainFile     string
    │         Files        map[absPath]content   ← every transitive .vint file
    │         IncludeFiles map[absPath]bool      ← which ones are `include`d
    │       }
    ▼
transpiler.GenerateTranspiledCode(bundle, version, buildTime)
    │   ──► output.go      (files inlined per Section 4/5)
    ▼
bundler.BuildMain(output.go, ...)     ──►  go build -trimpath  ──►  ./app
```

Only `transpiler/` (and two small upstream shims) are new.

---

## 3. How imports / includes become self-contained

Every Vint-side dependency is resolved **at build time** by `DependencyAnalyzer` and emitted into `output.go`. The final binary never touches a `.vint` file.

| Vint construct | What the bundler does | What `output.go` contains |
|---|---|---|
| `import http` (built-in module) | keeps the line; runtime sees `module.Mapper["http"]` | `env.Define("http", module.Mapper["http"])` |
| `import helper` (a `.vint` file) | wraps file body in `package helper { }` | a `*object.Package` built inline from `helper.vint`'s transpiled statements (`helperPkg := ...`), then `env.Define("helper", helperPkg)` |
| `include "greetings.vint"` | embeds file body verbatim into the main body | `greetings.vint`'s statements emitted inline, in the same scope (matches `include` semantics see docs/include.md) |
| `package p { ... }` | runtime `evalPackage` creates a scoped object | generated `*object.Package` block + member definitions + private-name tracking + auto-`init` |
| circular imports | `processed` map dedupes during analysis | impossible at runtime the graph is already flattened at build time |

---

## 4. Value mapping (Vint expression → generated Go)

Every Vint value is an `object.VintObject`. The transpiler generates Go that constructs or reuses these values:

| Vint | Generated Go |
|---|---|
| `5` | `&object.Integer{Value: 5}` |
| `3.14` | `&object.Float{Value: 3.14}` |
| `"hi"` | `&object.String{Value: "hi"}` |
| `true` / `false` | `evaluator.TRUE` / `evaluator.FALSE` |
| `null` | `evaluator.NULL` |
| `[1, 2]` | `&object.Array{Elements: []atomic.VintObject{&object.Integer{Value: 1}, ...}}` |
| `x + y` | `evaluator.ApplyInfixExpression("+", {x}, {y}, line)` |
| `x as int` | `evaluator.ApplyTypeCast({x}, &ast.BasicType{Name:"int"}, line)` |
| `::println(v)` | `evaluator.ApplyFunction(builtins.BuiltinRegistry["println"], []object.VintObject{{v}}, line)` |
| `f(args)` | `evaluator.ApplyFunction({f}, []object.VintObject{{args}}, line)` |
| `obj.method(args)` | `evaluator.ApplyMethod({obj}, &ast.Identifier{Value:"method"}, {args}, nil, line)` |
| `let x: int = 5` | `env.DefineTyped("x", &object.Integer{Value: 5}, &ast.BasicType{Name:"int"})` |
| `let x = 5` | `env.Define("x", &object.Integer{Value: 5})` |
| `let x: int` | `env.DefineTyped("x", evaluator.ZeroValueForType(&ast.BasicType{Name:"int"}), ...)` |
| `const X = 7` | `env.DefineConst("X", &object.Integer{Value: 7})` |
| `func(a: int): int { ... }` | an `*object.Function{ ParamTypes, ReturnType, Body: <rebuilt AST> , Env: env}` |
| `x = 10` | `env.Assign("x", &object.Integer{Value: 10})` |

**Key principle:** complex operations dispatch into the runtime evaluator's (exported) helper functions, so *all* semantics overloading, keyword args, type checks, method dispatch, string/array coercion, error formatting stay identical to `vint run` by construction. Only structural statements (let/const/if/while/for/import/include…) are emitted as real Go.

---

## 5. Worked example

### 5.1 The Vint project

**helper.v**  (a file you `import` as a package)

```v
// a file import: `import helper` looks for helper.v (or helper.VINT / helper.Vint)
package helper {
    let version = "1.0"

    let shout = func(msg: string): string {
        return msg + "!"
    }
}
```

**main.v**

```v
import helper

let greet: string = "hello, vint!"
let count: int = 3

let add = func(a: int, b: int): int {
    return a + b
}

::println(greet)
::println(add(count, 4))
::println(helper.shout("transpiled"))

let items = [1, 2, 3]
::println(items.len())
```

### 5.2 Expected output

```
hello, vint!
7
transpiled!
3
```

### 5.3 Generated `output.go` (annotated)

> The comments show which Vint source a block came from. Error propagation is omitted for readability; the emitted pattern is shown in Section 6.

```go
package main

import (
    "flag"
    "fmt"

    "github.com/vintlang/vintlang/internal/ast"
    "github.com/vintlang/vintlang/internal/evaluator"
    "github.com/vintlang/vintlang/internal/evaluator/builtins"
    "github.com/vintlang/vintlang/internal/module"
    "github.com/vintlang/vintlang/internal/object"
    "github.com/vintlang/vintlang/internal/toolkit"
)

var TranspilerVersion = "v1.x.x"
var BuildTime        = "2026-07-01T00:00:00Z"

func main() {
    details := flag.Bool("i", false, "Show build details")
    flag.Parse()
    if *details {
        fmt.Printf("[Vint Transpiler v%s | built %s]\n", TranspilerVersion, BuildTime)
        return
    }
    toolkit.CLI_ARGS = flag.Args()

    // ═══ runtime environment (same as `vint run` creates) ═══
    env := object.NewEnvironment()

    // ═══ import helper  (helper.v embedded here, nothing read at runtime) ═══
    helperPkg := &object.Package{
        Name:         &ast.Identifier{Value: "helper"},
        Env:          env,
        Scope:        object.NewEnclosedEnvironment(env),
        PrivateNames: map[string]bool{},
    }
    helperPkg.Scope.Define("@", helperPkg)            // package self-reference

    //   package helper { let version = "1.0" → helperPkg.Scope
    helperPkg.Scope.Define("version", &object.String{Value: "1.0"})
    //   let shout = func(msg: string): string { return msg + "!" }
    shoutFn := &object.Function{
        Name:       "shout",
        Parameters: []*ast.Identifier{{Value: "msg"}},
        ParamTypes: []ast.Type{&ast.BasicType{Name: "string"}},
        ReturnType: &ast.BasicType{Name: "string"},
        // body `return msg + "!"` is rebuilt as AST so the runtime can run it
        Env:        helperPkg.Scope,
    }
    helperPkg.Scope.Define("shout", shoutFn)
    env.Define("helper", helperPkg)                   // ← import binding

    // ═══ let greet: string = "hello, vint!" ═══
    env.DefineTyped("greet", &object.String{Value: "hello, vint!"}, &ast.BasicType{Name: "string"})

    // ═══ let count: int = 3 ═══
    env.DefineTyped("count", &object.Integer{Value: 3}, &ast.BasicType{Name: "int"})

    // ═══ let add = func(a: int, b: int): int { return a + b } ═══
    addFn := &object.Function{
        Parameters: []*ast.Identifier{{Value: "a"}, {Value: "b"}},
        ParamTypes: []*ast.BasicType{Name: "int"}, [] ...  // one per param
        ReturnType: &ast.BasicType{Name: "int"},
        Body:       /* rebuilt AST of { return a + b } */,
        Env:        env,
    }
    env.Define("add", addFn)

    // ═══ ::println(greet) ═══
    {
        greetVal, _ := env.Get("greet")
        evaluator.ApplyFunction(builtins.BuiltinRegistry["println"],
            []object.VintObject{greetVal}, 14)
    }

    // ═══ ::println(add(count, 4)) ═══
    {
        countVal, _ := env.Get("count")
        addVal, _   := env.Get("add")
        sum := evaluator.ApplyFunction(addVal,
            []object.VintObject{countVal, &object.Integer{Value: 4}}, 18)
        evaluator.ApplyFunction(builtins.BuiltinRegistry["println"],
            []object.VintObject{sum}, 18)
    }

    // ═══ ::println(helper.shout("transpiled")) ═══
    helperVal, _ := env.Get("helper")
    shouted := evaluator.ApplyMethod(helperVal,
        &ast.Identifier{Value: "shout"},
        []object.VintObject{&object.String{Value: "transpiled"}},
        nil,
        21)
    evaluator.ApplyFunction(builtins.BuiltinRegistry["println"],
        []object.VintObject{shouted}, 21)

    // ═══ let items = [1, 2, 3] ═══
    env.Define("items", &object.Array{Elements: []object.VintObject{
        &object.Integer{Value: 1},
        &object.Integer{Value: 2},
        &object.Integer{Value: 3},
    }})

    // ═══ ::println(items.len()) ═══
    itemsVal, _ := env.Get("items")
    length := evaluator.ApplyMethod(itemsVal,
        &ast.Identifier{Value: "len"}, []object.VintObject{}, nil, 24)
    evaluator.ApplyFunction(builtins.BuiltinRegistry["println"],
        []object.VintObject{length}, 24)
}
```

That is the whole binary. No embedded Vint text, no `repl.Read`, no file-lookup loop just one `main()` that builds runtime objects and calls the runtime's helpers.

---

## 6. Execution model & error handling

Generated code mirrors what the interpreter does, so behavior matches `vint run`:

1. `env := object.NewEnvironment()` same as `repl.ReadWithFilename` (repl.go:31).
2. Builtins (`::name`) resolve from `builtins.BuiltinRegistry` via `evaluator.ApplyFunction(...)` never from env (that's the point of `::`).
3. Modules (`import http`) resolve from `module.Mapper`; file imports and includes are the inlined package blocks from Section 5.
4. **Errors are values.** Every runtime call can return `*object.Error`. Generated code checks it like the evaluator does:

```go
res := evaluator.ApplyFunction(fn, args, 9)
if res != nil && res.Type() == object.ERROR_OBJ {
    fmt.Fprintln(os.Stderr, res.Inspect())
    return
}
```

Line numbers are threaded through every emitted call so runtime errors keep their `line N` formatting.

---

## 7. New Go code needed

### 7.1 `internal/transpiler/` (new package)

| File | Responsibility |
|------|----------------|
| `transpiler.go` | `Transpile(mainFile string) (string, error)` run analyzer, walk bundle (main file last, deps first), assemble `output.go` |
| `main_template.go` | the `package main` + `main()` skeleton (Section 5.3 / 6) |
| `statements.go` | `let`/`const`/typed-`let`/assign/`return`/blocks/`if`/`while`/`repeat`/`for..in`/`switch`/`match`/`break`/`continue` |
| `expressions.go` | identifiers, literals, calls, `::builtin`, index/slice/property/method, type cast/check, ranges |
| `functions.go` | `object.Function` construction + rebuilt-body emission; later: native bodies |
| `modules.go` | import/include/package emission (Section 3) |
| `helpers.go` | indent, string escaping, `isError` check helper, line numbers, import-name dedup |
| `transpiler_test.go` | parity tests (Section 8) |

### 7.2 Upstream changes (small, backward-compatible)

- **`internal/evaluator`**: export thin wrappers for the helpers generated code calls today they are unexported:
  - `func ApplyFunction(fn object.VintObject, args []object.VintObject, line int) object.VintObject` (wraps `applyFunction`)
  - `func ApplyMethod(obj object.VintObject, m ast.Expression, args []object.VintObject, defs map[string]object.VintObject, l int) object.VintObject` (wraps `applyMethod`)
  - `func ApplyInfixExpression(op string, l, r object.VintObject, line int) object.VintObject` (wraps `evalInfixExpression`)
  - `func ZeroValue(t ast.Type) object.VintObject` (wraps `zeroValueFromType`)
- **`internal/bundler/bundler.go`** extract the temp-dir / `go.mod` / `go build` / move-binary steps into a shared `func BuildMain(goCode string, opts ...) error`; `bundle` calls it too.
- **`main.go`** add:

```go
case "build", "-build", "--build":
    // vint build file.vint [-o name] [GOOS] [GOARCH] [quiet|keep...]  same shape as bundle
    if err := transpiler.Build(args[2:]); err != nil {
        fmt.Println(err)
        os.Exit(1)
    }
```

---

## 8. Milestones / rollout

| # | Milestone | Verifiable by |
|---|-----------|---------------|
| 1 | **Single-file transpilation** `main()` skeleton, literals, `let`/`const`/typed-let, identifiers, `::builtin`, `vint build` CLI | `examples/typed/01_basics.vint` builds & output matches `vint run` byte-for-byte |
| 2 | **Imports, includes, packages** (Section 3) the self-containment core | multi-file project builds; delete the `.vint` files and the binary still runs |
| 3 | **Control flow** if/while/repeat/for-in/switch/match/break/continue/return | `examples/for_loops.vint`, `if_expression.vint` parity |
| 4 | **Functions & the rest** function literals, closures, structs, enums, errors, `go`/`chan`/`await`/`defer` | `examples/typed/08_…`, `concurrency.vint`, async demos parity |
| 5 | **Type-specialized native codegen** (optional, incremental) annotated ints/strings compile to `int64`/`string` locals, calls dispatch natively | parity retained + benchmark: typed loops faster than interpreter |

---

## 9. Key design decisions

1. **Reuse the runtime, don't reimplement it.** Leaf operators (`+`, `==`, casts, method dispatch, builtin calls) dispatch into `evaluator`'s exported helpers. This guarantees parity by construction including error text. Native codegen (Milestone 4+) upgrades specific sites later.
2. **Self-contained by inlining, not embedding.** Imports/includes become transpiled Go in `output.go`. The binary ships alone.
3. **Environment at runtime, not Go scopes.** Functions carry `*object.Environment`s (Vint closure/overloading semantics live there environment.go). Go closures can't express Vint's env-chain lookup.
4. **One flat `.go` file.** Simple, debuggable, inspectable. Splitting later is additive.
5. **Errors are values.** Every statement follows the interpreter's `isError` pattern.

---

## 9. Risks + mitigations

| Risk | Impact | Mitigation |
|------|--------|------------|
| Semantic drift vs `vint run` | Wrong behavior | Parity harness per milestone (Section 8); keep dispatch on runtime helpers |
| `go` statements transferring a shared `*Environment` | Data race (existing in interpreter too) | guard shared env with mutex; document in generated comment |
| Vint function bodies rebuilt as AST at emit-time bloat the file | Correct but large `output.go` / slower emit | Phase-10 native bodies replace this per-function |
| Rebuilt-AST bodies must match parser exactly | Trust drift | `func emitASTLiteral` is a direct rewrite of the parsed node snapshot-tested |
| `go build` fails in user envs | Broken `vint build` | Reuse bundler's exact temp-dir/`go.mod`/`-mod=mod` recipe; `-S` flag to output `.go` only |

---

## 10. Future work

- **Native function bodies + typed locals** Phase-10 style: `let x: int` → `var x int64`, hot loops in native Go.
- **`-S` mode** emit `output.go` without invoking `go build` (debugging).
- **Cross-compile** `vint build` already reuses `GOOS`/`GOARCH` from the bundler.
- **WASM target** transpile + `GOOS=js GOARCH=wasm go build`.
- **Plugin / shared libraries** `vint build -buildmode=c-shared` for functions exposure.
- **Source maps** line-number mapping `.vim` → `.go` for runtime stack traces.

---

## 11. File map (measured against source)

| File | Est. lines |
|------|-----------|
| `transpiler/transpiler.go` | 120 |
| `transpiler/main_template.go` | 90 |
| `transpiler/statements.go` | 320 |
| `transpiler/expressions.go` | 260 |
| `transpiler/functions.go` | 160 |
| `transpiler/modules.go` | 130 |
| `transpiler/helpers.go` | 80 |
| `transpiler/transpiler_test.go` | 400+ |
| `bundler/bundler.go` (extract `BuildMain`) | ±60 |
| `evaluator/evaluator.go` (4 export wrappers) | ±20 |
| `main.go` (build command) | ±10 |

**Total to write: ~1400–1700 lines + tests**, of which most is mechanical statement emission the two real pieces of new logic are the import/include/package inliner (modules.go) and the AST-rebuild emitter for function bodies.