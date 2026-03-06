# Structs in Vint

Structs provide a way to define custom data types with named fields and methods. They let you model real-world entities, group related data together, and attach behavior — making your code more organized, readable, and reusable.

## Why Use Structs?

Structs are useful for:

- **Data modeling**: Representing entities like users, products, or configurations
- **Encapsulation**: Grouping related data and behavior together
- **Reusability**: Creating multiple instances from a single blueprint
- **Readability**: Using descriptive field names instead of array indices or dictionary keys
- **Methods**: Attaching functions directly to your data types

## Basic Syntax

```vint
struct StructName {
    field1: defaultValue1
    field2: defaultValue2

    func methodName(params) {
        // method body — use 'this' to access fields
    }
}
```

### Example: Simple Point

```vint
struct Point {
    x: 0
    y: 0
}

let p = Point(x = 10, y = 20)
print(p.x)  // Output: 10
print(p.y)  // Output: 20
```

---

## Defining Structs

### Fields with Default Values

Every field in a struct must have a default value. This serves as the fallback when no value is provided during instantiation:

```vint
struct Config {
    host: "localhost"
    port: 8080
    verbose: false
}

let cfg = Config()  // All defaults used
print(cfg.host)     // Output: localhost
print(cfg.port)     // Output: 8080
```

### Fields with Methods

Structs can contain both fields and methods. Methods are defined inline using the `func` keyword and can access the struct's fields through the `this` keyword:

```vint
struct User {
    name: ""
    age: 0

    func greet() {
        return "Hello, I'm " + this.name
    }

    func isAdult() {
        return this.age >= 18
    }
}

let user = User(name = "Alice", age = 30)
print(user.greet())    // Output: Hello, I'm Alice
print(user.isAdult())  // Output: true
```

### Fields Only (No Methods)

Structs don't need methods — they work perfectly as simple data containers:

```vint
struct Color {
    r: 0
    g: 0
    b: 0
}

let red = Color(r = 255, g = 0, b = 0)
print("R=" + string(red.r) + " G=" + string(red.g) + " B=" + string(red.b))
// Output: R=255 G=0 B=0
```

---

## Creating Instances

### Named Arguments

The clearest way to create an instance — specify field names explicitly:

```vint
struct User {
    name: ""
    age: 0
}

let user = User(name = "Alice", age = 30)
```

### Positional Arguments

Arguments are matched to fields in declaration order:

```vint
struct Point {
    x: 0
    y: 0
}

let p = Point(5, 15)  // x = 5, y = 15
print(p.x)  // Output: 5
print(p.y)  // Output: 15
```

### Using Default Values

Omit arguments to use the default values:

```vint
struct Config {
    host: "localhost"
    port: 8080
    verbose: false
}

// All defaults
let cfg1 = Config()

// Override some, keep rest as defaults
let cfg2 = Config(host = "example.com")
print(cfg2.host)  // Output: example.com
print(cfg2.port)  // Output: 8080
```

### Partial Overrides

You can override only the fields you care about:

```vint
struct Server {
    host: "localhost"
    port: 8080
    protocol: "http"
}

let prod = Server(host = "api.example.com", port = 443, protocol = "https")
let dev = Server(port = 3000)

print(dev.host)      // Output: localhost
print(dev.port)      // Output: 3000
print(dev.protocol)  // Output: http
```

---

## The `this` Keyword

Inside methods, `this` refers to the current struct instance. Use it to:

- **Read fields**: `this.fieldName`
- **Write fields**: `this.fieldName = newValue`
- **Call other methods**: `this.methodName()`

```vint
struct Counter {
    count: 0

    func increment() {
        this.count = this.count + 1
    }

    func decrement() {
        this.count = this.count - 1
    }

    func value() {
        return this.count
    }
}

let c = Counter()
c.increment()
c.increment()
c.increment()
c.decrement()
print(c.value())  // Output: 2
```

---

## Property Access and Assignment

### Reading Fields

Use dot notation to access a struct instance's fields:

```vint
struct Person {
    name: "Unknown"
    age: 0
}

let p = Person(name = "Alice", age = 25)
print(p.name)  // Output: Alice
print(p.age)   // Output: 25
```

