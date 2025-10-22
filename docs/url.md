# URL Module

The `url` module provides a set of functions for working with URLs. You can use it to parse, encode, decode, join, build, and validate URLs.

## Functions

### `parse(urlString)`

Parses a URL string and returns its components.

- `urlString` (string): The URL to parse.

**Returns:** A string containing the URL components (scheme, host, path, query, fragment).

**Usage:**

```js
import url

let components = url.parse("https://example.com/path?query=value#fragment")
println(components)
// Output: scheme:https host:example.com path:/path query:query=value fragment:fragment
```

### `build(components)`

Builds a URL from a dictionary of components.

- `components` (dict): A dictionary containing URL components.

**Valid components:**

- `scheme` (string): The URL scheme (e.g., "https", "http", "ftp")
- `host` (string): The hostname (e.g., "example.com", "localhost")
- `path` (string): The path component (e.g., "/api/v1")
- `query` (string): The query string (e.g., "limit=10&offset=0")
- `fragment` (string): The fragment identifier (e.g., "section1")
- `port` (string): The port number (e.g., "8080")
- `user` (string): User information for the URL

**Returns:** The constructed URL string.

**Usage:**

```js
import url

let components = {"scheme": "https", "host": "api.example.com", "path": "/v1/users", "query": "limit=10"}
let built_url = url.build(components)
println(built_url) // "https://api.example.com/v1/users?limit=10"

// Minimal example
let minimal = {"scheme": "http", "host": "localhost"}
println(url.build(minimal)) // "http://localhost"
```

### `encode(text)`

URL-encodes a string.

- `text` (string): The string to encode.

**Returns:** The URL-encoded string.

**Usage:**

```js
import url

let encoded = url.encode("hello world!")
println(encoded) // "hello%20world%21"
```

### `decode(encodedText)`

Decodes a URL-encoded string.

- `encodedText` (string): The URL-encoded string to decode.

**Returns:** The decoded string.

**Usage:**

```js
import url

let decoded = url.decode("hello%20world%21")
println(decoded) // "hello world!"
```

### `join(baseURL, path)`

Joins a base URL and a relative path to create a full URL.

- `baseURL` (string): The base URL.
- `path` (string): The relative path to join.

**Returns:** The full URL.

**Usage:**

```js
import url

let fullURL = url.join("https://example.com/", "/path/to/resource")
println(fullURL) // "https://example.com/path/to/resource"
```

### `isValid(urlString)`

Checks if a string is a valid URL.

- `urlString` (string): The URL to validate.

**Returns:** `true` if the URL is valid, `false` otherwise.

**Usage:**

```js
import url

println(url.isValid("https://example.com")) // true
println(url.isValid("not a url")) // false
```
