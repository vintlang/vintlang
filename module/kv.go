package module

import (
	"sync"
	"time"

	"github.com/vintlang/vintlang/object"
)

var KvFunctions = map[string]object.ModuleFunction{}

// KvStore represents a thread-safe in-memory key-value store
type KvStore struct {
	data   map[string]*KvItem
	mutex  sync.RWMutex
	ttlMap map[string]*time.Timer // For TTL functionality
}

// KvItem represents a stored item with optional TTL
type KvItem struct {
	Value     object.VintObject
	ExpiresAt *time.Time
}

// Global store instance
var globalStore = &KvStore{
	data:   make(map[string]*KvItem),
	ttlMap: make(map[string]*time.Timer),
}

func init() {
	// Basic operations
	KvFunctions["set"] = kvSet
	KvFunctions["get"] = kvGet
	KvFunctions["delete"] = kvDelete
	KvFunctions["exists"] = kvExists
	KvFunctions["clear"] = kvClear

	// Advanced operations
	KvFunctions["keys"] = kvKeys
	KvFunctions["values"] = kvValues
	KvFunctions["size"] = kvSize
	KvFunctions["isEmpty"] = kvIsEmpty

	// TTL operations
	KvFunctions["setTTL"] = kvSetTTL
	KvFunctions["getTTL"] = kvGetTTL
	KvFunctions["expire"] = kvExpire

	// Bulk operations
	KvFunctions["mget"] = kvMget
	KvFunctions["mset"] = kvMset

	// Atomic operations
	KvFunctions["increment"] = kvIncrement
	KvFunctions["decrement"] = kvDecrement

	// Utility operations
	KvFunctions["dump"] = kvDump
	KvFunctions["stats"] = kvStats
}

// Helper function to check if a key has expired
func (store *KvStore) isExpired(key string) bool {
	if item, exists := store.data[key]; exists && item.ExpiresAt != nil {
		return time.Now().After(*item.ExpiresAt)
	}
	return false
}

// Helper function to clean up expired keys
func (store *KvStore) cleanupExpired(key string) {
	if store.isExpired(key) {
		delete(store.data, key)
		if timer, exists := store.ttlMap[key]; exists {
			timer.Stop()
			delete(store.ttlMap, key)
		}
	}
}

// kvSet stores a key-value pair
func kvSet(args []object.VintObject, defs map[string]object.VintObject) object.VintObject {
	if len(args) != 2 || len(defs) != 0 {
		return ErrorMessage(
			"kv", "set",
			"2 arguments (key: string, value: any)",
			formatArgs(args),
			`kv.set("user:123", userData)`,
		)
	}

	if args[0].Type() != object.STRING_OBJ {
		return ErrorMessage(
			"kv", "set",
			"key must be a string",
			"key is "+string(args[0].Type()),
			`kv.set("mykey", "myvalue")`,
		)
	}

	key := args[0].(*object.String).Value
	value := args[1]

	globalStore.mutex.Lock()
	defer globalStore.mutex.Unlock()

	globalStore.data[key] = &KvItem{
		Value:     value,
		ExpiresAt: nil,
	}

	return &object.Boolean{Value: true}
}

// kvGet retrieves a value by key
func kvGet(args []object.VintObject, defs map[string]object.VintObject) object.VintObject {
	if len(args) != 1 || len(defs) != 0 {
		return ErrorMessage(
			"kv", "get",
			"1 argument (key: string)",
			formatArgs(args),
			`kv.get("user:123")`,
		)
	}

	if args[0].Type() != object.STRING_OBJ {
		return ErrorMessage(
			"kv", "get",
			"key must be a string",
			"key is "+string(args[0].Type()),
			`kv.get("mykey")`,
		)
	}

	key := args[0].(*object.String).Value

	globalStore.mutex.RLock()
	defer globalStore.mutex.RUnlock()

	// Check if key exists and clean up if expired
	if globalStore.isExpired(key) {
		globalStore.mutex.RUnlock()
		globalStore.mutex.Lock()
		globalStore.cleanupExpired(key)
		globalStore.mutex.Unlock()
		globalStore.mutex.RLock()
	}

	if item, exists := globalStore.data[key]; exists {
		return item.Value
	}

	return &object.Null{}
}

