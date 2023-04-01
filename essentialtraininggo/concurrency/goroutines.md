```go
package main

import (
 "fmt"
 "net/http"
 "sync"
 "time"
)

func main() {
 var wg sync.WaitGroup
 urls := []string{
  "https://amazon.com",
  "https://facebook.com",
  "https://google.com",
  "https://netflix.com",
 }

 //serial
 start := time.Now()
 siteSerial(urls)
 elapsed := time.Since(start)
 fmt.Println(elapsed)

 //concurrent
 start = time.Now()
 concurrentSite(urls, &wg)
 fmt.Println(time.Since(start))

}

func GetContenType(url string) {
 resp, err := http.Get(url)
 if err != nil {
  fmt.Errorf("error", err)
  return
 }

 defer resp.Body.Close()

 ctype := resp.Header.Get("content-type")

 if ctype == "" {
  fmt.Println("content type is not found")
  return
 }
 fmt.Println(ctype)
}
func siteSerial(urls []string) {
 for _, url := range urls {
  GetContenType(url)
 }
}

// start a goroutine per url
func concurrentSite(urls []string, wg *sync.WaitGroup) {
 for _, url := range urls {
  wg.Add(1)
  go func(url string) {
   GetContenType(url)
   wg.Done()
  }(url)
 }

 wg.Wait()
}
```
