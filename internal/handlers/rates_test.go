package handlers

import (
    _"database/sql"
    "encoding/json"
    "net/http"
    "net/http/httptest"
    "testing"

    "github.com/DATA-DOG/go-sqlmock"
    "github.com/stretchr/testify/assert"
)

func TestGetRatesByDate_NoData(t *testing.T) {
    db, mock, err := sqlmock.New()
    if err != nil {
        t.Fatalf("Ошибка создания mock БД: %v", err)
    }
    defer db.Close()

 
    mockRows := sqlmock.NewRows([]string{
        "Cur_ID", "Date", "Cur_Abbreviation", "Cur_Scale", "Cur_Name", "Cur_OfficialRate",
    })

    testDate := "2099-12-31"
    mock.ExpectQuery(`SELECT (.+) FROM exchange_rates WHERE Date = STR_TO_DATE\(\?, '%Y-%m-%d'\)`).
        WithArgs(testDate).
        WillReturnRows(mockRows)

    
    req := httptest.NewRequest("GET", "/api/rates?date="+testDate, nil)
    w := httptest.NewRecorder()


    handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        mux := http.NewServeMux()
        RegisterRoutes(mux, db)
        mux.ServeHTTP(w, r)
    })

    handler.ServeHTTP(w, req)

    
    assert.Equal(t, http.StatusOK, w.Code)

    
    var response map[string]string
    err = json.NewDecoder(w.Body).Decode(&response)
    assert.NoError(t, err)
    assert.Equal(t, "нет информации за эту дату", response["message"])

    
    if err := mock.ExpectationsWereMet(); 
    err != nil {
        t.Errorf("Остались невыполненные ожидания: %v", err)
    }
}
