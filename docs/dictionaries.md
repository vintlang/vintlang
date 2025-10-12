Here’s a detailed explanation of dictionaries in Vint, without the Swahili terms:

### Dictionaries in Vint

In the Vint programming language, dictionaries are key-value data structures that allow you to store and manage data efficiently. These dictionaries can store any type of value (such as strings, integers, booleans, or even functions) and are incredibly useful for organizing and accessing data. 

### Creating Dictionaries

In Vint, dictionaries are created using curly braces `{}`. Each key is followed by a colon `:` and the corresponding value. Here's an example of a dictionary:

```js
dict = {"name": "John", "age": 30}
```

In this dictionary:
- `"name"` is the key, and `"John"` is the value.
- `"age"` is the key, and `30` is the value.

Keys can be of various data types like strings, integers, floats, or booleans, and values can be anything, including strings, integers, booleans, `null`, or even functions.

### Accessing Elements

You can access individual elements in a dictionary by using the key. For example:

```js
print(dict["name"]) // John
```

This will print `"John"`, the value associated with the key `"name"`.

### Updating Elements

To update the value of an existing key, simply assign a new value to the key:

```js
dict["age"] = 35
print(dict["age"]) // 35
```

This updates the `"age"` key to have the value `35`.

### Adding New Elements

To add a new key-value pair to a dictionary, assign a value to a new key:

```js
dict["city"] = "Dar es Salaam"
print(dict["city"]) // Dar es Salaam
```

This adds a new key `"city"` with the value `"Dar es Salaam"`.

### Concatenating Dictionaries

You can combine two dictionaries into one using the `+` operator:

```js
dict1 = {"a": "apple", "b": "banana"}
dict2 = {"c": "cherry", "d": "date"}
combined = dict1 + dict2
print(combined) // {"a": "apple", "b": "banana", "c": "cherry", "d": "date"}
```

In this case, `dict1` and `dict2` are merged into a new dictionary called `combined`.

### Checking If a Key Exists in a Dictionary

To check if a particular key exists in a dictionary, you can use the `in` keyword:

```js
"age" in dict // true
"salary" in dict // false
```

This checks whether the key `"age"` exists in the dictionary, which returns `true`, and checks whether the key `"salary"` exists, which returns `false`.

### Looping Over a Dictionary

You can loop over the keys and values of a dictionary using the `for` keyword:

```js
hobby = {"a": "reading", "b": "cycling", "c": "eating"}
for key, value in hobby {
    print(key, "=>", value)
}
```

This will output:

```
a => reading
b => cycling
c => eating
```

You can also loop over just the values without the keys:

```js
for value in hobby {
    print(value)
}
```

This will output:

```
reading
cycling
eating
```

## Dictionary Methods

Vint dictionaries come with several powerful built-in methods that make data manipulation easy and efficient:

### keys()

Get all keys from the dictionary as an array:

```js
contacts = {"Alice": "alice@email.com", "Bob": "bob@email.com"}
keyList = contacts.keys()
print(keyList)  // ["Alice", "Bob"]
```

### values()

Get all values from the dictionary as an array:

```js
contacts = {"Alice": "alice@email.com", "Bob": "bob@email.com"}
valueList = contacts.values()
print(valueList)  // ["alice@email.com", "bob@email.com"]
```

### size()

Get the number of key-value pairs in the dictionary:

```js
contacts = {"Alice": "alice@email.com", "Bob": "bob@email.com"}
print(contacts.size())  // 2
```

### has()

Check if a key exists in the dictionary:

```js
contacts = {"Alice": "alice@email.com", "Bob": "bob@email.com"}
print(contacts.has("Alice"))   // true
print(contacts.has("Charlie")) // false
```

### get()

Get a value by key with an optional default value:

```js
contacts = {"Alice": "alice@email.com", "Bob": "bob@email.com"}
email = contacts.get("Alice", "unknown")        // "alice@email.com"
unknownEmail = contacts.get("Charlie", "unknown") // "unknown"
print(email)        // alice@email.com
print(unknownEmail) // unknown
```

### set()

Set a key-value pair in the dictionary:

```js
contacts = {"Alice": "alice@email.com"}
contacts.set("Bob", "bob@email.com")
print(contacts)  // {"Alice": "alice@email.com", "Bob": "bob@email.com"}

// Method chaining is supported
contacts.set("Charlie", "charlie@email.com").set("Dave", "dave@email.com")
```

### remove()

Remove a key-value pair from the dictionary:

```js
contacts = {"Alice": "alice@email.com", "Bob": "bob@email.com"}
contacts.remove("Bob")
print(contacts)  // {"Alice": "alice@email.com"}
```

### clear()

Remove all key-value pairs from the dictionary:

```js
contacts = {"Alice": "alice@email.com", "Bob": "bob@email.com"}
contacts.clear()
print(contacts)  // {}
```

### merge()

Merge another dictionary into this one:

```js
contacts = {"Alice": "alice@email.com"}
newContacts = {"Bob": "bob@email.com", "Charlie": "charlie@email.com"}
contacts.merge(newContacts)
print(contacts)  // {"Alice": "alice@email.com", "Bob": "bob@email.com", "Charlie": "charlie@email.com"}
```

