# Vint NET Module Documentation (`net`)

The `net` module in Vint provides functionalities for performing HTTP requests. This module allows you to make various types of HTTP requests such as GET, POST, PUT, DELETE, and PATCH. It supports adding headers and request bodies, and the responses from the HTTP requests are returned as strings.

## Importing the `net` Module

To use the functionality provided by the `net` module in your Vint program, you must first import it:

```vint
import net
```

## Methods Overview

The following HTTP methods are available in the `net` module:

- **`net.get()`**: Perform a GET request.
- **`net.post()`**: Perform a POST request.
- **`net.put()`**: Perform a PUT request.
- **`net.delete()`**: Perform a DELETE request.
- **`net.patch()`**: Perform a PATCH request.

---

## Method Details

### `net.get()`

Performs an HTTP GET request. The method accepts a URL and optional headers or body data.

#### Usage

```vint
response = net.get("http://example.com")
```

Or with keyword arguments:

```vint
response = net.get(url="http://mysite.com", headers={"Authorization": "Bearer token"}, body={"key": "value"})
```

#### Parameters
- **`url`** *(string)*: The URL for the GET request.
- **`headers`** *(dictionary, optional)*: Headers to include in the request.
- **`body`** *(dictionary, optional)*: Data to send in the request body.

#### Returns
A `string` containing the response body.

---

### `net.post()`

Performs an HTTP POST request. This is useful for sending data to a server.

#### Usage

```vint
response = net.post("http://example.com", headers={"Content-Type": "application/json"}, body={"key": "value"})
```

#### Parameters
- **`url`** *(string)*: The URL for the POST request.
- **`headers`** *(dictionary, optional)*: Headers to include in the request.
- **`body`** *(dictionary, optional)*: Data to send in the request body.

#### Returns
A `string` containing the response body.

---

### `net.put()`

Performs an HTTP PUT request to update resources on a server.

#### Usage

```vint
response = net.put("http://example.com/resource", headers={"Authorization": "Bearer token"}, body={"updated_key": "new_value"})
```

#### Parameters
- **`url`** *(string)*: The URL for the PUT request.
- **`headers`** *(dictionary, optional)*: Headers to include in the request.
- **`body`** *(dictionary, optional)*: Data to send in the request body.

#### Returns
A `string` containing the response body.

---

### `net.delete()`

Performs an HTTP DELETE request to remove a resource from the server.

#### Usage

```vint
response = net.delete("http://example.com/resource", headers={"Authorization": "Bearer token"})
```

#### Parameters
- **`url`** *(string)*: The URL for the DELETE request.
- **`headers`** *(dictionary, optional)*: Headers to include in the request.
- **`body`** *(dictionary, optional)*: Data to send in the request body, if supported by the API.

#### Returns
A `string` containing the response body.

---

### `net.patch()`

Performs an HTTP PATCH request to partially update a resource on the server.

#### Usage

```vint
response = net.patch("http://example.com/resource", headers={"Authorization": "Bearer token"}, body={"key": "updated_value"})
```

#### Parameters
- **`url`** *(string)*: The URL for the PATCH request.
- **`headers`** *(dictionary, optional)*: Headers to include in the request.
- **`body`** *(dictionary, optional)*: Data to send in the request body.

#### Returns
A `string` containing the response body.

---


## Example Usage

### Basic GET Request

```vint
import net

response = net.get("http://example.com")
print(response)
```

### POST Request with Headers and Body

```vint
import net

url = "http://example.com/api"
headers = {"Authorization": "Bearer token"}
data = {"key": "value"}

response = net.post(url=url, headers=headers, body=data)
print(response)
```

---

## Return Type

All methods return a `string` containing the response body. In case of an error, an error object is returned with the error details.

---
