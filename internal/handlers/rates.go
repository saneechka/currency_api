package handlers

import (
	"database/sql"
	"encoding/json"
	_ "log"
	"net/http"
	"testProject/internal/repository"
	"testProject/internal/service"
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
	repo := &repository.RateRepository{DB: db}
	nbrbService := service.NewNBRBService()

	mux.HandleFunc("/api/rates", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		switch r.Method {
		case http.MethodGet:
			date := r.URL.Query().Get("date")
			if date != "" {
				// Если указана дата, получаем курсы за эту дату
				rates, err := repo.GetRatesByDate(date)
				if err != nil {
					http.Error(w, "Ошибка получения данных", http.StatusInternalServerError)
					return
				}
				if len(rates) == 0 {
					w.WriteHeader(http.StatusOK)
					json.NewEncoder(w).Encode(map[string]string{
						"message": "нет информации за эту дату",
					})
					return
				}
				json.NewEncoder(w).Encode(rates)
				return
			}

			// Если дата не указана, получаем все курсы из базы
			rates, err := repo.GetAllRates()
			if err != nil {
				// Если не удалось получить из базы, пробуем получить текущие из API
				currentRates, apiErr := nbrbService.GetCurrentRates()
				if apiErr != nil {
					http.Error(w, "Ошибка получения данных", http.StatusInternalServerError)
					return
				}
				json.NewEncoder(w).Encode(currentRates)
				return
			}
			json.NewEncoder(w).Encode(rates)

		default:
			http.Error(w, "Метод не поддерживается", http.StatusMethodNotAllowed)
		}
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
