package handlers

import (
	"dev11/dev11/internal/models"
	"encoding/json"
	"net/http"
)

// UpdateEventHandler обрабатывает запросы на обновление события.
func UpdateEventHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid method", http.StatusMethodNotAllowed)
		return
	}

	updatedEvent, err := parseAndUpdateEvent(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = models.Cache.UpdateEvent(updatedEvent.ID, updatedEvent)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"result": "Event updated"})
}

// parseAndUpdateEvent парсит обновленное событие из тела запроса.
func parseAndUpdateEvent(r *http.Request) (models.Event, error) {
	var updatedEvent models.Event
	err := json.NewDecoder(r.Body).Decode(&updatedEvent)
	if err != nil {
		return models.Event{}, err
	}
	return updatedEvent, nil
}
