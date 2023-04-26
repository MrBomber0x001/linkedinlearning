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

Working with Timeout and Size limits

```go
package main

import (
 "context"
 "io"
 "log"
 "net/http"
 "os"
 "time"
)

//TODO: install NerdFonts

func main() {
 ctx, cancel := context.WithTimeout(context.Background(), 30*time.Millisecond)
 defer cancel()

 req, err := http.NewRequestWithContext(ctx, http.MethodGet, "https://httpbin.org/ip", nil)
 if err != nil {
  log.Fatal(err)
 }

 resp, err := http.DefaultClient.Do(req)
 if err != nil {
  log.Fatal(err)
 }

 defer resp.Body.Close()

 //define a single megabyte
 const mb = 1 << 20
 r := io.LimitReader(resp.Body, mb)
 io.Copy(os.Stdout, r)
}




```

GITHUB API TASK

```go
package main

//TODO: use /pkg/errors and errors.Wrap() to capture stack trace!
import (
 "encoding/json"
 "fmt"
 "log"
 "net/http"
 "net/url"
 "os"
)

// Print the errors with their stack trace!
func setupLoggin() {
 out, err := os.OpenFile("github_api.log", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
 if err != nil {
  return
 }

 log.SetOutput(out)
}

type User struct {
 Login    string `json:"login"`
 Name     string `json:"name"`
 NumRepos int    `json:"public_repos"`
}

// const (
//  ErrUserNotFound   error.New("User not found!")
//  ErrInvalidRequest error.New("Invalid Request")
// )

//UserInfo: return information of the user!
func UserInfo(login string) (*User, error) {
 // call the github API for a given login
 // and return the user struct
 if login == "" {
  return nil, fmt.Errorf("error - please pass a username")
 }
 endpoint := fmt.Sprintf("https://api.github.com/users/%s", url.PathEscape(login))
 resp, err := http.Get(endpoint)

 if err != nil {
  return nil, fmt.Errorf("error making request - %s", err.Error())
 }
 defer resp.Body.Close()

 if resp.StatusCode != http.StatusOK {
  return nil, fmt.Errorf("error making request - %s", err.Error())
 }

 // Decode JSON: Should be placed on it's own function
 user := User{Login: login}
 dec := json.NewDecoder(resp.Body)
 if err := dec.Decode(&user); err != nil {
  return nil, err
 }

 return &user, nil
}

//return slice of informations
// func GetUserFollowersInformation(login string) (*User, error) {
//  return nil, nil
// }

func main() {

 // if len(os.Args) < 2 {
 //  log.Fatal("Please provide argumentas!")
 // }
 setupLoggin()
 user, err := UserInfo("mrbomber0x001")
 if err != nil {
  fmt.Fprintf(os.Stderr, "error: %s\n", err)
  log.Printf("error: %+v", err) // will use setupLogging();
  os.Exit(1)
 }
 fmt.Println(user, err)
}
```

```go
package main

import (
 "fmt"
 "io"
 "log"
 "net/http"
)

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
 key := r.URL.Path[1:] //time leading /
 data := s.db.Get(key)
 if data == nil {
  http.Error(w, "not found", http.StatusNotFound)
  return
 }
 w.Write(data)
}

func (s *Server) postHandler(w http.ResponseWriter, r *http.Request) {
 key := r.URL.Path[1:]
 defer r.Body.Close()
 rdr := io.LimitReader(r.Body, 1<<10)
 data, err := io.ReadAll(rdr)

 if err != nil {
  http.Error(w, "can't read", http.StatusBadRequest)
  return
 }
 s.db.Set(key, data)
 fmt.Fprintf(w, "%s set\n", key)
}
func main() {
 var s Server
 http.HandleFunc("/", s.Handler)
 if err := http.ListenAndServe(":8080", nil); err != nil {
  log.Fatal(err)
 }
}


```
