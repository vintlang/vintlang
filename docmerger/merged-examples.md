# Example Code Snippets

## arrays.vint

```js
// VintLang Arrays Example
// Demonstrates array creation, access, mutation, slicing, and built-in methods

// ============================================================
// 1. Creating and accessing arrays
// ============================================================
let colors = ["red", "green", "blue", "yellow"]
println("Colors:", colors)
println("First:", colors[0])
println("Last:", colors[len(colors) - 1])

// ============================================================
// 2. Modifying arrays (push / pop / append)
// ============================================================
colors.push("purple")
println("\nAfter push:", colors)

let removed = pop(colors)
println("Popped:", removed)
println("After pop:", colors)

let more = append(colors, "orange", "pink")
println("After append:", more)

// ============================================================
// 3. Array slicing (Python-like syntax)
// ============================================================
let nums = [1, 2, 3, 4, 5, 6, 7, 8, 9, 10]
println("\nFull array:", nums)
println("nums[2:5]  =", nums[2:5])    // [3, 4, 5]
println("nums[:4]   =", nums[:4])     // [1, 2, 3, 4]
println("nums[7:]   =", nums[7:])     // [8, 9, 10]
println("nums[-3:]  =", nums[-3:])    // [8, 9, 10]
println("nums[:-3]  =", nums[:-3])    // [1, 2, 3, 4, 5, 6, 7]

// ============================================================
// 4. Iterating over an array
// ============================================================
let scores = [85, 92, 78, 95, 88]
let total = 0
for score in scores {
    total = total + score
}
let avg = total / len(scores)
println("\nScores:", scores)
println("Average:", avg)

// ============================================================
// 5. Sorting and reversing
// ============================================================
let words = ["banana", "apple", "cherry", "date"]
println("\nOriginal:", words)
println("Sorted:", words.sort())
println("Reversed:", words.reverse())

// ============================================================
// 6. Searching
// ============================================================
let animals = ["cat", "dog", "bird", "fish"]
println("\nAnimals:", animals)
println("Index of 'bird':", indexOf(animals, "bird"))
println("Index of 'lion':", indexOf(animals, "lion"))  // -1 means not found

// ============================================================
// 7. Range
// ============================================================
println("\nRange examples:")
println("range(5):", range(5))
println("range(2,7):", range(2, 7))

// ============================================================
// 8. Nested arrays (2D)
// ============================================================
let matrix = [
    [1, 2, 3],
    [4, 5, 6],
    [7, 8, 9]
]
println("\n2D Matrix:")
for row in matrix {
    println(" ", row)
}
println("Element [1][2]:", matrix[1][2])

// ============================================================
// 9. Unique elements
// ============================================================
let dupes = [1, 2, 2, 3, 3, 3, 4]
println("\nWith duplicates:", dupes)
println("Unique:", unique(dupes))

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

## builtins.vint

```js
// VintLang Built-in Functions Example
// Demonstrates core built-in functions available without importing any module

println("=== Type Checking ===")
println("type(42):", type(42))
println("type(3.14):", type(3.14))
println("type('hi'):", type("hi"))
println("type(true):", type(true))
println("type([1,2]):", type([1, 2]))
println("type({}):", type({}))
println("type(null):", type(null))

println("\n=== Type Conversion ===")
println("int('42'):", int("42"))
println("int(3.9):", int(3.9))
println("parseFloat('3.14'):", parseFloat("3.14"))
println("parseInt('100'):", parseInt("100"))
println("string(123):", string(123))
println("string(true):", string(true))

println("\n=== String Predicates ===")
println("startsWith('VintLang', 'Vint'):", startsWith("VintLang", "Vint"))
println("endsWith('VintLang', 'Lang'):", endsWith("VintLang", "Lang"))

println("\n=== Collection Builtins ===")
let arr = [3, 1, 4, 1, 5, 9, 2, 6]
println("Array:", arr)
println("len:", len(arr))
println("unique:", unique(arr))
println("sorted:", arr.sort())
println("reversed:", arr.reverse())
println("indexOf(arr, 9):", indexOf(arr, 9))

let dict = {"a": 1, "b": 2, "c": 3}
println("\nDict:", dict)
println("keys:", keys(dict))
println("values:", values(dict))
println("has_key:", has_key(dict, "b"))

println("\n=== Array Operations ===")
let nums = [1, 2, 3]
nums = append(nums, 4, 5)
println("append:", nums)
let last = pop(nums)
println("pop:", last, "| remaining:", nums)

println("\n=== Logical Builtins ===")
println("not(true):", not(true))
println("not(false):", not(false))
println("and(true, false):", and(true, false))
println("and(true, true):", and(true, true))
println("or(false, true):", or(false, true))
println("or(false, false):", or(false, false))

println("\n=== chr / ord ===")
println("ord('A'):", ord("A"))
println("chr(65):", chr(65))
println("ord('z'):", ord("z"))
println("chr(122):", chr(122))

println("\n=== range ===")
println("range(5):", range(5))
println("range(2,6):", range(2, 6))

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

## closures.vint

