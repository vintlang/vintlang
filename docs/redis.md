# Redis Module

The Redis module provides comprehensive Redis database functionality for VintLang, allowing you to interact with Redis servers for caching, data storage, and more.

## Connection Management

### `redis.connect(address, [password], [db])`

Establishes a connection to a Redis server.

- `address`: Redis server address (e.g., "localhost:6379")
- `password`: Optional password for authentication
- `db`: Optional database number (default: 0)

**Returns**: Connection object

**Example**:

```js
conn = redis.connect("localhost:6379");
auth_conn = redis.connect("localhost:6379", "mypassword", 1);
```

### `redis.close(connection)`

Closes the Redis connection.

**Example**:

```js
redis.close(conn);
```

### `redis.ping(connection)`

Tests the connection to Redis.

**Returns**: "PONG" if successful

**Example**:

```js
result = redis.ping(conn); // Returns "PONG"
```

## String Operations

### `redis.set(connection, key, value)`

Sets a string value.

**Example**:

```js
redis.set(conn, "mykey", "myvalue");
```

### `redis.get(connection, key)`

Gets a string value.

**Returns**: String value or null if key doesn't exist

**Example**:

```js
value = redis.get(conn, "mykey");
```

### `redis.setex(connection, key, value, seconds)`

Sets a string value with expiration time.

**Example**:

```js
redis.setex(conn, "session:123", "userdata", 3600); // Expires in 1 hour
```

### `redis.mset(connection, key1, value1, key2, value2, ...)`

Sets multiple key-value pairs.

**Example**:

```js
redis.mset(conn, "key1", "value1", "key2", "value2");
```

### `redis.mget(connection, key1, key2, ...)`

Gets multiple values.

**Returns**: Array of values (null for non-existent keys)

**Example**:

```js
values = redis.mget(conn, "key1", "key2");
```

## Numeric Operations

### `redis.incr(connection, key)`

Increments the integer value of a key by 1.

**Example**:

```js
counter = redis.incr(conn, "visits"); // Increments and returns new value
```

### `redis.decr(connection, key)`

Decrements the integer value of a key by 1.

### `redis.incrby(connection, key, increment)`

Increments the integer value of a key by the given amount.

### `redis.decrby(connection, key, decrement)`

Decrements the integer value of a key by the given amount.

## Key Operations

### `redis.exists(connection, key)`

Checks if a key exists.

**Returns**: true/false

### `redis.del(connection, key1, [key2, ...])`

Deletes one or more keys.

**Returns**: Number of keys that were deleted

### `redis.expire(connection, key, seconds)`

Sets expiration time for a key.

**Returns**: true if timeout was set, false if key doesn't exist

### `redis.ttl(connection, key)`

Gets the time to live of a key in seconds.

**Returns**: TTL in seconds (-1 if no timeout, -2 if key doesn't exist)

### `redis.keys(connection, pattern)`

Returns all keys matching a pattern.

**Example**:

```js
keys = redis.keys(conn, "user:*"); // All keys starting with "user:"
```

## Hash Operations

### `redis.hset(connection, key, field, value)`

Sets a field in a hash.

### `redis.hget(connection, key, field)`

Gets a field from a hash.

### `redis.hgetall(connection, key)`

Gets all fields and values from a hash.

**Returns**: Dictionary with field-value pairs

### `redis.hdel(connection, key, field1, [field2, ...])`

Deletes one or more hash fields.

### `redis.hexists(connection, key, field)`

Determines if a hash field exists.

### `redis.hkeys(connection, key)`

Gets all field names in a hash.

### `redis.hvals(connection, key)`

Gets all values in a hash.

## List Operations

### `redis.lpush(connection, key, value1, [value2, ...])`

Prepends one or more values to a list.

### `redis.rpush(connection, key, value1, [value2, ...])`

Appends one or more values to a list.

### `redis.lpop(connection, key)`

Removes and returns the first element of a list.

### `redis.rpop(connection, key)`

Removes and returns the last element of a list.

### `redis.llen(connection, key)`

Returns the length of a list.

### `redis.lrange(connection, key, start, stop)`

Returns a range of elements from a list.

**Example**:

```js
elements = redis.lrange(conn, "mylist", 0, -1); // All elements
```

## Set Operations

### `redis.sadd(connection, key, member1, [member2, ...])`

Adds one or more members to a set.

### `redis.srem(connection, key, member1, [member2, ...])`

Removes one or more members from a set.

### `redis.smembers(connection, key)`

Returns all members of a set.

### `redis.scard(connection, key)`

Returns the number of members in a set.

### `redis.sismember(connection, key, member)`

Determines if a member is in a set.

## Sorted Set Operations

### `redis.zadd(connection, key, score1, member1, [score2, member2, ...])`

Adds one or more members to a sorted set with scores.

**Example**:

```js
redis.zadd(conn, "leaderboard", 100, "player1", 85, "player2");
```

### `redis.zrem(connection, key, member1, [member2, ...])`

Removes one or more members from a sorted set.

### `redis.zrange(connection, key, start, stop)`

Returns a range of members in a sorted set by index.

### `redis.zcard(connection, key)`

Returns the number of members in a sorted set.

### `redis.zscore(connection, key, member)`

Returns the score of a member in a sorted set.

## Usage Example

```js
// Connect to Redis
conn = redis.connect("localhost:6379");

// Basic string operations
redis.set(conn, "greeting", "Hello, World!");
message = redis.get(conn, "greeting");
print(message); // Output: Hello, World!

// Working with hashes
redis.hset(conn, "user:1", "name", "John Doe");
redis.hset(conn, "user:1", "email", "john@example.com");
user = redis.hgetall(conn, "user:1");
print(user); // Output: {"name": "John Doe", "email": "john@example.com"}

// Working with lists
redis.rpush(conn, "tasks", "task1", "task2", "task3");
task = redis.lpop(conn, "tasks");
print(task); // Output: task1

// Close connection
redis.close(conn);
```