### Mutating Fields

Fields can be reassigned after instantiation:

```vint
let user = User(name = "Alice", age = 30)
print(user.name)  // Output: Alice

user.name = "Bob"
user.age = 25
print(user.name)  // Output: Bob
print(user.age)   // Output: 25
```

### Field Validation

Assigning to a field that doesn't exist on the struct produces an error:

```vint
struct Point {
    x: 0
    y: 0
}

let p = Point(x = 1, y = 2)
// p.z = 3  // ❌ Error: Struct 'Point' has no field 'z'
```

---

## Methods

### Basic Methods

Methods are defined inside the struct body with the `func` keyword:

```vint
struct Rectangle {
    width: 0
    height: 0

    func area() {
        return this.width * this.height
    }

    func perimeter() {
        return 2 * (this.width + this.height)
    }
}

let rect = Rectangle(width = 10, height = 5)
print(rect.area())       // Output: 50
print(rect.perimeter())  // Output: 30
```

### Methods with Parameters

Methods can accept parameters just like regular functions:

```vint
struct Calculator {
    value: 0

    func add(n) {
        return this.value + n
    }

    func multiply(n) {
        return this.value * n
    }
}

let calc = Calculator(value = 10)
print(calc.add(5))       // Output: 15
print(calc.multiply(3))  // Output: 30
```

### Methods with Default Parameters

Method parameters can have default values:

```vint
struct Greeter {
    name: "World"

    func greet(greeting = "Hello") {
        return greeting + ", " + this.name + "!"
    }
}

let g = Greeter(name = "VintLang")
print(g.greet())       // Output: Hello, VintLang!
print(g.greet("Hi"))   // Output: Hi, VintLang!
```

### Methods That Mutate State

Methods can modify the instance's fields through `this`:

```vint
struct Account {
    owner: ""
    balance: 0

    func deposit(amount) {
        this.balance = this.balance + amount
        return this.balance
    }

    func withdraw(amount) {
        if (amount > this.balance) {
            return "Insufficient funds"
        }
        this.balance = this.balance - amount
        return this.balance
    }

    func summary() {
        return this.owner + ": $" + string(this.balance)
    }
}

let acc = Account(owner = "Alice", balance = 100)
print(acc.summary())    // Output: Alice: $100

acc.deposit(50)
print(acc.summary())    // Output: Alice: $150

acc.withdraw(30)
print(acc.summary())    // Output: Alice: $120
```

### Methods Calling Other Methods

Methods can call other methods on the same instance using `this`:

```vint
struct Rectangle {
    width: 0
    height: 0

    func area() {
        return this.width * this.height
    }

    func perimeter() {
        return 2 * (this.width + this.height)
    }

    func describe() {
        return "Rectangle(" + string(this.width) + "x" + string(this.height) +
               ") area=" + string(this.area()) +
               " perimeter=" + string(this.perimeter())
    }
}

let rect = Rectangle(width = 10, height = 5)
print(rect.describe())
// Output: Rectangle(10x5) area=50 perimeter=30
```

---

## Instance Independence

Each instance has its own copy of fields. Mutating one instance does not affect others:

```vint
struct User {
    name: ""
    age: 0

    func greet() {
        return "Hello, I'm " + this.name
    }
}

let alice = User(name = "Alice", age = 30)
let bob = User(name = "Bob", age = 25)

print(alice.greet())  // Output: Hello, I'm Alice
print(bob.greet())    // Output: Hello, I'm Bob

alice.name = "Charlie"
print(alice.greet())  // Output: Hello, I'm Charlie
print(bob.greet())    // Output: Hello, I'm Bob (unchanged)
```

---

## Type Checking

Use the built-in `type()` function to inspect struct types:

```vint
struct User {
    name: ""
    age: 0
}

let u = User(name = "Alice", age = 30)

print(type(u))     // Output: User
print(type(User))  // Output: struct:User
```

- `type(instance)` returns the struct's name (e.g., `"User"`)
- `type(StructDefinition)` returns `"struct:Name"` (e.g., `"struct:User"`)

---

## Methods Returning New Instances

