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
	c := New(20, doNothingWithItems)
	var n int64
	c.SetCallback(func(items []interface{}) {
		atomic.AddInt64(&n, 1)
	})
	for i := 0; i < 2010; i++ {
		c.Push(i)
	}
	time.Sleep(time.Millisecond * 50)
	if int(n*20)+c.Len() != 2010 {
		t.FailNow()
	}
}

func TestConcurrentBucket(t *testing.T) {
	c := New(70, doNothingWithItems)
	var n int64
	c.SetCallback(func(items []interface{}) {
		atomic.AddInt64(&n, 1)
	})
	var wg sync.WaitGroup
	for i := 0; i < 8313; i++ {
		wg.Add(1)
		go func() {
			for j := 0; j < 170; j++ {
				c.Push(j)
			}
			wg.Done()
		}()
	}
	wg.Wait()
	time.Sleep(time.Millisecond * 10)
	if int(n) != c.Calls() {
		t.FailNow()
	}
	if int(n*70)+c.Len() != 8313*170 {
		t.FailNow()
	}
}

func TestOrder(t *testing.T) {

	var check, turn int64

	c := New(100, func(items []interface{}) {
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