// kvDelete removes a key-value pair
func kvDelete(args []object.VintObject, defs map[string]object.VintObject) object.VintObject {
	if len(args) != 1 || len(defs) != 0 {
		return ErrorMessage(
			"kv", "delete",
			"1 argument (key: string)",
			formatArgs(args),
			`kv.delete("user:123")`,
		)
	}

	if args[0].Type() != object.STRING_OBJ {
		return ErrorMessage(
			"kv", "delete",
			"key must be a string",
			"key is "+string(args[0].Type()),
			`kv.delete("mykey")`,
		)
	}

	key := args[0].(*object.String).Value

	globalStore.mutex.Lock()
	defer globalStore.mutex.Unlock()

	if _, exists := globalStore.data[key]; exists {
		delete(globalStore.data, key)
		if timer, exists := globalStore.ttlMap[key]; exists {
			timer.Stop()
			delete(globalStore.ttlMap, key)
		}
		return &object.Boolean{Value: true}
	}

	return &object.Boolean{Value: false}
}

// kvExists checks if a key exists
func kvExists(args []object.VintObject, defs map[string]object.VintObject) object.VintObject {
	if len(args) != 1 || len(defs) != 0 {
		return ErrorMessage(
			"kv", "exists",
			"1 argument (key: string)",
			formatArgs(args),
			`kv.exists("user:123")`,
		)
	}

	if args[0].Type() != object.STRING_OBJ {
		return ErrorMessage(
			"kv", "exists",
			"key must be a string",
			"key is "+string(args[0].Type()),
			`kv.exists("mykey")`,
		)
	}

	key := args[0].(*object.String).Value

	globalStore.mutex.RLock()
	defer globalStore.mutex.RUnlock()

	// Check expiration
	if globalStore.isExpired(key) {
		return &object.Boolean{Value: false}
	}

	_, exists := globalStore.data[key]
	return &object.Boolean{Value: exists}
}

// kvClear removes all key-value pairs
func kvClear(args []object.VintObject, defs map[string]object.VintObject) object.VintObject {
	if len(args) != 0 || len(defs) != 0 {
		return ErrorMessage(
			"kv", "clear",
			"no arguments",
			formatArgs(args),
			`kv.clear()`,
		)
	}

	globalStore.mutex.Lock()
	defer globalStore.mutex.Unlock()

	// Stop all TTL timers
	for _, timer := range globalStore.ttlMap {
		timer.Stop()
	}

	globalStore.data = make(map[string]*KvItem)
	globalStore.ttlMap = make(map[string]*time.Timer)

	return &object.Boolean{Value: true}
}

// kvKeys returns all keys in the store
func kvKeys(args []object.VintObject, defs map[string]object.VintObject) object.VintObject {
	if len(args) != 0 || len(defs) != 0 {
		return ErrorMessage(
			"kv", "keys",
			"no arguments",
			formatArgs(args),
			`kv.keys()`,
		)
	}

	globalStore.mutex.RLock()
	defer globalStore.mutex.RUnlock()

	var keys []object.VintObject
	for key, item := range globalStore.data {
		// Skip expired items
		if item.ExpiresAt != nil && time.Now().After(*item.ExpiresAt) {
			continue
		}
		keys = append(keys, &object.String{Value: key})
	}

	return &object.Array{Elements: keys}
}

// kvValues returns all values in the store
func kvValues(args []object.VintObject, defs map[string]object.VintObject) object.VintObject {
	if len(args) != 0 || len(defs) != 0 {
		return ErrorMessage(
			"kv", "values",
			"no arguments",
			formatArgs(args),
			`kv.values()`,
		)
	}

	globalStore.mutex.RLock()
	defer globalStore.mutex.RUnlock()

	var values []object.VintObject
	for _, item := range globalStore.data {
		// Skip expired items
		if item.ExpiresAt != nil && time.Now().After(*item.ExpiresAt) {
			continue
		}
		values = append(values, item.Value)
	}

	return &object.Array{Elements: values}
}

