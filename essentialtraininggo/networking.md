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

## HTTP REQUEST AND RESPONSE

```go
package main

import (
 "encoding/json"
 "fmt"
 "log"
 "net/http"
 "strings"
)

func healthHandler(w http.ResponseWriter, r *http.Request) {
 fmt.Fprintln(w, "OK!")
}

type MathResponse struct {
 Error  string  `json:"error"`
 Result float64 `json:"result"`
}

type MathRequest struct {
 Op    string  `json:"op"`
 Left  float64 `json:"left"`
 Right float64 `json:"right"`
}

func MathHandler(w http.ResponseWriter, r *http.Request) {
 //Step 1: Decode & Validate [Request] from json into struct
 defer r.Body.Close()
 dec := json.NewDecoder(r.Body)
 req := &MathRequest{}
 //Decode json
 if err := dec.Decode(req); err != nil {
  log.Printf("error: badJSON: %s", err)
  http.Error(w, "bad json", http.StatusBadRequest)
  return
 }
 //validate
 if !strings.Contains("+-*/", req.Op) {
  log.Printf("error: bad operator: %q", req.Op)
  http.Error(w, "unknown operator", http.StatusBadRequest)
  return
 }

 // Step 2: work
 resp := &MathResponse{}
 switch req.Op {
 case "+":
  resp.Result = req.Left + req.Right
 case "-":
  resp.Result = req.Left - req.Right
 case "*":
  resp.Result = req.Left * req.Right
 case "/":
  if req.Right == 0.0 {
   resp.Error = "division by 9"
  } else {
   resp.Result = req.Left / req.Right
  }
 default:
  resp.Error = fmt.Sprintf("unknown error %s", req.Op)
 }

 // Step 3: Encode result into json
 w.Header().Set("Content-Type", "application/json")
 if resp.Error != "" {
  w.WriteHeader(http.StatusBadRequest)
 }

 enc := json.NewEncoder(w)
 if err := enc.Encode(resp); err != nil {
  log.Printf("can't encode %v -%s", resp, err)
 }
}

func main() {
 http.HandleFunc("/health", healthHandler)
 http.HandleFunc("/math", MathHandler)

 addr := ":8080"
 if err := http.ListenAndServe(addr, nil); err != nil {
  log.Fatal(err.Error())
 }
}

```

TASK

```go
package main

import "net/http"

type Server struct {
 db DB
}

// POST /key
// GET /<key>
func (s *Server) Handler(w http.ResponseWriter, r *http.Request) {
 //TODO
 switch r.Method {
 case http.MethodGet:
  s.getHandler(w, r)
  return
 case http.MethodPost:
  s.postHandler(w, r)
  return
 }
 http.Error(w, "bad request", http.StatusMethodNotAllowed)
}

func (s *Server) getHandler(w http.ResponseWriter, r *http.Request) {

}

func (s *Server) postHandler(w http.ResponseWriter, r *http.Request) {

}
func main() {

}

```
