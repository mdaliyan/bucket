package collector

import (
	"sync"
	"sync/atomic"
	"testing"
	"time"
)

func TestCollector(t *testing.T) {
	c := New(20, func(items []interface{}) {})
	var n int64
	c.NewCallback(func(items []interface{}) {
		atomic.AddInt64(&n, 1)
	})
	for i := 0; i < 2010; i++ {
		c.Append(i)
	}
	time.Sleep(time.Millisecond * 50)
	if int(n*20)+c.Len() != 2010 {
		t.FailNow()
	}
}

func TestConcurrentCollector(t *testing.T) {
	c := New(70, func(items []interface{}) {})
	var n int64
	c.NewCallback(func(items []interface{}) {
		atomic.AddInt64(&n, 1)
	})
	var wg sync.WaitGroup
	for i := 0; i < 8313; i++ {
		wg.Add(1)
		go func() {
			for j := 0; j < 170; j++ {
				c.Append(j)
			}
			wg.Done()
		}()
	}
	wg.Wait()
	time.Sleep(time.Millisecond * 50)
	if int(n*70)+c.Len() != 8313*170 {
		t.FailNow()
	}
}
