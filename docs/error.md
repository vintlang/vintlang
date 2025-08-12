# Error

The `error` keyword allows you to raise a fatal runtime error that stops the execution of the script.

## Syntax

`error "Your error message here"`

When the interpreter encounters an `error` statement, it will print a formatted error message to the console and halt execution immediately. This is useful for handling critical problems where the program cannot safely continue.

### Example

```js
let file_path = "data.json"
if !fs.exists(file_path) {
    error "Critical file 'data.json' not found."
}
println("This will not be printed if the file is missing.")
```
If `data.json` does not exist, running this script will output:
```
Error: Critical file 'data.json' not found.
```
And the script will stop. 