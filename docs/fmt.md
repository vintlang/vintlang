# fmt Module

The `fmt` module provides comprehensive text formatting capabilities for VintLang, including string formatting, padding, alignment, number formatting, and utility functions.

## Overview

- **Purpose**: Advanced text formatting and manipulation
- **Module Name**: `fmt`
- **Usage**: `import fmt` or `let fmt = import("fmt")`

## Functions

### String Formatting Functions

#### `sprintf(format, ...args) -> string`

Formats a string using Go's fmt.Sprintf format specifiers.

```vint
import fmt
let result = fmt.sprintf("Hello, %s! You have %d messages.", "Alice", 5)
print(result)  // "Hello, Alice! You have 5 messages."
```

#### `printf(format, ...args) -> nil`

Prints formatted text to stdout.

```vint
import fmt
fmt.printf("Temperature: %.2f°C\n", 23.456)  // "Temperature: 23.46°C"
```

#### `fprintf(file, format, ...args) -> nil`

Prints formatted text to a file or writer (currently outputs to stdout).

```vint
import fmt
fmt.fprintf(null, "Debug: %s = %d\n", "count", 42)
```

#### `errorf(format, ...args) -> error`

Creates a formatted error object.

```vint
import fmt
let err = fmt.errorf("failed to process %s: %s", "file.txt", "not found")
```

### Padding and Alignment Functions

#### `padLeft(str, width, [padChar]) -> string`

Pads a string to the left with spaces or specified character.

```vint
import fmt
let padded = fmt.padLeft("hello", 10)        // "     hello"
let custom = fmt.padLeft("test", 8, "0")     // "0000test"
```

#### `padRight(str, width, [padChar]) -> string`

Pads a string to the right with spaces or specified character.

```vint
import fmt
let padded = fmt.padRight("hello", 10)       // "hello     "
let custom = fmt.padRight("test", 8, "-")    // "test----"
```

#### `padCenter(str, width, [padChar]) -> string`

Centers a string with padding.

```vint
import fmt
let centered = fmt.padCenter("hello", 11)    // "   hello   "
let custom = fmt.padCenter("test", 10, "*")  // "***test***"
```

### Number Formatting Functions

#### `formatInt(number, [base], [width]) -> string`

Formats an integer with specified base and width.

```vint
import fmt
let decimal = fmt.formatInt(42, 10, 5)       // "   42"
let binary = fmt.formatInt(15, 2)            // "1111"
let hex = fmt.formatInt(255, 16)             // "ff"
```

#### `formatFloat(number, [precision]) -> string`

Formats a float with specified precision.

```vint
import fmt
let rounded = fmt.formatFloat(3.14159, 2)    // "3.14"
let precise = fmt.formatFloat(1.2345, 4)     // "1.2345"
```

#### `formatHex(number, [uppercase]) -> string`

Formats an integer as hexadecimal.

```vint
import fmt
let lower = fmt.formatHex(255)               // "ff"
let upper = fmt.formatHex(255, true)         // "FF"
```

#### `formatOct(number) -> string`

Formats an integer as octal.

```vint
import fmt
let octal = fmt.formatOct(64)                // "100"
```

#### `formatBin(number) -> string`

Formats an integer as binary.

```vint
import fmt
let binary = fmt.formatBin(15)               // "1111"
```

### Width and Precision Functions

#### `width(str, width) -> string`

Formats a string to a specific width (truncates if too long).

```vint
import fmt
let fixed = fmt.width("hello world", 5)      // "hello"
let padded = fmt.width("hi", 8)              // "hi      "
```

#### `precision(number, precision) -> string`

Formats a float with specific precision.

```vint
import fmt
let precise = fmt.precision(3.14159, 2)      // "3.14"
```

### Utility Functions

#### `repeat(str, count) -> string`

Repeats a string n times.

```vint
import fmt
let repeated = fmt.repeat("ha", 3)           // "hahaha"
let line = fmt.repeat("-", 20)               // "--------------------"
```

#### `truncate(str, maxLength) -> string`

Limits a string to a maximum length.

```vint
import fmt
let short = fmt.truncate("hello world", 5)   // "hello"
let unchanged = fmt.truncate("hi", 10)       // "hi"
```

## Format Specifiers

When using `sprintf`, `printf`, and `fprintf`, you can use Go's format specifiers:

- `%s` - string
- `%d` - decimal integer
- `%f` - floating point
- `%.2f` - floating point with 2 decimal places
- `%x` - hexadecimal (lowercase)
- `%X` - hexadecimal (uppercase)
- `%o` - octal
- `%b` - binary
- `%t` - boolean
- `%v` - default format
- `%%` - literal %

## Complete Example

```vint
import fmt

// Basic formatting
let name = "Alice"
let age = 30
let height = 5.75

let intro = fmt.sprintf("Hi, I'm %s, %d years old, %.1f feet tall", name, age, height)
print(intro)

// Number formatting
let num = 255
print("Decimal:", fmt.formatInt(num, 10))     // "255"
print("Binary:", fmt.formatBin(num))          // "11111111"
print("Hex:", fmt.formatHex(num, true))       // "FF"
print("Octal:", fmt.formatOct(num))           // "377"

// Padding and alignment
let title = "REPORT"
let border = fmt.repeat("=", 20)

print(border)
print(fmt.padCenter(title, 20))
print(border)

// Table formatting
let items = [
    {"name": "Apple", "price": 1.25, "qty": 10},
    {"name": "Banana", "price": 0.75, "qty": 25},
    {"name": "Orange", "price": 1.50, "qty": 15}
]

fmt.printf("%-10s %8s %5s\n", "Item", "Price", "Qty")
fmt.printf("%s\n", fmt.repeat("-", 25))

for item in items {
    fmt.printf("%-10s $%7.2f %5d\n",
        item["name"],
        item["price"],
        item["qty"])
}
```

## Error Handling

All functions validate their input types and return descriptive error messages for invalid usage:

```vint
import fmt
let err = fmt.padLeft(123, "invalid")  // Returns error object
```

## Notes

- The `fprintf` function currently writes to stdout but can be extended to write to actual file objects
- Padding functions default to space character if no pad character is specified
- Number formatting functions handle negative numbers appropriately
- Width and truncation functions preserve string integrity while enforcing limits
