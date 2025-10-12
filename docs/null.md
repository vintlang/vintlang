# Null in Vint

The `null` data type in Vint represents the absence of a value or the concept of "nothing" or "empty." This page covers the syntax and usage of the `null` data type in Vint, including its definition and evaluation.

## Definition

A `null` data type is a data type with no value, defined with the `null` keyword:

```js
let a = null
```

## Evaluation

When evaluating a `null` data type in a conditional expression, it will evaluate to `false`:

```js
if (a) {
    print("a is null")
} else {
    print("a has a value")
}

// Output: a has a value
```

## Null Methods

The `null` data type in Vint comes with several utility methods:

### isNull()

Always returns `true` for null values:

```js
let value = null
print(value.isNull())  // true
```

### coalesce()

Returns the first non-null value from the arguments:

```js
let value = null
let result = value.coalesce("default", "backup")
print(result)  // "default"
```

### ifNull()

Returns the provided value if this is null:

```js
let value = null
let result = value.ifNull("default value")
print(result)  // "default value"
```

### toString()

Returns the string representation of null:

```js
let value = null
print(value.toString())  // "null"
```

### equals()

Checks if another value is also null:

```js
let value1 = null
let value2 = null
let value3 = "something"
print(value1.equals(value2))  // true
print(value1.equals(value3))  // false
```

The `null` data type is useful in Vint when you need to represent an uninitialized, missing, or undefined value in your programs. By understanding the `null` data type and its methods, you can create more robust and flexible code.