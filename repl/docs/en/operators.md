# Operators in Vint

Operators are a core feature of any programming language, enabling you to perform various operations on variables and values. This page details the syntax and usage of operators in Vint, including assignment, arithmetic, comparison, membership, and logical operators.

---

## Assignment Operators

Assignment operators are used to assign values to variables. The following are supported in Vint:

- `i = v`: Assigns the value of `v` to `i`.
- `i += v`: Equivalent to `i = i + v`.
- `i -= v`: Equivalent to `i = i - v`.
- `i *= v`: Equivalent to `i = i * v`.
- `i /= v`: Equivalent to `i = i / v`.

For strings, arrays, and dictionaries, the `+=` operator is also valid. For example:

```vint
list1 += list2 // Equivalent to list1 = list1 + list2
```

---

## Arithmetic Operators

Vint supports the following arithmetic operations:

| Operator | Description                          | Example          |
|----------|--------------------------------------|------------------|
| `+`      | Addition                             | `2 + 3 = 5`      |
| `-`      | Subtraction                          | `5 - 2 = 3`      |
| `*`      | Multiplication                       | `3 * 4 = 12`     |
| `/`      | Division                             | `10 / 2 = 5`     |
| `%`      | Modulo (remainder of a division)     | `7 % 3 = 1`      |
| `**`     | Exponential power                   | `2 ** 3 = 8`     |

---

## Comparison Operators

Comparison operators evaluate relationships between two values. These return `true` or `false`:

| Operator | Description                     | Example            |
|----------|---------------------------------|--------------------|
| `==`     | Equal to                        | `5 == 5 // true`   |
| `!=`     | Not equal to                    | `5 != 3 // true`   |
| `>`      | Greater than                    | `5 > 3 // true`    |
| `>=`     | Greater than or equal to        | `5 >= 5 // true`   |
| `<`      | Less than                       | `3 < 5 // true`    |
| `<=`     | Less than or equal to           | `3 <= 3 // true`   |

---

## Membership Operator

The membership operator `in` checks if an item exists within a collection:

```vint
names = ['juma', 'asha', 'haruna']

"haruna" in names // true
"halima" in names // false
```

---

## Logical Operators

Logical operators allow you to combine or invert conditions:

| Operator | Description              | Example                   |
|----------|--------------------------|---------------------------|
| `&&`     | Logical AND              | `true && false // false` |
| `||`     | Logical OR               | `true || false // true`  |
| `!`      | Logical NOT (negation)   | `!true // false`         |

---

## Precedence of Operators

When multiple operators are used in an expression, operator precedence determines the order of execution. Below is the precedence order, from highest to lowest:

1. `()` : Parentheses
2. `!`  : Logical NOT
3. `%`  : Modulo
4. `**` : Exponential power
5. `/`, `*` : Division and multiplication
6. `+`, `+=`, `-`, `-=` : Addition and subtraction
7. `>`, `>=`, `<`, `<=` : Comparison operators
8. `==`, `!=` : Equality and inequality
9. `=` : Assignment
10. `in` : Membership operator
11. `&&`, `||` : Logical AND and OR

---
