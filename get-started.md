# VintLang Programming Documentation
This documentation introduces VintLang programming concepts and progresses from basic to advanced usage with well-structured examples.

---

## 1. Basics

### 1.1 Variables and Data Types
```vint
// Define a variable
let name = "VintLang"
print(name)  // Output: VintLang
```

### 1.2 String Operations
```vint
let name = "VintLang"
let characters = name.split("")
for i in characters {
    print(i)  // Outputs each character of the string
}

// Define the variable 'name' and assign the string "Tachera Sasi" to it
name = "Tachera Sasi"

// Split the string into an array of characters and print the result
print(name.split("")) 

// Reverse the string and print the result
print(name.reverse()) 

// Get the length of the string and print it
print(name.len()) 

// Convert the string to uppercase and print it
print(name.upper()) 

// Convert the string to lowercase and print it
print(name.lower()) 

// Check if the string contains the substring "sasi" (case-sensitive) and print the result
print(name.contains("sasi")) 

// Convert the string to uppercase and check if it contains the substring "SASI" (case-sensitive), then print the result
print(name.upper().contains("SASI")) 

// Replace the substring "Sasi" with "Vint" and print the result
print(name.replace("Sasi", "Vint")) 

// Trim any occurrence of the character "a" from the start and end of the string and print the result
print(name.trim("a"))

print(string(123))           // "123"
print(string(true))          // "true"
print(string(12.34))         // "12.34"
print(string("Hello World")) // "Hello World"


print(int("123"))    // 123
print(int(12.34))    // 12
print(int(true))     // 1
print(int(false))    // 0

```

---

## 2. Control Flow

### 2.1 Conditional Statements
```vint
let age = "10"
age = convert(age, "INTEGER")  // Convert string to integer

if (age == 20) {
    print("Age is 20")
} else if (age == 10) {
    print("Age is 10")
} else {
    print("Age is unknown")
}
```

### 2.2 Loops
```vint
let numbers = [1, 2, 3, 4]
for n in numbers {
    print(n)
}
```

---

## 3. Functions

### 3.1 Defining Functions
```vint
let printDetails = func(name, age, height) {
    print("My name is " + name + ", I am " + age + " years old, and my height is " + height + " feet.")
}

// Function call
let name = "VintLang"
let age = 10
let height = "6.0"
printDetails(name, age, height)
```

---

## 4. Built-in Modules

### 4.1 Time Module
```vint
import time
print(time.now())  // Print current timestamp

let year = 2024
print("Is", year, "Leap year:", time.isLeapYear(year))
```

### 4.2 Networking (HTTP Requests)
```vint
import net
let res = net.get("https://tachera.com")
print(res)  // Prints the response
```

---

## 5. Advanced Features

### 5.1 String Manipulation
```vint
import string
let trimmed = string.trim("  Hello, World!  ")
print(trimmed)  // Output: "Hello, World!"

let upper = string.toUpper("hello")
print(upper)  // Output: "HELLO"

let replaced = string.replace("Hello, World!", "World", "Vint")
print(replaced)  // Output: "Hello, Vint!"
```

### 5.2 Regex Module
```vint
// Sample usage of the Regex module in Vint
import regex

// 1. Using match to check if a string matches a pattern
result = regex.match("^Hello", "Hello World")
print(result)  // Expected output: true

// 2. Using match to check if a string does not match a pattern
result = regex.match("^World", "Hello World")
print(result)  // Expected output: false

// 3. Using replaceString to replace part of a string with a new value
newString = regex.replaceString("World", "VintLang", "Hello World")
print(newString)  // Expected output: "Hello VintLang"

// 4. Using splitString to split a string by a regex pattern
words = regex.splitString("\\s+", "Hello World VintLang")
print(words)  // Expected output: ["Hello", "World", "VintLang"]

// 5. Using splitString to split a string by a comma
csv = regex.splitString(",", "apple,banana,orange")
print(csv)  // Expected output: ["apple", "banana", "orange"]

// 6. Using match with a more complex regex pattern
emailMatch = regex.match("^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\\.[a-zA-Z]{2,}$", "test@example.com")
print(emailMatch)  // Expected output: true

// 7. Using replaceString with a regex pattern to replace digits in a string
maskedString = regex.replaceString("\\d", "*", "My phone number is 123456789")
print(maskedString)  // Expected output: "My phone number is *********"

```

---

## 6. File System Operations

### 6.1 Using the OS Module
```vint
import os
os.writeFile("example.txt", "Hello, Vint!")
let content = os.readFile("example.txt")
print(content)  // Output: "Hello, Vint!"

let files = os.listDir(".")
print(files)  // List directory contents
```

---

## 7. Error Handling and Debugging

### 7.1 Switch Case
```vint
let n = 1
switch (n) {
    case 1 {
        print("is", n)
        break
    }
    default {
        print("is not")
        break
    }
}
```

---

