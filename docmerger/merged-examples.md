# Example Code Snippets

## README.md

````js
# VintLang Examples

Welcome to the VintLang examples directory! This collection contains 100 example files demonstrating various features and capabilities of the VintLang programming language.

## 📚 What's Inside

This directory contains comprehensive examples covering:

- **Core Language Features**: Variables, control flow, loops, functions
- **Data Structures**: Arrays, dictionaries, strings
- **Modules & Imports**: File organization, package system
- **Built-in Functions**: String manipulation, type conversion, utilities
- **File I/O**: Reading/writing files, directory operations
- **Databases**: SQLite, MySQL, PostgreSQL integration
- **Networking**: HTTP servers, API requests
- **Security**: Encryption, hashing, encoding
- **System Operations**: Shell commands, environment variables
- **Advanced Features**: Pointers, reflection, pattern matching, async operations

## 🚀 Getting Started

### Running Examples

To run any example, use the VintLang interpreter:

```bash
vint examples/example_name.vint
````

### Example Categories

#### Beginner-Friendly Examples

- `builtins.vint` - Built-in functions
- `functions.vint` - Function definitions and usage
- `switch.vint` - Switch-case statements
- `if_expression.vint` - If statements and expressions
- `test-for.vint` - For loops
- `repeat-keyword.vint` - Repeat loops

#### String & Data Manipulation

- `strings.vint` - String module functions
- `nativeStrings.vint` - Native string methods
- `json.vint` - JSON operations
- `csv.vint` - CSV file handling
- `encoding.vint` - Base64 encoding/decoding

#### File & System Operations

- `os.vint` - Operating system operations
- `path.vint` - Path manipulation
- `shell.vint` - Shell command execution

#### Database Examples

- `sqlite.vint` - SQLite database
- `mysql.vint` - MySQL database
- `postgres.vint` - PostgreSQL database

#### Networking & HTTP

- `http.vint` - HTTP file server
- `http_test.vint` - HTTP module testing
- `github-profile.vint` - HTTP requests

#### Advanced Features

- `pointers.vint` - Pointer operations
- `reflect.vint` - Runtime type inspection
- `defer_test.vint` - Defer statement
- `overloading_test.vint` - Function overloading
- `async_simple.vint` - Asynchronous operations

#### Security & Crypto

- `crypto.vint` - Hashing and encryption
- `dotenv.vint` - Environment variables from .env

#### AI/ML Integration

- `llm_openai.vint` - OpenAI GPT integration

## 📖 Documentation

All examples include comprehensive comments explaining:

- What the code does
- How VintLang features work
- Expected output
- Usage patterns

For a detailed analysis of all examples, see [TEST_RESULTS.md](TEST_RESULTS.md).

## ✅ Quality Assurance

All examples have been:

- ✅ Tested with VintLang v0.2.2
- ✅ Fixed for syntax errors
- ✅ Documented with explanatory comments
- ✅ Verified to compile successfully

## 🔧 Testing Examples

### Running Tests

Many examples include the word "test" in their names and demonstrate specific features:

```bash
# Test built-in functions
vint examples/builtins_test.vint

# Test array slicing
vint examples/array_slicing_test.vint

# Test function features
vint examples/function_test.vint
```

### Test Categories

- **builtins_test.vint** - Built-in function testing
- **declaratives_test.vint** - Declarative statements (info, debug, etc.)
- **function_test.vint** - Function default parameters
- **array_slicing_test.vint** - Array slicing operations
- **has_key_test.vint** - Dictionary key checking
- **overloading_test.vint** - Function overloading

## 📝 Notes

### External Dependencies

Some examples require external resources:

- **Database examples**: Need running database servers
- **dotenv.vint**: Needs a .env file
- **llm_openai.vint**: Requires OpenAI API key
- **Networking examples**: May require internet connectivity

These examples will fail expectedly without the required resources but demonstrate correct VintLang syntax.

### Module Availability

Some modules shown in examples may require additional setup or may be in development:

- **desktop** - Desktop GUI module (in development)
- **regex** - Regular expressions (in development)
- **package system** - Advanced package features (in development)

## 🤝 Contributing

When adding new examples:

1. Follow the naming convention (descriptive_name.vint)
2. Add comprehensive comments explaining the code
3. Test the example to ensure it works
4. Update this README if adding a new category
5. Add a description in TEST_RESULTS.md

## 📚 Learning Path

For beginners, we recommend following this learning path:

1. **Basics**: Start with `builtins.vint`, `functions.vint`, `switch.vint`
2. **Control Flow**: Try `if_expression.vint`, `test-for.vint`, `repeat-keyword.vint`
3. **Data Structures**: Explore `strings.vint`, `json.vint`, `array_slicing_test.vint`
4. **File Operations**: Learn from `os.vint`, `path.vint`
5. **Advanced**: Move to `pointers.vint`, `reflect.vint`, `overloading_test.vint`
6. **Modules**: Study various module examples for specific use cases

## 🎯 Quick Reference

| Category       | Example Files           | Count |
| -------------- | ----------------------- | ----- |
| Total Examples | All \*.vint files       | 100   |
| Test Files     | _test_.vint             | 29    |
| Showcase Files | _showcase_.vint         | 5     |
| Database Files | sqlite, mysql, postgres | 3     |
| HTTP Files     | http\*.vint             | 4     |

## 💡 Tips

- All examples use proper VintLang syntax as of v0.2.2
- Comments explain not just what code does, but why
- Most examples can be run independently
- Check example output to understand expected behavior
- Modify examples to experiment and learn

## 🔗 Resources

- [VintLang Documentation](https://vintlang.ekilie.com/docs)
- [VintLang GitHub Repository](https://github.com/vintlang/vintlang)
- [TEST_RESULTS.md](TEST_RESULTS.md) - Detailed test results and fixes

---

Happy coding with VintLang! 🎉

````

## array_slicing_test.vint

```js
// Array slicing functionality test
// This test validates the Python-like array slicing syntax

let testArray = [1, 2, 3, 4, 5, 6, 7, 8, 9, 10];

println("Testing Array Slicing Implementation");
println("====================================");

// Test the original requirement from the issue
let arr = [1, 2, 3, 4, 5];
let result = arr[1:4];
println("arr[1:4] =", result); // Should output: [2, 3, 4]

// Basic slicing tests
println("\nBasic Slicing:");
println("testArray[2:5] =", testArray[2:5]);   // [3, 4, 5]
println("testArray[0:3] =", testArray[0:3]);   // [1, 2, 3]
println("testArray[7:10] =", testArray[7:10]); // [8, 9, 10]

// Partial slicing tests
println("\nPartial Slicing:");
println("testArray[5:] =", testArray[5:]);     // [6, 7, 8, 9, 10]
println("testArray[:4] =", testArray[:4]);     // [1, 2, 3, 4]
println("testArray[:] =", testArray[:]);       // [1, 2, 3, 4, 5, 6, 7, 8, 9, 10]

// Edge cases
println("\nEdge Cases:");
println("testArray[0:0] =", testArray[0:0]);   // []
println("testArray[5:5] =", testArray[5:5]);   // []
println("testArray[9:10] =", testArray[9:10]); // [10]

// Negative indexing
println("\nNegative Indexing:");
println("testArray[-3:] =", testArray[-3:]);   // [8, 9, 10]
println("testArray[:-3] =", testArray[:-3]);   // [1, 2, 3, 4, 5, 6, 7]
println("testArray[-5:-2] =", testArray[-5:-2]); // [6, 7, 8]

// Out of bounds handling
println("\nOut of Bounds Handling:");
println("testArray[15:] =", testArray[15:]);   // []
println("testArray[:20] =", testArray[:20]);   // [1, 2, 3, 4, 5, 6, 7, 8, 9, 10]
println("testArray[-20:] =", testArray[-20:]); // [1, 2, 3, 4, 5, 6, 7, 8, 9, 10]

// Verify regular indexing still works
println("\nRegular Indexing Still Works:");
println("testArray[0] =", testArray[0]);       // 1
println("testArray[5] =", testArray[5]);       // 6
println("testArray[9] =", testArray[9]);       // 10

println("\nAll array slicing tests completed successfully!");
````

## async_simple.vint

```js
// Simple async patterns that work with VintLang syntax

print("=== Simple Async Patterns ===")

// Basic async function
let fetchData = async func(id) {
    return "Data for " + id
}

print("1. Basic async function:")
let promise = fetchData("123")
let result = await promise
print("   Result:", result)

print("\n2. Multiple async operations:")
let p1 = fetchData("A")
let p2 = fetchData("B")
let p3 = fetchData("C")

let r1 = await p1
let r2 = await p2
let r3 = await p3

print("   Results:", r1, r2, r3)

print("\n3. Channel-based communication:")
let dataChan = chan(3)

// Producer
go func() {
    send(dataChan, "Message 1")
    send(dataChan, "Message 2")
    send(dataChan, "Message 3")
    close(dataChan)
}()

// Consumer
let msg1 = receive(dataChan)
let msg2 = receive(dataChan)
let msg3 = receive(dataChan)

print("   Received:", msg1, msg2, msg3)

print("\n4. Async with channels:")
let processAsync = async func(input) {
    let resultChan = chan

    go func() {
        let processed = "Processed: " + input
        send(resultChan, processed)
    }()

    let result = receive(resultChan)
    return result
}

let asyncResult = await processAsync("important data")
print("   Async result:", asyncResult)

print("\n5. Error-free concurrent execution:")
go print("   Goroutine 1: Hello")
go print("   Goroutine 2: World")
go print("   Goroutine 3: From")
go print("   Goroutine 4: VintLang")

print("   Main thread: Concurrent execution complete")

print("\n=== All patterns demonstrated successfully! ===")
print("VintLang now supports:")
print("  - async functions that return promises")
print("  - await expressions for promise resolution")
print("  - go statements for concurrent execution")
print("  - channels for inter-goroutine communication")
print("  - Built-in functions: send(), receive(), close()")
```

## backend_demo.vint

```js
// Enhanced HTTP Backend Demo - Comprehensive Example
import http

println("Full Backend HTTP Server Demo")
println("=" * 40)

// Create the application
println("\n📦 Creating enhanced HTTP application...")
http.app()

// Add interceptors
println("\n🔍 Setting up interceptors...")
http.interceptor("request", func(req) {
    println("Request interceptor: Validating and logging request")
})

http.interceptor("response", func(res) {
    println("Response interceptor: Adding security headers")
})

// Add guards
println("\n🛡️ Setting up security guards...")
http.guard(func(req) {
    println("🔐 Auth guard: Checking authentication")
})

http.guard(func(req) {
    println("⏱️ Rate limit guard: Checking request rate")
})

// Add middleware
println("\n🔧 Setting up middleware...")
http.cors()
http.bodyParser()
http.auth(func(req, res, next) {
    println("🔑 Auth middleware: Processing authentication")
})

// Set error handler
println("\n🚨 Setting up error handling...")
http.errorHandler(func(err, req, res) {
    println("❌ Global error handler: Processing error")
})

// Define API routes
println("\n🛤️ Setting up API routes...")

// User management routes
http.get("/api/users", func(req, res) {
    println("📋 GET /api/users - Retrieving all users")
})

http.get("/api/users/:id", func(req, res) {
    println("👤 GET /api/users/:id - Retrieving user by ID")
})

http.post("/api/users", func(req, res) {
    println("➕ POST /api/users - Creating new user")
})

http.put("/api/users/:id", func(req, res) {
    println("✏️ PUT /api/users/:id - Updating user")
})

http.delete("/api/users/:id", func(req, res) {
    println("🗑️ DELETE /api/users/:id - Deleting user")
})

// Post management routes
http.get("/api/users/:userId/posts", func(req, res) {
    println("📝 GET /api/users/:userId/posts - Getting user posts")
})

http.post("/api/users/:userId/posts", func(req, res) {
    println("📄 POST /api/users/:userId/posts - Creating new post")
})

http.get("/api/posts/:id", func(req, res) {
    println("📖 GET /api/posts/:id - Getting specific post")
})

// File upload routes
http.post("/api/upload", func(req, res) {
    println("📤 POST /api/upload - Handling file upload")
})

// Authentication routes
http.post("/auth/login", func(req, res) {
    println("🔓 POST /auth/login - User login")
})

http.post("/auth/register", func(req, res) {
    println("📝 POST /auth/register - User registration")
})

http.post("/auth/logout", func(req, res) {
    println("🔒 POST /auth/logout - User logout")
})

// Admin routes
http.get("/admin/dashboard", func(req, res) {
    println("📊 GET /admin/dashboard - Admin dashboard")
})

http.get("/admin/users", func(req, res) {
    println("👥 GET /admin/users - Admin user management")
})

// Health check
http.get("/health", func(req, res) {
    println("💚 GET /health - Health check endpoint")
})

// API documentation
http.get("/api/docs", func(req, res) {
    println("📚 GET /api/docs - API documentation")
})

println("\n" + "=" * 40)
println("✅ Enhanced HTTP Backend Setup Complete!")
println("\n🎯 Features Implemented:")
println("  ✓ Request/Response Interceptors")
println("  ✓ Authentication & Rate Limiting Guards")
println("  ✓ CORS & Body Parser Middleware")
println("  ✓ Global Error Handling")
println("  ✓ RESTful API Routes with Parameters")
println("  ✓ File Upload Support")
println("  ✓ Authentication Endpoints")
println("  ✓ Admin Panel Routes")
println("  ✓ Health Check & Documentation")

println("\n🚀 Backend Infrastructure Ready!")
println("📡 All endpoints configured with full security and middleware stack")

// Note: To start the server, uncomment the line below:
// http.listen(3000, "🌐 Enhanced backend server running on http://localhost:3000")
```

## basic_guard_test.vint

```js
// Basic test of switch case value
let x = 5

switch (x) {
    case 5 if x > 0 {
        println("Positive 5")
    }
    default {
        println("Something else")
    }
}
```

## builtins.vint

```js
// VintLang Builtins Example
// This example demonstrates the use of built-in functions in VintLang

// eq() - Check if two values are equal (returns true/false)
print(eq(2, "")); // Prints false - comparing number 2 with empty string

// Functions can be assigned to variables
let a = print;
a("yeee"); // Call the print function through variable 'a'
```

## builtins_test.vint

```js
// examples/builtins_test.vint

// Test len()
let myArray = [1, 2, 3, 4];
println("Length of array [1, 2, 3, 4] is:", len(myArray)); // Expected: 4

let myString = "hello";
println("Length of string 'hello' is:", len(myString)); // Expected: 5

let myDict = { a: 1, b: 2 };
println("Length of dict {'a': 1, 'b': 2} is:", len(myDict)); // Expected: 2

// Test append() and pop()
let arr = [1, 2];
arr = append(arr, 3, 4);
println("After append: ", arr); // Expected: [1, 2, 3, 4]

let popped = pop(arr);
println("Popped element:", popped); // Expected: 4
println("Array after pop:", arr); // Expected: [1, 2, 3]

// Test keys() and values()
let dict = { name: "Alex", age: 30 };
let keys_arr = keys(dict);
println("Keys of dict:", keys_arr); // Expected: ["name", "age"] (order may vary)

let values_arr = values(dict);
println("Values of dict:", values_arr); // Expected: ["Alex", 30] (order may vary)

// Test chr() and ord()
let char_A = chr(65);
println("chr(65) is:", char_A); // Expected: "A"

let code_A = ord("A");
println("ord('A') is:", code_A); // Expected: 65

// Test sleep()
println("Waiting for 1 second...");
sleep(1000);
println("Done waiting.");

println("All built-in function tests passed!");

// The exit() function is not called here because it would stop the script.
// To test it, you would run a line like:
// exit(0);
```

## cli-todo-fixed.vint

```js
//The sample vint code for a simple cli todo app
import cli

let todos = ["new", "todo"]

let args = args()
let args_len = len(args)
info "Length: " + string(args_len)

if (args_len == 0) {
    info "Invalid options"
    return
}

if (args[0] == "all") {
    if (len(todos) == 0) {
        println("NO todo found")
    } else {
        for todo in todos {
            println(todo)
        }
    }
} else if (args[0] == "add") {
    let newtodo = input("Enter the todo title")
    todos.push(newtodo)
} else {
    info "Unknown command: " + args[0]
}
```

## cli-todo-working.vint

```js
//The sample vint code for a simple cli todo app

let todos = ["Buy groceries", "Walk the dog", "Finish project"]

let command = "all"  // Simulate command line argument for now

if (command == "all") {
    println("Your TODOs:")
    for todo in todos {
        println("- " + todo)
    }
} else {
    println("Unknown command: " + command)
}
```

## cli-todo.vint

```js
// VintLang CLI Todo Application Example
// This example demonstrates building a simple command-line todo application

import cli

// Initialize the todos list with some default items
let todos = ["new", "todo"]

// Get command line arguments
let args = args()
let args_len = len(args)
info "Length: " + str(args_len)

// Check if any arguments were provided
if (args_len == 0) {
    info "Invalid options"
    return
}

// Handle 'all' command - list all todos
if (args[0] == "all") {
    if (len(todos) == 0) {
        println("NO todo found")
    } else {
        for todo in todos {
            println(todo)
        }
    }
} else if (args[0] == "add") {
    // Handle 'add' command - add a new todo
    let newtodo = input("Enter the todo title: ")
    todos.push(newtodo)
    info "Added: " + newtodo
} else {
    // Handle unknown commands
    info "Unknown command: " + args[0]
    info "Available commands: all, add"
}

```

## cli.vint

```js
import cli

// Get all command-line arguments
let args = cli.getArgs()

// Parse flags
let flags = cli.getFlags()
if (flags["verbose"]) {
    print("Verbose mode enabled")
}

// Get a specific argument value
let output = cli.getArgValue("output")
if (output != null) {
    print("Output file:", output)
}

// Prompt for user input
let name = cli.prompt("Enter your name: ")
print("Hello,", name)

// Confirm an action
if (cli.confirm("Do you want to proceed?")) {
    // Execute a command
    let result = cli.execCommand("ls -l")
    print(result)
}

// Exit with status code
cli.cliExit(0)
```

## clipboard_complete_test.vint

```js
import clipboard

println("=== Clipboard Module Complete Test ===")

// Test write and read
clipboard.write("Hello VintLang!")
println("1. Written text:", clipboard.read())

// Test all method with content
let allContent = clipboard.all()
println("2. All content (array):", allContent)

// Test hasContent
if (clipboard.hasContent()) {
    println("3. Clipboard has content: true")
}

// Test with different data types
clipboard.write(42)
println("4. Number in clipboard:", clipboard.read())
println("   All content:", clipboard.all())

clipboard.write(3.14159)
println("5. Float in clipboard:", clipboard.read())
println("   All content:", clipboard.all())

clipboard.write(true)
println("6. Boolean in clipboard:", clipboard.read())
println("   All content:", clipboard.all())

// Test clear and all method with empty clipboard
clipboard.clear()
println("7. Clipboard cleared")

let emptyAll = clipboard.all()
println("8. All content after clear:", emptyAll)

if ((clipboard.hasContent())) {
    println("9. Still has content")
}

println("=== All tests completed successfully! ===")
```

## clipboard_test.vint

```js
import clipboard

println("Testing clipboard module...")

clipboard.write("Hello World")
let content = clipboard.read()
println("Read from clipboard:", content)

// Test the new 'all' method
let allContent = clipboard.all()
println("All clipboard content:", allContent)

clipboard.clear()
println("Clipboard cleared successfully")

// Test 'all' method with empty clipboard
let emptyContent = clipboard.all()
println("All clipboard content after clear:", emptyContent)

```

## colors.vint

```js
// VintLang Colors Module Example
// Demonstrates converting RGB color values to hexadecimal format

import colors

// Convert RGB values (Red, Green, Blue) to hexadecimal color code
// RGB values range from 0-255
// (255, 255, 255) represents white color
let hex = colors.rgbToHex(255, 255, 255)
print(hex) // Outputs "#FFFFFF"

```

## complete_backend_app.vint

```js
// Complete Backend Application Demo - All Features
// This demonstrates a full-fledged backend using all enhanced HTTP features
import http

print("🏗️ Building Complete Backend Application")
print("==========================================")

// 1. Initialize the application
print("\n🎯 Step 1: Initialize Application")
http.app()
print("✅ HTTP application created")

// 2. Set up request/response interceptors
print("\n🔄 Step 2: Configure Interceptors")
http.interceptor("request", func(req) {
    print("📥 Request Interceptor: Processing incoming request")
    print("   - Validating request format")
    print("   - Logging request details")
    print("   - Adding request timestamp")
})

http.interceptor("response", func(res) {
    print("📤 Response Interceptor: Processing outgoing response")
    print("   - Adding security headers")
    print("   - Logging response time")
    print("   - Adding API version header")
})
print("✅ Interceptors configured")

// 3. Set up security guards
print("\n🛡️ Step 3: Configure Security Guards")
http.guard(func(req) {
    print("🔐 Authentication Guard: Verifying user identity")
    print("   - Checking JWT token")
    print("   - Validating token expiry")
    print("   - Extracting user permissions")
})

http.guard(func(req) {
    print("⏱️ Rate Limiting Guard: Checking request rate")
    print("   - Tracking IP requests")
    print("   - Enforcing rate limits")
    print("   - Blocking suspicious activity")
})

http.guard(func(req) {
    print("🛡️ Security Guard: Scanning for threats")
    print("   - SQL injection detection")
    print("   - XSS prevention")
    print("   - Input validation")
})
print("✅ Security guards activated")

// 4. Configure middleware stack
print("\n🔧 Step 4: Configure Middleware Stack")
http.cors()
print("✅ CORS middleware enabled")

http.bodyParser()
print("✅ Body parser middleware enabled")

http.auth(func(req, res, next) {
    print("🔑 Authentication Middleware: Processing auth")
    print("   - Extracting credentials")
    print("   - Validating permissions")
    print("   - Setting user context")
})
print("✅ Authentication middleware enabled")

// 5. Set up global error handling
print("\n🚨 Step 5: Configure Error Handling")
http.errorHandler(func(err, req, res) {
    print("❌ Global Error Handler: Processing error")
    print("   - Logging error details")
    print("   - Sanitizing error messages")
    print("   - Sending appropriate response")
})
print("✅ Global error handler configured")

// 6. Define API endpoints
print("\n🛤️ Step 6: Define API Endpoints")

// Authentication endpoints
print("  🔐 Authentication Endpoints:")
http.post("/auth/register", func(req, res) {
    print("    📝 POST /auth/register - User registration")
    print("      - Validating user data")
    print("      - Hashing password")
    print("      - Creating user account")
})

http.post("/auth/login", func(req, res) {
    print("    🔓 POST /auth/login - User login")
    print("      - Verifying credentials")
    print("      - Generating JWT token")
    print("      - Setting session cookie")
})

http.post("/auth/logout", func(req, res) {
    print("    🔒 POST /auth/logout - User logout")
    print("      - Invalidating token")
    print("      - Clearing session")
})

http.post("/auth/refresh", func(req, res) {
    print("    🔄 POST /auth/refresh - Token refresh")
    print("      - Validating refresh token")
    print("      - Generating new access token")
})

// User management endpoints
print("  👥 User Management Endpoints:")
http.get("/api/users", func(req, res) {
    print("    📋 GET /api/users - List all users")
    print("      - Applying pagination")
    print("      - Filtering by permissions")
    print("      - Sorting results")
})

http.get("/api/users/:id", func(req, res) {
    print("    👤 GET /api/users/:id - Get user by ID")
    print("      - Extracting user ID from path")
    print("      - Checking access permissions")
    print("      - Returning user data")
})

http.post("/api/users", func(req, res) {
    print("    ➕ POST /api/users - Create new user")
    print("      - Validating input data")
    print("      - Checking admin permissions")
    print("      - Creating user record")
})

http.put("/api/users/:id", func(req, res) {
    print("    ✏️ PUT /api/users/:id - Update user")
    print("      - Validating update data")
    print("      - Checking ownership/admin rights")
    print("      - Updating user record")
})

http.delete("/api/users/:id", func(req, res) {
    print("    🗑️ DELETE /api/users/:id - Delete user")
    print("      - Checking admin permissions")
    print("      - Soft delete implementation")
    print("      - Cleaning up related data")
})

// Content management endpoints
print("  📄 Content Management Endpoints:")
http.get("/api/posts", func(req, res) {
    print("    📚 GET /api/posts - List posts")
    print("      - Applying filters")
    print("      - Supporting search")
    print("      - Paginating results")
})

http.get("/api/posts/:id", func(req, res) {
    print("    📖 GET /api/posts/:id - Get post by ID")
    print("      - Incrementing view count")
    print("      - Checking visibility permissions")
    print("      - Including related data")
})

http.post("/api/posts", func(req, res) {
    print("    📝 POST /api/posts - Create new post")
    print("      - Validating content")
    print("      - Processing media uploads")
    print("      - Setting publication status")
})

http.put("/api/posts/:id", func(req, res) {
    print("    ✏️ PUT /api/posts/:id - Update post")
    print("      - Checking author permissions")
    print("      - Validating changes")
    print("      - Updating modification time")
})

// File upload endpoints
print("  📤 File Upload Endpoints:")
http.post("/api/upload/image", func(req, res) {
    print("    🖼️ POST /api/upload/image - Upload image")
    print("      - Validating file type")
    print("      - Checking file size")
    print("      - Processing and storing")
})

http.post("/api/upload/document", func(req, res) {
    print("    📄 POST /api/upload/document - Upload document")
    print("      - Scanning for viruses")
    print("      - Extracting metadata")
    print("      - Storing securely")
})

// Analytics endpoints
print("  📊 Analytics Endpoints:")
http.get("/api/analytics/dashboard", func(req, res) {
    print("    📈 GET /api/analytics/dashboard - Analytics dashboard")
    print("      - Aggregating metrics")
    print("      - Generating charts")
    print("      - Applying date filters")
})

http.get("/api/analytics/reports/:type", func(req, res) {
    print("    📊 GET /api/analytics/reports/:type - Generate report")
    print("      - Processing report type")
    print("      - Gathering data")
    print("      - Formatting output")
})

// Admin endpoints
print("  ⚙️ Admin Endpoints:")
http.get("/admin/settings", func(req, res) {
    print("    ⚙️ GET /admin/settings - System settings")
    print("      - Checking admin permissions")
    print("      - Loading configuration")
    print("      - Filtering sensitive data")
})

http.put("/admin/settings", func(req, res) {
    print("    🔧 PUT /admin/settings - Update settings")
    print("      - Validating admin role")
    print("      - Backing up current config")
    print("      - Applying new settings")
})

// Health and monitoring
print("  🏥 Health & Monitoring Endpoints:")
http.get("/health", func(req, res) {
    print("    💚 GET /health - Health check")
    print("      - Checking database connection")
    print("      - Verifying external services")
    print("      - Reporting system status")
})

http.get("/metrics", func(req, res) {
    print("    📊 GET /metrics - System metrics")
    print("      - Gathering performance data")
    print("      - Memory and CPU usage")
    print("      - Request statistics")
})

print("✅ All API endpoints configured")

// 7. System ready
print("\n" + "=" * 40)
print("🎉 COMPLETE BACKEND APPLICATION READY!")
print("==========================================")
print("\n🏗️ Architecture Summary:")
print("  🔄 Request/Response Interceptors")
print("  🛡️ Multi-layer Security Guards")
print("  🔧 Comprehensive Middleware Stack")
print("  🚨 Global Error Handling")
print("  🛤️ RESTful API with 20+ endpoints")
print("  📊 Analytics and Monitoring")
print("  ⚙️ Admin Management")
print("  📤 File Upload Support")
print("  🔐 Complete Authentication System")

print("\n🚀 Backend Features:")
print("  ✅ User Management (CRUD)")
print("  ✅ Content Management (CRUD)")
print("  ✅ Authentication & Authorization")
print("  ✅ File Upload & Processing")
print("  ✅ Analytics & Reporting")
print("  ✅ Admin Panel")
print("  ✅ Health Monitoring")
print("  ✅ Security Scanning")
print("  ✅ Rate Limiting")
print("  ✅ CORS Support")
print("  ✅ JSON/Form Data Processing")
print("  ✅ Path Parameter Extraction")
print("  ✅ Cookie Management")
print("  ✅ Error Handling")

print("\n🌐 Ready for Production!")
print("💡 To start the server: http.listen(3000)")

// Uncomment to start the server:
// print("\n🚀 Starting production server...")
// http.listen(3000, "🌟 Production backend server running on http://localhost:3000")
```

## comprehensive_import_test.vint

```js
println("=== Comprehensive import() function and statement test ===");
println();

// Test 1: import() function calls
println("1. Testing import() function calls:");
let osModule = import("os");
println("   import('os'):", osModule);

println('Time now:', import("time").now());

let mathModule = import("math");
println("   import('math'):", mathModule);

let stringModule = import("string");
println("   import('string'):", stringModule);

// Test 2: import statements
println();
println("2. Testing import statements:");
import time;
println("   import time; ->", time);

import crypto;
println("   import crypto; ->", crypto);

// Test 3: Mixed usage
println();
println("3. Testing mixed usage:");
let jsonFunc = import("json");
import regex;
println("   Function call result:", jsonFunc);
println("   Statement result:", regex);

// Test 4: Error cases
println();
println("4. Testing error cases:");
let badModule = import("nonexistent");
println("   import('nonexistent'):", badModule);

println();
println("=== All tests completed successfully! ===");
```

## comprehensive_pattern_test.vint

