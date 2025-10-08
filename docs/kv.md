# KV Module

The **KV (Key-Value)** module provides a high-performance, thread-safe, in-memory key-value store for VintLang applications. It offers enterprise-grade features including TTL (Time-To-Live), atomic operations, bulk operations, and comprehensive statistics.

## Import

```vint
import kv
```

## Features

- 🔒 **Thread-Safe**: Concurrent read/write operations with proper locking
- ⏰ **TTL Support**: Automatic key expiration with customizable time-to-live
- 🚀 **Atomic Operations**: Increment/decrement operations for counters
- 📦 **Bulk Operations**: Efficient multi-key get/set operations
- 📊 **Statistics**: Real-time metrics and store monitoring
- 💾 **Data Dump**: Export all data for debugging and backup
- 🧹 **Memory Management**: Automatic cleanup of expired keys

## Basic Operations

### set(key, value)

Sets a key-value pair in the store.

**Parameters:**
- `key` (string): The key to set
- `value` (any): The value to store

**Returns:** `boolean` - `true` if successful

**Example:**
```vint
import kv

kv.set("user:123", {"name": "John", "age": 30})
kv.set("session:456", "active")
kv.set("counter", 42)
```

### get(key)

Retrieves a value by its key.

**Parameters:**
- `key` (string): The key to retrieve

**Returns:** The stored value, or `null` if not found or expired

**Example:**
```vint
let user = kv.get("user:123")
println("User:", user) // {name: John, age: 30}

let missing = kv.get("nonexistent")
println("Missing:", missing) // null
```

### delete(key)

Removes a key-value pair from the store.

**Parameters:**
- `key` (string): The key to delete

**Returns:** `boolean` - `true` if key existed and was deleted

**Example:**
```vint
let deleted = kv.delete("session:456")
println("Deleted:", deleted) // true
```

### exists(key)

Checks if a key exists in the store (and is not expired).

**Parameters:**
- `key` (string): The key to check

**Returns:** `boolean` - `true` if key exists and is not expired

**Example:**
```vint
if (kv.exists("user:123")) {
    println("User exists")
}
```

### clear()

Removes all key-value pairs from the store.

**Returns:** `boolean` - Always `true`

**Example:**
```vint
kv.clear()
println("Store cleared")
```

## Store Information

### keys()

Returns an array of all keys in the store (excluding expired keys).

**Returns:** `array` - Array of string keys

**Example:**
```vint
kv.set("key1", "value1")
kv.set("key2", "value2")
let allKeys = kv.keys()
println("Keys:", allKeys) // ["key1", "key2"]
```

### values()

Returns an array of all values in the store (excluding expired values).

**Returns:** `array` - Array of stored values

**Example:**
```vint
let allValues = kv.values()
println("Values:", allValues) // ["value1", "value2"]
```

### size()

Returns the number of key-value pairs in the store.

**Returns:** `integer` - Number of stored pairs

**Example:**
```vint
println("Store size:", kv.size()) // 2
```

### isEmpty()

Checks if the store is empty.

**Returns:** `boolean` - `true` if store has no keys

**Example:**
```vint
if (kv.isEmpty()) {
    println("Store is empty")
}
```

## TTL (Time-To-Live) Operations

### setTTL(key, value, ttl_seconds)

Sets a key-value pair with automatic expiration.

**Parameters:**
- `key` (string): The key to set
- `value` (any): The value to store
- `ttl_seconds` (integer): Time-to-live in seconds

**Returns:** `boolean` - `true` if successful

**Example:**
```vint
// Set a session that expires in 5 minutes
kv.setTTL("session:temp", "temporary_data", 300)

// Set a cache entry that expires in 1 hour
kv.setTTL("cache:user:123", userData, 3600)
```

### getTTL(key)

Gets the remaining time-to-live for a key.

**Parameters:**
- `key` (string): The key to check

**Returns:** `integer` - Remaining seconds, or `-1` if no TTL set, or `null` if key doesn't exist

**Example:**
```vint
let remaining = kv.getTTL("session:temp")
if (remaining != null && remaining > 0) {
    println("Session expires in", remaining, "seconds")
}
```

