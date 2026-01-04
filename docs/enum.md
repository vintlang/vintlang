# Enums in Vint

Enums (enumerations) provide a way to define a type that has a fixed set of named constant values. They make code more readable, maintainable, and type-safe by representing a group of related constants with meaningful names.

## Why Use Enums?

Enums are useful for:
- **State management**: Representing states like `PENDING`, `ACTIVE`, `COMPLETED`
- **Configuration**: Defining fixed sets of options or modes
- **Type safety**: Preventing invalid values by restricting choices to predefined constants
- **Readability**: Using descriptive names instead of magic numbers or strings

## Basic Syntax

```vint
enum EnumName {
    MEMBER1 = value1,
    MEMBER2 = value2,
    MEMBER3 = value3
}
```

### Example: Integer Enum

```vint
enum Status {
    PENDING = 0,
    ACTIVE = 1,
    COMPLETED = 2,
    FAILED = 3
}

let currentStatus = Status.ACTIVE
print(currentStatus)  // Output: 1
```

## Enum Types

### Integer Enums

Integer enums are useful for numeric codes, states, or priority levels:

```vint
enum Priority {
    LOW = 0,
    MEDIUM = 1,
    HIGH = 2,
    CRITICAL = 3
}

let taskPriority = Priority.HIGH
print("Priority level: " + taskPriority)  // Output: Priority level: 2
```

### String Enums

String enums are ideal for configuration values, roles, or descriptive constants:

```vint
enum Environment {
    DEV = "development",
    STAGING = "staging",
    PROD = "production"
}

let currentEnv = Environment.PROD
print("Running in: " + currentEnv)  // Output: Running in: production
```

### HTTP Status Codes

```vint
enum HttpStatus {
    OK = 200,
    CREATED = 201,
    NOTFOUND = 404,
    UNAUTHORIZED = 401,
    SERVERERROR = 500
}

let statusCode = HttpStatus.OK
if statusCode == HttpStatus.OK {
    print("Request successful!")
}
```

## Accessing Enum Members

Enum members are accessed using dot notation:

```vint
enum Color {
    RED = "red",
    GREEN = "green",
    BLUE = "blue"
}

let myColor = Color.RED
let anotherColor = Color.BLUE
```

## Using Enums in Expressions

### Comparisons

Enums can be compared using standard comparison operators:

```vint
enum OrderStatus {
    PENDING = 0,
    CONFIRMED = 1,
    SHIPPED = 2,
    DELIVERED = 3
}

let order = OrderStatus.CONFIRMED

if order == OrderStatus.CONFIRMED {
    print("Order has been confirmed")
}

if order >= OrderStatus.CONFIRMED {
    print("Order is confirmed or further along")
}
```

### Arithmetic Operations

Integer enum values can be used in arithmetic:

```vint
enum Level {
    BEGINNER = 1,
    INTERMEDIATE = 2,
    ADVANCED = 3,
    EXPERT = 4
}

let currentLevel = Level.INTERMEDIATE
let nextLevel = currentLevel + 1
print("Next level: " + nextLevel)  // Output: Next level: 3
```

### String Concatenation

String enum values can be concatenated:

```vint
enum Greeting {
    HELLO = "Hello",
    GOODBYE = "Goodbye",
    WELCOME = "Welcome"
}

let message = Greeting.HELLO + ", World!"
print(message)  // Output: Hello, World!
```

## Immutability

Enums are declared as constants and cannot be reassigned:

```vint
enum Status {
    ACTIVE = 1,
    INACTIVE = 0
}

// This will cause an error:
// Status = 5
// Error: Cannot assign to constant 'Status'
```

Enum members also cannot be modified:

```vint
enum Config {
    MAX_USERS = 100,
    TIMEOUT = 30
}

let maxUsers = Config.MAX_USERS  // ✅ Allowed
// Config.MAX_USERS = 200         // ❌ Not allowed
```

## Common Use Cases

### 1. Application States

```vint
enum AppState {
    INITIALIZING = 0,
    READY = 1,
    LOADING = 2,
    ERROR = 3
}

let state = AppState.INITIALIZING

if state == AppState.ERROR {
    print("Application encountered an error")
}
```

### 2. User Roles

```vint
enum Role {
    ADMIN = "admin",
    MODERATOR = "moderator",
    USER = "user",
    GUEST = "guest"
}

let userRole = Role.ADMIN

if userRole == Role.ADMIN {
    print("Admin access granted")
}
```

### 3. Log Levels

```vint
enum LogLevel {
    DEBUG = 0,
    INFO = 1,
    WARN = 2,
    ERROR = 3,
    FATAL = 4
}

let currentLogLevel = LogLevel.INFO

if currentLogLevel >= LogLevel.WARN {
    print("Warning or higher level message")
}
```

### 4. Direction or Orientation

```vint
enum Direction {
    NORTH = 0,
    EAST = 1,
    SOUTH = 2,
    WEST = 3
}

let facing = Direction.NORTH
```

### 5. API Response Codes