// kvSize returns the number of keys in the store
func kvSize(args []object.VintObject, defs map[string]object.VintObject) object.VintObject {
	if len(args) != 0 || len(defs) != 0 {
		return ErrorMessage(
			"kv", "size",
			"no arguments",
			formatArgs(args),
			`kv.size()`,
		)
	}

	globalStore.mutex.RLock()
	defer globalStore.mutex.RUnlock()

	count := 0
	for _, item := range globalStore.data {
		// Skip expired items
		if item.ExpiresAt != nil && time.Now().After(*item.ExpiresAt) {
			continue
		}
		count++
	}

	return &object.Integer{Value: int64(count)}
}

// kvIsEmpty checks if the store is empty
func kvIsEmpty(args []object.VintObject, defs map[string]object.VintObject) object.VintObject {
	if len(args) != 0 || len(defs) != 0 {
		return ErrorMessage(
			"kv", "isEmpty",
			"no arguments",
			formatArgs(args),
			`kv.isEmpty()`,
		)
	}

	globalStore.mutex.RLock()
	defer globalStore.mutex.RUnlock()

	for _, item := range globalStore.data {
		// Skip expired items
		if item.ExpiresAt != nil && time.Now().After(*item.ExpiresAt) {
			continue
		}
		return &object.Boolean{Value: false}
	}

	return &object.Boolean{Value: true}
}

// kvSetTTL sets a key with TTL (time to live) in seconds
func kvSetTTL(args []object.VintObject, defs map[string]object.VintObject) object.VintObject {
	if len(args) != 3 || len(defs) != 0 {
		return ErrorMessage(
			"kv", "setTTL",
			"3 arguments (key: string, value: any, ttl_seconds: integer)",
			formatArgs(args),
			`kv.setTTL("session:123", userData, 300)`,
		)
	}

	if args[0].Type() != object.STRING_OBJ {
		return ErrorMessage(
			"kv", "setTTL",
			"key must be a string",
			"key is "+string(args[0].Type()),
			`kv.setTTL("mykey", "value", 60)`,
		)
	}

	if args[2].Type() != object.INTEGER_OBJ {
		return ErrorMessage(
			"kv", "setTTL",
			"TTL must be an integer (seconds)",
			"TTL is "+string(args[2].Type()),
			`kv.setTTL("mykey", "value", 60)`,
		)
	}

	key := args[0].(*object.String).Value
	value := args[1]
	ttlSeconds := args[2].(*object.Integer).Value

	if ttlSeconds <= 0 {
		return ErrorMessage(
			"kv", "setTTL",
			"TTL must be positive",
			"TTL is "+args[2].Inspect(),
			`kv.setTTL("mykey", "value", 60)`,
		)
	}

	globalStore.mutex.Lock()
	defer globalStore.mutex.Unlock()

	expiresAt := time.Now().Add(time.Duration(ttlSeconds) * time.Second)

	// Stop existing timer if any
	if timer, exists := globalStore.ttlMap[key]; exists {
		timer.Stop()
	}

	globalStore.data[key] = &KvItem{
		Value:     value,
		ExpiresAt: &expiresAt,
	}

	// Set up TTL timer
	globalStore.ttlMap[key] = time.AfterFunc(time.Duration(ttlSeconds)*time.Second, func() {
		globalStore.mutex.Lock()
		defer globalStore.mutex.Unlock()
		delete(globalStore.data, key)
		delete(globalStore.ttlMap, key)
	})

	return &object.Boolean{Value: true}
}

// kvGetTTL gets the remaining TTL for a key in seconds
func kvGetTTL(args []object.VintObject, defs map[string]object.VintObject) object.VintObject {
	if len(args) != 1 || len(defs) != 0 {
		return ErrorMessage(
			"kv", "getTTL",
			"1 argument (key: string)",
			formatArgs(args),
			`kv.getTTL("session:123")`,
		)
	}

	if args[0].Type() != object.STRING_OBJ {
		return ErrorMessage(
			"kv", "getTTL",
			"key must be a string",
			"key is "+string(args[0].Type()),
			`kv.getTTL("mykey")`,
		)
	}

	key := args[0].(*object.String).Value

	globalStore.mutex.RLock()
	defer globalStore.mutex.RUnlock()

	if item, exists := globalStore.data[key]; exists {
		if item.ExpiresAt == nil {
			return &object.Integer{Value: -1} // No TTL set
		}

		remaining := item.ExpiresAt.Sub(time.Now()).Seconds()
		if remaining <= 0 {
			return &object.Integer{Value: 0} // Expired
		}

		return &object.Integer{Value: int64(remaining)}
	}

	return &object.Integer{Value: -2} // Key doesn't exist
}

