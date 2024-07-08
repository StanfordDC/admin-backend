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
	router.HandleFunc("/responses/history", h.handleGetHistory).Methods("GET","OPTIONS")
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

func (h* Handler) handleGetHistory(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Access-Control-Allow-Origin", "*")
	iter := h.store.GetAll()
	months := []map[string]int{
		{"month": 1, "good": 0, "bad": 0, "feature": 0},
		{"month": 2, "good": 0, "bad": 0, "feature": 0},
		{"month": 3, "good": 0, "bad": 0, "feature": 0},
		{"month": 4, "good": 0, "bad": 0, "feature": 0},
		{"month": 5, "good": 0, "bad": 0, "feature": 0},
		{"month": 6, "good": 0, "bad": 0, "feature": 0},
		{"month": 7, "good": 0, "bad": 0, "feature": 0},
		{"month": 8, "good": 0, "bad": 0, "feature": 0},
		{"month": 9, "good": 0, "bad": 0, "feature": 0},
		{"month": 10, "good": 0, "bad": 0, "feature": 0},
		{"month": 11, "good": 0, "bad": 0, "feature": 0},
		{"month": 12, "good": 0, "bad": 0, "feature": 0},
	}
	index := map[string]int{
		"January":   0,
		"February":  1,
		"March":     2,
		"April":     3,
		"May":       4,
		"June":      5,
		"July":      6,
		"August":    7,
		"September": 8,
		"October":   9,
		"November":  10,
		"December":  11,
	}
	for {
		doc, err := iter.Next()
		if err == iterator.Done {
				break
		}
		var item types.WastetypeResponse
		doc.DataTo(&item)

		//Get the created month
		month := doc.CreateTime.Month().String()

		months[index[month]]["feature"]++
		objects := item.Objects
		for _, value := range objects{
			if value == 1 {
				months[index[month]]["good"]++
			} else if value == 2 {
				months[index[month]]["bad"]++
			}
		} 
	}
	json.NewEncoder(w).Encode(months)
}