### expire(key, ttl_seconds)

Sets or updates the TTL for an existing key.

**Parameters:**
- `key` (string): The key to set expiration for
- `ttl_seconds` (integer): Time-to-live in seconds (must be positive)

**Returns:** `boolean` - `true` if key exists and TTL was set

**Example:**
```vint
kv.set("temp:data", "some data")
kv.expire("temp:data", 60) // Expire in 1 minute
```

## Bulk Operations

### mget(keys)

Gets multiple values in a single operation.

**Parameters:**
- `keys` (array): Array of string keys to retrieve

**Returns:** `array` - Array of values in same order as keys (`null` for missing/expired keys)

**Example:**
```vint
kv.set("user:1", "Alice")
kv.set("user:2", "Bob")

let users = kv.mget(["user:1", "user:2", "user:3"])
println("Users:", users) // ["Alice", "Bob", null]
```

### mset(pairs)

Sets multiple key-value pairs in a single operation.

**Parameters:**
- `pairs` (dictionary): Dictionary of key-value pairs to set

**Returns:** `boolean` - `true` if all pairs were set successfully

**Example:**
```vint
let bulkData = {
    "config:theme": "dark",
    "config:language": "en",
    "config:notifications": true
}
kv.mset(bulkData)
```

## Atomic Operations

### increment(key, [delta])

Atomically increments a numeric value.

**Parameters:**
- `key` (string): The key to increment
- `delta` (integer, optional): Amount to increment by (default: 1)

**Returns:** `integer` - The new value after increment

**Notes:**
- If key doesn't exist, creates it with the delta value
- If key exists but value is not an integer, returns an error
- Thread-safe for concurrent increments

**Example:**
```vint
// Simple counter
let count = kv.increment("page:views")
println("Views:", count) // 1

// Increment by custom amount
let score = kv.increment("user:score", 10)
println("Score:", score) // 10

// Increment existing value
kv.set("counter", 5)
let newCount = kv.increment("counter", 3)
println("New count:", newCount) // 8
```

### decrement(key, [delta])

Atomically decrements a numeric value.

**Parameters:**
- `key` (string): The key to decrement
- `delta` (integer, optional): Amount to decrement by (default: 1)

**Returns:** `integer` - The new value after decrement

**Notes:**
- If key doesn't exist, creates it with the negative delta value
- If key exists but value is not an integer, returns an error
- Thread-safe for concurrent decrements

**Example:**
```vint
// Simple countdown
let remaining = kv.decrement("lives")
println("Lives remaining:", remaining) // -1

// Decrement by custom amount
kv.set("inventory", 100)
let newInventory = kv.decrement("inventory", 15)
println("Inventory:", newInventory) // 85
```

## Utility Functions

### dump()

Returns all key-value pairs in the store as a dictionary (excluding expired keys).

**Returns:** `dictionary` - All stored key-value pairs

**Example:**
```vint
kv.set("key1", "value1")
kv.set("key2", 42)
let allData = kv.dump()
println("All data:", allData) // {key1: value1, key2: 42}
```

### stats()

Returns statistics about the KV store.

**Returns:** `dictionary` - Statistics including:
- `total_keys`: Total number of keys (including expired)
- `active_keys`: Number of active (non-expired) keys
- `expired_keys`: Number of expired keys
- `keys_with_ttl`: Number of keys that have TTL set

**Example:**
```vint
let statistics = kv.stats()
println("Store stats:", statistics)
// {total_keys: 10, active_keys: 8, expired_keys: 2, keys_with_ttl: 5}
```

## Common Use Cases

### Session Management

```vint
import kv

// Store user session with 30-minute expiration
func createSession(userId, sessionData) {
    let sessionId = "session:" + userId
    kv.setTTL(sessionId, sessionData, 1800) // 30 minutes
    return sessionId
}

// Check if session is valid
func isSessionValid(sessionId) {
    return kv.exists(sessionId)
}

// Extend session
func extendSession(sessionId) {
    if (kv.exists(sessionId)) {
        kv.expire(sessionId, 1800) // Extend by 30 minutes
        return true
    }
    return false
}
```

### Caching

