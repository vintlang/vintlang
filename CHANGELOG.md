# Changelog

## [September-2025] – Major Refactoring & New Features

### Major Refactoring (September 30, 2025)

- **Builtin Functions Restructuring**
  Complete refactoring of the builtin functions system from a monolithic structure to a modular, maintainable architecture:
  - Split 800+ lines `builtins.go` into 9 categorized files under `evaluator/builtins/`
  - Implemented centralized registration system with automatic function discovery
  - Moved specialized functions to appropriate modules for better organization
  - Created comprehensive developer documentation for the new structure
  
  **New Structure:**
  ```
  evaluator/builtins/
  ├── registry.go         # Central registration system
  ├── helpers.go          # Common helper functions  
  ├── core.go            # Essential functions (print, len, type, input)
  ├── io.go              # I/O functions (open, write)
  ├── type_conversion.go # Type conversion functions
  ├── logic.go           # Logic operations (and, or, not, xor, etc.)
  ├── arrays.go          # Array manipulation functions
  ├── dict.go            # Dictionary operations
  ├── system.go          # System functions (exit, sleep, args)
  ├── channels.go        # Channel operations
  └── imports.go         # Import functions
  ```
  
  **Benefits:**
  - Better maintainability with small, focused files
  - Reduced merge conflicts for multi-developer work
  - Clear guidelines for adding new functions
  - Improved code organization and navigation
  - Backward compatibility maintained - no breaking changes

### New Features (September 2025)

- **Dynamic Import Function**
  Added `import()` builtin function for runtime module importing:
  - Syntax: `import("module_name")` - dynamically loads built-in modules at runtime
  - Supports all existing built-in modules (os, math, string, etc.)
  - Enhanced lexer with lookahead capabilities for better parsing
  - Comprehensive error handling and validation
  - Works alongside traditional `import` statements
  
  **Sample:**
  ```js
  let os_module = import("os")
  let current_dir = os_module.Functions["getwd"]([])
  
  let math_module = import("math")
  let result = math_module.Functions["abs"]([-5])
  ```

- **YAML Module**
  Complete YAML processing capabilities with comprehensive functionality:
  - `yaml.decode(yamlString)` - parses YAML strings into VintLang objects
  - `yaml.encode(object)` - converts VintLang objects to YAML format
  - `yaml.merge(yaml1, yaml2)` - merges two YAML objects
  - `yaml.get(yamlObj, key)` - retrieves values from YAML objects with dot notation
  - Support for complex nested structures, arrays, and mappings
  - Proper error handling for malformed YAML
  - Integration with existing file I/O for YAML file processing
  
  **Sample:**
  ```js
  import yaml
  
  let config_yaml = `
  database:
    host: localhost
    port: 5432
  features:
    - authentication
    - logging
  `
  
  let config = yaml.decode(config_yaml)
  let db_host = yaml.get(config, "database.host")  // "localhost"
  let features = yaml.get(config, "features")      // ["authentication", "logging"]
  
  // Modify and encode back
  config["environment"] = "production"
  let output_yaml = yaml.encode(config)
  ```

- **Enhanced String Module**
  Added new string manipulation functions moved from builtins:
  - `string.startsWith(str, prefix)` - checks if string starts with prefix
  - `string.endsWith(str, suffix)` - checks if string ends with suffix
  - `string.chr(code)` - converts ASCII/Unicode code to character
  - `string.ord(char)` - converts character to ASCII/Unicode code
  - Better integration with existing string functions
  - Consistent error handling and validation
  
  **Sample:**
  ```js
  import string
  
  println(string.startsWith("hello world", "hello"))  // true
  println(string.endsWith("hello world", "world"))    // true
  println(string.chr(65))                             // "A"
  println(string.ord("A"))                            // 65
  ```

- **Enhanced Reflect Module**
  Expanded runtime type inspection capabilities:
  - Enhanced existing functions: `typeOf`, `valueOf`, `isNil`, `isArray`, `isObject`, `isFunction`
  - Better error messages and validation
  - Improved integration with the new builtin system
  - Comprehensive type checking for all VintLang data types
  
  **Sample:**
  ```js
  import reflect
  
  let arr = [1, 2, 3]
  println(reflect.typeOf(arr))    // "ARRAY"
  println(reflect.isArray(arr))   // true
  println(reflect.isNil(null))    // true
  ```

