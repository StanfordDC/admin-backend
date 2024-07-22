package wastetyperesponse

import (
	"admin-backend/types"
	"encoding/json"
	"net/http"
	"github.com/gorilla/mux"
	"google.golang.org/api/iterator"
	"time"
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
		objects := item.Items
		for _, obj := range objects{
			if obj.Feedback == 1 {
				items["goodResponse"]++
			} else if obj.Feedback == 2 {
				items["badResponse"]++
			}
			items["count"]++
		} 
	}
	json.NewEncoder(w).Encode(items)
}

func (h* Handler) handleGetHistory(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Access-Control-Allow-Origin", "*")
	var payload types.WasteTypeResponseRange
	err := json.NewDecoder(r.Body).Decode(&payload)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	startYear := payload.StartYear
	endYear := payload.EndYear
	startMonth := payload.StartMonth
	endMonth := payload.EndMonth
	iter := h.store.GetAll()
	totalMonths := (endYear - startYear) * 12
	totalMonths += endMonth - startMonth
	metrics := make([]types.WasteTypeResponseMetric, totalMonths)
	currentYear := startYear
	currentMonth := startMonth
	for i := 0; i < totalMonths; i++ {
        metrics[i] = types.WasteTypeResponseMetric{
            Year:    currentYear,
            Month:   currentMonth,
            Good:    0, // Initialize Good with 0 or other default value
            Bad:     0, // Initialize Bad with 0 or other default value
            Feature: 0, // Initialize Feature with 0 or other default value
        }
        // Increment month and adjust year if necessary
        currentMonth++
        if currentMonth > 12 {
            currentMonth = 1
            currentYear++
        }
    }
	indices := map[string]int{
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

		year := doc.CreateTime.Year()
		month := doc.CreateTime.Month()
		if year < startYear || year > endYear{
			continue
		}
		if year == startYear && month < time.Month(startMonth) || year == endYear && month > time.Month(endMonth){
			continue
		}
		index := (year - startYear) * 12 + startMonth - indices[month.String()]
		metrics[index].Feature++
		objects := item.Items
		for _, obj := range objects{
			if obj.Feedback == 1 {
				metrics[index].Good++
			} else if obj.Feedback == 2 {
				metrics[index].Bad++
			}
		} 
	}
	json.NewEncoder(w).Encode(metrics)
}
