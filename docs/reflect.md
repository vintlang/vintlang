# Reflect Module

The `reflect` module provides runtime reflection capabilities for VintLang, allowing you to inspect and analyze objects at runtime.

## Functions

### `typeof(obj)`
Returns the type of an object as a string.

```js
import reflect

reflect.typeof(42)          // "INTEGER"
reflect.typeof("hello")     // "STRING"
reflect.typeof([1, 2, 3])   // "ARRAY"
reflect.typeof({"a": 1})    // "DICT"
reflect.typeof(true)        // "BOOLEAN"
```

### `inspect(obj)`
Returns the inspection string representation of an object.

```js
import reflect

reflect.inspect([1, 2, 3])           // "[1, 2, 3]"
reflect.inspect({"name": "John"})    // "{name: John}"
```

### `isType(obj, typeString)`
Checks if an object is of a specific type.

```js
import reflect

reflect.isType(42, "INTEGER")        // true
reflect.isType("hello", "STRING")    // true
reflect.isType([1, 2], "ARRAY")      // true
reflect.isType({}, "DICT")           // true
```

### `size(obj)`
Returns the size/length of collections (arrays, dicts, strings).

```js
import reflect

reflect.size([1, 2, 3, 4, 5])       // 5
reflect.size("VintLang")             // 8
reflect.size({"a": 1, "b": 2})      // 2
```

### `keys(obj)`
Returns the keys of a dictionary as an array.

```js
import reflect

let person = {"name": "Alice", "age": 25}
reflect.keys(person)                 // ["age", "name"]
```

### `hasMethod(obj, methodName)`
Checks if an object has a specific method.

```js
import reflect

reflect.hasMethod([1, 2, 3], "push")     // true
reflect.hasMethod({"a": 1}, "keys")      // true
reflect.hasMethod(42, "push")            // false
```

## Usage Examples

### Basic Type Inspection
```js
import reflect

let data = [42, "hello", [1, 2], {"key": "value"}]

for item in data {
    let type = reflect.typeof(item)
    print("Item:", reflect.inspect(item), "Type:", type)
}
```

### Conditional Processing Based on Type
```js
import reflect

func processData(obj) {
    if reflect.isType(obj, "ARRAY") {
        print("Processing array of size:", reflect.size(obj))
    } else if reflect.isType(obj, "DICT") {
        print("Processing dict with keys:", reflect.keys(obj))
    } else {
        print("Processing", reflect.typeof(obj), ":", reflect.inspect(obj))
    }
}
```

### Method Capability Detection
```js
import reflect

func canPush(obj) {
    return reflect.hasMethod(obj, "push")
}

let arr = [1, 2, 3]
let dict = {"a": 1}

print("Array can push:", canPush(arr))    // true
print("Dict can push:", canPush(dict))    // false
```

The reflect module follows VintLang's error handling conventions and provides detailed error messages for incorrect usage.