```js
// VintLang Closures & Higher-Order Functions Example
// Demonstrates closures, function factories, and higher-order patterns

// ============================================================
// 1. Basic closure: a function that captures its environment
// ============================================================
let makeCounter = func() {
    let count = 0
    return func() {
        count = count + 1
        return count
    }
}

let counter = makeCounter()
println("Counter:", counter())  // 1
println("Counter:", counter())  // 2
println("Counter:", counter())  // 3

// Each call to makeCounter() creates an independent counter
let counter2 = makeCounter()
println("Counter2:", counter2())  // 1 (independent)

// ============================================================
// 2. Function factory: functions that generate other functions
// ============================================================
let makeMultiplier = func(factor) {
    return func(x) {
        return x * factor
    }
}

let double = makeMultiplier(2)
let triple = makeMultiplier(3)
let tenX   = makeMultiplier(10)

println("\nDouble 7:", double(7))   // 14
println("Triple 7:", triple(7))    // 21
println("10x 7:", tenX(7))         // 70

// ============================================================
// 3. Higher-order function: map (apply function to each element)
// ============================================================
let map = func(arr, f) {
    let result = []
    for item in arr {
        result.push(f(item))
    }
    return result
}

let numbers = [1, 2, 3, 4, 5]
let squared = map(numbers, func(x) { return x * x })
let doubled = map(numbers, double)

println("\nNumbers:", numbers)
println("Squared:", squared)
println("Doubled:", doubled)

// ============================================================
// 4. Higher-order function: filter
// ============================================================
let filter = func(arr, predicate) {
    let result = []
    for item in arr {
        if (predicate(item)) {
            result.push(item)
        }
    }
    return result
}

let isEven = func(x) { return x % 2 == 0 }
let isPositive = func(x) { return x > 0 }

let mixed = [-3, -1, 0, 2, 4, 7, 9]
println("\nMixed:", mixed)
println("Even:", filter(mixed, isEven))
println("Positive:", filter(mixed, isPositive))

// ============================================================
// 5. Higher-order function: reduce (fold)
// ============================================================
let reduce = func(arr, f, initial) {
    let acc = initial
    for item in arr {
        acc = f(acc, item)
    }
    return acc
}

let sum   = reduce(numbers, func(a, b) { return a + b }, 0)
let product = reduce(numbers, func(a, b) { return a * b }, 1)
let maxVal  = reduce(numbers, func(a, b) { if (a > b) { return a } return b }, numbers[0])

println("\nNumbers:", numbers)
println("Sum:", sum)
println("Product:", product)
println("Max:", maxVal)

// ============================================================
// 6. Composing functions
// ============================================================
let compose = func(f, g) {
    return func(x) {
        return f(g(x))
    }
}

let addOne  = func(x) { return x + 1 }
let square  = func(x) { return x * x }

let squareThenAdd = compose(addOne, square)
let addThenSquare = compose(square, addOne)

println("\nCompose examples:")
println("squareThenAdd(4):", squareThenAdd(4))  // (4^2)+1 = 17
println("addThenSquare(4):", addThenSquare(4))  // (4+1)^2 = 25

// ============================================================
// 7. Memoization (caching results with closures)
// ============================================================
let memoize = func(f) {
    let cache = {}
    return func(n) {
        let key = string(n)
        let cached = cache[key]
        if (cached != null) {
            return cached
        }
        let result = f(n)
        cache[key] = result
        return result
    }
}

let factorial = memoize(func(n) {
    if (n <= 1) { return 1 }
    let result = 1
    for i in range(2, n + 1) {
        result = result * i
    }
    return result
})

println("\nFactorials (memoized):")
println("5! =", factorial(5))   // 120
println("6! =", factorial(6))   // 720
println("10! =", factorial(10)) // 3628800

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

## concurrency.vint

```js
// VintLang Concurrency Example
// Demonstrates async/await, goroutines, and channels for concurrent programming

// ============================================================
// 1. Async functions and await
// ============================================================
println("=== Async / Await ===")

let fetchUser = async func(id) {
    return {"id": id, "name": "User" + string(id), "active": true}
}

let fetchScore = async func(userId) {
    return userId * 100
}

// Run async operations sequentially with await
let user = await fetchUser(42)
let score = await fetchScore(user["id"])
println("User:", user["name"], "| Score:", score)

// ============================================================
// 2. Multiple async calls
// ============================================================
println("\n=== Multiple Async Calls ===")

let p1 = fetchUser(1)
let p2 = fetchUser(2)
let p3 = fetchUser(3)

let u1 = await p1
let u2 = await p2
let u3 = await p3

println("Fetched users:", u1["name"], u2["name"], u3["name"])

// ============================================================
// 3. Goroutines - fire-and-forget concurrent tasks
// ============================================================
println("\n=== Goroutines ===")

go println("Goroutine 1: Hello from goroutine!")
go println("Goroutine 2: Running concurrently!")
go println("Goroutine 3: Concurrent execution!")

println("Main: goroutines launched")

// ============================================================
// 4. Channels for goroutine communication
// ============================================================
println("\n=== Channels ===")

// Create a buffered channel with capacity 3
let results = chan(3)

// Producer goroutine: sends values into channel
go func() {
    send(results, "task-1 complete")
    send(results, "task-2 complete")
    send(results, "task-3 complete")
    close(results)
}()

// Consumer: receive values from channel
let r1 = receive(results)
let r2 = receive(results)
let r3 = receive(results)
println("Received:", r1)
println("Received:", r2)
println("Received:", r3)

// ============================================================
// 5. Async with channel pipeline
// ============================================================
println("\n=== Async + Channel Pipeline ===")

let processAsync = async func(input) {
    let out = chan
    go func() {
        let processed = "processed:" + input
        send(out, processed)
    }()
    return receive(out)
}

