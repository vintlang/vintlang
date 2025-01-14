# Conditional Statements in Vint

Conditional statements in **Vint** allow you to perform different actions based on specific conditions. The `if/else` structure is fundamental for controlling the flow of your code. Here's a simple guide to using conditional statements in **Vint**.

## If Statement (`if`)

The `if` statement checks a condition inside parentheses `()`. If the condition evaluates to true, the code block inside curly braces `{}` will execute:

```vint
if (2 > 1) {
    print(true)  // Output: true
}
```

In this example, the condition `2 > 1` is true, so `print(true)` is executed, and the output is `true`.

## Else If and Else Blocks (`else if` and `else`)

You can use `else if` to test additional conditions after an `if` statement. The `else` block specifies code to execute if none of the previous conditions are met:

```vint
let a = 10

if (a > 100) {
    print("a is greater than 100")
} else if (a < 10) {
    print("a is less than 10")
} else {
    print("The value of a is", a)
}

// Output: The value of a is 10
```

### Explanation:
1. The condition `a > 100` is false.
2. The next condition `a < 10` is also false.
3. Therefore, the `else` block is executed, and the output is `The value of a is 10`.

## Summary

- **`if`**: Executes code if the condition is true.
- **`else if`**: Tests another condition if the previous `if` condition is false.
- **`else`**: Executes code if none of the above conditions are true.

By using `if`, `else if`, and `else`, you can make decisions and control the flow of your **Vint** programs based on dynamic conditions.