# Constants in Vint

Constants are used to declare variables with values that cannot be changed once assigned. This feature helps ensure immutability and prevents accidental reassignments, making your code more robust and predictable.

## Syntax Rules

The `const` keyword is used to declare a constant. It follows the same naming rules as `let`, but with one critical difference: its value is immutable.

- **Must be initialized at declaration.**
- **Cannot be reassigned.**

### Examples of Valid `const` Declarations:

```js
const PI = 3.14159
print(PI)  // Output: 3.14159

const GREETING = "Hello, Vint!"
print(GREETING)  // Output: "Hello, Vint!"
```

In the examples above, `PI` and `GREETING` are declared as constants and can be used throughout the program.

## Immutability

Once a constant is declared, its value cannot be changed. Attempting to reassign a `const` variable will result in an error.

### Example of an Invalid Reassignment:

```js
const MAX_CONNECTIONS = 5
print(MAX_CONNECTIONS) // Output: 5

// This will cause an error
MAX_CONNECTIONS = 10 
// Error: Cannot assign to constant 'MAX_CONNECTIONS'
```

This immutability ensures that critical values in your program remain constant, preventing bugs and making your code easier to reason about.

## Best Practices

1. **Use for Unchanging Values:** Use `const` for values that should not change during the execution of your program, such as mathematical constants, configuration settings, or fixed values.

2. **Use Uppercase for Global Constants:** It's a common convention to use `UPPER_SNAKE_CASE` for global constants to make them easily distinguishable from regular variables.

   ```js
   const API_KEY = "your-secret-key"
   ```

3. **Prefer `const` Over `let`:** Whenever possible, prefer `const` over `let` to make your code safer and more predictable. Only use `let` when you know a variable's value needs to change.

## Constants in Packages

Constants work seamlessly with VintLang's package system and support the same access control features:

```js
package Config {
    // Public constants (accessible from outside)
    const VERSION = "2.1.0"
    const MAX_USERS = 1000
    
    // Private constants (internal use only)
    const _SECRET_KEY = "internal-key-abc123"
    const _DEBUG_MODE = true
    
    let getPublicConfig = func() {
        return {
            "version": VERSION,
            "max_users": MAX_USERS
            // Note: private constants are not exposed
        }
    }
}
```

### Package Constant Usage

```js
import "Config"

// Accessing public constants
print("App Version:", Config.VERSION)  // ✅ Works
print("Max Users:", Config.MAX_USERS)  // ✅ Works

// Attempting to access private constants
print(Config._SECRET_KEY)  // ❌ Error: cannot access private property
```

For more information about packages and access control, see the [Packages documentation](packages.md). 