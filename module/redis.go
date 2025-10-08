package module

import (
	"context"
	"fmt"
	"time"

	"github.com/redis/go-redis/v9"
	"github.com/vintlang/vintlang/object"
)

var RedisFunctions = map[string]object.ModuleFunction{}

func init() {
	// Connection management
	RedisFunctions["connect"] = redisConnect
	RedisFunctions["close"] = redisClose
	RedisFunctions["ping"] = redisPing

	// String operations
	RedisFunctions["set"] = redisSet
	RedisFunctions["get"] = redisGet
	RedisFunctions["setex"] = redisSetEx
	RedisFunctions["mset"] = redisMSet
	RedisFunctions["mget"] = redisMGet
	RedisFunctions["incr"] = redisIncr
	RedisFunctions["decr"] = redisDecr
	RedisFunctions["incrby"] = redisIncrBy
	RedisFunctions["decrby"] = redisDecrBy

	// Key operations
	RedisFunctions["exists"] = redisExists
	RedisFunctions["del"] = redisDel
	RedisFunctions["expire"] = redisExpire
	RedisFunctions["ttl"] = redisTTL
	RedisFunctions["keys"] = redisKeys

	// Hash operations
	RedisFunctions["hset"] = redisHSet
	RedisFunctions["hget"] = redisHGet
	RedisFunctions["hgetall"] = redisHGetAll
	RedisFunctions["hdel"] = redisHDel
	RedisFunctions["hexists"] = redisHExists
	RedisFunctions["hkeys"] = redisHKeys
	RedisFunctions["hvals"] = redisHVals

	// List operations
	RedisFunctions["lpush"] = redisLPush
	RedisFunctions["rpush"] = redisRPush
	RedisFunctions["lpop"] = redisLPop
	RedisFunctions["rpop"] = redisRPop
	RedisFunctions["llen"] = redisLLen
	RedisFunctions["lrange"] = redisLRange

	// Set operations
	RedisFunctions["sadd"] = redisSAdd
	RedisFunctions["srem"] = redisSRem
	RedisFunctions["smembers"] = redisSMembers
	RedisFunctions["scard"] = redisSCard
	RedisFunctions["sismember"] = redisSIsMember

	// Sorted set operations
	RedisFunctions["zadd"] = redisZAdd
	RedisFunctions["zrem"] = redisZRem
	RedisFunctions["zrange"] = redisZRange
	RedisFunctions["zcard"] = redisZCard
	RedisFunctions["zscore"] = redisZScore
}

type RedisConnection struct {
	client *redis.Client
	ctx    context.Context
}

// redisConnect establishes a connection to Redis
func redisConnect(args []object.VintObject, defs map[string]object.VintObject) object.VintObject {
	if len(args) < 1 || len(args) > 3 {
		return ErrorMessage(
			"redis",
			"connect",
			"1-3 arguments (address, [password], [db])",
			formatArgs(args),
			`redis.connect("localhost:6379") or redis.connect("localhost:6379", "password", 0)`,
		)
	}

	if args[0].Type() != object.STRING_OBJ {
		return ErrorMessage(
			"redis",
			"connect",
			"string address as first argument",
			formatArgs(args),
			`redis.connect("localhost:6379")`,
		)
	}

	address := args[0].(*object.String).Value
	password := ""
	db := 0

	// Parse optional password
	if len(args) > 1 {
		if args[1].Type() != object.STRING_OBJ {
			return ErrorMessage(
				"redis",
				"connect",
				"string password as second argument",
				formatArgs(args),
				`redis.connect("localhost:6379", "password")`,
			)
		}
		password = args[1].(*object.String).Value
	}

	// Parse optional database number
	if len(args) > 2 {
		if args[2].Type() != object.INTEGER_OBJ {
			return ErrorMessage(
				"redis",
				"connect",
				"integer db number as third argument",
				formatArgs(args),
				`redis.connect("localhost:6379", "password", 0)`,
			)
		}
		db = int(args[2].(*object.Integer).Value)
	}

	rdb := redis.NewClient(&redis.Options{
		Addr:     address,
		Password: password,
		DB:       db,
	})

	ctx := context.Background()

	// Test the connection
	_, err := rdb.Ping(ctx).Result()
	if err != nil {
		return &object.Error{Message: fmt.Sprintf("Failed to connect to Redis: %s", err)}
	}

	return &object.NativeObject{Value: &RedisConnection{client: rdb, ctx: ctx}}
}

// redisClose closes the Redis connection
func redisClose(args []object.VintObject, defs map[string]object.VintObject) object.VintObject {
	if len(args) != 1 || args[0].Type() != object.NATIVE_OBJ {
		return ErrorMessage(
			"redis",
			"close",
			"1 Redis connection",
			formatArgs(args),
			`redis.close(conn)`,
		)
	}

	conn, ok := args[0].(*object.NativeObject).Value.(*RedisConnection)
	if !ok {
		return &object.Error{Message: "Invalid Redis connection"}
	}

	err := conn.client.Close()
	if err != nil {
		return &object.Error{Message: fmt.Sprintf("Failed to close Redis connection: %s", err)}
	}

	return &object.Null{}
}

// redisPing tests the connection to Redis
func redisPing(args []object.VintObject, defs map[string]object.VintObject) object.VintObject {
	if len(args) != 1 || args[0].Type() != object.NATIVE_OBJ {
		return ErrorMessage(
			"redis",
			"ping",
			"1 Redis connection",
			formatArgs(args),
			`redis.ping(conn) -> "PONG"`,
		)
	}

	conn, ok := args[0].(*object.NativeObject).Value.(*RedisConnection)
	if !ok {
		return &object.Error{Message: "Invalid Redis connection"}
	}

	result, err := conn.client.Ping(conn.ctx).Result()
	if err != nil {
		return &object.Error{Message: fmt.Sprintf("Redis ping failed: %s", err)}
	}

	return &object.String{Value: result}
}

