package handlers

import (
	"encoding/json"
	"net/http"
	"testProject/internal/repository"
	"database/sql"
)

func GetAllRates(repo *repository.RateRepository) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		rates, err := repo.GetAllRates()
		if err != nil {
			http.Error(w, "Ошибка получения данных", http.StatusInternalServerError)
			return
		}
		json.NewEncoder(w).Encode(rates)
	}
}
func RegisterRoutes(mux *http.ServeMux, db *sql.DB) {
	mux.HandleFunc("/rates", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Получение всех курсов валют"))
	})
	mux.HandleFunc("/rates/by-date", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Получение курсов валют за определённую дату"))
	})
}
func GetRatesByDate(repo *repository.RateRepository) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		date := r.URL.Query().Get("date")
		if date == "" {
			http.Error(w, "Не указана дата", http.StatusBadRequest)
			return
		}
		rates, err := repo.GetRatesByDate(date)
		if err != nil {
			http.Error(w, "Ошибка получения данных", http.StatusInternalServerError)
			return
		}
		json.NewEncoder(w).Encode(rates)
	}
}

