# URL Module in Vint

The URL module in Vint provides comprehensive URL manipulation and validation utilities. This module helps you parse URLs, encode/decode text for URLs, join URL components, and validate URL formats.

---

## Importing the URL Module

To use the URL module, simply import it:
```js
import url
```

---

## Functions and Examples

### 1. Parse URL (`parse`)
The `parse` function breaks down a URL into its component parts (scheme, host, path, query, fragment).

**Syntax**:
```js
parse(urlString)
```

**Example**:
```js
import url

print("=== URL Parse Example ===")
test_url = "https://example.com/path/to/page?query=value&page=1#section"
parsed = url.parse(test_url)
print("Parsed URL:", parsed)
// Output: Parsed URL: scheme:https host:example.com path:/path/to/page query:query=value&page=1 fragment:section
```

---

### 2. URL Encode Text (`encode`)
The `encode` function converts text to URL-safe format by encoding special characters.

**Syntax**:
```js
encode(text)
```

**Example**:
```js
import url

print("=== URL Encode Example ===")
text_to_encode = "hello world! & special chars"
encoded = url.encode(text_to_encode)
print("Original:", text_to_encode)
print("Encoded: ", encoded)
// Output: Encoded: hello%20world%21%20%26%20special%20chars
```

---

### 3. URL Decode Text (`decode`)
The `decode` function converts URL-encoded text back to its original format.

**Syntax**:
```js
decode(encodedText)
```

**Example**:
```js
import url

print("=== URL Decode Example ===")
encoded_text = "hello%20world%21%20%26%20special%20chars"
decoded = url.decode(encoded_text)
print("Encoded: ", encoded_text)
print("Decoded: ", decoded)
// Output: Decoded: hello world! & special chars
```

---

### 4. Join URL Components (`join`)
The `join` function combines a base URL with a relative path to create a complete URL.

**Syntax**:
```js
join(baseUrl, path)
```

**Example**:
```js
import url

print("=== URL Join Example ===")
base = "https://api.example.com"
path = "/users/123/profile"
full_url = url.join(base, path)
print("Base URL:", base)
print("Path:    ", path)
print("Full URL:", full_url)
// Output: Full URL: https://api.example.com/users/123/profile
```

---

### 5. Validate URL (`isValid`)
The `isValid` function checks if a given string is a valid URL format.

**Syntax**:
```js
isValid(urlString)
```

**Example**:
```js
import url

print("=== URL Validation Example ===")

// Valid URLs
valid_urls = [
    "https://example.com",
    "http://localhost:8080/path",
    "ftp://files.example.com/download"
]

// Invalid URLs
invalid_urls = [
    "not-a-url",
    "missing-protocol.com",
    "http://",
    ""
]

print("Valid URLs:")
for valid_url in valid_urls {
    result = url.isValid(valid_url)
    print("  " + valid_url + " -> " + string(result))
}

print("\nInvalid URLs:")
for invalid_url in invalid_urls {
    result = url.isValid(invalid_url)
    print("  " + invalid_url + " -> " + string(result))
}
```

---

## Complete Usage Example

```js
import url

print("=== URL Module Complete Example ===")

// Building a search URL
base_url = "https://search.example.com"
search_path = "/search"
search_query = "VintLang programming language"

// Create the search URL
search_url = url.join(base_url, search_path)
encoded_query = url.encode(search_query)
full_search_url = search_url + "?q=" + encoded_query

print("Search URL:", full_search_url)

// Validate the URL
if url.isValid(full_search_url) {
    print("✓ URL is valid")
    
    // Parse the URL to see components
    parsed = url.parse(full_search_url)
    print("URL Components:", parsed)
} else {
    print("✗ URL is invalid")
}

// Decode a received URL parameter
received_param = "VintLang%20programming%20language"
decoded_param = url.decode(received_param)
print("Received parameter:", decoded_param)
```

---

## Use Cases

- **API Integration**: Build and validate API endpoint URLs
- **Web Scraping**: Construct URLs for web scraping tasks
- **Form Processing**: Encode form data for URL parameters
- **Link Generation**: Create dynamic links in web applications
- **URL Validation**: Verify user-provided URLs before processing

---

## Summary of Functions

| Function   | Description                                    | Return Type |
|------------|------------------------------------------------|-------------|
| `parse`    | Breaks URL into components                     | String      |
| `encode`   | Encodes text for safe URL usage               | String      |
| `decode`   | Decodes URL-encoded text                       | String      |
| `join`     | Combines base URL with relative path           | String      |
| `isValid`  | Validates URL format                           | Boolean     |

The URL module provides essential functionality for working with URLs safely and efficiently in web-related VintLang applications.