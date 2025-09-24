
# Crypto Module

The `crypto` module provides a set of functions for common cryptographic operations, including hashing and encryption.

## Hashing Functions

### `hashMD5(data)`

Computes the MD5 hash of a string.

- `data` (string): The input string to hash.

**Returns:** A string representing the 32-character hexadecimal MD5 hash.

**Usage:**

```vint
import crypto

let hashed = crypto.hashMD5("hello world")
println(hashed) // "5eb63bbbe01eeed093cb22bb8f5acdc3"
```

### `hashSHA256(data)`

Computes the SHA-256 hash of a string.

- `data` (string): The input string to hash.

**Returns:** A string representing the 64-character hexadecimal SHA-256 hash.

**Usage:**

```vint
import crypto

let hashed = crypto.hashSHA256("hello world")
println(hashed) // "b94d27b9934d3e08a52e52d7da7dabfac484efe37a5380ee9088f7ace2efcde9"
```

## Encryption & Decryption Functions

### `encryptAES(data, key)`

Encrypts a string using AES (Advanced Encryption Standard).

- `data` (string): The plaintext string to encrypt.
- `key` (string): The encryption key. **Must be 16, 24, or 32 bytes long** for AES-128, AES-192, or AES-256 respectively.

**Returns:** A hexadecimal string representing the encrypted data.

**Usage:**

```vint
import crypto

let secret = "this is a secret message"
let key = "a_16_byte_secret_key"

let encrypted = crypto.encryptAES(secret, key)
println("Encrypted:", encrypted)
```

### `decryptAES(encryptedData, key)`

Decrypts an AES-encrypted hexadecimal string.

- `encryptedData` (string): The hexadecimal string to decrypt.
- `key` (string): The decryption key. Must be the same key used for encryption.

**Returns:** The original plaintext string.

**Usage:**

```vint
import crypto

let key = "a_16_byte_secret_key"
let encrypted = "..." // Result from encryptAES

let decrypted = crypto.decryptAES(encrypted, key)
println("Decrypted:", decrypted)
```
