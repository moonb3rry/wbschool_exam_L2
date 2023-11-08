package handlers

import (
	"dev11/dev11/internal/models"
	"encoding/json"
	"net/http"
	"time"
)

func CreateEventHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Only POST method is allowed", http.StatusMethodNotAllowed)
		return
	}

	// Декодируем JSON тело запроса в структуру Event
	var input struct {
		Title       string `json:"title"`
		Description string `json:"description"`
		Start       string `json:"start"`
		End         string `json:"end"`
	}

	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	// Парсим время начала и окончания события
	start, err := time.Parse(time.RFC3339, input.Start)
	if err != nil {
		http.Error(w, "Invalid start time format", http.StatusBadRequest)
		return
	}
	end, err := time.Parse(time.RFC3339, input.End)
	if err != nil {
		http.Error(w, "Invalid end time format", http.StatusBadRequest)
		return
	}

	// Создаем событие
	event, err := models.Cache.CreateEvent(input.Title, input.Description, start, end)
	if err != nil {
		http.Error(w, err.Error(), http.StatusServiceUnavailable)
		return
	}

	// Отправляем ответ с созданным событием
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(event)
}
