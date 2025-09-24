
# URL Module

The `url` module provides a set of functions for working with URLs. You can use it to parse, encode, decode, join, and validate URLs.

## Functions

### `parse(urlString)`

Parses a URL string and returns its components.

- `urlString` (string): The URL to parse.

**Returns:** A string containing the URL components (scheme, host, path, query, fragment).

**Usage:**

```vint
import url

let components = url.parse("https://example.com/path?query=value#fragment")
println(components)
// Output: scheme:https host:example.com path:/path query:query=value fragment:fragment
```

### `encode(text)`

URL-encodes a string.

- `text` (string): The string to encode.

**Returns:** The URL-encoded string.

**Usage:**

```vint
import url

let encoded = url.encode("hello world!")
println(encoded) // "hello%20world%21"
```

### `decode(encodedText)`

Decodes a URL-encoded string.

- `encodedText` (string): The URL-encoded string to decode.

**Returns:** The decoded string.

**Usage:**

```vint
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

```vint
import url

let fullURL = url.join("https://example.com/", "/path/to/resource")
println(fullURL) // "https://example.com/path/to/resource"
```

### `isValid(urlString)`

Checks if a string is a valid URL.

- `urlString` (string): The URL to validate.

**Returns:** `true` if the URL is valid, `false` otherwise.

**Usage:**

```vint
import url

println(url.isValid("https://example.com")) // true
println(url.isValid("not a url")) // false
```
