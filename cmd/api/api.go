package api

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/NayanPahuja/ecom/service/product"
	"github.com/NayanPahuja/ecom/service/user"
	"github.com/gorilla/mux"
)

/*
This is a type API Server kind of an abstract type for making
a new server with these specifications.
address -> string : Address of the server
database -> sql.DB : Database used by the server
*/
type APIServer struct {
	address  string
	database *sql.DB
}

/*
This is a func which takes thr address and database, returns us a server
of our abstract type APIServer and initializes address and database
with command line param addr and db
*/
func NewAPIServer(addr string, db *sql.DB) *APIServer {
	return &APIServer{
		address:  addr,
		database: db,
	}
}

func (s *APIServer) Run() error {
	router := mux.NewRouter()
	subrouter := router.PathPrefix("/api/v1").Subrouter()

	userStore := user.NewStore(s.database)
	userHandler := user.NewHandler(userStore)
	userHandler.RegisterRoutes(subrouter)
	productStore := product.NewStore(s.database)
	productHandler := product.NewHandler(productStore, userStore)
	productHandler.RegisterRoutes(subrouter)
	log.Println("Listening on: ", s.address)
	return http.ListenAndServe(s.address, router)
}
