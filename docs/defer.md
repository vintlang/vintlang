# Defer

The `defer` keyword provides a convenient way to schedule a function call to be executed just before the surrounding function returns. This is particularly useful for cleanup tasks, such as closing files or releasing resources, ensuring that they are always executed, regardless of how the function exits.

## Syntax

The `defer` keyword is followed by a function call:

```js
defer functionCall()
```

## Example

Here’s a simple example that demonstrates how `defer` works. The deferred `println` call is executed after the function body has completed but before the function returns.

```js
let my_function = func() {
    defer println("This will be printed last");
    println("This will be printed first");
};

my_function();
// Output:
// This will be printed first
// This will be printed last
```

## Multiple Defer Statements

If a function has multiple `defer` statements, they are pushed onto a stack. When the function returns, the deferred calls are executed in last-in, first-out (LIFO) order.

```js
let another_function = func() {
    defer println("deferred: 1");
    defer println("deferred: 2");
    println("function body");
};

another_function();
// Output:
// function body
// deferred: 2
// deferred: 1
```

This LIFO order is intuitive for managing resources. For example, if you acquire a resource and then lock it, you would want to unlock it first and then release it, which `defer` handles naturally.
