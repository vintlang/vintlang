# TODOs

The `todo` keyword allows you to leave compiler-visible TODOs that warn at runtime.

## Syntax

```vint
todo "Your todo message here"
```

When the Vint interpreter encounters a `todo` statement, it will print a warning to the console with your message, and then continue execution.

### Example

```vint
todo "Implement user authentication"

let x = 10
println(x)
```

Running this script will output:

```
TODO: "Implement user authentication"
10
``` 