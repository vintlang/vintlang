# While Loops in Vint

While loops in Vint are used to execute a block of code repeatedly, as long as a given condition is true. This page covers the basics of while loops, including how to use the `break` and `continue` keywords within them.

## Basic Syntax

A while loop is executed when a specified condition is true. You initialize a while loop with the `while` keyword followed by the condition in parentheses `()`. The consequence of the loop should be enclosed in curly braces `{}`.

```vint
let i = 1

while (i <= 5) {
    print(i)
    i++
}
```

### Output:
```vint
1
2
3
4
5
```

## Break and Continue

### Break

Use the `break` keyword to terminate a loop:

```vint
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
```vint
1
2
broken
```

### Continue

Use the `continue` keyword to skip a specific iteration:

```vint
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
```vint
1
2
skipped
4
5
```

By understanding while loops in Vint, you can create code that repeats a specific action or checks for certain conditions, offering more flexibility and control over your code execution.