# Random Module in VintLang

The `random` module provides functions for generating random numbers, strings, OTPs, tokens, passwords, and more.

## Functions

### `random.int(min, max)`

Returns a random integer in the range `[min, max]`, inclusive.

```js
num = random.int(1, 100);
print(num); // Outputs a random number between 1 and 100
```

### `random.float()`

Returns a random float in the range `[0.0, 1.0)`.

```js
f = random.float();
print(f);
```

### `random.string(length)`

Returns a random alphabetic string of a given length.

```js
s = random.string(12);
print(s); // Outputs a random 12-character string
```

### `random.choice(array)`

Returns a random element from an array.

```js
items = ["apple", "banana", "cherry"];
item = random.choice(items);
print(item); // Outputs one of the fruits
```

### `random.otp(length)`

Generates a cryptographically secure numeric OTP (One-Time Password) of the given length.

```js
code = random.otp(6);
print(code); // e.g. "482031"
```

### `random.token(byteLength)`

Generates a cryptographically secure hex token. The resulting string is twice the byte length (since each byte becomes 2 hex characters).

```js
tok = random.token(16);
print(tok); // e.g. "a3f1b2c4d5e6f7089012abcd3456ef78" (32 hex chars)

apiKey = random.token(32);
print(apiKey); // 64 hex chars, suitable for API keys
```

### `random.password(length)`

Generates a strong random password containing lowercase, uppercase, digits, and special characters. Minimum length is 4 to guarantee at least one of each type. Uses cryptographically secure randomness.

```js
pw = random.password(16);
print(pw); // e.g. "aB3!xK9@mP2#nQ5&"
```

### `random.shuffle(array)`

Returns a new shuffled copy of the array (does not modify the original).

```js
arr = [1, 2, 3, 4, 5];
shuffled = random.shuffle(arr);
print(shuffled); // e.g. [3, 1, 5, 2, 4]
print(arr); // [1, 2, 3, 4, 5] (unchanged)
```

### `random.sample(array, count)`

Returns `count` unique random elements from an array (sampling without replacement).

```js
nums = [1, 2, 3, 4, 5, 6, 7, 8, 9, 10];
picked = random.sample(nums, 3);
print(picked); // e.g. [7, 2, 9]
```

### `random.bool()`

Returns a random boolean value (`true` or `false`).

```js
val = random.bool();
print(val); // true or false
```

### `random.range(min, max, count)`

Returns an array of `count` random integers, each in the range `[min, max]`.

```js
nums = random.range(1, 100, 5);
print(nums); // e.g. [42, 17, 88, 3, 56]
```
