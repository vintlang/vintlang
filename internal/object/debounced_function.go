package object

import (
	"sync"
	"time"
)

// ApplyFunctionCallback is a type for the function that applies functions
type ApplyFunctionCallback func(fn VintObject, args []VintObject, line int) VintObject

// DebouncedFunction wraps a function with debouncing functionality
type DebouncedFunction struct {
	Fn         VintObject            // The original function (Function or Builtin)
	Duration   time.Duration         // The debounce delay
	timer      *time.Timer           // Internal timer for debouncing
	mutex      sync.Mutex            // Mutex to handle concurrent calls
	lastArgs   []VintObject          // Store last arguments for execution
	applyFunc  ApplyFunctionCallback // Callback to apply the function
	lastResult VintObject            // Store the last result
}

func (df *DebouncedFunction) Type() VintObjectType { return DEBOUNCED_FUNC_OBJ }

func (df *DebouncedFunction) Inspect() string {
	return "debounced function"
}

// SetApplyFunction sets the callback function for applying functions
func (df *DebouncedFunction) SetApplyFunction(applyFunc ApplyFunctionCallback) {
	df.mutex.Lock()
	defer df.mutex.Unlock()
	df.applyFunc = applyFunc
}

// Call executes the debounced function
func (df *DebouncedFunction) Call(args ...VintObject) VintObject {
	df.mutex.Lock()
	defer df.mutex.Unlock()

	// Store the arguments for later execution
	df.lastArgs = make([]VintObject, len(args))
	copy(df.lastArgs, args)

	// If there's an existing timer, stop it
	if df.timer != nil {
		df.timer.Stop()
	}

	// Create a new timer that will execute the function after the delay
	df.timer = time.AfterFunc(df.Duration, func() {
		df.executeFunction()
	})

	// Return the last result (or null if no previous execution)
	if df.lastResult != nil {
		return df.lastResult
	}
	return &Null{}
}

// executeFunction runs the actual function with the last set of arguments
func (df *DebouncedFunction) executeFunction() {
	df.mutex.Lock()
	args := df.lastArgs
	applyFunc := df.applyFunc
	fn := df.Fn
	df.mutex.Unlock()

	if applyFunc != nil {
		// Use the provided callback to execute the function
		result := applyFunc(fn, args, 0)
		df.mutex.Lock()
		df.lastResult = result
		df.mutex.Unlock()
	} else {
		// Fallback to direct execution for builtin functions
		switch fnTyped := fn.(type) {
		case *Builtin:
			result := fnTyped.Fn(args...)
			df.mutex.Lock()
			df.lastResult = result
			df.mutex.Unlock()
		}
	}
}
