# YAML Module

The YAML module provides functionality to work with YAML (YAML Ain't Markup Language) data in Vint. It allows you to parse YAML strings, convert Vint objects to YAML format, and manipulate YAML data structures.

## Functions

### yaml.decode(yamlString)

Parses a YAML string and converts it to Vint objects.

**Parameters:**

- `yamlString` (string): A valid YAML string to parse

**Returns:** Vint object (dict, array, string, number, boolean, or null)

**Example:**

```js
import yaml

let yamlData = yaml.decode("name: John\nage: 30\nactive: true")
print(yamlData) // {"name": "John", "age": 30, "active": true}
```

### yaml.encode(object)

Converts a Vint object to YAML format string.

**Parameters:**

- `object`: Any Vint object (dict, array, string, number, boolean, or null)

**Returns:** String containing YAML representation

**Example:**

```js
import yaml

let data = {
    "name": "Alice",
    "skills": ["Python", "Go", "Vint"],
    "active": true
}
let yamlString = yaml.encode(data)
print(yamlString)
```

### yaml.merge(object1, object2)

Merges two YAML-compatible objects into one. Properties from the second object will overwrite properties from the first object with the same key.

**Parameters:**

- `object1`: First object to merge
- `object2`: Second object to merge

**Returns:** New merged object

**Example:**

```js
import yaml

let obj1 = {"name": "John", "age": 30}
let obj2 = {"city": "NYC", "age": 35}
let merged = yaml.merge(obj1, obj2)
print(merged) // {"name": "John", "age": 35, "city": "NYC"}
```

### yaml.get(object, key)

Retrieves a value from a YAML-compatible object by key.

**Parameters:**

- `object`: The object to search in
- `key` (string): The key to look for

**Returns:** The value associated with the key, or null if not found

**Example:**

```js
import yaml

let data = yaml.decode("person:\n  name: Jane\n  age: 25")
let name = yaml.get(data.person, "name")
print(name) // "Jane"
```

## Supported YAML Features

The YAML module supports:

- Scalar values (strings, numbers, booleans, null)
- Sequences (arrays)
- Mappings (dictionaries/objects)
- Nested structures
- Multi-line strings

## Error Handling

All YAML functions return error objects when:

- Invalid YAML syntax is provided to `decode()`
- Incorrect argument types are passed
- Wrong number of arguments are provided

## Notes

- YAML keys are converted to strings in Vint objects
- The module handles YAML's flexible key types by converting them to strings
- Complex YAML features like anchors, references, and custom tags are not supported
- The implementation uses the gopkg.in/yaml.v3 library for robust YAML processing

## Usage with Files

You can combine the YAML module with file operations to work with YAML configuration files:

```js
import yaml
import os

// Read YAML config file
let configContent = os.read("config.yaml")
let config = yaml.decode(configContent)

// Modify and save back
config.updated = true
let newYaml = yaml.encode(config)
os.write("config.yaml", newYaml)
```
