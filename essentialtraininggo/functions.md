go pass parameters to function by value

```go
package main

import (
 "fmt"
 "log"
 "math"
 "time"
)

func main() {
 nums := []int{1, 2, 3}
 doubleSlice(nums)
 fmt.Println(nums)
 val := 10
 doubleValue(val)
 fmt.Println(val) //10
 doublePtr(&val)  // 20
 fmt.Println(val)

 num, err := sqrt(25)
 if err != nil {
  fmt.Printf("ERROR: %s\n", err)
 } else {
  fmt.Println(num)
 }
}

func timeTracker(start time.Time, name string) {
 elapsed := time.Since(start)
 log.Printf("%s taken %s", name, elapsed)
}

func doubleSlice(values []int) {
 for i := range values {
  values[i] *= 2
 }
}

func doubleValue(num int) {
 num *= 2
}

func doublePtr(num *int) {
 *num *= 2
}

func sqrt(n float64) (float64, error) {
 if n < 0 {
  return 0.0, fmt.Errorf("sqrt of negative value (%f)", n)
 }

 return math.Sqrt(n), nil
}

// we don't use panics in go!

```

defer, somtimes we need to work with some resources and `release` them after finishing the task
this is a common mission to do so!
the code below just simulates that job

```go
package main

import "fmt"

func main() {
 worker()
}

func worker() {
 r1, err := acquire("A")
 if err != nil {
  fmt.Println("ERROR", err)
  return
 }
 defer release(r1)

 r2, err := acquire("B")
 if err != nil {
  fmt.Println("ERROR", err)
  return
 }
 defer release(r2)

 fmt.Println("worker!")

}
func acquire(name string) (string, error) {
 return name, nil
}

func release(name string) {
 fmt.Printf("Cleaning up %s\n", name)
}
```

defer are called in reverse order

TASK: get content-type of http get request

```go
func main() {
 ctype, err := getContentType("https://yousefmeska.tech/")
 if err != nil {
  fmt.Println(err.Error())
 }
 fmt.Println(ctype)
}

func getContentType(url string) (string, error) {
 r, err := http.Get(url)
 if err != nil {
  return "", err
 }
 defer r.Body.Close() // close the body
 ctype := r.Header.Get("Content-Type")
 if ctype == "" {
  return "", fmt.Errorf("can't find content type")
 }
 return ctype, nil
}
```



```go

func calculateDimensions(length, width, height int) (area int, volume int) {
	area = length * width * height
	volume = length * height
	return
}
```