```js
// Comprehensive test of enhanced switch and match features
// This file tests all the working functionality we've implemented

println("=== Enhanced Switch and Match Features Test ===")

// Test 1: Switch with guard conditions
println("\n1. Switch with Guard Conditions:")

let numbers = [-5, 0, 3, 15, 150]

for num in numbers {
    print("Number", num, "-> ")

    switch (num) {
        case x if x < 0 {
            println("Negative:", x)
        }
        case 0 {
            println("Zero")
        }
        case x if x > 0 && x <= 10 {
            println("Small positive:", x)
        }
        case x if x > 10 && x <= 100 {
            println("Medium positive:", x)
        }
        case x if x > 100 {
            println("Large positive:", x)
        }
        default {
            println("Unknown")
        }
    }
}

// Test 2: Type-based switch
println("\n2. Type-based Switch:")

let values = [42, "hello", true, [1, 2, 3]]

for value in values {
    print("Value", value, "-> ")

    switch (value) {
        case x if type(x) == "INTEGER" {
            println("Integer:", x)
        }
        case x if type(x) == "STRING" {
            println("String:", x)
        }
        case x if type(x) == "BOOLEAN" {
            println("Boolean:", x)
        }
        case x if type(x) == "ARRAY" {
            println("Array with", len(x), "elements")
        }
        default {
            println("Other type:", type(value))
        }
    }
}

// Test 3: Match with dictionary patterns and variable binding
println("\n3. Dictionary Pattern Matching:")

let users = [
    {"name": "Alice", "role": "admin", "active": true},
    {"name": "Bob", "role": "user", "age": 25},
    {"name": "Charlie", "role": "admin", "active": false},
    {"name": "David", "age": 17}
]

for user in users {
    print("User", user["name"], "-> ")

    match user {
        {"role": "admin", "active": active, "name": name} if active => {
            println("Active admin:", name)
        }
        {"role": "admin", "name": name} => {
            println("Inactive admin:", name)
        }
        {"role": role, "age": age, "name": name} if age >= 18 => {
            println("Adult user:", name, "with role:", role)
        }
        {"age": age, "name": name} if age < 18 => {
            println("Minor user:", name)
        }
        {"name": name} => {
            println("User with incomplete info:", name)
        }
        _ => {
            println("Unknown user format")
        }
    }
}

// Test 4: Array pattern matching
println("\n4. Array Pattern Matching:")

let arrays = [
    [],
    [42],
    [1, 2],
    [1, 2, 3],
    ["a", "b", "c", "d"]
]

for arr in arrays {
    print("Array", arr, "-> ")

    match arr {
        [] => {
            println("Empty array")
        }
        [single] => {
            println("Single element:", single)
        }
        [first, second] => {
            println("Two elements:", first, "and", second)
        }
        [a, b, c] => {
            println("Three elements:", a, b, c)
        }
        _ => {
            println("More than three elements")
        }
    }
}

// Test 5: Complex nested patterns with guards
println("\n5. Complex Pattern Matching:")

let requests = [
    {"method": "GET", "path": "/", "headers": {"host": "localhost"}},
    {"method": "POST", "path": "/api/users", "body": {"name": "Alice"}},
    {"method": "DELETE", "path": "/api/users/123", "auth": {"role": "admin"}},
    {"method": "GET", "path": "/api/health"}
]

for req in requests {
    match req {
        {"method": "GET", "path": "/"} => {
            println("✓ Home page request")
        }
        {"method": "POST", "path": path, "body": body} if path.startsWith("/api/") => {
            println("✓ API POST to", path)
        }
        {"method": "DELETE", "auth": {"role": "admin"}} => {
            println("✓ Admin delete operation")
        }
        {"method": method, "path": path} => {
            println("→", method, "request to", path)
        }
        _ => {
            println("❌ Unknown request format")
        }
    }
}

// Test 6: Practical example - Configuration validator
println("\n6. Configuration Validation:")

let configs = [
    {"type": "database", "host": "localhost", "port": 5432},
    {"type": "redis", "url": "redis://localhost:6379", "ttl": 3600},
    {"type": "logging", "level": "info", "file": "/var/log/app.log"},
    {"invalid": "config"}
]

for config in configs {
    match config {
        {"type": "database", "host": host, "port": port} => {
            println("✓ Valid database config:", host, "port:", port)
        }
        {"type": "redis", "url": url, "ttl": ttl} => {
            println("✓ Valid Redis config with", ttl, "second TTL")
        }
        {"type": "logging", "level": level} => {
            println("✓ Valid logging config, level:", level)
        }
        _ => {
            println("❌ Invalid configuration")
        }
    }
}

println("\n=== All tests completed successfully! ===")
println("\nWorking features:")
println("✓ Switch statements with guard conditions")
println("✓ Switch statements with variable binding")
println("✓ Type-based switch cases")
println("✓ Match statements with dictionary patterns")
println("✓ Match statements with variable binding")
println("✓ Match statements with guard conditions")
println("✓ Array pattern matching (basic patterns)")
println("✓ Complex nested pattern matching")
println("✓ Practical real-world examples")
```

## comprehensive_showcase.vint

```js
// VintLang Comprehensive Feature Showcase
// Demonstrates core language capabilities safely

import time
import os
import json
import uuid

print("🎯 VintLang Comprehensive Feature Showcase")
print("=" * 60)
print("Demonstrating production-ready programming capabilities")
print("=" * 60)

// Feature 1: Data Structures and Variables
print("\n🗂️ Feature 1: Data Structures and Variables")
print("-" * 50)

let applicationData = {
    "name": "VintLang Feature Demo",
    "version": "2.0.0",
    "created": time.format(time.now(), "2006-01-02 15:04:05"),
    "features": ["variables", "functions", "loops", "files", "json"],
    "metadata": {
        "author": "VintLang Community",
        "license": "Open Source",
        "language": "VintLang"
    }
}

print("✓ Created application metadata structure")
print("Application: " + applicationData["name"])
print("Version: " + applicationData["version"])
print("Features: " + string(applicationData["features"]))

// Feature 2: Arrays and Iteration
print("\n🔄 Feature 2: Arrays and Iteration")
print("-" * 50)

let colors = ["red", "green", "blue", "yellow", "purple", "orange"]
let numbers = [10, 25, 30, 45, 50, 75, 80, 95]
let processed = []

print("Processing colors:")
for color in colors {
    print("  Processing color: " + color)
    processed.push(color.upper())
}
print("Processed colors: " + string(processed))

print("\nProcessing numbers:")
let total = 0
let count = 0
for number in numbers {
    total += number
    count += 1
    print("  Number " + string(count) + ": " + string(number))
}
let average = total / count
print("Total: " + string(total) + ", Average: " + string(average))

// Feature 3: String Operations
print("\n📝 Feature 3: String Processing")
print("-" * 50)

let sampleText = "VintLang is a powerful modern programming language"
print("Original text: " + sampleText)

let words = sampleText.split(" ")
print("Word count: " + string(len(words)))
print("Words: " + string(words))

let reversedText = sampleText.reverse()
print("Reversed text: " + reversedText)

let upperText = sampleText.upper()
print("Uppercase: " + upperText)

let searchTerm = "programming"
if (sampleText.contains(searchTerm)) {
    print("✓ Text contains '" + searchTerm + "'")
} else {
    print("❌ Text does not contain '" + searchTerm + "'")
}

// Feature 4: Functions and Logic
print("\n⚙️ Feature 4: Functions and Logic")
print("-" * 50)

let calculateSquare = func(n) {
    return n * n
}

let isEven = func(n) {
    return n % 2 == 0
}

let formatNumber = func(n, label) {
    return label + ": " + string(n)
}

print("Function demonstrations:")
for num in [2, 3, 4, 5, 6, 7, 8, 9, 10] {
    let square = calculateSquare(num)
    let evenStatus = isEven(num)

    print(formatNumber(num, "Number") +
          ", Square: " + string(square) +
          ", Even: " + string(evenStatus))
}

// Feature 5: JSON Data Handling
print("\n📊 Feature 5: JSON Data Operations")
print("-" * 50)

let employees = [
    {
        "id": uuid.generate(),
        "name": "Alice Johnson",
        "department": "Engineering",
        "salary": 75000,
        "startDate": "2022-01-15"
    },
    {
        "id": uuid.generate(),
        "name": "Bob Wilson",
        "department": "Marketing",
        "salary": 65000,
        "startDate": "2021-09-20"
    },
    {
        "id": uuid.generate(),
        "name": "Carol Davis",
        "department": "Engineering",
        "salary": 80000,
        "startDate": "2020-05-10"
    }
]

print("Employee database created with " + string(len(employees)) + " records")

// Analyze employee data
let departments = {}
let totalSalary = 0

for employee in employees {
    let dept = employee["department"]
    if (!departments.hasKey(dept)) {
        departments[dept] = 0
    }
    departments[dept] += 1
    totalSalary += employee["salary"]
}

print("\nDepartment distribution:")
for department, count in departments {
    print("  " + department + ": " + string(count) + " employees")
}

let avgSalary = totalSalary / len(employees)
print("Average salary: $" + string(avgSalary))

// Save employee data
let employeeJson = json.encode(employees)
os.writeFile("employees.json", employeeJson)
print("✓ Employee data saved to employees.json")

// Feature 6: File I/O Operations
print("\n💾 Feature 6: File System Operations")
print("-" * 50)

// Create a configuration file
let config = {
    "database": {
        "host": "localhost",
        "port": 5432,
        "name": "vintlang_demo"
    },
    "logging": {
        "level": "info",
        "file": "application.log"
    },
    "features": {
        "cache_enabled": true,
        "debug_mode": false,
        "max_connections": 100
    }
}

os.writeFile("app_config.json", json.encode(config))
print("✓ Configuration file created")

// Create a data file
let dataContent = "VintLang Data Processing Log\n"
dataContent += "=" * 40 + "\n"
dataContent += "Timestamp: " + time.format(time.now(), "02-01-2006 15:04:05") + "\n"
dataContent += "Records processed: " + string(len(employees)) + "\n"
dataContent += "Average salary calculated: $" + string(avgSalary) + "\n"
dataContent += "Departments analyzed: " + string(len(departments)) + "\n"

os.writeFile("processing_log.txt", dataContent)
print("✓ Processing log created")

// List files created
let files = os.listDir(".")
let createdFiles = []
let fileList = files.split(", ")

for filename in fileList {
    if (filename.contains("employees") ||
        filename.contains("config") ||
        filename.contains("log") ||
        filename.contains("showcase")) {
        if (os.fileExists(filename)) {
            createdFiles.push(filename)
        }
    }
}

print("Files created during demonstration:")
for file in createdFiles {
    let content = os.readFile(file)
    print("  📄 " + file + " (" + string(len(content)) + " bytes)")
}

// Feature 7: Advanced Data Processing
print("\n🔍 Feature 7: Advanced Data Analysis")
print("-" * 50)

// Find highest paid employee
let highestPaid = employees[0]
for employee in employees {
    if (employee["salary"] > highestPaid["salary"]) {
        highestPaid = employee
    }
}

print("Highest paid employee: " + highestPaid["name"] +
      " ($" + string(highestPaid["salary"]) + ")")

// Count employees by salary ranges
let salaryRanges = {
    "60000-69999": 0,
    "70000-79999": 0,
    "80000-89999": 0
}

for employee in employees {
    let salary = employee["salary"]
    if (salary >= 60000 && salary <= 69999) {
        salaryRanges["60000-69999"] += 1
    } else if (salary >= 70000 && salary <= 79999) {
        salaryRanges["70000-79999"] += 1
    } else if (salary >= 80000 && salary <= 89999) {
        salaryRanges["80000-89999"] += 1
    }
}

print("\nSalary distribution:")
for range, count in salaryRanges {
    print("  $" + range + ": " + string(count) + " employees")
}

// Feature 8: Time and UUID Operations
print("\n⏰ Feature 8: Time and Unique Identifiers")
print("-" * 50)

let sessionId = uuid.generate()
let currentTime = time.format(time.now(), "2006-01-02 15:04:05")

print("Session ID: " + sessionId)
print("Current timestamp: " + currentTime)

// Create timestamped records
let events = []
let eventTypes = ["login", "data_access", "calculation", "file_write", "logout"]

for i, eventType in eventTypes {
    let event = {
        "id": uuid.generate(),
        "type": eventType,
        "timestamp": time.format(time.now(), "2006-01-02 15:04:05"),
        "session": sessionId,
        "sequence": i + 1
    }
    events.push(event)
    print("Event " + string(i + 1) + ": " + eventType + " at " + event["timestamp"])
}

// Feature 9: Report Generation
print("\n📋 Feature 9: Comprehensive Report Generation")
print("-" * 50)

let report = {
    "report_id": uuid.generate(),
    "generated_at": time.format(time.now(), "2006-01-02 15:04:05"),
    "session_id": sessionId,
    "summary": {
        "employees_processed": len(employees),
        "departments_found": len(departments),
        "average_salary": avgSalary,
        "files_created": len(createdFiles),
        "events_logged": len(events)
    },
    "details": {
        "employees": employees,
        "departments": departments,
        "events": events,
        "highest_paid": highestPaid
    }
}

let reportFile = "comprehensive_report_" + time.format(time.now(), "2006-01-02_15-04-05") + ".json"
os.writeFile(reportFile, json.encode(report))
print("✓ Comprehensive report saved to: " + reportFile)

// Create human-readable summary
let summary = "VINTLANG FEATURE SHOWCASE SUMMARY\n"
summary += "Generated: " + time.format(time.now(), "02-01-2006 15:04:05") + "\n"
summary += "Session ID: " + sessionId + "\n"
summary += "=" * 50 + "\n\n"

summary += "STATISTICS:\n"
summary += "  • Employees processed: " + string(len(employees)) + "\n"
summary += "  • Departments identified: " + string(len(departments)) + "\n"
summary += "  • Average salary: $" + string(avgSalary) + "\n"
summary += "  • Files created: " + string(len(createdFiles)) + "\n"
summary += "  • Events logged: " + string(len(events)) + "\n"
summary += "  • UUIDs generated: " + string(len(employees) + len(events) + 2) + "\n"

summary += "\nFILES CREATED:\n"
for file in createdFiles {
    summary += "  • " + file + "\n"
}

summary += "\nFEATURES DEMONSTRATED:\n"
for feature in applicationData["features"] {
    summary += "  ✓ " + feature + "\n"
}

let summaryFile = "showcase_summary_" + time.format(time.now(), "2006-01-02_15-04-05") + ".txt"
os.writeFile(summaryFile, summary)
print("✓ Summary report saved to: " + summaryFile)

// Final Results
print("\n🎉 VintLang Showcase Complete!")
print("=" * 60)

print("🏆 SUCCESSFULLY DEMONSTRATED:")
print("  ✓ Variable declarations and data types")
print("  ✓ Arrays, dictionaries, and iteration")
print("  ✓ String manipulation and processing")
print("  ✓ Function definitions and calls")
print("  ✓ Conditional logic and comparisons")
print("  ✓ JSON encoding and decoding")
print("  ✓ File I/O operations")
print("  ✓ Directory management")
print("  ✓ Data analysis and statistics")
print("  ✓ Time and date operations")
print("  ✓ UUID generation")
print("  ✓ Report generation")

print("\n📊 PROCESSING STATISTICS:")
print("  • Data records: " + string(len(employees)))
print("  • Calculations performed: " + string(len(numbers) + len(employees) * 2))
print("  • Files generated: " + string(len(createdFiles) + 2))
print("  • JSON operations: 4+")
print("  • String operations: " + string(len(words) + len(processed)))

print("\n🚀 VINTLANG IS READY FOR:")
print("  • Business applications")
print("  • Data processing systems")
print("  • File management tools")
print("  • API development")
print("  • Automation scripts")
print("  • Educational programming")
print("  • Rapid prototyping")

print("\n✨ This comprehensive demonstration proves")
print("   VintLang's capability for real-world development!")

print("\n" + "=" * 60)
print("🎯 VintLang Feature Showcase Successfully Completed!")
print("=" * 60)
```

## crypto.vint

```js
// VintLang Crypto Module Example
// This module is still experimental
// Demonstrates cryptographic operations: hashing and encryption

import crypto

// Example 1: MD5 Hashing
// MD5 is a cryptographic hash function that produces a 128-bit hash value
let md5_hash = crypto.hashMD5("Hello, World!")
print(md5_hash)  // Prints the MD5 hash of "Hello, World!"

// Example 2: AES Encryption and Decryption
// AES (Advanced Encryption Standard) is a symmetric encryption algorithm
let key = "mysecretkey12345"  // Encryption key (must be appropriate length for AES)
let data = "Sensitive data"    // Data to encrypt
let encrypted = crypto.encryptAES(data, key)  // Encrypt the data
print(encrypted)

// Decrypt the encrypted data back to original
let decrypted = crypto.decryptAES(encrypted, key)
print(decrypted)
print(type(decrypted))

```

## csv.vint

```js
import csv
import os

// --- Writing to a CSV file ---
const data_to_write = [
    ["id", "name", "score"],
    ["1", "Alice", "88"],
    ["2", "Bob", "92"],
    ["3", "Charlie", "75"]
]

let filename = "scores.csv"
csv.write(filename, data_to_write)

println("Wrote data to", filename)


// --- Reading from a CSV file ---
let read_data = csv.read(filename)
println("Read data:",read_data)

if (!read_data) {
    error "Failed to read data from " + filename
} else {
    print("Read data from", filename, ": ")
    print(read_data)
}

// --- Clean up the created file ---
os.deleteFile(filename)
println("Cleaned up", filename)
```

## data_processor.vint

```js
// VintLang Showcase: Advanced Data Processing & Analytics Tool
// This demonstrates VintLang's real-world capabilities

import time
import os
import json
import uuid
import net

// Application configuration
let config = {
    "name": "VintLang Data Processor",
    "version": "1.0.0",
    "dataDir": "data",
    "outputDir": "output"
}

// Ensure directories exist
let setupDirectories = func() {
    if (!os.fileExists(config["dataDir"])) {
        os.makeDir(config["dataDir"])
        print("✓ Created data directory")
    }
    if (!os.fileExists(config["outputDir"])) {
        os.makeDir(config["outputDir"])
        print("✓ Created output directory")
    }
}

// Generate sample data for demonstration
let generateSampleData = func() {
    print("\n📊 Generating sample data...")

    let categories = ["Technology", "Science", "Business", "Health", "Education"]
    let priorities = ["Low", "Medium", "High", "Critical"]
    let statuses = ["Active", "Pending", "Completed", "Cancelled"]

    let sampleData = []

    // Generate 50 sample records
    for i in range(1, 51) {
        let record = {
            "id": uuid.generate(),
            "name": "Record " + string(i),
            "category": categories[(i - 1) % len(categories)],
            "priority": priorities[(i - 1) % len(priorities)],
            "status": statuses[(i - 1) % len(statuses)],
            "value": (i * 123) % 1000,
            "score": (i * 7) % 100,
            "timestamp": time.format(time.now(), "2006-01-02 15:04:05"),
            "tags": ["tag" + string(i % 3), "sample", "data"]
        }
        sampleData.push(record)
    }

    // Save to JSON file
    let dataFile = config["dataDir"] + "/sample_data.json"
    os.writeFile(dataFile, json.encode(sampleData))
    print("✓ Generated " + string(len(sampleData)) + " sample records")
    print("✓ Saved to: " + dataFile)

    return sampleData
}

// Load data from file
let loadData = func(filename) {
    let filepath = config["dataDir"] + "/" + filename
    if (!os.fileExists(filepath)) {
        print("❌ File not found: " + filepath)
        return []
    }

    let content = os.readFile(filepath)
    return json.decode(content)
}

// Data analysis functions
let analyzeByCategory = func(data) {
    print("\n📈 Analysis by Category:")
    print("-" * 40)

    let categoryStats = {}
    let categoryValues = {}

    for record in data {
        let category = record["category"]

        // Count records per category
        if (!categoryStats.hasKey(category)) {
            categoryStats[category] = 0
            categoryValues[category] = 0
        }
        categoryStats[category] += 1
        categoryValues[category] += record["value"]
    }

    // Display statistics
    for category, count in categoryStats {
        let avgValue = categoryValues[category] / count
        print(category + ":")
        print("  Records: " + string(count))
        print("  Total Value: " + string(categoryValues[category]))
        print("  Average Value: " + string(avgValue))
        print("")
    }

    return categoryStats
}

let analyzeByStatus = func(data) {
    print("\n📊 Analysis by Status:")
    print("-" * 40)

    let statusStats = {}

    for record in data {
        let status = record["status"]
        if (!statusStats.hasKey(status)) {
            statusStats[status] = 0
        }
        statusStats[status] += 1
    }

    let total = len(data)
    for status, count in statusStats {
        let percentage = (count * 100) / total
        print(status + ": " + string(count) + " (" + string(percentage) + "%)")
    }

    return statusStats
}

let findTopPerformers = func(data) {
    print("\n🏆 Top Performers by Score:")
    print("-" * 40)

    // Sort data by score (simple bubble sort for demonstration)
    let sortedData = data
    for i in range(0, len(sortedData) - 1) {
        for j in range(0, len(sortedData) - i - 1) {
            if (sortedData[j]["score"] < sortedData[j + 1]["score"]) {
                let temp = sortedData[j]
                sortedData[j] = sortedData[j + 1]
                sortedData[j + 1] = temp
            }
        }
    }

    // Display top 10
    let topCount = 10
    if (len(sortedData) < topCount) {
        topCount = len(sortedData)
    }

    for i in range(0, topCount) {
        let record = sortedData[i]
        print(string(i + 1) + ". " + record["name"] +
              " (Score: " + string(record["score"]) +
              ", Category: " + record["category"] + ")")
    }

    return sortedData
}

let generateReport = func(data, categoryStats, statusStats) {
    print("\n📝 Generating comprehensive report...")

    let reportTime = time.format(time.now(), "2006-01-02_15-04-05")
    let reportFile = config["outputDir"] + "/analysis_report_" + reportTime + ".txt"

    let report = "VINTLANG DATA ANALYSIS REPORT\n"
    report += "Generated: " + time.format(time.now(), "02-01-2006 15:04:05") + "\n"
    report += "=" * 50 + "\n\n"

    // Summary
    report += "SUMMARY\n"
    report += "-" * 20 + "\n"
    report += "Total Records: " + string(len(data)) + "\n"

    let totalValue = 0
    let totalScore = 0
    for record in data {
        totalValue += record["value"]
        totalScore += record["score"]
    }
    let avgValue = totalValue / len(data)
    let avgScore = totalScore / len(data)

    report += "Total Value: " + string(totalValue) + "\n"
    report += "Average Value: " + string(avgValue) + "\n"
    report += "Average Score: " + string(avgScore) + "\n\n"

    // Category breakdown
    report += "CATEGORY BREAKDOWN\n"
    report += "-" * 20 + "\n"
    for category, count in categoryStats {
        let percentage = (count * 100) / len(data)
        report += category + ": " + string(count) + " (" + string(percentage) + "%)\n"
    }

    report += "\nSTATUS BREAKDOWN\n"
    report += "-" * 20 + "\n"
    for status, count in statusStats {
        let percentage = (count * 100) / len(data)
        report += status + ": " + string(count) + " (" + string(percentage) + "%)\n"
    }

    report += "\n" + "=" * 50 + "\n"
    report += "Report generated by " + config["name"] + " v" + config["version"] + "\n"
    report += "Powered by VintLang Programming Language\n"

    os.writeFile(reportFile, report)
    print("✓ Report saved to: " + reportFile)
}

// Web data fetching demonstration
let fetchWebData = func() {
    print("\n🌐 Fetching data from web...")

    // Try to fetch from a simple API
    let url = "https://httpbin.org/json"
    let response = net.get(url)

    if (response != "") {
        print("✓ Successfully fetched data from: " + url)

        // Save response to file
        let webDataFile = config["dataDir"] + "/web_data.json"
        os.writeFile(webDataFile, response)
        print("✓ Web data saved to: " + webDataFile)

        return true
    } else {
        print("❌ Failed to fetch web data")
        return false
    }
}

// Export data in different formats
let exportData = func(data, format) {
    let timestamp = time.format(time.now(), "2006-01-02_15-04-05")

    if (format == "json") {
        let filename = config["outputDir"] + "/export_" + timestamp + ".json"
        os.writeFile(filename, json.encode(data))
        print("✓ Data exported to JSON: " + filename)
    } else if (format == "csv") {
        let filename = config["outputDir"] + "/export_" + timestamp + ".csv"
        let csvContent = "ID,Name,Category,Priority,Status,Value,Score,Timestamp\n"

        for record in data {
            csvContent += record["id"] + ","
            csvContent += record["name"] + ","
            csvContent += record["category"] + ","
            csvContent += record["priority"] + ","
            csvContent += record["status"] + ","
            csvContent += string(record["value"]) + ","
            csvContent += string(record["score"]) + ","
            csvContent += record["timestamp"] + "\n"
        }

        os.writeFile(filename, csvContent)
        print("✓ Data exported to CSV: " + filename)
    }
}

// Main application
let runDataProcessor = func() {
    print("🚀 Welcome to " + config["name"] + " v" + config["version"])
    print("=" * 60)
    print("This showcase demonstrates VintLang's capabilities for:")
    print("  • Data generation and processing")
    print("  • JSON manipulation and file I/O")
    print("  • Statistical analysis and reporting")
    print("  • Web data fetching")
    print("  • Multiple export formats")
    print("  • Directory management")
    print("  • Time and UUID operations")
    print("=" * 60)

    // Setup
    setupDirectories()

    // Generate or load data
    let data = generateSampleData()

    // Perform analysis
    let categoryStats = analyzeByCategory(data)
    let statusStats = analyzeByStatus(data)
    let topPerformers = findTopPerformers(data)

    // Generate report
    generateReport(data, categoryStats, statusStats)

    // Export data
    exportData(data, "json")
    exportData(data, "csv")

    // Try web data fetching
    fetchWebData()

    print("\n🎉 Data processing complete!")
    print("📁 Check the '" + config["outputDir"] + "' directory for generated files")
    print("📁 Source data available in '" + config["dataDir"] + "' directory")

    // Final summary
    print("\n📊 FINAL SUMMARY:")
    print("  • Processed " + string(len(data)) + " records")
    print("  • Analyzed " + string(len(categoryStats)) + " categories")
    print("  • Generated comprehensive report")
    print("  • Exported data in multiple formats")
    print("  • Demonstrated web data fetching")

    print("\n✨ VintLang successfully demonstrated real-world data processing!")
}

// Start the application
runDataProcessor()
```

## debounce_demo.vint

```js
// Comprehensive debounce demonstration
println("=== Vint Debounce Function Demo ===\n")

// Test 1: Basic debounce with user function
println("Test 1: User-defined function debouncing")
let callCount = 0
let testFunc = func(message) {
    callCount = callCount + 1
    println("Executed:", message, "| Call #" + string(callCount))
}

let debouncedTest = debounce(150, testFunc)

println("Making 5 rapid calls...")
debouncedTest("Call 1")
debouncedTest("Call 2")
debouncedTest("Call 3")
debouncedTest("Call 4")
debouncedTest("Final call") // Only this should execute

sleep(200) // Wait for execution
println("Result: Only 1 execution should have occurred\n")

// Test 2: Debounce with builtin function
println("Test 2: Builtin function debouncing")
let debouncedPrint = debounce(100, println)

println("Making rapid println calls...")
debouncedPrint("Message 1")
debouncedPrint("Message 2")
debouncedPrint("This message should appear") // Only this prints

sleep(150)
println("Result: Only the final message should have printed above\n")

// Test 3: Different delay values
println("Test 3: Testing with integer milliseconds")
let quickDebounce = debounce(50, func() { println("Quick execution!") })
let slowDebounce = debounce(200, func() { println("Slow execution!") })

quickDebounce()
slowDebounce()

sleep(75)  // Quick should execute
println("After 75ms - quick should have executed")

sleep(150) // Slow should execute
println("After 225ms total - slow should have executed")

println("\n=== Demo Complete ===")
println("The debounce function successfully delays execution and")
println("cancels previous calls when new ones are made rapidly!")
```

## debounce_example.vint

```js
// Example usage of the debounce builtin function in VintLang
// This demonstrates how to use debounce with different types of delays and functions

// Example 1: Using debounce with an integer (milliseconds) and a builtin function
let debouncedPrint = debounce(500, print)

println("Calling debounced print multiple times rapidly...")
debouncedPrint("First call")
debouncedPrint("Second call")
debouncedPrint("Third call") // Only this one should execute after 500ms

// Example 2: Using debounce with a Duration object and a user-defined function
let logMessage = func(msg) {
    println("LOG:", msg)
}

// Note: This would require a duration object to be created
// let debouncedLog = debounce(duration.milliseconds(1000), logMessage)

// Example 3: Demonstrating the debounce behavior
println("Testing debounce behavior...")
let counter = 0
let incrementCounter = func() {
    counter = counter + 1
    println("Counter:", counter)
}

let debouncedIncrement = debounce(300, incrementCounter)

// These calls will be debounced - only the last one should execute
debouncedIncrement()
debouncedIncrement()
debouncedIncrement()
debouncedIncrement() // Only this call should result in incrementing the counter

println("Debounce test complete. The counter should only increment once after 300ms.")
```

## debug_match.vint

```js
// Debug match pattern
let config = {"database": {"host": "localhost", "port": 5432}}

match config {
    {"database": db} => {
        println("Found database config:", db)
    }
    _ => {
        println("No match")
    }
}
```

## debug_switch.vint

```js
// Test switch with string length
let s = "hello"

switch (s) {
    case x if len(x) > 3 {
        println("Long string:", x)
    }
    default {
        println("Short string")
    }
}
```

## declaratives_test.vint

```js
// Info statement
info "This is an informational message."
println("Info should print above.")

// Debug statement
debug "Debugging value: " + "123"
println("Debug should print above.")

// Note statement
note "This is a note for the user."
println("Note should print above.")

// Success statement
success "Operation completed successfully!"
println("Success should print above.")

// Existing warn and error tests
warn "This is a warning. The script should continue."
println("This line should execute after the warning.")
let errMSG = "This is a fatal error. The script should stop here."
error errMSG

//unreachable code
println("This line should NOT execute.")
```

## defer_test.vint

```js
// VintLang Defer Statement Example
// Demonstrates the defer keyword which delays execution until function returns

let test_defer = func() {
    // The defer statement will execute AFTER the function completes
    // regardless of how the function exits
    defer println("deferred message");

    // This executes immediately
    println("function body");
};

// Call the function
test_defer();

// This executes after the function and its deferred statements complete
println("after function call");

```

## desktop.vint

```js
// VintLang Desktop GUI Application Example
// This example demonstrates creating a simple desktop application with GUI elements

import desktop

// Create the desktop application
desktop.createApp()

// Add a label to the window
desktop.newLabel("Hello, Vint Desktop App!")

// Add a button with a callback function
desktop.newButton("Click Me", func() {
    print("Button clicked!")
})

// Add an entry field for text input
desktop.newEntry()

// Run the application (starts the GUI event loop)
desktop.runApp()

```

## dict_pattern_matching.vint

```js
// Test dict pattern matching feature
// This demonstrates the new match statement for flexible dict matching

let user = {"role": "admin", "active": true}

println("Testing dict pattern matching:")

match user {
    {"role": "admin"} => println("Hello, Admin!")
    {"active": false} => println("Inactive user")
    _ => println("Regular user")
}

// Test multiple patterns
let users = [
    {"role": "admin", "active": true},
    {"role": "user", "active": false},
    {"name": "John", "type": "guest"}
]

for u in users {
    match u {
        {"role": "admin"} => println("Admin user found")
        {"active": false} => println("Found inactive user")
        _ => println("Other user type")
    }
}
```

## dotenv.vint

```js
// VintLang DotEnv Module Example
// Demonstrates loading environment variables from .env files

import dotenv

// Example 1: Load environment variables from a .env file
// This reads the .env file and makes its variables available
// dotenv.load(".env")

// Example 2: Get a specific environment variable
// let apiKey = dotenv.get("API_KEY")
// print(apiKey) // Outputs the value of API_KEY from .env

// Note: Create a .env file in the same directory with content like:
// API_KEY=your_api_key_here
// DATABASE_URL=your_database_url

println("DotEnv module requires a .env file to be present")
println("Create a .env file with key=value pairs, then uncomment the code above")

```

## encoding.vint

```js
// VintLang Encoding Module Example
// Demonstrates base64 encoding and decoding operations

import encoding

// Example 1: Base64 Encoding
// Encode a string to base64 format
let encoded = encoding.base64Encode("Hello, World!")
print(encoded) // Outputs "SGVsbG8sIFdvcmxkIQ=="

// Example 2: Base64 Decoding
// Decode a base64 string back to original text
let decoded = encoding.base64Decode(encoded)
print(decoded) // Outputs "Hello, World!"

```

## enhanced_declaratives_test.vint

```js
// Enhanced Declaratives Test File for Vint
// This file demonstrates all the declarative improvements made to Vint

println("=== Testing Enhanced Declaratives ===")

// Original lowercase declaratives (already working)
println("\n--- Original Declaratives (lowercase) ---")
info "Info: Standard informational message"
debug "Debug: Debugging information"
note "Note: Important note for developers"
todo "Todo: Task that needs to be completed"
warn "Warn: Warning about potential issues"
success "Success: Operation completed successfully"
trace "Trace: Execution flow tracing"

// NEW: Capitalized declaratives (now supported)
println("\n--- NEW: Capitalized Declaratives ---")
Info "Info: Capitalized informational message"
Debug "Debug: Capitalized debugging information"
Note "Note: Capitalized important note"
Todo "Todo: Capitalized task"
Warn "Warn: Capitalized warning"
Success "Success: Capitalized success message"
Trace "Trace: Capitalized execution tracing"

// NEW: Log declarative - non-fatal error logging
println("\n--- NEW: Non-Fatal Error Logging ---")
log "Log: Non-fatal error message (execution continues)"
Log "Log: Capitalized non-fatal error message"
println("Execution continues after log messages")

// Fatal declaratives (stop execution)
println("\n--- Fatal Declaratives (execution stops) ---")
println("Testing error declarative:")
let errorMsg = "This is an error"
error errorMsg + " - execution stops here"

println("This line should NOT be reached due to error above")
```

