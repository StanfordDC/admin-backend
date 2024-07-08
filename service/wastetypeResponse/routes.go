package wastetyperesponse

import (
	"admin-backend/types"
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"google.golang.org/api/iterator"
)

type Handler struct {
	store types.WastetypeResponseStore
}

func NewHandler(store types.WastetypeResponseStore) *Handler {
	return &Handler{store: store}
}

func (h *Handler) RegisterRoutes(router *mux.Router){
	router.HandleFunc("/responses", h.handleGetAll).Methods("GET", "OPTIONS")
	router.HandleFunc("/responses/metrics", h.handleGetMetrics).Methods("GET", "OPTIONS")
}

func (h *Handler) handleGetAll(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Access-Control-Allow-Origin", "*")
	iter := h.store.GetAll()
	var items []types.WastetypeResponse
	for {
		doc, err := iter.Next()
		if err == iterator.Done {
				break
		}
		createTime := doc.CreateTime
		var item types.WastetypeResponse
		doc.DataTo(&item)
		item.CreateTime = createTime
		items = append(items, item)
	}
	json.NewEncoder(w).Encode(items)
}

func (h *Handler) handleGetMetrics(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Access-Control-Allow-Origin", "*")
	iter := h.store.GetAll()
	items := make(map[string]int)
	for {
		doc, err := iter.Next()
		if err == iterator.Done {
				break
		}
		items["feature"]++
		var item types.WastetypeResponse
		doc.DataTo(&item)
		objects := item.Objects
		for _, value := range objects{
			if value == 1 {
				items["goodResponse"]++
			} else if value == 2 {
				items["badResponse"]++
			}
			items["count"]++
		} 
	}
	json.NewEncoder(w).Encode(items)
}