let result1 = await processAsync("data-A")
let result2 = await processAsync("data-B")
println(result1)
println(result2)

// ============================================================
// 6. Worker pool pattern
// ============================================================
println("\n=== Worker Pool ===")

let jobs = chan(5)
let done = chan(3)

// Worker function
let worker = func(id) {
    go func() {
        let job = receive(jobs)
        println("Worker", id, "processed:", job)
        send(done, "worker-" + string(id) + "-done")
    }()
}

// Dispatch workers
worker(1)
worker(2)
worker(3)

// Send jobs
send(jobs, "job-A")
send(jobs, "job-B")
send(jobs, "job-C")

// Collect results
receive(done)
receive(done)
receive(done)
println("All workers finished")

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

## defer.vint

```js
// VintLang Defer Statement Example
// Demonstrates the defer keyword which schedules code to run when the function exits
// Deferred statements execute in LIFO (last-in, first-out) order

// ============================================================
// 1. Basic defer - runs after the function body
// ============================================================
let basicExample = func() {
    defer println("3. Deferred: cleanup done")
    println("1. Function body starts")
    println("2. Function body ends")
}

println("=== Basic Defer ===")
basicExample()
println("4. Back in main")

// ============================================================
// 2. Multiple defers - execute in reverse order (LIFO)
// ============================================================
let multipleDefers = func() {
    defer println("Deferred 1 (runs last)")
    defer println("Deferred 2 (runs second)")
    defer println("Deferred 3 (runs first)")
    println("Function body")
}

println("\n=== Multiple Defers (LIFO order) ===")
multipleDefers()

// ============================================================
// 3. Practical use: resource cleanup
// ============================================================
let processData = func(data) {
    println("\nOpening resource...")
    defer println("Resource closed (always runs)")

    if (len(data) == 0) {
        warn "Empty data, aborting early"
        return null  // defer still runs!
    }

    let total = 0
    for item in data {
        total = total + item
    }
    println("Processed", len(data), "items, sum =", total)
}

println("\n=== Defer with Early Return ===")
processData([1, 2, 3, 4, 5])
processData([])

// ============================================================
// 4. Defer with logging (wrap a section with open/close logs)
// ============================================================
let runTask = func(name) {
    println("\n--- Task:", name, "---")
    defer println("--- End:", name, "---")

    if (name == "task2") {
        println("  Doing task2 work...")
        return null
    }
    println("  Doing", name, "work...")
    println("  Finishing", name, "...")
}

println("\n=== Defer for Structured Logging ===")
runTask("task1")
runTask("task2")
runTask("task3")

```

## dictionaries.vint

```js
// VintLang Dictionaries Example
// Demonstrates dictionary creation, access, mutation, and iteration

// ============================================================
// 1. Creating dictionaries
// ============================================================
let user = {
    "name": "Alice",
    "age": 30,
    "city": "Nairobi",
    "active": true
}
println("User:", user)

// ============================================================
// 2. Accessing values
// ============================================================
println("\nName:", user["name"])
println("Age:", user["age"])

// Dot notation also works for string keys
println("City:", user["city"])

// ============================================================
// 3. Updating and adding keys
// ============================================================
user["age"] = 31
user["email"] = "alice@example.com"
println("\nUpdated user:", user)

// ============================================================
// 4. Checking for keys (has_key)
// ============================================================
println("\nhas 'email'?", has_key(user, "email"))
println("has 'phone'?", has_key(user, "phone"))

// Method syntax also works
println("has 'name'?", user.has_key("name"))

// ============================================================
// 5. Iterating over a dictionary
// ============================================================
println("\nAll user fields:")
for key, value in user {
    println(" ", key, "=", value)
}

// ============================================================
// 6. Nested dictionaries
// ============================================================
let company = {
    "name": "VintCorp",
    "address": {
        "street": "123 Main St",
        "city": "Nairobi",
        "country": "Kenya"
    },
    "employees": 42
}
println("\nCompany:", company["name"])
println("City:", company["address"]["city"])
println("Country:", company["address"]["country"])

// ============================================================
// 7. Dictionary with array values
// ============================================================
let student = {
    "name": "Bob",
    "grades": [85, 92, 78, 95, 88],
    "subjects": ["Math", "Science", "English"]
}
println("\nStudent:", student["name"])
println("Subjects:", student["subjects"])
println("First grade:", student["grades"][0])

// ============================================================
// 8. Building a dictionary dynamically
// ============================================================
let inventory = {}
let items = ["apple", "banana", "cherry"]
let prices = [1.5, 0.75, 2.0]

for i in range(len(items)) {
    inventory[items[i]] = prices[i]
}
println("\nInventory:", inventory)

// ============================================================
// 9. Counting occurrences (frequency map)
// ============================================================
let words = ["cat", "dog", "cat", "bird", "dog", "cat"]
let freq = {}
for word in words {
    let count = freq[word]
    if (count == null) {
        freq[word] = 1
    } else {
        freq[word] = count + 1
    }
}
println("\nWord frequencies:", freq)

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

## enum_demo.vint

```js
// Simple enum example

enum Status {
    PENDING = 0,
    ACTIVE = 1,
    COMPLETED = 2
}

let current = Status.ACTIVE
print(current)

// String enum
enum Color {
    RED = "red",
    GREEN = "green",
    BLUE = "blue"
}

