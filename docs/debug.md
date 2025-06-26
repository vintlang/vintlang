# Debug

The `debug` keyword allows you to print debug messages at runtime for troubleshooting and development.

## Syntax

```vint
debug "Your debug message here"
```

When the Vint interpreter encounters a `debug` statement, it prints a magenta-colored debug message to the console and continues execution. This is useful for inspecting variable values or program flow during development.

### Example

```vint
let value = 42
debug "Current value is: " + value
println("Done.")
```

Running this script will output:

```
[DEBUG]: Current value is: 42
Done.
``` 