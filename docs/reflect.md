# Reflect Module

The `reflect` module provides runtime type inspection and reflection utilities for VintLang. It allows you to examine the type and structure of values, check for null, and determine if a value is an array, object, or function.

## Importing
```vint
import reflect
```

## Functions

### reflect.typeOf(value)
Returns the type name of the given value as a string.

- **Arguments:**
  - `value`: Any value
- **Returns:** String (e.g., "STRING", "ARRAY", "DICT", "NULL", "FUNCTION", etc.)
- **Example:**
  ```vint
  reflect.typeOf("hello")        // "STRING"
  reflect.typeOf([1,2,3])        // "ARRAY"
  reflect.typeOf({"a": 1})      // "DICT"
  reflect.typeOf(null)           // "NULL"
  reflect.typeOf(func() {})      // "FUNCTION"
  ```

### reflect.valueOf(value)
Returns the raw value passed in (identity function).

- **Arguments:**
  - `value`: Any value
- **Returns:** The same value
- **Example:**
  ```vint
  reflect.valueOf(42)            // 42
  reflect.valueOf("foo")         // "foo"
  ```

### reflect.isNil(value)
Checks if the value is `null`.

- **Arguments:**
  - `value`: Any value
- **Returns:** Boolean
- **Example:**
  ```vint
  reflect.isNil(null)            // true
  reflect.isNil(123)             // false
  ```

### reflect.isArray(value)
Checks if the value is an array.

- **Arguments:**
  - `value`: Any value
- **Returns:** Boolean
- **Example:**
  ```vint
  reflect.isArray([1,2,3])       // true
  reflect.isArray("not array")   // false
  ```

### reflect.isObject(value)
Checks if the value is a dictionary/object.

- **Arguments:**
  - `value`: Any value
- **Returns:** Boolean
- **Example:**
  ```vint
  reflect.isObject({"a": 1})    // true
  reflect.isObject([1,2,3])      // false
  ```

### reflect.isFunction(value)
Checks if the value is a function.

- **Arguments:**
  - `value`: Any value
- **Returns:** Boolean
- **Example:**
  ```vint
  let f = func(x) { x * 2 }
  reflect.isFunction(f)          // true
  reflect.isFunction(123)        // false
  ```

## Example Usage
See `examples/reflect.vint` for a full demonstration of all reflect module functions.
