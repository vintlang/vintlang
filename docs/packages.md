# Packages in VintLang

Packages in VintLang provide a powerful way to organize, encapsulate, and reuse your code. They allow you to group related functions, variables, and state into a single, importable unit, similar to modules or libraries in other languages.

---

## Defining a Package

You can define a package using the `package` keyword, followed by the package name and a block of code enclosed in curly braces `{}`.

A single `.vint` file can contain one package definition. The name of the file does not need to match the package name, but it is good practice to keep them related.

**Syntax:**
```js
package MyPackage {
    // ... package members ...
}
```

### Package Members

Inside a package block, you can define:

- **Variables**: To hold the package's state using `let`.
- **Constants**: Immutable values using the `const` keyword.
- **Functions**: To provide the package's functionality.

#### Constants in Packages

VintLang supports package-level constants using the `const` keyword. Constants are immutable and are often used for configuration values, version numbers, or other fixed data:

```js
package Config {
    const VERSION = "1.2.3"
    const MAX_CONNECTIONS = 100
    const API_BASE_URL = "https://api.example.com"
    
    let getConfig = func() {
        return {
            "version": VERSION,
            "max_conn": MAX_CONNECTIONS,
            "api_url": API_BASE_URL
        }
    }
}
```

#### Public vs Private Members

VintLang supports access control for package members using a naming convention:

- **Public members**: Names that do NOT start with an underscore `_` are accessible from outside the package.
- **Private members**: Names that start with an underscore `_` are only accessible within the package itself.

```js
package MyPackage {
    // Public members (accessible from outside)
    let publicVariable = "I'm accessible"
    const PUBLIC_CONSTANT = 42
    let publicFunction = func() { return "Hello!" }
    
    // Private members (internal use only)
    let _privateVariable = "Internal only"
    const _PRIVATE_KEY = "secret-key-123"
    let _privateFunction = func() { return "Internal helper" }
}
```

Attempting to access private members from outside the package will result in an error:
```js
import "MyPackage"

print(MyPackage.publicVariable)  // ✅ Works
print(MyPackage._privateVariable) // ❌ Error: cannot access private property

---

## The Automatic `init` Function

VintLang's package system includes a special feature for initialization. If you define a function named `init` inside your package, the Vint interpreter will **automatically execute it** when the package is first loaded.

This is useful for setting up initial state, connecting to services, or performing any other setup work the package needs before it can be used.

**Example:**
```js
package Counter {
    let count = 0

    // This function will run automatically
    let init = func() {
        print("Counter package has been initialized!")
        @.count = 100 // Set initial state
    }

    let getCount = func() {
        return @.count
    }
}
```

---

## The `@` Operator: Self-Reference

Inside a package, you may need to refer to the package's own members or state. VintLang provides the special `@` operator for this purpose. The `@` operator is a reference to the package's own scope.

This is similar to `this` or `self` in other object-oriented languages.

You use it with dot notation to access other members within the same package.

**Example:**
```js
package Greeter {
    let greeting = "Hello"

    let setGreeting = func(newGreeting) {
        // Use @ to access the 'greeting' variable
        @.greeting = newGreeting
    }

    let sayHello = func(name) {
        // Use @ to access the 'greeting' variable
        print(@.greeting + ", " + name + "!")
    }
}
```
Using `@` is necessary to distinguish between a package-level variable and a local variable with the same name.

---

## Importing and Using Packages

To use a package, you import the file that contains its definition. The package object is then assigned to a variable with the same name as the package.

1. **Create your package file** (e.g., `utils.vint`).
2. **Import it in another file** (e.g., `main.vint`).
3. **Access its members** using dot notation.

If `utils.vint` contains `package utils { ... }`, you would use it like this:

```js
// main.vint

// Import the file containing the package
import "utils"

// Now you can use the 'utils' package
utils.doSomething()
```

---

## Complete Examples

For comprehensive, runnable examples that demonstrate all package features including constants, private members, and initialization, see the files in the `examples/packages_example/` directory:

- **`enhanced_test.vint`**: Showcases constants, private members, auto-initialization, and complex package functionality.
- **`greeter_pkg.vint`**: Simple package with state management and initialization.
- **`enhanced_system_test.vint`**: Demonstrates how to use packages with private member protection.

### Key Features Summary

✅ **Package-level constants** with `const` keyword
✅ **Private member protection** using underscore `_` prefix
✅ **Auto-initialization** with `init()` functions  
✅ **State management** with the `@` operator
✅ **Comprehensive access control** for variables, constants, and functions

```js
// Complete example demonstrating all features
package EnhancedExample {
    // Public constants
    const VERSION = "2.0.0"
    const MAX_ITEMS = 100
    
    // Private constants  
    const _SECRET_KEY = "internal-key-123"
    
    // Public variables
    let counter = 0
    
    // Private variables
    let _internalState = "hidden"
    
    // Auto-initialization
    let init = func() {
        print("Package initialized! Version:", VERSION)
        @.counter = 10
    }
    
    // Public functions
    let increment = func() {
        @.counter = @.counter + 1
        return @.counter
    }
    
    // Private functions
    let _validate = func(value) {
        return value != null && value > 0
    }
    
    let processValue = func(value) {
        if (!_validate(value)) {
            return "Invalid value"
        }
        return "Processed: " + string(value)
    }
}
```
