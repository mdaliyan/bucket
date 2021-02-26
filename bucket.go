// Package bucket sends your queued items to your callback function in chunks
package bucket

import (
	`sync`
	`sync/atomic`
	`time`
)

// Bucket collects items and sends them to the defined callback when the size is reached
type Bucket interface {
	// Len returns the number of remaining items in the queue
	Len() int
	// Calls returns number of calls the bucket have called
	Calls() uint64
	// Push adds new item in the queue
	Push(i interface{})
	// SetCallback replaces the callback function
	SetCallback(callback func([]interface{}))
}

type bucket struct {
	calls         uint64
	size          int
	ticker        *time.Ticker
	closed        chan bool
	items         []interface{}
	callback      func([]interface{})
	itemsQueue    chan interface{}
	callbackQueue chan []interface{}
	write         sync.RWMutex
}

// Calls returns number of calls the bucket have called
func (c *bucket) Calls() uint64 {
	return atomic.LoadUint64(&c.calls)
}

// SetCallback replaces the callback function
func (c *bucket) SetCallback(callback func([]interface{})) {
	c.callback = callback
}

// Len returns the number of remaining items in the queue
func (c *bucket) Len() int {
	c.write.RLock()
	size := len(c.items)
	c.write.RUnlock()
	return size
}

// Push adds new item in the queue
func (c *bucket) Push(i interface{}) {
	c.itemsQueue <- i
}

func (c *bucket) pop(size int) []interface{} {
	c.write.Lock()
	items := c.items[:size]
	c.items = c.items[c.size:]
	c.write.Unlock()
	return items
}

// New returns a fixed size bucket
func New(setup Setup, callback func([]interface{})) (Bucket, error) {
	var c = &bucket{
		itemsQueue:    make(chan interface{}, 1000),
		callbackQueue: make(chan []interface{}, 10000),
		callback:      callback,
	}
	err := setup(c)
	if err != nil {
		return nil, err
	}
	go func() {
		for {
			// better not to call this with goroutine
			// cause the order of the buckets might change
			c.callback(<-c.callbackQueue)
			atomic.AddUint64(&c.calls, 1)
		}
	}()
	return c, nil
}
