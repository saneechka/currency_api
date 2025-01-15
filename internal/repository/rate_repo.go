package repository

import (
	"database/sql"
	"testProject/internal/models"
)

type RateRepository struct {
	DB *sql.DB
}

func (r *RateRepository) GetAllRates() ([]models.Rate, error) {
	rows, err := r.DB.Query(`SELECT 
		Cur_ID, 
		Date, 
		Cur_Abbreviation, 
		Cur_Scale, 
		Cur_Name, 
		Cur_OfficialRate 
		FROM exchange_rates`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var rates []models.Rate
	for rows.Next() {
		var rate models.Rate
		if err := rows.Scan(
			&rate.Cur_ID,
			&rate.Date,
			&rate.Cur_Abbreviation,
			&rate.Cur_Scale,
			&rate.Cur_Name,
			&rate.Cur_OfficialRate,
		); err != nil {
			return nil, err
		}
		rates = append(rates, rate)
	}
	return rates, nil
}

func (r *RateRepository) GetRatesByDate(date string) ([]models.Rate, error) {
	rows, err := r.DB.Query(`SELECT 
		Cur_ID, 
		Date, 
		Cur_Abbreviation, 
		Cur_Scale, 
		Cur_Name, 
		Cur_OfficialRate 
		FROM exchange_rates WHERE date = ?`, date)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var rates []models.Rate
	for rows.Next() {
		var rate models.Rate
		if err := rows.Scan(
			&rate.Cur_ID,
			&rate.Date,
			&rate.Cur_Abbreviation,
			&rate.Cur_Scale,
			&rate.Cur_Name,
			&rate.Cur_OfficialRate,
		); err != nil {
			return nil, err
		}
		rates = append(rates, rate)
	}
	return rates, nil
}

func (r *RateRepository) SaveRate(rate models.Rate) error {
    query := `
        INSERT INTO exchange_rates 
        (Cur_ID, Date, Cur_Abbreviation, Cur_Scale, Cur_Name, Cur_OfficialRate)
        VALUES (?, ?, ?, ?, ?, ?)
        ON DUPLICATE KEY UPDATE
        Cur_Scale = VALUES(Cur_Scale),
        Cur_Name = VALUES(Cur_Name),
        Cur_OfficialRate = VALUES(Cur_OfficialRate)
    `
    
    _, err := r.DB.Exec(query,
        rate.Cur_ID,
        rate.Date,
        rate.Cur_Abbreviation,
        rate.Cur_Scale,
        rate.Cur_Name,
        rate.Cur_OfficialRate,
    )
    return err
}