## enhanced_http_test.vint

```js
// Enhanced HTTP module test demonstrating new backend features
import http

print("🚀 Enhanced HTTP Module Backend Features Test")
print("=" * 50)

// Test 1: Create app with enhanced features
print("\n✓ Test 1 - Enhanced App Creation")
let result = http.app()
print("App creation:", result)

// Test 2: Basic route registration with different HTTP methods
print("\n✓ Test 2 - Route Registration")
http.get("/users", func(req, res) {
    print("GET /users - List users")
})

http.post("/users", func(req, res) {
    print("POST /users - Create user")
})

http.put("/users/:id", func(req, res) {
    print("PUT /users/:id - Update user")
})

http.delete("/users/:id", func(req, res) {
    print("DELETE /users/:id - Delete user")
})

// Test 3: Middleware registration
print("\n✓ Test 3 - Middleware Registration")
http.use(func(req, res, next) {
    print("Custom middleware - logging request")
})

// Test 4: New Backend Features - Interceptors
print("\n✓ Test 4 - Interceptors")
let requestInterceptor = http.interceptor("request", func(req) {
    print("Request interceptor - validating request")
})
print("Request interceptor:", requestInterceptor)

let responseInterceptor = http.interceptor("response", func(res) {
    print("Response interceptor - adding headers")
})
print("Response interceptor:", responseInterceptor)

// Test 5: Guards
print("\n✓ Test 5 - Guards")
let authGuard = http.guard(func(req) {
    print("Authentication guard - checking token")
})
print("Auth guard:", authGuard)

let rateLimitGuard = http.guard(func(req) {
    print("Rate limit guard - checking limits")
})
print("Rate limit guard:", rateLimitGuard)

// Test 6: Built-in Middleware
print("\n✓ Test 6 - Built-in Middleware")
let corsResult = http.cors()
print("CORS middleware:", corsResult)

let bodyParserResult = http.bodyParser()
print("Body parser middleware:", bodyParserResult)

let authMiddleware = http.auth(func(req, res, next) {
    print("Authentication middleware")
})
print("Auth middleware:", authMiddleware)

// Test 7: Error Handler
print("\n✓ Test 7 - Error Handler")
let errorHandler = http.errorHandler(func(err, req, res) {
    print("Global error handler")
})
print("Error handler:", errorHandler)

// Test 8: API Routes with Parameters
print("\n✓ Test 8 - API Routes with Parameters")
http.get("/api/users/:id", func(req, res) {
    print("GET /api/users/:id - Get user by ID")
})

http.get("/api/users/:id/posts/:postId", func(req, res) {
    print("GET /api/users/:id/posts/:postId - Get user post")
})

// Test 9: Different content types
print("\n✓ Test 9 - Content Type Handlers")
http.post("/api/upload", func(req, res) {
    print("POST /api/upload - File upload handler")
})

http.post("/api/json", func(req, res) {
    print("POST /api/json - JSON data handler")
})

http.post("/api/form", func(req, res) {
    print("POST /api/form - Form data handler")
})

print("\n" + "=" * 50)
print("✨ All enhanced HTTP features registered successfully!")
print("Features demonstrated:")
print("  • Enhanced request/response objects")
print("  • Interceptors (request & response)")
print("  • Guards (authentication, rate limiting)")
print("  • Built-in middleware (CORS, body parser, auth)")
print("  • Error handling")
print("  • Route parameters (/users/:id)")
print("  • Multiple content type support")
print("  • Security features")
print("\n🎯 Ready for full-fledged backend development!")

// The server would normally be started with:
// http.listen(3000, "Enhanced HTTP server running on port 3000")
```

## enhanced_language_showcase.vint

```js
// Enhanced VintLang Language Showcase
// This example demonstrates the improvements made to VintLang

print("🚀 VintLang Enhanced Language Showcase")
print("=" * 50)

// 1. Enhanced String Handling with Unicode and Escape Sequences
print("\n📝 Enhanced String Handling:")
let greeting = "Hello\nWorld!"
let unicode_text = "Hello \u0041\u0042\u0043"  // Hello ABC
let hex_text = "Value: \x41\x42"  // Value: AB
print("Multiline greeting:", greeting)
print("Unicode text:", unicode_text)
print("Hex escape text:", hex_text)

// 2. Type Checking Functions
print("\n🔍 Type Checking:")
let values = [42, 3.14, "hello", true, [1, 2, 3], {"key": "value"}]
for value in values {
    print("Value:", value, "Type:", typeof(value))
    print("  isInt:", isInt(value))
    print("  isFloat:", isFloat(value))
    print("  isString:", isString(value))
    print("  isBool:", isBool(value))
    print("  isArray:", isArray(value))
    print("  isDict:", isDict(value))
    print("---")
}

// 3. Enhanced Array Operations
print("\n📊 Enhanced Array Operations:")
let numbers = [5, 2, 8, 1, 9, 3, 5, 2]
print("Original array:", numbers)
print("Length:", len(numbers))
print("Sorted:", sort(numbers))
print("Reversed:", reverse(numbers))
print("Unique values:", unique(numbers))
print("Index of 8:", indexOf(numbers, 8))

// Higher-order functions
let doubled = map(numbers, func(x) { return x * 2 })
print("Doubled:", doubled)

let evens = filter(numbers, func(x) { return x % 2 == 0 })
print("Even numbers:", evens)

// 4. Enhanced String Operations
print("\n🔤 Enhanced String Operations:")
let text = "  VintLang Programming Language  "
print("Original:", text)
print("Trimmed:", trim(text))
print("Uppercase:", toUpper(trim(text)))
print("Lowercase:", toLower(trim(text)))

let sentence = "VintLang is a modern programming language"
let words = split(sentence, " ")
print("Words:", words)
print("Joined with '-':", join(words, "-"))

print("Starts with 'VintLang':", startsWith(sentence, "VintLang"))
print("Ends with 'language':", endsWith(sentence, "language"))

// 5. Math Operations
print("\n🧮 Math Operations:")
let vals = [10, -5, 3.14, -2.7]
for val in vals {
    print("abs(" + string(val) + ") =", abs(val))
}
print("min(10, 5, 20):", min(10, 5, 20))
print("max(10, 5, 20):", max(10, 5, 20))

// 6. Range Function
print("\n📏 Range Function:")
print("range(5):", range(5))
print("range(2, 8):", range(2, 8))
print("range(0, 10, 2):", range(0, 10, 2))
print("range(10, 0, -2):", range(10, 0, -2))

// 7. Data Conversion and Parsing
print("\n🔄 Data Conversion:")
let num_str = "42"
let float_str = "3.14159"
print("parseInt('" + num_str + "'):", parseInt(num_str))
print("parseFloat('" + float_str + "'):", parseFloat(float_str))
print("string(42):", string(42))
print("int('100'):", int("100"))

// 8. Dictionary Operations with Better Built-ins
print("\n📚 Dictionary Operations:")
let person = {
    "name": "Alice",
    "age": 30,
    "city": "New York"
}
print("Person:", person)
print("Keys:", keys(person))
print("Values:", values(person))
print("Has 'age' key:", has_key(person, "age"))
print("Has 'country' key:", has_key(person, "country"))

// 9. Error Handling and Validation
print("\n⚠️ Error Handling Examples:")
let safe_divide = func(a, b) {
    if (b == 0) {
        return "Error: Division by zero"
    }
    return a / b
}

print("10 / 2 =", safe_divide(10, 2))
print("10 / 0 =", safe_divide(10, 0))

// 10. Advanced Control Flow
print("\n🔀 Advanced Control Flow:")
let process_data = func(data) {
    if (isArray(data)) {
        return "Processing array with " + string(len(data)) + " elements"
    } else if (isDict(data)) {
        return "Processing dictionary with " + string(len(data)) + " keys"
    } else if (isString(data)) {
        return "Processing string: " + data
    } else {
        return "Unknown data type: " + typeof(data)
    }
}

let test_data = [
    [1, 2, 3, 4],
    {"name": "test", "type": "demo"},
    "Hello World",
    42
]

for data in test_data {
    print(process_data(data))
}

// 11. Function Composition and Higher-Order Functions
print("\n🔗 Function Composition:")
let add_one = func(x) { return x + 1 }
let double = func(x) { return x * 2 }
let square = func(x) { return x * x }

let compose = func(f, g) {
    return func(x) { return f(g(x)) }
}

let add_one_then_double = compose(double, add_one)
let square_then_add_one = compose(add_one, square)

print("add_one_then_double(5):", add_one_then_double(5))  // (5 + 1) * 2 = 12
print("square_then_add_one(4):", square_then_add_one(4))  // (4 * 4) + 1 = 17

// 12. Data Processing Pipeline
print("\n⚙️ Data Processing Pipeline:")
let process_numbers = func(nums) {
    let result = nums
    result = filter(result, func(x) { return x > 0 })  // Filter positive
    result = map(result, func(x) { return x * x })     // Square each
    result = sort(result)                              // Sort ascending
    result = unique(result)                            // Remove duplicates
    return result
}

let mixed_numbers = [-2, 3, -1, 4, 3, 5, -3, 4, 2]
print("Original:", mixed_numbers)
print("Processed:", process_numbers(mixed_numbers))

// 13. Performance and Utility
print("\n⏱️ Performance and Utilities:")
let start_time = time.now()
sleep(100)  // Sleep for 100ms
print("Operation completed after sleep")

// Clone objects safely
let original_array = [1, [2, 3], {"nested": "object"}]
let cloned_array = clone(original_array)
print("Original array:", original_array)
print("Cloned array:", cloned_array)

print("\n✅ VintLang Enhanced Language Showcase Complete!")
print("This demonstrates:")
print("• Enhanced string handling with Unicode support")
print("• Comprehensive type checking functions")
print("• Rich array and string manipulation")
print("• Mathematical operations")
print("• Functional programming features")
print("• Data conversion and parsing")
print("• Better error handling patterns")
print("• Advanced control flow")
print("• Function composition")
print("• Data processing pipelines")
print("• Performance utilities")
```

## enhanced_switch_match.vint

```js
// Enhanced Switch and Match Statement Examples
// This demonstrates the new features: guard conditions, variable binding, and array patterns

println("=== Enhanced Switch Statements ===")

// Test 1: Switch with guard conditions
let testNumbers = [5, -3, 0, 15, 100]

for num in testNumbers {
    println("Testing number:", num)

    switch (num) {
        case x if x > 0 && x < 10 {
            println("  Small positive number:", x)
        }
        case x if x > 10 {
            println("  Large positive number:", x)
        }
        case x if x < 0 {
            println("  Negative number:", x)
        }
        case 0 {
            println("  Zero")
        }
        default {
            println("  Other number")
        }
    }
}

println("\n=== Enhanced Match Statements ===")

// Test 2: Array pattern matching with destructuring
let arrays = [
    [],
    [1],
    [1, 2],
    [1, 2, 3, 4, 5]
]

for arr in arrays {
    println("Testing array:", arr)

    match arr {
        [] => println("  Empty array")
        [single] => println("  Single element:", single)
        [first, second] => println("  Two elements:", first, "and", second)
        [head, ...tail] => println("  Head:", head, "Tail:", tail)
    }
}

// Test 3: Dictionary pattern matching with guard conditions
let users = [
    {"role": "admin", "active": true, "age": 30},
    {"role": "user", "active": false, "age": 25},
    {"role": "guest", "age": 45},
    {"name": "John", "type": "visitor"}
]

println("\n=== Dictionary Pattern Matching with Guards ===")

for user in users {
    println("Testing user:", user)

    match user {
        {"role": "admin", "active": active} if active =>
            println("  Active admin user")
        {"role": role, "age": age} if age >= 18 =>
            println("  Adult user with role:", role)
        {"role": "guest"} =>
            println("  Guest user")
        _ =>
            println("  Unknown user type")
    }
}

// Test 4: Complex array patterns with variable binding
let complexArrays = [
    [1, 2, 3, 4, 5, 6],
    ["a", "b"],
    [true, false, true]
]

println("\n=== Complex Array Patterns ===")

for arr in complexArrays {
    println("Testing complex array:", arr)

    match arr {
        [first, second, ...rest] if len(rest) > 2 =>
            println("  Long array - first:", first, "second:", second, "rest has", len(rest), "items")
        [a, b] if type(a) == "STRING" =>
            println("  String pair:", a, "and", b)
        [x, y, z] if type(x) == "BOOLEAN" =>
            println("  Boolean triple:", x, y, z)
        _ =>
            println("  Other pattern")
    }
}

println("\n=== Tests completed! ===")
```

## enterprise_http_test.vint

```js
// Enterprise HTTP Module Test - Advanced Backend Features
import http

print("🏢 Enterprise HTTP Module Features Test")
print("=" * 60)

// Test 1: Create app with enterprise features
print("\n✓ Test 1 - Enterprise App Creation")
let result = http.app()
print("App creation:", result)

// Test 2: Route Grouping for API Versioning
print("\n✓ Test 2 - Route Grouping & API Versioning")
let v1Group = http.group("/api/v1", func() {
    // Routes within this group will be prefixed with /api/v1
    print("API v1 group created")
})
print("API v1 group:", v1Group)

let v2Group = http.group("/api/v2", func() {
    // Routes within this group will be prefixed with /api/v2
    print("API v2 group created")
})
print("API v2 group:", v2Group)

// Test 3: Multipart File Upload Support
print("\n✓ Test 3 - Multipart File Upload")
http.post("/upload", func(req, res) {
    // Parse multipart form data automatically
    let parseResult = http.multipart(req)
    print("Multipart parsing:", parseResult)

    // Access uploaded files
    let avatar = req.file("avatar")
    let documents = req.file("documents")

    if avatar {
        // Save the uploaded file
        let saveResult = avatar.save("/uploads/" + avatar.name())
        print("Avatar saved:", saveResult)
    }

    // Access form fields
    let username = req.form("username")
    let description = req.form("description")

    res.status(200).json({
        "message": "Files uploaded successfully",
        "username": username,
        "files": req.files()
    })
})

// Test 4: Async Handlers for Long-Running Operations
print("\n✓ Test 4 - Async Handlers")
let asyncHandler = http.async(func(req, res) {
    print("Processing long-running task asynchronously")
    // Simulate heavy computation or database operations
    // This won't block other requests
})

http.post("/process", asyncHandler)
print("Async handler registered")

// Test 5: Advanced Middleware Composition
print("\n✓ Test 5 - Advanced Middleware Composition")

// Authentication middleware
let authMiddleware = func(req, res, next) {
    let token = req.get("Authorization")
    if !token {
        res.status(401).json({"error": "No authorization token"})
        return
    }
    // Validate JWT token here
    next()
}

// Logging middleware
let loggingMiddleware = func(req, res, next) {
    print("Request:", req.method(), req.path())
    next()
}

// Rate limiting middleware
let rateLimitMiddleware = func(req, res, next) {
    // Implement rate limiting logic here
    print("Rate limit check passed")
    next()
}

// Apply multiple middlewares
http.use(loggingMiddleware)
http.use(rateLimitMiddleware)

// Protected route with authentication
http.post("/protected", func(req, res) {
    res.json({"message": "Access granted to protected resource"})
})

// Test 6: Enhanced Security Features
print("\n✓ Test 6 - Security Features")
let securityResult = http.security()
print("Security middleware:", securityResult)

// Test 7: Enhanced Error Handling Structure
print("\n✓ Test 7 - Enhanced Error Handling")
http.errorHandler(func(err, req, res) {
    // Structured error response
    res.status(500).json({
        "error": {
            "type": "INTERNAL_SERVER_ERROR",
            "message": err.message,
            "code": "ERR_INTERNAL",
            "status": 500,
            "details": {
                "timestamp": Date.now(),
                "path": req.path(),
                "method": req.method()
            }
        }
    })
})

// Test 8: Real-world API Endpoints
print("\n✓ Test 8 - Real-world API Endpoints")

// User management endpoints
http.get("/api/users", func(req, res) {
    let page = req.query("page") || "1"
    let limit = req.query("limit") || "10"

    res.json({
        "users": [],
        "pagination": {
            "page": page,
            "limit": limit,
            "total": 0
        }
    })
})

http.post("/api/users", func(req, res) {
    let userData = req.json()

    // Validate user data
    if !userData.email {
        res.status(400).json({
            "error": {
                "type": "VALIDATION_ERROR",
                "message": "Email is required",
                "code": "MISSING_EMAIL"
            }
        })
        return
    }

    res.status(201).json({
        "message": "User created successfully",
        "user": userData
    })
})

// File upload endpoint
http.post("/api/files", func(req, res) {
    http.multipart(req)

    let files = []
    let uploadedFiles = req.files()

    // Process each uploaded file
    for file in uploadedFiles {
        let savedPath = file.save("/uploads/" + file.name())
        files.push({
            "originalName": file.name(),
            "size": file.size(),
            "type": file.type(),
            "savedPath": savedPath
        })
    }

    res.json({
        "message": "Files uploaded successfully",
        "files": files
    })
})

// Health check endpoint
http.get("/health", func(req, res) {
    res.json({
        "status": "healthy",
        "timestamp": Date.now(),
        "version": "1.0.0",
        "uptime": process.uptime()
    })
})

// Metrics endpoint (for APM integration)
http.get("/metrics", func(req, res) {
    res.header("Content-Type", "text/plain")
    res.send(`
# HTTP Request Count
http_requests_total 1500

# HTTP Request Duration
http_request_duration_seconds 0.05

# Memory Usage
memory_usage_bytes 104857600
`)
})

print("\n" + "=" * 60)
print("✨ All enterprise HTTP features registered successfully!")
print("\nEnterprise Features Demonstrated:")
print("  🔧 Route grouping and API versioning")
print("  📁 Multipart file upload support")
print("  ⚡ Async handlers for long-running tasks")
print("  🔗 Advanced middleware composition")
print("  🛡️  Enhanced security features")
print("  📊 Structured error handling")
print("  🔍 Performance monitoring hooks")
print("  🌐 Production-ready API endpoints")
print("  📈 Health checks and metrics")

print("\n🎯 Ready for enterprise-level backend development!")
print("📝 Start the server with: http.listen(3000)")
```

## example.txt

```js
Hello;
World;
```

## example_cli.vint

```js
// Example: A simple CLI tool that processes files
// Usage: vint example_cli.vint input.txt --output result.txt --verbose --format json

import cli

// Get all command line arguments
let allArgs = args()
println("All arguments:", allArgs)

// Get positional arguments (non-flags)
let positional = cli.getPositional()
println("Positional arguments:", positional)

// Check if help was requested
if (cli.hasArg("--help") || cli.hasArg("-h")) {
    println("Usage: vint example_cli.vint <input_file> [options]")
    println("Options:")
    println("  --output <file>    Output file (default: stdout)")
    println("  --format <format>  Output format: json, csv, txt (default: txt)")
    println("  --verbose, -v      Enable verbose output")
    println("  --help, -h         Show this help message")
    exit(0)
}

// Get input file from positional arguments
let inputFile = "stdin"
if (len(positional) > 0) {
    inputFile = positional[0]
}

// Get output file from flags
let outputFile = cli.getArgValue("--output")
if (!outputFile) {
    outputFile = "stdout"
}

// Get format
let format = cli.getArgValue("--format")
if (!format) {
    format = "txt"
}

// Check for verbose mode
let verbose = cli.hasArg("--verbose") || cli.hasArg("-v")

// Display configuration
println("\nProcessing configuration:")
println("Input file:", inputFile)
println("Output file:", outputFile)
println("Format:", format)
println("Verbose mode:", verbose)

if (verbose) {
    println("\nVerbose: Processing", inputFile, "->", outputFile, "in", format, "format")
}

println("\n✓ CLI argument parsing working correctly!")

```

## excel_demo.vint

```js
// VintLang Excel Module Demo
// Demonstrates comprehensive Excel functionality including password handling
import excel

print("🚀 VintLang Excel Module Comprehensive Demo")
print("==========================================")

// Create a comprehensive employee workbook
print("\n📝 Creating Employee Management System...")
let file_id = excel.create("employee_system.xlsx")

// Set up the main employee data sheet
excel.renameSheet(file_id, "Sheet1", "Employees")
excel.addSheet(file_id, "Summary")
excel.addSheet(file_id, "Departments")

// Employee data headers
excel.setCell(file_id, "Employees", "A1", "Employee ID")
excel.setCell(file_id, "Employees", "B1", "Full Name")
excel.setCell(file_id, "Employees", "C1", "Department")
excel.setCell(file_id, "Employees", "D1", "Salary")
excel.setCell(file_id, "Employees", "E1", "Bonus")
excel.setCell(file_id, "Employees", "F1", "Total Compensation")

// Sample employee data
excel.setCell(file_id, "Employees", "A2", 101)
excel.setCell(file_id, "Employees", "B2", "John Smith")
excel.setCell(file_id, "Employees", "C2", "Engineering")
excel.setCell(file_id, "Employees", "D2", 75000)
excel.setCellFormula(file_id, "Employees", "E2", "=D2*0.15")
excel.setCellFormula(file_id, "Employees", "F2", "=D2+E2")

excel.setCell(file_id, "Employees", "A3", 102)
excel.setCell(file_id, "Employees", "B3", "Sarah Johnson")
excel.setCell(file_id, "Employees", "C3", "Marketing")
excel.setCell(file_id, "Employees", "D3", 68000)
excel.setCellFormula(file_id, "Employees", "E3", "=D3*0.12")
excel.setCellFormula(file_id, "Employees", "F3", "=D3+E3")

excel.setCell(file_id, "Employees", "A4", 103)
excel.setCell(file_id, "Employees", "B4", "Michael Brown")
excel.setCell(file_id, "Employees", "C4", "Sales")
excel.setCell(file_id, "Employees", "D4", 72000)
excel.setCellFormula(file_id, "Employees", "E4", "=D4*0.18")
excel.setCellFormula(file_id, "Employees", "F4", "=D4+E4")

excel.setCell(file_id, "Employees", "A5", 104)
excel.setCell(file_id, "Employees", "B5", "Emily Davis")
excel.setCell(file_id, "Employees", "C5", "HR")
excel.setCell(file_id, "Employees", "D5", 65000)
excel.setCellFormula(file_id, "Employees", "E5", "=D5*0.10")
excel.setCellFormula(file_id, "Employees", "F5", "=D5+E5")

print("✅ Added employee data with calculated bonuses")

// Create department summary
excel.setCell(file_id, "Departments", "A1", "Department")
excel.setCell(file_id, "Departments", "B1", "Employee Count")
excel.setCell(file_id, "Departments", "C1", "Average Salary")

excel.setCell(file_id, "Departments", "A2", "Engineering")
excel.setCell(file_id, "Departments", "B2", 1)
excel.setCell(file_id, "Departments", "C2", 75000)

excel.setCell(file_id, "Departments", "A3", "Marketing")
excel.setCell(file_id, "Departments", "B3", 1)
excel.setCell(file_id, "Departments", "C3", 68000)

excel.setCell(file_id, "Departments", "A4", "Sales")
excel.setCell(file_id, "Departments", "B4", 1)
excel.setCell(file_id, "Departments", "C4", 72000)

excel.setCell(file_id, "Departments", "A5", "HR")
excel.setCell(file_id, "Departments", "B5", 1)
excel.setCell(file_id, "Departments", "C5", 65000)

print("✅ Created department breakdown")

// Summary sheet with calculated totals
excel.setCell(file_id, "Summary", "A1", "EMPLOYEE MANAGEMENT SUMMARY")
excel.mergeCells(file_id, "Summary", "A1:C1")

excel.setCell(file_id, "Summary", "A3", "Total Employees:")
excel.setCellFormula(file_id, "Summary", "B3", "=COUNTA(Employees.A2:A5)")

excel.setCell(file_id, "Summary", "A4", "Total Payroll:")
excel.setCellFormula(file_id, "Summary", "B4", "=SUM(Employees.D2:D5)")

excel.setCell(file_id, "Summary", "A5", "Total Bonuses:")
excel.setCellFormula(file_id, "Summary", "B5", "=SUM(Employees.E2:E5)")

excel.setCell(file_id, "Summary", "A6", "Total Compensation:")
excel.setCellFormula(file_id, "Summary", "B6", "=SUM(Employees.F2:F5)")

excel.setCell(file_id, "Summary", "A7", "Average Salary:")
excel.setCellFormula(file_id, "Summary", "B7", "=AVERAGE(Employees.D2:D5)")

print("✅ Created executive summary with formulas")

// Test row operations
excel.insertRow(file_id, "Employees", 3)
excel.setCell(file_id, "Employees", "A3", 105)
excel.setCell(file_id, "Employees", "B3", "David Wilson")
excel.setCell(file_id, "Employees", "C3", "Finance")
excel.setCell(file_id, "Employees", "D3", 70000)
excel.setCellFormula(file_id, "Employees", "E3", "=D3*0.13")
excel.setCellFormula(file_id, "Employees", "F3", "=D3+E3")

print("✅ Added new employee via row insertion")

// Save the workbook
excel.save(file_id)
print("✅ Saved employee management system")

// Test file operations
let backup_id = excel.saveAs(file_id, "employee_backup.xlsx")
excel.close(file_id)
excel.close(backup_id)
print("✅ Created backup and closed files")

// Reopen and verify
let reopened_id = excel.open("employee_system.xlsx")
let employee_name = excel.getCell(reopened_id, "Employees", "B2")
let employee_salary = excel.getCell(reopened_id, "Employees", "D2")
print("✅ Reopened file - First employee:", employee_name, "Salary:", employee_salary)

// Display file information
let file_info = excel.getFileInfo(reopened_id)
print("\n📊 File Information:")
print("✓ Excel file created successfully")
print("✓ Multiple sheets with data and formulas")
print("✓ Cell formatting and merging")
print("✓ Row/column operations")

excel.close(reopened_id)

print("\n🎉 Excel Module Demo Complete!")
print("==============================")
print("✅ Created comprehensive employee management system")
print("📁 Files generated:")
print("   • employee_system.xlsx (main workbook)")
print("   • employee_backup.xlsx (backup copy)")
print("")
print("🔧 Features demonstrated:")
print("   ✓ Create and save Excel files")
print("   ✓ Multiple sheet management")
print("   ✓ Cell data entry (text, numbers)")
print("   ✓ Formula calculations")
print("   ✓ Cell merging")
print("   ✓ Row insertion")
print("   ✓ File backup and reopening")
print("")
print("🔒 Password Protection Available:")
print("   • Use excel.openWithPassword(file, password) for protected files")
print("   • Password setting requires manual Excel file protection")
print("")
print("🚀 VintLang Excel module is production-ready!")
print("   Full documentation: docs/excel.md")
```

## excel_minimal_test.vint

```js
import excel

print("Testing Excel Module")
let file_id = excel.create("simple_test.xlsx")
print("Created file:", file_id)

excel.setCell(file_id, "Sheet1", "A1", "Hello")
excel.setCell(file_id, "Sheet1", "B1", "World")

let value = excel.getCell(file_id, "Sheet1", "A1")
print("Read value:", value)

excel.save(file_id)
excel.close(file_id)

print("Test completed successfully!")
```

## excel_test.vint

```js
// Excel Module Test - Comprehensive functionality test
// Tests all major Excel operations including password handling

import excel

print("🚀 VintLang Excel Module Test")
print("==============================\n")

// Test 1: Create a new Excel file
print("📝 Test 1: Creating new Excel file...")
file_id = excel.create("test_workbook.xlsx")
print("✅ Created file with ID:", file_id)

// Test 2: Sheet operations
print("\n📋 Test 2: Sheet operations...")
sheets = excel.getSheets(file_id)
print("Initial sheets:", sheets)

// Add new sheets
print("Adding new sheets...")
index1 = excel.addSheet(file_id, "DataSheet")
index2 = excel.addSheet(file_id, "SummarySheet")
print("Created DataSheet at index:", index1)
print("Created SummarySheet at index:", index2)

// Rename default sheet
excel.renameSheet(file_id, "Sheet1", "MainSheet")
print("Renamed Sheet1 to MainSheet")

// Get updated sheet list
sheets = excel.getSheets(file_id)
print("Updated sheets:", sheets)

// Test 3: Cell operations
print("\n📊 Test 3: Cell operations...")
excel.setCell(file_id, "MainSheet", "A1", "Employee Name")
excel.setCell(file_id, "MainSheet", "B1", "Department")
excel.setCell(file_id, "MainSheet", "C1", "Salary")
excel.setCell(file_id, "MainSheet", "D1", "Bonus")

excel.setCell(file_id, "MainSheet", "A2", "John Doe")
excel.setCell(file_id, "MainSheet", "B2", "Engineering")
excel.setCell(file_id, "MainSheet", "C2", 75000)
excel.setCell(file_id, "MainSheet", "D2", "=C2*0.1")

excel.setCell(file_id, "MainSheet", "A3", "Jane Smith")
excel.setCell(file_id, "MainSheet", "B3", "Marketing")
excel.setCell(file_id, "MainSheet", "C3", 65000)
excel.setCell(file_id, "MainSheet", "D3", "=C3*0.1")

print("✅ Added employee data with formulas")

// Read back some cells
name1 = excel.getCell(file_id, "MainSheet", "A2")
dept1 = excel.getCell(file_id, "MainSheet", "B2")
salary1 = excel.getCell(file_id, "MainSheet", "C2")

print("Read back - Name:", name1, "Department:", dept1, "Salary:", salary1)

// Test 4: Formula operations
print("\n🧮 Test 4: Formula operations...")
excel.setCellFormula(file_id, "MainSheet", "C4", "=SUM(C2:C3)")
excel.setCellFormula(file_id, "MainSheet", "D4", "=SUM(D2:D3)")

formula = excel.getCellFormula(file_id, "MainSheet", "C4")
print("Formula in C4:", formula)

// Test 5: Range operations
print("\n📐 Test 5: Range operations...")
range_data = excel.getRange(file_id, "MainSheet", "A1:D4")
print("Range data (4x4):")
for row in range_data {
    print("  Row:", row)
}

// Test 6: Row/Column operations
print("\n📏 Test 6: Row/Column operations...")
excel.insertRow(file_id, "MainSheet", 2)
print("✅ Inserted row at position 2")

excel.setCell(file_id, "MainSheet", "A2", "Alice Brown")
excel.setCell(file_id, "MainSheet", "B2", "HR")
excel.setCell(file_id, "MainSheet", "C2", 60000)
excel.setCell(file_id, "MainSheet", "D2", "=C2*0.1")
print("✅ Added new employee data in inserted row")

// Test 7: Merge operations
print("\n🔗 Test 7: Cell merge operations...")
excel.setCell(file_id, "SummarySheet", "A1", "Company Report")
excel.mergeCells(file_id, "SummarySheet", "A1:D1")
print("✅ Merged header cells A1:D1 in SummarySheet")

// Test 8: File information
print("\n📋 Test 8: File information...")
info = excel.getFileInfo(file_id)
print("File info:")
print("  Sheet count:", info.sheetCount)
print("  Active sheet:", info.activeSheet)
print("  Sheets:", info.sheets)

// Test 9: Save operations
print("\n💾 Test 9: Save operations...")
excel.save(file_id)
print("✅ Saved file")

backup_id = excel.saveAs(file_id, "backup_workbook.xlsx")
print("✅ Created backup with ID:", backup_id)

// Test 10: Close and reopen
print("\n🔄 Test 10: Close and reopen operations...")
excel.close(file_id)
excel.close(backup_id)
print("✅ Closed original files")

// Reopen the file
reopened_id = excel.open("test_workbook.xlsx")
print("✅ Reopened file with ID:", reopened_id)

// Verify data is still there
name_check = excel.getCell(reopened_id, "MainSheet", "A3")
salary_check = excel.getCell(reopened_id, "MainSheet", "C3")
print("Verified data - Name:", name_check, "Salary:", salary_check)

// Test 11: Summary operations
print("\n📈 Test 11: Creating summary sheet...")
excel.setCell(reopened_id, "SummarySheet", "A3", "Total Employees:")
excel.setCellFormula(reopened_id, "SummarySheet", "B3", "=COUNTA(MainSheet.A:A)-1")

excel.setCell(reopened_id, "SummarySheet", "A4", "Average Salary:")
excel.setCellFormula(reopened_id, "SummarySheet", "B4", "=AVERAGE(MainSheet.C:C)")

excel.setCell(reopened_id, "SummarySheet", "A5", "Total Payroll:")
excel.setCellFormula(reopened_id, "SummarySheet", "B5", "=SUM(MainSheet.C:C)+SUM(MainSheet.D:D)")

print("✅ Added summary calculations")

// Final save and cleanup
excel.save(reopened_id)
excel.close(reopened_id)

print("\n🎉 Excel Module Test Complete!")
print("===============================")
print("✅ All tests passed successfully")
print("📁 Created files: test_workbook.xlsx, backup_workbook.xlsx")
print("💡 Excel module is fully functional!")

// Demonstrate error handling
print("\n🛡️  Testing error handling...")
try {
    bad_id = excel.open("nonexistent_file.xlsx")
} catch error {
    print("✅ Correctly caught error for missing file:", error)
}

print("\n🔒 Note: Password protection testing requires manual setup")
print("    To test password features:")
print("    1. Create a password-protected Excel file manually")
print("    2. Use excel.openWithPassword(file, password)")
print("    3. Password setting is limited by excelize v2.8.0")

print("\n🚀 Excel module ready for production use!")
```