// redisSet sets a string value
func redisSet(args []object.VintObject, defs map[string]object.VintObject) object.VintObject {
	if len(args) != 3 || args[0].Type() != object.NATIVE_OBJ || args[1].Type() != object.STRING_OBJ || args[2].Type() != object.STRING_OBJ {
		return ErrorMessage(
			"redis",
			"set",
			"connection, key (string), value (string)",
			formatArgs(args),
			`redis.set(conn, "key", "value") -> "OK"`,
		)
	}

	conn, ok := args[0].(*object.NativeObject).Value.(*RedisConnection)
	if !ok {
		return &object.Error{Message: "Invalid Redis connection"}
	}

	key := args[1].(*object.String).Value
	value := args[2].(*object.String).Value

	result, err := conn.client.Set(conn.ctx, key, value, 0).Result()
	if err != nil {
		return &object.Error{Message: fmt.Sprintf("Redis SET failed: %s", err)}
	}

	return &object.String{Value: result}
}

// redisGet gets a string value
func redisGet(args []object.VintObject, defs map[string]object.VintObject) object.VintObject {
	if len(args) != 2 || args[0].Type() != object.NATIVE_OBJ || args[1].Type() != object.STRING_OBJ {
		return ErrorMessage(
			"redis",
			"get",
			"connection, key (string)",
			formatArgs(args),
			`redis.get(conn, "key") -> "value"`,
		)
	}

	conn, ok := args[0].(*object.NativeObject).Value.(*RedisConnection)
	if !ok {
		return &object.Error{Message: "Invalid Redis connection"}
	}

	key := args[1].(*object.String).Value

	result, err := conn.client.Get(conn.ctx, key).Result()
	if err == redis.Nil {
		return &object.Null{}
	}
	if err != nil {
		return &object.Error{Message: fmt.Sprintf("Redis GET failed: %s", err)}
	}

	return &object.String{Value: result}
}

// redisSetEx sets a string value with expiration
func redisSetEx(args []object.VintObject, defs map[string]object.VintObject) object.VintObject {
	if len(args) != 4 || args[0].Type() != object.NATIVE_OBJ || args[1].Type() != object.STRING_OBJ || args[2].Type() != object.STRING_OBJ || args[3].Type() != object.INTEGER_OBJ {
		return ErrorMessage(
			"redis",
			"setex",
			"connection, key (string), value (string), seconds (integer)",
			formatArgs(args),
			`redis.setex(conn, "key", "value", 60) -> "OK"`,
		)
	}

	conn, ok := args[0].(*object.NativeObject).Value.(*RedisConnection)
	if !ok {
		return &object.Error{Message: "Invalid Redis connection"}
	}

	key := args[1].(*object.String).Value
	value := args[2].(*object.String).Value
	seconds := time.Duration(args[3].(*object.Integer).Value) * time.Second

	result, err := conn.client.Set(conn.ctx, key, value, seconds).Result()
	if err != nil {
		return &object.Error{Message: fmt.Sprintf("Redis SETEX failed: %s", err)}
	}

	return &object.String{Value: result}
}

// redisExists checks if a key exists
func redisExists(args []object.VintObject, defs map[string]object.VintObject) object.VintObject {
	if len(args) != 2 || args[0].Type() != object.NATIVE_OBJ || args[1].Type() != object.STRING_OBJ {
		return ErrorMessage(
			"redis",
			"exists",
			"connection, key (string)",
			formatArgs(args),
			`redis.exists(conn, "key") -> true/false`,
		)
	}

	conn, ok := args[0].(*object.NativeObject).Value.(*RedisConnection)
	if !ok {
		return &object.Error{Message: "Invalid Redis connection"}
	}

	key := args[1].(*object.String).Value

	result, err := conn.client.Exists(conn.ctx, key).Result()
	if err != nil {
		return &object.Error{Message: fmt.Sprintf("Redis EXISTS failed: %s", err)}
	}

	return &object.Boolean{Value: result > 0}
}

// redisDel deletes one or more keys
func redisDel(args []object.VintObject, defs map[string]object.VintObject) object.VintObject {
	if len(args) < 2 || args[0].Type() != object.NATIVE_OBJ {
		return ErrorMessage(
			"redis",
			"del",
			"connection, key1 (string), [key2, ...]",
			formatArgs(args),
			`redis.del(conn, "key1", "key2") -> 2`,
		)
	}

	conn, ok := args[0].(*object.NativeObject).Value.(*RedisConnection)
	if !ok {
		return &object.Error{Message: "Invalid Redis connection"}
	}

	keys := make([]string, len(args)-1)
	for i, arg := range args[1:] {
		if arg.Type() != object.STRING_OBJ {
			return ErrorMessage(
				"redis",
				"del",
				"all keys must be strings",
				formatArgs(args),
				`redis.del(conn, "key1", "key2")`,
			)
		}
		keys[i] = arg.(*object.String).Value
	}

	result, err := conn.client.Del(conn.ctx, keys...).Result()
	if err != nil {
		return &object.Error{Message: fmt.Sprintf("Redis DEL failed: %s", err)}
	}

	return &object.Integer{Value: result}
}

// redisExpire sets expiration on a key
func redisExpire(args []object.VintObject, defs map[string]object.VintObject) object.VintObject {
	if len(args) != 3 || args[0].Type() != object.NATIVE_OBJ || args[1].Type() != object.STRING_OBJ || args[2].Type() != object.INTEGER_OBJ {
		return ErrorMessage(
			"redis",
			"expire",
			"connection, key (string), seconds (integer)",
			formatArgs(args),
			`redis.expire(conn, "key", 60) -> true/false`,
		)
	}

	conn, ok := args[0].(*object.NativeObject).Value.(*RedisConnection)
	if !ok {
		return &object.Error{Message: "Invalid Redis connection"}
	}

	key := args[1].(*object.String).Value
	seconds := time.Duration(args[2].(*object.Integer).Value) * time.Second

	result, err := conn.client.Expire(conn.ctx, key, seconds).Result()
	if err != nil {
		return &object.Error{Message: fmt.Sprintf("Redis EXPIRE failed: %s", err)}
	}

	return &object.Boolean{Value: result}
}

