package wastetype

import (
	"admin-backend/types"
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"google.golang.org/api/iterator"
)

type Handler struct {
	store types.WasteTypeStore
}

func NewHandler(store types.WasteTypeStore) *Handler {
	return &Handler{store: store}
}

func (h *Handler) RegisterRoutes(router *mux.Router){
	router.HandleFunc("/waste-type", h.handleCreate).Methods("POST")
	router.HandleFunc("/waste-type", h.handleUpdate).Methods("PUT", "OPTIONS")
	router.HandleFunc("/waste-type/{item}", h.handleDeleteItemByName).Methods("DELETE", "OPTIONS")
	router.HandleFunc("/waste-type", h.handleGetAll).Methods("GET", "OPTIONS")
	router.HandleFunc("/waste-type/{item}", h.handleGetByItemName).Methods("GET", "OPTIONS")
}

func (h *Handler) handleCreate(w http.ResponseWriter, r *http.Request) {
	var payload types.WasteType
	err := json.NewDecoder(r.Body).Decode(&payload)
	if err != nil{
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	err = h.store.Create(payload)
	if err != nil{
		http.Error(w, "Creation failed", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated) 
    w.Header().Set("Content-Type", "application/json") 
    response := map[string]string{"message": "Waste type created successfully"}
    json.NewEncoder(w).Encode(response) 
}

func (h *Handler) handleGetAll(w http.ResponseWriter, r *http.Request){
	iter := h.store.GetAll()
	var items []types.WasteType
	for {
		doc, err := iter.Next()
		if err == iterator.Done {
				break
		}
		var item types.WasteType
		doc.DataTo(&item)
		items = append(items, item)
	}
	json.NewEncoder(w).Encode(items)
}

func (h *Handler) handleGetByItemName(w http.ResponseWriter, r *http.Request){
	vars := mux.Vars(r)
	item, ok := vars["item"]
	if !ok {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	doc := h.store.GetAllByItem(item)
	if doc != nil{
		json.NewEncoder(w).Encode(doc.Data())
	} else{
		http.Error(w, "Waste type not found", http.StatusNotFound)
	}
}

func (h *Handler) handleDeleteItemByName(w http.ResponseWriter, r *http.Request){
	vars := mux.Vars(r)
	item, ok := vars["item"]
	if !ok {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	result, err := h.store.DeleteItemByName(item)
	if err != nil{
		http.Error(w, "Deletion failed", http.StatusInternalServerError)
		return
	} else if !result {
		http.Error(w, "Waste type not found", http.StatusNotFound)
		return
	}
	w.WriteHeader(http.StatusOK) 
    w.Header().Set("Content-Type", "application/json") 
    response := map[string]interface{}{
        "message": "Waste type deleted successfully",
    }
    json.NewEncoder(w).Encode(response) 
}

func (h *Handler) handleUpdate(w http.ResponseWriter, r *http.Request){
	var payload types.WasteType
	err := json.NewDecoder(r.Body).Decode(&payload)
	if err != nil{
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	err = h.store.Update(payload)
	if err != nil{
		http.Error(w, "Update failed", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK) 
    w.Header().Set("Content-Type", "application/json") 
    response := map[string]interface{}{
        "message": "Waste type updated successfully",
    }
    json.NewEncoder(w).Encode(response) 
}