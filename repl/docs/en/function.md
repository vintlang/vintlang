# Functions in Vint

Functions in **Vint** allow you to encapsulate code and execute it when needed. Here's a simple guide to understanding how functions work in **Vint**.

## Immediately Invoked Function

You can define and immediately execute a function:

```vint
let go = func() {
    print("this is a function")
}()
```

This function `go` is defined and executed immediately upon declaration.

## Declared but Not Immediately Invoked Function

Functions can also be declared without being executed immediately:

```vint
let vint = func() {
    print("This is also a function\nBut not invoked immediately after being declared")
}

vint()  // Executes the function
```

The function `vint` is called later using `vint()`.

## Passing Functions as Arguments

Functions in **Vint** can be passed as arguments to other functions:

```vint
let w = func() {
    print("w function")
}

func(w) {
    w()  // Executes the function passed as an argument
    print("func")
}(w)  // Passes `w` as an argument and immediately invokes the outer function
```

In this example, the function `w` is passed to another function and executed within it.

By understanding these basic concepts, you can start creating reusable and flexible code using functions in **Vint**.