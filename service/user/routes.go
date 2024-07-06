package user

import (
	"admin-backend/types"
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"google.golang.org/api/iterator"
)

type Handler struct {
	store types.User
}

func NewHandler(store types.User) *Handler {
	return &Handler{store: store}
}

func (h *Handler) RegisterRoutes(router *mux.Router){
	router.HandleFunc("/user", h.getAllUser).Methods("GET","OPTIONS")
	router.HandleFunc("/user/{id}", h.getUserById).Methods("GET","OPTIONS")
	router.HandleFunc("/user", h.createUser).Methods("POST", "OPTIONS")
	router.HandleFunc("/user", h.updateUser).Methods("PUT", "OPTIONS")
}
