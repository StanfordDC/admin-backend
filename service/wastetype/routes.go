package wastetype

import (
	"admin-backend/types"
	"encoding/json"
	"fmt"
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
	router.HandleFunc("/waste-type/create", h.handleCreate).Methods("POST")
	router.HandleFunc("/waste-type/{item}", h.handleDeleteItemByName).Methods("DELETE")
	router.HandleFunc("/waste-type", h.handleGetAll).Methods("GET", "OPTIONS")
	router.HandleFunc("/waste-type/{item}", h.handleGetByItemName).Methods("GET", "OPTIONS")
}

func (h *Handler) handleCreate(w http.ResponseWriter, r *http.Request) {
	var payload types.WasteType
	err := json.NewDecoder(r.Body).Decode(&payload)
	if err != nil{
		fmt.Println(err.Error())
	}
	h.store.Create(payload)
}

func (h* Handler) handleGetAll(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Access-Control-Allow-Origin", "*")
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

func (h* Handler) handleGetByItemName(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Access-Control-Allow-Origin", "*")
	vars := mux.Vars(r)
	item, ok := vars["item"]
	if !ok {
		fmt.Println("Something is wrong")
	}
	doc := h.store.GetAllByItem(item)
	if doc != nil{
		json.NewEncoder(w).Encode(doc.Data())
	} else{
		http.Error(w, "Waste type not found", http.StatusNotFound)
	}
}

func (h* Handler) handleDeleteItemByName(w http.ResponseWriter, r *http.Request){
	vars := mux.Vars(r)
	item, ok := vars["item"]
	if !ok {
		fmt.Println("Something is wrong")
	}
	h.store.DeleteItemByName(item)
}