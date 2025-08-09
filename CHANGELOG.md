# Changelog

## [Unreleased] – Recent Major & Minor Updates

### New Features (Week of August 9, 2025)

- **Dict Pattern Matching**
  Introduced a powerful new `match` statement for pattern matching on dictionaries:
  - Syntax: `match value { pattern => action }`
  - Supports dictionary pattern matching with specific key-value pairs
  - Wildcard pattern `_` for default cases
  - Full lexer, parser, and evaluator integration with new `MATCH` keyword and `=>` arrow token
  
  **Sample:**
  ```vint
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
  ```vint
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
  ```vint
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

- **Declarative Statements**
  Introduced new declarative keywords for expressive, styled runtime messages:
  - `info`: Print informational messages in cyan.
  - `debug`: Print debug messages in magenta for troubleshooting.
  - `note`: Print general notes in blue for context or reminders.
  - `success`: Print success messages in green to indicate successful operations.
  
  **Sample:**
  ```vint
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
  ```vint
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
  ```vint
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
  ```vint
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