### copy()

Create a shallow copy of the dictionary:

```js
original = {"Alice": "alice@email.com", "Bob": "bob@email.com"}
backup = original.copy()
backup.set("Charlie", "charlie@email.com")
print(original)  // {"Alice": "alice@email.com", "Bob": "bob@email.com"}
print(backup)    // {"Alice": "alice@email.com", "Bob": "bob@email.com", "Charlie": "charlie@email.com"}
```

### filter()

Create a new dictionary with key-value pairs that pass a test function:

```js
scores = {"Alice": 85, "Bob": 92, "Charlie": 78, "Diana": 95}
highScores = scores.filter(func(key, value) { return value >= 90 })
print(highScores)  // {"Bob": 92, "Diana": 95}
```

### map()

Create a new dictionary with transformed values:

```js
prices = {"apple": 1.5, "banana": 0.8, "orange": 2.0}
discountedPrices = prices.map(func(key, value) { return value * 0.9 })
print(discountedPrices)  // {"apple": 1.35, "banana": 0.72, "orange": 1.8}
```

### forEach()

Execute a function for each key-value pair:

```js
contacts = {"Alice": "alice@email.com", "Bob": "bob@email.com"}
contacts.forEach(func(key, value) { 
    print("Name:", key, "Email:", value) 
})
// Output:
// Name: Alice Email: alice@email.com
// Name: Bob Email: bob@email.com
```

### find()

Find the first key-value pair that satisfies a test function:

```js
users = {"user1": 25, "user2": 17, "user3": 32}
adult = users.find(func(key, value) { return value >= 18 })
print(adult)  // ["user1", 25] or null if not found
```

### some()

Test whether at least one key-value pair passes the test:

```js
scores = {"Alice": 85, "Bob": 72, "Charlie": 95}
hasHighScore = scores.some(func(key, value) { return value >= 90 })
print(hasHighScore)  // true
```

### every()

Test whether all key-value pairs pass the test:

```js
scores = {"Alice": 85, "Bob": 92, "Charlie": 95}
allPassed = scores.every(func(key, value) { return value >= 80 })
print(allPassed)  // true
```

### pick()

Create a new dictionary with only specified keys:

```js
user = {"name": "Alice", "age": 25, "email": "alice@email.com", "password": "secret"}
publicInfo = user.pick("name", "age", "email")
print(publicInfo)  // {"name": "Alice", "age": 25, "email": "alice@email.com"}
```

### omit()

Create a new dictionary excluding specified keys:

```js
user = {"name": "Alice", "age": 25, "email": "alice@email.com", "password": "secret"}
safeInfo = user.omit("password")
print(safeInfo)  // {"name": "Alice", "age": 25, "email": "alice@email.com"}
```

### isEmpty()

Check if the dictionary is empty:

```js
emptyDict = {}
filledDict = {"key": "value"}
print(emptyDict.isEmpty())   // true
print(filledDict.isEmpty())  // false
```

### equals()

Check if two dictionaries are equal:

```js
dict1 = {"name": "Alice", "age": 25}
dict2 = {"name": "Alice", "age": 25}
dict3 = {"name": "Bob", "age": 30}
print(dict1.equals(dict2))   // true
print(dict1.equals(dict3))   // false
```

### entries()

Get an array of [key, value] pairs:

```js
contacts = {"Alice": "alice@email.com", "Bob": "bob@email.com"}
entryList = contacts.entries()
print(entryList)  // [["Alice", "alice@email.com"], ["Bob", "bob@email.com"]]
```

### flatten()

Flatten nested dictionaries (one level deep):

```js
nested = {
    "user": {"name": "Alice", "age": 25},
    "status": "active"
}
flattened = nested.flatten()
print(flattened)  // {"user.name": "Alice", "user.age": 25, "status": "active"}
```

### deepMerge()

Recursively merge dictionaries:

```js
dict1 = {"user": {"name": "Alice"}, "status": "active"}
dict2 = {"user": {"age": 25}, "role": "admin"}
merged = dict1.deepMerge(dict2)
print(merged)  // {"user": {"name": "Alice", "age": 25}, "status": "active", "role": "admin"}
```

## Practical Examples
```

## Advanced Dictionary Usage

Here are some practical examples of using dictionaries with their methods:

```js
// Building a user database
users = {}
users.set("john", {"name": "John Doe", "age": 30, "city": "New York"})
users.set("jane", {"name": "Jane Smith", "age": 25, "city": "Los Angeles"})

// Check if user exists
if (users.has("john")) {
    user = users.get("john")
    print("User found:", user["name"])
}

// Get all usernames
usernames = users.keys()
print("All users:", usernames)

// Create settings with defaults
settings = {"theme": "dark", "notifications": true}
getTheme = settings.get("theme", "light")           // "dark"
getLanguage = settings.get("language", "english")   // "english" (default)

// Configuration management
config = {}
config.set("database", "localhost")
      .set("port", 5432)
      .set("timeout", 30)
print("Config:", config)
```