// redisTTL gets the time to live of a key
func redisTTL(args []object.VintObject, defs map[string]object.VintObject) object.VintObject {
	if len(args) != 2 || args[0].Type() != object.NATIVE_OBJ || args[1].Type() != object.STRING_OBJ {
		return ErrorMessage(
			"redis",
			"ttl",
			"connection, key (string)",
			formatArgs(args),
			`redis.ttl(conn, "key") -> 60`,
		)
	}

	conn, ok := args[0].(*object.NativeObject).Value.(*RedisConnection)
	if !ok {
		return &object.Error{Message: "Invalid Redis connection"}
	}

	key := args[1].(*object.String).Value

	result, err := conn.client.TTL(conn.ctx, key).Result()
	if err != nil {
		return &object.Error{Message: fmt.Sprintf("Redis TTL failed: %s", err)}
	}

	return &object.Integer{Value: int64(result.Seconds())}
}

// redisIncr increments the integer value of a key by one
func redisIncr(args []object.VintObject, defs map[string]object.VintObject) object.VintObject {
	if len(args) != 2 || args[0].Type() != object.NATIVE_OBJ || args[1].Type() != object.STRING_OBJ {
		return ErrorMessage(
			"redis",
			"incr",
			"connection, key (string)",
			formatArgs(args),
			`redis.incr(conn, "counter") -> 1`,
		)
	}

	conn, ok := args[0].(*object.NativeObject).Value.(*RedisConnection)
	if !ok {
		return &object.Error{Message: "Invalid Redis connection"}
	}

	key := args[1].(*object.String).Value

	result, err := conn.client.Incr(conn.ctx, key).Result()
	if err != nil {
		return &object.Error{Message: fmt.Sprintf("Redis INCR failed: %s", err)}
	}

	return &object.Integer{Value: result}
}

// redisDecr decrements the integer value of a key by one
func redisDecr(args []object.VintObject, defs map[string]object.VintObject) object.VintObject {
	if len(args) != 2 || args[0].Type() != object.NATIVE_OBJ || args[1].Type() != object.STRING_OBJ {
		return ErrorMessage(
			"redis",
			"decr",
			"connection, key (string)",
			formatArgs(args),
			`redis.decr(conn, "counter") -> -1`,
		)
	}

	conn, ok := args[0].(*object.NativeObject).Value.(*RedisConnection)
	if !ok {
		return &object.Error{Message: "Invalid Redis connection"}
	}

	key := args[1].(*object.String).Value

	result, err := conn.client.Decr(conn.ctx, key).Result()
	if err != nil {
		return &object.Error{Message: fmt.Sprintf("Redis DECR failed: %s", err)}
	}

	return &object.Integer{Value: result}
}

// redisIncrBy increments the integer value of a key by the given amount
func redisIncrBy(args []object.VintObject, defs map[string]object.VintObject) object.VintObject {
	if len(args) != 3 || args[0].Type() != object.NATIVE_OBJ || args[1].Type() != object.STRING_OBJ || args[2].Type() != object.INTEGER_OBJ {
		return ErrorMessage(
			"redis",
			"incrby",
			"connection, key (string), increment (integer)",
			formatArgs(args),
			`redis.incrby(conn, "counter", 5) -> 5`,
		)
	}

	conn, ok := args[0].(*object.NativeObject).Value.(*RedisConnection)
	if !ok {
		return &object.Error{Message: "Invalid Redis connection"}
	}

	key := args[1].(*object.String).Value
	increment := args[2].(*object.Integer).Value

	result, err := conn.client.IncrBy(conn.ctx, key, increment).Result()
	if err != nil {
		return &object.Error{Message: fmt.Sprintf("Redis INCRBY failed: %s", err)}
	}

	return &object.Integer{Value: result}
}

// redisDecrBy decrements the integer value of a key by the given amount
func redisDecrBy(args []object.VintObject, defs map[string]object.VintObject) object.VintObject {
	if len(args) != 3 || args[0].Type() != object.NATIVE_OBJ || args[1].Type() != object.STRING_OBJ || args[2].Type() != object.INTEGER_OBJ {
		return ErrorMessage(
			"redis",
			"decrby",
			"connection, key (string), decrement (integer)",
			formatArgs(args),
			`redis.decrby(conn, "counter", 3) -> -3`,
		)
	}

	conn, ok := args[0].(*object.NativeObject).Value.(*RedisConnection)
	if !ok {
		return &object.Error{Message: "Invalid Redis connection"}
	}

	key := args[1].(*object.String).Value
	decrement := args[2].(*object.Integer).Value

	result, err := conn.client.DecrBy(conn.ctx, key, decrement).Result()
	if err != nil {
		return &object.Error{Message: fmt.Sprintf("Redis DECRBY failed: %s", err)}
	}

	return &object.Integer{Value: result}
}

// redisMSet sets multiple key-value pairs
func redisMSet(args []object.VintObject, defs map[string]object.VintObject) object.VintObject {
	if len(args) < 3 || (len(args)-1)%2 != 0 || args[0].Type() != object.NATIVE_OBJ {
		return ErrorMessage(
			"redis",
			"mset",
			"connection, key1, value1, key2, value2, ...",
			formatArgs(args),
			`redis.mset(conn, "key1", "value1", "key2", "value2") -> "OK"`,
		)
	}

	conn, ok := args[0].(*object.NativeObject).Value.(*RedisConnection)
	if !ok {
		return &object.Error{Message: "Invalid Redis connection"}
	}

	pairs := make([]interface{}, len(args)-1)
	for i, arg := range args[1:] {
		if arg.Type() != object.STRING_OBJ {
			return ErrorMessage(
				"redis",
				"mset",
				"all keys and values must be strings",
				formatArgs(args),
				`redis.mset(conn, "key1", "value1", "key2", "value2")`,
			)
		}
		pairs[i] = arg.(*object.String).Value
	}

	result, err := conn.client.MSet(conn.ctx, pairs...).Result()
	if err != nil {
		return &object.Error{Message: fmt.Sprintf("Redis MSET failed: %s", err)}
	}

	return &object.String{Value: result}
}

