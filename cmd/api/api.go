package api

import (
	"admin-backend/service/user"
	wastetype "admin-backend/service/wastetype"
	wastetypeResponse "admin-backend/service/wastetypeResponse"
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
	
	wastetypeResponseStore := wastetypeResponse.NewStore(s.db)
	wastetypeResponseHandler := wastetypeResponse.NewHandler(wastetypeResponseStore)
	wastetypeResponseHandler.RegisterRoutes(router)

	userStore := user.NewStore(s.db)
	userHandler := user.NewHandler(userStore)
	userHandler.RegisterRoutes(router)
	return http.ListenAndServe(s.addr, router)
}
