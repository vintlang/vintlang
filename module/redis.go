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