### Language Improvements

- **For Loop Enhancements**
  Improved for..in loop functionality with better iteration handling:
  - Implemented `IsolatedIterator` for safe nested iteration
  - Enhanced error messages for iteration issues
  - Better handling of different iterable types
  - Improved loop variable scoping and safety
  - Fixed edge cases in nested loop scenarios

- **Documentation System**
  Enhanced embedded documentation support:
  - Interactive documentation command in REPL
  - Embedded docs in binary for offline access
  - Improved documentation generation and categorization
  - Better help system integration
  - Streamlined documentation structure

### Error Handling & Developer Experience

- **Enhanced Error Messages**
  Comprehensive improvements to error reporting across the language:
  - Added filename support to lexer, parser, and REPL
  - Enhanced column and line tracking for precise error location
  - Improved error messages for type mismatches and operations
  - Better context in error messages with code snippets
  - Structured error handling throughout the codebase

- **Code Quality Improvements**
  Extensive refactoring and cleanup:
  - Standardized function declarations to use `let` syntax
  - Improved variable naming consistency
  - Enhanced code formatting and whitespace handling
  - Better separation of concerns in modules
  - Removed outdated and obsolete files

### Bug Fixes

- **Import System Fixes**
  - Fixed import statement parsing and evaluation
  - Improved module loading reliability
  - Better error handling for missing modules
  - Enhanced import path resolution

- **Loop and Control Flow**
  - Fixed for..in loop edge cases and nested iteration issues
  - Improved control flow handling in various scenarios
  - Better variable scoping in loops and functions

- **Documentation and Examples**
  - Updated examples to use current syntax and best practices
  - Fixed various syntax issues in example code
  - Improved consistency across code samples

### Development Tools

- **Bundler Improvements**
  - Enhanced bundler architecture with better string processing
  - Improved package processing capabilities
  - Added bundler visualization and documentation
  - Better integration with the overall build system

- **VSCode Extension**
  - Added VSCode extension as submodule for better integration
  - Improved development workflow
  - Better syntax highlighting and language support

## [August-2025] – Previous Major & Minor Updates

### New Features (Week of August 9, 2025)

- **Dict Pattern Matching**
  Introduced a powerful new `match` statement for pattern matching on dictionaries:
  - Syntax: `match value { pattern => action }`
  - Supports dictionary pattern matching with specific key-value pairs
  - Wildcard pattern `_` for default cases
  - Full lexer, parser, and evaluator integration with new `MATCH` keyword and `=>` arrow token
  
  **Sample:**
  ```js
  let user = {"role": "admin", "active": true}
  match user {
      {"role": "admin"} => print("Admin user!")
      {"active": false} => print("Inactive user")
      _ => print("Regular user")
  }
  ```

- **Enhanced Scheduling Module**
  Complete rewrite of the schedule module with full ticker and cron functionality:
  - `ticker(intervalSeconds, callback)` - executes functions at regular intervals
  - `schedule(cronExpr, callback)` - cron-based scheduling with proper expression parsing
  - Helper functions: `everySecond()`, `everyMinute()`, `everyHour()`, `daily(hour, minute)`
  - Support for step values in cron expressions (e.g., `*/5` for every 5 seconds)
  - Comprehensive error handling and validation
  - Proper cleanup with `stopTicker()` and `stopSchedule()` functions
  
  **Sample:**
  ```js
  import schedule
  
  // Every 5 seconds
  let ticker = schedule.ticker(5, func() {
      print("Tick!")
  })
  
  // Daily at 9:30 AM
  let job = schedule.schedule("0 30 9 * * *", func() {
      print("Good morning!")
  })
  
  // Using helper functions
  let minutely = schedule.everyMinute(func() {
      print("New minute!")
  })
  ```

- **HTTP Module Improvements**
  Enhanced Express.js-like HTTP server functionality:
  - `http.app()` - creates Express-like application instances
  - Route registration methods: `http.get()`, `http.post()`, `http.put()`, `http.delete()`, `http.patch()`
  - Middleware support with `http.use()`
  - Improved request/response handling with better route matching
  - Enhanced HTTP handler execution with request information display
  - Better 404 handling and error responses
  
  **Sample:**
  ```js
  import http
  
  http.app()
  
  http.get("/", func(req, res) {
      print("Home page accessed")
  })
  
  http.post("/api/users", func(req, res) {
      print("Creating new user")
  })
  
  http.use(func(req, res, next) {
      print("Middleware executed")
  })
  
  http.listen(3000, "Server running on port 3000")
  ```

