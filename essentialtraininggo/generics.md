an interface can also be a set of types!

```go
package main

import "fmt"

type Ordered interface {
 int | float64 | string
}

func min[T Ordered](values []T) (T, error) {
 if len(values) == 0 {
  var zero T
  return zero, fmt.Errorf("min of empty slice")
 }

 m := values[0]
 for _, v := range values[:1] {
  if v < m {
   m = v
  }
 }
 return m, nil
}
func main() {
 fmt.Println(min([]int{1, 2, 3}))
 fmt.Println(min([]float64{12, 24, 53}))
 fmt.Println(min([]string{"A", "B", "C"}))
}
```
