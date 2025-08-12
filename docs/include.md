# Include

The `include` keyword is a language construct in Vint that allows you to include and evaluate code from another file into the current file. This is useful for organizing code into reusable modules, separating concerns, and managing larger projects more effectively. When a file is included, its code is executed in the same scope as the `include` statement, meaning any variables, functions, or other constructs defined in the included file become available in the including file.

## Syntax

```js
include "path/to/your/file.vint"
```

The path to the file can be relative or absolute. The file extension is not mandatory but is recommended for clarity.

## Example

Let's say you have a file named `greetings.vint` with the following content:

**greetings.vint**
```js
let greeting = "Hello, Vint!"

func sayHello() {
    println(greeting)
}
```

You can include this file in another file, for instance, `main.vint`, and use the `greeting` variable and the `sayHello` function:

**main.vint**
```js
include "greetings.vint"

sayHello() // Output: Hello, Vint!

let customMessage = greeting + " How are you?"
println(customMessage) // Output: Hello, Vint! How are you?
```

In this example, the `include` statement at the beginning of `main.vint` makes the `greeting` variable and the `sayHello` function from `greetings.vint` available for use. This helps in keeping the code modular and easy to manage. 