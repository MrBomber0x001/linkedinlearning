```go
package main

import (
 "fmt"
 "log"
 "os"

 "github.com/pkg/errors"
)

type FileConfig struct {
}

func readConfig(filepath string) (*FileConfig, error) {
 file, err := os.Open(filepath)
 if err != nil {
  return nil, errors.Wrap(err, "can't open configuration file")
 }
 defer file.Close()

 cfg := &FileConfig{}
 return cfg, nil
}
func setupLogging() {
 out, err := os.OpenFile("app.log", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
 if err != nil {
  return
 }
 log.SetOutput(out)
}
func main() {
 setupLogging()
 cfg, err := readConfig("/path/to/config.tml")
 if err != nil {
  fmt.Fprintf(os.Stderr, "error: %s\n", err)
  log.Printf("error: %+v", err) // will use setupLogging();
  os.Exit(1)
 }

 fmt.Println(cfg)

}
```

using panic <Don't use panic>
the way we go againse panic is inside defer

recover returns nil if there's no panic, and return something if there's a panic

```go
package main

import "fmt"

func main() {
 vals := []int{1, 2, 3}
 v, err := safeValues(vals, 10)
 if err != nil {
  fmt.Println("ERROR:", err)
 }

 fmt.Println("v:", v)
}

func safeValues(vals []int, index int) (n int, err error) {
 defer func() {
  if e := recover(); e != nil {
   err = fmt.Errorf("%v", e)
  }
 }()
 return vals[index], nil
}

```
