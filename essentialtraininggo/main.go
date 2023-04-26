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
