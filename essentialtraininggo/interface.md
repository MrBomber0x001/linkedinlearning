An interface defines a set of method that a <type> should implement
any type!

```go
package main

import (
 "fmt"
 "math"
)

type Shape interface {
 Area() float64
}
type Circle struct {
 Radius float64
}

type Square struct {
 Length float64
}

func (c Circle) Area() float64 {
 return math.Pi * c.Radius * c.Radius
}

func (s Square) Area() float64 {
 return s.Length * s.Length
}

func main() {
 s := Square{10}
 c := Circle{20}

 shapes := []Shape{s, c}
 sa := sumAreas(shapes)
 fmt.Println(sa)
}

//sumAreas: return sum of all areas in the slice
func sumAreas(shapes []Shape) float64 {
 total := 0.0
 for _, shape := range shapes {
  fmt.Printf("The area of %s is %f\n", shape, shape.Area())
  total += shape.Area()
 }
 return total
}
```

using interface for example `Reader.IO` interface can enable you to work with multiple and different readers sources like
files, sockets, digital signatures!

TASK:
implement IO writer!

```go
package main

import (
 "fmt"
 "io"
 "os"
)

// implements io.Writer and turns everything into uppercase!

type Capper struct {
 wtr io.Writer
}

func main() {
 c := &Capper{os.Stdout}
 fmt.Fprintf(c, "This is going to be printed!")
}

func (c *Capper) Write(p []byte) (n int, err error) {
 diff := byte('a' - 'A')

 out := make([]byte, len(p)) // turn into slice of bytes
 for i, c := range p {
  if c >= 'a' && c <= 'z' {
   c -= diff
  }
  out[i] = c
 }

 return c.wtr.Write(out)
}
```
