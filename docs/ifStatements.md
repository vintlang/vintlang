# Conditional Statements in Vint

Conditional statements in **Vint** allow you to perform different actions based on specific conditions. The `if/else` structure is fundamental for controlling the flow of your code. Here's a simple guide to using conditional statements in **Vint**.

## If Statement (`if`)

The `if` statement checks a condition inside parentheses `()`. If the condition evaluates to true, the code block inside curly braces `{}` will execute:

```js
if (2 > 1) {
    print(true)  // Output: true
}
```

In this example, the condition `2 > 1` is true, so `print(true)` is executed, and the output is `true`.

## Else If and Else Blocks (`else if` and `else`)

You can use `else if` to test additional conditions after an `if` statement. The `else` block specifies code to execute if none of the previous conditions are met:

```js
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

# If Statements and If Expressions in Vint

Vint supports both classic if statements and the new if expressions, allowing you to use conditional logic in both statement and expression positions.

---

## Classic If Statement

The classic if statement executes a block of code if a condition is true. You can optionally provide an `else` block.

**Syntax:**
```vint
if (condition) {
    // code to run if condition is true
} else {
    // code to run if condition is false
}
```

**Example:**
```vint
let x = 0
if (true) {
    x = 42
}
print("Classic if statement result: ", x)
```

---

## If as an Expression (New Feature)

You can now use `if` as an expression, which returns a value. This allows you to assign the result of a conditional directly to a variable, or use it in any expression context.

**Syntax:**
```vint
let result = if (condition) { valueIfTrue } else { valueIfFalse }
```

- The `if` expression evaluates to the value of the first block if the condition is true, or the value of the `else` block if provided.
- If the condition is false and there is no `else`, the result is `null`.

**Examples:**
```vint
let status = ""
status = if (x > 0) { "Online" } else { "Offline" }
print("If as an expression result: ", status)

let y = if (false) { 123 }
print("If as an expression with no else: ", y) // prints: null
```

---

## Notes
- Parentheses around the condition are required: `if (condition) { ... }`.
- Both the classic statement and the new expression form are fully supported and can be mixed in your code.
- Use `//` for single-line comments and `/* ... */` for multi-line comments in Vint.

---

## See Also
- [Switch Statements](switch.md)
- [Operators](operators.md)