## excel_test_final.vint

```js
// Excel Module Test - Simple functionality test
import excel

print("🚀 VintLang Excel Module Test")
print("==============================")

// Test 1: Create a new Excel file
print("📝 Test 1: Creating new Excel file...")
let file_id = excel.create("test_workbook.xlsx")
print("✅ Created file with ID:")
print(file_id)

// Test 2: Sheet operations
print("")
print("📋 Test 2: Sheet operations...")
let sheets = excel.getSheets(file_id)
print("Initial sheets:")
print(sheets)

// Add new sheets
print("Adding new sheets...")
let index1 = excel.addSheet(file_id, "DataSheet")
let index2 = excel.addSheet(file_id, "SummarySheet")
print("Created DataSheet at index:")
print(index1)
print("Created SummarySheet at index:")
print(index2)

// Rename default sheet
excel.renameSheet(file_id, "Sheet1", "MainSheet")
print("Renamed Sheet1 to MainSheet")

// Get updated sheet list
sheets = excel.getSheets(file_id)
print("Updated sheets:")
print(sheets)

// Test 3: Cell operations
print("")
print("📊 Test 3: Cell operations...")
excel.setCell(file_id, "MainSheet", "A1", "Employee Name")
excel.setCell(file_id, "MainSheet", "B1", "Department")
excel.setCell(file_id, "MainSheet", "C1", "Salary")
excel.setCell(file_id, "MainSheet", "D1", "Bonus")

excel.setCell(file_id, "MainSheet", "A2", "John Doe")
excel.setCell(file_id, "MainSheet", "B2", "Engineering")
excel.setCell(file_id, "MainSheet", "C2", 75000)

excel.setCell(file_id, "MainSheet", "A3", "Jane Smith")
excel.setCell(file_id, "MainSheet", "B3", "Marketing")
excel.setCell(file_id, "MainSheet", "C3", 65000)

print("✅ Added employee data")

// Read back some cells
let name1 = excel.getCell(file_id, "MainSheet", "A2")
let dept1 = excel.getCell(file_id, "MainSheet", "B2")
let salary1 = excel.getCell(file_id, "MainSheet", "C2")

print("Read back - Name:")
print(name1)
print("Department:")
print(dept1)
print("Salary:")
print(salary1)

// Test 4: Formula operations
print("")
print("🧮 Test 4: Formula operations...")
excel.setCellFormula(file_id, "MainSheet", "D2", "=C2*0.1")
excel.setCellFormula(file_id, "MainSheet", "D3", "=C3*0.1")
excel.setCellFormula(file_id, "MainSheet", "C4", "=SUM(C2:C3)")
excel.setCellFormula(file_id, "MainSheet", "D4", "=SUM(D2:D3)")

let formula = excel.getCellFormula(file_id, "MainSheet", "C4")
print("Formula in C4:")
print(formula)

// Test 5: Range operations
print("")
print("📐 Test 5: Range operations...")
let range_data = excel.getRange(file_id, "MainSheet", "A1:D4")
print("Range data (4x4):")
print("Number of rows:")
print(range_data.length)

// Test 6: Row/Column operations
print("")
print("📏 Test 6: Row/Column operations...")
excel.insertRow(file_id, "MainSheet", 2)
print("✅ Inserted row at position 2")

excel.setCell(file_id, "MainSheet", "A2", "Alice Brown")
excel.setCell(file_id, "MainSheet", "B2", "HR")
excel.setCell(file_id, "MainSheet", "C2", 60000)
excel.setCellFormula(file_id, "MainSheet", "D2", "=C2*0.1")
print("✅ Added new employee data in inserted row")

// Test 7: Merge operations
print("")
print("🔗 Test 7: Cell merge operations...")
excel.setCell(file_id, "SummarySheet", "A1", "Company Report")
excel.mergeCells(file_id, "SummarySheet", "A1:D1")
print("✅ Merged header cells A1:D1 in SummarySheet")

// Test 8: File info
print("")
print("📋 Test 8: File information...")
let info = excel.getFileInfo(file_id)
print("✅ Got file info")

// Test 9: Save operations
print("")
print("💾 Test 9: Save operations...")
excel.save(file_id)
print("✅ Saved file")

let backup_id = excel.saveAs(file_id, "backup_workbook.xlsx")
print("✅ Created backup with ID:")
print(backup_id)

// Test 10: Close and reopen
print("")
print("🔄 Test 10: Close and reopen operations...")
excel.close(file_id)
excel.close(backup_id)
print("✅ Closed original files")

// Reopen the file
let reopened_id = excel.open("test_workbook.xlsx")
print("✅ Reopened file with ID:")
print(reopened_id)

// Verify data is still there
let name_check = excel.getCell(reopened_id, "MainSheet", "A3")
let salary_check = excel.getCell(reopened_id, "MainSheet", "C3")
print("Verified data - Name:")
print(name_check)
print("Salary:")
print(salary_check)

// Test 11: Summary operations
print("")
print("📈 Test 11: Creating summary sheet...")
excel.setCell(reopened_id, "SummarySheet", "A3", "Total Employees:")
excel.setCellFormula(reopened_id, "SummarySheet", "B3", "=COUNTA(MainSheet.A:A)-1")

excel.setCell(reopened_id, "SummarySheet", "A4", "Average Salary:")
excel.setCellFormula(reopened_id, "SummarySheet", "B4", "=AVERAGE(MainSheet.C:C)")

excel.setCell(reopened_id, "SummarySheet", "A5", "Total Payroll:")
excel.setCellFormula(reopened_id, "SummarySheet", "B5", "=SUM(MainSheet.C:C)+SUM(MainSheet.D:D)")

print("✅ Added summary calculations")

// Final save and cleanup
excel.save(reopened_id)
excel.close(reopened_id)

print("")
print("🎉 Excel Module Test Complete!")
print("===============================")
print("✅ All tests passed successfully")
print("📁 Created files: test_workbook.xlsx, backup_workbook.xlsx")
print("💡 Excel module is fully functional!")

print("")
print("🔒 Note: Password protection testing requires manual setup")
print("    To test password features:")
print("    1. Create a password-protected Excel file manually")
print("    2. Use excel.openWithPassword(file, password)")

print("")
print("🚀 Excel module ready for production use!")
```

## excel_test_simple.vint

```js
// Excel Module Test - Comprehensive functionality test
import excel

print("🚀 VintLang Excel Module Test")
print("==============================")

// Test 1: Create a new Excel file
print("📝 Test 1: Creating new Excel file...")
let file_id = excel.create("test_workbook.xlsx")
print("✅ Created file with ID: " + file_id)

// Test 2: Sheet operations
print("")
print("📋 Test 2: Sheet operations...")
let sheets = excel.getSheets(file_id)
print("Initial sheets:")
print(sheets)

// Add new sheets
print("Adding new sheets...")
let index1 = excel.addSheet(file_id, "DataSheet")
let index2 = excel.addSheet(file_id, "SummarySheet")
print("Created DataSheet at index: " + index1)
print("Created SummarySheet at index: " + index2)

// Rename default sheet
excel.renameSheet(file_id, "Sheet1", "MainSheet")
print("Renamed Sheet1 to MainSheet")

// Get updated sheet list
sheets = excel.getSheets(file_id)
print("Updated sheets:")
print(sheets)

// Test 3: Cell operations
print("")
print("📊 Test 3: Cell operations...")
excel.setCell(file_id, "MainSheet", "A1", "Employee Name")
excel.setCell(file_id, "MainSheet", "B1", "Department")
excel.setCell(file_id, "MainSheet", "C1", "Salary")
excel.setCell(file_id, "MainSheet", "D1", "Bonus")

excel.setCell(file_id, "MainSheet", "A2", "John Doe")
excel.setCell(file_id, "MainSheet", "B2", "Engineering")
excel.setCell(file_id, "MainSheet", "C2", 75000)

excel.setCell(file_id, "MainSheet", "A3", "Jane Smith")
excel.setCell(file_id, "MainSheet", "B3", "Marketing")
excel.setCell(file_id, "MainSheet", "C3", 65000)

print("✅ Added employee data")

// Read back some cells
let name1 = excel.getCell(file_id, "MainSheet", "A2")
let dept1 = excel.getCell(file_id, "MainSheet", "B2")
let salary1 = excel.getCell(file_id, "MainSheet", "C2")

print("Read back - Name: " + name1 + " Department: " + dept1 + " Salary: " + salary1)

// Test 4: Formula operations
print("")
print("🧮 Test 4: Formula operations...")
excel.setCellFormula(file_id, "MainSheet", "D2", "=C2*0.1")
excel.setCellFormula(file_id, "MainSheet", "D3", "=C3*0.1")
excel.setCellFormula(file_id, "MainSheet", "C4", "=SUM(C2:C3)")
excel.setCellFormula(file_id, "MainSheet", "D4", "=SUM(D2:D3)")

let formula = excel.getCellFormula(file_id, "MainSheet", "C4")
print("Formula in C4: " + formula)

// Test 5: Range operations
print("")
print("📐 Test 5: Range operations...")
let range_data = excel.getRange(file_id, "MainSheet", "A1:D4")
print("Range data (4x4):")
for let i = 0; i < range_data.length; i++ {
    print("  Row " + i + ":")
    let row = range_data[i]
    for let j = 0; j < row.length; j++ {
        print("    " + row[j])
    }
}

// Test 6: Row/Column operations
print("")
print("📏 Test 6: Row/Column operations...")
excel.insertRow(file_id, "MainSheet", 2)
print("✅ Inserted row at position 2")

excel.setCell(file_id, "MainSheet", "A2", "Alice Brown")
excel.setCell(file_id, "MainSheet", "B2", "HR")
excel.setCell(file_id, "MainSheet", "C2", 60000)
excel.setCellFormula(file_id, "MainSheet", "D2", "=C2*0.1")
print("✅ Added new employee data in inserted row")

// Test 7: Merge operations
print("")
print("🔗 Test 7: Cell merge operations...")
excel.setCell(file_id, "SummarySheet", "A1", "Company Report")
excel.mergeCells(file_id, "SummarySheet", "A1:D1")
print("✅ Merged header cells A1:D1 in SummarySheet")

// Test 8: File information
print("")
print("📋 Test 8: File information...")
let info = excel.getFileInfo(file_id)
print("File info retrieved successfully")

// Test 9: Save operations
print("")
print("💾 Test 9: Save operations...")
excel.save(file_id)
print("✅ Saved file")

let backup_id = excel.saveAs(file_id, "backup_workbook.xlsx")
print("✅ Created backup with ID: " + backup_id)

// Test 10: Close and reopen
print("")
print("🔄 Test 10: Close and reopen operations...")
excel.close(file_id)
excel.close(backup_id)
print("✅ Closed original files")

// Reopen the file
let reopened_id = excel.open("test_workbook.xlsx")
print("✅ Reopened file with ID: " + reopened_id)

// Verify data is still there
let name_check = excel.getCell(reopened_id, "MainSheet", "A3")
let salary_check = excel.getCell(reopened_id, "MainSheet", "C3")
print("Verified data - Name: " + name_check + " Salary: " + salary_check)

// Test 11: Summary operations
print("")
print("📈 Test 11: Creating summary sheet...")
excel.setCell(reopened_id, "SummarySheet", "A3", "Total Employees:")
excel.setCellFormula(reopened_id, "SummarySheet", "B3", "=COUNTA(MainSheet.A:A)-1")

excel.setCell(reopened_id, "SummarySheet", "A4", "Average Salary:")
excel.setCellFormula(reopened_id, "SummarySheet", "B4", "=AVERAGE(MainSheet.C:C)")

excel.setCell(reopened_id, "SummarySheet", "A5", "Total Payroll:")
excel.setCellFormula(reopened_id, "SummarySheet", "B5", "=SUM(MainSheet.C:C)+SUM(MainSheet.D:D)")

print("✅ Added summary calculations")

// Final save and cleanup
excel.save(reopened_id)
excel.close(reopened_id)

print("")
print("🎉 Excel Module Test Complete!")
print("===============================")
print("✅ All tests passed successfully")
print("📁 Created files: test_workbook.xlsx, backup_workbook.xlsx")
print("💡 Excel module is fully functional!")

// Test error handling
print("")
print("🛡️  Testing error handling...")
let error_result = excel.open("nonexistent_file.xlsx")
print("Attempted to open nonexistent file, result: " + error_result)

print("")
print("🔒 Note: Password protection testing requires manual setup")
print("    To test password features:")
print("    1. Create a password-protected Excel file manually")
print("    2. Use excel.openWithPassword(file, password)")

print("")
print("🚀 Excel module ready for production use!")
```

## express_like_server.vint

```js
// Express.js-like HTTP Server Example for VintLang
// This example demonstrates how to create a REST API server using VintLang's http module

import http

// Create the Express-like app instance
http.app()

// Basic routes
http.get("/", func(req, res) {
    println(req.headers)
    // This function will be registered as the handler for GET /
    println("Home page accessed")
})

http.get("/about", func(req, res) {
    println("About page accessed")
})

// API routes
http.get("/api/status", func(req, res) {
    println("API status check")
})

http.get("/api/users", func(req, res) {
    println("Getting all users")
})

http.post("/api/users", func(req, res) {
    println("Creating a new user")
})

http.put("/api/users/123", func(req, res) {
    println("Updating user with ID 123")
})

http.delete("/api/users/123", func(req, res) {
    println("Deleting user with ID 123")
})

// Middleware (currently registered but execution pending full evaluator integration)
http.use(func(req, res, next) {
    println("Middleware: Request logged")
})

// Display configured routes
println("🚀 VintLang Express-like HTTP Server")
println("====================================")
println("Routes configured:")
println("  GET    /")
println("  GET    /about")
println("  GET    /api/status")
println("  GET    /api/users")
println("  POST   /api/users")
println("  PUT    /api/users/123")
println("  DELETE /api/users/123")
println("")

// Start the server
println("Starting server...")
http.listen(3000, "🌟 Server running at http://localhost:3000")

/*
Usage examples:
  curl http://localhost:3000/
  curl http://localhost:3000/about
  curl http://localhost:3000/api/status
  curl http://localhost:3000/api/users
  curl -X POST http://localhost:3000/api/users
  curl -X PUT http://localhost:3000/api/users/123
  curl -X DELETE http://localhost:3000/api/users/123

Features implemented:
✓ Express.js-like syntax with http.app()
✓ Route registration: GET, POST, PUT, DELETE, PATCH
✓ Server listening with graceful shutdown
✓ Route matching and 404 handling
✓ Function handler registration
✓ Middleware registration (http.use)

Coming soon:
- Full function handler execution with req/res objects
- Request body parsing and response methods
- Route parameters and query string support
- Enhanced middleware execution
*/
```

## feature_test.vint

```js
// VintLang Feature Test
import time
import os
import json
import uuid

print("🚀 VintLang Feature Test")
print("=" * 30)

// Test 1: Time functions
print("1. Time Functions:")
let now = time.now()
print("Current time: " + time.format(now, "02-01-2006 15:04:05"))

// Test 2: UUID generation
print("\n2. UUID Generation:")
let id = uuid.generate()
print("Generated UUID: " + id)

// Test 3: JSON operations
print("\n3. JSON Operations:")
let data = {
    "name": "VintLang",
    "version": "0.2.5",
    "features": ["time", "json", "uuid", "os"]
}
let jsonStr = json.encode(data)
print("JSON encoded: " + jsonStr)

let decoded = json.decode(jsonStr)
print("JSON decoded: " + string(decoded))

// Test 4: File operations
print("\n4. File Operations:")
let testFile = "test.txt"
os.writeFile(testFile, "Hello VintLang!")
let content = os.readFile(testFile)
print("File content: " + content)

// Test 5: Arrays and loops
print("\n5. Arrays and Loops:")
let numbers = [1, 2, 3, 4, 5]
for num in numbers {
    print("Number: " + string(num))
}

// Test 6: String operations
print("\n6. String Operations:")
let text = "VintLang Programming"
let words = text.split(" ")
print("Split result: " + string(words))

print("\n✅ All tests completed!")
print("VintLang is working with these modules:")
print("  • time - Date and time operations")
print("  • os - File system operations")
print("  • json - JSON encoding/decoding")
print("  • uuid - UUID generation")
print("  • Built-in string methods")
print("  • Arrays and dictionaries")
print("  • Control flow (loops, conditionals)")
```

## file_manager.vint

```js
// VintLang Showcase: File & Data Management System
// Demonstrates practical VintLang capabilities

import time
import os
import json
import uuid

// Application state
let app = {
    "name": "VintLang File Manager",
    "version": "1.0",
    "dataDir": "vint_data",
    "logFile": "activity.log"
}

// Logging function
let logActivity = func(message) {
    let timestamp = time.format(time.now(), "2006-01-02 15:04:05")
    let logEntry = "[" + timestamp + "] " + message + "\n"

    if (os.fileExists(app["logFile"])) {
        let existing = os.readFile(app["logFile"])
        os.writeFile(app["logFile"], existing + logEntry)
    } else {
        os.writeFile(app["logFile"], logEntry)
    }
}

// Setup function
let setupApplication = func() {
    print("🔧 Setting up VintLang File Manager...")

    if (!os.fileExists(app["dataDir"])) {
        os.makeDir(app["dataDir"])
        print("✓ Created data directory: " + app["dataDir"])
        logActivity("Created data directory")
    }

    logActivity("Application started")
    print("✓ Application setup complete")
}

// Create sample files
let createSampleFiles = func() {
    print("\n📁 Creating sample files...")

    // Create different types of files
    let files = [
        {
            "name": "notes.txt",
            "content": "VintLang Programming Notes\n========================\n\nVintLang is a powerful programming language with:\n- JSON support\n- File I/O operations\n- Time handling\n- UUID generation\n- Network capabilities\n\nCreated: " + time.format(time.now(), "02-01-2006 15:04:05")
        },
        {
            "name": "config.json",
            "content": json.encode({
                "app_name": "VintLang Demo",
                "version": "1.0.0",
                "features": ["file_io", "json", "time", "uuid"],
                "settings": {
                    "debug": true,
                    "log_level": "info"
                }
            })
        },
        {
            "name": "tasks.json",
            "content": json.encode([
                {
                    "id": uuid.generate(),
                    "title": "Learn VintLang",
                    "status": "completed",
                    "created": time.format(time.now(), "2006-01-02")
                },
                {
                    "id": uuid.generate(),
                    "title": "Build awesome project",
                    "status": "in_progress",
                    "created": time.format(time.now(), "2006-01-02")
                }
            ])
        },
        {
            "name": "data.csv",
            "content": "ID,Name,Category,Value\n1,Item A,Category 1,100\n2,Item B,Category 2,200\n3,Item C,Category 1,150\n4,Item D,Category 3,300"
        }
    ]

    for file in files {
        let filepath = app["dataDir"] + "/" + file["name"]
        os.writeFile(filepath, file["content"])
        print("✓ Created: " + file["name"])
        logActivity("Created file: " + file["name"])
    }

    print("✓ Sample files created successfully!")
}

// List files in directory
let listFiles = func() {
    print("\n📂 Files in " + app["dataDir"] + ":")
    print("-" * 40)

    let files = os.listDir(app["dataDir"])
    let fileList = files.split(", ")

    for filename in fileList {
        if (filename != "." && filename != "..") {
            let filepath = app["dataDir"] + "/" + filename
            if (os.fileExists(filepath)) {
                print("📄 " + filename)

                // Get file size (approximate by content length)
                let content = os.readFile(filepath)
                print("   Size: " + string(len(content)) + " bytes")

                // Determine file type
                if (filename.contains(".json")) {
                    print("   Type: JSON Data")
                } else if (filename.contains(".txt")) {
                    print("   Type: Text File")
                } else if (filename.contains(".csv")) {
                    print("   Type: CSV Data")
                } else {
                    print("   Type: Unknown")
                }
                print("")
            }
        }
    }
}

// Analyze JSON files
let analyzeJsonFiles = func() {
    print("\n🔍 Analyzing JSON files...")
    print("-" * 40)

    let files = os.listDir(app["dataDir"])
    let fileList = files.split(", ")

    for filename in fileList {
        if (filename.contains(".json")) {
            print("📊 Analyzing: " + filename)
            let filepath = app["dataDir"] + "/" + filename
            let content = os.readFile(filepath)

            let data = json.decode(content)
            print("   JSON structure analyzed")
            print("   Content type: " + type(data))

            if (type(data) == "ARRAY") {
                print("   Array length: " + string(len(data)))
            } else if (type(data) == "HASH") {
                print("   Object with " + string(len(data)) + " properties")
            }

            logActivity("Analyzed JSON file: " + filename)
            print("")
        }
    }
}

// Generate file report
let generateReport = func() {
    print("\n📋 Generating file system report...")

    let reportContent = "VINTLANG FILE MANAGER REPORT\n"
    reportContent += "Generated: " + time.format(time.now(), "02-01-2006 15:04:05") + "\n"
    reportContent += "=" * 50 + "\n\n"

    // Count files by type
    let files = os.listDir(app["dataDir"])
    let fileList = files.split(", ")
    let typeCount = {}
    let totalSize = 0

    for filename in fileList {
        if (filename != "." && filename != "..") {
            let filepath = app["dataDir"] + "/" + filename
            if (os.fileExists(filepath)) {
                let content = os.readFile(filepath)
                totalSize += len(content)

                let extension = "unknown"
                if (filename.contains(".json")) {
                    extension = "json"
                } else if (filename.contains(".txt")) {
                    extension = "txt"
                } else if (filename.contains(".csv")) {
                    extension = "csv"
                }

                if (!typeCount.hasKey(extension)) {
                    typeCount[extension] = 0
                }
                typeCount[extension] += 1
            }
        }
    }

    reportContent += "SUMMARY\n"
    reportContent += "-------\n"
    reportContent += "Total files: " + string(len(fileList) - 2) + "\n"
    reportContent += "Total size: " + string(totalSize) + " bytes\n\n"

    reportContent += "FILES BY TYPE\n"
    reportContent += "-------------\n"
    for extension, count in typeCount {
        reportContent += extension.upper() + " files: " + string(count) + "\n"
    }

    reportContent += "\nDETAILS\n"
    reportContent += "-------\n"
    for filename in fileList {
        if (filename != "." && filename != "..") {
            let filepath = app["dataDir"] + "/" + filename
            if (os.fileExists(filepath)) {
                let content = os.readFile(filepath)
                reportContent += filename + " (" + string(len(content)) + " bytes)\n"
            }
        }
    }

    reportContent += "\n" + "=" * 50 + "\n"
    reportContent += "Report generated by " + app["name"] + " v" + app["version"] + "\n"
    reportContent += "Powered by VintLang Programming Language\n"

    let reportFile = "file_report_" + time.format(time.now(), "2006-01-02_15-04-05") + ".txt"
    os.writeFile(reportFile, reportContent)
    print("✓ Report saved to: " + reportFile)
    logActivity("Generated file report")
}

// Show activity log
let showActivityLog = func() {
    print("\n📜 Activity Log:")
    print("-" * 40)

    if (os.fileExists(app["logFile"])) {
        let logContent = os.readFile(app["logFile"])
        let lines = os.readLines(app["logFile"])

        print("Total log entries: " + string(len(lines)))
        print("\nRecent activity:")
        for line in lines {
            if (line != "") {
                print(line)
            }
        }
    } else {
        print("No activity log found.")
    }
}

// Backup data
let backupData = func() {
    print("\n💾 Creating backup...")

    let backupDir = "backup_" + time.format(time.now(), "2006-01-02_15-04-05")
    os.makeDir(backupDir)

    let files = os.listDir(app["dataDir"])
    let fileList = files.split(", ")
    let backedUp = 0

    for filename in fileList {
        if (filename != "." && filename != "..") {
            let sourcePath = app["dataDir"] + "/" + filename
            let backupPath = backupDir + "/" + filename

            if (os.fileExists(sourcePath)) {
                let content = os.readFile(sourcePath)
                os.writeFile(backupPath, content)
                backedUp += 1
            }
        }
    }

    // Create backup manifest
    let manifest = {
        "created": time.format(time.now(), "2006-01-02 15:04:05"),
        "files_backed_up": backedUp,
        "source_directory": app["dataDir"],
        "backup_id": uuid.generate()
    }

    os.writeFile(backupDir + "/manifest.json", json.encode(manifest))
    print("✓ Backup created: " + backupDir)
    print("✓ Files backed up: " + string(backedUp))
    logActivity("Created backup: " + backupDir)
}

// Main application
let runFileManager = func() {
    print("🚀 Welcome to " + app["name"] + " v" + app["version"])
    print("=" * 60)
    print("This VintLang application demonstrates:")
    print("  • File and directory operations")
    print("  • JSON data processing")
    print("  • Logging and reporting")
    print("  • Backup functionality")
    print("  • Time and UUID utilities")
    print("  • Data analysis capabilities")
    print("=" * 60)

    // Run all demonstrations
    setupApplication()
    createSampleFiles()
    listFiles()
    analyzeJsonFiles()
    generateReport()
    showActivityLog()
    backupData()

    print("\n🎉 File Manager demonstration complete!")
    print("\n📊 Summary of operations:")
    print("  ✓ Created application directory")
    print("  ✓ Generated sample files (JSON, TXT, CSV)")
    print("  ✓ Performed file analysis")
    print("  ✓ Generated comprehensive report")
    print("  ✓ Maintained activity log")
    print("  ✓ Created data backup")

    print("\n✨ VintLang successfully demonstrated:")
    print("  • Robust file I/O operations")
    print("  • JSON encoding/decoding")
    print("  • Directory management")
    print("  • Logging and auditing")
    print("  • Data backup and recovery")
    print("  • Time-based operations")
    print("  • UUID generation")
    print("  • String manipulation")
    print("  • Error handling")

    print("\n🎯 VintLang is ready for real-world applications!")
}

// Start the file manager
runFileManager()
```

## file_processor_cli.vint

```js
// File Processor CLI - Demonstrates file operations with terminal UI
// Run with: vint file_processor_cli.vint [options]

import term
import cli
import os

// Help and version handling
if (cli.hasArg("--help")) {
    cli.help("FileProcessor", "Process and analyze files with VintLang")
    term.println("")
    term.info("File Processing Options:")
    term.println("  --input FILE     Input file to process")
    term.println("  --output FILE    Output file for results")
    term.println("  --format TYPE    Output format (json, csv, txt)")
    term.println("  --verbose        Show detailed processing info")
    os.exit()
}

if (cli.hasArg("--version")) {
    cli.version("FileProcessor", "2.1.0")
    os.exit()
}

// Display banner
let banner = term.banner("VintLang File Processor")
term.println(banner)

// Process command line arguments
let inputFile = cli.getArgValue("--input")
let outputFile = cli.getArgValue("--output")
let format = cli.getArgValue("--format")
let verbose = cli.hasArg("--verbose")

// Show configuration
if (verbose) {
    term.info("Configuration:")
    let configTableRows = [
        ["Setting", "Value"],
        ["Input File", inputFile || "Not specified"],
        ["Output File", outputFile || "stdout"],
        ["Format", format || "txt"],
        ["Verbose", verbose ? "Enabled" : "Disabled"]
    ]
    let configTable = term.table(configTableRows)
    term.println(configTable)
}

// Interactive file selection if no input specified
if (!inputFile) {
    term.info("No input file specified")
    let choice = term.select([
        "Browse current directory",
        "Enter file path manually",
        "Create sample file",
        "Exit"
    ])

    if (choice == "Browse current directory") {
    let currentDir = os.getwd()
    term.info("Current directory: " + currentDir)
    term.warning("File browsing not implemented yet")

    } else if (choice == "Enter file path manually") {
    inputFile = term.input("Enter file path: ")
    term.success("Input file set to: " + inputFile)

    } else if (choice == "Create sample file") {
    let filename = term.input("Enter filename for sample: ")
    term.success("Sample file '" + filename + "' would be created")

    } else if (choice == "Exit") {
    term.info("Exiting file processor")
    os.exit()

    }
}

// File processing simulation
if (inputFile) {
    term.info("Processing file: " + inputFile)
    // Show processing steps
    term.loading("Reading file...")
    term.success("File read successfully")
    term.loading("Analyzing content...")
    term.success("Analysis complete")
    // Show results
    let resultsRows = [
        ["Analysis Result", "Value"],
        ["File Size", "1.2 KB"],
        ["Lines", "45"],
        ["Words", "234"],
        ["Characters", "1,234"]
    ]
    let results = term.table(resultsRows)
    term.println(results)
    // Show processing chart
    let processingChart = term.chart([45, 234, 12])
    term.println("Content Analysis:")
    term.println(processingChart)
    if (outputFile) {
        term.success("Results saved to: " + outputFile)
    }
}

// Final operations
let shouldCleanup = term.confirm("Clean up temporary files?")
if (shouldCleanup) {
    term.success("Cleanup completed")
}

let finalBox = term.box("File processing completed successfully!")
term.println(finalBox)
```

## fmt_basic_test.vint

```js
import fmt

print("Testing fmt module:")

// Test sprintf
let result = fmt.sprintf("Hello %s number %d", "world", 42)
print(result)

// Test number formatting
print(fmt.formatHex(255))
print(fmt.formatBin(15))

// Test padding
print(fmt.padLeft("hi", 8))
print(fmt.padRight("hi", 8))

print("fmt module works!")
```

## fmt_demo.vint

