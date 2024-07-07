package user

import (
	"admin-backend/service/auth"
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
	router.HandleFunc("/user/{email}", h.deleteUserByEmail).Methods("DELETE", "OPTIONS")
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
	if h.store.CheckIfUserExists(payload.Email) {
		http.Error(w, "Email has been used", http.StatusBadRequest)
	}
	payload.Password = auth.HashPassword(payload.Password)
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

func (h* Handler) deleteUserByEmail(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Access-Control-Allow-Methods", "DELETE") 
	w.Header().Set("Access-Control-Allow-Origin", "*")
	vars := mux.Vars(r)
	email, ok := vars["email"]
	if !ok{
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
	result, err := h.store.DeleteUserByEmail(email)
	if err != nil{
		http.Error(w, "Deletion failed", http.StatusInternalServerError)
	} else if !result {
		http.Error(w, "No user found with this email", http.StatusNotFound)
	}
}
