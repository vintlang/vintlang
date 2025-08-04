# Enhanced Error Messages in Vint Core Modules

This document describes the enhanced error handling system implemented across Vint's core modules. The improvements provide more descriptive, helpful error messages that guide users toward correct usage.

## What Was Enhanced

### 1. **CLI Module** (`import cli`)
- **Function**: `cli.prompt()`, `cli.confirm()`, `cli.execCommand()`, `cli.exit()`, `cli.hasArg()`, `cli.getArgValue()`, `cli.getPositional()`
- **Improvements**:
  - Specific argument count errors with usage examples
  - Type validation with expected types
  - Empty string validation
  - Status code range validation for `exit()`
  - Command existence checks

**Before**: `"Argument must be a string"`
**After**: `"cli.prompt() expects a string argument, but received INTEGER. Usage: cli.prompt(\"Enter your name: \")"`

### 2. **Net Module** (`import net`)
- **Functions**: `net.get()`, `net.post()`, `net.put()`, `net.delete()`, `net.patch()`
- **Improvements**:
  - Parameter-specific error messages
  - URL validation with examples
  - Body serialization error details
  - Network connectivity guidance
  - HTTP status context

**Before**: `"URL must be a string"`
**After**: `"net.get() 'url' parameter must be a string, but received INTEGER. Usage: net.get(url=\"https://example.com\")"`

### 3. **OS Module** (`import os`)
- **Functions**: `os.run()`, `os.getEnv()`
- **Improvements**:
  - Command execution error details
  - Exit code information
  - Empty command validation
  - Environment variable name validation

**Before**: `"Failed to execute command: error"`
**After**: `"os.run() failed to execute 'invalidcmd': command exited with status 127. This usually indicates the command encountered an error."`

### 4. **Math Module** (`import math`)
- **Functions**: `math.abs()` and similar numeric functions
- **Improvements**:
  - Function-specific error messages
  - Usage examples
  - Type expectations
  - Keyword argument validation

**Before**: `"The argument must be a number"`
**After**: `"math.abs() expects a number argument, but received STRING. Usage: math.abs(-5)"`

### 5. **Time Module** (`import time`)
- **Functions**: `time.now()`, `time.sleep()`
- **Improvements**:
  - No-argument validation for `now()`
  - Negative duration validation
  - Type checking with examples

**Before**: `"Only numbers are allowed as arguments"`
**After**: `"time.sleep() expects a number argument, but received 'not a number'. Usage: time.sleep(5) to sleep for 5 seconds"`

## Error Message Format

All enhanced error messages follow this consistent pattern:

```
[module].[function]() [description of issue] [received vs expected] [usage example]
```

### Examples:

1. **Argument Count Errors**:
   ```
   cli.prompt() expects exactly 1 argument (prompt message), but received 2. 
   Usage: cli.prompt("Enter your name: ")
   ```

2. **Type Errors**:
   ```
   math.abs() expects a number argument, but received STRING. 
   Usage: math.abs(-5)
   ```

3. **Value Validation Errors**:
   ```
   time.sleep() cannot sleep for negative duration (-5 seconds). 
   Please provide a positive number.
   ```

4. **Network Errors**:
   ```
   net.get() failed to execute HTTP request to 'https://invalid.url': 
   connection timeout. Please check your internet connection and ensure the server is accessible.
   ```

## Benefits of Enhanced Errors

### For Developers:
- **Faster Debugging**: Immediate understanding of what went wrong
- **Learning Aid**: Usage examples help learn correct syntax
- **Context Awareness**: Understand why something failed, not just that it failed
- **Type Safety**: Clear indication of expected vs received types

### For the Language:
- **Better Developer Experience**: Reduces frustration and development time
- **Self-Documenting**: Error messages serve as inline documentation
- **Consistency**: All modules follow the same error message patterns
- **Professionalism**: Makes Vint feel more polished and production-ready

## Testing Enhanced Errors

Use the provided `enhanced_error_test.vint` file to see all the improved error messages in action:

```bash
vint enhanced_error_test.vint
```

## Future Enhancements

The error enhancement framework is designed to be extensible. Future improvements could include:

1. **Error Codes**: Numeric codes for programmatic error handling
2. **Suggestions**: Auto-suggestions for common mistakes
3. **Stack Traces**: Better debugging with call stack information
4. **Localization**: Multi-language error messages
5. **Error Recovery**: Suggested fixes or alternatives

## Contributing

When adding new functions or modules, follow these error message guidelines:

1. Always include the module and function name
2. Specify exactly what was expected vs what was received
3. Provide a concrete usage example
4. Include context about why the error occurred when possible
5. Use consistent formatting and terminology

This enhanced error system makes Vint more user-friendly and professional, helping developers write better code faster.