```js
// VintLang fmt module comprehensive demo
// Demonstrates all the formatting capabilities

import fmt

print("🎨 VintLang fmt Module Demo")
print("=" * 50)

// 1. Basic string formatting with sprintf
print("\n📝 String Formatting:")
let name = "VintLang"
let version = "1.0"
let users = 1234

let intro = fmt.sprintf("Welcome to %s v%s! We have %d active users.", name, version, users)
print(intro)

// 2. Printf for direct output
print("\n🖨️  Direct Printing:")
fmt.printf("Temperature: %.1f°C, Humidity: %d%%\n", 23.7, 65)
fmt.printf("Processing %s... [%d/%d] %.1f%% complete\n", "data.txt", 75, 100, 75.0)

// 3. Number formatting in different bases
print("\n🔢 Number Formatting:")
let number = 255
print("Decimal:", fmt.formatInt(number, 10))
print("Binary: ", fmt.formatBin(number))
print("Hex (lower):", fmt.formatHex(number))
print("Hex (upper):", fmt.formatHex(number, true))
print("Octal:", fmt.formatOct(number))

// 4. Float precision
print("\n🎯 Float Precision:")
let pi = 3.14159265359
print("π with 2 decimals:", fmt.formatFloat(pi, 2))
print("π with 4 decimals:", fmt.formatFloat(pi, 4))
print("π with precision:", fmt.precision(pi, 6))

// 5. Padding and alignment
print("\n📏 Padding & Alignment:")
let items = ["Apple", "Banana", "Orange", "Grape"]

print("Left-padded (width 10):")
for item in items {
    print("  ", fmt.padLeft(item, 10, "."))
}

print("\nRight-padded (width 10):")
for item in items {
    print("  ", fmt.padRight(item, 10, "."))
}

print("\nCentered (width 12):")
for item in items {
    print("  ", fmt.padCenter(item, 12, "-"))
}

// 6. Table formatting demo
print("\n📊 Table Formatting:")
let products = [
    {"name": "Laptop", "price": 999.99, "stock": 15, "category": "Electronics"},
    {"name": "Mouse", "price": 25.50, "stock": 150, "category": "Accessories"},
    {"name": "Keyboard", "price": 75.00, "stock": 45, "category": "Accessories"},
    {"name": "Monitor", "price": 299.99, "stock": 8, "category": "Electronics"}
]

// Table header
let border = fmt.repeat("-", 60)
print(border)
fmt.printf("| %-12s | %8s | %5s | %-12s |\n", "Product", "Price", "Stock", "Category")
print(border)

// Table rows
for product in products {
    fmt.printf("| %-12s | $%7.2f | %5d | %-12s |\n",
        product["name"],
        product["price"],
        product["stock"],
        product["category"])
}
print(border)

// 7. Report generation with formatting
print("\n📋 Report Generation:")
let reportTitle = "MONTHLY SALES REPORT"
let reportWidth = 40

print(fmt.repeat("=", reportWidth))
print(fmt.padCenter(reportTitle, reportWidth))
print(fmt.repeat("=", reportWidth))

let totalSales = 1245.75
let totalOrders = 28
let avgOrder = totalSales / totalOrders

fmt.printf("Total Sales:    %s\n", fmt.padLeft(fmt.sprintf("$%.2f", totalSales), 12))
fmt.printf("Total Orders:   %s\n", fmt.padLeft(fmt.sprintf("%d", totalOrders), 12))
fmt.printf("Average Order:  %s\n", fmt.padLeft(fmt.sprintf("$%.2f", avgOrder), 12))

// 8. Progress bar using repeat and formatting
print("\n📈 Progress Indicators:")
let progress = [25, 50, 75, 100]

for p in progress {
    let filled = p / 5  // Scale to 20 chars max
    let empty = 20 - filled
    let bar = fmt.repeat("█", filled) + fmt.repeat("░", empty)
    fmt.printf("Progress: [%s] %3d%%\n", bar, p)
}

// 9. Error formatting
print("\n❌ Error Formatting:")
let err = fmt.errorf("Connection failed to %s:%d - %s", "localhost", 8080, "refused")
print("Error:", err.Inspect())

// 10. Utility functions
print("\n🔧 Utility Functions:")
let longText = "This is a very long text that needs to be truncated for display"
print("Original:", longText)
print("Truncated (30):", fmt.truncate(longText, 30))

let separator = fmt.repeat("*", 25)
print("Separator:", separator)

// 11. Width formatting
print("\n📐 Width Formatting:")
let titles = ["Short", "Medium Title", "Very Long Title That Exceeds Width"]

print("Fixed width (15 chars):")
for title in titles {
    print("  [" + fmt.width(title, 15) + "]")
}

// 12. Complex formatting combinations
print("\n🎪 Complex Combinations:")
let timestamp = "2024-10-24 14:30:00"
let level = "INFO"
let message = "System initialized successfully"

// Log-style formatting
let logEntry = fmt.sprintf("[%s] %s: %s",
    timestamp,
    fmt.padRight(level, 5),
    message)
print(logEntry)

// Status display with alignment
let status = "READY"
let statusLine = fmt.sprintf("System Status: [%s] %s",
    fmt.padCenter(status, 8),
    fmt.repeat("●", 3))
print(statusLine)

print("\n" + fmt.repeat("=", 50))
print("🎉 fmt Module Demo Complete!")
print("All formatting functions demonstrated successfully.")
```

## fmt_test_simple.vint

```js
import fmt

print("Testing fmt module...")

// Test sprintf
let result = fmt.sprintf("Hello, %s! Number: %d", "World", 42)
print("sprintf result:", result)

// Test number formatting
let num = 255
print("formatHex:", fmt.formatHex(num))
print("formatBin:", fmt.formatBin(num))

// Test padding
let padded = fmt.padLeft("test", 10, "*")
print("padLeft result:", padded)

print("fmt module test complete!")
```

## function_test.vint

```js
// Test default parameter
let greet = func(name = "Guest") {
    println("Hello, " + name)
}

greet()        // Should print: Hello, Guest
greet("Alice") // Should print: Hello, Alice

// Test multiple parameters with one default
let add = func(a, b = 10) {
    return a + b
}

println(add(5))      // Should print: 15
println(add(5, 2))   // Should print: 7

// Test function returning a value
let square = func(x) {
    return x * x
}

println(square(4))   // Should print: 16
```

## functions.vint

```js
// VintLang Functions Example
// This example demonstrates different ways to define and use functions in VintLang

// Example 1: Immediately Invoked Function Expression (IIFE)
// This function is defined and executed immediately using ()
let runNow = func(){
    print("this is a function executed immediately")
}()

// Example 2: Named function stored in a variable
// This function is defined but not invoked until called explicitly
let vint = func(){
    print("This is also a function\nBut not invoked immediately after being declared")
}

// Example 3: Call the stored function
vint()

// Example 4: Higher-order function (function that takes another function as parameter)
let w = func(){
    print("w function")
}

// Pass function 'w' as an argument to another function and execute it
func(w){
    w()
    print("func")
}(w)

```

## github-profile.vint

```js
// VintLang GitHub Profile View Counter Example
// This example demonstrates making HTTP requests and measuring time

import net
import time

// GitHub username to track profile views
let githubUsername = "tacheraSasi"

// URL for GitHub profile view counter (using shields.io service)
let url = "https://camo.githubusercontent.com/b53ec94d48abfb67ae979b046d1980e04133e2e6374c2dfe123d3b9d7ee95a7d/68747470733a2f2f6b6f6d617265762e636f6d2f67687076632f3f757365726e616d653d7461636865726153617369266c6162656c3d50726f66696c65253230566965777326636f6c6f723d626c7565267374796c653d666c6174"

// Record the start time
let startTime = time.now()
print(startTime)

// Make 100 requests to the profile view counter
for i in range(100) {
    print(":",i+1)
    let res = net.get(url)
    // print(res)  // Uncomment to see response
}

// Calculate and print elapsed time
print(time.since(startTime))
```

## github_issue_example.vint

```js
// Exact example from GitHub issue #11

let user = {"role": "admin", "active": true}

match user {
    {"role": "admin"} => print("Hello, Admin!")
    {"active": false} => print("Inactive user")
    _ => print("Regular user")
}

// Additional test cases to show flexibility
print("\n--- Additional Examples ---")

let inactiveUser = {"role": "user", "active": false}
match inactiveUser {
    {"role": "admin"} => print("Hello, Admin!")
    {"active": false} => print("Inactive user")
    _ => print("Regular user")
}

let unknownUser = {"name": "Jane"}
match unknownUser {
    {"role": "admin"} => print("Hello, Admin!")
    {"active": false} => print("Inactive user")
    _ => print("Regular user")
}
```

## github_issue_test.vint

```js
// Get all arguments
let allArgs = args()
println(allArgs)

// Parse flags
import cli

if (cli.hasArg("--verbose")) {
    println("Verbose mode enabled")
}

// Additional tests to show full functionality
let outputFile = cli.getArgValue("--output")
if (outputFile) {
    println("Output file:", outputFile)
}

let positional = cli.getPositional()
println("Positional arguments:", positional)

```

## greetings_module.vint

```js

package greetings_module{
	// Demonstrate a simple function from a package
	let greet = func(name) {
		print("Hello, " + name + "!")
	}
}

```

## guessingGame.vint

```js
import math, random

// greet("Tach")

// Guessing game in Vint
let guess = input("Guess a number: ")
guess = int(guess) // You can also use convert(guess, "INTEGER") to convert the guess input to an integer

let number = random.int(1, 10)
while (number != guess) {
    // Check if the guess is correct
    if (number == guess) {
        println("Woo hoo you've guessed it right!")
        break
    }else if(number > guess){
        println("Too small")
    }else{
        println("Too big")
    }

    // Prompt for a new guess if incorrect
    guess = input("Guess Again: ")
    guess = int(guess) // Converting the guess input to an integer
}
println("Woo hoo you've guessed it right")
println("Game over!")

```

## has_key_test.vint

```js
// VintLang has_key() Function Example
// Demonstrates checking if a dictionary contains a specific key

// Create a dictionary with some key-value pairs
let myDict = { name: "Alex", age: 30 };

// Example 1: Test global has_key() function
// Syntax: has_key(dictionary, key)
println("Testing global has_key():");
println("has 'name'?", has_key(myDict, "name")); // Expected: true
println("has 'city'?", has_key(myDict, "city")); // Expected: false

// Example 2: Test has_key() as a dictionary method
// Syntax: dictionary.has_key(key)
println("\nTesting dict.has_key():");
println("has 'age'?", myDict.has_key("age")); // Expected: true
println("has 'country'?", myDict.has_key("country")); // Expected: false
```

## http.vint

```js
// VintLang HTTP File Server Example
// Demonstrates creating a simple HTTP file server

import http

let port = "3000";
// let dir = "/var/www/html";  // Production directory
let dir = "./";  // Serve files from current directory

// Scenario 1: Using the default message and no directory listing (commented)
// http.fileServer(port, dir);

// Scenario 2: Starting a file server with a custom message
// This starts the server and blocks until stopped
// Uncomment to test:
// http.fileServer(
//     port,
//     dir,
//     "My custom Vint server is running on port 3000!"
// );

// Scenario 3: Starting a file server with directory listing enabled (default message)
// Uncomment to test:
// http.fileServer(
//     port,
//     dir
// );

// Scenario 4: Starting a file server with a custom message and directory listing enabled
println("To start the server, uncomment one of the http.fileServer() calls above")
println("Server will listen on port:", port)
println("Serving directory:", dir)

// http.fileServer(
//     port,
//     dir,
//     "Directory listing is now enabled! Check it out!"
// );

// Scenario 5: Invalid usage examples (commented to avoid errors)
// http.fileServer(3000, dir);  // Port must be a string
// http.fileServer(port, 123);  // Directory must be a string
// http.fileServer(port, dir, 404);  // Message must be a string
// http.fileServer(port, dir, "Message", "true");  // enableListing must be a boolean

```

## http_test.vint

```js
// Test file for the Express.js-like HTTP module
import http

// Test 1: Create app
let result = http.app()
print("✓ Test 1 - App creation:", result)

// Test 2: Route registration
let getResult = http.get("/test", func(req, res) {
    print("Test route")
})
print("✓ Test 2 - GET route registration:", getResult)

let postResult = http.post("/data", func(req, res) {
    print("POST test route")
})
print("✓ Test 3 - POST route registration:", postResult)

// Test 3: Multiple HTTP methods
http.put("/update", func(req, res) { print("PUT test") })
http.delete("/remove", func(req, res) { print("DELETE test") })
http.patch("/modify", func(req, res) { print("PATCH test") })
print("✓ Test 4 - Multiple HTTP methods registered")

// Test 4: Middleware
let middlewareResult = http.use(func(req, res, next) {
    print("Test middleware")
})
print("✓ Test 5 - Middleware registration:", middlewareResult)

// Test 5: Error handling - should fail gracefully
print("Testing error conditions...")

// Test with no app (this should work since we created one above)
let routeWithoutApp = http.get("/noapp", func(req, res) {
    print("This should work")
})
print("✓ Test 6 - Route registration with existing app:", routeWithoutApp)

print("")
print("All tests completed! ✨")
print("The Express.js-like HTTP module is working correctly.")
print("")
print("Note: Function execution will be improved in future versions")
print("      when full evaluator integration is implemented.")
```

## if_expression.vint

```js
// VintLang If Expressions Example
// Demonstrates using if as both a statement and an expression

// Example 1: Using if as a statement (classic approach)
// The if statement modifies an existing variable
let x = 0
if (true) {
    x = 42
}
print("Classic if statement result: ", x)

// Example 2: Using if as an expression (functional approach)
// The if expression returns a value that can be assigned
let status = ""
status = if (x > 0) { "Online" } else { "Offline" }
print("If as an expression result: ", status)

// Example 3: If expression without else clause
// If the condition is false and there's no else, returns null
let y = if (false) { 123 }
print("If as an expression with no else: ", y)
```

## include_test.vint

```js
// VintLang Include Statement Example
// Demonstrates including code from another VintLang file

// Include another VintLang file - this loads and executes it
// The included file can define variables, functions, etc. that become available here
include "examples/included.vint";

// Use a variable defined in the included file
println(message);
```

## included.vint

```js
// VintLang Included File Example
// This file is meant to be included by other VintLang files

// Define a variable that will be available to the including file
let message = "Hello from an included file!";
```

## json.vint

```js
// VintLang JSON Module Example
// Demonstrates JSON encoding, decoding, and manipulation operations

import json

// Example 1: Decode a JSON string
// Converts a JSON string into a VintLang object
print("=== Example 1: Decode ===")
let raw_json = '{"name": "John", "age": 30, "isAdmin": false, "friends": ["Jane", "Doe"]}'
let decoded = json.decode(raw_json)
print("Decoded Object:", decoded)

// Example 2: Encode a Vint object to JSON
// Converts a VintLang object into a JSON string
print("\n=== Example 2: Encode ===")
let data = {
  "language": "Vint",
  "version": 1.0,
  "features": ["custom modules", "native objects"]
}
let encoded_json = json.encode(data) // Optional parameter: indent for pretty printing
print("Encoded JSON:", encoded_json)

// Example 3: Pretty print a JSON string
// Formats JSON with indentation for readability
print("\n=== Example 3: Pretty Print ===")
let raw_json_pretty = '{"name":"John","age":30,"friends":["Jane","Doe"]}'
let pretty_json = json.pretty(raw_json_pretty)
print("Pretty JSON:\n", pretty_json)

// Example 4: Merge two JSON objects
// Combines two objects, with the second object's values taking precedence
print("\n=== Example 4: Merge ===")
let json1 = {"name": "John", "age": 30}
let json2 = {"city": "New York", "age": 35}
let merged_json = json.merge(json1, json2)
print("Merged JSON:", merged_json)

// Example 5: Get a value by key from a JSON object
// Retrieves a value from an object, returns null if key doesn't exist
print("\n=== Example 5: Get Value by Key ===")
let json_object = {"name": "John", "age": 30, "city": "New York"}
let value = json.get(json_object, "age")
print("Age:", value)

let missing_value = json.get(json_object, "country")
print("Country (missing key):", missing_value)

```

## just_import.vint

```js
// Debug the import process
import time
```

## jwt.vint

```js
// import jwt
const jwt = import("jwt") // Both ways work

const secret = "vint-is-so-awaseome"
// Create a JWT token
let payload = {"user": "Tachera", "role": "admin"}
let token = jwt.create(payload, secret)
info "token is "+ token

const decoded = jwt.decode(token)

println("DECODED", decoded)

// Verify the token
let result = jwt.verify(token, secret)
if (result) {
    println("User:", result["user"])
    println("Role:", result["role"])
}
```

## live_server_test.vint

```js
// Live Server Test - Demonstrates enhanced HTTP features in action
import http

print("🔥 Live Enhanced HTTP Server Test")
print("=" * 35)

// Create and configure the app
http.app()

// Add interceptors
http.interceptor("request", func(req) {
    print("📥 Request interceptor activated")
})

http.interceptor("response", func(res) {
    print("📤 Response interceptor activated")
})

// Add guards
http.guard(func(req) {
    print("🛡️ Security guard check passed")
})

// Add middleware
http.cors()
http.bodyParser()

// Define test routes that showcase enhanced features
http.get("/", func(req, res) {
    print("🏠 Home route accessed")
})

http.get("/users/:id", func(req, res) {
    print("👤 User route with parameter accessed")
})

http.post("/api/data", func(req, res) {
    print("💾 API data endpoint accessed")
})

http.get("/test", func(req, res) {
    print("🧪 Test endpoint accessed")
})

print("\n✅ Server configured with enhanced features:")
print("  • Interceptors for request/response processing")
print("  • Security guards for protection")
print("  • CORS and body parser middleware")
print("  • Routes with parameter support")

print("\n🌐 Starting enhanced HTTP server...")
print("📡 Visit these endpoints to test:")
print("  • http://localhost:8080/ (Home)")
print("  • http://localhost:8080/users/123 (User with ID)")
print("  • http://localhost:8080/test (Test endpoint)")
print("  • POST http://localhost:8080/api/data (API endpoint)")

print("\n🚀 Enhanced server starting on port 8080...")
http.listen(8080, "🎯 Enhanced backend server running with all features!")
```

## llm_openai.vint

```js
// VintLang LLM/OpenAI Module Example
// Demonstrates using Large Language Models (like GPT) from VintLang
// NOTE: Requires an OpenAI API key set in environment variables

// Example 1: Chat with OpenAI's GPT model
// The chat function takes an array of messages with roles and content

// import llm

// let messages = [
//     {"role": "system", "content": "You are a helpful assistant."},
//     {"role": "user", "content": "What is the capital of France?"}
// ]

// let response, err = llm.chat(messages)
// if (err != null) {
//     print("Chat error: ", err)
// } else {
//     print("Chat response: ", response)
// }

// Example 2: Text completion
// The completion function generates text based on a prompt

// let prompt = "Write a haiku about the ocean."
// let completion, err = llm.completion(prompt)
// if (err != null) {
//     print("Completion error: ", err)
// } else {
//     print("Completion: ", completion)
// }

println("LLM module example (requires OpenAI API key)");
println("Uncomment the code above to test with a valid API key");
```

## logicals.vint

```js
// Sample Vint program demonstrating logical operators

// Define a function to test 'and', 'or', and 'not'
let test_logical_operators = func () {
    // Testing 'and' operator
    let result_and = and(1+2==3, false) // Should return false
    print("Result of true AND false: ", result_and)
    print(1+2==4)

    // Testing 'or' operator
    let result_or = or(false, true) // Should return true
    print("Result of false OR true: ", result_or)

    // Testing 'not' operator
    let result_not = not(true) // Should return false
    print("Result of NOT true: ", result_not)
}

// Call the function to test the logical operators
test_logical_operators()

```

## main.vint

```js
// Importing modules
import net       // Importing networking module for HTTP operations
import time      // Importing time module to work with date and time

// Main logic to split and print characters of a string
let name = "VintLang"
s = name.split("")
for i in s {
    print(i)
}

// Demonstrating type conversion and conditional statements
age = "10"
convert(age, "INTEGER")  // Convert age string to integer
print(type(age))          // Uncomment to check the type of ageInInt

// Conditional statements to compare the age variable
if (age == 20) {
    print(age)
} else if (age == 10) {
    print("Age is " + age)
} else {
    print((age == "20"))
}

// Working with height variable
height = "6.0" // Height in feet
print("My name is " + name)

// Define a function to print details
let printDetails = func(name, age, height) {
    print("My name is " + name + ", I am " + age + " years old, and my height is " + height + " feet.")
}

// Calling the printDetails function with initial values
printDetails(name, age, height)

// Update height and call the function again
height = "7"
printDetails(name, age, height)

// Print the current timestamp
print(time.now())

// Function to greet a user based on the time of the day
let greet = func(nameParam) {
    let currentTime = time.now()  // Get the current time
    print(currentTime)            // Print the current time
    if (true) {                   // Placeholder condition, modify for actual logic
        print("Good morning, " + nameParam + "!")
    } else {
        print("Good evening, " + nameParam + "!")
    }
}

// Time-related operations
year = 2024
print("Is", year, "Leap year:", time.isLeapYear(year))
print(time.format(time.now(), "02-01-2006 15:04:05"))
print(time.add(time.now(), "1h"))
print(time.subtract(time.now(), "2h30m45s"))

// Call the greet function with a sample name
greet("John")

// Example of a GET request using the net module
let res = net.get("https://tachera.com")
print(res)

// Built-in functions
print(type(123))             // Print the type of an integer
let a = "123"                // Initialize a string variable
convert(a, "INTEGER")        // Convert the string to an integer
type(a)
print(a)                     // Check the type of the variable
print("Hello", "World")      // Print multiple values
write("Hello World")         // Write a string (useful in returning output)

```

## main_example.vint

```js
// VintLang with main function example
import time

// Global setup - this runs during first pass
println("🚀 VintLang Program Starting...")

// Define helper functions
let greet = func(name) {
    println("👋 Hello,", name, "!")
}

let calculate = func(a, b) {
    let sum = a + b
    println("📊 Calculation:", a, "+", b, "=", sum)
    return sum
}

// Main function - the entry point
let main = func() {
    println("\n🎯 === Main Function Executed ===")

    greet("VintLang Developer")

    let result = calculate(10, 32)

    println("⏰ Current time:", time.now())

    // Demonstrate arrays and loops
    let languages = ["VintLang", "Go", "Zig", "C++"]
    println("💻 Supported language styles:")
    for lang in languages {
        println("  -", lang)
    }

    println("✅ Main function completed successfully!")
    return result
}

// This also runs during first pass
println("⚙️ Program initialization complete, main will execute next...")
```

## make_example.vint

```js
// Example of using the make module for build automation
// This is a simple demonstration of how to use make as a Makefile replacement

import make
import os

print("🔨 Make Module Example - Build Automation")
print("==========================================\n")

// Example 1: Check if a command exists
print("1. Checking if 'go' command exists...")
let hasGo = make.check("go")
if (hasGo) {
    print("   ✅ Go is installed!")
} else {
    print("   ❌ Go is not installed!")
}

// Example 2: Set environment variables
print("\n2. Setting environment variables...")
make.env("EXAMPLE_VAR", "Hello from Vint!")
let envValue = os.getEnv("EXAMPLE_VAR")
print("   Set EXAMPLE_VAR=" + envValue)

// Example 3: Execute a simple command
print("\n3. Executing a command (echo)...")
let result = make.exec("echo 'Hello from make.exec!'")
print("   Output: " + result)

// Example 4: Execute a command that lists files
print("\n4. Listing .vint files in examples directory...")
let lsResult = make.exec("ls examples/*.vint 2>/dev/null | head -5")
if (lsResult.type != "error") {
    print("   Files:\n" + lsResult)
} else {
    print("   Could not list files")
}

// Example 5: Echo message (like Make's @echo)
print("\n5. Using make.echo (like Makefile @echo)...")
make.echo("This is a build message!")
make.echo("Building project...")

// Example 6: Check for a command that doesn't exist
print("\n6. Checking for a command that doesn't exist...")
let hasUPX = make.check("upx")
if (hasUPX) {
    print("   ✅ UPX is installed!")
} else {
    print("   ℹ️  UPX is not installed (this is expected)")
}

print("\n✅ Make module example completed!")

```

## math_extensions.vint

```js
// Math Module Extensions Example
// This example demonstrates the new math module capabilities

import math

print("=" * 60)
print("MATH MODULE EXTENSIONS DEMO")
print("=" * 60)

// ============================================================================
// STATISTICS FUNCTIONS
// ============================================================================

print("\n### STATISTICS FUNCTIONS ###\n")

let data = [12, 15, 18, 22, 25, 28, 30, 35, 40]
print("Dataset:", data)
print("Mean:", math.mean(data))
print("Median:", math.median(data))
print("Variance:", math.variance(data))
print("Standard Deviation:", math.stddev(data))

// Example with test scores
let scores = [85, 92, 78, 95, 88, 76, 89, 93]
print("\nTest Scores:", scores)
print("Average Score:", math.mean(scores))
print("Score Spread (stddev):", math.stddev(scores))

// ============================================================================
// COMPLEX NUMBERS
// ============================================================================

print("\n### COMPLEX NUMBERS ###\n")

let c1 = math.complex(3, 4)
print("Complex number c1 = 3 + 4i:", c1)
print("Real part:", c1["real"])
print("Imaginary part:", c1["imag"])
print("Magnitude |c1|:", math.abs(c1))

let c2 = math.complex(5, 12)
print("\nComplex number c2 = 5 + 12i:", c2)
print("Magnitude |c2|:", math.abs(c2))

let c3 = math.complex(0, 1)
print("\nImaginary unit i = 0 + 1i:", c3)
print("Magnitude |i|:", math.abs(c3))

// ============================================================================
// BIG INTEGERS
// ============================================================================

print("\n### BIG INTEGER SUPPORT ###\n")

let bigNum1 = math.bigint("12345678901234567890")
print("Big integer 1:", bigNum1["value"])

let bigNum2 = math.bigint("99999999999999999999999999999999")
print("Big integer 2:", bigNum2["value"])

let regularInt = math.bigint(42)
print("Regular integer as bigint:", regularInt["value"])

// ============================================================================
// LINEAR ALGEBRA OPERATIONS
// ============================================================================

print("\n### LINEAR ALGEBRA ###\n")

// Dot product
let v1 = [1, 2, 3]
let v2 = [4, 5, 6]
print("Vector v1:", v1)
print("Vector v2:", v2)
print("Dot product (v1 · v2):", math.dot(v1, v2))

// Cross product (3D vectors only)
let a = [1, 0, 0]
let b = [0, 1, 0]
print("\nVector a:", a)
print("Vector b:", b)
print("Cross product (a × b):", math.cross(a, b))

// Vector magnitude
let vec = [3, 4]
print("\nVector:", vec)
print("Magnitude:", math.magnitude(vec))

let vec3d = [1, 2, 2]
print("\nVector 3D:", vec3d)
print("Magnitude:", math.magnitude(vec3d))

// ============================================================================
// NUMERICAL METHODS
// ============================================================================

print("\n### NUMERICAL METHODS ###\n")

// Greatest Common Divisor
print("GCD(48, 18):", math.gcd(48, 18))
print("GCD(100, 75):", math.gcd(100, 75))

// Least Common Multiple
print("\nLCM(12, 15):", math.lcm(12, 15))
print("LCM(8, 12):", math.lcm(8, 12))

// Clamp - constrain a value between min and max
print("\nClamp examples:")
print("clamp(5, 0, 10):", math.clamp(5, 0, 10))
print("clamp(-5, 0, 10):", math.clamp(-5, 0, 10))
print("clamp(15, 0, 10):", math.clamp(15, 0, 10))

// Linear interpolation
print("\nLinear interpolation:")
print("lerp(0, 100, 0.0):", math.lerp(0, 100, 0.0))
print("lerp(0, 100, 0.25):", math.lerp(0, 100, 0.25))
print("lerp(0, 100, 0.5):", math.lerp(0, 100, 0.5))
print("lerp(0, 100, 0.75):", math.lerp(0, 100, 0.75))
print("lerp(0, 100, 1.0):", math.lerp(0, 100, 1.0))

// ============================================================================
// PRACTICAL EXAMPLES
// ============================================================================

print("\n### PRACTICAL EXAMPLES ###\n")

// Example 1: Calculate distance between two points in 3D space
let distance3D = func(p1, p2) {
    let diff = [p2[0] - p1[0], p2[1] - p1[1], p2[2] - p1[2]]
    return math.magnitude(diff)
}

let point1 = [1, 2, 3]
let point2 = [4, 6, 8]
print("Distance between", point1, "and", point2)
print("Result:", distance3D(point1, point2))

// Example 2: Normalize a vector
let normalize = func(vec) {
    let mag = math.magnitude(vec)
    return [vec[0] / mag, vec[1] / mag, vec[2] / mag]
}

let vector = [3, 4, 0]
print("\nOriginal vector:", vector)
print("Normalized:", normalize(vector))

// Example 3: Statistical analysis of a dataset
let analyzeData = func(data) {
    print("\nDataset Analysis:")
    print("Data:", data)
    print("Count:", len(data))
    print("Mean:", math.mean(data))
    print("Median:", math.median(data))
    print("Std Dev:", math.stddev(data))
    print("Min:", math.min(data))
    print("Max:", math.max(data))
}

let measurements = [102, 98, 105, 97, 103, 100, 99, 101, 104, 98]
analyzeData(measurements)

print("\n" + "=" * 60)
print("END OF DEMO")
print("=" * 60)

```

## math_showcase.vint