### Language Features

- **Println**
  - `println`: Print value with a newline which was something `print` could do but now `print` just prints the value without new-lines

- **Declarative Statements**
  Introduced new declarative keywords for expressive, styled runtime messages:
  - `info`: Print informational messages in cyan.
  - `debug`: Print debug messages in magenta for troubleshooting.
  - `note`: Print general notes in blue for context or reminders.
  - `success`: Print success messages in green to indicate successful operations.

  **Sample:**
  ```js
  info "Starting backup..."
  debug "Current value: " + str(42)
  note "This script was last updated on 2024-06-01."
  success "Backup completed successfully!"
  ```

- **Repeat Loops**
  Added the `repeat` keyword for concise, fixed-count iteration:
  - Syntax: `repeat 5 { ... }` or `repeat n { ... }`
  - The default loop variable `i` is available inside the block, starting from 0.
  - Supports `break` and `continue` for flexible control flow.
  
  **Sample:**
  ```js
  repeat 3 {
      println("Iteration:", i)
  }
  // Output:
  // Iteration: 0
  // Iteration: 1
  // Iteration: 2
  ```

- **Function Default Parameters & Overloading**
  Enhanced function definitions and calls:
  - Functions can now specify default values for parameters (e.g., `func(name = "Guest")`).
  - Function overloading and default parameters work seamlessly together.
  - Improved error messages for missing or ambiguous arguments, including line numbers and code snippets.
  
  **Sample:**
  ```js
  let greet = func(name = "Guest") {
      println("Hello, " + name)
  }
  greet()        // Hello, Guest
  greet("Alice") // Hello, Alice

  let add = func(a, b = 10) {
      return a + b
  }
  println(add(5))      // 15
  println(add(5, 2))   // 7
  ```

- **Pointer Improvements**
  Improved pointer handling and documentation:
  - Added clear, safe pointer operations and examples.
  - See the updated `pointers.md` for details and best practices.
  
  **Sample:**
  ```js
  let x = 10
  let p = &x      // Create a pointer to x
  println(*p)     // Dereference pointer, prints 10
  *p = 20         // Set value via pointer
  println(x)      // Prints 20
  ```

- **Error Messages**
  All runtime and compile-time errors are now more descriptive:
  - Function call errors include the function name, argument count, line number, and a code snippet.
  - General error messages are clearer and more actionable for users.
  
  **Sample:**
  ```
  greet(1, 2, 3)
  // Error: No matching overload for function 'greet' with 3 arguments at line 5. Source: greet(1, 2, 3)
  ```

### Documentation & Examples
- Comprehensive documentation for all new features, including declaratives, repeat loops, and pointer usage.
- Updated and expanded code examples and tests to cover new language constructs and behaviors.

### Refactoring & Code Quality
- Reduced code duplication in built-in functions and improved internal consistency.
- Refactored database modules (`mysql`, `postgres`) for clarity and reliability.
- Enhanced error handling and reporting throughout the codebase.

### Bug Fixes
- Fixed error message formatting in the evaluator and math module.
- Improved handling of `nil`/`null` values and edge cases in core modules.

### Miscellaneous
- Cleaned up obsolete files and improved project structure for maintainability.
- Commented out the main function in `test_vm.go` for easier manual testing.

---

### Compiler & Virtual Machine (VM) Foundation

- **New Bytecode Compiler & VM**
  - Introduced a modern bytecode compiler and virtual machine for VintLang.
  - Supports integer arithmetic, boolean logic, and comparison operators.
  - Comprehensive tests for the compiler and VM ensure correctness and performance.
  - Lays the groundwork for future optimizations and advanced language features.

  **Sample:**
  ```go
  // See test_vm.go for bytecode/VM usage and manual tests
  ```

---

**Summary:**
These updates make VintLang more expressive, robust, and user-friendly. The language now supports modern declarative statements, flexible iteration, advanced function features, and improved error reporting—all backed by a new VM/compiler foundation for future growth. 