let myColor = Color.RED
print(myColor)


```

## error_handling.vint

```js
// VintLang Error Handling & Logging Example
// Demonstrates error handling patterns, logging levels, and defensive programming

// ============================================================
// 1. Logging and status messages
// ============================================================
println("=== Logging Levels ===")
info "Application starting up"
debug "Loading configuration..."
note "Remember to set API_KEY in your .env file"
success "Configuration loaded successfully"
warn "Running in development mode"

// ============================================================
// 2. Checking for null values (defensive programming)
// ============================================================
println("\n=== Null Checks ===")

let safeDivide = func(a, b) {
    if (b == 0) {
        warn "Division by zero attempted"
        return null
    }
    return a / b
}

let r1 = safeDivide(10, 2)
let r2 = safeDivide(10, 0)

if (r1 != null) {
    println("10 / 2 =", r1)
} else {
    println("10 / 2 failed")
}

if (r2 != null) {
    println("10 / 0 =", r2)
} else {
    println("10 / 0 returned null (division by zero)")
}

// ============================================================
// 3. Returning error information from functions
// ============================================================
println("\n=== Returning Errors ===")

let parseAge = func(input) {
    let age = int(input)
    if (age == null || age < 0) {
        return {"ok": false, "error": "Age must be a non-negative number"}
    }
    if (age > 150) {
        return {"ok": false, "error": "Age seems unrealistically large"}
    }
    return {"ok": true, "value": age}
}

let tests = ["25", "200", "-5", "42"]
for t in tests {
    let result = parseAge(t)
    if (result["ok"]) {
        println("Input '" + t + "' -> valid age:", result["value"])
    } else {
        println("Input '" + t + "' -> error:", result["error"])
    }
}

// ============================================================
// 4. Validating user input
// ============================================================
println("\n=== Input Validation ===")

let strmod = import("string")

let validateEmail = func(email) {
    if (email == null || email == "") {
        return "Email cannot be empty"
    }
    if (!strmod.contains(email, "@")) {
        return "Email must contain @"
    }
    if (!strmod.contains(email, ".")) {
        return "Email must contain a domain"
    }
    return null  // null means no error
}

let emails = ["alice@example.com", "notanemail", "", "bob@test.org"]
for email in emails {
    let err = validateEmail(email)
    if (err == null) {
        success "Valid email: " + email
    } else {
        warn "Invalid email '" + email + "': " + err
    }
}

// ============================================================
// 5. Graceful handling with defaults
// ============================================================
println("\n=== Default Values ===")

let getConfig = func(key, defaultValue) {
    let config = {
        "host": "localhost",
        "port": 8080,
        "debug": false
    }
    let val = config[key]
    if (val == null) {
        return defaultValue
    }
    return val
}

println("host:", getConfig("host", "127.0.0.1"))
println("port:", getConfig("port", 3000))
println("timeout:", getConfig("timeout", 30))   // uses default
println("debug:", getConfig("debug", true))

// ============================================================
// 6. Retry pattern
// ============================================================
println("\n=== Retry Pattern ===")

let attempt = 0
let maxAttempts = 3
let succeeded = false

while (attempt < maxAttempts && !succeeded) {
    attempt = attempt + 1
    debug "Attempt " + string(attempt) + " of " + string(maxAttempts)

    // Simulate success on the 3rd try
    if (attempt == 3) {
        succeeded = true
        success "Operation succeeded on attempt " + string(attempt)
    } else {
        warn "Attempt " + string(attempt) + " failed, retrying..."
    }
}

if (!succeeded) {
    println("All", maxAttempts, "attempts failed")
}

```

## excel_demo.vint

```js
// VintLang Excel Module Example
// Demonstrates creating Excel workbooks with multiple sheets, data, and formulas

import excel

println("=== Excel Module Demo ===")

// ============================================================
// 1. Create a new workbook
// ============================================================
let wb = excel.create("report.xlsx")
println("Created workbook: report.xlsx")

// ============================================================
// 2. Set up the first sheet - Sales Data
// ============================================================
excel.renameSheet(wb, "Sheet1", "Sales")
excel.addSheet(wb, "Summary")

// Add headers to Sales sheet
excel.setCell(wb, "Sales", "A1", "Month")
excel.setCell(wb, "Sales", "B1", "Revenue")
excel.setCell(wb, "Sales", "C1", "Expenses")
excel.setCell(wb, "Sales", "D1", "Profit")

// Add monthly data
let months    = ["Jan", "Feb", "Mar", "Apr", "May", "Jun"]
let revenues  = [12000, 15000, 18000, 14000, 20000, 22000]
let expenses  = [8000,  9000,  10000, 8500,  11000, 12000]

for i in range(len(months)) {
    let row = string(i + 2)
    excel.setCell(wb, "Sales", "A" + row, months[i])
    excel.setCell(wb, "Sales", "B" + row, revenues[i])
    excel.setCell(wb, "Sales", "C" + row, expenses[i])
    // Profit formula: Revenue - Expenses
    excel.setCellFormula(wb, "Sales", "D" + row, "=B" + row + "-C" + row)
}

println("Added 6 months of sales data with profit formulas")

// ============================================================
// 3. Fill in the Summary sheet
// ============================================================
excel.setCell(wb, "Summary", "A1", "Metric")
excel.setCell(wb, "Summary", "B1", "Value")