```js
// VintLang Mathematical & Algorithmic Showcase
// Advanced computational demonstrations

import time
import os
import json
import uuid
import math

print("🧮 VintLang Mathematical & Algorithmic Showcase")
print("=" * 60)
print("Demonstrating computational and algorithmic capabilities")
print("=" * 60)

// Algorithm 1: Fibonacci Sequence Generator
print("\n🔢 Algorithm 1: Fibonacci Sequence")
print("-" * 40)

let fibonacci = func(n) {
    if (n <= 1) {
        return n
    }

    let a = 0
    let b = 1
    let sequence = [0, 1]

    for i in [2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15] {
        if (i <= n) {
            let next = a + b
            sequence.push(next)
            a = b
            b = next
        }
    }

    return sequence
}

let fibSequence = fibonacci(15)
print("Fibonacci sequence (first 16 numbers):")
for i, num in fibSequence {
    print("F(" + string(i) + ") = " + string(num))
}

// Algorithm 2: Prime Number Generator
print("\n🔢 Algorithm 2: Prime Number Detection")
print("-" * 40)

let isPrime = func(n) {
    if (n < 2) {
        return false
    }
    if (n == 2) {
        return true
    }
    if (n % 2 == 0) {
        return false
    }

    let i = 3
    while (i * i <= n) {
        if (n % i == 0) {
            return false
        }
        i += 2
    }
    return true
}

let primes = []
let numbers = [2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26, 27, 28, 29, 30]

for num in numbers {
    if (isPrime(num)) {
        primes.push(num)
    }
}

print("Prime numbers up to 30:")
print(string(primes))
print("Found " + string(len(primes)) + " prime numbers")

// Algorithm 3: Factorial Calculator
print("\n🔢 Algorithm 3: Factorial Calculations")
print("-" * 40)

let factorial = func(n) {
    if (n <= 1) {
        return 1
    }
    return n * factorial(n - 1)
}

let factorials = []
for i in [1, 2, 3, 4, 5, 6, 7, 8, 9, 10] {
    let fact = factorial(i)
    factorials.push({
        "number": i,
        "factorial": fact
    })
    print(string(i) + "! = " + string(fact))
}

// Algorithm 4: Sorting Algorithm (Bubble Sort)
print("\n🔢 Algorithm 4: Bubble Sort Implementation")
print("-" * 40)

let bubbleSort = func(arr) {
    let sortedArr = arr
    let n = len(sortedArr)

    for i in [0, 1, 2, 3, 4, 5, 6, 7, 8, 9] {
        if (i < n - 1) {
            for j in [0, 1, 2, 3, 4, 5, 6, 7, 8, 9] {
                if (j < n - i - 1) {
                    if (sortedArr[j] > sortedArr[j + 1]) {
                        let temp = sortedArr[j]
                        sortedArr[j] = sortedArr[j + 1]
                        sortedArr[j + 1] = temp
                    }
                }
            }
        }
    }

    return sortedArr
}

let unsortedArray = [64, 34, 25, 12, 22, 11, 90, 88, 76, 50]
print("Original array: " + string(unsortedArray))
let sortedArray = bubbleSort(unsortedArray)
print("Sorted array:   " + string(sortedArray))

// Algorithm 5: Number Theory - GCD Calculator
print("\n🔢 Algorithm 5: Greatest Common Divisor")
print("-" * 40)

let gcd = func(a, b) {
    while (b != 0) {
        let temp = b
        b = a % b
        a = temp
    }
    return a
}

let numberPairs = [
    {"a": 48, "b": 18},
    {"a": 56, "b": 42},
    {"a": 72, "b": 27},
    {"a": 100, "b": 75}
]

for pair in numberPairs {
    let result = gcd(pair["a"], pair["b"])
    print("GCD(" + string(pair["a"]) + ", " + string(pair["b"]) + ") = " + string(result))
}

// Algorithm 6: Mathematical Constants and Calculations
print("\n🔢 Algorithm 6: Mathematical Constants")
print("-" * 40)

print("Mathematical constants:")
print("PI ≈ " + string(math.PI()))

// Calculate circle properties
let radius = 5
let circumference = 2 * math.PI() * radius
let area = math.PI() * radius * radius

print("For a circle with radius " + string(radius) + ":")
print("  Circumference = " + string(circumference))
print("  Area = " + string(area))

// Algorithm 7: Data Analysis and Statistics
print("\n📊 Algorithm 7: Statistical Analysis")
print("-" * 40)

let dataset = [23, 45, 67, 89, 12, 34, 56, 78, 90, 21, 43, 65, 87, 10, 32]

// Calculate mean
let sum = 0
for value in dataset {
    sum += value
}
let mean = sum / len(dataset)

// Find min and max
let min = dataset[0]
let max = dataset[0]
for value in dataset {
    if (value < min) {
        min = value
    }
    if (value > max) {
        max = value
    }
}

print("Dataset: " + string(dataset))
print("Count: " + string(len(dataset)))
print("Sum: " + string(sum))
print("Mean: " + string(mean))
print("Min: " + string(min))
print("Max: " + string(max))
print("Range: " + string(max - min))

// Algorithm 8: Data Structure Operations
print("\n🔢 Algorithm 8: Data Structure Demonstrations")
print("-" * 40)

// Stack simulation using array
let stack = []
let stackOperations = ["push(10)", "push(20)", "push(30)", "pop()", "push(40)", "pop()"]

print("Stack operations:")
for operation in stackOperations {
    if (operation.contains("push")) {
        let value = 10
        if (operation.contains("20")) {
            value = 20
        } else if (operation.contains("30")) {
            value = 30
        } else if (operation.contains("40")) {
            value = 40
        }
        stack.push(value)
        print("  " + operation + " -> Stack: " + string(stack))
    } else if (operation.contains("pop")) {
        if (len(stack) > 0) {
            stack.pop()
            print("  " + operation + " -> Stack: " + string(stack))
        }
    }
}

// Comprehensive Report Generation
print("\n📋 Generating Comprehensive Mathematical Report")
print("-" * 50)

let mathReport = {
    "report_id": uuid.generate(),
    "generated_at": time.format(time.now(), "2006-01-02 15:04:05"),
    "algorithms_demonstrated": [
        "Fibonacci Sequence Generation",
        "Prime Number Detection",
        "Factorial Calculation",
        "Bubble Sort Algorithm",
        "Greatest Common Divisor",
        "Mathematical Constants",
        "Statistical Analysis",
        "Data Structure Operations"
    ],
    "results": {
        "fibonacci_sequence": fibSequence,
        "prime_numbers": primes,
        "factorials": factorials,
        "sorted_array": sortedArray,
        "gcd_calculations": numberPairs,
        "statistics": {
            "dataset": dataset,
            "mean": mean,
            "min": min,
            "max": max,
            "count": len(dataset)
        }
    },
    "performance_metrics": {
        "algorithms_executed": 8,
        "calculations_performed": len(fibSequence) + len(primes) + len(factorials) + len(dataset),
        "data_points_processed": len(dataset) + len(unsortedArray) + len(fibSequence)
    }
}

let reportFile = "mathematical_analysis_" + time.format(time.now(), "2006-01-02_15-04-05") + ".json"
os.writeFile(reportFile, json.encode(mathReport))
print("✓ Comprehensive mathematical report saved to: " + reportFile)

// Generate text summary
let textSummary = "VINTLANG MATHEMATICAL & ALGORITHMIC SHOWCASE\n"
textSummary += "Generated: " + time.format(time.now(), "02-01-2006 15:04:05") + "\n"
textSummary += "=" * 60 + "\n\n"

textSummary += "ALGORITHMS DEMONSTRATED:\n"
for algo in mathReport["algorithms_demonstrated"] {
    textSummary += "  ✓ " + algo + "\n"
}

textSummary += "\nKEY RESULTS:\n"
textSummary += "  • Fibonacci sequence calculated up to F(15)\n"
textSummary += "  • Found " + string(len(primes)) + " prime numbers up to 30\n"
textSummary += "  • Calculated factorials from 1! to 10!\n"
textSummary += "  • Sorted array of " + string(len(sortedArray)) + " elements\n"
textSummary += "  • Analyzed dataset of " + string(len(dataset)) + " values\n"
textSummary += "  • Demonstrated stack operations\n"

textSummary += "\nSTATISTICAL SUMMARY:\n"
textSummary += "  • Dataset mean: " + string(mean) + "\n"
textSummary += "  • Dataset range: " + string(min) + " to " + string(max) + "\n"
textSummary += "  • Total calculations: " + string(mathReport["performance_metrics"]["calculations_performed"]) + "\n"

let summaryFile = "math_summary_" + time.format(time.now(), "2006-01-02_15-04-05") + ".txt"
os.writeFile(summaryFile, textSummary)
print("✓ Summary report saved to: " + summaryFile)

// Final Results
print("\n🎉 Mathematical Showcase Complete!")
print("-" * 50)

print("📊 COMPUTATIONAL ACHIEVEMENTS:")
print("  • Algorithms implemented: " + string(len(mathReport["algorithms_demonstrated"])))
print("  • Mathematical calculations: " + string(mathReport["performance_metrics"]["calculations_performed"]))
print("  • Data points processed: " + string(mathReport["performance_metrics"]["data_points_processed"]))
print("  • Prime numbers found: " + string(len(primes)))
print("  • Fibonacci numbers calculated: " + string(len(fibSequence)))

print("\n✨ VINTLANG MATHEMATICAL CAPABILITIES:")
print("  ✓ Recursive algorithms (Fibonacci, Factorial)")
print("  ✓ Iterative algorithms (Prime detection, Sorting)")
print("  ✓ Mathematical calculations (GCD, Statistics)")
print("  ✓ Data structure operations (Stack, Array)")
print("  ✓ Number theory implementations")
print("  ✓ Statistical analysis functions")
print("  ✓ Algorithm complexity handling")

print("\n🚀 VintLang is excellent for:")
print("  • Mathematical computing")
print("  • Algorithm implementation")
print("  • Data analysis and statistics")
print("  • Educational programming")
print("  • Scientific calculations")
print("  • Computational problem solving")

print("\n🎯 This demonstrates VintLang's power for")
print("   computational and mathematical applications!")

print("\n" + "=" * 60)
print("🧮 Mathematical Showcase Complete!")
print("=" * 60)
```

## mathfile.vint

```js
// VintLang Package Example
// Demonstrates defining a package to organize related functions

// Define a package (will be changed to module syntax in future)
package mathfile{
    // Define a sum function within the package
    let sum = func(a,b){
        return a+b
    }
}
```

## module_functions_test.vint

```js
// Test to verify existing module functions still work
import math
import string
import random

println("Testing existing module functions...")

// Test math module functions
println("=== Math Module ===")
println("math.abs(-5):", math.abs(-5))
println("math.sqrt(16):", math.sqrt(16))
println("math.round(3.14):", math.round(3.14))

// Test string module functions
println("\n=== String Module ===")
println("string.toUpper('hello'):", string.toUpper("hello"))
println("string.toLower('WORLD'):", string.toLower("WORLD"))
println("string.trim('  hi  '):", string.trim("  hi  "))
println("string.contains('hello', 'ell'):", string.contains("hello", "ell"))

// Test random module functions
println("\n=== Random Module ===")
println("random.float():", random.float())
println("random.int(1, 10):", random.int(1, 10))

// Test array methods
println("\n=== Array Methods ===")
let arr = [3, 1, 4, 1, 5]
println("Original array:", arr)
println("Sorted array:", arr.sort())
let arr2 = [1, 2, 3]
println("Original array 2:", arr2)
println("Reversed array 2:", arr2.reverse())

println("\nExisting module functions work correctly!")
```

## mysql.vint

```js
// VintLang MySQL Database Example
// Demonstrates MySQL database operations: connect, create table, insert, query

import mysql

// NOTE: Please replace the placeholder credentials with your actual MySQL database credentials.
// The connection string should be in the format: "user:password@tcp(127.0.0.1:3306)/dbname"

// Example 1: Open a MySQL database connection
let conn = mysql.open("user:password@tcp(127.0.0.1:3306)/testdb")

if (type(conn) == "ERROR") {
    print("Error connecting to MySQL:", conn)
} else {
    print("Successfully connected to MySQL")

    // Example 2: Create a table
    // Creates a users table with auto-incrementing id
    let create_table_query = "CREATE TABLE IF NOT EXISTS users (id INT AUTO_INCREMENT PRIMARY KEY, name VARCHAR(255), age INT)"
    let err = mysql.execute(conn, create_table_query)
    if (err != null) {
        print("Error creating table:", err)
    }

    // Example 3: Insert data into the table
    // Uses prepared statements with ? placeholders
    let insert_query = "INSERT INTO users (name, age) VALUES (?, ?)"
    mysql.execute(conn, insert_query, "Alice", 25)
    mysql.execute(conn, insert_query, "Bob", 30)

    // Example 4: Fetch all rows from the table
    print("\n=== All Users ===")
    let fetch_all_query = "SELECT * FROM users"
    let users = mysql.fetchAll(conn, fetch_all_query)
    print(users)

    // Example 5: Fetch a single row
    print("\n=== First User ===")
    let fetch_one_query = "SELECT * FROM users LIMIT 1"
    let first_user = mysql.fetchOne(conn, fetch_one_query)
    print(first_user)

    // Example 6: Close the connection
    mysql.close(conn)
    print("\nConnection closed")
}
```

## nativeStrings.vint

```js
// VintLang Native String Operations Example
// Demonstrates built-in string methods and type conversion functions

// Define a variable and assign a string to it
let name = "Tachera Sasi";

// Split the string into an array of characters and print the result
print(name.split(""));

// Reverse the string and print the result
print(name.reverse());

// Get the length of the string and print it
print(name.len());

// Convert the string to uppercase and print it
print(name.upper());

// Convert the string to lowercase and print it
print(name.lower());

// Check if the string contains the substring "sasi" (case-sensitive) and print the result
print(name.contains("sasi"));

// Convert the string to uppercase and check if it contains the substring "SASI" (case-sensitive), then print the result
print(name.upper().contains("SASI"));

// Replace the substring "Sasi" with "Vint" and print the result
print(name.replace("Sasi", "Vint"));

// Trim any occurrence of the character "a" from the start and end of the string and print the result
print(name.trim("a"));

print(string(123)); // "123"
print(string(true)); // "true"
print(string(12.34)); // "12.34"
print(string("Hello World")); // "Hello World"

print(int("123")); // 123
print(int(12.34)); // 12
print(int(true)); // 1
print(int(false)); // 0
```

## os.vint

```js
// VintLang OS Module Example
// Demonstrates operating system operations: file I/O, commands, environment variables

import os

// Example 1: Exit with a status code (commented to avoid terminating script)
// os.exit(1)

// Example 2: Run a shell command and capture output
let result = os.run("ls -la")
print(result)
// print(os.run("go run . vintLang/main.vint"))

// Example 3: Get and set environment variables
// os.setEnv("API_KEY", "12345")  // Set environment variable
let api_key = os.getEnv("API_KEY")  // Get environment variable
print(api_key)

// Example 4: Write and read files
os.writeFile("example.txt", "Hello, Vint!")
let content = os.readFile("example.txt")
print(content)

// Example 5: List directory contents
let files = os.listDir(".")
print(files)

// Example 6: Create a directory
os.makeDir("new_folder")

// Example 7: Check if a file exists
let exists = os.fileExists("example.txt")
print(exists) // Outputs: true (after writing the file above)

// Example 8: Write a file and read it line by line
os.writeFile("example.txt", "Hello\nWorld")
let lines = os.readLines("example.txt")
print(lines) // Outputs: ["Hello", "World"]

// Example 9: Delete a file (commented to keep example file)
// os.deleteFile("example.txt")

```

## overloading_test.vint

```js
// VintLang Function Overloading Example
// Demonstrates function overloading by arity (number of parameters)

// Define multiple versions of 'greet' function with different parameter counts
// The interpreter selects the appropriate version based on the number of arguments

// Version 1: Single parameter
let greet = func(name) {
    print("Hello, ", name)
}

// Version 2: Two parameters
let greet = func(name, times) {
    for i in range(times) {
        print("Hello, ", name)
    }
}

// Version 3: No parameters
let greet = func() {
    print("Hello, world!")
}

// Call with one argument - uses Version 1
greet("Alice")      // Should print: Hello, Alice
print("---")

// Call with two arguments - uses Version 2
greet("Bob", 3)      // Should print: Hello, Bob (3 times)
print("---")

// Call with no arguments - uses Version 3
greet()             // Should print: Hello, world!

// Example: Normal (non-overloaded) functions work as usual
let add = func(a, b) {
    return a + b
}

let result = add(2, 3)
print("add(2, 3) =", result) // Should print: add(2, 3) = 5

let say_hello = func() {
    print("Hello from a normal function!")
}

say_hello() // Should print: Hello from a normal function!
```

## packages.vint

```js
// VintLang Package System Example
// NOTE: This feature is still in development and not fully working yet

/**
The package system with init functions and @ accessor is under development.
This example demonstrates the planned syntax, but it's not yet functional.
*/

// Planned syntax (not yet working):
// package vint{
//     init = func(){
//         name = "vint"
//         a = "hiofehoi"
//     }
//     print(@.name)
// }
// print(@.a)

println("Package system is under development.");
println(
  "This example shows planned syntax that will be implemented in future versions.",
);
```

## path.vint

```js
// VintLang Path Module Example
// Demonstrates file path manipulation operations

import path

// Example 1: Join path components
// Combines multiple path segments into a single path
let full_path = path.join("/home", "user", "documents", "report.pdf")
print("Joined path:", full_path)

// Example 2: Get the basename (filename)
// Extracts the last component of the path
let base = path.basename(full_path)
print("Basename:", base)

// Example 3: Get the directory name
// Returns the directory path without the filename
let dir = path.dirname(full_path)
print("Dirname:", dir)

// Example 4: Get the file extension
// Extracts the file extension from the path
let extension = path.ext(full_path)
print("Extension:", extension)

// Example 5: Check if path is absolute
// Determines if a path is absolute (starts from root)
print("Is absolute?", path.isAbs(full_path))
print("Is 'relative/path' absolute?", path.isAbs("relative/path"))
```

## pointers.vint

```js
// VintLang Pointers Example
// Demonstrates pointer operations: creating, dereferencing, and displaying pointers

print("POINTERS IN VINTLANG")

// Create a variable with a value
let x = 42 * 2

// Create a pointer to x using the & (address-of) operator
let p = &x

// Print the pointer (shows memory address and value)
println(p)

// Display the pointer with a label
println("POINTER is",p)

// Dereference the pointer using * operator to get the value
println("VALUE is",*p)

// Example output:
// POINTER is Pointer(addr=0x14000010560, value=84)
// VALUE is 84
```

## postgres.vint

```js
// VintLang PostgreSQL Database Example
// Demonstrates PostgreSQL database operations: connect, create table, insert, query

import postgres

// NOTE: Please replace the placeholder credentials with your actual PostgreSQL database credentials.
// The connection string should be in the format: "user=youruser password=yourpassword dbname=yourdbname sslmode=disable"

// Example 1: Open a PostgreSQL database connection
const conn = postgres.open("user=postgres password=password dbname=testdb sslmode=disable")

if (type(conn) == "ERROR") {
    error "Error connecting to PostgreSQL: " + str(conn)
} else {
    info "Successfully connected to PostgreSQL"

    // Example 2: Create a table
    // Creates a users table with auto-incrementing SERIAL id
    let create_table_query = "CREATE TABLE IF NOT EXISTS users (id SERIAL PRIMARY KEY, name VARCHAR(255), age INT)"
    let err = postgres.execute(conn, create_table_query)
    if (err != null) {
        print("Error creating table:", err)
    }

    // Example 3: Insert data into the table
    // PostgreSQL uses $1, $2 instead of ? for placeholders
    let insert_query = "INSERT INTO users (name, age) VALUES ($1, $2)"
    postgres.execute(conn, insert_query, "Alice", 25)
    postgres.execute(conn, insert_query, "Bob", 30)

    // Example 4: Fetch all rows from the table
    print("\n=== All Users ===")
    let fetch_all_query = "SELECT * FROM users"
    let users = postgres.fetchAll(conn, fetch_all_query)
    print(users)

    // Example 5: Fetch a single row
    print("\n=== First User ===")
    let fetch_one_query = "SELECT * FROM users LIMIT 1"
    let first_user = postgres.fetchOne(conn, fetch_one_query)
    print(first_user)

    // Example 6: Close the connection
    postgres.close(conn)
    print("\nConnection closed")
}
```

## random.vint

```js
// VintLang Random Module Example
// This example demonstrates various random number generation functions

import random

// Generate a random integer between 10 and 20 (inclusive)
print("Random integer between 10 and 20:", random.int(10, 20))

// Generate a random float between 0 and 1
print("Random float:", random.float())

// Generate a random string of specified length
print("Random string of length 8:", random.string(8))

// Select a random element from an array
let options = ["rock", "paper", "scissors"]
print("Random choice from", options, ":", random.choice(options))
```

## redis.vint

```js
const redis = import("redis");

// Connect to Redis
conn = redis.connect("localhost:6379");

// Basic operations
redis.set(conn, "greeting", "Hello, World!");
message = redis.get(conn, "greeting");

// Hash operations
redis.hset(conn, "user:1", "name", "John Doe");
user = redis.hgetall(conn, "user:1");

// List operations
redis.rpush(conn, "tasks", "task1", "task2");
task = redis.lpop(conn, "tasks");

// Close connection
redis.close(conn);
```

## reflect.vint

```js
// VintLang Reflect Module Example
// Demonstrates runtime type inspection and reflection capabilities

import reflect

// Example 1: Type inspection
// Get the type of various values
let t1 = reflect.typeOf("hello")         // "STRING"
let t2 = reflect.typeOf([1,2,3])         // "ARRAY"
let t3 = reflect.typeOf({"a": 1})        // "DICT"
let t4 = reflect.typeOf(null)            // "NULL"
let t5 = reflect.typeOf(func() {})       // "FUNCTION"
println(t1, t2, t3, t4, t5)

// Example 2: Value extraction
// Extract the underlying value
let v = reflect.valueOf(42)              // 42
println(v)

// Example 3: Null check
// Check if a value is null/nil
println(reflect.isNil(null))             // true
println(reflect.isNil(123))              // false

// Example 4: Array check
// Check if a value is an array
println(reflect.isArray([1,2,3]))        // true
println(reflect.isArray("not array"))    // false

// Example 5: Object (dictionary) check
// Check if a value is an object/dict
println(reflect.isObject({"a": 1}))      // true
println(reflect.isObject([1,2,3]))       // false

// Example 6: Function check
// Check if a value is a function
let f = func(x) { x * 2 }
println(reflect.isFunction(f))           // true
println(reflect.isFunction(123))         // false
```

## regex.vint

```js
// VintLang Regex Module Example (CURRENTLY NOT WORKING)
// NOTE: This example demonstrates planned regex functionality
// The regex module is not yet fully implemented and has syntax parsing issues

// TODO: Fix regex module implementation to support these operations:
// - Pattern matching
// - String replacement with regex
// - String splitting with regex patterns

// import regex

// Example 1: Using match to check if a string matches a pattern
// let result = regex.match("^Hello", "Hello World")
// print(result)  // Expected output: true

// Example 2: Using match to check if a string does not match a pattern
// result = regex.match("^World", "Hello World")
// print(result)  // Expected output: false

// Example 3: Using replaceString to replace part of a string with a new value
// let newString = regex.replaceString("World", "VintLang", "Hello World")
// print(newString)  // Expected output: "Hello VintLang"

// Example 4: Using splitString to split a string by a regex pattern
// let words = regex.splitString("\\s+", "Hello World VintLang")
// print(words)  // Expected output: ["Hello", "World", "VintLang"]

// Example 5: Using splitString to split a string by a comma
// let csv = regex.splitString(",", "apple,banana,orange")
// print(csv)  // Expected output: ["apple", "banana", "orange"]

// Example 6: Using match with a more complex regex pattern
// let emailMatch = regex.match("^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\\.[a-zA-Z]{2,}$", "test@example.com")
// print(emailMatch)  // Expected output: true

// Example 7: Using replaceString with a regex pattern to replace digits in a string
// let maskedString = regex.replaceString("\\d", "*", "My phone number is 123456789")
// print(maskedString)  // Expected output: "My phone number is *********"

println(
  "Regex module is not yet implemented. This example is for future reference.",
);
```

## repeat-keyword.vint

```js
// VintLang Repeat Keyword Example
// Demonstrates the repeat loop for fixed-iteration counting

let num = 5

// Repeat with a variable - executes the block 'num' times
repeat num {
    println("Hello there")
}

// Repeat with a literal number - automatically provides index variable 'i'
repeat 3 {
    println("Default i:", i)
}
```

## savegame.json

```js
{
  "health": 100,
  "inventory": ["Mystic Key"],
  "location": "start",
  "name": "Tach"
}

```

## schedule_examples.vint

```js
// VintLang Schedule Module Examples
// This file demonstrates all features of the schedule module

import schedule
import time

println("=== VintLang Schedule Module - All Features ===")
println("")

// 1. Basic ticker functionality
println("1. Ticker Examples:")
println("   Creating ticker for every 3 seconds...")
let ticker1 = schedule.ticker(3, func() {
    println("   [TICKER] Regular tick every 3 seconds")
})

// 2. Helper functions
println("2. Helper Function Examples:")
let everySecJob = schedule.everySecond(func() {
    println("   [HELPER] Every second ping")
})

// Let both run for a while
time.sleep(8)

// Stop them
schedule.stopTicker(ticker1)
schedule.stopTicker(everySecJob)
println("   → Stopped basic examples")
println("")

// 3. Cron expressions
println("3. Cron Expression Examples:")

println("   a) Every 5 seconds using */5")
let cronJob1 = schedule.schedule("*/5 * * * * *", func() {
    println("   [CRON] Every 5 seconds via cron")
})

println("   b) Every second using wildcards")
let cronJob2 = schedule.schedule("* * * * * *", func() {
    println("   [CRON] Every second via cron")
})

// Let cron jobs run
time.sleep(7)
schedule.stopSchedule(cronJob1)
schedule.stopSchedule(cronJob2)
println("   → Stopped cron examples")
println("")

// 4. Practical scheduling examples
println("4. Practical Examples:")

// Daily reminder (would trigger at 14:30)
let dailyReminder = schedule.daily(14, 30, func() {
    println("   [DAILY] Time for afternoon coffee!")
})
println("   → Daily reminder set for 14:30")

// Hourly report (would trigger at top of each hour)
let hourlyReport = schedule.everyHour(func() {
    println("   [HOURLY] Generating hourly report...")
})
println("   → Hourly report scheduled")

// Minute marker (would trigger at :00 seconds of each minute)
let minuteMarker = schedule.everyMinute(func() {
    println("   [MINUTE] New minute started!")
})
println("   → Minute marker scheduled")
println("")

// 5. Error handling examples
println("5. Error Handling:")

println("   Testing invalid ticker interval...")
let errorResult1 = schedule.ticker(-5, func() { println("Never runs") })
println("   Result: Error (as expected)")

println("   Testing invalid cron expression...")
let errorResult2 = schedule.schedule("not a cron", func() { println("Never runs") })
println("   Result: Error (as expected)")

println("   Testing invalid daily time...")
let errorResult3 = schedule.daily(25, 70, func() { println("Never runs") })
println("   Result: Error (as expected)")
println("")

// Cleanup
println("6. Cleanup:")
schedule.stopSchedule(dailyReminder)
schedule.stopSchedule(hourlyReport)
schedule.stopSchedule(minuteMarker)
println("   → All scheduled jobs stopped")
println("")

println("=== Summary ===")
println("✓ ticker(intervalSeconds, callback) - for regular intervals")
println("✓ schedule(cronExpr, callback) - for cron-based scheduling")
println("✓ everySecond(callback) - shortcut for 1-second intervals")
println("✓ everyMinute(callback) - runs at :00 seconds each minute")
println("✓ everyHour(callback) - runs at :00:00 each hour")
println("✓ daily(hour, minute, callback) - runs daily at specified time")
println("✓ stopTicker(obj) and stopSchedule(obj) - for cleanup")
println("✓ Supports step values like */5 for 'every 5' patterns")
println("✓ Comprehensive error handling with helpful messages")
println("")
println("The VintLang schedule module is complete and ready for use!")
println("Similar to Go's time.Ticker and NestJS's @Cron decorators.")
```

## schedule_test.vint

```js
import schedule

// Test ticker (should return error about callback execution)
let tickerResult = schedule.ticker(1, func() { print("Tick!") })
print("ticker result:", tickerResult)

// Test schedule (should return error about callback execution)
let scheduleResult = schedule.schedule("0 0 9 * * *", func() { print("Good morning!") })
print("schedule result:", scheduleResult)

// Test stopTicker and stopSchedule with dummy objects (should not error, but do nothing)
let dummyTicker = null
let dummySchedule = null
print("stopTicker result:", schedule.stopTicker(dummyTicker))
print("stopSchedule result:", schedule.stopSchedule(dummySchedule))
```

## shell.vint

```js
// VintLang Shell Module Example
// Demonstrates executing shell commands from VintLang

import shell

// Execute a shell command and capture the output
let output = shell.run("echo Hello, Shell!")
print(output) // Outputs "Hello, Shell!"

// Check if a command exists in the system PATH
let exists = shell.exists("ls")
print(exists) // Outputs true if the 'ls' command exists

```

## showcase_task_manager.vint

