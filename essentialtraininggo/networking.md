## JSON

```go
package main

import (
 "encoding/json"
 "fmt"
 "log"
 "os"
 "strings"
)

type Request struct {
 Login  string  `json:"user"`
 Type   string  `json:"type"`
 Amount float64 `json:"amount"`
}

const data = `
{
 "user": "yousef meska",
 "type": "deposit",
 "amount": 123.4
}`

func main() {
 reader := strings.NewReader(data) // simulate a file/socket
 // fmt.Println(reader.Len())
 //decode request
 dec := json.NewDecoder(reader)

 var req Request
 if err := dec.Decode(&req); err != nil {
  log.Fatalf("error: can't decode -%s", err)
 }

 fmt.Printf("got: %+v ", req)

 // Create response
 prevBalanc := 1_000_000.0
 resp := map[string]interface{}{
  "ok":      true,
  "balance": prevBalanc + req.Amount,
 }

 // Encode response
 enc := json.NewEncoder(os.Stdout)

 if err := enc.Encode(resp); err != nil {
  log.Fatalf("error: can't encode - %s", err)
 }

 fmt.Printf("got: %+v ", req)

}

```
