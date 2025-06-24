# Math Module

## Usage

To use the `math` module, import it into your Vint script:

```vint
import math
```

You can then call functions and access constants from the module:

```vint
println(math.PI())
println(math.sqrt(16))
```

## Contents

This module provides a wide range of mathematical functions and constants, including:

- Basic Mathematical Functions
- Hyperbolic & Trigonometric Functions
- Exponential & Logarithmic Functions
- Rounding & Comparison Functions

Here is a complete list of the available functions and constants:

### Constants

- **PI**: Represents the mathematical constant `π` (3.14159...).
- **e**: Represents Euler's Number (2.71828...).
- **phi**: Represents the Golden Ratio (1.61803...).
- **ln10**: Represents the natural logarithm of 10.
- **ln2**: Represents the natural logarithm of 2.
- **log10e**: Represents the base-10 logarithm of `e`.
- **log2e**: Represents the base-2 logarithm of `e`.
- **sqrt1_2**: Represents the square root of 1/2.
- **sqrt2**: Represents the square root of 2.
- **sqrt3**: Represents the square root of 3.
- **sqrt5**: Represents the square root of 5.
- **EPSILON**: Represents a very small number, often used for float comparisons.

### Functions

#### `abs(n)`
- **Description**: Calculates the absolute value of a number.
- **Example**: `math.abs(-42)` returns `42`.

#### `acos(n)`
- **Description**: Calculates the arccosine (inverse cosine) of a number in radians.
- **Example**: `math.acos(0.5)` returns `1.047...`.

#### `acosh(n)`
- **Description**: Calculates the inverse hyperbolic cosine of a number.
- **Example**: `math.acosh(2.0)` returns `1.316...`.

#### `asin(n)`
- **Description**: Calculates the arcsine (inverse sine) of a number in radians.
- **Example**: `math.asin(0.5)` returns `0.523...`.

#### `asinh(n)`
- **Description**: Calculates the inverse hyperbolic sine of a number.
- **Example**: `math.asinh(2.0)` returns `1.443...`.

#### `atan(n)`
- **Description**: Calculates the arctangent (inverse tangent) of a number in radians.
- **Example**: `math.atan(1.0)` returns `0.785...`.

#### `atan2(y, x)`
- **Description**: Calculates the arctangent of the quotient of its arguments (`y/x`) in radians.
- **Example**: `math.atan2(1.0, 1.0)` returns `0.785...`.

#### `atanh(n)`
- **Description**: Calculates the inverse hyperbolic tangent of a number.
- **Example**: `math.atanh(0.5)` returns `0.549...`.

#### `cbrt(n)`
- **Description**: Calculates the cube root of a number.
- **Example**: `math.cbrt(8)` returns `2.0`.

#### `ceil(n)`
- **Description**: Rounds a number up to the nearest integer.
- **Example**: `math.ceil(4.3)` returns `5`.

#### `cos(n)`
- **Description**: Calculates the cosine of an angle (in radians).
- **Example**: `math.cos(0.0)` returns `1.0`.

#### `cosh(n)`
- **Description**: Calculates the hyperbolic cosine of a number.
- **Example**: `math.cosh(0.0)` returns `1.0`.

#### `exp(n)`
- **Description**: Calculates `e` raised to the power of `n`.
- **Example**: `math.exp(2.0)` returns `7.389...`.

#### `expm1(n)`
- **Description**: Calculates `e` raised to the power of a number, minus 1.
- **Example**: `math.expm1(1.0)` returns `1.718...`.

#### `factorial(n)`
- **Description**: Calculates the factorial of a non-negative integer.
- **Example**: `math.factorial(5)` returns `120`.

#### `floor(n)`
- **Description**: Rounds a number down to the nearest integer.
- **Example**: `math.floor(4.7)` returns `4`.

#### `hypot(numbers)`
- **Description**: Calculates the square root of the sum of the squares of the numbers in an array.
- **Example**: `math.hypot([3, 4])` returns `5.0`.

#### `log10(n)`
- **Description**: Calculates the base-10 logarithm of a number.
- **Example**: `math.log10(100.0)` returns `2.0`.

#### `log1p(n)`
- **Description**: Calculates the natural logarithm of 1 plus the given number.
- **Example**: `math.log1p(1.0)` returns `0.693...`.

#### `log2(n)`
- **Description**: Calculates the base-2 logarithm of a number.
- **Example**: `math.log2(8)` returns `3.0`.

#### `max(numbers)`
- **Description**: Finds the maximum value in an array of numbers.
- **Example**: `math.max([4, 2, 9, 5])` returns `9.0`.

#### `min(numbers)`
- **Description**: Finds the minimum value in an array of numbers.
- **Example**: `math.min([4, 2, 9, 5])` returns `2.0`.

#### `random()`
- **Description**: Returns a random floating-point number between 0.0 and 1.0.
- **Example**: `math.random()` returns a value like `0.12345...`.

#### `round(n)`
- **Description**: Rounds a floating-point number to the nearest integer.
- **Example**: `math.round(4.6)` returns `5`.

#### `root(x, n)`
- **Description**: Calculates the nth root of a number `x`.
- **Example**: `math.root(27, 3)` returns `3.0`.

#### `sign(n)`
- **Description**: Returns the sign of a number (`-1` for negative, `0` for zero, `1` for positive).
- **Example**: `math.sign(-5)` returns `-1`.

#### `sin(n)`
- **Description**: Calculates the sine of an angle (in radians).
- **Example**: `math.sin(1.0)` returns `0.841...`.

#### `sinh(n)`
- **Description**: Calculates the hyperbolic sine of a number.
- **Example**: `math.sinh(1.0)` returns `1.175...`.

#### `sqrt(n)`
- **Description**: Calculates the square root of a number.
- **Example**: `math.sqrt(4)` returns `2.0`.

#### `tan(n)`
- **Description**: Calculates the tangent of an angle (in radians).
- **Example**: `math.tan(1.0)` returns `1.557...`.

#### `tanh(n)`
- **Description**: Calculates the hyperbolic tangent of a number.
- **Example**: `math.tanh(1.0)` returns `0.761...`.