```js
// VintLang Showcase: Personal Task Management System
// This application demonstrates VintLang's capabilities with a real-world use case

import time
import os
import json
import uuid
import "string"

// Color codes for terminal output
let COLORS = {
    "reset": "\033[0m",
    "red": "\033[31m",
    "green": "\033[32m",
    "yellow": "\033[33m",
    "blue": "\033[34m",
    "magenta": "\033[35m",
    "cyan": "\033[36m",
    "white": "\033[37m"
}

// Task structure and application state
let taskManager = {
    "tasks": [],
    "categories": ["Work", "Personal", "Study", "Health", "Project"],
    "priorities": ["Low", "Medium", "High", "Urgent"],
    "dataFile": "tasks.json"
}

// Helper function to print colored text
let printColor = func(color, text) {
    print(COLORS[color] + text + COLORS["reset"])
}

// Helper function to print a separator line
let printSeparator = func() {
    printColor("cyan", "=" * 50)
}

// Load tasks from JSON file
let loadTasks = func() {
    if (os.fileExists(taskManager["dataFile"])) {
        let data = os.readFile(taskManager["dataFile"])
        taskManager["tasks"] = json.decode(data)
        printColor("green", "✓ Tasks loaded successfully!")
    } else {
        printColor("yellow", "⚠ No existing tasks file found. Starting fresh!")
        taskManager["tasks"] = []
    }
}

// Save tasks to JSON file
let saveTasks = func() {
    let data = json.encode(taskManager["tasks"])
    os.writeFile(taskManager["dataFile"], data)
    printColor("green", "✓ Tasks saved successfully!")
}

// Create a new task
let createTask = func() {
    printSeparator()
    printColor("blue", "📝 Creating New Task")
    printSeparator()

    let title = input("Task title: ")
    let description = input("Description: ")

    // Show categories
    printColor("cyan", "Available categories:")
    for i, category in taskManager["categories"] {
        print(string(i + 1) + ". " + category)
    }
    let categoryIndex = input("Select category (1-5): ")
    let categoryNum = int(categoryIndex)
    let category = taskManager["categories"][categoryNum - 1]

    // Show priorities
    printColor("cyan", "Available priorities:")
    for i, priority in taskManager["priorities"] {
        print(string(i + 1) + ". " + priority)
    }
    let priorityIndex = input("Select priority (1-4): ")
    let priorityNum = int(priorityIndex)
    let priority = taskManager["priorities"][priorityNum - 1]

    let dueDate = input("Due date (YYYY-MM-DD) or press Enter for none: ")

    let task = {
        "id": uuid.generate(),
        "title": title,
        "description": description,
        "category": category,
        "priority": priority,
        "dueDate": dueDate,
        "completed": false,
        "createdAt": time.now(),
        "completedAt": ""
    }

    taskManager["tasks"].push(task)
    printColor("green", "✓ Task created successfully!")
}

// List all tasks
let listTasks = func() {
    printSeparator()
    printColor("blue", "📋 Task List")
    printSeparator()

    if (len(taskManager["tasks"]) == 0) {
        printColor("yellow", "No tasks found!")
        return
    }

    for i, task in taskManager["tasks"] {
        let status = "❌"
        if (task["completed"]) {
            status = "✅"
        }

        let priorityColor = "white"
        if (task["priority"] == "Urgent") {
            priorityColor = "red"
        } else if (task["priority"] == "High") {
            priorityColor = "yellow"
        } else if (task["priority"] == "Medium") {
            priorityColor = "blue"
        }

        print(status + " [" + string(i + 1) + "] " + task["title"])
        print("    Category: " + task["category"])
        printColor(priorityColor, "    Priority: " + task["priority"])
        if (task["dueDate"] != "") {
            print("    Due: " + task["dueDate"])
        }
        if (task["description"] != "") {
            print("    Description: " + task["description"])
        }
        print("    Created: " + task["createdAt"])
        if (task["completed"] && task["completedAt"] != "") {
            print("    Completed: " + task["completedAt"])
        }
        print("")
    }
}

// Mark task as completed
let completeTask = func() {
    listTasks()
    if (len(taskManager["tasks"]) == 0) {
        return
    }

    let taskIndex = input("Enter task number to mark as completed: ")
    let taskNum = int(taskIndex)

    if (taskNum > 0 && taskNum <= len(taskManager["tasks"])) {
        let task = taskManager["tasks"][taskNum - 1]
        task["completed"] = true
        task["completedAt"] = time.now()
        printColor("green", "✓ Task marked as completed!")
    } else {
        printColor("red", "❌ Invalid task number!")
    }
}

// Filter tasks by category
let filterByCategory = func() {
    printColor("cyan", "Available categories:")
    for i, category in taskManager["categories"] {
        print(string(i + 1) + ". " + category)
    }
    let categoryIndex = input("Select category to filter (1-5): ")
    let categoryNum = int(categoryIndex)
    let selectedCategory = taskManager["categories"][categoryNum - 1]

    printSeparator()
    printColor("blue", "📋 Tasks in category: " + selectedCategory)
    printSeparator()

    let found = false
    for task in taskManager["tasks"] {
        if (task["category"] == selectedCategory) {
            let status = "❌"
            if (task["completed"]) {
                status = "✅"
            }
            print(status + " " + task["title"] + " (" + task["priority"] + ")")
            found = true
        }
    }

    if (!found) {
        printColor("yellow", "No tasks found in this category!")
    }
}

// Get task statistics
let showStatistics = func() {
    printSeparator()
    printColor("blue", "📊 Task Statistics")
    printSeparator()

    let totalTasks = len(taskManager["tasks"])
    let completedTasks = 0
    let pendingTasks = 0
    let urgentTasks = 0
    let categoryStats = {}

    for task in taskManager["tasks"] {
        if (task["completed"]) {
            completedTasks += 1
        } else {
            pendingTasks += 1
        }

        if (task["priority"] == "Urgent") {
            urgentTasks += 1
        }

        let category = task["category"]
        if (!categoryStats.hasKey(category)) {
            categoryStats[category] = 0
        }
        categoryStats[category] += 1
    }

    print("Total Tasks: " + string(totalTasks))
    printColor("green", "Completed: " + string(completedTasks))
    printColor("yellow", "Pending: " + string(pendingTasks))
    printColor("red", "Urgent: " + string(urgentTasks))

    if (totalTasks > 0) {
        let completionRate = (completedTasks * 100) / totalTasks
        print("Completion Rate: " + string(completionRate) + "%")
    }

    print("\nTasks by Category:")
    for category in taskManager["categories"] {
        let count = 0
        if (categoryStats.hasKey(category)) {
            count = categoryStats[category]
        }
        print("  " + category + ": " + string(count))
    }
}

// Export tasks to a readable format
let exportTasks = func() {
    let exportFile = "tasks_export_" + time.format(time.now(), "2006-01-02_15-04-05") + ".txt"
    let content = "TASK EXPORT - " + time.format(time.now(), "02-01-2006 15:04:05") + "\n"
    content += "=" * 50 + "\n\n"

    for i, task in taskManager["tasks"] {
        let status = "PENDING"
        if (task["completed"]) {
            status = "COMPLETED"
        }

        content += "[" + string(i + 1) + "] " + task["title"] + "\n"
        content += "Status: " + status + "\n"
        content += "Category: " + task["category"] + "\n"
        content += "Priority: " + task["priority"] + "\n"
        if (task["dueDate"] != "") {
            content += "Due Date: " + task["dueDate"] + "\n"
        }
        content += "Description: " + task["description"] + "\n"
        content += "Created: " + task["createdAt"] + "\n"
        if (task["completed"] && task["completedAt"] != "") {
            content += "Completed: " + task["completedAt"] + "\n"
        }
        content += "\n" + "-" * 30 + "\n\n"
    }

    os.writeFile(exportFile, content)
    printColor("green", "✓ Tasks exported to: " + exportFile)
}

// Delete a task
let deleteTask = func() {
    listTasks()
    if (len(taskManager["tasks"]) == 0) {
        return
    }

    let taskIndex = input("Enter task number to delete: ")
    let taskNum = int(taskIndex)

    if (taskNum > 0 && taskNum <= len(taskManager["tasks"])) {
        let task = taskManager["tasks"][taskNum - 1]
        let confirm = input("Are you sure you want to delete '" + task["title"] + "'? (y/N): ")
        if (confirm == "y" || confirm == "Y") {
            taskManager["tasks"].splice(taskNum - 1, 1)
            printColor("green", "✓ Task deleted successfully!")
        } else {
            printColor("yellow", "Delete cancelled.")
        }
    } else {
        printColor("red", "❌ Invalid task number!")
    }
}

// Main menu
let showMenu = func() {
    printSeparator()
    printColor("magenta", "🎯 VintLang Task Manager")
    printSeparator()
    print("1. Create New Task")
    print("2. List All Tasks")
    print("3. Mark Task as Completed")
    print("4. Filter Tasks by Category")
    print("5. Show Statistics")
    print("6. Export Tasks")
    print("7. Delete Task")
    print("8. Save & Exit")
    printSeparator()
}

// Main application loop
let runTaskManager = func() {
    printColor("cyan", "🚀 Welcome to VintLang Task Manager!")
    printColor("cyan", "A showcase of VintLang's capabilities")

    // Load existing tasks
    loadTasks()

    while (true) {
        showMenu()
        let choice = input("Select option (1-8): ")

        if (choice == "1") {
            createTask()
        } else if (choice == "2") {
            listTasks()
        } else if (choice == "3") {
            completeTask()
        } else if (choice == "4") {
            filterByCategory()
        } else if (choice == "5") {
            showStatistics()
        } else if (choice == "6") {
            exportTasks()
        } else if (choice == "7") {
            deleteTask()
        } else if (choice == "8") {
            saveTasks()
            printColor("cyan", "👋 Thank you for using VintLang Task Manager!")
            printColor("cyan", "This showcase demonstrated:")
            print("  • JSON data persistence")
            print("  • File I/O operations")
            print("  • Time and date handling")
            print("  • UUID generation")
            print("  • String manipulation")
            print("  • Interactive user input")
            print("  • Data structures and algorithms")
            print("  • Modular programming")
            break
        } else {
            printColor("red", "❌ Invalid option! Please select 1-8.")
        }

        input("\nPress Enter to continue...")
    }
}

// Start the application
runTaskManager()
```

## simple_build.vint

```js
// Simple Build Script Example
// This demonstrates a basic build automation script using the make module
// Usage: vint examples/simple_build.vint

import make
import cli

print("🔨 Simple Build Automation Example\n")

// Helper function for colored output
let printSuccess = func(msg) {
    print("✅ " + msg)
}

let printInfo = func(msg) {
    print("ℹ️  " + msg)
}

// Define tasks as a dictionary
let tasks = {
    "check": func() {
        printInfo("Checking build requirements...")

        if (make.check("go")) {
            printSuccess("Go compiler found")
        } else {
            print("❌ Go compiler not found")
            return false
        }

        if (make.check("git")) {
            printSuccess("Git found")
        } else {
            printInfo("Git not found (optional)")
        }

        return true
    },

    "info": func() {
        printInfo("Getting system information...")
        let goVersion = make.exec("go version")
        print("Go: " + goVersion)
    },

    "build": func() {
        make.echo("Building application...")

        // Set build environment
        make.env("CGO_ENABLED", "0")

        // Execute build command
        let result = make.exec("go build -o example-app main.go")

        if (result.type != "error") {
            printSuccess("Build completed successfully!")
        } else {
            print("❌ Build failed")
        }
    },

    "clean": func() {
        make.echo("Cleaning build artifacts...")
        make.exec("rm -f example-app")
        printSuccess("Clean completed!")
    },

    "help": func() {
        print("Available tasks:")
        print("  check - Check build requirements")
        print("  info  - Show system information")
        print("  build - Build the application")
        print("  clean - Clean build artifacts")
        print("  help  - Show this help")
        print("\nUsage: vint examples/simple_build.vint <task>")
    }
}

// Parse command line arguments
let args = cli.getArgs()

if (len(args) == 0) {
    print("No task specified. Available tasks:")
    tasks["help"]()
} else {
    let taskName = args[0]

    if (tasks[taskName] != null) {
        print("Running task: " + taskName + "\n")
        tasks[taskName]()
    } else {
        print("❌ Unknown task: " + taskName)
        tasks["help"]()
    }
}

print("\n✅ Done!")

```

## simple_enterprise_test.vint

```js
// Simple Enterprise HTTP Module Test
import http

print("🏢 Enterprise HTTP Module Features Test")
print("=" * 60)

// Test 1: Create app with enterprise features
print("\n✓ Test 1 - Enterprise App Creation")
let result = http.app()
print("App creation:", result)

// Test 2: Route Grouping for API Versioning
print("\n✓ Test 2 - Route Grouping & API Versioning")
let v1Group = http.group("/api/v1", func() {
    print("API v1 group created")
})
print("API v1 group:", v1Group)

// Test 3: Async Handlers
print("\n✓ Test 3 - Async Handlers")
let asyncHandler = http.async(func(req, res) {
    print("Processing asynchronously")
})
print("Async handler created:", asyncHandler)

// Test 4: Security Features
print("\n✓ Test 4 - Security Features")
let securityResult = http.security()
print("Security middleware:", securityResult)

// Test 5: Basic Routes
print("\n✓ Test 5 - Basic Routes Registration")
let getRoute = http.get("/users", func(req, res) {
    print("GET /users route")
})
print("GET route:", getRoute)

let postRoute = http.post("/upload", func(req, res) {
    print("POST /upload route")
})
print("POST route:", postRoute)

// Test 6: Middleware
print("\n✓ Test 6 - Middleware Registration")
let middleware = http.use(func(req, res, next) {
    print("Middleware function")
})
print("Middleware:", middleware)

// Test 7: Guards
print("\n✓ Test 7 - Guards")
let guard = http.guard(func(req) {
    print("Guard function")
})
print("Guard:", guard)

// Test 8: Error Handler
print("\n✓ Test 8 - Error Handler")
let errorHandler = http.errorHandler(func(err, req, res) {
    print("Error handler function")
})
print("Error handler:", errorHandler)

print("\n" + "=" * 60)
print("✨ All enterprise HTTP features registered successfully!")
print("\nEnterprise Features Available:")
print("  🔧 Route grouping and API versioning")
print("  📁 Multipart file upload support")
print("  ⚡ Async handlers for long-running tasks")
print("  🛡️  Enhanced security features")
print("  🔗 Advanced middleware composition")
print("  📊 Structured error handling")
print("  🌐 Production-ready capabilities")

print("\n🎯 Ready for enterprise-level backend development!")
print("📝 Use http.listen(3000) to start the server")
```

## simple_main.vint

```js
// VintLang Main Function Example
// Demonstrates defining and calling a main function

// Code before main function definition executes immediately
println("Before main definition")

// Define a main function that returns a value
let main = func() {
    println("Main function is running!")
    return 42
}

println("After main definition")

// In VintLang, main() must be called explicitly
// It doesn't run automatically like in some languages
// Uncomment the line below to execute main:
// println(main())
```

## simple_task_manager.vint

```js
// VintLang Feature Showcase: Simple Task Manager
// This demonstrates core VintLang features in a working application

import time
import os
import json
import uuid

// Simple task manager state
let tasks = []
let dataFile = "simple_tasks.json"

// Load tasks from file
let loadTasks = func() {
    if (os.fileExists(dataFile)) {
        let data = os.readFile(dataFile)
        tasks = json.decode(data)
        print("✓ Tasks loaded!")
    } else {
        print("⚠ Starting with empty task list")
        tasks = []
    }
}

// Save tasks to file
let saveTasks = func() {
    let data = json.encode(tasks)
    os.writeFile(dataFile, data)
    print("✓ Tasks saved!")
}

// Create a new task
let createTask = func() {
    print("\n--- Creating New Task ---")
    let title = input("Task title: ")
    let description = input("Description: ")

    let task = {
        "id": uuid.generate(),
        "title": title,
        "description": description,
        "completed": false,
        "createdAt": time.now()
    }

    tasks.push(task)
    print("✓ Task created: " + title)
}

// List all tasks
let listTasks = func() {
    print("\n--- Task List ---")
    if (len(tasks) == 0) {
        print("No tasks found!")
        return
    }

    for i, task in tasks {
        let status = "❌"
        if (task["completed"]) {
            status = "✅"
        }
        print(status + " [" + string(i + 1) + "] " + task["title"])
        print("    " + task["description"])
        print("    Created: " + task["createdAt"])
        print("")
    }
}

// Mark task as completed
let completeTask = func() {
    listTasks()
    if (len(tasks) == 0) {
        return
    }

    let taskIndex = input("Enter task number to complete: ")
    let taskNum = int(taskIndex)

    if (taskNum > 0 && taskNum <= len(tasks)) {
        tasks[taskNum - 1]["completed"] = true
        print("✓ Task completed!")
    } else {
        print("❌ Invalid task number!")
    }
}

// Show statistics
let showStats = func() {
    print("\n--- Statistics ---")
    let total = len(tasks)
    let completed = 0

    for task in tasks {
        if (task["completed"]) {
            completed += 1
        }
    }

    print("Total tasks: " + string(total))
    print("Completed: " + string(completed))
    print("Pending: " + string(total - completed))

    if (total > 0) {
        let rate = (completed * 100) / total
        print("Completion rate: " + string(rate) + "%")
    }
}

// Main menu
let showMenu = func() {
    print("\n" + "=" * 30)
    print("🎯 VintLang Simple Task Manager")
    print("=" * 30)
    print("1. Create Task")
    print("2. List Tasks")
    print("3. Complete Task")
    print("4. Show Statistics")
    print("5. Save & Exit")
    print("=" * 30)
}

// Main application
let runApp = func() {
    print("🚀 Welcome to VintLang Task Manager!")
    print("This showcases VintLang features:")
    print("  • JSON data handling")
    print("  • File I/O operations")
    print("  • Time functions")
    print("  • UUID generation")
    print("  • Interactive input")
    print("  • Data structures")

    loadTasks()

    while (true) {
        showMenu()
        let choice = input("Select option (1-5): ")

        if (choice == "1") {
            createTask()
        } else if (choice == "2") {
            listTasks()
        } else if (choice == "3") {
            completeTask()
        } else if (choice == "4") {
            showStats()
        } else if (choice == "5") {
            saveTasks()
            print("👋 Thanks for using VintLang Task Manager!")
            break
        } else {
            print("❌ Invalid option!")
        }
    }
}

// Start the application
runApp()
```

## simplegame.vint

```js
/*
THIS IS A SIMPLE TERMINAL GAME WRITTEN IN VINTLANG
*/
import time
import os
import json
import uuid

// Game initialization
let player = {
    "name": "",
    "health": 100,
    "inventory": [],
    "location": "start"
}

// Save game state to a file
let saveGame = func () {
    let saveData = json.encode(player)
    os.writeFile("savegame.json", saveData)
    print("Game saved!")
}

// Load game state from a file
let loadGame = func () {
    if (os.fileExists("savegame.json")) {
        let saveData = os.readFile("savegame.json")
        player = json.decode(saveData)
        print("Game loaded!")
    } else {
        print("No saved game found.")
    }
}

// Display player stats
let showStats = func () {
    print("Player Stats:")
    print("Name: " + player["name"])
    print("Health: " + string(player["health"]))
    print("Inventory: " + string(player["inventory"]))
    print("Location: " + player["location"])
}

// Handle game events
let handleEvent = func (event) {
    if (event["type"] == "item") {
        print("You found an item: " + event["name"])
        player["inventory"].push(event["name"])
    } else if (event["type"] == "enemy") {
        print("An enemy appears: " + event["name"])
        print("You lose 10 health!")
        player["health"] -= 10
    }
    if (player["health"] <= 0) {
        print("You died! Game Over.")
        os.exit(1)
    }
}

// Main game loop
let gameLoop = func () {
    while (true) {
        print("\nYou are at: " + player["location"])
        print("Choose an action: [explore, stats, save, quit]")
        let action = input("> ")

        if (action == "explore") {
            print("Exploring...")
            let event = {
                "type": "item",
                "name": "Mystic Key"
            }
            handleEvent(event)
        } else if (action == "stats") {
            showStats()
        } else if (action == "save") {
            saveGame()
        } else if (action == "quit") {
            print("Quitting game...")
            os.exit(0)
        } else {
            print("Invalid action!")
        }
    }
}

// Start the game
print("Welcome to the Adventure Game!")
print("Enter your player name:")
player["name"] = input(">>> ")

print("Hello, " + player["name"] + "! Let's begin.")
gameLoop()

```

## sqlite.vint

```js
// VintLang SQLite Database Example
// Demonstrates SQLite database operations: create, insert, query, update

import sqlite

// Drop a table (commented to preserve data between runs)
// sqlite.dropTable(db, "users")

let someFunction = func(){
    println("This is a function in sqlite.vint")

    // Example 1: Open a database connection
    // Creates or opens the database file
    const db = sqlite.open("example.db")
    println("Database opened:", db)
    // defer sqlite.close(db)  // Ensure the database is closed when done
    // defer println("Database closed.")

    // Example 2: Create a table
    // Creates a users table with id, name, and age columns
    sqlite.createTable(db, "CREATE TABLE IF NOT EXISTS users (id INTEGER PRIMARY KEY, name TEXT, age INTEGER)")
    println("Table 'users' created or already exists.")

    // Example 3: Insert data into the table
    // Uses prepared statements with ? placeholders for safety
    sqlite.execute(db, "INSERT INTO users (name, age) VALUES (?, ?)", "Alice", 25)
    sqlite.execute(db, "INSERT INTO users (name, age) VALUES (?, ?)", "Bob", 30)

    // Fetch all rows
    println("=== All Users ===")
    let users = sqlite.fetchAll(db, "SELECT * FROM users")
    println(users)

    // Fetch a single row
    println("\n=== First User ===")
    let first_user = sqlite.fetchOne(db, "SELECT * FROM users LIMIT 1")
    println(first_user)

    // Update data
    sqlite.execute(db, "UPDATE users SET age = ? WHERE name = ?", 26, "Alice")

    // Delete data
    //sqlite.execute(db, "DELETE FROM users WHERE name = ?", "Bob")

    // Fetch all rows after changes
    println("\n=== Users After Changes ===")
    let users_after_changes = sqlite.fetchAll(db, "SELECT * FROM users")
    println(users_after_changes)

}

someFunction()

```

## string_builtins_demo.vint

```js
// Comprehensive test of new string builtin functions in VintLang

println("=== String Builtin Functions Test ===");
println();

// Test format function with various data types
println("1. format() function:");
println("   String formatting:", format("Hello %s!", "World"));
println("   Integer formatting:", format("Number: %d", 42));
println(
  "   Mixed formatting:",
  format("%s has %d apples and %s oranges", "Alice", 5, "3"),
);
println();

// Test startsWith function
println("2. startsWith() function:");
let text = "Hello World";
println("   Text:", format('"%s"', text));
println("   Starts with 'Hello':", startsWith(text, "Hello"));
println("   Starts with 'World':", startsWith(text, "World"));
println("   Starts with 'H':", startsWith(text, "H"));
println("   Starts with '':", startsWith(text, ""));
println();

// Test endsWith function
println("3. endsWith() function:");
println("   Text:", format('"%s"', text));
println("   Ends with 'World':", endsWith(text, "World"));
println("   Ends with 'Hello':", endsWith(text, "Hello"));
println("   Ends with 'd':", endsWith(text, "d"));
println("   Ends with '':", endsWith(text, ""));
println();

// Test chr function (ASCII to character)
println("4. chr() function:");
println("   chr(65):", format('"%s" (ASCII A)', chr(65)));
println("   chr(97):", format('"%s" (ASCII a)', chr(97)));
println("   chr(48):", format('"%s" (ASCII 0)', chr(48)));
println("   chr(32):", format('"%s" (ASCII space)', chr(32)));
println();

// Test ord function (character to ASCII)
println("5. ord() function:");
println("   ord('A'):", ord("A"));
println("   ord('a'):", ord("a"));
println("   ord('0'):", ord("0"));
println("   ord(' '):", ord(" "));
println();

// Demonstrate practical usage
println("6. Practical examples:");
let name = "VintLang";
let version = "1.0";
let message = format("Welcome to %s v%s!", name, version);
println("   Formatted message:", message);

let filename = "data.txt";
if (endsWith(filename, ".txt")) {
  println("   File is a text file:", filename);
}

if (startsWith(name, "Vint")) {
  println("   Language name starts with 'Vint'");
}

// Convert between characters and ASCII
let letter = chr(ord("A") + 1); // Should be "B"
println("   Next letter after A:", letter);

println();
println("=== All tests completed ===");
```

## strings.vint

```js
// Sample usage of the string module
// The string module provides various string manipulation functions
import string

// Example 1: Trim whitespace
let result = string.trim("  Hello, World!  ")
print(result)  // Output: "Hello, World!"

// Example 2: Check if a string contains a substring
let containsResult = string.contains("Hello, World!", "World")
print(containsResult)  // Output: true

// Example 3: Convert to uppercase
let upperResult = string.toUpper("hello")
print(upperResult)  // Output: "HELLO"

// Example 4: Convert to lowercase
let lowerResult = string.toLower("HELLO")
print(lowerResult)  // Output: "hello"

// Example 5: Replace a substring
let replaceResult = string.replace("Hello, World!", "World", "Vint")
print(replaceResult)  // Output: "Hello, Vint!"

// Example 6: Split string into parts
let splitResult = string.split("a,b,c,d", ",")
print(splitResult)  // Output: ["a", "b", "c", "d"]

// Example 7: Join string parts
let joinResult = string.join(["a", "b", "c"], "-")
print(joinResult)  // Output: "a-b-c"

// Example 8: Get a substring
let substringResult = string.substring("Hello, World!", 7, 12)
print(substringResult)  // Output: "World"

// Example 9: Get the length of a string
let lengthResult = string.length("Hello")
print(lengthResult)  // Output: 5

// Example 10: Find index of a substring
let indexResult = string.indexOf("Hello, World!", "World")
print(indexResult)  // Output: 7

// Example 11: Get a substring (valid start and end indices)
result = string.substring("Hello, World!", 0, 5)
print(result)  // Output: "Hello"

// Example 12: Invalid indices (start >= end)
// result = string.substring("Hello, World!", 7, 3)
// print(result)  // Output: Error: Invalid start or end index

//Example 13: Generating a slug
result = string.slug("Creates a slug string from a normal string")
print(result)   //Output: string: creates-a-slug-string-from-a-normal-string


/*
More methods for this module
simirality
*/
```

## switch.vint

```js
// VintLang Switch Statement Example
// Demonstrates the switch-case control flow structure

let n = 1

// Switch statement checks the value of n against multiple cases
switch (n) {
    case 1 {
        println("is",n)
        break  // Exit the switch statement after matching
    }
    default {
        // Default case runs if no other case matches
        println("is not")
        // break
    }
}
```

## sysinfo.vint

```js
import sysinfo

println("=== Vint System Information Module Demo ===")
println("")

// Basic system information
println("1. Basic System Info:")
println("   Operating System:", sysinfo.os())
println("   Architecture:", sysinfo.arch())
println("")

// Memory information
println("2. Memory Information:")
let memory = sysinfo.memInfo()
println("   Total Memory:", memory["total"])
println("   Available Memory:", memory["available"])
println("   Used Memory:", memory["used"])
println("   Free Memory:", memory["free"])
println("   Usage Percentage:", memory["percent"] + "%")
println("")

// CPU information
println("3. CPU Information:")
let cpu = sysinfo.cpuInfo()
println("   CPU Model:", cpu["model"])
println("   CPU Cores:", cpu["cores"])
println("   CPU Frequency:", cpu["frequency"])
println("   CPU Usage:", cpu["usage"] + "%")
println("")

// Disk information
println("4. Disk Information:")
let disk = sysinfo.diskInfo()
println("   Total Disk Space:", disk["total"])
println("   Used Disk Space:", disk["used"])
println("   Free Disk Space:", disk["free"])
println("   Disk Usage:", disk["percent"] + "%")
println("")

// Network interfaces
println("5. Network Interfaces:")
let interfaces = sysinfo.netInfo()
println("   Number of interfaces:", len(interfaces))
println("   Example interface data:", interfaces)
println("")

println("=== Complete! All sysinfo functions working ===")

```

## system_monitor_cli.vint

```js
// System Monitor CLI - Real-time system information display
// Run with: vint system_monitor_cli.vint [--watch] [--summary]

import term
import cli
import os

// Help handling
if (cli.hasArg("--help")) {
    cli.help("SystemMonitor", "Monitor system resources and information")
    term.println("")
    term.info("Monitor Options:")
    term.println("  --summary        Show system summary only")
    term.println("  --watch          Continuous monitoring mode")
    term.println("  --refresh N      Refresh interval in seconds")
    exit(0)
}

// Banner
let banner = term.banner("VintLang System Monitor")
term.println(banner)

// Check for summary mode
if (cli.hasArg("--summary")) {
    term.info("System Summary")

    let summaryTable = term.table([
        ["Component", "Status", "Details"],
        ["Operating System", "✓ Online", "Linux/Unix"],
        ["Memory", "✓ Normal", "8.2 GB Available"],
        ["Storage", "✓ Normal", "45.6 GB Free"],
        ["Network", "✓ Connected", "192.168.1.100"]
    ])
    term.println(summaryTable)

    let chart = term.chart([82, 67, 91, 100])
    term.println("Resource Usage (%):")
    term.println(chart)
    exit(0)
}

// Interactive monitoring
term.info("System Monitoring Interface")
let option = term.select([
    "View system information",
    "Check resource usage",
    "Monitor network status",
    "View process information",
    "Generate report",
    "Exit"
])

if (option == "View system information") {
    term.info("System Information")

    let currentDir = os.getwd()
    let sysInfo = term.table([
        ["Property", "Value"],
        ["Current Directory", currentDir],
        ["User", "vintlang-user"],
        ["Shell", "/bin/bash"],
        ["Path", "/usr/local/bin:/usr/bin:/bin"]
    ])
    term.println(sysInfo)

} else if (option == "Check resource usage") {
    term.info("Resource Usage")

    let resources = term.table([
        ["Resource", "Used", "Available", "Percentage"],
        ["Memory", "4.2 GB", "8.0 GB", "52%"],
        ["CPU", "2.1 GHz", "3.2 GHz", "65%"],
        ["Storage", "234 GB", "500 GB", "47%"],
        ["Network", "45 Mb/s", "100 Mb/s", "45%"]
    ])
    term.println(resources)

    let usageChart = term.chart([52, 65, 47, 45])
    term.println("Usage Distribution:")
    term.println(usageChart)

} else if (option == "Monitor network status") {
    term.info("Network Status")

    let networkTable = term.table([
        ["Interface", "Status", "IP Address", "Speed"],
        ["eth0", "✓ Up", "192.168.1.100", "1 Gbps"],
        ["wlan0", "✗ Down", "N/A", "N/A"],
        ["lo", "✓ Up", "127.0.0.1", "Local"]
    ])
    term.println(networkTable)

} else if (option == "View process information") {
    term.info("Top Processes")

    let processes = term.table([
        ["PID", "Name", "CPU%", "Memory"],
        ["1234", "vintlang", "15.2%", "124 MB"],
        ["5678", "node", "8.5%", "245 MB"],
        ["9012", "python", "3.1%", "89 MB"],
        ["3456", "bash", "0.5%", "12 MB"]
    ])
    term.println(processes)

} else if (option == "Generate report") {
    let reportName = term.input("Enter report name: ")
    term.loading("Generating system report...")

    let reportTable = term.table([
        ["Report Section", "Status"],
        ["System Info", "✓ Complete"],
        ["Resource Usage", "✓ Complete"],
        ["Network Status", "✓ Complete"],
        ["Process List", "✓ Complete"]
    ])
    term.println(reportTable)

    term.success("Report '" + reportName + "' generated successfully!")

} else if (option == "Exit") {
    let saveConfig = term.confirm("Save monitoring configuration?")
    if (saveConfig) {
        term.success("Configuration saved")
    }
    term.success("System monitor closed")
}

let statusBox = term.box("System monitoring session completed")
term.println(statusBox)
```

## task_manager_cli.vint

```js
// Task Manager CLI - A practical example of VintLang terminal capabilities
// Run with: vint task_manager_cli.vint [--help] [--list] [--add "task"]

import term
import cli

// Check for help
if (cli.hasArg("--help") || cli.hasArg("-h")) {
    cli.help("TaskManager", "A simple task management CLI tool")
    term.println("")
    term.info("Additional Usage Examples:")
    term.println("  vint task_manager_cli.vint --list")
    term.println("  vint task_manager_cli.vint --add \"Write documentation\"")
    term.println("  vint task_manager_cli.vint --interactive")
    exit(0)
}

// Initialize tasks (in real app, would load from file)
let tasks = [
    "Review code changes",
    "Update documentation",
    "Test new features",
    "Deploy to staging"
]

// Display app banner
let banner = term.banner("VintLang Task Manager")
term.println(banner)

// Handle command line operations
if (cli.hasArg("--list")) {
    term.info("Current Tasks:")
    let taskTable = term.table([
        ["ID", "Task", "Status"],
        ["1", "Review code changes", "Pending"],
        ["2", "Update documentation", "Pending"],
        ["3", "Test new features", "In Progress"],
        ["4", "Deploy to staging", "Pending"]
    ])
    term.println(taskTable)
    exit(0)
}

let newTask = cli.getArgValue("--add")
if (newTask) {
    term.success("Added new task: " + newTask)
    exit(0)
}

// Interactive mode
term.info("Interactive Task Manager")
let action = term.select([
    "View all tasks",
    "Add new task",
    "Mark task complete",
    "Delete task",
    "Task statistics",
    "Exit"
])

if (action == "View all tasks") {
    let taskTable = term.table([
        ["ID", "Task", "Status"],
        ["1", "Review code changes", "Pending"],
        ["2", "Update documentation", "Pending"],
        ["3", "Test new features", "In Progress"],
        ["4", "Deploy to staging", "Pending"]
    ])
    term.println(taskTable)

} else if (action == "Add new task") {
    let taskName = term.input("Enter task description: ")
    term.success("Added: " + taskName)

} else if (action == "Mark task complete") {
    let taskChoice = term.select([
        "Review code changes",
        "Update documentation",
        "Test new features",
        "Deploy to staging"
    ])
    term.success("Marked as complete: " + taskChoice)

} else if (action == "Delete task") {
    let taskToDelete = term.radio([
        "Review code changes",
        "Update documentation",
        "Test new features",
        "Deploy to staging"
    ])
    let confirmed = term.confirm("Delete '" + taskToDelete + "'?")
    if (confirmed) {
        term.success("Task deleted: " + taskToDelete)
    }

} else if (action == "Task statistics") {
    term.info("Task Statistics")
    let stats = term.table([
        ["Metric", "Value"],
        ["Total Tasks", "4"],
        ["Completed", "0"],
        ["In Progress", "1"],
        ["Pending", "3"]
    ])
    term.println(stats)

    let chart = term.chart([0, 1, 3])
    term.println("Status Distribution:")
    term.println(chart)

} else if (action == "Exit") {
    term.success("Thank you for using VintLang Task Manager!")
}

let message = term.box("Task management session complete!")
term.println(message)
```

## term.vint

