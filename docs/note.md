# Note

The `note` keyword allows you to print general notes at runtime.

## Syntax

```vint
note "Your note message here"
```

When the Vint interpreter encounters a `note` statement, it prints a blue-colored note message to the console and continues execution. This is useful for providing additional context or reminders in your scripts.

### Example

```vint
note "This script was last updated on 2024-06-01."
println("Script running...")
```

Running this script will output:

```
[NOTE]: This script was last updated on 2024-06-01.
Script running...
``` 