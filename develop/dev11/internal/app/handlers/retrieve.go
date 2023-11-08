package handlers

import (
	"dev11/dev11/internal/models"
	"encoding/json"
	"errors"
	"net/http"
	"time"
)

// GetEventsForDayHandler обрабатывает запросы на получение событий за день.
func GetEventsForDayHandler(w http.ResponseWriter, r *http.Request) {
	day, err := parseDateQuery(r, "day")
	if err != nil {
		http.Error(w, "Invalid day query parameter", http.StatusBadRequest)
		return
	}

	events := models.Cache.GetEventsForDay(day)
	json.NewEncoder(w).Encode(events)
}

// GetEventsForWeekHandler обрабатывает запросы на получение событий за неделю.
func GetEventsForWeekHandler(w http.ResponseWriter, r *http.Request) {
	// Предполагается, что параметр week - это начало недели в формате YYYY-MM-DD
	startOfWeek, err := parseDateQuery(r, "week")
	if err != nil {
		http.Error(w, "Invalid week query parameter", http.StatusBadRequest)
		return
	}

	events := models.Cache.GetEventsForWeek(startOfWeek)
	json.NewEncoder(w).Encode(events)
}

// GetEventsForMonthHandler обрабатывает запросы на получение событий за месяц.
func GetEventsForMonthHandler(w http.ResponseWriter, r *http.Request) {
	// Предполагается, что параметр month - это месяц в формате YYYY-MM
	year, month, err := parseMonthQuery(r, "month")
	if err != nil {
		http.Error(w, "Invalid month query parameter", http.StatusBadRequest)
		return
	}

	events := models.Cache.GetEventsForMonth(year, month)
	json.NewEncoder(w).Encode(events)
}

// Вспомогательная функция для парсинга даты из запроса.
func parseDateQuery(r *http.Request, queryParam string) (time.Time, error) {
	dateStr := r.URL.Query().Get(queryParam)
	if dateStr == "" {
		return time.Time{}, errors.New("query parameter missing")
	}
	return time.Parse("2006-01-02", dateStr)
}

// Вспомогательная функция для парсинга месяца из запроса.
func parseMonthQuery(r *http.Request, queryParam string) (int, time.Month, error) {
	monthStr := r.URL.Query().Get(queryParam)
	if monthStr == "" {
		return 0, 0, errors.New("query parameter missing")
	}
	date, err := time.Parse("2006-01", monthStr)
	if err != nil {
		return 0, 0, err
	}
	return date.Year(), date.Month(), nil
}
