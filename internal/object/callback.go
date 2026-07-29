package object

// FuncCaller is a callback type that allows modules to invoke Vint functions
// from Go code (e.g., HTTP handlers running in separate goroutines).
// The evaluator registers an implementation of this callback.
type FuncCaller func(fn *Function, args []VintObject) VintObject

// globalFuncCaller holds the registered function caller callback.
var globalFuncCaller FuncCaller

// RegisterFuncCaller registers a callback that modules can use to call Vint functions.
func RegisterFuncCaller(caller FuncCaller) {
	globalFuncCaller = caller
}

// CallFunction invokes a Vint function using the registered callback.
// Returns nil if no callback is registered.
func CallFunction(fn *Function, args []VintObject) VintObject {
	if globalFuncCaller == nil {
		return &Error{Message: "Function caller not registered"}
	}
	return globalFuncCaller(fn, args)
}
