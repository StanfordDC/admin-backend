package user

import (
	"admin-backend/types"
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"google.golang.org/api/iterator"
)

type Handler struct {
	store types.UserStore
}

func NewHandler(store types.UserStore) *Handler {
	return &Handler{store: store}
}

func (h *Handler) RegisterRoutes(router *mux.Router){
	router.HandleFunc("/user", h.getAllUsers).Methods("GET","OPTIONS")
	router.HandleFunc("/user", h.createUser).Methods("POST")
	router.HandleFunc("/user", h.updateUser).Methods("PUT", "OPTIONS")
}

func (h* Handler) getAllUsers(w http.ResponseWriter,  r *http.Request){
	w.Header().Set("Access-Control-Allow-Origin", "*")
	iter := h.store.GetAllUsers()
	var users []types.User
	for {
		doc, err := iter.Next()
		if err == iterator.Done {
				break
		}
		var user types.User
		doc.DataTo(&user)
		users = append(users, user)
	}
	json.NewEncoder(w).Encode(users)
}

func (h* Handler) createUser(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Access-Control-Allow-Origin", "*")
	var payload types.User
	err := json.NewDecoder(r.Body).Decode(&payload)
	if err != nil{
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	err = h.store.CreateUser(payload)
	if err != nil{
		http.Error(w, "Creation failed", http.StatusInternalServerError)
	}
}

func (h* Handler) updateUser(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Access-Control-Allow-Methods", "PUT") 
	w.Header().Set("Access-Control-Allow-Origin", "*")
	if r.Method == "OPTIONS" {
		w.WriteHeader(http.StatusOK)
		return
	}
	var payload types.User
	err := json.NewDecoder(r.Body).Decode(&payload)
	if err != nil{
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	err = h.store.UpdateUser(payload)
	if err != nil{
		http.Error(w, "Update failed", http.StatusInternalServerError)
	}
}

