package object

import (
	"fmt"
	"sync"
)

// Promise represents a value that will be available in the future
type Promise struct {
	Value     VintObject
	Error     VintObject
	Done      bool
	mu        sync.Mutex
	callbacks []func(VintObject, VintObject)
	waitChan  chan struct{}
}

func (p *Promise) Type() VintObjectType {
	return PROMISE_OBJ
}

func (p *Promise) Inspect() string {
	p.mu.Lock()
	defer p.mu.Unlock()

	if p.Done {
		if p.Error != nil {
			return fmt.Sprintf("Promise{rejected: %s}", p.Error.Inspect())
		}
		return fmt.Sprintf("Promise{resolved: %s}", p.Value.Inspect())
	}
	return "Promise{pending}"
}

// Resolve sets the promise value and notifies callbacks
func (p *Promise) Resolve(value VintObject) {
	p.mu.Lock()
	defer p.mu.Unlock()

	if p.Done {
		return
	}

	p.Value = value
	p.Done = true

	// Notify waiting goroutines
	close(p.waitChan)

	for _, callback := range p.callbacks {
		go callback(value, nil)
	}
	p.callbacks = nil
}

// Reject sets the promise error and notifies callbacks
func (p *Promise) Reject(err VintObject) {
	p.mu.Lock()
	defer p.mu.Unlock()

	if p.Done {
		return
	}

	p.Error = err
	p.Done = true

	// Notify waiting goroutines
	close(p.waitChan)

	for _, callback := range p.callbacks {
		go callback(nil, err)
	}
	p.callbacks = nil
}

// Then adds a callback to be executed when the promise resolves
func (p *Promise) Then(callback func(VintObject, VintObject)) {
	p.mu.Lock()
	defer p.mu.Unlock()

	if p.Done {
		go callback(p.Value, p.Error)
		return
	}

	p.callbacks = append(p.callbacks, callback)
}

// Wait blocks until the promise is resolved or rejected
func (p *Promise) Wait() {
	<-p.waitChan
}

// NewPromise creates a new Promise
func NewPromise() *Promise {
	return &Promise{
		Done:      false,
		callbacks: make([]func(VintObject, VintObject), 0),
		waitChan:  make(chan struct{}),
	}
}
