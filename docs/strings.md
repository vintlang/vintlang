# Strings in Vint

Strings are a sequence of characters used to represent text in the Vint programming language. Here’s a detailed explanation of how to work with strings, including syntax, manipulation, and useful built-in methods.

## Basic Syntax

In Vint, strings can be enclosed in either single quotes (`''`) or double quotes (`""`):

```js
print("Hello")  // Output: Hello

let name = 'Tachera'

print("Hello", name)  // Output: Hello Tachera
```

## Concatenating Strings

Strings can be concatenated using the `+` operator:

```js
let greeting = "Hello" + " " + "World"
print(greeting)  // Output: Hello World

let message = "Hello"
message += " World"
// Output: Hello World
```

You can also repeat a string a specific number of times using the `*` operator:

```js
print("Hello " * 3)  // Output: Hello Hello Hello

let repeated = "World"
repeated *= 2
// Output: WorldWorld
```

## Looping Over a String

You can loop through each character of a string using the `for` keyword:

```js
let name = "Avicenna"

for char in name {
    print(char)
}
// Output:
// A
// v
// i
// c
// e
// n
// n
// a
```

You can also loop through the string using its index and character:

```js
for i, char in name {
    print(i, "=>", char)
}
// Output:
// 0 => A
// 1 => v
// 2 => i
// 3 => c
// 4 => e
// 5 => n
// 6 => n
// 7 => a
```

## Comparing Strings

You can compare two strings using the `==` operator:

```js
let a = "Vint"
print(a == "Vint")  // Output: true
print(a == "vint")  // Output: false
```

## String Methods

### Length of a String (`length`)

You can find the length of a string using the `length` method. It does not accept any parameters:

```js
let message = "Vint"
print(message.length())  // Output: 4
```

### Convert to Uppercase (`upper`)

This method converts the string to uppercase:

```js
let text = "vint"
print(text.upper())  // Output: VINT
```

### Convert to Lowercase (`lower`)

This method converts the string to lowercase:

```js
let text = "VINT"
print(text.lower())  // Output: vint
```

### Split a String (`split`)

The `split` method splits a string into an array based on a specified delimiter. If no delimiter is provided, it splits by whitespace.

Example without a delimiter:

```js
let sentence = "Vint programming language"
let words = sentence.split()
print(words)  // Output: ["Vint", "programming", "language"]
```

Example with a delimiter:

```js
let sentence = "Vint,programming,language"
let words = sentence.split(",")
print(words)  // Output: ["Vint", "programming", "language"]
```

### Replace Substrings (`replace`)

You can replace a substring with another string using the `replace` method:

```js
let greeting = "Hello World"
let newGreeting = greeting.replace("World", "Vint")
print(newGreeting)  // Output: Hello Vint
```

### Trim Whitespace (`trim`)

You can remove whitespace from the start and end of a string using the `trim` method:

```js
let message = "  Hello World  "
print(message.trim())  // Output: Hello World
```

### Get a Substring (`substring`)

You can extract a substring from a string by specifying the starting and ending indices:

```js
let sentence = "Vint programming"
print(sentence.substring(0, 4))  // Output: Vint
```

### Find the Index of a Substring (`indexOf`)

You can find the index of a substring within a string using the `indexOf` method:

```js
let sentence = "Vint programming"
print(sentence.indexOf("programming"))  // Output: 5
```

### Slugify a String (`slug`)

You can convert a string into a URL-friendly format (slug) using the `slug` method:

```js
let title = "Creating a Slug String"
print(title.slug())  // Output: creating-a-slug-string
```

### Checking Substring Presence (`contains`)

Check if a string contains a specific substring:

```js
let name = "Tachera Sasi"
print(name.contains("Sasi"))  // Output: true
```

### Get Character at Index (`charAt`)

Get the character at a specific index:

```js
let word = "Hello"
print(word.charAt(1))  // Output: e
print(word.charAt(10)) // Output: "" (empty string for out of bounds)
```

### Repeat String (`times`)

Repeat a string a specified number of times:

```js
let pattern = "Ha"
print(pattern.times(3))  // Output: HaHaHa
```

### Pad String Start (`padStart`)

Pad the string to a target length from the beginning:

```js
let num = "5"
print(num.padStart(3, "0"))  // Output: 005

let word = "hi"
print(word.padStart(5, "*"))  // Output: ***hi
```

### Pad String End (`padEnd`)

Pad the string to a target length from the end:

```js
let num = "5"
print(num.padEnd(3, "0"))  // Output: 500

let word = "hi"
print(word.padEnd(5, "*"))  // Output: hi***
```

### Check String Start (`startsWith`)

Check if a string starts with a specified prefix:

```js
let message = "Hello World"
print(message.startsWith("Hello"))  // Output: true
print(message.startsWith("World"))  // Output: false
```

### Check String End (`endsWith`)

Check if a string ends with a specified suffix:

```js
let filename = "document.pdf"
print(filename.endsWith(".pdf"))  // Output: true
print(filename.endsWith(".txt"))  // Output: false
```

### Extract Slice (`slice`)

Extract a section of the string:

```js
let text = "Hello World"
print(text.slice(0, 5))   // Output: Hello
print(text.slice(6))      // Output: World
print(text.slice(-5))     // Output: World
```

## Example Usage

Here’s an example of how you might use these string operations in Vint:

```js
import "string"

// Example: Trim whitespace
let trimmed = string.trim("  Hello, World!  ")
print(trimmed)  // Output: "Hello, World!"

// Example: Check if a string contains a substring
let containsResult = string.contains("Hello, World!", "World")
print(containsResult)  // Output: true

// Example: Convert to uppercase
let upperResult = string.toUpper("hello")
print(upperResult)  // Output: "HELLO"

// Example: Convert to lowercase
let lowerResult = string.toLower("HELLO")
print(lowerResult)  // Output: "hello"

// Example: Replace a substring
let replaceResult = string.replace("Hello, World!", "World", "Vint")
print(replaceResult)  // Output: "Hello, Vint!"

// Example: Split a string into parts
let splitResult = string.split("a,b,c,d", ",")
print(splitResult)  // Output: ["a", "b", "c", "d"]

// Example: Join string parts
let joinResult = string.join(["a", "b", "c"], "-")
print(joinResult)  // Output: "a-b-c"

// Example: Get the length of a string
let lengthResult = string.length("Hello")
print(lengthResult)  // Output: 5
```

## Example with Vint Data

Here's an example using Vint-specific strings:

```js
let name = "Tachera Sasi"
let reversed = name.reverse()
print(reversed)  // Output: "isaS arehcaT"

let upperName = name.upper()
print(upperName)  // Output: "TACHERA SASI"

let trimmedName = name.trim("T")
print(trimmedName)  // Output: "achera Sasi"
```

Understanding how to manipulate and work with strings in Vint allows you to efficiently handle text data in your programs.