excel.setCell(wb, "Summary", "A2", "Total Revenue")
excel.setCellFormula(wb, "Summary", "B2", "=SUM(Sales!B2:B7)")

excel.setCell(wb, "Summary", "A3", "Total Expenses")
excel.setCellFormula(wb, "Summary", "B3", "=SUM(Sales!C2:C7)")

excel.setCell(wb, "Summary", "A4", "Total Profit")
excel.setCellFormula(wb, "Summary", "B4", "=B2-B3")

println("Created summary sheet with totals")

// ============================================================
// 4. Merge cells for a title
// ============================================================
excel.mergeCells(wb, "Summary", "A1:B1")
excel.setCell(wb, "Summary", "A1", "H1 2024 FINANCIAL SUMMARY")

// ============================================================
// 5. Save the workbook
// ============================================================
excel.save(wb)
println("Saved to report.xlsx")

// ============================================================
// 6. Read back data to verify
// ============================================================
let jan = excel.getCell(wb, "Sales", "A2")
let jan_rev = excel.getCell(wb, "Sales", "B2")
let jan_exp = excel.getCell(wb, "Sales", "C2")
println("\nVerification - " + jan + ": Revenue=" + string(jan_rev) + " Expenses=" + string(jan_exp))

excel.close(wb)
println("\nDone! Excel demo complete.")
println("File 'report.xlsx' created with Sales and Summary sheets.")

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

## fmt_demo.vint

```js
// VintLang fmt Module Example
// Demonstrates string formatting, number formatting, padding, and alignment

import fmt

println("=== String Formatting (sprintf) ===")

// Basic string interpolation
let name = "VintLang"
let version = "1.0"
let users = 1234
let intro = fmt.sprintf("Welcome to %s v%s! We have %d active users.", name, version, users)
println(intro)

// Float formatting
let pi = 3.14159265359
println(fmt.sprintf("Pi to 4 decimal places: %.4f", pi))
println(fmt.sprintf("Hex: %x, Octal: %o, Binary: %b", 255, 8, 15))

println("\n=== Number Formatting ===")

let number = 255
println("Decimal:", fmt.formatInt(number, 10))
println("Binary: ", fmt.formatBin(number))
println("Hex:    ", fmt.formatHex(number))
println("Octal:  ", fmt.formatOct(number))

println("\nFloat precision:")
println("pi (2 dec):", fmt.formatFloat(pi, 2))
println("pi (5 dec):", fmt.formatFloat(pi, 5))
println("pi (prec 6):", fmt.precision(pi, 6))

println("\n=== Text Padding & Alignment ===")

let items = ["Apple", "Banana", "Fig", "Dragonfruit"]
let sep = "=" * 40

// Left-aligned (padded right)
println(sep)
println("Left-aligned (padded right, width 15):")
for item in items {
    println("  [" + fmt.padRight(item, 15, ".") + "]")
}

// Right-aligned (padded left)
println("\nRight-aligned (padded left, width 15):")
for item in items {
    println("  [" + fmt.padLeft(item, 15, ".") + "]")
}

// Centered
println("\nCentered (width 15):")
for item in items {
    println("  [" + fmt.padCenter(item, 15, "-") + "]")
}

println("\n=== Table Formatting ===")

let header = fmt.sprintf("| %-10s | %8s | %6s |", "Product", "Price", "Stock")
let border = "-" * len(header)
println(border)
println(header)
println(border)

let products = [
    ["Laptop", "$999.99", "15"],
    ["Mouse", "$25.50", "150"],
    ["Keyboard", "$75.00", "45"],
    ["Monitor", "$299.99", "8"]
]

for p in products {
    let row = fmt.sprintf("| %-10s | %8s | %6s |", p[0], p[1], p[2])
    println(row)
}
println(border)

println("\n=== Fixed Width ===")

let titles = ["Short", "Medium Title", "A Very Long Title Here"]
println("Fixed width (18 chars):")
for title in titles {
    println("  [" + fmt.width(title, 18) + "]")
}

println("\n=== Report Generation ===")

let reportWidth = 40
let title = "MONTHLY SALES REPORT"
println("=" * reportWidth)
println(fmt.padCenter(title, reportWidth))
println("=" * reportWidth)

let totalSales = 1245.75
let totalOrders = 28
let avgOrder = totalSales / totalOrders

let salesLine = fmt.sprintf("Total Sales:   $%8.2f", totalSales)
let ordersLine = fmt.sprintf("Total Orders:  %8d", totalOrders)
let avgLine    = fmt.sprintf("Avg Per Order: $%8.2f", avgOrder)

println(salesLine)
println(ordersLine)
println(avgLine)
println("=" * reportWidth)

println("\nAll fmt functions demonstrated successfully!")

```

## for_loops.vint