// redisMGet gets multiple values
func redisMGet(args []object.VintObject, defs map[string]object.VintObject) object.VintObject {
	if len(args) < 2 || args[0].Type() != object.NATIVE_OBJ {
		return ErrorMessage(
			"redis",
			"mget",
			"connection, key1, key2, ...",
			formatArgs(args),
			`redis.mget(conn, "key1", "key2") -> ["value1", "value2"]`,
		)
	}

	conn, ok := args[0].(*object.NativeObject).Value.(*RedisConnection)
	if !ok {
		return &object.Error{Message: "Invalid Redis connection"}
	}

	keys := make([]string, len(args)-1)
	for i, arg := range args[1:] {
		if arg.Type() != object.STRING_OBJ {
			return ErrorMessage(
				"redis",
				"mget",
				"all keys must be strings",
				formatArgs(args),
				`redis.mget(conn, "key1", "key2")`,
			)
		}
		keys[i] = arg.(*object.String).Value
	}

	result, err := conn.client.MGet(conn.ctx, keys...).Result()
	if err != nil {
		return &object.Error{Message: fmt.Sprintf("Redis MGET failed: %s", err)}
	}

	elements := make([]object.VintObject, len(result))
	for i, val := range result {
		if val == nil {
			elements[i] = &object.Null{}
		} else {
			elements[i] = &object.String{Value: val.(string)}
		}
	}

	return &object.Array{Elements: elements}
}

// redisKeys returns all keys matching a pattern
func redisKeys(args []object.VintObject, defs map[string]object.VintObject) object.VintObject {
	if len(args) != 2 || args[0].Type() != object.NATIVE_OBJ || args[1].Type() != object.STRING_OBJ {
		return ErrorMessage(
			"redis",
			"keys",
			"connection, pattern (string)",
			formatArgs(args),
			`redis.keys(conn, "*") -> ["key1", "key2"]`,
		)
	}

	conn, ok := args[0].(*object.NativeObject).Value.(*RedisConnection)
	if !ok {
		return &object.Error{Message: "Invalid Redis connection"}
	}

	pattern := args[1].(*object.String).Value

	result, err := conn.client.Keys(conn.ctx, pattern).Result()
	if err != nil {
		return &object.Error{Message: fmt.Sprintf("Redis KEYS failed: %s", err)}
	}

	elements := make([]object.VintObject, len(result))
	for i, key := range result {
		elements[i] = &object.String{Value: key}
	}

	return &object.Array{Elements: elements}
}

// Hash operations

// redisHSet sets field in hash stored at key
func redisHSet(args []object.VintObject, defs map[string]object.VintObject) object.VintObject {
	if len(args) != 4 || args[0].Type() != object.NATIVE_OBJ || args[1].Type() != object.STRING_OBJ || args[2].Type() != object.STRING_OBJ || args[3].Type() != object.STRING_OBJ {
		return ErrorMessage(
			"redis",
			"hset",
			"connection, key (string), field (string), value (string)",
			formatArgs(args),
			`redis.hset(conn, "hash", "field", "value") -> 1`,
		)
	}

	conn, ok := args[0].(*object.NativeObject).Value.(*RedisConnection)
	if !ok {
		return &object.Error{Message: "Invalid Redis connection"}
	}

	key := args[1].(*object.String).Value
	field := args[2].(*object.String).Value
	value := args[3].(*object.String).Value

	result, err := conn.client.HSet(conn.ctx, key, field, value).Result()
	if err != nil {
		return &object.Error{Message: fmt.Sprintf("Redis HSET failed: %s", err)}
	}

	return &object.Integer{Value: result}
}

// redisHGet gets field from hash stored at key
func redisHGet(args []object.VintObject, defs map[string]object.VintObject) object.VintObject {
	if len(args) != 3 || args[0].Type() != object.NATIVE_OBJ || args[1].Type() != object.STRING_OBJ || args[2].Type() != object.STRING_OBJ {
		return ErrorMessage(
			"redis",
			"hget",
			"connection, key (string), field (string)",
			formatArgs(args),
			`redis.hget(conn, "hash", "field") -> "value"`,
		)
	}

	conn, ok := args[0].(*object.NativeObject).Value.(*RedisConnection)
	if !ok {
		return &object.Error{Message: "Invalid Redis connection"}
	}

	key := args[1].(*object.String).Value
	field := args[2].(*object.String).Value

	result, err := conn.client.HGet(conn.ctx, key, field).Result()
	if err == redis.Nil {
		return &object.Null{}
	}
	if err != nil {
		return &object.Error{Message: fmt.Sprintf("Redis HGET failed: %s", err)}
	}

	return &object.String{Value: result}
}

// redisHGetAll gets all fields and values from hash
func redisHGetAll(args []object.VintObject, defs map[string]object.VintObject) object.VintObject {
	if len(args) != 2 || args[0].Type() != object.NATIVE_OBJ || args[1].Type() != object.STRING_OBJ {
		return ErrorMessage(
			"redis",
			"hgetall",
			"connection, key (string)",
			formatArgs(args),
			`redis.hgetall(conn, "hash") -> {"field1": "value1", "field2": "value2"}`,
		)
	}

	conn, ok := args[0].(*object.NativeObject).Value.(*RedisConnection)
	if !ok {
		return &object.Error{Message: "Invalid Redis connection"}
	}

	key := args[1].(*object.String).Value

	result, err := conn.client.HGetAll(conn.ctx, key).Result()
	if err != nil {
		return &object.Error{Message: fmt.Sprintf("Redis HGETALL failed: %s", err)}
	}

	pairs := make(map[object.HashKey]object.DictPair)
	for field, value := range result {
		fieldKey := (&object.String{Value: field}).HashKey()
		pairs[fieldKey] = object.DictPair{Key: &object.String{Value: field}, Value: &object.String{Value: value}}
	}

	return &object.Dict{Pairs: pairs}
}

