# INTEGERS (NAMBA) AND FLOATS (DESIMALI)

Integers and floats are the basic numeric data types in vint, used for representing whole numbers and decimal numbers, respectively. This page covers the syntax and usage of integers and floats in vint, including precedence, unary increments, shorthand assignments, and negative numbers.

## PRECEDENCE

Integers and floats behave as expected in mathematical operations, following the BODMAS rule:
```go
2 + 3 * 5 // 17

let a = 2.5
let b = 3/5

a + b // 2.8
```

## UNARY INCREMENTS

You can perform unary increments (++ and --) on both floats and integers. These will add or subtract 1 from the current value. Note that the float or int have to be assigned to a variable for this operation to work. Here's an example:

```go
let i = 2.4

i++ // 3.4
```

## SHORTHAND ASSIGNMENT

vint supports shorthand assignments with +=, -=, /=, *=, and %=:
You
```go
let i = 2

i *= 3 // 6
i /= 2 // 3
i += 100 // 103
i -= 10 // 93
i %= 90 // 3
```

## NEGATIVE NUMBERS

Negative numbers also behave as expected:

```go
let i = -10

wakati (i < 0) {
    print(i)
    i++
}

```
Output:
```s
-10
-9
-8
-7
-6
-5
-4
-3
-2
-1
0
1
2
3
4
5
6
7
8
9 
```
