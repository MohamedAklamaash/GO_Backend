package api

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/MohamedAklamaash/GO_Simple_CRUD_Backend_WIth_SQL/services/user"
	"github.com/gorilla/mux"
)

type APIServer struct {
	addr string
	db   *sql.DB
}

func NewAPIServer(addr string, db *sql.DB) *APIServer {

	return &APIServer{
		addr: addr,
		db:   db,
	}
}

func (s *APIServer) Run() error {
	router := mux.NewRouter()
	subrouter := router.PathPrefix("/api/v1").Subrouter()
	userHandler := user.NewHandler()
	userHandler.RegisterRoutes(subrouter)
	log.Println("\nRunning on", s.addr)
	return http.ListenAndServe(s.addr, router)
}