// redisHDel deletes one or more hash fields
func redisHDel(args []object.VintObject, defs map[string]object.VintObject) object.VintObject {
	if len(args) < 3 || args[0].Type() != object.NATIVE_OBJ || args[1].Type() != object.STRING_OBJ {
		return ErrorMessage(
			"redis",
			"hdel",
			"connection, key (string), field1 (string), [field2, ...]",
			formatArgs(args),
			`redis.hdel(conn, "hash", "field1", "field2") -> 2`,
		)
	}

	conn, ok := args[0].(*object.NativeObject).Value.(*RedisConnection)
	if !ok {
		return &object.Error{Message: "Invalid Redis connection"}
	}

	key := args[1].(*object.String).Value
	fields := make([]string, len(args)-2)
	for i, arg := range args[2:] {
		if arg.Type() != object.STRING_OBJ {
			return ErrorMessage(
				"redis",
				"hdel",
				"all fields must be strings",
				formatArgs(args),
				`redis.hdel(conn, "hash", "field1", "field2")`,
			)
		}
		fields[i] = arg.(*object.String).Value
	}

	result, err := conn.client.HDel(conn.ctx, key, fields...).Result()
	if err != nil {
		return &object.Error{Message: fmt.Sprintf("Redis HDEL failed: %s", err)}
	}

	return &object.Integer{Value: result}
}

// redisHExists determines if hash field exists
func redisHExists(args []object.VintObject, defs map[string]object.VintObject) object.VintObject {
	if len(args) != 3 || args[0].Type() != object.NATIVE_OBJ || args[1].Type() != object.STRING_OBJ || args[2].Type() != object.STRING_OBJ {
		return ErrorMessage(
			"redis",
			"hexists",
			"connection, key (string), field (string)",
			formatArgs(args),
			`redis.hexists(conn, "hash", "field") -> true/false`,
		)
	}

	conn, ok := args[0].(*object.NativeObject).Value.(*RedisConnection)
	if !ok {
		return &object.Error{Message: "Invalid Redis connection"}
	}

	key := args[1].(*object.String).Value
	field := args[2].(*object.String).Value

	result, err := conn.client.HExists(conn.ctx, key, field).Result()
	if err != nil {
		return &object.Error{Message: fmt.Sprintf("Redis HEXISTS failed: %s", err)}
	}

	return &object.Boolean{Value: result}
}

// redisHKeys gets all field names in hash
func redisHKeys(args []object.VintObject, defs map[string]object.VintObject) object.VintObject {
	if len(args) != 2 || args[0].Type() != object.NATIVE_OBJ || args[1].Type() != object.STRING_OBJ {
		return ErrorMessage(
			"redis",
			"hkeys",
			"connection, key (string)",
			formatArgs(args),
			`redis.hkeys(conn, "hash") -> ["field1", "field2"]`,
		)
	}

	conn, ok := args[0].(*object.NativeObject).Value.(*RedisConnection)
	if !ok {
		return &object.Error{Message: "Invalid Redis connection"}
	}

	key := args[1].(*object.String).Value

	result, err := conn.client.HKeys(conn.ctx, key).Result()
	if err != nil {
		return &object.Error{Message: fmt.Sprintf("Redis HKEYS failed: %s", err)}
	}

	elements := make([]object.VintObject, len(result))
	for i, field := range result {
		elements[i] = &object.String{Value: field}
	}

	return &object.Array{Elements: elements}
}

// redisHVals gets all values in hash
func redisHVals(args []object.VintObject, defs map[string]object.VintObject) object.VintObject {
	if len(args) != 2 || args[0].Type() != object.NATIVE_OBJ || args[1].Type() != object.STRING_OBJ {
		return ErrorMessage(
			"redis",
			"hvals",
			"connection, key (string)",
			formatArgs(args),
			`redis.hvals(conn, "hash") -> ["value1", "value2"]`,
		)
	}

	conn, ok := args[0].(*object.NativeObject).Value.(*RedisConnection)
	if !ok {
		return &object.Error{Message: "Invalid Redis connection"}
	}

	key := args[1].(*object.String).Value

	result, err := conn.client.HVals(conn.ctx, key).Result()
	if err != nil {
		return &object.Error{Message: fmt.Sprintf("Redis HVALS failed: %s", err)}
	}

	elements := make([]object.VintObject, len(result))
	for i, value := range result {
		elements[i] = &object.String{Value: value}
	}

	return &object.Array{Elements: elements}
}

// List operations

// redisLPush prepends one or more values to list
func redisLPush(args []object.VintObject, defs map[string]object.VintObject) object.VintObject {
	if len(args) < 3 || args[0].Type() != object.NATIVE_OBJ || args[1].Type() != object.STRING_OBJ {
		return ErrorMessage(
			"redis",
			"lpush",
			"connection, key (string), value1 (string), [value2, ...]",
			formatArgs(args),
			`redis.lpush(conn, "list", "value1", "value2") -> 2`,
		)
	}

	conn, ok := args[0].(*object.NativeObject).Value.(*RedisConnection)
	if !ok {
		return &object.Error{Message: "Invalid Redis connection"}
	}

	key := args[1].(*object.String).Value
	values := make([]interface{}, len(args)-2)
	for i, arg := range args[2:] {
		if arg.Type() != object.STRING_OBJ {
			return ErrorMessage(
				"redis",
				"lpush",
				"all values must be strings",
				formatArgs(args),
				`redis.lpush(conn, "list", "value1", "value2")`,
			)
		}
		values[i] = arg.(*object.String).Value
	}

	result, err := conn.client.LPush(conn.ctx, key, values...).Result()
	if err != nil {
		return &object.Error{Message: fmt.Sprintf("Redis LPUSH failed: %s", err)}
	}

	return &object.Integer{Value: result}
}

