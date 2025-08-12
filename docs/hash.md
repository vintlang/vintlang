# Hash Module in Vint

The Hash module in Vint provides additional hashing algorithms that complement the existing crypto module. It currently supports SHA1 and SHA512 hashing functions for generating secure hash values from text input.

---

## Importing the Hash Module

To use the Hash module, simply import it:
```js
import hash
```

---

## Functions and Examples

### 1. SHA1 Hash (`sha1`)
The `sha1` function generates a SHA1 hash from the provided string data.

**Syntax**:
```js
sha1(data)
```

**Example**:
```js
import hash

print("=== SHA1 Hash Example ===")
data = "hello world"
sha1_hash = hash.sha1(data)
print("Input:", data)
print("SHA1 Hash:", sha1_hash)
// Output: SHA1 Hash: 2aae6c35c94fcfb415dbe95f408b9ce91ee846ed
```

---

### 2. SHA512 Hash (`sha512`)
The `sha512` function generates a SHA512 hash from the provided string data.

**Syntax**:
```js
sha512(data)
```

**Example**:
```js
import hash

print("=== SHA512 Hash Example ===")
data = "hello world"
sha512_hash = hash.sha512(data)
print("Input:", data)
print("SHA512 Hash:", sha512_hash)
// Output: SHA512 Hash: 309ecc489c12d6eb4cc40f50c902f2b4d0ed77ee511a7c7a9bcd3ca86d4cd86f989dd35bc5ff499670da34255b45b0cfd830e81f605dcf7dc5542e93ae9cd76f
```

---

## Complete Usage Example

```js
import hash

print("=== Hash Module Complete Example ===")

// Test different types of data
test_data = [
    "hello",
    "hello world",
    "VintLang Programming Language",
    "The quick brown fox jumps over the lazy dog"
]

for data in test_data {
    print("\nInput:", data)
    print("SHA1:  ", hash.sha1(data))
    print("SHA512:", hash.sha512(data))
}

// Example with password hashing
password = "mySecretPassword123"
print("\n=== Password Hashing ===")
print("Password SHA1:  ", hash.sha1(password))
print("Password SHA512:", hash.sha512(password))
```

---

## Use Cases

- **Password Storage**: Hash passwords before storing them in databases
- **Data Integrity**: Verify file integrity by comparing hash values
- **Digital Signatures**: Generate unique identifiers for data
- **Checksums**: Create checksums for data validation
- **Security**: Generate secure hash values for authentication

---

## Summary of Functions

| Function | Description                               | Output Length    |
|----------|-------------------------------------------|------------------|
| `sha1`   | Generates SHA1 hash from string data      | 40 characters    |
| `sha512` | Generates SHA512 hash from string data    | 128 characters   |

Both functions return hexadecimal string representations of the hash values, making them easy to store and compare.