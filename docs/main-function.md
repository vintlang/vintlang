# Main Function in VintLang

VintLang now supports a **main function** as an entry point for your programs, similar to Go, Zig, C, and C++. This provides a structured way to organize your code while maintaining full backward compatibility.

## How It Works

### Two-Phase Execution Model

1. **Setup Phase**: All top-level statements are executed to define functions, variables, and perform initialization
2. **Main Phase**: If a `main` function is found, it's automatically called as the program's entry point

### Main Function Syntax

Use VintLang's existing function syntax to define a main function:

```javascript
let main = func() {
    println("Hello from main!")
    return 0
}
```

Or with `const`:

```javascript
const main = func() {
    println("Hello from main!")
    return "success"
}
```

## Complete Example

```javascript
// Setup phase - runs first
import time
println("🚀 Program starting...")

// Define helper functions
let greet = func(name) {
    println("Hello,", name, "!")
}

let calculate = func(a, b) {
    return a + b
}

// Main function - entry point
let main = func() {
    println("=== Main Function ===")
    
    greet("Developer")
    let result = calculate(10, 20)
    println("Result:", result)
    println("Time:", time.now())
    
    return result
}

// More setup
println("⚙️ Setup complete")
```

**Output:**
```
🚀 Program starting...
⚙️ Setup complete
=== Main Function ===
Hello, Developer !
Result: 30
Time: 16:42:00 26-09-2025
30
```

## Main Function Features

### Parameters
Main functions currently don't receive command-line arguments, but this could be extended in the future:

```javascript
let main = func() {
    // No parameters for now
    println("Main executed")
}
```

### Return Values
The main function's return value becomes the program's final result:

```javascript
let main = func() {
    return 42  // This will be printed as the program output
}
```

### Error Handling
If the main function returns an error, it will be propagated:

```javascript
let main = func() {
    if (someCondition) {
        return error("Something went wrong")
    }
    return "success"
}
```

## Backward Compatibility

Programs **without** a main function continue to work exactly as before:

```javascript
// This still works - no main function needed
println("Hello World")
let x = 42
println("x =", x)
```

## When to Use Main Functions

### Use main functions when:
- Building larger, structured programs
- You want clear separation between setup and execution
- Coming from Go, C, C++, or similar languages
- Building command-line tools or applications

### Stick with the traditional approach when:
- Writing simple scripts
- Prototyping or testing small code snippets
- Using VintLang in REPL mode
- Personal preference for simpler structure

## Execution Flow

1. **Parse** the entire program
2. **Setup Phase**: Execute all top-level statements
   - Define variables with `let` and `const`
   - Define functions
   - Run imports
   - Execute setup code
3. **Main Phase**: If `main` function exists, call it
4. **Return** either the main function's result or the last statement's result

## Migration Guide

To convert existing VintLang programs to use main functions:

**Before:**
```javascript
import time
let x = 42
println("Hello World")
println("Time:", time.now())
```

**After:**
```javascript
import time  // Still runs in setup phase

let main = func() {
    let x = 42
    println("Hello World")  
    println("Time:", time.now())
}
```

The behavior remains identical, but the code is now more structured and explicit about the entry point.