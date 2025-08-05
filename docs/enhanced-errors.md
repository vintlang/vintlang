# Enhanced Error Messages in Vint Core Modules

This document describes the enhanced error handling system implemented across Vint's core modules using the new `ErrorMessage` helper function. The improvements provide consistent, colorized, and highly descriptive error messages that guide users toward correct usage.

## New ErrorMessage Helper Function

All modules now use a centralized `ErrorMessage` function from `module/module.go`:

```go
func ErrorMessage(module, function, expected, received, usage string) *object.Error {
    return &object.Error{
        Message: fmt.Sprintf(
            "\033[1;31mError in %s.%s()\033[0m:\n"+
                "  Expected: %s\n"+
                "  Received: %s\n"+
                "  Usage: %s\n"+
                "  See documentation for details.\n",
            module, function, expected, received, usage,
        ),
    }
}
```

## What Was Enhanced

### 1. **CLI Module** (`import cli`)
- **Functions**: `cli.prompt()`, `cli.confirm()`, `cli.execCommand()`, `cli.exit()`, `cli.hasArg()`, `cli.getArgValue()`, `cli.getPositional()`
- **Improvements**:
  - Consistent error formatting with colors
  - Clear expected vs received information
  - Practical usage examples
  - Descriptive parameter names

**New Format Example**:
```
Error in cli.prompt():
  Expected: 1 string argument (prompt message)
  Received: 2 arguments
  Usage: cli.prompt("Enter your name: ") -> returns user input
  See documentation for details.
```

### 2. **Net Module** (`import net`)
- **Functions**: `net.get()`, `net.post()`, `net.put()`, `net.delete()`, `net.patch()`
- **Improvements**:
  - Parameter-specific error messages
  - Network operation context
  - HTTP method-specific examples

**New Format Example**:
```
Error in net.get():
  Expected: string value for 'url' parameter
  Received: INTEGER
  Usage: net.get(url="https://example.com")
  See documentation for details.
```

### 3. **OS Module** (`import os`)
- **Functions**: `os.run()`, `os.getEnv()`
- **Improvements**:
  - Command execution context
  - System operation guidance
  - Clear parameter descriptions

### 4. **Math Module** (`import math`)
- **Functions**: `math.abs()` and similar numeric functions
- **Improvements**:
  - Mathematical operation context
  - Numeric type specifications
  - Calculation examples

### 5. **Time Module** (`import time`)
- **Functions**: `time.now()`, `time.sleep()`
- **Improvements**:
  - Time operation context
  - Duration specifications
  - Temporal examples

### 6. **Crypto Module** (`import crypto`)
- **Functions**: `crypto.hashMD5()`, `crypto.hashSHA256()`, `crypto.encryptAES()`, `crypto.decryptAES()`
- **Improvements**:
  - Cryptographic operation context
  - Security-related guidance
  - Encryption examples

### 7. **Colors Module** (`import colors`)
- **Functions**: `colors.rgbToHex()`
- **Improvements**:
  - Color value specifications
  - Range validation
  - Visual examples

### 8. **String Module** (`import string`)
- **Functions**: `string.slug()` and others
- **Improvements**:
  - Text processing context
  - String manipulation examples

## Error Message Format

All enhanced error messages follow this consistent pattern:

```
Error in [module].[function]():
  Expected: [clear description of expected input]
  Received: [what was actually provided]
  Usage: [practical example with expected output]
  See documentation for details.
```

### Key Features:

1. **🎨 Color Coding**: Red highlighting for error identification
2. **📝 Clear Structure**: Consistent four-line format
3. **🔍 Specific Details**: Exact expected vs received information
4. **💡 Usage Examples**: Practical code examples
5. **📚 Documentation Reference**: Pointer to additional help

## Benefits of the New System

### For Developers:
- **Instant Recognition**: Red coloring makes errors immediately visible
- **Clear Guidance**: Know exactly what's expected vs what was provided
- **Learn by Example**: Usage examples teach correct syntax
- **Consistent Experience**: Same error format across all modules
- **Reduced Debugging Time**: Precise error information speeds up fixes

### For the Language:
- **Professional Appearance**: Consistent, polished error messages
- **Better Learning Curve**: New users learn faster with clear examples
- **Maintainability**: Centralized error formatting makes updates easier
- **Extensibility**: Easy to add new modules using the same pattern

## Examples by Category

### Argument Count Errors:
```
Error in time.sleep():
  Expected: 1 numeric argument (seconds to sleep)
  Received: 2 arguments
  Usage: time.sleep(5) -> sleeps for 5 seconds
  See documentation for details.
```

### Type Errors:
```
Error in math.abs():
  Expected: numeric argument (integer or float)
  Received: STRING
  Usage: math.abs(-5) -> 5
  See documentation for details.
```

### Parameter-Specific Errors:
```
Error in net.post():
  Expected: dictionary value for 'headers' parameter
  Received: STRING
  Usage: net.post(headers={"Content-Type": "application/json"})
  See documentation for details.
```

### Range Validation Errors:
```
Error in colors.rgbToHex():
  RGB values must be in the range 0-255.
  Usage: colors.rgbToHex(255, 0, 128) -> "#FF0080"
```

## Testing Enhanced Errors

Use the provided `new_error_format_test.vint` file to see all the improved error messages in action:

```bash
vint new_error_format_test.vint
```

## Implementation Guidelines

When adding new functions or modules, use the `ErrorMessage` helper:

```go
if len(args) != expectedCount {
    return ErrorMessage(
        "moduleName", "functionName",
        "description of expected arguments",
        fmt.Sprintf("%d arguments", len(args)),
        "usage.example() -> expected output",
    )
}

if args[0].Type() != expectedType {
    return ErrorMessage(
        "moduleName", "functionName", 
        "description of expected type",
        string(args[0].Type()),
        "usage.example() -> expected output",
    )
}
```

## Future Enhancements

The ErrorMessage system enables future improvements:

1. **Error Codes**: Add numeric codes for programmatic handling
2. **Suggestions**: Auto-suggest corrections for common mistakes
3. **Localization**: Multi-language error messages
4. **Context Awareness**: Errors that understand the calling context
5. **Interactive Help**: Links to relevant documentation sections

This enhanced error system represents a significant improvement in Vint's developer experience, making the language more approachable and professional.
