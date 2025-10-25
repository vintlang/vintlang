package object

import (
	"fmt"
)

// Channel represents a communication channel between goroutines
type Channel struct {
	ch       chan VintObject
	closed   bool
	buffered bool
	size     int
}

func (c *Channel) Type() VintObjectType {
	return CHANNEL_OBJ
}

func (c *Channel) Inspect() string {
	if c.buffered {
		return fmt.Sprintf("chan(buffered:%d)", c.size)
	}
	return "chan(unbuffered)"
}

// Send sends a value to the channel
func (c *Channel) Send(value VintObject) error {
	if c.closed {
		return fmt.Errorf("send on closed channel")
	}

	select {
	case c.ch <- value:
		return nil
	default:
		// Channel is full or no receiver
		c.ch <- value // This will block if unbuffered
		return nil
	}
}

// Receive receives a value from the channel
func (c *Channel) Receive() (VintObject, bool) {
	value, ok := <-c.ch
	return value, ok
}

// Close closes the channel
func (c *Channel) Close() {
	if !c.closed {
		close(c.ch)
		c.closed = true
	}
}

// IsClosed returns whether the channel is closed
func (c *Channel) IsClosed() bool {
	return c.closed
}

// NewChannel creates a new unbuffered channel
func NewChannel() *Channel {
	return &Channel{
		ch:       make(chan VintObject),
		closed:   false,
		buffered: false,
		size:     0,
	}
}

// NewBufferedChannel creates a new buffered channel with given size
func NewBufferedChannel(size int) *Channel {
	return &Channel{
		ch:       make(chan VintObject, size),
		closed:   false,
		buffered: true,
		size:     size,
	}
}