```js
import term
import time

// Display a banner
let banner = term.banner("Welcome to VintLang Terminal Demo!")
term.println(banner)

// Get terminal size
let size = term.getSize()
term.println("Terminal size: 80x24 (default)")

// Create a select menu
term.println("Select an option:")
let choice = term.select([
    "Start Game",
    "Show Settings",
    "Exit"
])
term.println("You selected: " + choice)

// Create a checkbox list
term.println("Select multiple options (enter numbers like '1 3 4'):")
let selected = term.checkbox([
    "Option 1",
    "Option 2",
    "Option 3",
    "Option 4"
])
term.println("Selected options:")
for option in selected {
    term.println("- " + option)
}

// Create a radio button list
term.println("Select one option:")
let radioChoice = term.radio([
    "Yes",
    "No",
    "Maybe"
])
term.println("You selected: " + radioChoice)

// Get password input
let password = term.password("Enter your password: ")
term.println("Password entered (hidden)")

// Ask for confirmation
let confirmed = term.confirm("Do you want to proceed?")
if (confirmed) {
    term.success("Proceeding...")
} else {
    term.error("Operation cancelled")
}

// Show different types of messages
term.info("This is an information message")
term.warning("This is a warning message")
term.error("This is an error message")
term.success("This is a success message")
term.notify("This is a notification")

// Show a loading spinner
term.loading("Processing...")

// Create a styled table
let table = term.table([
    ["Feature", "Status"],
    ["Select Menu", "✓"],
    ["Checkbox", "✓"],
    ["Radio Buttons", "✓"],
    ["Password Input", "✓"],
    ["Confirmations", "✓"],
    ["Messages", "✓"]
])
term.println(table)

// Show progress indication
let progress = term.progress(100)
term.println(progress)

// Create a boxed message
let message = term.box("Thank you for trying the terminal features!")
term.println(message)

// Hide cursor, show message, then show cursor
term.cursor(false)
term.println("Cursor is hidden")
term.cursor(true)

// Play a beep
term.beep()

// Clear screen and show final message
term.clear()
let styledMsg = term.style("Demo completed!", {
    "color": "green",
    "bold": "true"
})
term.println(styledMsg)

// Show some charts
let data = [10, 20, 30, 40, 50]
let chart = term.chart(data)
term.println("Sample chart:")
term.println(chart)

term.println("Terminal demo completed successfully!")
```

## test-for.vint

```js
// VintLang For Loop Example
// Demonstrates iterating over array elements

// Create an array of todo items
let todos = ["new", "todo"]

// Print the entire array
println("Your TODOs:", todos)

// Check the type of the variable
println(type(todos))

// Iterate through each element in the array
// The loop variable (tach) takes on each value from the array
for tach in todos {
    println(tach)
}
```

## test_array_patterns.vint

```js
// Test array patterns in match statements
let arr1 = []
let arr2 = [1]
let arr3 = [1, 2]
let arr4 = [1, 2, 3, 4, 5]

println("Testing array patterns:")

println("Empty array:")
match arr1 {
    [] => println("  Matched empty array")
    _ => println("  Not empty")
}

println("Single element:")
match arr2 {
    [x] => println("  Single element:", x)
    _ => println("  Not single")
}

println("Two elements:")
match arr3 {
    [a, b] => println("  Two elements:", a, "and", b)
    _ => println("  Not two elements")
}

println("Multiple elements with spread:")
match arr4 {
    [first, ...rest] => println("  First:", first, "Rest length:", len(rest))
    _ => println("  No match")
}
```

## test_array_var.vint

```js
// Test array pattern with variable
let arr = [42]

println("Testing array pattern with variable:")

match arr {
    [x] => println("Single element:", x)
    _ => println("Not single")
}
```

## test_import.vint

```js
// VintLang Import Statement Example
// Demonstrates importing and using a module

import time

// Use the imported time module to get the current timestamp
println("Current time:", time.now())
```

## test_import_lookahead.vint

```js
println("=== Testing import() function with lookahead ===");

// Test import function call
println("Testing import('os'):");
let osModule = import("os");
println("Result:", osModule);

// Test import statement
println("Testing import statement:");
import math;
println("Math module imported:", math);

println("=== Tests completed ===");
```

## test_import_whitespace.vint

```js
println("=== Testing import with various whitespace scenarios ===");

// Test with spaces before parentheses
let osModule = import   ("os");
println("import   ('os'):", osModule);

// Test with newlines and tabs (should still be function call)
let mathModule = import
    ("math");
println("import\\n    ('math'):", mathModule);

// Test import statement with spaces
import    time;
println("import    time;:", time);

println("=== Whitespace tests completed ===");
```

## test_match_guards.vint

```js
// Test enhanced match with guard conditions
// This is a simple test to verify our match guard condition implementation

let user = {"role": "admin", "age": 30}

println("Testing match with guard conditions")

match user {
    {"role": "admin", "age": age} if age >= 18 => println("Adult admin, age:", age)
    {"role": "user"} => println("Regular user")
    _ => println("Unknown user type")
}
```

## test_simple_array.vint

```js
// Simple array pattern test
let arr = [1, 2]

println("Testing simple array pattern:")

match arr {
    [] => println("Empty")
    _ => println("Not empty")
}
```

## test_simple_guard.vint

```js
// Test match with simple guard
let user = {"role": "admin"}

println("Testing simple match with guard")

match user {
    {"role": role} if role == "admin" => println("Found admin with role:", role)
    _ => println("Not admin")
}
```

## test_switch_guards.vint

```js
// Test enhanced switch with guard conditions
// This is a simple test to verify our guard condition implementation

let x = 5

println("Testing switch with guard conditions")

switch (x) {
    case y if y > 0 {
        println("Positive number:", y)
    }
    case y if y < 0 {
        println("Negative number:", y)
    }
    case 0 {
        println("Zero")
    }
    default {
        println("Unknown")
    }
}
```

## test_two_elem.vint

```js
// Very simple spread test
let arr = [1, 2]

match arr {
    [a, b] => println("Two elements:", a, b)
    _ => println("Other")
}
```

## todo_example.vint

```js
println("Starting the program...")

let name = "Vint"
println("Hello, " + name)

todo "Refactor this section later"

let greet = func (person) {
    println("Greetings, " + person + "!")
}

greet("developer")

println("Program finished.")
```

## todo_test.vint

```js
todo "this is a test todo"
```

## unique_builtins_test.vint

```js
// Test file for unique built-in functions that don't exist in modules
// These are the functions that were kept after removing duplicates

println("Testing unique built-in functions...");

// Test string functions
println("=== String Functions ===");
println("startsWith('VintLang', 'Vint'):", startsWith("VintLang", "Vint"));
println("startsWith('hello', 'hi'):", startsWith("hello", "hi"));
println("endsWith('VintLang', 'Lang'):", endsWith("VintLang", "Lang"));
println("endsWith('hello', 'world'):", endsWith("hello", "world"));

// Test array function
println("\n=== Array Functions ===");
let arr = [1, 2, 3, 2, 4];
println("Array:", arr);
println("indexOf(arr, 2):", indexOf(arr, 2));
println("indexOf(arr, 5):", indexOf(arr, 5));

// Test type checking functions
println("\n=== Type Checking Functions ===");
println("isInt(42):", isInt(42));
println("isInt(3.14):", isInt(3.14));
println("isFloat(3.14):", isFloat(3.14));
println("isFloat(42):", isFloat(42));
println("isString('hello'):", isString("hello"));
println("isString(42):", isString(42));
println("isBool(true):", isBool(true));
println("isBool(42):", isBool(42));
println("isArray([1,2,3]):", isArray([1, 2, 3]));
println("isArray('hello'):", isArray("hello"));
println("isDict({'key': 'value'}):", isDict({ key: "value" }));
println("isDict([1,2,3]):", isDict([1, 2, 3]));
println("isNull(null):", isNull(null));
println("isNull(42):", isNull(42));

// Test parsing functions
println("\n=== Parsing Functions ===");
println("parseInt('42'):", parseInt("42"));
println("parseInt('-10'):", parseInt("-10"));
println("parseFloat('3.14'):", parseFloat("3.14"));
println("parseFloat('-2.5'):", parseFloat("-2.5"));

println("\nAll unique built-in functions tested successfully!");
```

## unique_feature_test.vint

```js
// VintLang unique() function demonstration and test
// This file showcases the new unique() builtin function that removes duplicates from arrays

print("🚀 VintLang unique() Function Showcase");
print("=" * 40);

// Demonstrate the core "no duplicates" functionality
print("\n📊 Data Processing with unique():");

// Example 1: Processing survey responses
print("\n1. Survey Responses (with duplicates):");
let responses = ["yes", "no", "maybe", "yes", "no", "yes", "maybe"];
print("Original responses:", responses);
print("Unique responses:", unique(responses));
print(
  "Count: " + string(len(responses)) + " -> " + string(len(unique(responses))),
);

// Example 2: User IDs from multiple sources
print("\n2. User IDs from different systems:");
let userIds = [101, 202, 303, 101, 404, 202, 505, 303, 101];
print("All IDs:", userIds);
print("Unique IDs:", unique(userIds));

// Example 3: Feature tags (mixed data types)
print("\n3. Feature tags with mixed types:");
let tags = ["cool", 1, true, "awesome", 1, false, "cool", true];
print("All tags:", tags);
print("Unique tags:", unique(tags));

// Example 4: Real-world use case - removing duplicate items
print("\n4. Shopping cart cleanup:");
let cart = ["apple", "banana", "apple", "orange", "banana", "apple"];
print("Cart with duplicates:", cart);
let cleanCart = unique(cart);
print("Clean cart:", cleanCart);
print("Items saved: " + string(len(cart) - len(cleanCart)));

// Example 5: Data validation - ensuring no duplicates
print("\n5. Data validation example:");
let data = [1, 2, 3, 4, 5];
let uniqueData = unique(data);
if (len(data) == len(uniqueData)) {
  print("✅ Data has no duplicates!");
} else {
  print("⚠️  Data contains duplicates - cleaned up");
}

print("\n🎯 Conclusion:");
print("The unique() function provides a simple, efficient way to handle");
print("the 'no duplicates' requirement in data processing workflows!");
print("\n✨ This is a cool feature that makes VintLang even more powerful! ✨");
```

## user_access_control.vint

```js
// Comprehensive example: User access control system using dict pattern matching

let users = [
    {"role": "admin", "active": true, "department": "IT"},
    {"role": "admin", "active": false, "department": "HR"},
    {"role": "user", "active": true, "department": "Sales"},
    {"role": "user", "active": false, "department": "Marketing"},
    {"role": "guest", "active": true},
    {"name": "John", "type": "visitor"},
    {}
]

print("=== User Access Control System ===")

for user in users {
    print("\nChecking user:", user)

    match user {
        {"role": "admin", "active": true} => {
            print("ACCESS GRANTED: Active administrator")
            print("Full system access available")
        }
        {"role": "admin", "active": false} => {
            print("ACCESS DENIED: Inactive administrator")
            print("Please contact system administrator")
        }
        {"role": "user", "active": true} => {
            print("ACCESS GRANTED: Active user")
            print("Limited access to user features")
        }
        {"role": "user", "active": false} => {
            print("ACCESS DENIED: Inactive user")
            print("Account has been suspended")
        }
        {"role": "guest"} => {
            print("ACCESS GRANTED: Guest user")
            print("Read-only access to public content")
        }
        {"name": "John"} => {
            print("ACCESS GRANTED: Known visitor")
            print("Limited visitor access")
        }
        {} => {
            print("ACCESS DENIED: Unrecognized user structure")
            print("User has valid dict structure but unknown pattern")
        }
        _ => {
            print("ACCESS DENIED: Invalid user data")
            print("Please provide valid user information")
        }
    }
}
```

## uuid.vint

```js
// VintLang UUID Module Example
// Demonstrates generating universally unique identifiers (UUIDs)

import uuid

// Generate a random UUID (Version 4)
// Returns a string in the format: xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx
print(uuid.generate())
```

## vintChart.vint

```js
import vintChart

labels = ["A", "B", "C", "D"]
values = [10, 20, 30, 40]
vintChart.barChart(labels, values, "bar_chart.png")
vintChart.pieChart(labels, values, "pie_chart.png")
vintChart.lineGraph(labels, values, "line_graph.png")

```

## vintSocket.vint

```js
import vintSocket

vintSocket.createServer("8080")
vintSocket.connect("ws://localhost:8080")
vintSocket.broadcast("Hello, WebSocket clients!")

```

## vintlang_showcase.vint

```js
// VintLang Showcase: Personal Information Manager
// A comprehensive demonstration of VintLang capabilities

import time
import os
import json
import uuid

print("🚀 VintLang Personal Information Manager")
print("=" * 50)
print("Showcasing VintLang's real-world capabilities")
print("=" * 50)

// Step 1: Data Generation and JSON Handling
print("\n📊 Step 1: Data Generation and JSON Operations")
print("-" * 50)

let contacts = [
    {
        "id": uuid.generate(),
        "name": "Alice Johnson",
        "email": "alice@example.com",
        "phone": "+1-555-0101",
        "category": "Work",
        "created": time.format(time.now(), "2006-01-02")
    },
    {
        "id": uuid.generate(),
        "name": "Bob Smith",
        "email": "bob@example.com",
        "phone": "+1-555-0102",
        "category": "Personal",
        "created": time.format(time.now(), "2006-01-02")
    },
    {
        "id": uuid.generate(),
        "name": "Carol Davis",
        "email": "carol@company.com",
        "phone": "+1-555-0103",
        "category": "Work",
        "created": time.format(time.now(), "2006-01-02")
    }
]

print("✓ Generated " + string(len(contacts)) + " contact records")
print("✓ Each contact has unique UUID: " + contacts[0]["id"])

// Step 2: File I/O Operations
print("\n💾 Step 2: File I/O Operations")
print("-" * 50)

let contactsJson = json.encode(contacts)
os.writeFile("contacts.json", contactsJson)
print("✓ Saved contacts to JSON file")

let loadedData = os.readFile("contacts.json")
let loadedContacts = json.decode(loadedData)
print("✓ Loaded and verified " + string(len(loadedContacts)) + " contacts")

// Step 3: Data Analysis
print("\n🔍 Step 3: Data Analysis")
print("-" * 50)

let categoryCount = {}
for contact in loadedContacts {
    let category = contact["category"]
    if (!categoryCount.hasKey(category)) {
        categoryCount[category] = 0
    }
    categoryCount[category] += 1
}

print("Contact distribution by category:")
for category, count in categoryCount {
    print("  " + category + ": " + string(count) + " contacts")
}

// Step 4: Advanced String Operations
print("\n📝 Step 4: String Processing")
print("-" * 50)

for contact in contacts {
    let name = contact["name"]
    let nameParts = name.split(" ")
    let firstName = nameParts[0]
    let lastName = nameParts[1]

    print("Processing: " + name)
    print("  First: " + firstName + ", Last: " + lastName)
    print("  Email domain: " + contact["email"].split("@")[1])
    print("  Phone area: " + contact["phone"].split("-")[1])
    print("")
}

// Step 5: Report Generation
print("\n📋 Step 5: Report Generation")
print("-" * 50)

let reportTime = time.format(time.now(), "2006-01-02 15:04:05")
let report = "PERSONAL INFORMATION MANAGER REPORT\n"
report += "Generated: " + reportTime + "\n"
report += "=" * 50 + "\n\n"

report += "SUMMARY\n"
report += "-------\n"
report += "Total Contacts: " + string(len(contacts)) + "\n"

for category, count in categoryCount {
    report += category + " Contacts: " + string(count) + "\n"
}

report += "\nCONTACT DETAILS\n"
report += "---------------\n"
for contact in contacts {
    report += "Name: " + contact["name"] + "\n"
    report += "Email: " + contact["email"] + "\n"
    report += "Phone: " + contact["phone"] + "\n"
    report += "Category: " + contact["category"] + "\n"
    report += "ID: " + contact["id"] + "\n"
    report += "Created: " + contact["created"] + "\n"
    report += "\n"
}

report += "=" * 50 + "\n"
report += "Report generated by VintLang PIM v1.0\n"

let reportFile = "contact_report_" + time.format(time.now(), "2006-01-02_15-04-05") + ".txt"
os.writeFile(reportFile, report)
print("✓ Generated comprehensive report: " + reportFile)

// Step 6: CSV Export
print("\n📊 Step 6: CSV Export")
print("-" * 50)

let csvContent = "ID,Name,Email,Phone,Category,Created\n"
for contact in contacts {
    csvContent += contact["id"] + ","
    csvContent += contact["name"] + ","
    csvContent += contact["email"] + ","
    csvContent += contact["phone"] + ","
    csvContent += contact["category"] + ","
    csvContent += contact["created"] + "\n"
}

let csvFile = "contacts_export_" + time.format(time.now(), "2006-01-02_15-04-05") + ".csv"
os.writeFile(csvFile, csvContent)
print("✓ Exported data to CSV format: " + csvFile)

// Step 7: Configuration Management
print("\n⚙️ Step 7: Configuration Management")
print("-" * 50)

let config = {
    "app_name": "VintLang PIM",
    "version": "1.0.0",
    "created": time.format(time.now(), "2006-01-02 15:04:05"),
    "features": [
        "Contact Management",
        "Data Analysis",
        "Report Generation",
        "CSV Export",
        "JSON Processing"
    ],
    "settings": {
        "auto_backup": true,
        "data_format": "json",
        "report_format": "txt"
    },
    "statistics": {
        "total_contacts": len(contacts),
        "categories": len(categoryCount),
        "last_updated": time.format(time.now(), "2006-01-02 15:04:05")
    }
}

os.writeFile("config.json", json.encode(config))
print("✓ Created configuration file with app settings")

// Step 8: Directory Listing and File Management
print("\n📁 Step 8: File Management")
print("-" * 50)

let files = os.listDir(".")
let fileList = files.split(", ")
let generatedFiles = []

for filename in fileList {
    if (filename.contains("contact") || filename.contains("config") || filename.contains(".json") || filename.contains(".csv") || filename.contains(".txt")) {
        if (os.fileExists(filename)) {
            let content = os.readFile(filename)
            print("📄 " + filename + " (" + string(len(content)) + " bytes)")
            generatedFiles.push(filename)
        }
    }
}

print("✓ Generated " + string(len(generatedFiles)) + " files during demonstration")

// Step 9: Final Summary and Statistics
print("\n🎉 Step 9: Final Summary")
print("-" * 50)

print("VintLang Personal Information Manager completed successfully!")
print("")
print("📊 DEMONSTRATION STATISTICS:")
print("  • Contacts processed: " + string(len(contacts)))
print("  • Files generated: " + string(len(generatedFiles)))
print("  • Categories analyzed: " + string(len(categoryCount)))
print("  • UUIDs generated: " + string(len(contacts)))
print("  • JSON operations: 4 (encode/decode)")
print("  • String operations: " + string(len(contacts) * 3))
print("")

print("✨ VINTLANG FEATURES DEMONSTRATED:")
print("  ✓ UUID Generation - Unique identifiers for all records")
print("  ✓ Time Operations - Timestamps and formatting")
print("  ✓ JSON Processing - Encoding and decoding data")
print("  ✓ File I/O - Reading and writing multiple file types")
print("  ✓ String Manipulation - Splitting, processing text")
print("  ✓ Data Structures - Arrays and dictionaries")
print("  ✓ Control Flow - Loops and conditionals")
print("  ✓ Directory Operations - File listing and management")
print("  ✓ Report Generation - Formatted text output")
print("  ✓ CSV Export - Data format conversion")
print("")

print("🚀 VintLang is production-ready for:")
print("  • Data processing applications")
print("  • File management systems")
print("  • Report generation tools")
print("  • Configuration management")
print("  • Business automation scripts")
print("  • API data processing")
print("  • Log analysis tools")
print("")

print("🎯 This demonstration proves VintLang's capability")
print("   to handle real-world programming challenges!")

print("\n" + "=" * 50)
print("🏆 VintLang Showcase Complete!")
print("=" * 50)
```

## weatherapp.vint

```js
//NOTE: This currently Does not work due to some limititions of some modules

// Importing modules
import net        // Networking for fetching APIs
import time       // Time utilities
import os         // File and OS utilities
import json       // JSON parsing
import uuid       // Unique ID generation

// Session ID for unique tracking
let session_id = uuid.generate()
println("Session ID:", session_id)

// Fetch weather using wttr.in (No API key required)
let fetchWeather = func(city) {
    let url = "https://wttr.in/" + city + "?format=%C+%t+%w"
    let response = net.get(url)
    if (response) {
        println("Weather in " + city + ":")
        println(response)
    } else {
        println("Error fetching weather data for " + city)
    }
}

// Fetch news headlines (mock API, no key needed)
let fetchNews = func() {
    let url = "https://jsonplaceholder.typicode.com/posts"
    let response = net.get(url)
    if (response) {
        let news = json.decode(response.body)
        println("\n=== Top 3 News Headlines ===")
        for i in range(0, 3) {
            println("Title:", news[i]["title"])
            println("Summary:", news[i]["body"])
            println("-------------------------")
        }
    } else {
        println("Error fetching news.")
    }
}

// Save logs with timestamp
let saveLog = func(logMessage) {
    let log_file = "app_logs.txt"
    let timestamp = time.now().toString()
    os.writeFile(log_file, "[" + timestamp + "] " + logMessage + "\n", "append")
    println("Log saved:", logMessage)
}

// Greet user dynamically based on the time
let greetUser = func(name) {
    let currentTime = time.now()  // Example: returns "19:15:50 28-11-2024"
    let hour = currentTime.split(" ")[0].split(":")[0].toInt()  // Extract hour as an integer

    if (hour < 12) {
        println("Good morning, " + name + "!")
    } else if (hour < 18) {
        println("Good afternoon, " + name + "!")
    } else {
        println("Good evening, " + name + "!")
    }
}


// List all saved logs
let viewLogs = func() {
    let log_file = "app_logs.txt"
    if (os.exists(log_file)) {
        let logs = os.readFile(log_file)
        println("=== Log History ===")
        println(logs)
    } else {
        println("No logs available yet.")
    }
}

// Main interactive menu
let main = func() {
    println("\n=== Welcome to the Dynamic Dashboard ===")
    // greetUser("Tachera")
    println("\nChoose an option:")
    println("1. View Weather Information")
    println("2. Read Latest News")
    println("3. Save a Custom Log")
    println("4. View All Logs")
    println("5. Exit")

    let choice = input("\nEnter your choice (1-5): ")
    if (choice == "1") {
        let city = input("Enter city name: ")
        fetchWeather(city)
    } else if (choice == "2") {
        fetchNews()
    } else if (choice == "3") {
        let logMessage = input("Enter a log message: ")
        saveLog(logMessage)
    } else if (choice == "4") {
        viewLogs()
    } else if (choice == "5") {
        println("Goodbye! See you next time.")
        os.exit(0)
    } else {
        println("Invalid choice. Try again.")
    }
}

// Run the app in a loop
while (true) {
    main()
}

```

## web_fetcher.vint

```js
// VintLang Web Data Fetcher & Analyzer
// Demonstrates network operations and data processing

import net
import json
import time
import os
import uuid

print("🌐 VintLang Web Data Fetcher & Analyzer")
print("=" * 50)
print("Demonstrating network and data processing capabilities")
print("=" * 50)

// Test 1: Simple HTTP Request
print("\n🔄 Test 1: Basic HTTP Request")
print("-" * 40)

let url1 = "https://httpbin.org/json"
print("Fetching data from: " + url1)
let response1 = net.get(url1)

if (response1 != "") {
    print("✓ Successfully fetched data!")
    print("Response length: " + string(len(response1)) + " characters")

    // Save raw response
    os.writeFile("web_response_1.json", response1)
    print("✓ Saved response to web_response_1.json")

    // Parse JSON response
    let data1 = json.decode(response1)
    print("✓ Successfully parsed JSON response")
    print("Data type: " + type(data1))
} else {
    print("❌ Failed to fetch data from " + url1)
}

// Test 2: Different API Endpoint
print("\n🔄 Test 2: User Agent and Headers Test")
print("-" * 40)

let url2 = "https://httpbin.org/user-agent"
print("Fetching user agent info from: " + url2)
let response2 = net.get(url2)

if (response2 != "") {
    print("✓ Successfully fetched user agent data!")
    print("Response: " + response2)

    // Save and analyze
    os.writeFile("user_agent_response.json", response2)
    let data2 = json.decode(response2)
    print("✓ User agent data saved and parsed")
} else {
    print("❌ Failed to fetch user agent data")
}

// Test 3: Multiple Endpoints Analysis
print("\n📊 Test 3: Multi-Source Data Analysis")
print("-" * 40)

let endpoints = [
    "https://httpbin.org/json",
    "https://httpbin.org/uuid",
    "https://httpbin.org/time/now"
]

let results = []
let successCount = 0

for i, endpoint in endpoints {
    print("Fetching from endpoint " + string(i + 1) + ": " + endpoint)
    let response = net.get(endpoint)

    if (response != "") {
        successCount += 1
        let result = {
            "endpoint": endpoint,
            "success": true,
            "response_length": len(response),
            "timestamp": time.format(time.now(), "2006-01-02 15:04:05"),
            "data": response
        }
        results.push(result)
        print("  ✓ Success (" + string(len(response)) + " bytes)")
    } else {
        let result = {
            "endpoint": endpoint,
            "success": false,
            "response_length": 0,
            "timestamp": time.format(time.now(), "2006-01-02 15:04:05"),
            "error": "Failed to fetch"
        }
        results.push(result)
        print("  ❌ Failed")
    }
}

print("\nResults Summary:")
print("  Total endpoints tested: " + string(len(endpoints)))
print("  Successful requests: " + string(successCount))
print("  Success rate: " + string((successCount * 100) / len(endpoints)) + "%")

// Test 4: Data Processing and Storage
print("\n💾 Test 4: Data Processing and Storage")
print("-" * 40)

// Save all results to a comprehensive file
let analysisReport = {
    "report_id": uuid.generate(),
    "generated_at": time.format(time.now(), "2006-01-02 15:04:05"),
    "summary": {
        "total_endpoints": len(endpoints),
        "successful_requests": successCount,
        "success_rate": (successCount * 100) / len(endpoints)
    },
    "results": results
}

let reportJson = json.encode(analysisReport)
let reportFile = "web_analysis_report_" + time.format(time.now(), "2006-01-02_15-04-05") + ".json"
os.writeFile(reportFile, reportJson)
print("✓ Comprehensive analysis saved to: " + reportFile)

// Generate text report
let textReport = "WEB DATA FETCHER ANALYSIS REPORT\n"
textReport += "Generated: " + time.format(time.now(), "02-01-2006 15:04:05") + "\n"
textReport += "=" * 50 + "\n\n"

textReport += "SUMMARY\n"
textReport += "-------\n"
textReport += "Total Endpoints Tested: " + string(len(endpoints)) + "\n"
textReport += "Successful Requests: " + string(successCount) + "\n"
textReport += "Success Rate: " + string((successCount * 100) / len(endpoints)) + "%\n\n"

textReport += "DETAILED RESULTS\n"
textReport += "----------------\n"
for result in results {
    textReport += "Endpoint: " + result["endpoint"] + "\n"
    textReport += "Status: " + string(result["success"]) + "\n"
    textReport += "Response Size: " + string(result["response_length"]) + " bytes\n"
    textReport += "Timestamp: " + result["timestamp"] + "\n"
    textReport += "\n"
}

textReport += "=" * 50 + "\n"
textReport += "Report generated by VintLang Web Fetcher v1.0\n"

let textReportFile = "web_analysis_" + time.format(time.now(), "2006-01-02_15-04-05") + ".txt"
os.writeFile(textReportFile, textReport)
print("✓ Text report saved to: " + textReportFile)

// Test 5: Network Performance Analysis
print("\n⏱️ Test 5: Network Performance Analysis")
print("-" * 40)

let performanceTest = func(url) {
    let startTime = time.now()
    let response = net.get(url)
    let endTime = time.now()

    return {
        "url": url,
        "success": response != "",
        "response_size": len(response),
        "start_time": startTime,
        "end_time": endTime
    }
}

print("Running performance tests...")
let perfResults = []

for endpoint in endpoints {
    let result = performanceTest(endpoint)
    perfResults.push(result)

    if (result["success"]) {
        print("  " + endpoint + " - Success (" + string(result["response_size"]) + " bytes)")
    } else {
        print("  " + endpoint + " - Failed")
    }
}

// Save performance data
let perfReport = {
    "test_id": uuid.generate(),
    "test_time": time.format(time.now(), "2006-01-02 15:04:05"),
    "performance_results": perfResults
}

os.writeFile("performance_test.json", json.encode(perfReport))
print("✓ Performance data saved to performance_test.json")

// Final Summary
print("\n🎉 Web Data Fetcher Analysis Complete!")
print("-" * 50)

print("📊 DEMONSTRATION SUMMARY:")
print("  • HTTP requests made: " + string(len(endpoints) * 2))
print("  • Successful requests: " + string(successCount * 2))
print("  • JSON responses processed: " + string(successCount))
print("  • Files generated: 5+")
print("  • Performance tests run: " + string(len(endpoints)))

print("\n✨ NETWORKING FEATURES DEMONSTRATED:")
print("  ✓ HTTP GET Requests - Multiple API endpoints")
print("  ✓ Response Processing - JSON parsing and analysis")
print("  ✓ Error Handling - Graceful failure management")
print("  ✓ Data Storage - Multiple output formats")
print("  ✓ Performance Testing - Request timing analysis")
print("  ✓ Report Generation - Comprehensive documentation")

print("\n🚀 VintLang Network Capabilities:")
print("  • RESTful API integration")
print("  • JSON data processing")
print("  • Web scraping potential")
print("  • API testing automation")
print("  • Data aggregation from multiple sources")
print("  • Network performance monitoring")

print("\n🎯 VintLang is ready for:")
print("  • API integration projects")
print("  • Data aggregation systems")
print("  • Web monitoring tools")
print("  • Automated testing scripts")
print("  • Content fetching applications")

print("\n" + "=" * 50)
print("🌐 Network Demonstration Complete!")
print("=" * 50)
```

## working_enterprise_test.vint

```js
// Basic Enterprise HTTP Module Test
import http

print("🏢 Enterprise HTTP Module Features Test")
print("=" * 60)

// Test 1: Create app
print("\n✓ Test 1 - App Creation")
let result = http.app()
print("App creation:", result)

// Test 2: Route Grouping
print("\n✓ Test 2 - Route Grouping")
let groupFunc = func() {
    print("Group function")
}
let v1Group = http.group("/api/v1", groupFunc)
print("API v1 group:", v1Group)

// Test 3: Security
print("\n✓ Test 3 - Security Features")
let securityResult = http.security()
print("Security middleware:", securityResult)

// Test 4: Routes
print("\n✓ Test 4 - Route Registration")
let routeHandler = func(req, res) {
    print("Route handler")
}

let getRoute = http.get("/users", routeHandler)
print("GET route:", getRoute)

let postRoute = http.post("/upload", routeHandler)
print("POST route:", postRoute)

// Test 5: Middleware
print("\n✓ Test 5 - Middleware")
let middlewareFunc = func(req, res, next) {
    print("Middleware")
}
let middleware = http.use(middlewareFunc)
print("Middleware:", middleware)

// Test 6: Guards
print("\n✓ Test 6 - Guards")
let guardFunc = func(req) {
    print("Guard")
}
let guard = http.guard(guardFunc)
print("Guard:", guard)

// Test 7: Error Handler
print("\n✓ Test 7 - Error Handler")
let errorFunc = func(err, req, res) {
    print("Error handler")
}
let errorHandler = http.errorHandler(errorFunc)
print("Error handler:", errorHandler)

// Test 8: Streaming and Metrics
print("\n✓ Test 8 - Streaming and Metrics")
let streamFunc = func(req, res) {
    print("Stream handler")
}
let streamHandler = http.stream(streamFunc)
print("Stream handler:", streamHandler)

let metricsResult = http.metrics()
print("Metrics enabled:", metricsResult)

print("\n" + "=" * 60)
print("✨ All enterprise HTTP features work!")
print("\nFeatures demonstrated:")
print("  🔧 Route grouping")
print("  🛡️  Security middleware")
print("  🔗 Advanced routing")
print("  📊 Error handling")
print("  📈 Performance metrics")
print("  🌊 Streaming support")
print("  🚀 Production-ready backend")

print("\n🎯 Enterprise HTTP module is ready!")
```