// redisRPush appends one or more values to list
func redisRPush(args []object.VintObject, defs map[string]object.VintObject) object.VintObject {
	if len(args) < 3 || args[0].Type() != object.NATIVE_OBJ || args[1].Type() != object.STRING_OBJ {
		return ErrorMessage(
			"redis",
			"rpush",
			"connection, key (string), value1 (string), [value2, ...]",
			formatArgs(args),
			`redis.rpush(conn, "list", "value1", "value2") -> 2`,
		)
	}

	conn, ok := args[0].(*object.NativeObject).Value.(*RedisConnection)
	if !ok {
		return &object.Error{Message: "Invalid Redis connection"}
	}

	key := args[1].(*object.String).Value
	values := make([]interface{}, len(args)-2)
	for i, arg := range args[2:] {
		if arg.Type() != object.STRING_OBJ {
			return ErrorMessage(
				"redis",
				"rpush",
				"all values must be strings",
				formatArgs(args),
				`redis.rpush(conn, "list", "value1", "value2")`,
			)
		}
		values[i] = arg.(*object.String).Value
	}

	result, err := conn.client.RPush(conn.ctx, key, values...).Result()
	if err != nil {
		return &object.Error{Message: fmt.Sprintf("Redis RPUSH failed: %s", err)}
	}

	return &object.Integer{Value: result}
}

// redisLPop removes and returns first element of list
func redisLPop(args []object.VintObject, defs map[string]object.VintObject) object.VintObject {
	if len(args) != 2 || args[0].Type() != object.NATIVE_OBJ || args[1].Type() != object.STRING_OBJ {
		return ErrorMessage(
			"redis",
			"lpop",
			"connection, key (string)",
			formatArgs(args),
			`redis.lpop(conn, "list") -> "value"`,
		)
	}

	conn, ok := args[0].(*object.NativeObject).Value.(*RedisConnection)
	if !ok {
		return &object.Error{Message: "Invalid Redis connection"}
	}

	key := args[1].(*object.String).Value

	result, err := conn.client.LPop(conn.ctx, key).Result()
	if err == redis.Nil {
		return &object.Null{}
	}
	if err != nil {
		return &object.Error{Message: fmt.Sprintf("Redis LPOP failed: %s", err)}
	}

	return &object.String{Value: result}
}

// redisRPop removes and returns last element of list
func redisRPop(args []object.VintObject, defs map[string]object.VintObject) object.VintObject {
	if len(args) != 2 || args[0].Type() != object.NATIVE_OBJ || args[1].Type() != object.STRING_OBJ {
		return ErrorMessage(
			"redis",
			"rpop",
			"connection, key (string)",
			formatArgs(args),
			`redis.rpop(conn, "list") -> "value"`,
		)
	}

	conn, ok := args[0].(*object.NativeObject).Value.(*RedisConnection)
	if !ok {
		return &object.Error{Message: "Invalid Redis connection"}
	}

	key := args[1].(*object.String).Value

	result, err := conn.client.RPop(conn.ctx, key).Result()
	if err == redis.Nil {
		return &object.Null{}
	}
	if err != nil {
		return &object.Error{Message: fmt.Sprintf("Redis RPOP failed: %s", err)}
	}

	return &object.String{Value: result}
}

// redisLLen returns length of list
func redisLLen(args []object.VintObject, defs map[string]object.VintObject) object.VintObject {
	if len(args) != 2 || args[0].Type() != object.NATIVE_OBJ || args[1].Type() != object.STRING_OBJ {
		return ErrorMessage(
			"redis",
			"llen",
			"connection, key (string)",
			formatArgs(args),
			`redis.llen(conn, "list") -> 5`,
		)
	}

	conn, ok := args[0].(*object.NativeObject).Value.(*RedisConnection)
	if !ok {
		return &object.Error{Message: "Invalid Redis connection"}
	}

	key := args[1].(*object.String).Value

	result, err := conn.client.LLen(conn.ctx, key).Result()
	if err != nil {
		return &object.Error{Message: fmt.Sprintf("Redis LLEN failed: %s", err)}
	}

	return &object.Integer{Value: result}
}

// redisLRange returns range of elements from list
func redisLRange(args []object.VintObject, defs map[string]object.VintObject) object.VintObject {
	if len(args) != 4 || args[0].Type() != object.NATIVE_OBJ || args[1].Type() != object.STRING_OBJ || args[2].Type() != object.INTEGER_OBJ || args[3].Type() != object.INTEGER_OBJ {
		return ErrorMessage(
			"redis",
			"lrange",
			"connection, key (string), start (integer), stop (integer)",
			formatArgs(args),
			`redis.lrange(conn, "list", 0, -1) -> ["value1", "value2"]`,
		)
	}

	conn, ok := args[0].(*object.NativeObject).Value.(*RedisConnection)
	if !ok {
		return &object.Error{Message: "Invalid Redis connection"}
	}

	key := args[1].(*object.String).Value
	start := args[2].(*object.Integer).Value
	stop := args[3].(*object.Integer).Value

	result, err := conn.client.LRange(conn.ctx, key, start, stop).Result()
	if err != nil {
		return &object.Error{Message: fmt.Sprintf("Redis LRANGE failed: %s", err)}
	}

	elements := make([]object.VintObject, len(result))
	for i, value := range result {
		elements[i] = &object.String{Value: value}
	}

	return &object.Array{Elements: elements}
}

// Set operations

