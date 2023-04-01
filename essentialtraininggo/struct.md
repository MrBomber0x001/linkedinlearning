you can only have reciever methods on types defined in the same package or module!

```go
package main

import (
 "fmt"
 "time"
)

type Budget struct {
 Balance   float64
 CompainID string
 ExpireAt  time.Time
}

// methods (function receiver)
func (b Budget) TimeLeft() time.Duration {
 return b.ExpireAt.Sub(time.Now().UTC())
}

// method (pointer receiver)
func (b *Budget) Update(sum float64) {
 b.Balance += sum
}
func main() {
 b1 := Budget{12.2, "Kitten", time.Now().Add(7 * 24 * time.Hour)}
 b1.Update(90)
 fmt.Println(b1.TimeLeft())
 fmt.Println(b1)
 fmt.Printf("%#v", b1)

 fmt.Println(b1.CompainID)

 b2 := Budget{
  Balance:   1,
  CompainID: "puppoes",
 }
 fmt.Println(b2)
 var b3 Budget
 fmt.Printf("%#v", b3)
}
```

we used on object oriented programming languages to have a contructor

```py
class Point:
        def __init__(self, x, y):
                self.x = x
                self.y = y
```

on Go we can simulate this behaviour, with new function that return nwe object and possible error

```go
package main

import (
 "fmt"
 "time"
)

type Budget struct {
 Balance   float64
 CompainID string
 ExpiresAt time.Time
}

func main() {
 expires := time.Now().Add(24 * 7 * time.Hour)
 b1, err := NewBudget(34.5, "puppies", expires)
 if err != nil {
  fmt.Println("ERROR", err)
 } else {
  fmt.Println(b1)
 }
}

func NewBudget(balance float64, compainid string, expiresat time.Time) (*Budget, error) {
 // do the validation neccessary first
 if compainid == "" {
  return nil, fmt.Errorf("empty CompainID")
 }
 if balance <= 0 {
  return nil, fmt.Errorf("balance must be bigger than 0")
 }

 if expiresat.Before(time.Now()) {
  return nil, fmt.Errorf("bad expiration time")
 }

 b := Budget{
  CompainID: compainid,
  Balance:   balance,
  ExpiresAt: expiresat,
 }

 return &b, nil // the compiler does here an escape anaylsis
}

```

TASK:

```go
package main

import "fmt"

type Square struct {
 X      int
 Y      int
 Length int
}

func NewSquare(x int, y int, length int) (*Square, error) {
 // validate the input!
 if x < 0 || y < 0 || length < 0 {
  return nil, fmt.Errorf("can't create with negative number")
 }

 sq := Square{
  X:      x,
  Y:      y,
  Length: length,
 }
 return &sq, nil
}

//Move: moves the square
func (s *Square) Move(dx int, dy int) {
 // validate the input!
 s.X += dx
 s.Y += dy
}

//Area: return the area of the square!

func (s *Square) Area() int {
 return s.Length * s.Length
}

func main() {
 s, err := NewSquare(1, 1, 10)
 if err != nil {
  fmt.Println("ERROR", err.Error())
 }
 s.Move(2, 3)
 fmt.Printf("%+v\n", s)
 fmt.Printf("The area of the square: %d", s.Area())

}

```
