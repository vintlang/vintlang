# VVault - Password & Token Generator

A CLI tool built with VintLang for generating passwords, OTPs, tokens, and random strings.

## Files

- `vvault.vint` — Package with generator functions
- `main.vint` — CLI entry point

## Usage

```bash
vint main.vint <command> <length>
```

## Commands

### `generate` — Random Alphanumeric String

```bash
vint main.vint generate 16
# Generated: aKx9mBvR3nQw7pLs
```

### `password` — Strong Password

Generates a password with lowercase, uppercase, digits, and special characters. Minimum length is 4.

```bash
vint main.vint password 20
# Password: aB3!xK9@mP2#nQ5&yZ8$
```

### `otp` — Numeric One-Time Password

Generates a digits-only OTP using cryptographically secure randomness.

```bash
vint main.vint otp 6
# OTP: 482031
```

### `token` — Hex Token

Generates a cryptographically secure hex token. The output is twice the given length (each byte = 2 hex chars).

```bash
vint main.vint token 32
# Token: a3f1b2c4d5e6f7089012abcd3456ef78a3f1b2c4d5e6f7089012abcd3456ef78
```

## Example

```bash
# Generate a 12-character random string
vint main.vint generate 12

# Generate a secure 16-character password
vint main.vint password 16

# Generate a 6-digit OTP for verification
vint main.vint otp 6

# Generate a 32-byte hex token for API keys
vint main.vint token 32
```

## Modules Used

- `random` — Random number/string generation (OTP, password, and token use `crypto/rand`)
- `cli` — Command-line argument parsing
- `os` — Process exit
