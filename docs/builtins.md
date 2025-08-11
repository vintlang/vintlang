# Built-in Functions in Vint

Vint has a number of built-in functions that are globally available to perform common tasks.

---

## I/O and System Functions

### `print(...)`
Prints messages to the standard output. It can take zero or more arguments, which will be printed with a space between them.
```vint
print("Hello,", "world!") // Output: Hello, world!
print(1, 2, 3)         // Output: 1 2 3
```

### `println(...)`
Similar to `print`, but it adds a newline character at the end of the output.

### `input(prompt)`
Reads a line of input from the user from standard input. It can optionally take a string argument to use as a prompt.
```vint
let name = input("Enter your name: ")
println("Hello,", name)
```

### `sleep(milliseconds)`
Pauses the program's execution for a specified duration in milliseconds.
```vint
println("Waiting for 1 second...")
sleep(1000)
println("Done.")
```

### `exit(code)`
Terminates the program with a specified exit code. An exit code of `0` typically indicates success, while any other number indicates an error.
```vint
if (some_error) {
    println("An error occurred!")
    exit(1)
}
```

---

## Type and General Information Functions

### `type(object)`
Returns a string representing the type of the given object.
```vint
type(10)      // Output: "INTEGER"
type("hello") // Output: "STRING"
type([])      // Output: "ARRAY"
```

### `len(object)`
Returns the length of a string, array, or dictionary.
```vint
len("hello")      // Output: 5
len([1, 2, 3])    // Output: 3
len({"a": 1})   // Output: 1
```

---

## Array Functions

### `append(array, element1, ...)`
Returns a *new* array with the given elements added to the end.
```vint
let arr = [1, 2]
let new_arr = append(arr, 3, 4)
println(new_arr) // Output: [1, 2, 3, 4]
```

### `pop(array)`
Removes the last element from an array and returns that element. This function modifies the array in-place.
```vint
let arr = [1, 2, 3]
let last = pop(arr)
println(last) // Output: 3
println(arr)  // Output: [1, 2]
```

---

## Dictionary Functions

### `keys(dictionary)`
Returns an array containing all the keys from a dictionary. The order is not guaranteed.
```vint
let dict = {"name": "Alex", "age": 30}
println(keys(dict)) // Output: ["name", "age"] (or ["age", "name"])
```

### `values(dictionary)`
Returns an array containing all the values from a dictionary. The order corresponds to the order of the keys returned by `keys()`.
```vint
let dict = {"name": "Alex", "age": 30}
println(values(dict)) // Output: ["Alex", 30] (or [30, "Alex"])
```

### `has_key(dictionary, key)`
Returns `true` if the dictionary contains the given key, and `false` otherwise. This is also available as a method on dictionary objects: `my_dict.has_key(key)`.
```vint
let dict = {"a": 1}
println(has_key(dict, "a")) // Output: true
println(dict.has_key("b"))  // Output: false
```

---

## String and Character Functions

### `chr(integer)`
Returns a single-character string corresponding to the given integer ASCII code.
```vint
println(chr(65)) // Output: "A"
```

### `ord(string)`
Returns the integer ASCII code of the first character of a given string.
```vint
println(ord("A")) // Output: 65
```

---

## File Functions

### `open(filepath)`
Opens a file and returns a file object. This is typically used for reading file contents.
```vint
let file = open("data.txt")
// You can then use methods on the file object
```

---

## Math Functions

### `abs(number)`
Returns the absolute value of a number (integer or float).
```vint
abs(5)      // Output: 5
abs(-5)     // Output: 5
abs(-3.14)  // Output: 3.14
```

### `min(number1, number2, ...)`
Returns the minimum value from the given arguments. Accepts multiple numbers.
```vint
min(5, 3, 8, 1)        // Output: 1
min(-5, -2, -10)       // Output: -10
min(3.14, 2.5, 4.0)    // Output: 2.5
```

### `max(number1, number2, ...)`
Returns the maximum value from the given arguments. Accepts multiple numbers.
```vint
max(5, 3, 8, 1)        // Output: 8
max(-5, -2, -10)       // Output: -2
max(3.14, 2.5, 4.0)    // Output: 4.0
```

### `round(number)`
Rounds a number to the nearest integer.
```vint
round(3.14)   // Output: 3
round(3.64)   // Output: 4
round(-2.7)   // Output: -3
```

### `floor(number)`
Rounds a number down to the nearest integer.
```vint
floor(3.99)   // Output: 3
floor(-2.1)   // Output: -3
```

