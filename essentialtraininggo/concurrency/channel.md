this will make a deadlock!

```go
package main

import "fmt"

func main() {
 ch := make(chan int)
 ch <- 33

 val := <-ch
 fmt.Println(val)
}

```

the goroutin will run concurrently with main one, so it will send through channel and at the same time the main is receiving!

```go
package main

import "fmt"

func main() {
 ch := make(chan int)

 go func() {
  ch <- 33
 }()

 val := <-ch
 fmt.Println(val)
}

```

1 and 2 will run concurrently to each other, each time one is sending, receiving is blocked until receiving is ready!
and each time 2 is receiving, 1 is blocked and vise versa.

```go
func main() {
 // send multiple
 ch := make(chan int)
 const count = 3
 go func() { // 1
  for i := 0; i < count; i++ {
   fmt.Printf("sending %d\n", i)
   ch <- i
   time.Sleep(time.Second)
  }
 }()

 for i := 0; i < count; i++ { // 2
  val := <-ch
  fmt.Printf("received %d\n", val)
 }
}
```

sending 0
sending 1
received 0
received 1
sending 2
received 2

A common thing to do, is to close the channel signaling the work is done!

```go

func main() {
 // closing, we're not receiving, so we must close the channel
 ch := make(chan int)
 go func() {
  for i := 0; i < 3; i++ {
   fmt.Printf("sending %d\n", i)
   ch <- i
   time.Sleep(time.Second)
  }

  close(ch)
 }()

 for i := range ch {
  fmt.Printf("received %d\n", i)
 }
}
```

buffered channel

```go
ch := make(chan int, 1)
 ch <- 19
 val := <-ch
 fmt.Printf("value: %d", val) // not goig to block as the unbuffered channel
 val2 := <-ch
 fmt.Printf("value: %d", val2) // deadlock, only one reciever!
```
