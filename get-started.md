# VintLang Programming Documentation
This documentation introduces VintLang programming concepts and progresses from basic to advanced usage with well-structured examples.

---

## 1. Basics

### 1.1 Variables and Data Types
```js
// Define a variable
let name = "VintLang"
print(name)  // Output: VintLang
```

### 1.2 String Operations
```js
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
```js
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
```js
let numbers = [1, 2, 3, 4]
for n in numbers {
    print(n)
}
```

---

## 3. Functions

### 3.1 Defining Functions
```js
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
```js
import time
print(time.now())  // Print current timestamp

let year = 2024
print("Is", year, "Leap year:", time.isLeapYear(year))
```

### 4.2 Networking (HTTP Requests)
```js
import net
let res = net.get("https://tachera.com")
print(res)  // Prints the response
```

---

## 5. Advanced Features

### 5.1 String Manipulation
```js
import string
let trimmed = string.trim("  Hello, World!  ")
print(trimmed)  // Output: "Hello, World!"

let upper = string.toUpper("hello")
print(upper)  // Output: "HELLO"

let replaced = string.replace("Hello, World!", "World", "Vint")
print(replaced)  // Output: "Hello, Vint!"
```

### 5.2 Regex Module
```js
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
```js
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
```js
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
```js
import json

//Example 1: Decode a JSON string
print("=== Example 1: Decode ===")
let raw_json = '{"name": "John", "age": 30, "isAdmin": false, "friends": ["Jane", "Doe"]}'
let decoded = json.decode(raw_json)
print("Decoded Object:", decoded)

//Example 2: Encode a Vint object to JSON
print("\n=== Example 2: Encode ===")
let data = {
  "language": "Vint",
  "version": 1.0,
  "features": ["custom modules", "native objects"]
}
let encoded_json = json.encode(data) //optional parameter indent
print("Encoded JSON:", encoded_json)

//Example 3: Pretty print a JSON string
print("\n=== Example 3: Pretty Print ===")
let raw_json_pretty = '{"name":"John","age":30,"friends":["Jane","Doe"]}'
let pretty_json = json.pretty(raw_json_pretty)
print("Pretty JSON:\n", pretty_json)

//Example 4: Merge two JSON objects
print("\n=== Example 4: Merge ===")
let json1 = {"name": "John", "age": 30}
let json2 = {"city": "New York", "age": 35}
let merged_json = json.merge(json1, json2)
print("Merged JSON:", merged_json)

//Example 5: Get a value by key from a JSON object
print("\n=== Example 5: Get Value by Key ===")
let json_object = {"name": "John", "age": 30, "city": "New York"}
let value = json.get(json_object, "age")
print("Age:", value)

let missing_value = json.get(json_object, "country")
print("Country (missing key):", missing_value)

```

## 9. Logical

```js
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

The OS module provides comprehensive system-level functionality for interacting with the operating system.

```js
import os

// Process information
print("Process ID:", os.getpid())
print("Parent Process ID:", os.getppid())
print("User ID:", os.getuid())

// System information  
print("CPU Count:", os.cpuCount())
print("Hostname:", os.hostname())
print("Page Size:", os.getpagesize())

// Environment variables
os.setEnv("MYVAR", "hello")
print("MYVAR:", os.getEnv("MYVAR"))

// Advanced environment functions
let pathResult = os.lookupEnv("PATH")
if (pathResult["exists"]) {
    print("PATH is set")
}

let expanded = os.expandEnv("Home is $HOME")
print(expanded)

// User directories
print("Home:", os.userHomeDir())
print("Cache:", os.userCacheDir())
print("Config:", os.userConfigDir())
print("Temp:", os.tempDir())

// File operations
os.writeFile("test.txt", "Hello, Vint!")
let content = os.readFile("test.txt")
print("Content:", content)

// File information
let fileInfo = os.stat("test.txt")
print("File size:", fileInfo["size"])
print("Is directory:", fileInfo["isDir"])

// Directory operations
os.mkdirAll("path/to/directory")
files = os.readDir(".")
print("Directory contents:", files)

// Cleanup
os.remove("test.txt")
print(content)

// List directory contents
let files = os.listDir(".")
print(files)

// Create a directory
os.makeDir("new_folder")

// Check if a file exists
let exists = os.fileExists("example.txt")
print(exists) // Outputs: false

// Write a file and read it line by line
os.writeFile("example.txt", "Hello\nWorld")
let lines = os.readLines("example.txt")
print(lines) // Outputs: ["Hello", "World"]

// Delete a file
//os.deleteFile("example.txt")

// Copy and move files
// os.copy("source.txt", "destination.txt")
// os.move("old_name.txt", "new_name.txt")


```

## 11.String module
```js
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
```js
import uuid

print(uuid.generate())
```

## GuessingGame

```js
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
---

## Inventory Game
```js
/*
THIS IS A SIMPLE TERMINAL GAME WRITTEN IN VINTLANG
*/
import time
import os
import json
import uuid

// Game initialization
let player = {
    "name": "",
    "health": 100,
    "inventory": [],
    "location": "start"
}

// Save game state to a file
let saveGame = func () {
    let saveData = json.encode(player)
    os.writeFile("savegame.json", saveData)
    print("Game saved!")
}

// Load game state from a file
let loadGame = func () {
    if (os.fileExists("savegame.json")) {
        let saveData = os.readFile("savegame.json")
        player = json.decode(saveData)
        print("Game loaded!")
    } else {
        print("No saved game found.")
    }
}

// Display player stats
let showStats = func () {
    print("Player Stats:")
    print("Name: " + player["name"])
    print("Health: " + string(player["health"]))
    print("Inventory: " + string(player["inventory"]))
    print("Location: " + player["location"])
}

// Handle game events
let handleEvent = func (event) {
    if (event["type"] == "item") {
        print("You found an item: " + event["name"])
        player["inventory"].push(event["name"])
    } else if (event["type"] == "enemy") {
        print("An enemy appears: " + event["name"])
        print("You lose 10 health!")
        player["health"] -= 10
    }
    if (player["health"] <= 0) {
        print("You died! Game Over.")
        os.exit(1)
    }
}

// Main game loop
let gameLoop = func () {
    while (true) {
        print("\nYou are at: " + player["location"])
        print("Choose an action: [explore, stats, save, quit]")
        let action = input("> ")

        if (action == "explore") {
            print("Exploring...")
            let event = {
                "type": "item",
                "name": "Mystic Key"
            }
            handleEvent(event)
        } else if (action == "stats") {
            showStats()
        } else if (action == "save") {
            saveGame()
        } else if (action == "quit") {
            print("Quitting game...")
            os.exit(0)
        } else {
            print("Invalid action!")
        }
    }
}

// Start the game
print("Welcome to the Adventure Game!")
print("Enter your player name:")
player["name"] = input(">>> ")

print("Hello, " + player["name"] + "! Let's begin.")
gameLoop()

```

## Comprehensive Example

This example integrates modules and features to create a simple application.
```js
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

