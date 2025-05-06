# Range Function (`range`)

The `range` function in Vint generates a sequence of numbers and is commonly used in loops or for creating arrays of sequential values.

---

## Syntax

```js
range(end)
range(start, end)
range(start, end, step)
```

---

## Parameters

- **`end`**: The upper limit of the sequence (exclusive).
- **`start`** (optional): The starting value of the sequence. Default is `0`.
- **`step`** (optional): The increment or decrement between each number in the sequence. Default is `1`.

---

## Return Value

The function returns an array of integers.

---

## Examples

### Basic Usage
```js
// Generate numbers from 0 to 4
for i in range(5) {
    print(i)
}
// Output: 0 1 2 3 4
```

### Specifying a Start and End
```js
// Generate numbers from 1 to 9
for i in range(1, 10) {
    print(i)
}
// Output: 1 2 3 4 5 6 7 8 9
```

### Using a Step Value
```js
// Generate even numbers from 0 to 8
for i in range(0, 10, 2) {
    print(i)
}
// Output: 0 2 4 6 8
```

### Generating a Reverse Sequence
```js
// Generate numbers in reverse order
for i in range(10, 0, -1) {
    print(i)
}
// Output: 10 9 8 7 6 5 4 3 2 1
```

---

## Notes

1. **Exclusive End**: The `end` value is not included in the sequence; the range stops before reaching it.
2. **Negative Steps**: If a negative `step` is provided, ensure `start` is greater than `end` to create a valid reverse sequence.
3. **Non-Zero Step**: The `step` value cannot be `0`, as it would result in an infinite loop or an error.

---
