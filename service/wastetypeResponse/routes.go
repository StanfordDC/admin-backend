package wastetyperesponse

import (
	"admin-backend/types"
	"encoding/json"
	"net/http"
	"github.com/gorilla/mux"
	"google.golang.org/api/iterator"
	"strconv"
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
	router.HandleFunc("/responses/history/startYear={startYear}&startMonth={startMonth}&endYear={endYear}&endMonth={endMonth}", h.handleGetHistory).Methods("GET", "OPTIONS")
}

func (h *Handler) handleGetAll(w http.ResponseWriter, r *http.Request){
	// w.Header().Set("Access-Control-Allow-Origin", "*")
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
	// w.Header().Set("Access-Control-Allow-Origin", "*")
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
	// w.Header().Set("Access-Control-Allow-Origin", "*")
	vars := mux.Vars(r)
	payload := parsePathVariables(vars)
	iter := h.store.GetAll()
	totalMonths := getTotalMonths(payload)
	metrics := initializeMetrics(payload, totalMonths)
	for {
		doc, err := iter.Next()
		if err == iterator.Done {
				break
		}
		var item types.WastetypeResponse
		doc.DataTo(&item)

		year := doc.CreateTime.Year()
		month := doc.CreateTime.Month()
		if !checkIfCreatedTimeIsValid(year, month, payload){
			continue
		}
		index := getIndex(year, month, payload)
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

func parsePathVariables(vars map[string]string) types.WasteTypeResponseRange{
	var payload types.WasteTypeResponseRange
	startYear, _ := strconv.Atoi(vars["startYear"])
	startMonth, _ := strconv.Atoi(vars["startMonth"])
	endYear, _ := strconv.Atoi(vars["endYear"])
	endMonth, _ := strconv.Atoi(vars["endMonth"])

	payload = types.WasteTypeResponseRange{
		StartYear:  startYear,
		StartMonth: startMonth,
		EndYear:    endYear,
		EndMonth:   endMonth,
	}

	return payload
}

func initializeMetrics(payload types.WasteTypeResponseRange, totalMonths int) []types.WasteTypeResponseMetric {
    metrics := make([]types.WasteTypeResponseMetric, totalMonths)
    currentYear := payload.StartYear
    currentMonth := payload.StartMonth

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

    return metrics
}

func getTotalMonths(payload types.WasteTypeResponseRange) int{
	startYear := payload.StartYear
	endYear := payload.EndYear
	startMonth := payload.StartMonth
	endMonth := payload.EndMonth
	totalMonths := (endYear - startYear) * 12
	totalMonths += endMonth - startMonth + 1
	return totalMonths
}

func checkIfCreatedTimeIsValid(year int, month time.Month, payload types.WasteTypeResponseRange) bool{
	startYear := payload.StartYear
	endYear := payload.EndYear
	startMonth := payload.StartMonth
	endMonth := payload.EndMonth
	if year < startYear || year > endYear{
		return false
	}
	if year == startYear && month < time.Month(startMonth) || year == endYear && month > time.Month(endMonth){
		return false
	}
	return true
}

func getIndex(year int, month time.Month, payload types.WasteTypeResponseRange) int{
	indices := map[string]int{
		"January":   1,
		"February":  2,
		"March":     3,
		"April":     4,
		"May":       5,
		"June":      6,
		"July":      7,
		"August":    8,
		"September": 9,
		"October":   10,
		"November":  11,
		"December":  12,
	}
	startYear := payload.StartYear
	startMonth := payload.StartMonth
	return (year - startYear) * 12 + indices[month.String()] - startMonth
}
