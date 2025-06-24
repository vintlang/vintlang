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