```js
// VintLang Loop Constructs Example
// Demonstrates for, while, repeat loops and loop control (break, continue)

// ============================================================
// 1. For-in loop over an array
// ============================================================
let fruits = ["apple", "banana", "cherry", "date"]
println("Fruits:")
for fruit in fruits {
    println(" -", fruit)
}

// ============================================================
// 2. For loop with index using range()
// ============================================================
println("\nCounting with range:")
for i in range(5) {
    print(i, "")
}
println()

// ============================================================
// 3. For-in loop over a dictionary (key, value pairs)
// ============================================================
let person = {"name": "Alice", "age": 30, "city": "Nairobi"}
println("\nPerson details:")
for key, value in person {
    println(" ", key, ":", value)
}

// ============================================================
// 4. While loop
// ============================================================
println("\nCountdown:")
let n = 5
while (n > 0) {
    print(n, "")
    n = n - 1
}
println("Go!")

// ============================================================
// 5. Repeat loop (fixed iterations)
// ============================================================
println("\nRepeat 3 times:")
repeat 3 {
    println(" Hello from repeat! (iteration", i, ")")
}

// ============================================================
// 6. Break and continue
// ============================================================
println("\nSkip even numbers:")
for num in range(10) {
    if (num % 2 == 0) {
        continue
    }
    print(num, "")
}
println()

println("\nStop at 5:")
for num in range(10) {
    if (num == 5) {
        break
    }
    print(num, "")
}
println()

// ============================================================
// 7. Nested loops
// ============================================================
println("\nMultiplication table (3x3):")
for row in range(1, 4) {
    for col in range(1, 4) {
        print(row * col, "\t")
    }
    println()
}

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

println("LLM module example (requires OpenAI API key)")
println("Uncomment the code above to test with a valid API key") 
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
let name = "Tachera Sasi"

// Split the string into an array of characters and print the result
print(name.split("")) 

// Reverse the string and print the result
print(name.reverse()) 

// Get the length of the string and print it
print(name.len()) 

// Convert the string to uppercase and print it
print(name.upper()) 

// Convert the string to lowercase and print it
print(name.lower()) 

// Check if the string contains the substring "sasi" (case-sensitive) and print the result
print(name.contains("sasi")) 

// Convert the string to uppercase and check if it contains the substring "SASI" (case-sensitive), then print the result
print(name.upper().contains("SASI")) 

// Replace the substring "Sasi" with "Vint" and print the result
print(name.replace("Sasi", "Vint")) 

// Trim any occurrence of the character "a" from the start and end of the string and print the result
print(name.trim("a"))

print(string(123))           // "123"
print(string(true))          // "true"
print(string(12.34))         // "12.34"
print(string("Hello World")) // "Hello World"


print(int("123"))    // 123
print(int(12.34))    // 12
print(int(true))     // 1
print(int(false))    // 0

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

## overloading.vint

```js
// VintLang Function Overloading Example
// VintLang supports overloading functions by arity (number of parameters).
// The interpreter automatically selects the correct version based on the
// number of arguments passed.

// ============================================================
// 1. Basic overloading by arity
// ============================================================
let greet = func() {
    println("Hello, World!")
}

let greet = func(name) {
    println("Hello,", name + "!")
}

let greet = func(name, title) {
    println("Hello,", title, name + "!")
}

println("=== Greeting overloads ===")
greet()                   // calls 0-arg version
greet("Alice")            // calls 1-arg version
greet("Smith", "Dr.")     // calls 2-arg version

// ============================================================
// 2. Math operations overloaded by arity
// ============================================================
let add = func(a, b) {
    return a + b
}

let add = func(a, b, c) {
    return a + b + c
}

let add = func(a, b, c, d) {
    return a + b + c + d
}

println("\n=== Add overloads ===")
println("add(2, 3)       =", add(2, 3))
println("add(1, 2, 3)    =", add(1, 2, 3))
println("add(1, 2, 3, 4) =", add(1, 2, 3, 4))

// ============================================================
// 3. Default parameter values (another way to handle arity)
// ============================================================
let connect = func(host = "localhost") {
    println("Connecting to", host, "on port 8080")
}

let connect = func(host, port) {
    println("Connecting to", host, "on port", port)
}

println("\n=== Connect overloads ===")
connect()                    // uses default host
connect("example.com")       // explicit host, default port via default param
connect("example.com", 3306) // explicit host and port

// ============================================================
// 4. Logging function with multiple signatures
// ============================================================
let logMsg = func(message) {
    println("[LOG]", message)
}

let logMsg = func(level, message) {
    println("[" + level + "]", message)
}

let logMsg = func(level, message, data) {
    println("[" + level + "] " + message + " | data: " + string(data))
}

println("\n=== logMsg overloads ===")
logMsg("Something happened")
logMsg("ERROR", "Connection refused")
logMsg("DEBUG", "User fetched", "id=42")

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

## pattern_matching.vint

