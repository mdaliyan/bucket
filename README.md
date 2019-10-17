# Bucket
[![Go Report Card](https://goreportcard.com/badge/github.com/mdaliyan/bucket)](https://goreportcard.com/report/github.com/mdaliyan/bucket)
[![Go Walker](http://gowalker.org/api/v1/badge)](https://gowalker.org/github.com/mdaliyan/bucket) 
[![godoc](https://godoc.org/github.com/mdaliyan/bucket.svg?status.svg)](https://godoc.org/github.com/mdaliyan/bucket)
[![License](http://img.shields.io/badge/license-mit-blue.svg?style=flat)](https://raw.githubusercontent.com/labstack/echo/master/LICENSE)

`bucket` sends your queued items to your callback function in chunks 

### Installation

```bash
go get github.com/mdaliyan/bucket
```

### Usage

```go
var callback = func(items []interface{}) {
    fmt.Println(items)
}

var bucket = bucket.New(10, callback)

for i:=0; i < 25; i++ {
    bucket.Push(i)
}

time.Sleep(time.Millisecond)

fmt.Println(bucket.Len())
```
this Prints
```
[0 1 2 3 4 5 6 7 8 9]
[10 11 12 13 14 15 16 17 18 19]
5
```
