package handlers

import (
	"dev11/dev11/internal/models"
	"encoding/json"
	"net/http"
	"strconv"
)

// DeleteEventHandler обрабатывает запросы на удаление события.
func DeleteEventHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid method", http.StatusMethodNotAllowed)
		return
	}

	eventIDStr := r.URL.Query().Get("id")
	if eventIDStr == "" {
		http.Error(w, "Event ID is required", http.StatusBadRequest)
		return
	}

	eventID, err := strconv.Atoi(eventIDStr)
	if err != nil {
		http.Error(w, "Invalid event ID", http.StatusBadRequest)
		return
	}

	err = models.Cache.DeleteEvent(eventID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"result": "Event deleted"})
}
