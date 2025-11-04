# Pattern Matching in VintLang

VintLang supports powerful pattern matching through the `match` statement, allowing you to destructure data, bind variables, and use guard conditions for complex control flow.

## Basic Match Syntax

The `match` statement uses the `=>` arrow syntax to map patterns to actions:

```js
let user = {"role": "admin", "name": "Alice"}

match user {
    {"role": "admin"} => print("Admin user!")
    {"role": "user"} => print("Regular user")
    _ => print("Unknown user type")
}
// Output: Admin user!
```

## Dictionary Pattern Matching

### Simple Dictionary Patterns

Match specific key-value pairs in dictionaries:

```js
let request = {"method": "GET", "path": "/api/users", "status": 200}

match request {
    {"method": "GET"} => print("GET request")
    {"method": "POST"} => print("POST request")
    {"status": 404} => print("Not found")
    _ => print("Other request")
}
// Output: GET request
```

### Variable Binding in Dictionary Patterns

Extract values from dictionaries and bind them to variables:

```js
let user = {"name": "Alice", "age": 30, "role": "admin"}

match user {
    {"name": name, "role": "admin"} => print("Admin:", name)
    {"name": name, "age": age} => print("User:", name, "Age:", age)
    _ => print("Unknown format")
}
// Output: Admin: Alice
```

### Nested Dictionary Patterns

Match complex nested structures:

```js
let config = {
    "database": {
        "host": "localhost",
        "port": 5432
    },
    "cache": {
        "enabled": true
    }
}

match config {
    {"database": {"host": host, "port": port}} => {
        print("Database at", host + ":" + str(port))
    }
    {"cache": {"enabled": enabled}} => {
        print("Cache enabled:", enabled)
    }
    _ => print("Unknown config")
}
// Output: Database at localhost:5432
```

## Array Pattern Matching

### Basic Array Patterns

Match arrays based on their structure:

```js
let arrays = [[], [1], [1, 2], [1, 2, 3]]

for arr in arrays {
    match arr {
        [] => print("Empty array")
        [single] => print("Single element:", single)
        [first, second] => print("Two elements:", first, second)
        _ => print("More than two elements")
    }
}
// Output:
// Empty array
// Single element: 1
// Two elements: 1 2
// More than two elements
```

### Variable Binding in Array Patterns

Extract and bind array elements:

```js
let coordinates = [[0, 0], [3, 4], [1, 2, 3]]

for coord in coordinates {
    match coord {
        [x, y] => print("2D coordinate:", x, y)
        [x, y, z] => print("3D coordinate:", x, y, z)
        _ => print("Unknown coordinate format")
    }
}
// Output:
// 2D coordinate: 0 0
// 2D coordinate: 3 4
// 3D coordinate: 1 2 3
```

## Guard Conditions

### Basic Guards

Add conditions to patterns using the `if` keyword:

```js
let users = [
    {"name": "Alice", "age": 25, "role": "admin"},
    {"name": "Bob", "age": 17, "role": "user"},
    {"name": "Charlie", "age": 35, "role": "user"}
]

for user in users {
    match user {
        {"role": "admin", "name": name} => print("Admin:", name)
        {"age": age, "name": name} if age < 18 => print("Minor:", name)
        {"age": age, "name": name} if age >= 18 => print("Adult:", name)
        _ => print("Unknown user")
    }
}
// Output:
// Admin: Alice
// Minor: Bob
// Adult: Charlie
```

### Complex Guards

Combine multiple conditions in guards:

```js
let events = [
    {"type": "error", "severity": "high", "message": "Database down"},
    {"type": "warning", "count": 5},
    {"type": "info", "user": "alice"}
]

for event in events {
    match event {
        {"type": "error", "severity": severity} if severity == "high" => {
            print("CRITICAL ERROR!")
        }
        {"type": "warning", "count": count} if count > 3 => {
            print("Many warnings:", count)
        }
        {"type": type, "user": user} => {
            print("User event:", type, "for", user)
        }
        _ => print("Other event")
    }
}
// Output:
// CRITICAL ERROR!
// Many warnings: 5
// User event: info for alice
```

## Advanced Pattern Matching

### Multiple Pattern Matching

Combine different pattern types:

