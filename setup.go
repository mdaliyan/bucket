package bucket

import (
	`errors`
	`time`
)

type Setup func(c *bucket) error

func BySize(size int) Setup {
	return func(c *bucket) error {
		if size < 1 {
			return errors.New("bucket size cannot be lt 1")
		}
		c.size = size
		go func() {
			for {
				select {
				case <-c.closed:
					return
				case i := <-c.itemsQueue:
					c.write.Lock()
					c.items = append(c.items, i)
					c.write.Unlock()
					if c.Len() >= c.size {
						c.callbackQueue <- c.pop(c.size)
					}
				}
			}
		}()
		return nil
	}
}

func ByTime(every time.Duration) Setup {
	return func(c *bucket) error {
		if every < 1 {
			return errors.New("call interval cannot be lt 1")
		}
		c.ticker = time.NewTicker(every)
		go func() {
			for {
				select {
				case <-c.closed:
					return
				case i := <-c.itemsQueue:
					c.write.Lock()
					c.items = append(c.items, i)
					c.write.Unlock()
				case <-c.ticker.C:
					c.callbackQueue <- c.pop(c.Len())
				}
			}
		}()
		return nil
	}
}
