package wastetyperesponse

import (
	"admin-backend/types"
	"encoding/json"
	"net/http"
	"fmt"
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
	months := map[string]map[string]int{
        "january":   {"good": 0, "bad": 0, "feature": 0},
        "february":  {"good": 0, "bad": 0, "feature": 0},
        "march":     {"good": 0, "bad": 0, "feature": 0},
        "april":     {"good": 0, "bad": 0, "feature": 0},
        "may":       {"good": 0, "bad": 0, "feature": 0},
        "june":      {"good": 0, "bad": 0, "feature": 0},
        "july":      {"good": 0, "bad": 0, "feature": 0},
        "august":    {"good": 0, "bad": 0, "feature": 0},
        "september": {"good": 0, "bad": 0, "feature": 0},
        "october":   {"good": 0, "bad": 0, "feature": 0},
        "november":  {"good": 0, "bad": 0, "feature": 0},
        "december":  {"good": 0, "bad": 0, "feature": 0},
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
		month = string(month[0]) + month[1:]
		month = fmt.Sprintf("%s%s", string(month[0]|0x20), month[1:])

		months[month]["feature"]++
		objects := item.Objects
		for _, value := range objects{
			if value == 1 {
				months[month]["good"]++
			} else if value == 2 {
				months[month]["bad"]++
			}
		} 
	}
	json.NewEncoder(w).Encode(months)
}
