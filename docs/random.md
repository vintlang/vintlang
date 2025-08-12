# Random Module in VintLang

The `random` module provides functions for generating random numbers and data.

## Functions

### `random.int(min, max)`
Returns a random integer in the range `[min, max]`, inclusive.

```js
num = random.int(1, 100)
print(num) // Outputs a random number between 1 and 100
```

### `random.float()`
Returns a random float in the range `[0.0, 1.0)`.

```js
f = random.float()
print(f)
```

### `random.string(length)`
Returns a random string of a given length.

```js
s = random.string(12)
print(s) // Outputs a random 12-character string
```

### `random.choice(array)`
Returns a random element from an array.

```js
items = ["apple", "banana", "cherry"]
item = random.choice(items)
print(item) // Outputs one of the fruits
``` 