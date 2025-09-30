# Builtin Functions Organization

This document describes the new organization of builtin functions in VintLang and provides guidance for developers on where to add new functions.

## Overview

The builtin functions have been restructured from a single large file (`builtins.go`) into a modular system with logical categorization. This improves maintainability, reduces merge conflicts, and makes the codebase easier to navigate.

## New Structure

```
evaluator/
  builtins.go           # Main interface file
  builtins/
    registry.go         # Central registration system
    helpers.go          # Common helper functions
    core.go            # Essential builtin functions
    io.go              # I/O related functions
    type_conversion.go # Type conversion functions
    logic.go           # Logical operations
    arrays.go          # Array manipulation functions
    dict.go            # Dictionary operations
    system.go          # System functions (exit, sleep, args)
    channels.go        # Channel operations
    imports.go         # Import functions
```

## Function Categories

### Core Functions (builtins/core.go)
Essential, frequently used functions that should remain as builtins:
- `print`, `println`, `printErr`, `printlnErr` - Output functions
- `input` - User input
- `len` - Get length of collections/strings
- `type` - Get type information

### I/O Functions (builtins/io.go)
File and data I/O operations:
- `open` - Read file contents
- `write` - Format data for output

### Type Conversion (builtins/type_conversion.go)
Functions for converting between types:
- `convert` - Generic type conversion
- `string`, `int` - Specific type converters
- `parseInt`, `parseFloat` - Parse from strings

### Logic Functions (builtins/logic.go)
Logical operations:
- `and`, `or`, `not` - Boolean logic
- `xor`, `nand`, `nor` - Extended logical operations
- `eq` - Equality comparison

### Array Functions (builtins/arrays.go)
Array manipulation:
- `range` - Generate number sequences
- `append`, `pop` - Array modification
- `indexOf`, `unique` - Array utilities

### Dictionary Functions (builtins/dict.go)
Dictionary operations:
- `keys`, `values` - Get dictionary contents
- `has_key` - Check key existence

### System Functions (builtins/system.go)
System-level operations:
- `exit` - Program termination
- `sleep` - Pause execution
- `args` - Command line arguments

### Channel Functions (builtins/channels.go)
Concurrency operations:
- `send`, `receive`, `close` - Channel operations

### Import Functions (builtins/imports.go)
Module importing:
- `import` - Dynamic module import

## Functions Moved to Modules

### String Functions → `string` module
Functions that were moved from builtins to the `string` module:
- `startsWith` → `string.startsWith`
- `endsWith` → `string.endsWith`
- `chr` → `string.chr`
- `ord` → `string.ord`

### Type Checking → `reflect` module
Type checking functions are available in the `reflect` module:
- `isInt` → `reflect.isInt`
- `isFloat` → `reflect.isFloat`
- `isString` → `reflect.isString`
- `isBool` → `reflect.isBool`
- `isArray` → `reflect.isArray` (already existed)
- `isDict` → `reflect.isObject` (already existed)
- `isNull` → `reflect.isNil` (already existed)

## Guidelines for Adding New Functions

### When to Add as Builtin
Functions should be builtin only if they are:
1. **Essential** - Used frequently across many programs
2. **Core language features** - Basic operations like `len`, `type`, `print`
3. **Cannot be implemented in modules** - Require evaluator access

### When to Add to Modules
Functions should go in modules if they are:
1. **Specialized** - Domain-specific functionality
2. **Complex** - Advanced features that aren't basic language operations
3. **Logical groupings** - Related to existing module functionality

### Adding New Builtin Functions

1. **Choose the appropriate category file** based on function purpose
2. **Add to the relevant `registerXXXBuiltins()` function** in that file
3. **Follow the existing pattern**:
   ```go
   RegisterBuiltin("functionName", &object.Builtin{
       Fn: func(args ...object.Object) object.Object {
           // Implementation
       },
   })
   ```
4. **Use helper functions** from `helpers.go` for common operations
5. **Add appropriate error checking** and validation

### Adding New Module Functions

1. **Choose the appropriate module** in the `module/` directory
2. **Add to the module's function map** in its `init()` function
3. **Follow the module's error handling pattern** using `ErrorMessage()`
4. **Update module documentation**

## Registry System

The new registry system (`builtins/registry.go`) provides:
- **Central registration** - All builtins register themselves via `RegisterBuiltin()`
- **Automatic discovery** - Functions are automatically available once registered
- **Backward compatibility** - Existing evaluator code continues to work

## Benefits of New Structure

1. **Better Organization** - Related functions are grouped together
2. **Easier Maintenance** - Smaller files are easier to work with
3. **Reduced Conflicts** - Multiple developers can work on different categories
4. **Clear Guidelines** - Developers know where to add new functions
5. **Logical Separation** - Core builtins vs. specialized module functions

## Migration Notes

- All existing builtin functions continue to work exactly the same way
- No changes required to existing VintLang programs
- The evaluator interface remains unchanged
- Imports and initialization happen automatically

## Future Considerations

Consider moving these builtin functions to modules in the future:
- More type checking functions → `reflect` module
- Math operations → `math` module (if any are currently builtins)
- Advanced array operations → dedicated `array` module
- File operations → `os` module (where appropriate)

The goal is to keep only truly essential functions as builtins while providing rich functionality through well-organized modules.