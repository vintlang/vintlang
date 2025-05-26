# JSON Module in Vint

The JSON module in Vint provides powerful and straightforward functions for working with JSON data, including decoding, encoding, formatting, merging, and retrieving values. Below is the detailed documentation, along with examples.

---

## Importing the JSON Module

To use the JSON module, simply import it:
```js
import json
```

---

## Functions and Examples

### 1. Decode JSON (`decode`)
The `decode` function parses a JSON string into a Vint dictionary or array.

**Syntax**:
```js
decode(jsonString)
```

**Example**:
```js
import json

print("=== Example 1: Decode ===")
raw_json = '{"name": "John", "age": 30, "isAdmin": false, "friends": ["Jane", "Doe"]}'
decoded = json.decode(raw_json)
print("Decoded Object:", decoded)
// Output: Decoded Object: {"name": "John", "age": 30, "isAdmin": false, "friends": ["Jane", "Doe"]}
```

---

### 2. Encode JSON (`encode`)
The `encode` function converts a Vint dictionary or array into a JSON string. It optionally supports pretty formatting with an `indent` parameter.

**Syntax**:
```js
encode(data, indent = 0)
```

**Example**:
```js
import json

print("\n=== Example 2: Encode ===")
data = {
  "language": "Vint",
  "version": 1.0,
  "features": ["custom modules", "native objects"]
}
encoded_json = json.encode(data, indent=2)
print("Encoded JSON:", encoded_json)
// Output:
// Encoded JSON: {
//   "language": "Vint",
//   "version": 1.0,
//   "features": ["custom modules", "native objects"]
// }
```

---

### 3. Pretty Print JSON (`pretty`)
The `pretty` function reformats a JSON string into a human-readable format with proper indentation.

**Syntax**:
```js
pretty(jsonString)
```

**Example**:
```js
import json

print("\n=== Example 3: Pretty Print ===")
raw_json_pretty = '{"name":"John","age":30,"friends":["Jane","Doe"]}'
pretty_json = json.pretty(raw_json_pretty)
print("Pretty JSON:\n", pretty_json)
// Output:
// Pretty JSON:
// {
//   "name": "John",
//   "age": 30,
//   "friends": ["Jane", "Doe"]
// }
```

---

### 4. Merge JSON Objects (`merge`)
The `merge` function combines two JSON objects. If both objects have the same key, the value from the second object overwrites the first.

**Syntax**:
```js
merge(json1, json2)
```

**Example**:
```js
import json

print("\n=== Example 4: Merge ===")
json1 = {"name": "John", "age": 30}
json2 = {"city": "New York", "age": 35}
merged_json = json.merge(json1, json2)
print("Merged JSON:", merged_json)
// Output: Merged JSON: {"name": "John", "age": 35, "city": "New York"}
```

---

### 5. Get Value by Key (`get`)
The `get` function retrieves a value associated with a key from a JSON object. If the key is not found, it returns `null`.

**Syntax**:
```js
get(jsonObject, key)
```

**Example**:
```js
import json

print("\n=== Example 5: Get Value by Key ===")
json_object = {"name": "John", "age": 30, "city": "New York"}

value = json.get(json_object, "age")
print("Age:", value)
// Output: Age: 30

missing_value = json.get(json_object, "country")
print("Country (missing key):", missing_value)
// Output: Country (missing key): null
```

---

## Summary of Functions

| Function         | Description                                         | Example Output                           |
|------------------|-----------------------------------------------------|------------------------------------------|
| `decode`         | Converts JSON string to a Vint object.             | `{"key": "value"}`                       |
| `encode`         | Converts Vint object to a JSON string.             | `{"key":"value"}`                        |
| `pretty`         | Formats JSON string for better readability.        | `{ "key": "value" }`                     |
| `merge`          | Combines two JSON objects, overwriting duplicates. | `{"key1": "value1", "key2": "value2"}`   |
| `get`            | Retrieves a value by key, returns `null` if absent.| `"value"` or `null`                      |

These functions make working with JSON in Vint easy, flexible, and efficient.