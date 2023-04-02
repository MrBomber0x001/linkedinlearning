## context

is used for cancellation and deadlines

ctx.Done() => returns a channel

everytime we create a context, we create it from another context!

the parent of all contexts is `context.Background()`

so we do a context with timeout starting from the context background and and give it a timeout of 100 milliseconds

```go
ctx, cancel := context.WithTimeout(context.Background(), 100*time.Millisecond)
```

in the next context, we create much longer one

```go
package main

import (
 "context"
 "fmt"
 "time"
)

type Bid struct {
 AdURL string
 Price float64
}

func bestBid(url string) Bid {
 time.Sleep(20 * time.Millisecond) // simulate running time of an algorithms
 return Bid{
  AdURL: "https://anything.com",
  Price: 2.9,
 }
}

var defaultBid = Bid{
 AdURL: "https://anything.com/default",
 Price: 12,
}

func findBid(ctx context.Context, url string) Bid {
 ch := make(chan Bid, 1) // buffered channel to avoid goroutine leak!
 go func() {
  ch <- bestBid(url)
 }()

 select {
 case bid := <-ch:
  return bid
 case <-ctx.Done():
  return defaultBid
 }
}
func main() {
 ctx, cancel := context.WithTimeout(context.Background(), 100*time.Millisecond)
 defer cancel()

 url := "https://http.cat/418"
 bid := findBid(ctx, url)
 fmt.Println(bid)

 ctx, cancel = context.WithTimeout(context.Background(), 10*time.Millisecond)
 defer cancel()
 url = "http://http.cat/404"
 bid = findBid(ctx, url)
 fmt.Println(bid)

}
```