Methods can create and return new struct instances. This is useful for immutable-style operations:

```vint
struct Vector {
    x: 0
    y: 0

    func scale(factor) {
        return Vector(x = this.x * factor, y = this.y * factor)
    }

    func display() {
        return "Vector(" + string(this.x) + ", " + string(this.y) + ")"
    }
}

let v1 = Vector(x = 3, y = 4)
let v2 = v1.scale(2)

print(v1.display())  // Output: Vector(3, 4) — original unchanged
print(v2.display())  // Output: Vector(6, 8)
```

---

## Common Use Cases

### 1. Data Models

```vint
struct Product {
    name: ""
    price: 0
    quantity: 0

    func total() {
        return this.price * this.quantity
    }

    func display() {
        return this.name + " - $" + string(this.price) + " x " + string(this.quantity) + " = $" + string(this.total())
    }
}

let item = Product(name = "Widget", price = 25, quantity = 4)
print(item.display())  // Output: Widget - $25 x 4 = $100
```

### 2. State Management

```vint
struct Counter {
    count: 0

    func increment() {
        this.count = this.count + 1
    }

    func decrement() {
        this.count = this.count - 1
    }

    func reset() {
        this.count = 0
    }

    func value() {
        return this.count
    }
}

let c = Counter()
c.increment()
c.increment()
c.increment()
c.decrement()
print(c.value())  // Output: 2
c.reset()
print(c.value())  // Output: 0
```

### 3. Configuration Objects

```vint
struct DatabaseConfig {
    host: "localhost"
    port: 5432
    name: "app_db"
    user: "admin"
    password: ""

    func connectionString() {
        return this.user + "@" + this.host + ":" + string(this.port) + "/" + this.name
    }
}

let devDB = DatabaseConfig()
print(devDB.connectionString())
// Output: admin@localhost:5432/app_db

let prodDB = DatabaseConfig(
    host = "db.production.com",
    port = 5433,
    name = "prod_db",
    user = "prod_admin",
    password = "secret"
)
print(prodDB.connectionString())
// Output: prod_admin@db.production.com:5433/prod_db
```

### 4. Builder Pattern

```vint
struct HTMLBuilder {
    tag: "div"
    content: ""

    func render() {
        return "<" + this.tag + ">" + this.content + "</" + this.tag + ">"
    }
}

let heading = HTMLBuilder(tag = "h1", content = "Welcome to VintLang")
let paragraph = HTMLBuilder(tag = "p", content = "Structs are powerful!")

print(heading.render())    // Output: <h1>Welcome to VintLang</h1>
print(paragraph.render())  // Output: <p>Structs are powerful!</p>
```

### 5. Todo List

```vint
struct TodoItem {
    title: ""
    completed: false

    func status() {
        if (this.completed) {
            return "[x] " + this.title
        } else {
            return "[ ] " + this.title
        }
    }
}

let todos = [
    TodoItem(title = "Buy groceries"),
    TodoItem(title = "Write code", completed = true),
    TodoItem(title = "Read docs")
]

for item in todos {
    print(item.status())
}
// Output:
// [ ] Buy groceries
// [x] Write code
// [ ] Read docs
```

### 6. Geometry

```vint
struct Circle {
    cx: 0
    cy: 0
    radius: 1

    func containsPoint(px, py) {
        let dx = this.cx - px
        let dy = this.cy - py
        return (dx * dx + dy * dy) <= (this.radius * this.radius)
    }
}

let c = Circle(cx = 0, cy = 0, radius = 5)
print(c.containsPoint(3, 4))  // Output: true
print(c.containsPoint(6, 0))  // Output: false
```

---

## Structs in Collections

### In Arrays

```vint
struct Student {
    name: ""
    grade: 0
}

let students = [
    Student(name = "Alice", grade = 95),
    Student(name = "Bob", grade = 82),
    Student(name = "Charlie", grade = 91)
]

for s in students {
    print(s.name + ": " + string(s.grade))
}
```

### In Dictionaries

