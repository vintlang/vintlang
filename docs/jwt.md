# JWT Module

The JWT module provides functions for creating, verifying, and decoding JSON Web Tokens (JWTs). It supports HMAC-SHA256 signing by default.

## Functions

### `jwt.create(payload, secret)`

Creates a JWT token with the provided payload and secret using HS256 signing method.

**Parameters:**

- `payload` (dict): The claims/payload to include in the JWT
- `secret` (string): The secret key used for signing

**Returns:**

- `string`: The JWT token string

**Example:**

```js
let payload = { user: "john", role: "admin" };
let token = jwt.create(payload, "my-secret-key");
print(token); // eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9...
```

### `jwt.createHS256(payload, secret, [expiration_hours])`

Creates a JWT token with HS256 signing method, with optional expiration.

**Parameters:**

- `payload` (dict): The claims/payload to include in the JWT
- `secret` (string): The secret key used for signing
- `expiration_hours` (number, optional): Token expiration time in hours

**Returns:**

- `string`: The JWT token string

**Example:**

```js
let payload = { user: "john", role: "admin" };
let token = jwt.createHS256(payload, "my-secret-key", 24); // Expires in 24 hours
print(token);
```

### `jwt.verify(token, secret)`

Verifies a JWT token and returns the payload if valid.

**Parameters:**

- `token` (string): The JWT token to verify
- `secret` (string): The secret key used for verification

**Returns:**

- `dict`: The token payload if valid
- `error`: Error if token is invalid or verification fails

**Example:**

```js
let token = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9...";
let result = jwt.verify(token, "my-secret-key");

if (result.type != "ERROR") {
  print("User:", result["user"]);
  print("Role:", result["role"]);
} else {
  print("Invalid token:", result.message);
}
```

### `jwt.verifyHS256(token, secret)`

Verifies a JWT token with explicit HS256 signing method check.

**Parameters:**

- `token` (string): The JWT token to verify
- `secret` (string): The secret key used for verification

**Returns:**

- `dict`: The token payload if valid
- `error`: Error if token is invalid or not HS256 signed

**Example:**

```js
let token = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9...";
let result = jwt.verifyHS256(token, "my-secret-key");

if (result.type != "ERROR") {
  print("Verified HS256 token:", result);
}
```

### `jwt.decode(token)`

Decodes a JWT token without verification (useful for inspecting headers/payload).

**Parameters:**

- `token` (string): The JWT token to decode

**Returns:**

- `dict`: Contains "header" and "payload" keys with their respective data
- `error`: Error if token format is invalid

**Example:**

```js
let token = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9...";
let decoded = jwt.decode(token);

print("Header:", decoded["header"]);
print("Payload:", decoded["payload"]);
```

## Error Handling

All JWT functions return error objects when something goes wrong:

```js
let result = jwt.verify("invalid-token", "secret");
if (result.type == "ERROR") {
  print("Error:", result.message);
}
```

## Security Considerations

1. **Keep secrets secure**: Never hardcode JWT secrets in your source code
2. **Use strong secrets**: Use cryptographically strong random strings
3. **Set expiration times**: Always include expiration claims for security
4. **Validate claims**: Always verify the token payload matches your expectations

## Common Use Cases

### User Authentication

```js
// Login endpoint - create token
let payload = {
  user_id: 123,
  username: "john_doe",
  exp: time.now() + 24 * 3600, // 24 hours
};
let token = jwt.create(payload, env.JWT_SECRET);

// Protected endpoint - verify token
let result = jwt.verify(request_token, env.JWT_SECRET);
if (result.type == "ERROR") {
  return { error: "Unauthorized" };
}
let user_id = result["user_id"];
```

### API Rate Limiting

```js
// Create token with rate limit info
let payload = {
  client_id: "app123",
  requests_remaining: 1000,
  exp: time.now() + 3600, // 1 hour
};
let token = jwt.create(payload, "rate-limit-secret");
```

### Session Management

```js
// Create session token
let session = {
  session_id: uuid.generate(),
  user_id: user.id,
  permissions: ["read", "write"],
  exp: time.now() + 8 * 3600, // 8 hours
};
let session_token = jwt.createHS256(session, "session-secret", 8);
```