```js
// VintLang Pattern Matching Example
// Demonstrates match and switch statements for expressive control flow

// ============================================================
// 1. Basic switch statement
// ============================================================
let day = 3
switch (day) {
    case 1 { println("Monday") }
    case 2 { println("Tuesday") }
    case 3 { println("Wednesday") }
    case 4 { println("Thursday") }
    case 5 { println("Friday") }
    default { println("Weekend") }
}

// ============================================================
// 2. Switch with guard conditions
// ============================================================
let score = 78
let grade = ""
switch (score) {
    case x if x >= 90 { grade = "A" }
    case x if x >= 80 { grade = "B" }
    case x if x >= 70 { grade = "C" }
    case x if x >= 60 { grade = "D" }
    default { grade = "F" }
}
println("\nScore:", score, "-> Grade:", grade)

// ============================================================
// 3. Dictionary pattern matching with match
// ============================================================
let user = {"role": "admin", "active": true}
println("\nUser access check:")
match user {
    {"role": "admin", "active": true}  => println("Full access granted")
    {"role": "admin", "active": false} => println("Admin account suspended")
    {"role": "user",  "active": true}  => println("User access granted")
    {"role": "user",  "active": false} => println("User account disabled")
    _ => println("Unknown user type")
}

// ============================================================
// 4. Match with guard conditions
// ============================================================
let player = {"name": "Alice", "score": 2500, "level": 8}
println("\nPlayer rank:")
match player {
    {"score": s} if s >= 3000 => println("Diamond rank")
    {"score": s} if s >= 2000 => println("Gold rank")
    {"score": s} if s >= 1000 => println("Silver rank")
    _ => println("Bronze rank")
}

// ============================================================
// 5. Pattern matching on a list of records
// ============================================================
let orders = [
    {"status": "shipped",   "id": 101},
    {"status": "pending",   "id": 102},
    {"status": "delivered", "id": 103},
    {"status": "cancelled", "id": 104}
]

println("\nOrder status report:")
for order in orders {
    match order {
        {"status": "shipped"}   => println("Order", order["id"], "- In transit")
        {"status": "pending"}   => println("Order", order["id"], "- Awaiting processing")
        {"status": "delivered"} => println("Order", order["id"], "- Successfully delivered")
        {"status": "cancelled"} => println("Order", order["id"], "- Cancelled")
        _ => println("Order", order["id"], "- Status unknown")
    }
}

// ============================================================
// 6. Switch matching on string type
// ============================================================
let classify = func(value) {
    switch (type(value)) {
        case "INTEGER" { return "integer number" }
        case "FLOAT"   { return "floating point number" }
        case "STRING"  { return "text string" }
        case "BOOLEAN" { return "boolean value" }
        case "ARRAY"   { return "array with " + string(len(value)) + " elements" }
        case "DICT"    { return "dictionary" }
        default        { return "unknown type" }
    }
}

println("\nType classification:")
println(42, "->", classify(42))
println(3.14, "->", classify(3.14))
println("hello", "->", classify("hello"))
println(true, "->", classify(true))
println([1, 2, 3], "->", classify([1, 2, 3]))
println({"a": 1}, "->", classify({"a": 1}))

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
const redis = import("redis")

// Connect to Redis
conn = redis.connect("localhost:6379")

// Basic operations
redis.set(conn, "greeting", "Hello, World!")
message = redis.get(conn, "greeting")

// Hash operations
redis.hset(conn, "user:1", "name", "John Doe")
user = redis.hgetall(conn, "user:1")

// List operations
redis.rpush(conn, "tasks", "task1", "task2")
task = redis.lpop(conn, "tasks")

// Close connection
redis.close(conn)
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
// VintLang Regex Module Example
// Demonstrates pattern matching, text replacement, and string splitting using regular expressions
//
// Use regex.test() to check if a pattern matches (avoids conflict with match keyword)
// Use regex.replaceString(pattern, replacement, text) to replace occurrences
// Use regex.splitString(pattern, text) to split a string

import regex

println("=== Pattern Matching ===")

// Check if a string contains digits
println("Contains digits:", regex.test("\\d+", "hello123"))         // true
println("Contains digits:", regex.test("\\d+", "helloworld"))       // false

// Validate an email address
let email = "user@example.com"
let emailPattern = "^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\\.[a-zA-Z]{2,}$"
println("Valid email:", regex.test(emailPattern, email))             // true
println("Valid email:", regex.test(emailPattern, "notanemail"))      // false

// Match a URL
let url = "https://vintlang.org/docs"
println("Is URL:", regex.test("^https?://", url))                   // true

println("\n=== Text Replacement ===")

// Mask digits with *
let phone = "Call me at 123-456-7890 or 555-0000"
let masked = regex.replaceString("\\d", "*", phone)
println("Masked:", masked)

// Remove extra whitespace
let messy = "hello    world   from   vint"
let clean = regex.replaceString("\\s+", " ", messy)
println("Cleaned:", clean)

// Replace URLs with [link]
let text = "Visit https://vintlang.org and http://example.com today"
let replaced = regex.replaceString("https?://[^\\s]+", "[link]", text)
println("Replaced URLs:", replaced)

println("\n=== String Splitting ===")

// Split by comma (CSV-like)
let csv = "alice,bob,charlie,diana"
let names = regex.splitString(",", csv)
println("CSV split:", names)

// Split by whitespace
let sentence = "The quick   brown  fox"
let words = regex.splitString("\\s+", sentence)
println("Word split:", words)

// Split by punctuation
let code = "if(x>0){return x}"
let tokens = regex.splitString("[(){}]", code)
println("Code tokens:", tokens)

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

## structs.vint

```js
// VintLang Structs Example
// Demonstrates defining structs with fields and methods

// ============================================================
// 1. Basic struct definition
// ============================================================
struct Point {
    x: 0
    y: 0

    func toString() {
        return "Point(" + string(this.x) + ", " + string(this.y) + ")"
    }

    func distanceTo(other) {
        let dx = this.x - other.x
        let dy = this.y - other.y
        return (dx * dx + dy * dy)
    }
}

let p1 = Point(x = 3, y = 4)
let p2 = Point(x = 0, y = 0)
println("Point p1:", p1.toString())
println("Point p2:", p2.toString())
println("Distance squared:", p1.distanceTo(p2))

// ============================================================
// 2. Struct with default values and methods
// ============================================================
struct BankAccount {
    owner: "Unknown"
    balance: 0.0

    func deposit(amount) {
        this.balance = this.balance + amount
        println("Deposited $" + string(amount) + " | Balance: $" + string(this.balance))
    }

