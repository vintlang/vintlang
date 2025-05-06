# Module math


## Usage

To use the `math` in-built module follow the steps below:

1. You directly import the `math` in-built module and any required in-built modules in your vint code using the `import` keyword.

   ```js
   import math
   ```

2. Calling the in-built module methods:

   ```js
   print(math.e())
   ```

## Yaliyomo

This in-built module covers a wide range of mathematical operations, including :

- `Basic Mathematical Functions:`
- `Hyperbolic` & `Trigonometric Functions`
- `Exponential` & `Logarithmic Functions`
- `Rounding` & `Comparison Functions`

Here is an in-depth classification of the methods:

1. Trigonometric Functions:

   - `cos(n)`
   - `sin(n)`
   - `tan(n)`
   - `acos(n)`
   - `asin(n)`
   - `atan(n)`
   - `hypot(numbers)`

2. Hyperbolic Functions:

   - `cosh(n)`
   - `sinh(n)`
   - `tanh(n)`
   - `acosh(n)`
   - `asinh(n)`
   - `atanh(n)`

3. Exponential and Logarithmic Functions:

   - `exp(n)`
   - `expm1(n)`
   - `log(n)`
   - `log2(n)`
   - `log10(n)`
   - `log1p(n)`

4. Basic Mathematical Functions:

   - `abs(n)`
   - `sqrt(n)`
   - `cbrt(n)`
   - `root(x, n)`
   - `factorial(n)`
   - `sign(n)`

5. Rounding and Comparison Functions:

   - `ceil(n)`
   - `floor(n)`
   - `round(n)`
   - `max(numbers)`
   - `min(numbers)`

### 1. Constants:

- **PI**: Represents the mathematical constant `π`.
- **e**: Represents `Euler's Number`.
- **phi**: Represents the `Golden Ratio`.
- **ln10**: Represents the `natural logarithm of 10`.
- **ln2**: Represents the `natural logarithm of 2`.
- **log10e**: Represents the `base 10 logarithms` of Euler's number `(e)`.
- **log2e**: Represents the `base 2 logarithm` of Euler's number` (e)`.
- **sqrt1_2**: Represents the `square root` of `1/2`.
- **sqrt2**: Represents the `square root` of `2`.
- **sqrt3**: Represents the `square root` of `3`.
- **sqrt5**: Represents the `square root` of `5`.
- **EPSILON**: Represents a small value `2.220446049250313e-16`.

### 2. Methods:

1. **abs(namba)**

   - Description: Calculates the absolute value of a number.
   - Example: `math.abs(-42)` returns `42`.

2. **acos(n)**

   - Description: Calculates the arccosine of a number.
   - Example: `math.acos(0.5)` returns `1.0471975511965979`.

3. **acosh(n)**

   - Description: Calculates the inverse hyperbolic cosine of a number.
   - Example: `math.acosh(2.0)` returns `1.3169578969248166`.

4. **asin(n)**

   - Description: Calculates the arcsine of a number using the Taylor series.
   - Example: `math.arcsin(0.5)` returns `0.5235987755982988`.

5. **asinh(n)**

   - Description: Calculates the inverse hyperbolic sine of a number.
   - Example: `math.arsinh(2.0)` returns `1.4436354751788103`.

6. **atan(n)**

   - Description: Calculates the arctangent of a number using the Taylor series.
   - Example: `math.atan(1.0)` returns `0.7853981633974483`.

7. **atan2(y, x)**

   - Description: Calculates the arctangent of the quotient of its arguments.
   - Example: `math.atan2(1.0, 1.0)` returns `0.7853981633974483`.

8. **atanh(n)**

   - Description: Calculates the inverse hyperbolic tangent of a number.
   - Example: `math.atanh(0.5)` returns `0.5493061443340549`.

9. **cbrt(n)**

   - Description: Calculates the cube root of a number.
   - Example: `math.cbrt(8)` returns `2`.

10. **root(x, n)**

    - Description: Calculates the nth root of a number using the Newton-Raphson method.
    - Example: `math.root(27, 3)` returns `3`.

11. **ceil(n)**

    - Description: Rounds up to the smallest integer greater than or equal to a given number.
    - Example: `math.ceil(4.3)` returns `5`.

12. **cos(n)**

    - Description: Calculates the cosine of an angle in radians using the Taylor series.
    - Example: `math.cos(0.0)` returns `1`.

13. **cosh(n)**

    - Description: Calculates the hyperbolic cosine of a number.
    - Example: `math.cosh(0.0)` returns `1`.

14. **exp(n)**

    - Description: Calculates the value of Euler's number raised to the power of a given number.
    - Example: `math.exp(2.0)` returns `7.38905609893065`.

15. **expm1(n)**

    - Description: Calculates Euler's number raised to the power of a number minus 1.
    - Example: `math.expm1(1.0)` returns `1.718281828459045`.

16. **floor(n)**

    - Description: Rounds down to the largest integer less than or equal to a given number.
    - Example: `math.floor(4.7)` returns `4`.

17. **hypot(values)**

    - Description: Calculates the square root of the sum of squares of the given values.
    - Example: `math.hypot([3, 4])` returns `5`.

18. **log(n)**

    - Description: Calculates the natural logarithm of a number.
    - Example: `math.log(1.0)` returns `0`.

19. **log10(n)**

    - Description: Calculates the base 10 logarithm of a number.
    - Example: `math.log10(100.0)` returns `2`.

20. **log1p(n)**

    - Description: Calculates the natural logarithm of 1 plus the given number.
    - Example: `math.log1p(1.0)` returns `0.6931471805599453`.

21. **log2(n)**

    - Description: Calculates the base 2 logarithm of a number.
    - Example: `math.log2(8)` returns `3`.

22. **max(numbers)**

    - Description: Finds the maximum value in a list of numbers.
    - Example: `math.max([4, 2, 9, 5])` returns `9`.

23. **min(numbers)**

    - Description: Finds the minimum value in a list of numbers.
    - Example: `math.min([4, 2, 9, 5])` returns `2`.

24. **round(x, method)**

    - Description: Rounds a number to the nearest integer using the specified method.
    - Example: `math.round(4.6)` returns `5`.

25. **sign(n)**

    - Description: Determines the sign of a number.
    - Example: `math.sign(-5)` returns `-1`.

26. **sin(n)**

    - Description: Calculates the sine of an angle in radians using the Taylor series.
    - Example: `math.sin(1.0)` returns `0.8414709848078965`.

27. **sinh(n)**

    - Description: Calculates the hyperbolic sine of a number.
    - Example: `math.sinh(1.0)` returns `1.1752011936438014`.

28. **sqrt(n)**

    - Description: Calculates the square root of a number.
    - Example: `math.sqrt(4)` returns `2`.

29. **tan(n)**

    - Description: Calculates the tangent of an angle in radians.
    - Example: `math.tan(1.0)` returns `1.557407724654902`.

30. **tanh(n)**

    - Description: Calculates the hyperbolic tangent of a number.
    - Example: `math.tanh(1.0)` returns `0.7615941559557649`.

31. **factorial(n)**

    - Description: Calculates the factorial of a number.
    - Example: `math.factorial(5)` returns `120`.

