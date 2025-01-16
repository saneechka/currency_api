package service

import (
    "encoding/json"
    "fmt"
    "net/http"
    "time"
    "testProject/internal/models"
)

type NBRBService struct {
    apiURL string
}

func NewNBRBService() *NBRBService {
    return &NBRBService{
        apiURL: "https://api.nbrb.by/exrates/rates?periodicity=0",
    }
}

func (s *NBRBService) GetCurrentRates() ([]models.Rate, error) {
    resp, err := http.Get(s.apiURL)
    if err != nil {
        return nil, fmt.Errorf("ошибка получения данных из API: %v", err)
    }
    defer resp.Body.Close()

    var rates []models.Rate
    if err := json.NewDecoder(resp.Body).Decode(&rates); err != nil {
        return nil, fmt.Errorf("ошибка декодирования ответа: %v", err)
    }

    
    for i := range rates {
       
        t, err := time.Parse("2006-01-02T15:04:05", rates[i].Date)
        if err != nil {
            continue
        }
      
        rates[i].Date = t.Format("2006-01-02")
    }

    return rates, nil
}