// kvExpire sets TTL for an existing key
func kvExpire(args []object.VintObject, defs map[string]object.VintObject) object.VintObject {
	if len(args) != 2 || len(defs) != 0 {
		return ErrorMessage(
			"kv", "expire",
			"2 arguments (key: string, ttl_seconds: integer)",
			formatArgs(args),
			`kv.expire("session:123", 300)`,
		)
	}

	if args[0].Type() != object.STRING_OBJ {
		return ErrorMessage(
			"kv", "expire",
			"key must be a string",
			"key is "+string(args[0].Type()),
			`kv.expire("mykey", 60)`,
		)
	}

	if args[1].Type() != object.INTEGER_OBJ {
		return ErrorMessage(
			"kv", "expire",
			"TTL must be an integer (seconds)",
			"TTL is "+string(args[1].Type()),
			`kv.expire("mykey", 60)`,
		)
	}

	key := args[0].(*object.String).Value
	ttlSeconds := args[1].(*object.Integer).Value

	if ttlSeconds <= 0 {
		return ErrorMessage(
			"kv", "expire",
			"TTL must be positive",
			"TTL is "+args[1].Inspect(),
			`kv.expire("mykey", 60)`,
		)
	}

	globalStore.mutex.Lock()
	defer globalStore.mutex.Unlock()

	if item, exists := globalStore.data[key]; exists {
		expiresAt := time.Now().Add(time.Duration(ttlSeconds) * time.Second)
		item.ExpiresAt = &expiresAt

		// Stop existing timer if any
		if timer, exists := globalStore.ttlMap[key]; exists {
			timer.Stop()
		}

		// Set up new TTL timer
		globalStore.ttlMap[key] = time.AfterFunc(time.Duration(ttlSeconds)*time.Second, func() {
			globalStore.mutex.Lock()
			defer globalStore.mutex.Unlock()
			delete(globalStore.data, key)
			delete(globalStore.ttlMap, key)
		})

		return &object.Boolean{Value: true}
	}

	return &object.Boolean{Value: false}
}

// kvMget gets multiple values by keys
func kvMget(args []object.VintObject, defs map[string]object.VintObject) object.VintObject {
	if len(args) != 1 || len(defs) != 0 {
		return ErrorMessage(
			"kv", "mget",
			"1 argument (keys: array of strings)",
			formatArgs(args),
			`kv.mget(["user:123", "session:456"])`,
		)
	}

	if args[0].Type() != object.ARRAY_OBJ {
		return ErrorMessage(
			"kv", "mget",
			"keys must be an array",
			"keys is "+string(args[0].Type()),
			`kv.mget(["key1", "key2"])`,
		)
	}

	keysArray := args[0].(*object.Array)
	results := make([]object.VintObject, len(keysArray.Elements))

	globalStore.mutex.RLock()
	defer globalStore.mutex.RUnlock()

	for i, keyObj := range keysArray.Elements {
		if keyObj.Type() != object.STRING_OBJ {
			return ErrorMessage(
				"kv", "mget",
				"all keys must be strings",
				"key at index "+string(rune(i+48))+" is "+string(keyObj.Type()),
				`kv.mget(["key1", "key2"])`,
			)
		}

		key := keyObj.(*object.String).Value

		// Check expiration and get value
		if item, exists := globalStore.data[key]; exists {
			if item.ExpiresAt == nil || time.Now().Before(*item.ExpiresAt) {
				results[i] = item.Value
			} else {
				results[i] = &object.Null{}
			}
		} else {
			results[i] = &object.Null{}
		}
	}

	return &object.Array{Elements: results}
}

