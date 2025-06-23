# INTEGERS  AND FLOATS 

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

while (i < 0) {
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

## Integer Methods

Integers in vint have several built-in methods:

### abs()

Returns the absolute value of the integer:

```s
let i = -42
print(i.abs())  // 42
```

### is_even()

Returns true if the integer is even, false otherwise:

```s
let i = 4
print(i.is_even())  // true
print((5).is_even())  // false
```

### is_odd()

Returns true if the integer is odd, false otherwise:

```s
let i = 7
print(i.is_odd())  // true
print((8).is_odd())  // false
```

### to_string()

Converts the integer to a string:

```s
let i = 123
print(i.to_string())  // "123"
```

### sign()

Returns 1 if the integer is positive, -1 if negative, or 0 if zero:

```s
print((10).sign())   // 1
print((-5).sign())   // -1
print((0).sign())    // 0
```
