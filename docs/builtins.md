# Built-in Functions in Vint

Vint has a number of built-in functions that are globally available to perform common tasks.

---

## I/O and System Functions

### `print(...)`
Prints messages to the standard output. It can take zero or more arguments, which will be printed with a space between them.
```js
print("Hello,", "world!") // Output: Hello, world!
print(1, 2, 3)         // Output: 1 2 3
```

### `println(...)`
Similar to `print`, but it adds a newline character at the end of the output.

### `input(prompt)`
Reads a line of input from the user from standard input. It can optionally take a string argument to use as a prompt.
```js
let name = input("Enter your name: ")
println("Hello,", name)
```

### `sleep(milliseconds)`
Pauses the program's execution for a specified duration in milliseconds.
```js
println("Waiting for 1 second...")
sleep(1000)
println("Done.")
```

### `exit(code)`
Terminates the program with a specified exit code. An exit code of `0` typically indicates success, while any other number indicates an error.
```js
if (some_error) {
    println("An error occurred!")
    exit(1)
}
```

---

## Type and General Information Functions

### `type(object)`
Returns a string representing the type of the given object.
```js
type(10)      // Output: "INTEGER"
type("hello") // Output: "STRING"
type([])      // Output: "ARRAY"
```

### `len(object)`
Returns the length of a string, array, or dictionary.
```js
len("hello")      // Output: 5
len([1, 2, 3])    // Output: 3
len({"a": 1})   // Output: 1
```

---

## Array Functions

### `append(array, element1, ...)`
Returns a *new* array with the given elements added to the end.
```js
let arr = [1, 2]
let new_arr = append(arr, 3, 4)
println(new_arr) // Output: [1, 2, 3, 4]
```

### `pop(array)`
Removes the last element from an array and returns that element. This function modifies the array in-place.
```js
let arr = [1, 2, 3]
let last = pop(arr)
println(last) // Output: 3
println(arr)  // Output: [1, 2]
```

---

## Dictionary Functions

### `keys(dictionary)`
Returns an array containing all the keys from a dictionary. The order is not guaranteed.
```js
let dict = {"name": "Alex", "age": 30}
println(keys(dict)) // Output: ["name", "age"] (or ["age", "name"])
```

### `values(dictionary)`
Returns an array containing all the values from a dictionary. The order corresponds to the order of the keys returned by `keys()`.
```js
let dict = {"name": "Alex", "age": 30}
println(values(dict)) // Output: ["Alex", 30] (or [30, "Alex"])
```

### `has_key(dictionary, key)`
Returns `true` if the dictionary contains the given key, and `false` otherwise. This is also available as a method on dictionary objects: `my_dict.has_key(key)`.
```js
let dict = {"a": 1}
println(has_key(dict, "a")) // Output: true
println(dict.has_key("b"))  // Output: false
```

---

## String and Character Functions

### `chr(integer)`
Returns a single-character string corresponding to the given integer ASCII code.
```js
println(chr(65)) // Output: "A"
```

### `ord(string)`
Returns the integer ASCII code of the first character of a given string.
```js
println(ord("A")) // Output: 65
```

---

## File Functions

### `open(filepath)`
Opens a file and returns a file object. This is typically used for reading file contents.
```js
let file = open("data.txt")
// You can then use methods on the file object
```

---

## Additional Built-in Functions

### String Functions

#### `startsWith(string, prefix)`
Checks if a string starts with the specified prefix.
```js
startsWith("VintLang", "Vint")    // Output: true
startsWith("hello", "hi")         // Output: false
```

#### `endsWith(string, suffix)`
Checks if a string ends with the specified suffix.
```js
endsWith("VintLang", "Lang")      // Output: true
endsWith("hello", "world")        // Output: false
```

### Array Functions

#### `indexOf(array, element)`
Returns the index of the first occurrence of an element in an array, or -1 if not found.
```js
let arr = [1, 2, 3, 2, 4]
indexOf(arr, 2)    // Output: 1
indexOf(arr, 5)    // Output: -1
```

### Type Checking Functions

#### `isInt(value)`
Returns true if the value is an integer.
```js
isInt(42)          // Output: true
isInt(3.14)        // Output: false
isInt("hello")     // Output: false
```

#### `isFloat(value)`
Returns true if the value is a float.
```js
isFloat(3.14)      // Output: true
isFloat(42)        // Output: false
isFloat("hello")   // Output: false
```

#### `isString(value)`
Returns true if the value is a string.
```js
isString("hello")  // Output: true
isString(42)       // Output: false
isString(3.14)     // Output: false
```

#### `isBool(value)`
Returns true if the value is a boolean.
```js
isBool(true)       // Output: true
isBool(false)      // Output: true
isBool(42)         // Output: false
```

#### `isArray(value)`
Returns true if the value is an array.
```js
isArray([1, 2, 3]) // Output: true
isArray("hello")   // Output: false
isArray(42)        // Output: false
```

#### `isDict(value)`
Returns true if the value is a dictionary.
```js
isDict({"key": "value"})  // Output: true
isDict([1, 2, 3])         // Output: false
isDict("hello")           // Output: false
```

#### `isNull(value)`
Returns true if the value is null.
```js
isNull(null)       // Output: true
isNull(42)         // Output: false
isNull("")         // Output: false
```

### Parsing Functions

#### `parseInt(string)`
Parses a string and returns an integer.
```js
parseInt("42")     // Output: 42
parseInt("-10")    // Output: -10
parseInt("abc")    // Error: cannot parse 'abc' as integer
```

#### `parseFloat(string)`
Parses a string and returns a float.
```js
parseFloat("3.14")    // Output: 3.14
parseFloat("-2.5")    // Output: -2.5
parseFloat("hello")   // Error: cannot parse 'hello' as float
```

---

## Note on Existing Modules

VintLang also provides specialized modules for advanced functionality:

- **Math functions** like `abs`, `min`, `max`, `sqrt`, etc. are available in the `math` module
- **String functions** like `toUpper`, `toLower`, `trim`, `contains`, etc. are available in the `string` module  
- **Random functions** like `random.int()` and `random.float()` are available in the `random` module
- **Array methods** like `reverse()` and `sort()` are available as methods on array objects

Use these modules for more advanced functionality:
```js
import math
import string
import random

let result = math.abs(-5)        // 5
let upper = string.toUpper("hi") // "HI"
let num = random.int(1, 10)      // Random number 1-10
let arr = [3, 1, 4].sort()       // [1, 3, 4]
```