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
			"key is " + string(args[0].Type()),
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
			"key is " + string(args[0].Type()),
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
			"key is " + string(args[0].Type()),
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
			"key is " + string(args[0].Type()),
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