```vint
import kv

// Cache expensive computation results
func getCachedResult(cacheKey, computeFunc) {
    // Check cache first
    let cached = kv.get(cacheKey)
    if (cached != null) {
        return cached
    }
    
    // Compute and cache result
    let result = computeFunc()
    kv.setTTL(cacheKey, result, 300) // Cache for 5 minutes
    return result
}
```

### Rate Limiting

```vint
import kv

// Simple rate limiter
func isRateLimited(userId, limit, windowSeconds) {
    let key = "rate:" + userId
    let current = kv.get(key)
    
    if (current == null) {
        // First request in window
        kv.setTTL(key, 1, windowSeconds)
        return false
    }
    
    if (current >= limit) {
        return true // Rate limited
    }
    
    // Increment counter
    kv.increment(key)
    return false
}

// Usage: Allow 100 requests per minute
if (isRateLimited("user123", 100, 60)) {
    println("Rate limited!")
}
```

### Counters and Metrics

```vint
import kv

// Track page views
func trackPageView(page) {
    kv.increment("views:" + page)
    kv.increment("total:views")
}

// Track user actions
func trackUserAction(userId, action) {
    kv.increment("user:" + userId + ":actions")
    kv.increment("action:" + action + ":count")
}

// Get metrics
func getMetrics() {
    return {
        "total_views": kv.get("total:views"),
        "stats": kv.stats(),
        "all_counters": kv.dump()
    }
}
```

### Configuration Management

```vint
import kv

// Load configuration
func loadConfig() {
    let defaultConfig = {
        "app:theme": "light",
        "app:language": "en",
        "app:debug": false,
        "cache:ttl": 3600
    }
    kv.mset(defaultConfig)
}

// Update configuration
func updateConfig(key, value) {
    kv.set("config:" + key, value)
}

// Get all configuration
func getConfig() {
    let allData = kv.dump()
    let config = {}
    
    for (key, value in allData) {
        if (key.startsWith("config:")) {
            config[key] = value
        }
    }
    
    return config
}
```

## Performance Considerations

### Thread Safety
- All operations are thread-safe using read-write mutexes
- Multiple concurrent reads are allowed
- Writes are exclusive and properly synchronized

### Memory Management
- Expired keys are automatically cleaned up when accessed
- Use `clear()` to free all memory when done
- Monitor using `stats()` to track memory usage

### Bulk Operations
- Use `mget()` and `mset()` for better performance with multiple keys
- Bulk operations are more efficient than multiple single operations

### TTL Best Practices
- Set appropriate TTL values to prevent memory leaks
- Use `getTTL()` to check remaining time before operations
- Consider using `expire()` to extend TTL for active sessions

## Error Handling

All KV functions return appropriate error messages for invalid usage:

```vint
// Invalid key type
let result = kv.get(123) // Error: key must be a string

// Invalid TTL
let result = kv.expire("key", -5) // Error: TTL must be positive

// Invalid increment target
kv.set("text", "hello")
let result = kv.increment("text") // Error: existing value must be an integer
```

## Integration Examples

### With HTTP Server

```vint
import kv
import http

// Simple API with caching
let app = http.app()

app.get("/api/user/:id", func(req, res) {
    let userId = req.params.id
    let cacheKey = "user:" + userId
    
    // Try cache first
    let user = kv.get(cacheKey)
    if (user != null) {
        return res.json({"user": user, "cached": true})
    }
    
    // Fetch from database (simulated)
    user = fetchUserFromDB(userId)
    
    // Cache for 10 minutes
    kv.setTTL(cacheKey, user, 600)
    
    res.json({"user": user, "cached": false})
})

// Track API calls
app.use(func(req, res, next) {
    kv.increment("api:calls")
    kv.increment("api:endpoint:" + req.path)
    next()
})
```

### With Async Operations

```vint
import kv

// Async cache implementation
async func getCachedOrFetch(key, fetchFunc) {
    let cached = kv.get(key)
    if (cached != null) {
        return cached
    }
    
    let result = await fetchFunc()
    kv.setTTL(key, result, 300)
    return result
}
```

The KV module provides a robust foundation for in-memory data storage and caching in VintLang applications, with enterprise-grade features suitable for production use.