# Net Module

The `net` module provides functions for making HTTP requests (GET, POST, PUT, DELETE, PATCH) from your Vint scripts. Each function supports both simple and advanced usage, allowing you to specify URLs, headers, and request bodies.

## Functions

### 1. `net.get`

**Description:**
Sends an HTTP GET request.

**Usage:**
```vint
net.get("https://api.example.com/data")
```
Or with named arguments:
```vint
net.get(
  url: "https://api.example.com/data",
  headers: {"Authorization": "Bearer token"},
  body: {"key": "value"}  # Optional, rarely used for GET
)
```

**Arguments:**
- `url` (string, required): The URL to request.
- `headers` (dict, optional): HTTP headers as key-value pairs.
- `body` (dict, optional): Data to send as JSON (rare for GET).

**Returns:**
Response body as a string, or an error object.

---

### 2. `net.post`

**Description:**
Sends an HTTP POST request.

**Usage:**
```vint
net.post(
  url: "https://api.example.com/data",
  headers: {"Authorization": "Bearer token"},
  body: {"key": "value"}
)
```

**Arguments:**
- `url` (string, required): The URL to request.
- `headers` (dict, optional): HTTP headers as key-value pairs.
- `body` (dict, optional): Data to send as JSON.

**Returns:**
Response body as a string, or an error object.

---

### 3. `net.put`

**Description:**
Sends an HTTP PUT request.

**Usage:**
```vint
net.put(
  url: "https://api.example.com/data/1",
  headers: {"Authorization": "Bearer token"},
  body: {"key": "new value"}
)
```

**Arguments:**
- `url` (string, required): The URL to request.
- `headers` (dict, optional): HTTP headers as key-value pairs.
- `body` (dict, optional): Data to send as JSON.

**Returns:**
Response body as a string, or an error object.

---

### 4. `net.delete`

**Description:**
Sends an HTTP DELETE request.

**Usage:**
```vint
net.delete(
  url: "https://api.example.com/data/1",
  headers: {"Authorization": "Bearer token"}
)
```

**Arguments:**
- `url` (string, required): The URL to request.
- `headers` (dict, optional): HTTP headers as key-value pairs.

**Returns:**
Response body as a string, or an error object.

---

### 5. `net.patch`

**Description:**
Sends an HTTP PATCH request.

**Usage:**
```vint
net.patch(
  url: "https://api.example.com/data/1",
  headers: {"Authorization": "Bearer token"},
  body: {"key": "patched value"}
)
```

**Arguments:**
- `url` (string, required): The URL to request.
- `headers` (dict, optional): HTTP headers as key-value pairs.
- `body` (dict, optional): Data to send as JSON.

**Returns:**
Response body as a string, or an error object.

---

## Notes

- All functions return the response body as a string, or an error object if something goes wrong.
- Named arguments (`url`, `headers`, `body`) are recommended for clarity.
- Headers and body must be dictionaries.
- For GET requests, the body is rarely used and may not be supported by all servers.
