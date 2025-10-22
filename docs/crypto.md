# Crypto Module

The `crypto` module provides a set of functions for common cryptographic operations, including hashing, symmetric encryption (AES), asymmetric encryption (RSA), and digital signatures.

## Hashing Functions

### `hashMD5(data)`

Computes the MD5 hash of a string.

- `data` (string): The input string to hash.

**Returns:** A string representing the 32-character hexadecimal MD5 hash.

**Usage:**

```js
import crypto

let hashed = crypto.hashMD5("hello world")
println(hashed) // "5eb63bbbe01eeed093cb22bb8f5acdc3"
```

### `hashSHA256(data)`

Computes the SHA-256 hash of a string.

- `data` (string): The input string to hash.

**Returns:** A string representing the 64-character hexadecimal SHA-256 hash.

**Usage:**

```js
import crypto

let hashed = crypto.hashSHA256("hello world")
println(hashed) // "b94d27b9934d3e08a52e52d7da7dabfac484efe37a5380ee9088f7ace2efcde9"
```

## Symmetric Encryption Functions (AES)

### `encryptAES(data, key)`

Encrypts a string using AES (Advanced Encryption Standard).

- `data` (string): The plaintext string to encrypt.
- `key` (string): The encryption key. **Must be 16, 24, or 32 bytes long** for AES-128, AES-192, or AES-256 respectively.

**Returns:** A hexadecimal string representing the encrypted data.

**Usage:**

```js
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

```js
import crypto

let key = "a_16_byte_secret_key"
let encrypted = "..." // Result from encryptAES

let decrypted = crypto.decryptAES(encrypted, key)
println("Decrypted:", decrypted)
```

## Asymmetric Encryption Functions (RSA)

### `generateRSA(keySize)`

Generates an RSA key pair with the specified bit size.

- `keySize` (integer, optional): The key size in bits. Must be between 1024 and 4096. Defaults to 2048.

**Returns:** A dictionary containing `"private"` and `"public"` keys in PEM format.

**Usage:**

```js
import crypto

// Generate 2048-bit RSA key pair (default)
let keys = crypto.generateRSA()

// Generate 4096-bit RSA key pair
let strongKeys = crypto.generateRSA(4096)

println("Private key:", keys["private"])
println("Public key:", keys["public"])
```

### `encryptRSA(data, publicKey)`

Encrypts data using RSA public key encryption.

- `data` (string): The plaintext string to encrypt.
- `publicKey` (string): The RSA public key in PEM format.

**Returns:** A hexadecimal string representing the encrypted data.

**Usage:**

```js
import crypto

let keys = crypto.generateRSA(2048)
let message = "Secret message"

let encrypted = crypto.encryptRSA(message, keys["public"])
println("Encrypted:", encrypted)
```

### `decryptRSA(encryptedData, privateKey)`

Decrypts data using RSA private key decryption.

- `encryptedData` (string): The encrypted data in hexadecimal format.
- `privateKey` (string): The RSA private key in PEM format.

**Returns:** The original plaintext string.

**Usage:**

```js
import crypto

let keys = crypto.generateRSA(2048)
let encrypted = "..." // Result from encryptRSA

let decrypted = crypto.decryptRSA(encrypted, keys["private"])
println("Decrypted:", decrypted)
```

## Digital Signature Functions

### `signRSA(data, privateKey)`

Creates a digital signature using RSA private key and SHA-256 hashing.

- `data` (string): The data to sign.
- `privateKey` (string): The RSA private key in PEM format.

**Returns:** A hexadecimal string representing the digital signature.

**Usage:**

```js
import crypto

let keys = crypto.generateRSA(2048)
let document = "Important document content"

let signature = crypto.signRSA(document, keys["private"])
println("Signature:", signature)
```

### `verifyRSA(data, signature, publicKey)`

Verifies a digital signature using RSA public key and SHA-256 hashing.

- `data` (string): The original data that was signed.
- `signature` (string): The signature in hexadecimal format.
- `publicKey` (string): The RSA public key in PEM format.

**Returns:** A boolean value indicating whether the signature is valid.

**Usage:**

```js
import crypto

let keys = crypto.generateRSA(2048)
let document = "Important document content"
let signature = crypto.signRSA(document, keys["private"])

let isValid = crypto.verifyRSA(document, signature, keys["public"])
println("Signature valid:", isValid) // true

let isTampered = crypto.verifyRSA("Tampered content", signature, keys["public"])
println("Tampered signature valid:", isTampered) // false
```

## Complete Example

```js
import crypto

// Generate RSA key pair
let keys = crypto.generateRSA(2048)

// Test message
let message = "Hello, secure world!"

// Test encryption/decryption
let encrypted = crypto.encryptRSA(message, keys["public"])
let decrypted = crypto.decryptRSA(encrypted, keys["private"])
println("Encryption test:", message == decrypted)

// Test digital signatures
let signature = crypto.signRSA(message, keys["private"])
let isValid = crypto.verifyRSA(message, signature, keys["public"])
println("Signature test:", isValid)

// Test hashing
println("MD5:", crypto.hashMD5(message))
println("SHA256:", crypto.hashSHA256(message))
```

## Security Notes

- **RSA Key Size**: Use at least 2048-bit keys for security. 4096-bit keys provide stronger security but slower performance.
- **AES Keys**: Use strong, randomly generated keys. Store keys securely and never hardcode them in source code.
- **Digital Signatures**: Always verify signatures before trusting signed data.
- **MD5**: Consider deprecated for security-critical applications. Use SHA-256 or stronger hash functions instead.
