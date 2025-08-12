# Success

The `success` keyword allows you to print success messages at runtime.

## Syntax

```js
success "Your success message here"
```

When the Vint interpreter encounters a `success` statement, it prints a green-colored success message to the console and continues execution. This is useful for indicating when an operation has completed successfully.

### Example

```js
success "Backup completed successfully!"
println("All done.")
```

Running this script will output:

```
[SUCCESS]: Backup completed successfully!
All done.
``` 