# Random Module in VintLang

The `random` module provides functions for generating random numbers and data.

## Functions

### `random.int(min, max)`
Returns a random integer in the range `[min, max]`, inclusive.

```vint
num = random.int(1, 100)
print(num) // Outputs a random number between 1 and 100
```

### `random.float()`
Returns a random float in the range `[0.0, 1.0)`.

```vint
f = random.float()
print(f)
```

### `random.string(length)`
Returns a random string of a given length.

```vint
s = random.string(12)
print(s) // Outputs a random 12-character string
```

### `random.choice(array)`
Returns a random element from an array.

```vint
items = ["apple", "banana", "cherry"]
item = random.choice(items)
print(item) // Outputs one of the fruits
``` 