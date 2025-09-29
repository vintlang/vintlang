package evaluator

import (
	"github.com/vintlang/vintlang/ast"
	"github.com/vintlang/vintlang/object"
)

func evalForInExpression(fie *ast.ForIn, env *object.Environment, line int) object.Object {
	// Evaluates the iterable expression
	iterable := Eval(fie.Iterable, env)

	// Check if the iterable object supports iteration
	switch i := iterable.(type) {
	case object.Iterable:
		// Create an isolated iterator to avoid conflicts with nested loops
		iterator := createIsolatedIterator(i)
		return loopIterable(iterator.Next, env, fie, line) // Start looping through the iterable
	default:
		// Returns an error if the iterable object does not support iteration
		return newError("Line %d: for..in loop requires an iterable object, but got %s", line, i.Type())
	}
}

// IsolatedIterator wraps an iterable with its own state to prevent nested loop conflicts
type IsolatedIterator struct {
	original object.Iterable
	index    int
	items    []IteratorItem
}

type IteratorItem struct {
	Key   object.Object
	Value object.Object
}

// createIsolatedIterator creates a snapshot of all items to iterate over
func createIsolatedIterator(original object.Iterable) *IsolatedIterator {
	iterator := &IsolatedIterator{
		original: original,
		index:    0,
		items:    make([]IteratorItem, 0),
	}

	// We Reset the original to start from beginning
	original.Reset()

	//  We Collect all items into our snapshot
	for {
		key, value := original.Next()
		if key == nil {
			break
		}
		iterator.items = append(iterator.items, IteratorItem{Key: key, Value: value})
	}

	// We Reset original again to not affect its state for other uses
	original.Reset()

	return iterator
}

// Next We returns the next key-value pair from the snapshot
func (iter *IsolatedIterator) Next() (object.Object, object.Object) {
	if iter.index >= len(iter.items) {
		return nil, nil
	}

	item := iter.items[iter.index]
	iter.index++
	return item.Key, item.Value
}

// Reset resets the iterator to the beginning
func (iter *IsolatedIterator) Reset() {
	iter.index = 0
}

func loopIterable(
	next func() (object.Object, object.Object),
	env *object.Environment,
	fi *ast.ForIn,
	line int,
) object.Object {
	var ret object.Object
	k, v := next()
	for k != nil {
		loopEnv := object.NewEnclosedEnvironment(env)
		loopEnv.Define(fi.Key, k)
		if fi.Value != "" {
			loopEnv.Define(fi.Value, v)
		}
		ret = Eval(fi.Block, loopEnv)
		if isError(ret) {
			return ret
		}
		if ret != nil {
			if ret.Type() == object.BREAK_OBJ {
				return NULL
			}
			if ret.Type() == object.CONTINUE_OBJ {
				k, v = next()
				continue
			}
			if ret.Type() == object.RETURN_VALUE_OBJ {
				return ret
			}
		}
		k, v = next()
	}
	return NULL
}