// kvMset sets multiple key-value pairs
func kvMset(args []object.VintObject, defs map[string]object.VintObject) object.VintObject {
	if len(args) != 1 || len(defs) != 0 {
		return ErrorMessage(
			"kv", "mset",
			"1 argument (pairs: dictionary of key-value pairs)",
			formatArgs(args),
			`kv.mset({"user:123": userData, "session:456": sessionData})`,
		)
	}

	if args[0].Type() != object.DICT_OBJ {
		return ErrorMessage(
			"kv", "mset",
			"pairs must be a dictionary",
			"pairs is "+string(args[0].Type()),
			`kv.mset({"key1": "value1", "key2": "value2"})`,
		)
	}

	pairsDict := args[0].(*object.Dict)

	globalStore.mutex.Lock()
	defer globalStore.mutex.Unlock()

	for _, pair := range pairsDict.Pairs {
		if pair.Key.Type() != object.STRING_OBJ {
			return ErrorMessage(
				"kv", "mset",
				"all keys must be strings",
				"found key of type "+string(pair.Key.Type()),
				`kv.mset({"key1": "value1", "key2": "value2"})`,
			)
		}

		key := pair.Key.(*object.String).Value
		globalStore.data[key] = &KvItem{
			Value:     pair.Value,
			ExpiresAt: nil,
		}
	}

	return &object.Boolean{Value: true}
}

// kvIncrement atomically increments a numeric value
func kvIncrement(args []object.VintObject, defs map[string]object.VintObject) object.VintObject {
	if len(args) < 1 || len(args) > 2 || len(defs) != 0 {
		return ErrorMessage(
			"kv", "increment",
			"1-2 arguments (key: string, [delta: integer = 1])",
			formatArgs(args),
			`kv.increment("counter") or kv.increment("counter", 5)`,
		)
	}

	if args[0].Type() != object.STRING_OBJ {
		return ErrorMessage(
			"kv", "increment",
			"key must be a string",
			"key is "+string(args[0].Type()),
			`kv.increment("counter")`,
		)
	}

	key := args[0].(*object.String).Value
	delta := int64(1)

	if len(args) == 2 {
		if args[1].Type() != object.INTEGER_OBJ {
			return ErrorMessage(
				"kv", "increment",
				"delta must be an integer",
				"delta is "+string(args[1].Type()),
				`kv.increment("counter", 5)`,
			)
		}
		delta = args[1].(*object.Integer).Value
	}

	globalStore.mutex.Lock()
	defer globalStore.mutex.Unlock()

	// Check if key exists and is not expired
	if item, exists := globalStore.data[key]; exists {
		if item.ExpiresAt != nil && time.Now().After(*item.ExpiresAt) {
			// Key expired, remove it and create new
			delete(globalStore.data, key)
			if timer, exists := globalStore.ttlMap[key]; exists {
				timer.Stop()
				delete(globalStore.ttlMap, key)
			}
		} else {
			// Key exists and valid, try to increment
			if item.Value.Type() == object.INTEGER_OBJ {
				currentValue := item.Value.(*object.Integer).Value
				newValue := currentValue + delta
				item.Value = &object.Integer{Value: newValue}
				return &object.Integer{Value: newValue}
			} else {
				return ErrorMessage(
					"kv", "increment",
					"existing value must be an integer",
					"existing value is "+string(item.Value.Type()),
					`kv.increment("counter")`,
				)
			}
		}
	}

	// Key doesn't exist or was expired, create new
	newValue := delta
	globalStore.data[key] = &KvItem{
		Value:     &object.Integer{Value: newValue},
		ExpiresAt: nil,
	}

	return &object.Integer{Value: newValue}
}

