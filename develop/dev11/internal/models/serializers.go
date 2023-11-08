package models

import (
	"encoding/json"
	"net/http"
)

// Вспомогательные функции
func SerializeToJSON(w http.ResponseWriter, data interface{}, statusCode int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(data)
}

func ParseAndValidateEvent(r *http.Request) (Event, error) {
	var e Event
	// Парсинг и валидация события из запроса
	// Возвращать ошибку, если валидация не пройдена
	return e, nil
}
