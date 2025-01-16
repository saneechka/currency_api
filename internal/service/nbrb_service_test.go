package service

import (
    "net/http"
    "net/http/httptest"
    "testing"
    _"time"
)

func TestGetCurrentRates(t *testing.T) {
    tests := []struct {
        name           string
        responseBody   string
        expectedError  bool
        expectedCount  int
        expectedCurr   string
        responseStatus int
    }{
        {
            name: "успешное получение курсов",
            responseBody: `[
                {"Cur_ID":431,"Date":"2024-01-16T00:00:00","Cur_Abbreviation":"USD","Cur_Scale":1,"Cur_Name":"Доллар США","Cur_OfficialRate":3.2345}
            ]`,
            expectedError: false,
            expectedCount: 1,
            expectedCurr:  "USD",
            responseStatus: http.StatusOK,
        },
        {
            name:           "ошибка сервера",
            responseBody:   `{"error": "server error"}`,
            expectedError:  true,
            responseStatus: http.StatusInternalServerError,
        },
        {
            name:           "неверный формат JSON",
            responseBody:   `invalid json`,
            expectedError:  true,
            responseStatus: http.StatusOK,
        },
    }

    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
                w.WriteHeader(tt.responseStatus)
                w.Write([]byte(tt.responseBody))
            }))
            defer server.Close()

            service := NBRBService{apiURL: server.URL}
            rates, err := service.GetCurrentRates()

            if tt.expectedError && err == nil {
                t.Error("ожидалась ошибка, но её не получили")
            }
            if !tt.expectedError && err != nil {
                t.Errorf("неожиданная ошибка: %v", err)
            }
            if !tt.expectedError && len(rates) != tt.expectedCount {
                t.Errorf("ожидалось %d курсов, получено %d", tt.expectedCount, len(rates))
            }
            if !tt.expectedError && len(rates) > 0 && rates[0].Cur_Abbreviation != tt.expectedCurr {
                t.Errorf("ожидалась валюта %s, получена %s", tt.expectedCurr, rates[0].Cur_Abbreviation)
            }
        })
    }
}