// redisSAdd adds one or more members to set
func redisSAdd(args []object.VintObject, defs map[string]object.VintObject) object.VintObject {
	if len(args) < 3 || args[0].Type() != object.NATIVE_OBJ || args[1].Type() != object.STRING_OBJ {
		return ErrorMessage(
			"redis",
			"sadd",
			"connection, key (string), member1 (string), [member2, ...]",
			formatArgs(args),
			`redis.sadd(conn, "set", "member1", "member2") -> 2`,
		)
	}

	conn, ok := args[0].(*object.NativeObject).Value.(*RedisConnection)
	if !ok {
		return &object.Error{Message: "Invalid Redis connection"}
	}

	key := args[1].(*object.String).Value
	members := make([]interface{}, len(args)-2)
	for i, arg := range args[2:] {
		if arg.Type() != object.STRING_OBJ {
			return ErrorMessage(
				"redis",
				"sadd",
				"all members must be strings",
				formatArgs(args),
				`redis.sadd(conn, "set", "member1", "member2")`,
			)
		}
		members[i] = arg.(*object.String).Value
	}

	result, err := conn.client.SAdd(conn.ctx, key, members...).Result()
	if err != nil {
		return &object.Error{Message: fmt.Sprintf("Redis SADD failed: %s", err)}
	}

	return &object.Integer{Value: result}
}

// redisSRem removes one or more members from set
func redisSRem(args []object.VintObject, defs map[string]object.VintObject) object.VintObject {
	if len(args) < 3 || args[0].Type() != object.NATIVE_OBJ || args[1].Type() != object.STRING_OBJ {
		return ErrorMessage(
			"redis",
			"srem",
			"connection, key (string), member1 (string), [member2, ...]",
			formatArgs(args),
			`redis.srem(conn, "set", "member1", "member2") -> 2`,
		)
	}

	conn, ok := args[0].(*object.NativeObject).Value.(*RedisConnection)
	if !ok {
		return &object.Error{Message: "Invalid Redis connection"}
	}

	key := args[1].(*object.String).Value
	members := make([]interface{}, len(args)-2)
	for i, arg := range args[2:] {
		if arg.Type() != object.STRING_OBJ {
			return ErrorMessage(
				"redis",
				"srem",
				"all members must be strings",
				formatArgs(args),
				`redis.srem(conn, "set", "member1", "member2")`,
			)
		}
		members[i] = arg.(*object.String).Value
	}

	result, err := conn.client.SRem(conn.ctx, key, members...).Result()
	if err != nil {
		return &object.Error{Message: fmt.Sprintf("Redis SREM failed: %s", err)}
	}

	return &object.Integer{Value: result}
}

// redisSMembers returns all members of set
func redisSMembers(args []object.VintObject, defs map[string]object.VintObject) object.VintObject {
	if len(args) != 2 || args[0].Type() != object.NATIVE_OBJ || args[1].Type() != object.STRING_OBJ {
		return ErrorMessage(
			"redis",
			"smembers",
			"connection, key (string)",
			formatArgs(args),
			`redis.smembers(conn, "set") -> ["member1", "member2"]`,
		)
	}

	conn, ok := args[0].(*object.NativeObject).Value.(*RedisConnection)
	if !ok {
		return &object.Error{Message: "Invalid Redis connection"}
	}

	key := args[1].(*object.String).Value

	result, err := conn.client.SMembers(conn.ctx, key).Result()
	if err != nil {
		return &object.Error{Message: fmt.Sprintf("Redis SMEMBERS failed: %s", err)}
	}

	elements := make([]object.VintObject, len(result))
	for i, member := range result {
		elements[i] = &object.String{Value: member}
	}

	return &object.Array{Elements: elements}
}

// redisSCard returns cardinality (number of members) of set
func redisSCard(args []object.VintObject, defs map[string]object.VintObject) object.VintObject {
	if len(args) != 2 || args[0].Type() != object.NATIVE_OBJ || args[1].Type() != object.STRING_OBJ {
		return ErrorMessage(
			"redis",
			"scard",
			"connection, key (string)",
			formatArgs(args),
			`redis.scard(conn, "set") -> 3`,
		)
	}

	conn, ok := args[0].(*object.NativeObject).Value.(*RedisConnection)
	if !ok {
		return &object.Error{Message: "Invalid Redis connection"}
	}

	key := args[1].(*object.String).Value

	result, err := conn.client.SCard(conn.ctx, key).Result()
	if err != nil {
		return &object.Error{Message: fmt.Sprintf("Redis SCARD failed: %s", err)}
	}

	return &object.Integer{Value: result}
}

// redisSIsMember determines if member is in set
func redisSIsMember(args []object.VintObject, defs map[string]object.VintObject) object.VintObject {
	if len(args) != 3 || args[0].Type() != object.NATIVE_OBJ || args[1].Type() != object.STRING_OBJ || args[2].Type() != object.STRING_OBJ {
		return ErrorMessage(
			"redis",
			"sismember",
			"connection, key (string), member (string)",
			formatArgs(args),
			`redis.sismember(conn, "set", "member") -> true/false`,
		)
	}

	conn, ok := args[0].(*object.NativeObject).Value.(*RedisConnection)
	if !ok {
		return &object.Error{Message: "Invalid Redis connection"}
	}

	key := args[1].(*object.String).Value
	member := args[2].(*object.String).Value

	result, err := conn.client.SIsMember(conn.ctx, key, member).Result()
	if err != nil {
		return &object.Error{Message: fmt.Sprintf("Redis SISMEMBER failed: %s", err)}
	}

	return &object.Boolean{Value: result}
}

// Sorted set operations

