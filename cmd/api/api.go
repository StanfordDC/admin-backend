package api

import (
	wastetype "admin-backend/service/wastetype"
	"net/http"

	"cloud.google.com/go/firestore"
	"github.com/gorilla/mux"
)

type APIServer struct {
	addr string
	db *firestore.Client
}

func NewAPIServer(addr string, db *firestore.Client) *APIServer {
	return &APIServer{
		addr: addr,
		db: db,
	}
}

func (s *APIServer) Run() error {
	router := mux.NewRouter()
	wastetypeStore := wastetype.NewStore(s.db)
	wastetypeHandler := wastetype.NewHandler(wastetypeStore)
	wastetypeHandler.RegisterRoutes(router)
	return http.ListenAndServe(s.addr, router)
}
