# encoding Module in Vint

The `encoding` module in Vint provides functions for encoding and decoding data in various formats. This includes commonly used encoding schemes like Base64.

---

## Importing the encoding Module

To use the `encoding` module, import it as follows:

```vint
import encoding
```

---

## Functions and Examples

### 1. Base64 Encoding with `base64Encode()`
The `base64Encode` function encodes a string into Base64 format. Base64 encoding is often used for encoding binary data as text, making it suitable for transmission over text-based protocols such as email or HTTP.

**Syntax**:
```vint
base64Encode(inputString)
```
- `inputString`: The string you want to encode.

**Example**:
```vint
import encoding

encoded = encoding.base64Encode("Hello, World!")
print(encoded)  // Expected output: "SGVsbG8sIFdvcmxkIQ=="
```
In this example, the string `"Hello, World!"` is encoded into Base64 format.

---

### 2. Base64 Decoding with `base64Decode()`
The `base64Decode` function decodes a Base64-encoded string back into its original format.

**Syntax**:
```vint
base64Decode(encodedString)
```
- `encodedString`: The Base64-encoded string that you want to decode.

**Example**:
```vint
import encoding

encoded = encoding.base64Encode("Hello, World!")
print(encoded)  // Expected output: "SGVsbG8sIFdvcmxkIQ=="

decoded = encoding.base64Decode(encoded)
print(decoded)  // Expected output: "Hello, World!"
```
In this example, the Base64-encoded string is decoded back to its original value.

---

## Summary of Functions

| Function               | Description                                        | Example Output                             |
|------------------------|----------------------------------------------------|--------------------------------------------|
| `base64Encode(input)`   | Encodes a string to Base64 format.                 | `"SGVsbG8sIFdvcmxkIQ=="`                   |
| `base64Decode(encoded)` | Decodes a Base64-encoded string back to its original form. | `"Hello, World!"`                           |

---

The `encoding` module in Vint is essential for working with different encoding schemes such as Base64. It simplifies the process of converting data between text and binary formats, making it easier to handle data transmission or storage in encoded formats.