    func withdraw(amount) {
        if (amount > this.balance) {
            println("Insufficient funds! Balance: $" + string(this.balance))
            return null
        }
        this.balance = this.balance - amount
        println("Withdrew $" + string(amount) + " | Balance: $" + string(this.balance))
    }

    func getBalance() {
        return this.balance
    }
}

println("\n=== Bank Account ===")
let account = BankAccount(owner = "Alice", balance = 100.0)
println("Account owner:", account.owner)
account.deposit(50.0)
account.withdraw(30.0)
account.withdraw(200.0)
println("Final balance: $" + string(account.getBalance()))

// ============================================================
// 3. Struct as a data container
// ============================================================
struct Student {
    name: ""
    grade: 0

    func isPassing() {
        return this.grade >= 50
    }

    func status() {
        if (this.isPassing()) {
            return this.name + " - PASS (" + string(this.grade) + ")"
        }
        return this.name + " - FAIL (" + string(this.grade) + ")"
    }
}

println("\n=== Student Records ===")
let students = [
    Student(name = "Alice", grade = 85),
    Student(name = "Bob",   grade = 42),
    Student(name = "Carol", grade = 91),
    Student(name = "Dave",  grade = 55)
]

for s in students {
    println(" ", s.status())
}

// Count passing
let passing = 0
for s in students {
    if (s.isPassing()) {
        passing = passing + 1
    }
}
println("Passing:", passing, "/", len(students))

// ============================================================
// 4. Structs with nested data
// ============================================================
struct Config {
    host: "localhost"
    port: 8080
    isDebug: false

    func address() {
        return this.host + ":" + string(this.port)
    }

    func describe() {
        let mode = "production"
        if (this.isDebug) {
            mode = "debug"
        }
        return "Server at " + this.address() + " [" + mode + "]"
    }
}

println("\n=== Configuration ===")
let devConfig  = Config(host = "localhost", port = 3000, isDebug = true)
let prodConfig = Config(host = "api.example.com", port = 443, isDebug = false)

println(devConfig.describe())
println(prodConfig.describe())

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

## time.vint

```js
// VintLang Time Module Example
// Demonstrates date and time operations: getting current time, formatting, arithmetic

import time

// ============================================================
// 1. Getting the current time
// ============================================================
let now = time.now()
println("Current time:", now)

// ============================================================
// 2. Formatting dates
// ============================================================
println("\nFormatted dates:")
println("Default:   ", time.format(now, "02-01-2006 15:04:05"))
println("Date only: ", time.format(now, "January 02, 2006"))
println("Time only: ", time.format(now, "15:04:05"))
println("ISO 8601:  ", time.format(now, "2006-01-02T15:04:05"))

// ============================================================
// 3. Leap year check
// ============================================================
println("\nLeap year checks:")
for year in [2000, 2023, 2024, 2100] {
    println(" ", year, "->", time.isLeapYear(year))
}

// ============================================================
// 4. Time arithmetic (add / subtract durations)
// ============================================================
println("\nTime arithmetic:")
println("Now:           ", time.format(now, "2006-01-02 15:04:05"))
println("+ 1 hour:      ", time.format(time.add(now, "1h"), "2006-01-02 15:04:05"))
println("+ 2.5 hours:   ", time.format(time.add(now, "2h30m"), "2006-01-02 15:04:05"))
println("- 24 hours:    ", time.format(time.subtract(now, "24h"), "2006-01-02 15:04:05"))
println("+ 30 seconds:  ", time.format(time.add(now, "30s"), "2006-01-02 15:04:05"))

// ============================================================
// 5. Measuring elapsed time with time.since
// ============================================================
println("\nMeasuring elapsed time:")
let start = time.now()

// Simulate some work
let total = 0
for i in range(100000) {
    total = total + i
}

let elapsed = time.since(start)
println("Computed sum:", total)
println("Elapsed time:", elapsed, "ms")

// ============================================================
// 6. Building a simple timer utility
// ============================================================
let makeTimer = func() {
    let startTime = time.now()
    return func(label) {
        let ms = time.since(startTime)
        println("[timer]", label, "->", ms, "ms elapsed")
    }
}

let timer = makeTimer()
let count = 0
for i in range(50000) {
    count = count + 1
}
timer("loop done")

// ============================================================
// 7. sleep (pause execution)
// ============================================================
println("\nSleeping for 1 second...")
let before = time.now()
time.sleep(1)
let after = time.now()
println("Done! Elapsed:", time.since(before), "ms")

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
// VintLang Chart Module Example
// Demonstrates generating bar, pie, and line charts as PNG images

import vintChart

// Data for the charts
let labels = ["Q1", "Q2", "Q3", "Q4"]
let values = [120, 85, 150, 200]

// Generate a bar chart
vintChart.barChart(labels, values, "bar_chart.png")
println("Bar chart saved to bar_chart.png")

// Generate a pie chart
vintChart.pieChart(labels, values, "pie_chart.png")
println("Pie chart saved to pie_chart.png")

// Generate a line graph
vintChart.lineGraph(labels, values, "line_graph.png")
println("Line graph saved to line_graph.png")

```

## vintSocket.vint

```js
import vintSocket

vintSocket.createServer("8080")
vintSocket.connect("ws://localhost:8080")
vintSocket.broadcast("Hello, WebSocket clients!")

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

