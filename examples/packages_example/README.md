# Package Examples

This directory contains comprehensive examples demonstrating VintLang's package system features.

## Files Overview

### `enhanced_test.vint`
**Advanced Package Demonstration**

This package showcases all the advanced features of VintLang's package system:

- ✅ **Package-level constants** using `const` keyword
- ✅ **Private member protection** with underscore `_` prefix
- ✅ **Auto-initialization** with `init()` function
- ✅ **State management** using the `@` operator
- ✅ **Complex functionality** with public/private functions

**Features demonstrated:**
- Public constants: `VERSION`, `AUTHOR`, `MAX_ITEMS`, `PI`, `DEBUG_MODE`
- Private constants: `_PRIVATE_KEY`, `_INTERNAL_LIMIT`
- Public variables with state management
- Private variables for internal use
- Auto-initialization setting up initial state
- Public API functions
- Private helper functions for internal logic
- Error handling and validation

### `greeter_pkg.vint`
**Simple Package Example**

A basic package demonstrating:
- Simple initialization
- State management
- Basic public functions

### `enhanced_system_test.vint`
**Package Usage Example**

Shows how to:
- Import and use packages
- Access public members
- Handle package state
- Work with package functions

## Key Features Demonstrated

### 1. Constants in Packages
```js
package MyPackage {
    const VERSION = "1.0.0"           // Public constant
    const _INTERNAL_VERSION = "dev"   // Private constant
}
```

### 2. Private Member Protection
```js
package SecurePackage {
    let publicData = "accessible"
    let _privateData = "hidden"       // Cannot be accessed from outside
    
    let publicFunction = func() { }
    let _privateHelper = func() { }   // Cannot be called from outside
}
```

### 3. Auto-Initialization
```js
package AutoInit {
    let state = 0
    
    // Runs automatically when package is loaded
    let init = func() {
        print("Package initialized!")
        @.state = 100  // Set initial state
    }
}
```

### 4. Access Control Errors

When trying to access private members from outside:

```js
import "./enhanced_test"

// ✅ These work - public access
print(enhanced_test.VERSION)
print(enhanced_test.getVersion())

// ❌ These fail - private access blocked
print(enhanced_test._privateSecret)    // Error: cannot access private property
enhanced_test._validateInput("test")   // Error: function not accessible
```

## Running the Examples

To run any of these examples:

```bash
# Run the enhanced package test
./vint examples/packages_example/enhanced_test.vint

# Test package usage
./vint examples/packages_example/enhanced_system_test.vint
```

## Security Model

| Member Type | Public (`name`) | Private (`_name`) |
|-------------|----------------|-------------------|
| Variables | ✅ Accessible | ❌ Blocked |
| Constants | ✅ Accessible | ❌ Blocked |
| Functions | ✅ Callable | ❌ Blocked |

The underscore `_` prefix convention provides enterprise-grade access control, ensuring that internal implementation details remain hidden while exposing a clean public API.