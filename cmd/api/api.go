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
	apiRouter := router.PathPrefix("/api").Subrouter()
	
	wastetypeStore := wastetype.NewStore(s.db)
	wastetypeHandler := wastetype.NewHandler(wastetypeStore)
	wastetypeHandler.RegisterRoutes(apiRouter)
	
	wastetypeResponseStore := wastetypeResponse.NewStore(s.db)
	wastetypeResponseHandler := wastetypeResponse.NewHandler(wastetypeResponseStore)
	wastetypeResponseHandler.RegisterRoutes(apiRouter)

	userStore := user.NewStore(s.db)
	userHandler := user.NewHandler(userStore)
	userHandler.RegisterRoutes(apiRouter)

	corsRouter := corsMiddleware(router)
	return http.ListenAndServe(s.addr, corsRouter)
}

func corsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusOK)
			return
		}
		next.ServeHTTP(w, r)
	})
}
