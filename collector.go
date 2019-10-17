package collector

type Collector interface {
	Len() int
	Append(i interface{})
	NewCallback(callback func([]interface{}))
}

type collector struct {
	calls         int
	count         int
	items         []interface{}
	callback      func([]interface{})
	itemsQueue    chan interface{}
	callbackQueue chan []interface{}
}

func (c *collector) init() {
	go func() {
		for {
			var i = <-c.itemsQueue
			if i != nil {
				c.items = append(c.items, i)
				if len(c.items) == c.count {
					var popped []interface{}
					for i := 0; i < c.count; i++ {
						popped = append(popped, c.items[i])
					}
					c.calls++
					c.items = c.items[c.count:]
					c.callbackQueue <- popped
				}
			}
		}
	}()
	go func() {
		for {
			var collection = <-c.callbackQueue
			c.callback(collection)
		}
	}()
}

func (c *collector) Calls() int {
	return c.calls
}

func (c *collector) NewCallback(callback func([]interface{})) {
	c.callback = callback
}

func (c *collector) Len() int {
	return len(c.items)
}

func (c *collector) Append(i interface{}) {
	c.itemsQueue <- i
}

func New(itemCount int, callback func([]interface{})) Collector {
	var c = collector{
		count:         itemCount,
		itemsQueue:    make(chan interface{}, 1000),
		callbackQueue: make(chan []interface{}, 10000),
		callback:      callback,
	}
	c.init()
	return &c
}
