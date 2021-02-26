package bucket

import (
	"sync"
	"sync/atomic"
	"testing"
	"time"
)

func doNothingWithItems(_ []interface{}) {
	// Do nothing because of testing purpose
}

func TestBucket(t *testing.T) {
	c, _ := New(BySize(20), doNothingWithItems)
	var n int64
	c.SetCallback(func(items []interface{}) {
		atomic.AddInt64(&n, 1)
	})
	for i := 0; i < 2010; i++ {
		c.Push(i)
	}
	time.Sleep(time.Millisecond * 50)
	if int(atomic.LoadInt64(&n)*20)+c.Len() != 2010 {
		t.FailNow()
	}
}

func TestConcurrentBucket(t *testing.T) {
	count := 1000
	size := 60
	c, _ := New(BySize(size), doNothingWithItems)
	var n int64
	c.SetCallback(func(items []interface{}) {
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
	time.Sleep(time.Millisecond * 10)
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
