# Modules in VintLang

Modules are a way to organize and reuse code in VintLang. This guide explains how to create and use modules.

## Creating a Module

To create a module, use the `module` keyword followed by the module name and a block of code:

```vint
module math {
    func add(a, b) {
        return a + b
    }
    
    func subtract(a, b) {
        return a - b
    }
}
```

## Using Modules

To use a module, you need to:

1. Create a file with the `.vint` extension in one of these locations:
   - Current working directory
   - `./modules` directory
   - `./vintLang/modules` directory

2. Import the module using the `import` statement:

```vint
import math

result = math.add(5, 3)
print(result)  // Output: 8
```

## Module Structure

A module can contain:
- Functions
- Variables
- Other modules
- Any valid VintLang code

Example of a more complex module:

```vint
module utils {
    let version = "1.0.0"
    
    func format(text) {
        return "Formatted: " + text
    }
    
    module helpers {
        func validate(input) {
            return input != null
        }
    }
}
```

## Best Practices

1. Keep modules focused on a single responsibility
2. Use descriptive names for modules
3. Document your modules with comments
4. Place related modules in the same directory
5. Use the `modules` directory for reusable code

## Error Handling

If a module is not found, you'll see an error message like this:

```
Module 'math' not found.

To fix this:
1. Create a file named 'math.vint' in one of these locations:
  1. /current/working/directory
  2. /current/working/directory/modules
2. Make sure the file contains valid VintLang code
3. Try importing again
```

## Module Scope

Variables and functions defined in a module are only accessible within that module unless explicitly exported. This helps prevent naming conflicts and keeps code organized.

## Example: Creating a Custom Module

Here's a complete example of creating and using a custom module:

```vint
// file: modules/calculator.vint
module calculator {
    func add(a, b) {
        return a + b
    }
    
    func subtract(a, b) {
        return a - b
    }
    
    func multiply(a, b) {
        return a * b
    }
    
    func divide(a, b) {
        if b == 0 {
            return "Error: Division by zero"
        }
        return a / b
    }
}

// file: main.vint
import calculator

result1 = calculator.add(10, 5)
result2 = calculator.multiply(4, 3)

print("Addition: " + result1)      // Output: Addition: 15
print("Multiplication: " + result2) // Output: Multiplication: 12
``` 