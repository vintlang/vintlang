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

### pow()

Raises the integer to the power of another number:

```s
let base = 2
print(base.pow(3))   // 8
print((5).pow(2))    // 25
```

### sqrt()

Returns the square root of the integer as a float:

```s
let num = 16
print(num.sqrt())    // 4.0
print((25).sqrt())   // 5.0
```

### gcd()

Returns the greatest common divisor of two integers:

```s
let a = 24
let b = 18
print(a.gcd(b))      // 6
print((48).gcd(18))  // 6
```

### lcm()

Returns the least common multiple of two integers:

```s
let a = 12
let b = 8
print(a.lcm(b))      // 24
print((15).lcm(20))  // 60
```

### factorial()

Returns the factorial of the integer:

```s
let n = 5
print(n.factorial()) // 120
print((4).factorial()) // 24
print((0).factorial()) // 1
```

### toBinary()

Converts the integer to binary representation:

```s
let num = 255
print(num.toBinary())  // "11111111"
print((5).toBinary())  // "101"
```

### toHex()

Converts the integer to hexadecimal representation:

```s
let num = 255
print(num.toHex())     // "ff"
print((16).toHex())    // "10"
```

### toOctal()

Converts the integer to octal representation:

```s
let num = 64
print(num.toOctal())   // "100"
print((8).toOctal())   // "10"
```

### isPrime()

Checks if the integer is a prime number:

```s
print((17).isPrime())  // true
print((4).isPrime())   // false
print((2).isPrime())   // true
print((1).isPrime())   // false
```

### nthRoot()

Calculates the nth root of the integer:

```s
let num = 8
print(num.nthRoot(3))  // 2.0 (cube root)
print((16).nthRoot(2)) // 4.0 (square root)
```

### mod()

Calculates the modulo (remainder) with another integer:

```s
let num = 10
print(num.mod(3))      // 1
print((15).mod(4))     // 3
```

### clamp()

Restricts the integer to be within specified bounds:

```s
let num = 15
print(num.clamp(1, 10))  // 10 (clamped to max)
print((-5).clamp(1, 10)) // 1 (clamped to min)
print((5).clamp(1, 10))  // 5 (within bounds)
```

### inRange()

Checks if the integer is within the specified range (inclusive):

```s
let num = 5
print(num.inRange(1, 10))  // true
print((15).inRange(1, 10)) // false
print((0).inRange(1, 10))  // false
```

### digits()

Returns an array of individual digits:

```s
let num = 123
print(num.digits())    // [1, 2, 3]
print((456).digits())  // [4, 5, 6]
```

## Float Methods

Floats in vint have powerful built-in methods for mathematical operations and utility functions:

### abs()

Returns the absolute value of the float:

```s
let f = -3.14
print(f.abs())       // 3.14
print((-2.5).abs())  // 2.5
```

### ceil()

Returns the smallest integer greater than or equal to the float:

```s
let price = 29.95
print(price.ceil())  // 30
print((4.1).ceil())  // 5
print((-2.1).ceil()) // -2
```

### floor()

Returns the largest integer less than or equal to the float:

```s
let price = 29.95
print(price.floor()) // 29
print((4.9).floor()) // 4
print((-2.1).floor()) // -3
```

### round()

Rounds the float to a specified number of decimal places:

```s
let pi = 3.14159
print(pi.round(2))   // 3.14
print(pi.round(0))   // 3
print((2.7).round()) // 3
```

### sqrt()

Returns the square root of the float:

```s
let num = 9.0
print(num.sqrt())    // 3.0
print((16.0).sqrt()) // 4.0
```

### pow()

Raises the float to the power of another number:

```s
let base = 2.5
print(base.pow(2))   // 6.25
print((3.0).pow(3))  // 27.0
```

### is_nan()

Checks if the float is NaN (Not a Number):

```s
let valid = 3.14
let invalid = 0.0 / 0.0
print(valid.is_nan())   // false
print(invalid.is_nan()) // true
```

### is_infinite()

Checks if the float is infinite:

```s
let normal = 3.14
let inf = 1.0 / 0.0
print(normal.is_infinite()) // false
print(inf.is_infinite())    // true
```

### to_string()

Converts the float to a string with optional precision:

```s
let price = 29.95
print(price.to_string())   // "29.95"
print(price.to_string(1))  // "30.0"
print((3.14159).to_string(2)) // "3.14"
```

### clamp()

Clamps the float between minimum and maximum values:

```s
let value = 75.5
print(value.clamp(0.0, 50.0))  // 50.0
print((25.3).clamp(30.0, 100.0)) // 30.0
print((45.7).clamp(10.0, 80.0))  // 45.7
```

### toPrecision()

Formats the float to specified precision:

```s
let num = 123.456789
print(num.toPrecision(4))    // "123.5"
print((0.123456).toPrecision(3)) // "0.123"
```

### toFixed()

Formats the float to fixed decimal places:

```s
let num = 123.456
print(num.toFixed(2))        // "123.46"
print((5.0).toFixed(3))      // "5.000"
```

### sign()

Returns the sign of the float:

```s
print((5.5).sign())          // 1.0
print((-3.2).sign())         // -1.0
print((0.0).sign())          // 0.0
```

### truncate()

Removes the fractional part:

```s
print((5.9).truncate())      // 5.0
print((-3.7).truncate())     // -3.0
```

### mod()

Calculates the floating-point remainder:

```s
let num = 5.5
print(num.mod(2.0))          // 1.5
print((10.7).mod(3.0))       // 1.7
```

### degrees()

Converts radians to degrees:

```s
import math
let pi = math.PI
print(pi.degrees())          // 180.0
print((pi / 2).degrees())    // 90.0
```

### radians()

Converts degrees to radians:

```s
print((180.0).radians())     // 3.141592653589793
print((90.0).radians())      // 1.5707963267948966
```

### sin()

Calculates the sine:

```s
print((0.0).sin())           // 0.0
print((math.PI / 2).sin())   // 1.0
```

### cos()

Calculates the cosine:

```s
print((0.0).cos())           // 1.0
print(math.PI.cos())         // -1.0
```

### tan()

Calculates the tangent:

```s
print((0.0).tan())           // 0.0
print((math.PI / 4).tan())   // 1.0
```

### log()

Calculates the natural logarithm:

```s
import math
print(math.E.log())          // 1.0
print((10.0).log())          // 2.302585092994046
```

### exp()

Calculates e raised to the power of the float:

```s
print((0.0).exp())           // 1.0
print((1.0).exp())           // 2.718281828459045
```

## Practical Examples

Here are some practical examples using integer and float methods:

```s
// Calculate compound interest
let principal = 1000.0
let rate = 0.05
let time = 3
let amount = principal * (1.0 + rate).pow(time)
print("Amount after", time, "years:", amount.round(2))

// Check if numbers are perfect squares
numbers = [16, 25, 30, 36]
for num in numbers {
    let sqrt_val = num.sqrt()
    if (sqrt_val.floor() == sqrt_val.ceil()) {
        print(num, "is a perfect square")
    }
}

// Mathematical calculations with bounds
let angle = 1.57079  // approximately π/2
let sin_approx = angle - angle.pow(3) / (3).factorial()
print("sin approximation:", sin_approx.round(6))

// Working with ranges and validation
let score = 87.5
let normalized = score.clamp(0.0, 100.0) / 100.0
print("Normalized score:", normalized.round(3))
```
