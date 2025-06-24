# Warn

The `warn` keyword allows you to emit non-fatal warnings at runtime.

## Syntax

`warn "Your warning message here"`

When the interpreter encounters a `warn` statement, it will print a formatted warning to the console and then continue execution. This is useful for alerting developers to potential issues that don't need to stop the program, such as using a deprecated feature or missing a configuration file.

### Example

```vint
warn "Configuration file not found, using default settings."
println("Program is running with default configuration.")
```
Running this will output:
```

[WARN]: Configuration file not found, using default settings.

Program is running with default configuration.
``` 