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
	router.HandleFunc("/user/login", h.userLogin).Methods("POST")
	router.HandleFunc("/user", h.updateUser).Methods("PUT", "OPTIONS")
	router.HandleFunc("/user/{username}", h.deleteUserByUsername).Methods("DELETE", "OPTIONS")
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
		return
	}
	if h.store.GetUserByUsername(payload.Username) != nil {
		http.Error(w, "Username has been used", http.StatusBadRequest)
		return
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
		return
	}
	err = h.store.UpdateUser(payload)
	if err != nil{
		http.Error(w, "Update failed", http.StatusInternalServerError)
	}
}

func (h* Handler) deleteUserByUsername(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Access-Control-Allow-Methods", "DELETE") 
	w.Header().Set("Access-Control-Allow-Origin", "*")
	vars := mux.Vars(r)
	username, ok := vars["username"]
	if !ok{
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	result, err := h.store.DeleteUserByUsername(username)
	if err != nil{
		http.Error(w, "Deletion failed", http.StatusInternalServerError)
	} else if !result {
		http.Error(w, "No user found with this username", http.StatusNotFound)
	}
}

func (h* Handler) userLogin(w http.ResponseWriter,  r *http.Request){
	w.Header().Set("Access-Control-Allow-Origin", "*")
	var payload types.User
	err := json.NewDecoder(r.Body).Decode(&payload)
	if err != nil{
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	user := h.store.GetUserByUsername(payload.Username)
	if user == nil {
		http.Error(w, "Username does not exist", http.StatusNotFound)
		return
	}
	if !auth.ComparePassword(user.Data()["password"].(string), []byte(payload.Password)) {
		http.Error(w, "Wrong password", http.StatusUnauthorized)
		return
	}
	json.NewEncoder(w).Encode("")
}
