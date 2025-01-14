# Regex Module in Vint

The `regex` module in Vint provides powerful functionality to perform regular expression operations like matching patterns, replacing strings, and splitting strings. Regular expressions are useful for text processing and pattern matching.

---

## Importing the Regex Module

To use the Regex module, import it as follows:

```vint
import regex
```

---

## Functions and Examples

### 1. Matching a Pattern with `match`
The `match` function checks if a string matches a specified pattern. It returns `true` if the string matches the pattern and `false` if it does not.

**Syntax**:
```vint
match(pattern, string)
```
- `pattern`: The regular expression pattern.
- `string`: The string to match against the pattern.

**Example**:
```vint
import regex

result = regex.match("^Hello", "Hello World")
print(result)  // Expected output: true
```
In this case, the string starts with `"Hello"`, so it matches the pattern.

---

### 2. Using `match` to Check Non-Matches
You can use `match` to check if a string does *not* match a given pattern. If the pattern is not found at the beginning of the string, it will return `false`.

**Example**:
```vint
import regex

result = regex.match("^World", "Hello World")
print(result)  // Expected output: false
```
Since the string does not start with `"World"`, the result is `false`.

---

### 3. Replacing Part of a String with `replaceString`
The `replaceString` function replaces parts of a string that match a pattern with a new value.

**Syntax**:
```vint
replaceString(pattern, replacement, string)
```
- `pattern`: The regular expression pattern.
- `replacement`: The string to replace the matched part with.
- `string`: The input string.

**Example**:
```vint
import regex

newString = regex.replaceString("World", "VintLang", "Hello World")
print(newString)  // Expected output: "Hello VintLang"
```
In this case, `"World"` is replaced by `"VintLang"`, resulting in `"Hello VintLang"`.

---

### 4. Splitting a String with `splitString`
The `splitString` function splits a string into a list of substrings based on a regular expression pattern.

**Syntax**:
```vint
splitString(pattern, string)
```
- `pattern`: The regular expression pattern used as a delimiter.
- `string`: The string to be split.

**Example**:
```vint
import regex

words = regex.splitString("\\s+", "Hello World VintLang")
print(words)  // Expected output: ["Hello", "World", "VintLang"]
```
Here, `\\s+` matches one or more whitespace characters, so the string is split into words.

---

### 5. Splitting a String by a Comma
You can also split a string by a specific delimiter, such as a comma, using `splitString`.

**Example**:
```vint
import regex

csv = regex.splitString(",", "apple,banana,orange")
print(csv)  // Expected output: ["apple", "banana", "orange"]
```
The string `"apple,banana,orange"` is split at each comma.

---

### 6. Matching a Complex Pattern
You can match more complex patterns, such as an email address, using `match` with a regex pattern.

**Example**:
```vint
import regex

emailMatch = regex.match("^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\\.[a-zA-Z]{2,}$", "test@example.com")
print(emailMatch)  // Expected output: true
```
This pattern matches valid email addresses, so `"test@example.com"` matches successfully.

---

### 7. Replacing Digits in a String
You can use `replaceString` to replace parts of a string that match a pattern, such as replacing digits with asterisks.

**Example**:
```vint
import regex

maskedString = regex.replaceString("\\d", "*", "My phone number is 123456789")
print(maskedString)  // Expected output: "My phone number is *********"
```
Here, `\\d` matches any digit, so the digits in the phone number are replaced with asterisks.

---

## Summary of Functions

| Function           | Description                                             | Example Output                             |
|--------------------|---------------------------------------------------------|--------------------------------------------|
| `match(pattern, string)`  | Checks if the string matches the given pattern.         | `true` or `false`                          |
| `replaceString(pattern, replacement, string)`  | Replaces parts of a string matching a pattern with a new value.  | A modified string                          |
| `splitString(pattern, string)` | Splits a string into a list of substrings based on the pattern.   | List of substrings                         |

---

The `regex` module is extremely useful for text manipulation, pattern matching, and string processing tasks in Vint. Whether you need to validate input, split strings, or replace parts of a string, the regex module provides powerful tools to handle these tasks efficiently.