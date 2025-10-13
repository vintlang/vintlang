# Math Module

## Usage

To use the `math` module, import it into your Vint script:

```js
import math
```

You can then call functions and access constants from the module:

```js
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

---

## Statistics Functions

#### `mean(array)`
- **Description**: Calculates the arithmetic mean (average) of an array of numbers.
- **Example**: `math.mean([1, 2, 3, 4, 5])` returns `3.0`.

#### `median(array)`
- **Description**: Calculates the median (middle value) of an array of numbers.
- **Example**: `math.median([1, 2, 3, 4, 5])` returns `3.0`.
- **Note**: For even-length arrays, returns the average of the two middle values.

#### `variance(array)`
- **Description**: Calculates the variance of an array of numbers.
- **Example**: `math.variance([1, 2, 3, 4, 5])` returns `2.0`.

#### `stddev(array)`
- **Description**: Calculates the standard deviation of an array of numbers.
- **Example**: `math.stddev([1, 2, 3, 4, 5])` returns `1.414...`.

---

## Complex Numbers

#### `complex(real, imag)`
- **Description**: Creates a complex number with the given real and imaginary parts.
- **Returns**: A dictionary with `real` and `imag` keys.
- **Example**: 
  ```js
  let c = math.complex(3, 4)
  print(c["real"])  // 3
  print(c["imag"])  // 4
  ```

#### `abs(n)` (extended)
- **Description**: Also works with complex numbers to calculate magnitude.
- **Example**: 
  ```js
  let c = math.complex(3, 4)
  math.abs(c)  // returns 5.0
  ```

---

## Arbitrary Precision

#### `bigint(value)`
- **Description**: Creates a big integer representation for arbitrary precision arithmetic.
- **Parameters**: A string or integer representing a large number.
- **Returns**: A dictionary with `value` (string) and `type` ("bigint") keys.
- **Example**: 
  ```js
  let big = math.bigint("999999999999999999999")
  print(big["value"])  // "999999999999999999999"
  ```

---

## Linear Algebra

#### `dot(array1, array2)`
- **Description**: Calculates the dot product of two vectors (arrays).
- **Example**: `math.dot([1, 2, 3], [4, 5, 6])` returns `32.0`.

#### `cross(array1, array2)`
- **Description**: Calculates the cross product of two 3D vectors.
- **Example**: `math.cross([1, 2, 3], [4, 5, 6])` returns `[-3, 6, -3]`.

#### `magnitude(array)`
- **Description**: Calculates the magnitude (length) of a vector.
- **Example**: `math.magnitude([3, 4])` returns `5.0`.

---

## Numerical Methods

#### `gcd(a, b)`
- **Description**: Calculates the greatest common divisor of two integers.
- **Example**: `math.gcd(48, 18)` returns `6`.

#### `lcm(a, b)`
- **Description**: Calculates the least common multiple of two integers.
- **Example**: `math.lcm(12, 15)` returns `60`.

#### `clamp(value, min, max)`
- **Description**: Clamps a value between a minimum and maximum.
- **Example**: `math.clamp(15, 0, 10)` returns `10.0`.

#### `lerp(start, end, t)`
- **Description**: Linear interpolation between two values.
- **Parameters**: `start` and `end` values, and `t` (0.0 to 1.0) as the interpolation factor.
- **Example**: `math.lerp(0, 10, 0.5)` returns `5.0`.