```vint
struct Student {
    name: ""
    grade: 0

    func passing() {
        return this.grade >= 60
    }
}

let roster = {
    "s1": Student(name = "Alice", grade = 95),
    "s2": Student(name = "Bob", grade = 55),
    "s3": Student(name = "Charlie", grade = 72)
}

for key, val in roster {
    let status = "PASS"
    if (!val.passing()) {
        status = "FAIL"
    }
    print(val.name + ": " + string(val.grade) + " - " + status)
}
```

---

## Error Handling

VintLang provides clear error messages for common struct mistakes:

### Unknown Field in Constructor

```vint
struct Point {
    x: 0
    y: 0
}

// let p = Point(x = 1, z = 3)
// ❌ Error: Struct 'Point' has no field 'z'
```

### Too Many Positional Arguments

```vint
struct Point {
    x: 0
    y: 0
}

// let p = Point(1, 2, 3)
// ❌ Error: Too many arguments for struct 'Point'
```

### Assigning to a Non-Existent Field

```vint
struct Point {
    x: 0
    y: 0
}

let p = Point(x = 1, y = 2)
// p.z = 3
// ❌ Error: Struct 'Point' has no field 'z'
```

### Calling a Non-Existent Method

```vint
struct Point {
    x: 0
    y: 0
}

let p = Point(x = 1, y = 2)
// p.move()
// ❌ Error: Struct 'Point' has no method 'move'
```

---

## Structs vs Packages

Both structs and packages group data and behavior, but they serve different purposes:

| Feature             | Struct                                 | Package                             |
| ------------------- | -------------------------------------- | ----------------------------------- |
| Multiple instances  | ✅ Yes — create as many as you need    | ❌ No — singleton                   |
| Fields per instance | ✅ Each instance has own fields        | N/A — shared state                  |
| Self-reference      | `this`                                 | `@`                                 |
| Constructor         | `StructName(args)`                     | N/A                                 |
| Use case            | Data modeling, entities, value objects | Organizing code, utilities, modules |

**Use structs** when you need multiple instances with their own state (users, products, shapes).

**Use packages** when you need a single namespace for utility functions, constants, or shared state.

---

## Best Practices

### 1. Use Descriptive Names

Choose clear, descriptive names for structs and their fields:

```vint
// ✅ Good
struct HttpRequest {
    method: "GET"
    url: ""
    headers: {}
}

// ❌ Avoid
struct Req {
    m: ""
    u: ""
    h: {}
}
```

### 2. Always Provide Defaults

Every field should have a sensible default value so that partial construction works:

```vint
struct Config {
    host: "localhost"
    port: 8080
    timeout: 30
    retries: 3
}

// Works — only override what you need
let cfg = Config(host = "production.com")
```

### 3. Keep Methods Focused

Each method should do one thing well:

```vint
// ✅ Good — single responsibility
struct User {
    name: ""
    email: ""

    func displayName() {
        return this.name
    }

    func contactInfo() {
        return this.name + " <" + this.email + ">"
    }
}
```

### 4. Use Named Arguments for Clarity

Prefer named arguments over positional when creating instances:

```vint
// ✅ Clear intent
let user = User(name = "Alice", age = 30)

// ❌ Less readable
let user = User("Alice", 30)
```

### 5. Avoid Reserved Words as Field Names

VintLang has reserved keywords (like `error`, `debug`, `info`, `break`, `return`) that cannot be used as field names. Use descriptive alternatives:

```vint
// ✅ Good
struct LogEntry {
    message: ""
    level: "info"
    verbose: false
}

// ❌ Won't work — 'error' and 'debug' are reserved keywords
// struct LogEntry {
//     error: ""
//     debug: false
// }
```

---

## Summary

Structs in Vint provide:

- ✅ Custom data types with named fields and default values
- ✅ Methods with `this` binding for accessing and mutating fields
- ✅ Named and positional argument constructors
- ✅ Instance independence — each instance owns its data
- ✅ Method parameters with default values
- ✅ Methods calling other methods via `this`
- ✅ Type checking with `type()`
- ✅ Clear error messages for invalid operations
- ✅ Full compatibility with arrays, dictionaries, and other Vint features

Use structs whenever you need to model entities with data and behavior — they're the foundation for building well-structured VintLang applications!
