package repository

import (
	"database/sql"
	"testProject/internal/models"
)

type RateRepository struct {
	DB *sql.DB
}

func (r *RateRepository) GetAllRates() ([]models.Rate, error) {
	rows, err := r.DB.Query("SELECT currency, rate, date FROM exchange_rates")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var rates []models.Rate
	for rows.Next() {
		var rate models.Rate
		if err := rows.Scan(&rate.Currency, &rate.Rate, &rate.Date); err != nil {
			return nil, err
		}
		rates = append(rates, rate)
	}
	return rates, nil
}

func (r *RateRepository) GetRatesByDate(date string) ([]models.Rate, error) {
	rows, err := r.DB.Query("SELECT currency, rate, date FROM exchange_rates WHERE date = ?", date)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var rates []models.Rate
	for rows.Next() {
		var rate models.Rate
		if err := rows.Scan(&rate.Currency, &rate.Rate, &rate.Date); err != nil {
			return nil, err
		}
		rates = append(rates, rate)
	}
	return rates, nil
}
