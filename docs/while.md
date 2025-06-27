# While Loops in Vint

While loops in Vint are used to execute a block of code repeatedly, as long as a given condition is true. This page covers the basics of while loops, including how to use the `break` and `continue` keywords within them.

## Basic Syntax

A while loop is executed when a specified condition is true. You initialize a while loop with the `while` keyword followed by the condition in parentheses `()`. The consequence of the loop should be enclosed in curly braces `{}`.

```js
let i = 1

while (i <= 5) {
    print(i)
    i++
}
```

### Output:
```js
1
2
3
4
5
```

## Break and Continue

### Break

Use the `break` keyword to terminate a loop:

```js
let i = 1

while (i < 5) {
    if (i == 3) {
        print("broken")
        break
    }
    print(i)
    i++
}
```

### Output:
```js
1
2
broken
```

### Continue

Use the `continue` keyword to skip a specific iteration:

```js
let i = 0

while (i < 5) {
    i++
    if (i == 3) {
        print("skipped")
        continue
    }
    print(i)
}
```

### Output:
```js
1
2
skipped
4
5
```

By understanding while loops in Vint, you can create code that repeats a specific action or checks for certain conditions, offering more flexibility and control over your code execution.

## Repeat Loops

The `repeat` keyword allows you to execute a block of code a specific number of times. The default loop variable `i` is available inside the block, representing the current iteration (starting from 0).

### Syntax

```vint
repeat 5 {
    println("Iteration:", i)
}
```

This will print:

```
Iteration: 0
Iteration: 1
Iteration: 2
Iteration: 3
Iteration: 4
```

You can also use an expression for the count:

```vint
let n = 3
repeat n {
    println(i)
}
```

The variable `i` is always available in the scope of the repeat block.