// Package bucket sends your queued items to your callback function in chunks
package bucket

// Bucket collects items and sends them to the defined callback when the size is reached
type Bucket interface {
	// Len returns the number of remaining items in the queue
	Len() int
	// Calls returns number of calls the bucket have called
	Calls() int
	// Push adds new item in the queue
	Push(i interface{})
	// SetCallback replaces the callback function
	SetCallback(callback func([]interface{}))
}

type bucket struct {
	calls         int
	size          int
	items         []interface{}
	callback      func([]interface{})
	itemsQueue    chan interface{}
	callbackQueue chan []interface{}
}

func (c *bucket) init() {
	// I think there is no need to implement a loop breaker system
	go func() {
		for {
			var i = <-c.itemsQueue
			c.items = append(c.items, i)
			if len(c.items) < c.size {
				continue
			}
			var items = make([]interface{}, c.size)
			for i := 0; i < c.size; i++ {
				items[i] = c.items[i]
				c.items[i] = nil // in case user has stored a pointer
			}
			c.calls++
			c.items = c.items[c.size:]
			c.callbackQueue <- items
		}
	}()
	go func() {
		for {
			// better not to call this with goroutine
			// cause the order of item may change
			c.callback(<-c.callbackQueue)
		}
	}()
}

// Calls returns number of calls the bucket have called
func (c *bucket) Calls() int {
	return c.calls
}

// SetCallback replaces the callback function
func (c *bucket) SetCallback(callback func([]interface{})) {
	c.callback = callback
}

// Len returns the number of remaining items in the queue
func (c *bucket) Len() int {
	return len(c.items)
}

// Push adds new item in the queue
func (c *bucket) Push(i interface{}) {
	c.itemsQueue <- i
}

// New returns a fixed size bucket
func New(size int, callback func([]interface{})) Bucket {
	var c = bucket{
		size:          size,
		itemsQueue:    make(chan interface{}, 1000),
		callbackQueue: make(chan []interface{}, 10000),
		callback:      callback,
	}
	c.init()
	return &c
}