```js
let data = [
    {"users": [{"name": "Alice"}, {"name": "Bob"}]},
    {"config": {"debug": true}},
    [1, 2, 3]
]

for item in data {
    match item {
        {"users": users} => {
            print("Found", len(users), "users")
        }
        {"config": {"debug": debug}} if debug => {
            print("Debug mode enabled")
        }
        [first, second, third] => {
            print("Array with three numbers:", first, second, third)
        }
        _ => print("Unknown data format")
    }
}
// Output:
// Found 2 users
// Debug mode enabled
// Array with three numbers: 1 2 3
```

### Type-Based Pattern Matching

Use guards to match based on types:

```js
let mixed = [42, "hello", true, {"key": "value"}, [1, 2, 3]]

for item in mixed {
    match item {
        x if type(x) == "INTEGER" && x > 0 => print("Positive integer:", x)
        x if type(x) == "STRING" => print("String:", x)
        x if type(x) == "BOOLEAN" => print("Boolean:", x)
        x if type(x) == "DICT" => print("Dictionary with keys:", keys(x))
        x if type(x) == "ARRAY" => print("Array with", len(x), "elements")
        _ => print("Unknown type")
    }
}
// Output:
// Positive integer: 42
// String: hello
// Boolean: true
// Dictionary with keys: ["key"]
// Array with 3 elements
```

## Practical Examples

### HTTP Request Router

```js
let handleRequest = func(request) {
    match request {
        {"method": "GET", "path": "/"} => "Home page"
        {"method": "GET", "path": path} if path.startsWith("/api/") => {
            "API endpoint: " + path
        }
        {"method": "POST", "path": "/users", "body": body} => {
            "Creating user: " + body["name"]
        }
        {"method": method, "path": path} => {
            "Unsupported: " + method + " " + path
        }
        _ => "Invalid request"
    }
}
```

### Configuration Validation

```js
let validateConfig = func(config) {
    match config {
        {"database": {"host": host, "port": port}} if len(host) > 0 && port > 0 => {
            print("✓ Valid database config")
        }
        {"redis": {"url": url}} if len(url) > 0 => {
            print("✓ Valid Redis config")
        }
        {"logging": {"level": level}} if level in ["debug", "info", "warn", "error"] => {
            print("✓ Valid logging config")
        }
        _ => {
            print("❌ Invalid configuration")
        }
    }
}
```

### Data Processing Pipeline

```js
let processMessage = func(message) {
    match message {
        {"type": "user_action", "action": "login", "user": user} => {
            logUserLogin(user)
        }
        {"type": "system_event", "event": event, "severity": severity} if severity >= 3 => {
            alertSystemEvent(event)
        }
        {"type": "data_update", "table": table, "records": records} if len(records) > 0 => {
            updateDatabase(table, records)
        }
        {"type": type} => {
            print("Unhandled message type:", type)
        }
        _ => {
            print("Invalid message format")
        }
    }
}
```

## Best Practices

### 1. Order Patterns from Specific to General

```js
// ✅ Correct - specific patterns first
match user {
    {"role": "admin", "active": true} => handleActiveAdmin()
    {"role": "admin"} => handleInactiveAdmin()
    {"role": role} => handleUser(role)
    _ => handleUnknown()
}
```

### 2. Use Meaningful Variable Names

```js
// ❌ Unclear
match request {
    {"a": a, "b": b} => process(a, b)
}

// ✅ Clear
match request {
    {"method": method, "body": body} => process(method, body)
}
```

### 3. Leverage Guards for Complex Logic

```js
// ✅ Use guards instead of nested conditions
match user {
    {"age": age, "role": role} if age >= 18 && role == "admin" => {
        grantAdminAccess()
    }
    {"age": age} if age < 18 => {
        requireParentalConsent()
    }
    _ => handleRegularUser()
}
```

### 4. Handle All Cases

Always include a wildcard pattern (`_`) to handle unexpected cases:

```js
match data {
    {"type": "A"} => handleA()
    {"type": "B"} => handleB()
    _ => handleUnknown() // Don't forget this!
}
```

Pattern matching in VintLang provides a powerful and expressive way to handle complex data structures and control flow, making your code more readable and maintainable.
