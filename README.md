# Bucket

![example workflow](https://github.com/mdaliyan/bucket/actions/workflows/test.yml/badge.svg)
[![Coverage Status](https://coveralls.io/repos/github/mdaliyan/bucket/badge.svg?branch=master)](https://coveralls.io/github/mdaliyan/bucket?branch=master)
[![Go Report Card](https://goreportcard.com/badge/github.com/mdaliyan/bucket/v2)](https://goreportcard.com/report/github.com/mdaliyan/bucket/v2)
[![Go Reference](https://pkg.go.dev/badge/github.com/mdaliyan/bucket/v2.svg)](https://pkg.go.dev/github.com/mdaliyan/bucket/v2)
[![License](http://img.shields.io/badge/license-mit-blue.svg?style=flat)](https://raw.githubusercontent.com/labstack/echo/master/LICENSE)

bucket queues your items and sends them to your callback function in chunks.

### Installation

```bash
go get github.com/mdaliyan/bucket
```

### Usage

```go
callback := func(items []interface{}) {
    fmt.Println(items)
}

b, _ := bucket.New(bucket.BySize(10), callback)

for i := 0; i < 25; i++ {
    b.Push(i)
}

time.Sleep(time.Microsecond * 100)

fmt.Println(b.Len())
```
this Prints
```
[0 1 2 3 4 5 6 7 8 9]
[10 11 12 13 14 15 16 17 18 19]
5
```