// kvDecrement atomically decrements a numeric value
func kvDecrement(args []object.VintObject, defs map[string]object.VintObject) object.VintObject {
	if len(args) < 1 || len(args) > 2 || len(defs) != 0 {
		return ErrorMessage(
			"kv", "decrement",
			"1-2 arguments (key: string, [delta: integer = 1])",
			formatArgs(args),
			`kv.decrement("counter") or kv.decrement("counter", 5)`,
		)
	}

	if args[0].Type() != object.STRING_OBJ {
		return ErrorMessage(
			"kv", "decrement",
			"key must be a string",
			"key is "+string(args[0].Type()),
			`kv.decrement("counter")`,
		)
	}

	key := args[0].(*object.String).Value
	delta := int64(1)

	if len(args) == 2 {
		if args[1].Type() != object.INTEGER_OBJ {
			return ErrorMessage(
				"kv", "decrement",
				"delta must be an integer",
				"delta is "+string(args[1].Type()),
				`kv.decrement("counter", 5)`,
			)
		}
		delta = args[1].(*object.Integer).Value
	}

	globalStore.mutex.Lock()
	defer globalStore.mutex.Unlock()

	// Check if key exists and is not expired
	if item, exists := globalStore.data[key]; exists {
		if item.ExpiresAt != nil && time.Now().After(*item.ExpiresAt) {
			// Key expired, remove it and create new
			delete(globalStore.data, key)
			if timer, exists := globalStore.ttlMap[key]; exists {
				timer.Stop()
				delete(globalStore.ttlMap, key)
			}
		} else {
			// Key exists and valid, try to decrement
			if item.Value.Type() == object.INTEGER_OBJ {
				currentValue := item.Value.(*object.Integer).Value
				newValue := currentValue - delta
				item.Value = &object.Integer{Value: newValue}
				return &object.Integer{Value: newValue}
			} else {
				return ErrorMessage(
					"kv", "decrement",
					"existing value must be an integer",
					"existing value is "+string(item.Value.Type()),
					`kv.decrement("counter")`,
				)
			}
		}
	}

	// Key doesn't exist or was expired, create new with negative delta
	newValue := -delta
	globalStore.data[key] = &KvItem{
		Value:     &object.Integer{Value: newValue},
		ExpiresAt: nil,
	}

	return &object.Integer{Value: newValue}
}

// kvDump returns all key-value pairs as a dictionary
func kvDump(args []object.VintObject, defs map[string]object.VintObject) object.VintObject {
	if len(args) != 0 || len(defs) != 0 {
		return ErrorMessage(
			"kv", "dump",
			"no arguments",
			formatArgs(args),
			`kv.dump()`,
		)
	}

	globalStore.mutex.RLock()
	defer globalStore.mutex.RUnlock()

	pairs := make(map[object.HashKey]object.DictPair)
	for key, item := range globalStore.data {
		// Skip expired items
		if item.ExpiresAt != nil && time.Now().After(*item.ExpiresAt) {
			continue
		}
		keyObj := &object.String{Value: key}
		hashKey := keyObj.HashKey()
		pairs[hashKey] = object.DictPair{Key: keyObj, Value: item.Value}
	}

	return &object.Dict{Pairs: pairs}
}

// kvStats returns statistics about the KV store
func kvStats(args []object.VintObject, defs map[string]object.VintObject) object.VintObject {
	if len(args) != 0 || len(defs) != 0 {
		return ErrorMessage(
			"kv", "stats",
			"no arguments",
			formatArgs(args),
			`kv.stats()`,
		)
	}

	globalStore.mutex.RLock()
	defer globalStore.mutex.RUnlock()

	totalKeys := int64(len(globalStore.data))
	expiredKeys := int64(0)
	keysWithTTL := int64(0)

	for _, item := range globalStore.data {
		if item.ExpiresAt != nil {
			keysWithTTL++
			if time.Now().After(*item.ExpiresAt) {
				expiredKeys++
			}
		}
	}

	activeKeys := totalKeys - expiredKeys

	stats := make(map[object.VintObject]*object.DictItem)
	stats[&object.String{Value: "total_keys"}] = &object.DictItem{Value: &object.Integer{Value: totalKeys}}
	stats[&object.String{Value: "active_keys"}] = &object.DictItem{Value: &object.Integer{Value: activeKeys}}
	stats[&object.String{Value: "expired_keys"}] = &object.DictItem{Value: &object.Integer{Value: expiredKeys}}
	stats[&object.String{Value: "keys_with_ttl"}] = &object.DictItem{Value: &object.Integer{Value: keysWithTTL}}

	return &object.Dict{Pairs: stats}
}
