# Info

The `info` keyword allows you to print informational messages at runtime.

## Syntax

```js
info "Your informational message here"
```

When the Vint interpreter encounters an `info` statement, it prints a cyan-colored informational message to the console and continues execution. This is useful for providing helpful context or status updates to users or developers.

### Example

```js
info "Starting the backup process."
println("Backup in progress...")
```

Running this script will output:

```
[INFO]: Starting the backup process.
Backup in progress...
``` 