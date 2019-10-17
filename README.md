# Bucket
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
