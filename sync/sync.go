package sync

import "sync"

// Counter is a simple counter.
type Counter struct {
	mu    sync.Mutex
	value int
}

// NewCounter creates a new Counter.
func NewCounter() *Counter {
	return &Counter{}
}

// Inc increments the counter.
func (c *Counter) Inc() {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.value++
}

// Value returns the counter's current value.
func (c *Counter) Value() int {
	return c.value
}