```vint
enum ResponseCode {
    SUCCESS = 200,
    CREATED = 201,
    ACCEPTED = 202,
    BADREQUEST = 400,
    UNAUTHORIZED = 401,
    FORBIDDEN = 403,
    NOTFOUND = 404,
    SERVERERROR = 500
}

let response = ResponseCode.SUCCESS

if response == ResponseCode.SUCCESS {
    print("Operation completed successfully")
}
```

## Working with Functions

Enums can be passed to and returned from functions:

```vint
enum Status {
    PENDING = 0,
    ACTIVE = 1,
    COMPLETED = 2
}

let checkStatus = func(status) {
    if status == Status.COMPLETED {
        return "Done!"
    } else if status == Status.ACTIVE {
        return "In progress..."
    } else {
        return "Waiting..."
    }
}

let result = checkStatus(Status.ACTIVE)
print(result)  // Output: In progress...
```

## Error Handling

Attempting to access a non-existent enum member will result in an error:

```vint
enum Status {
    PENDING = 0,
    ACTIVE = 1
}

// This will cause an error:
// let invalid = Status.INVALID
// Error: Enum 'Status' has no member 'INVALID'
```

## Best Practices

### 1. Use Descriptive Names

Choose clear, descriptive names for both the enum and its members:

```vint
// ✅ Good
enum OrderStatus {
    PENDING = 0,
    CONFIRMED = 1,
    SHIPPED = 2
}

// ❌ Avoid
enum Status {
    S1 = 0,
    S2 = 1,
    S3 = 2
}
```

### 2. Use UPPERCASE for Members

Follow the convention of using `UPPERCASE` for enum member names:

```vint
// ✅ Good
enum Priority {
    LOW = 1,
    MEDIUM = 2,
    HIGH = 3
}

// ❌ Avoid
enum Priority {
    low = 1,
    medium = 2,
    high = 3
}
```

### 3. Group Related Constants

Use enums to group related constants together:

```vint
// ✅ Good
enum FilePermission {
    READ = 4,
    WRITE = 2,
    EXECUTE = 1
}

// ❌ Avoid spreading them as separate constants
const READ = 4
const WRITE = 2
const EXECUTE = 1
```

### 4. Choose Appropriate Values

Use values that make sense for your use case:

```vint
// For states/flags: use 0, 1, 2...
enum State {
    IDLE = 0,
    RUNNING = 1,
    STOPPED = 2
}

// For HTTP codes: use actual HTTP status codes
enum HttpStatus {
    OK = 200,
    NOTFOUND = 404
}

// For descriptive values: use strings
enum Mode {
    LIGHT = "light",
    DARK = "dark"
}
```

### 5. Document Complex Enums

Add comments to explain enum purposes and member meanings:

```vint
// User permission levels for access control
enum Permission {
    NONE = 0,      // No access
    READ = 1,      // Read-only access
    WRITE = 2,     // Read and write access
    ADMIN = 3      // Full administrative access
}
```

## Comparison with Constants

While you can use `const` for individual constants, enums are better for related groups:

```vint
// Using const (verbose)
const STATUS_PENDING = 0
const STATUS_ACTIVE = 1
const STATUS_COMPLETED = 2

// Using enum (cleaner and grouped)
enum Status {
    PENDING = 0,
    ACTIVE = 1,
    COMPLETED = 2
}

// Access is clearer with enums
let status = Status.ACTIVE  // Clear namespace
```

## Advanced Examples

### Switch Statements with Enums

```vint
enum TaskStatus {
    TODO = 0,
    INPROGRESS = 1,
    REVIEW = 2,
    DONE = 3
}

let status = TaskStatus.INPROGRESS

switch status {
    case TaskStatus.TODO {
        print("Task not started")
    }
    case TaskStatus.INPROGRESS {
        print("Task in progress")
    }
    case TaskStatus.REVIEW {
        print("Task under review")
    }
    case TaskStatus.DONE {
        print("Task completed")
    }
}
```

### Using Enums in Loops

```vint
enum Status {
    PENDING = 0,
    ACTIVE = 1,
    COMPLETED = 2
}

let statuses = [Status.PENDING, Status.ACTIVE, Status.COMPLETED]

for status in statuses {
    print("Processing status: " + status)
}
```

### Enums in Dictionaries

```vint
enum Status {
    PENDING = 0,
    ACTIVE = 1,
    COMPLETED = 2
}

let statusMessages = {
    Status.PENDING: "Waiting to start",
    Status.ACTIVE: "Currently processing",
    Status.COMPLETED: "All done!"
}

print(statusMessages[Status.ACTIVE])  // Output: Currently processing
```

## Limitations

1. **No Auto-increment**: Each member must have an explicit value
2. **No Reverse Mapping**: You cannot get member name from value directly
3. **No Methods**: Enums cannot have associated methods

## Summary

Enums in Vint provide:
- ✅ Named constants for fixed sets of values
- ✅ Support for both integer and string values
- ✅ Type safety through immutability
- ✅ Clear, readable code
- ✅ Easy namespace management
- ✅ Integration with all Vint language features

Use enums whenever you have a fixed set of related constants to make your code more maintainable and self-documenting!

