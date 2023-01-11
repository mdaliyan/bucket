package bucket

import (
	"sync"
	"sync/atomic"
	"testing"
	"time"
)

func TestCountCallbackCalls(t *testing.T) {
	var n int64
	c, _ := New(BySize(100), func(items []interface{}) {
		atomic.AddInt64(&n, 1)
	})
	for i := 0; i < 2010; i++ {
		c.Push(i)
	}
	time.Sleep(time.Millisecond * 100)
	callbacksCount := atomic.LoadInt64(&n)
	if callbacksCount != 20 {
		t.Errorf("callback should have been called 20 times but it was called %d times", callbacksCount)
	}
	itemsLeftInBucket := c.Len()
	if itemsLeftInBucket != 10 {
		t.Errorf("there should be 10 items left in bucket but there are %d", itemsLeftInBucket)
	}
}

func TestConcurrentBucket(t *testing.T) {
	count := 1000
	size := 60
	var n int64
	c, _ := New(BySize(size), func(items []interface{}) {
		atomic.AddInt64(&n, 1)
	})
	var wg sync.WaitGroup
	for i := 0; i < count; i++ {
		wg.Add(1)
		go func() {
			for j := 0; j < size; j++ {
				c.Push(j)
			}
			wg.Done()
		}()
	}
	wg.Wait()
	time.Sleep(time.Millisecond * 100)
	if uint64(atomic.LoadInt64(&n)) != c.Calls() {
		t.FailNow()
	}
	if int(atomic.LoadInt64(&n)*int64(size))+c.Len() != count*size {
		t.FailNow()
	}
}

func TestOrder(t *testing.T) {

	var check, turn int64

	c, _ := New(BySize(100), func(items []interface{}) {
		for _, i := range items {
			atomic.AddInt64(&check, 1)
			if check != i.(int64) {
				t.FailNow()
			}
		}
	})
	for i := 0; i < 8313; i++ {
		for j := 1; j < 171; j++ {
			c.Push(atomic.AddInt64(&turn, 1))
		}
	}
}
