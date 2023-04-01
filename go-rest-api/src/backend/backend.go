package backend

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	_ "github.com/mattn/go-sqlite3"
)

type App struct {
	DB   *sql.DB
	Port string
	Router *mux.Router
}

func (a *App) Initialize() {
	DB, err := sql.Open("sqlite3", "../../practiceit.db")
	if err != nil {
		log.Fatal(err.Error())
	}

	a.DB = DB
	a.Router = mux.NewRouter()

	a.initializeRoutes();

}

func (a *App) initializeRoutes(){
	a.Router.HandleFunc("/", getRepos).Methods("GET")
	a.Router.HandleFunc("/", postRepos).Methods("POST")
	a.Router.HandleFunc("/", deleteRepos).Methods("DELETE")
}

func (a *App) ProductRoutes(){
	a.Router.HandleFunc("/", getRepos).Methods("GET")
	a.Router.HandleFunc("/", postRepos).Methods("POST")
	a.Router.HandleFunc("/", deleteRepos).Methods("DELETE")
}

func (a *App) OrderRoutes(){
	a.Router.HandleFunc("/", getRepos).Methods("GET")
	a.Router.HandleFunc("/", postRepos).Methods("POST")
	a.Router.HandleFunc("/", deleteRepos).Methods("DELETE")
}

func (a *App) UserRoutes(){
	a.Router.HandleFunc("/", getRepos).Methods("GET")
	a.Router.HandleFunc("/", postRepos).Methods("POST")
	a.Router.HandleFunc("/", deleteRepos).Methods("DELETE")
}
func welcome(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello world!")
}

// GET
func getRepos(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "[GET] repos")
}

func postRepos(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "[POST] repos")
}

func deleteRepos(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "[DELETE] repos")
}
func (a *App) Run() {
	fmt.Println("listning on port 9003", a.Port)
	log.Fatal(http.ListenAndServe(a.Port, a.Router))
}
