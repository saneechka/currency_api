package service

import (
    "encoding/json"
    "fmt"
    "net/http"
    "time"
    "testProject/internal/models"
    "testProject/internal/repository"
)

type RateUpdater struct {
    repo *repository.RateRepository
}

func NewRateUpdater(repo *repository.RateRepository) *RateUpdater {
    return &RateUpdater{repo: repo}
}

func (ru *RateUpdater) StartDailyUpdate() {
   
    ru.updateRates()

    ticker := time.NewTicker(24 * time.Hour)
    go func() {
        for range ticker.C {
            ru.updateRates()
        }
    }()
}

func (ru *RateUpdater) updateRates() error {
    resp, err := http.Get("https://api.nbrb.by/exrates/rates?periodicity=0")
    if (err != nil) {
        return fmt.Errorf("ошибка получения данных из API: %v", err)
    }
    defer resp.Body.Close()

    var rates []models.Rate
    if err := json.NewDecoder(resp.Body).Decode(&rates); err != nil {
        return fmt.Errorf("ошибка декодирования ответа: %v", err)
    }

    for _, rate := range rates {
        if err := ru.repo.SaveRate(rate); err != nil {
            return fmt.Errorf("ошибка сохранения курса %s: %v", rate.Cur_Abbreviation, err)
        }
    }
    return nil
}
