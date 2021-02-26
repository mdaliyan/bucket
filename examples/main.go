package main

import (
	`fmt`
	`time`

	"github.com/mdaliyan/bucket"
)

func main() {
	callback := func(items []interface{}) {
		fmt.Println(items)
	}

	b, _ := bucket.New(bucket.BySize(10), callback)

	for i := 0; i < 25; i++ {
		b.Push(i)
	}

	time.Sleep(time.Microsecond * 100)

	fmt.Println(b.Len())
}