// redisZAdd adds one or more members to sorted set
func redisZAdd(args []object.VintObject, defs map[string]object.VintObject) object.VintObject {
	if len(args) < 4 || (len(args)-2)%2 != 0 || args[0].Type() != object.NATIVE_OBJ || args[1].Type() != object.STRING_OBJ {
		return ErrorMessage(
			"redis",
			"zadd",
			"connection, key (string), score1 (float/int), member1 (string), [score2, member2, ...]",
			formatArgs(args),
			`redis.zadd(conn, "zset", 1.0, "member1", 2.0, "member2") -> 2`,
		)
	}

	conn, ok := args[0].(*object.NativeObject).Value.(*RedisConnection)
	if !ok {
		return &object.Error{Message: "Invalid Redis connection"}
	}

	key := args[1].(*object.String).Value
	members := make([]redis.Z, 0)

	for i := 2; i < len(args); i += 2 {
		var score float64
		switch args[i].Type() {
		case object.INTEGER_OBJ:
			score = float64(args[i].(*object.Integer).Value)
		case object.FLOAT_OBJ:
			score = args[i].(*object.Float).Value
		default:
			return ErrorMessage(
				"redis",
				"zadd",
				"scores must be numbers",
				formatArgs(args),
				`redis.zadd(conn, "zset", 1.0, "member1")`,
			)
		}

		if args[i+1].Type() != object.STRING_OBJ {
			return ErrorMessage(
				"redis",
				"zadd",
				"members must be strings",
				formatArgs(args),
				`redis.zadd(conn, "zset", 1.0, "member1")`,
			)
		}

		member := args[i+1].(*object.String).Value
		members = append(members, redis.Z{Score: score, Member: member})
	}

	result, err := conn.client.ZAdd(conn.ctx, key, members...).Result()
	if err != nil {
		return &object.Error{Message: fmt.Sprintf("Redis ZADD failed: %s", err)}
	}

	return &object.Integer{Value: result}
}

// redisZRem removes one or more members from sorted set
func redisZRem(args []object.VintObject, defs map[string]object.VintObject) object.VintObject {
	if len(args) < 3 || args[0].Type() != object.NATIVE_OBJ || args[1].Type() != object.STRING_OBJ {
		return ErrorMessage(
			"redis",
			"zrem",
			"connection, key (string), member1 (string), [member2, ...]",
			formatArgs(args),
			`redis.zrem(conn, "zset", "member1", "member2") -> 2`,
		)
	}

	conn, ok := args[0].(*object.NativeObject).Value.(*RedisConnection)
	if !ok {
		return &object.Error{Message: "Invalid Redis connection"}
	}

	key := args[1].(*object.String).Value
	members := make([]interface{}, len(args)-2)
	for i, arg := range args[2:] {
		if arg.Type() != object.STRING_OBJ {
			return ErrorMessage(
				"redis",
				"zrem",
				"all members must be strings",
				formatArgs(args),
				`redis.zrem(conn, "zset", "member1", "member2")`,
			)
		}
		members[i] = arg.(*object.String).Value
	}

	result, err := conn.client.ZRem(conn.ctx, key, members...).Result()
	if err != nil {
		return &object.Error{Message: fmt.Sprintf("Redis ZREM failed: %s", err)}
	}

	return &object.Integer{Value: result}
}

// redisZRange returns range of members in sorted set by index
func redisZRange(args []object.VintObject, defs map[string]object.VintObject) object.VintObject {
	if len(args) != 4 || args[0].Type() != object.NATIVE_OBJ || args[1].Type() != object.STRING_OBJ || args[2].Type() != object.INTEGER_OBJ || args[3].Type() != object.INTEGER_OBJ {
		return ErrorMessage(
			"redis",
			"zrange",
			"connection, key (string), start (integer), stop (integer)",
			formatArgs(args),
			`redis.zrange(conn, "zset", 0, -1) -> ["member1", "member2"]`,
		)
	}

	conn, ok := args[0].(*object.NativeObject).Value.(*RedisConnection)
	if !ok {
		return &object.Error{Message: "Invalid Redis connection"}
	}

	key := args[1].(*object.String).Value
	start := args[2].(*object.Integer).Value
	stop := args[3].(*object.Integer).Value

	result, err := conn.client.ZRange(conn.ctx, key, start, stop).Result()
	if err != nil {
		return &object.Error{Message: fmt.Sprintf("Redis ZRANGE failed: %s", err)}
	}

	elements := make([]object.VintObject, len(result))
	for i, member := range result {
		elements[i] = &object.String{Value: member}
	}

	return &object.Array{Elements: elements}
}

// redisZCard returns cardinality (number of members) of sorted set
func redisZCard(args []object.VintObject, defs map[string]object.VintObject) object.VintObject {
	if len(args) != 2 || args[0].Type() != object.NATIVE_OBJ || args[1].Type() != object.STRING_OBJ {
		return ErrorMessage(
			"redis",
			"zcard",
			"connection, key (string)",
			formatArgs(args),
			`redis.zcard(conn, "zset") -> 5`,
		)
	}

	conn, ok := args[0].(*object.NativeObject).Value.(*RedisConnection)
	if !ok {
		return &object.Error{Message: "Invalid Redis connection"}
	}

	key := args[1].(*object.String).Value

	result, err := conn.client.ZCard(conn.ctx, key).Result()
	if err != nil {
		return &object.Error{Message: fmt.Sprintf("Redis ZCARD failed: %s", err)}
	}

	return &object.Integer{Value: result}
}

// redisZScore returns score of member in sorted set
func redisZScore(args []object.VintObject, defs map[string]object.VintObject) object.VintObject {
	if len(args) != 3 || args[0].Type() != object.NATIVE_OBJ || args[1].Type() != object.STRING_OBJ || args[2].Type() != object.STRING_OBJ {
		return ErrorMessage(
			"redis",
			"zscore",
			"connection, key (string), member (string)",
			formatArgs(args),
			`redis.zscore(conn, "zset", "member") -> 1.5`,
		)
	}

	conn, ok := args[0].(*object.NativeObject).Value.(*RedisConnection)
	if !ok {
		return &object.Error{Message: "Invalid Redis connection"}
	}

	key := args[1].(*object.String).Value
	member := args[2].(*object.String).Value

	result, err := conn.client.ZScore(conn.ctx, key, member).Result()
	if err == redis.Nil {
		return &object.Null{}
	}
	if err != nil {
		return &object.Error{Message: fmt.Sprintf("Redis ZSCORE failed: %s", err)}
	}

	return &object.Float{Value: result}
}