## 8. Json
```
import json

//Example 1: Decode a JSON string
print("=== Example 1: Decode ===")
raw_json = '{"name": "John", "age": 30, "isAdmin": false, "friends": ["Jane", "Doe"]}'
decoded = json.decode(raw_json)
print("Decoded Object:", decoded)

//Example 2: Encode a Vint object to JSON
print("\n=== Example 2: Encode ===")
data = {
  "language": "Vint",
  "version": 1.0,
  "features": ["custom modules", "native objects"]
}
encoded_json = json.encode(data) //optional parameter indent
print("Encoded JSON:", encoded_json)

//Example 3: Pretty print a JSON string
print("\n=== Example 3: Pretty Print ===")
raw_json_pretty = '{"name":"John","age":30,"friends":["Jane","Doe"]}'
pretty_json = json.pretty(raw_json_pretty)
print("Pretty JSON:\n", pretty_json)

//Example 4: Merge two JSON objects
print("\n=== Example 4: Merge ===")
json1 = {"name": "John", "age": 30}
json2 = {"city": "New York", "age": 35}
merged_json = json.merge(json1, json2)
print("Merged JSON:", merged_json)

//Example 5: Get a value by key from a JSON object
print("\n=== Example 5: Get Value by Key ===")
json_object = {"name": "John", "age": 30, "city": "New York"}
value = json.get(json_object, "age")
print("Age:", value)

missing_value = json.get(json_object, "country")
print("Country (missing key):", missing_value)

```

## 9. Logical

```
// Sample Vint program demonstrating logical operators

// Define a function to test 'and', 'or', and 'not'
let test_logical_operators = func () {
    // Testing 'and' operator
    let result_and = and(1+2==3, false) // Should return false
    print("Result of true AND false: ", result_and)
    print(1+2==4)

    // Testing 'or' operator
    let result_or = or(false, true) // Should return true
    print("Result of false OR true: ", result_or)

    // Testing 'not' operator
    let result_not = not(true) // Should return false
    print("Result of NOT true: ", result_not)
}

// Call the function to test the logical operators
test_logical_operators()

```

## 10. OS

```
import os

// Exit with a status code
// os.exit(1)

// Run a shell command
result = os.run("ls -la")
print(result)
// print(os.run("go run . vintLang/main.vint"))

// Get and set environment variables
// os.setEnv("API_KEY", "12345")
api_key = os.getEnv("API_KEY")
print(api_key)

// Read and write files
os.writeFile("example.txt", "Hello, Vint!")
content = os.readFile("example.txt")
print(content)

// List directory contents
files = os.listDir(".")
print(files)

// Create a directory
os.makeDir("new_folder")

// Check if a file exists
exists = os.fileExists("example.txt")
print(exists) // Outputs: false

// Write a file and read it line by line
os.writeFile("example.txt", "Hello\nWorld")
lines = os.readLines("example.txt")
print(lines) // Outputs: ["Hello", "World"]

// Delete a file
//os.deleteFile("example.txt")


```

## 11.String module
```
// Sample usage of the string module
import "string"

// Example 1: Trim whitespace
result = string.trim("  Hello, World!  ")
print(result)  // Output: "Hello, World!"

// Example 2: Check if a string contains a substring
containsResult = string.contains("Hello, World!", "World")
print(containsResult)  // Output: true

// Example 3: Convert to uppercase
upperResult = string.toUpper("hello")
print(upperResult)  // Output: "HELLO"

// Example 4: Convert to lowercase
lowerResult = string.toLower("HELLO")
print(lowerResult)  // Output: "hello"

// Example 5: Replace a substring
replaceResult = string.replace("Hello, World!", "World", "Vint")
print(replaceResult)  // Output: "Hello, Vint!"

// Example 6: Split string into parts
splitResult = string.split("a,b,c,d", ",")
print(splitResult)  // Output: ["a", "b", "c", "d"]

// Example 7: Join string parts
joinResult = string.join(["a", "b", "c"], "-")
print(joinResult)  // Output: "a-b-c"

// Example 8: Get a substring
substringResult = string.substring("Hello, World!", 7, 12)
print(substringResult)  // Output: "World"

// Example 9: Get the length of a string
lengthResult = string.length("Hello")
print(lengthResult)  // Output: 5

// Example 10: Find index of a substring
indexResult = string.indexOf("Hello, World!", "World")
print(indexResult)  // Output: 7

// Example 11: Get a substring (valid start and end indices)
result = string.substring("Hello, World!", 0, 5)
print(result)  // Output: "Hello"

// Example 12: Invalid indices (start >= end)
result = string.substring("Hello, World!", 7, 3)
print(result)  // Output: Error: Invalid start or end index


/*
More methods for this module
simirality
*/

```

## 12. UUID
```
import uuid

print(uuid.generate())
```

## GuessingGame

```vint
import math

// greet("Tach")

// Guessing game in Vint
guess = input("Guess a number: ")
guess = convert(guess, "INTEGER")

number = 5  // Predefined correct number
print(math.PI())
while (number != guess) {
    // Check if the guess is correct
    if (number == guess) {
        print("Woo hoo you've guessed it right!")
        break
    }else if(number > guess){
        print("Too small")
    }else{
        print("Too big")
    }

    // Prompt for a new guess if incorrect
    guess = input("Guess Again: ")
    guess = convert(guess, "INTEGER") // Converting the guess input to an integer
}
print("Woo hoo you've guessed it right")
print("Game over!")

```

## Comprehensive Example

This example integrates modules and features to create a simple application.
```vint
import net
import time
import string
import regex
import os

// Function to greet based on time
let greet = func(nameParam) {
    let currentTime = time.now()
    print("Hello, " + nameParam + "! The current time is " + currentTime)
}

// Greet a user
greet("Vint User")

// Perform a network request
let response = net.get("https://example.com")
print(response)

// Write and read a file
os.writeFile("log.txt", "VintLang Log: " + time.now())
let logContent = os.readFile("log.txt")
print(logContent)
```

---

