# VintLang Examples

Welcome to the VintLang examples directory! This collection demonstrates the features and capabilities of the VintLang programming language.

## 🚀 Running Examples

```bash
vint examples/<filename>.vint
```

---

## 📚 Examples by Category

### Core Language

| File | Description |
|------|-------------|
| `builtins.vint` | Built-in functions: type checking, conversion, collections, math |
| `functions.vint` | Function definitions, IIFE, higher-order functions |
| `closures.vint` | Closures, function factories, map/filter/reduce, memoization |
| `overloading.vint` | Function overloading by arity (number of arguments) |
| `if_expression.vint` | if as statement and as expression |
| `switch.vint` | Switch-case control flow |
| `for_loops.vint` | for, while, repeat loops; break and continue |
| `repeat-keyword.vint` | The repeat loop |
| `logicals.vint` | Logical operators: and, or, not |
| `pointers.vint` | Pointer operations with & and * |
| `defer.vint` | Defer statement — deferred execution (LIFO) |

### Data Types & Structures

| File | Description |
|------|-------------|
| `arrays.vint` | Array creation, slicing, sorting, searching, iteration |
| `dictionaries.vint` | Dict creation, access, iteration, has_key, nested dicts |
| `strings.vint` | String module: trim, split, join, replace, etc. |
| `nativeStrings.vint` | Native string methods: .upper(), .split(), .replace(), etc. |
| `json.vint` | JSON encode, decode, pretty-print, merge |
| `enum_demo.vint` | Enum types with integer and string values |
| `structs.vint` | Struct definitions with fields and methods |

### Pattern Matching

| File | Description |
|------|-------------|
| `pattern_matching.vint` | switch guards, match on dicts, type-based matching |
| `user_access_control.vint` | Dict pattern matching for access control logic |

### Modules & Standard Library

| File | Description |
|------|-------------|
| `math_extensions.vint` | Statistics, complex numbers, linear algebra, GCD/LCM |
| `random.vint` | Random integers, floats, strings, choices |
| `time.vint` | Current time, formatting, arithmetic, sleep |
| `encoding.vint` | Base64 encode and decode |
| `uuid.vint` | UUID generation |
| `regex.vint` | Pattern matching, text replacement, splitting |
| `colors.vint` | RGB to hex color conversion |
| `reflect.vint` | Runtime type inspection |
| `fmt_demo.vint` | String formatting, padding, alignment, tables |

### Concurrency

| File | Description |
|------|-------------|
| `concurrency.vint` | async/await, goroutines, channels, worker pool |

### System & File Operations

| File | Description |
|------|-------------|
| `os.vint` | File I/O, directory listing, environment variables |
| `path.vint` | File path manipulation: join, basename, dirname, ext |
| `shell.vint` | Run shell commands from VintLang |
| `sysinfo.vint` | OS, CPU, memory, disk, network information |
| `make_example.vint` | Build automation with the make module |
| `dotenv.vint` | Load environment variables from a .env file |

### Error Handling & Debugging

| File | Description |
|------|-------------|
| `error_handling.vint` | Null checks, error returns, validation, retry pattern |

### Data Formats

| File | Description |
|------|-------------|
| `csv.vint` | Write and read CSV files |
| `excel_demo.vint` | Create Excel workbooks with sheets and formulas |

### Security & Auth

| File | Description |
|------|-------------|
| `crypto.vint` | MD5 hashing, AES encryption/decryption |
| `jwt.vint` | Create, decode, and verify JWT tokens |

### Databases

| File | Description |
|------|-------------|
| `sqlite.vint` | SQLite: create, insert, fetch, update |
| `mysql.vint` | MySQL database operations *(requires MySQL server)* |
| `postgres.vint` | PostgreSQL database operations *(requires PostgreSQL server)* |
| `redis.vint` | Redis key-value operations *(requires Redis server)* |

### Networking & HTTP

| File | Description |
|------|-------------|
| `http.vint` | HTTP file server setup |
| `backend_demo.vint` | HTTP backend with routes, middleware, interceptors |
| `complete_backend_app.vint` | Full REST API backend definition |
| `express_like_server.vint` | Express-like HTTP server with routes |
| `web_fetcher.vint` | HTTP GET requests and JSON parsing *(requires network)* |
| `github-profile.vint` | Profile view counter via HTTP *(requires network)* |
| `vintSocket.vint` | WebSocket server and client |

### Visualization

| File | Description |
|------|-------------|
| `vintChart.vint` | Bar chart, pie chart, line graph generation |

### Modules & Packages

| File | Description |
|------|-------------|
| `greetings_module.vint` | Package definition with functions |

### CLI Applications

| File | Description |
|------|-------------|
| `cli-todo.vint` | Todo list CLI with add/list commands |
| `cli.vint` | CLI argument parsing with the cli module |

### Terminal UI

| File | Description |
|------|-------------|
| `term.vint` | Terminal banners, menus, and prompts |

### AI / LLM

| File | Description |
|------|-------------|
| `llm_openai.vint` | OpenAI GPT integration *(requires API key)* |

### Interactive Games

| File | Description |
|------|-------------|
| `guessingGame.vint` | Number guessing game *(interactive)* |
| `simplegame.vint` | Text adventure with save/load *(interactive)* |

---

## 📝 Notes

### External Dependencies

Some examples require external resources and will not run without them:
- **Database examples** (`mysql`, `postgres`, `redis`): Need running servers
- **dotenv**: Needs a `.env` file
- **llm_openai**: Requires `OPENAI_API_KEY` environment variable
- **web_fetcher**, **github-profile**: Require internet connectivity
- **Interactive examples** (`guessingGame`, `simplegame`, `cli-todo`, `cli`): Need user input

### All other examples run without external dependencies.

---

## 🎯 Learning Path

1. **Start here**: `builtins.vint`, `functions.vint`, `if_expression.vint`
2. **Control flow**: `for_loops.vint`, `switch.vint`, `repeat-keyword.vint`
3. **Data structures**: `arrays.vint`, `dictionaries.vint`, `strings.vint`
4. **Functions**: `closures.vint`, `overloading.vint`
5. **Pattern matching**: `pattern_matching.vint`
6. **OOP-style**: `structs.vint`, `enum_demo.vint`
7. **Modules**: `math_extensions.vint`, `time.vint`, `regex.vint`
8. **System**: `os.vint`, `shell.vint`, `sysinfo.vint`
9. **Concurrency**: `concurrency.vint`
10. **Real-world**: `sqlite.vint`, `http.vint`, `json.vint`
