# Working with Booleans in vint

Boolean objects in vint are truthy, meaning that any value is true, except tupu and false. They are used to evaluate expressions that return true or false values.

## Evaluating Boolean Expressions

### Evaluating Simple Expressions

In vint, you can evaluate simple expressions that return a boolean value:

```s
print(1 > 2) // Output: `false`

print(1 + 3 < 10) // Output: `true`
```

### Evaluating Complex Expressions

In vint, you can use boolean operators to evaluate complex expressions:

```s
a = 5
b = 10
c = 15

result = (a < b) && (b < c)

if (result) {
    print("Both conditions are true")
} else {
    print("At least one condition is false")
}
// Output: "Both conditions are true"
```

Here, we create three variables a, b, and c. We then evaluate the expression (a < b) && (b < c). Since both conditions are true, the output will be "Both conditions are true".

## Boolean Operators

vint has several boolean operators that you can use to evaluate expressions:

### The && Operator

The && operator evaluates to true only if both operands are true. Here's an example:

```s
print(true && true) // Output: `true`

print(true && false) // Output: `false`
```

### The and() function

```s
print(and(true,true)) // Output: `true`

print(and(true,false)) // Output: `false`
```

### The || Operator

The || operator evaluates to true if at least one of the operands is true. Here's an example:

```s
print(true || false) // Output: `true`

print(false || false) // Output: `false`
```
### The or() Function
```s
print(or(true,false)) // Output: `true`

print(or(false,false)) // Output: `false`
```

### The ! Operator

The ! operator negates the value of the operand. Here's an example:

```s
print(!true) // Output: `false`

print(!false) // Output: `true`
```

### The not() function

```s
print(not(true)) // Output: `false`

print(not(false)) // Output: `true`
```

## Working with Boolean Values in Loops

In vint, you can use boolean expressions in loops to control their behavior. Here's an example:

```s
num = [1, 2, 3, 4, 5]

for v in num {
    if (v % 2 == 0) {
        print(v, "is even")
    } else {
        print(v, "is odd")
    }
}
// Output:
// 1 is odd
// 2 is even
// 3 is odd
// 4 is even
// 5 is odd
```