### `ceil(number)`
Rounds a number up to the nearest integer.
```vint
ceil(3.01)    // Output: 4
ceil(-2.9)    // Output: -2
```

### `sqrt(number)`
Returns the square root of a number. The number must be non-negative.
```vint
sqrt(4)       // Output: 2.0
sqrt(9)       // Output: 3.0
sqrt(2)       // Output: 1.4142135623730951
```

---

## Enhanced String Functions

### `upper(string)`
Converts a string to uppercase.
```vint
upper("hello")    // Output: "HELLO"
```

### `lower(string)`
Converts a string to lowercase.
```vint
lower("WORLD")    // Output: "world"
```

### `trim(string)`
Removes leading and trailing whitespace from a string.
```vint
trim("  hello  ")  // Output: "hello"
```

### `contains(string, substring)` or `contains(array, element)`
Checks if a string contains a substring, or if an array contains an element.
```vint
contains("hello", "ell")           // Output: true
contains(["a", "b", "c"], "b")     // Output: true
```

### `startsWith(string, prefix)`
Checks if a string starts with a given prefix.
```vint
startsWith("hello", "he")          // Output: true
startsWith("hello", "lo")          // Output: false
```

### `endsWith(string, suffix)`
Checks if a string ends with a given suffix.
```vint
endsWith("hello", "lo")            // Output: true
endsWith("hello", "he")            // Output: false
```

---

## Enhanced Array Functions

### `reverse(array)`
Returns a new array with elements in reverse order.
```vint
let arr = [1, 2, 3, 4, 5]
reverse(arr)      // Output: [5, 4, 3, 2, 1]
```

### `indexOf(array, element)`
Returns the index of the first occurrence of an element in an array, or -1 if not found.
```vint
let arr = [1, 2, 3, 4, 5]
indexOf(arr, 3)   // Output: 2
indexOf(arr, 6)   // Output: -1
```

### `sort(array)`
Returns a new sorted array. Works with numbers, strings, and mixed arrays.
```vint
let numbers = [5, 2, 8, 1, 9]
sort(numbers)     // Output: [1, 2, 5, 8, 9]

let strings = ["banana", "apple", "cherry"]
sort(strings)     // Output: ["apple", "banana", "cherry"]
```

---

## Random Functions

### `rand()`
Returns a random floating-point number between 0.0 and 1.0.
```vint
rand()            // Output: 0.7326588891711826 (example)
```

### `randInt(max)`
Returns a random integer from 0 to max-1.
```vint
randInt(10)       // Output: 7 (example, 0-9)
```

### `randInt(min, max)`
Returns a random integer from min to max-1.
```vint
randInt(5, 15)    // Output: 12 (example, 5-14)
```

---

## String Parsing Functions

### `parseInt(string)`
Parses a string and returns an integer. Throws an error if the string is not a valid integer.
```vint
parseInt("123")   // Output: 123
parseInt("-456")  // Output: -456
```

### `parseFloat(string)`
Parses a string and returns a float. Throws an error if the string is not a valid number.
```vint
parseFloat("3.14")  // Output: 3.14
parseFloat("-2.5")  // Output: -2.5
parseFloat("123")   // Output: 123.0
```

---

## Type Checking Functions

### `isInt(value)`
Returns `true` if the value is an integer, `false` otherwise.
```vint
isInt(42)         // Output: true
isInt(3.14)       // Output: false
```

### `isFloat(value)`
Returns `true` if the value is a float, `false` otherwise.
```vint
isFloat(3.14)     // Output: true
isFloat(42)       // Output: false
```

### `isString(value)`
Returns `true` if the value is a string, `false` otherwise.
```vint
isString("hello") // Output: true
isString(42)      // Output: false
```

### `isBool(value)`
Returns `true` if the value is a boolean, `false` otherwise.
```vint
isBool(true)      // Output: true
isBool("true")    // Output: false
```

### `isArray(value)`
Returns `true` if the value is an array, `false` otherwise.
```vint
isArray([1, 2, 3]) // Output: true
isArray("hello")   // Output: false
```

### `isDict(value)`
Returns `true` if the value is a dictionary, `false` otherwise.
```vint
isDict({"a": 1})   // Output: true
isDict([1, 2])     // Output: false
```

### `isNull(value)`
Returns `true` if the value is null, `false` otherwise.
```vint
let nullVar = null
isNull(nullVar)    // Output: true
isNull(42)         // Output: false
```