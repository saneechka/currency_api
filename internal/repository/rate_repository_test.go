package repository

import (
    "testing"
    "database/sql"
    "github.com/DATA-DOG/go-sqlmock"
    "github.com/stretchr/testify/assert"
    _"testProject/internal/models"
)

func TestRateRepository_GetAllRates(t *testing.T) {
    // Создаем мок базы данных
    db, mock, err := sqlmock.New()
    if err != nil {
        t.Fatalf("Ошибка создания мока БД: %v", err)
    }
    defer db.Close()

    repo := &RateRepository{DB: db}

    // Подготавливаем тестовые данные
    mockRows := sqlmock.NewRows([]string{
        "Cur_ID", "Date", "Cur_Abbreviation", "Cur_Scale", "Cur_Name", "Cur_OfficialRate",
    }).AddRow(
        431, "2024-01-16", "USD", 1, "Доллар США", 3.2345,
    ).AddRow(
        451, "2024-01-16", "EUR", 1, "Евро", 3.5678,
    )

    // Ожидаем SQL запрос
    mock.ExpectQuery("^SELECT (.+) FROM exchange_rates$").
        WillReturnRows(mockRows)

    // Выполняем тест
    rates, err := repo.GetAllRates()

    // Проверяем результаты
    assert.NoError(t, err)
    assert.Len(t, rates, 2)
    assert.Equal(t, "USD", rates[0].Cur_Abbreviation)
    assert.Equal(t, "EUR", rates[1].Cur_Abbreviation)
    assert.Equal(t, 3.2345, rates[0].Cur_OfficialRate)
    assert.Equal(t, 3.5678, rates[1].Cur_OfficialRate)

    // Проверяем, что все ожидания были выполнены
    if err := mock.ExpectationsWereMet(); err != nil {
        t.Errorf("Остались невыполненные ожидания: %v", err)
    }
}

func TestRateRepository_GetRatesByDate(t *testing.T) {
    // Создаем мок базы данных
    db, mock, err := sqlmock.New()
    if err != nil {
        t.Fatalf("Ошибка создания мока БД: %v", err)
    }
    defer db.Close()

    repo := &RateRepository{DB: db}
    testDate := "2024-01-16"

    // Подготавливаем тестовые данные
    mockRows := sqlmock.NewRows([]string{
        "Cur_ID", "Date", "Cur_Abbreviation", "Cur_Scale", "Cur_Name", "Cur_OfficialRate",
    }).AddRow(
        431, testDate, "USD", 1, "Доллар США", 3.2345,
    ).AddRow(
        451, testDate, "EUR", 1, "Евро", 3.5678,
    )

    // Ожидаем SQL запрос с параметром даты
    mock.ExpectQuery(`SELECT (.+) FROM exchange_rates WHERE Date = STR_TO_DATE\(\?, '%Y-%m-%d'\)`).
        WithArgs(testDate).
        WillReturnRows(mockRows)

    // Выполняем тест
    rates, err := repo.GetRatesByDate(testDate)

    // Проверяем результаты
    assert.NoError(t, err)
    assert.Len(t, rates, 2)
    assert.Equal(t, testDate, rates[0].Date)
    assert.Equal(t, "USD", rates[0].Cur_Abbreviation)
    assert.Equal(t, testDate, rates[1].Date)
    assert.Equal(t, "EUR", rates[1].Cur_Abbreviation)

    // Проверяем ошибки SQL
    if err := mock.ExpectationsWereMet(); err != nil {
        t.Errorf("Остались невыполненные ожидания: %v", err)
    }
}

// Тест на ошибку БД
func TestRateRepository_GetAllRates_DBError(t *testing.T) {
    db, mock, err := sqlmock.New()
    if err != nil {
        t.Fatalf("Ошибка создания мока БД: %v", err)
    }
    defer db.Close()

    repo := &RateRepository{DB: db}

    // Симулируем ошибку БД
    mock.ExpectQuery("^SELECT (.+) FROM exchange_rates$").
        WillReturnError(sql.ErrConnDone)

    // Выполняем тест
    rates, err := repo.GetAllRates()

    // Проверяем результаты
    assert.Error(t, err)
    assert.Nil(t, rates)
    assert.Equal(t, sql.ErrConnDone, err)
}

// Тест на пустой результат
func TestRateRepository_GetRatesByDate_EmptyResult(t *testing.T) {
    db, mock, err := sqlmock.New()
    if err != nil {
        t.Fatalf("Ошибка создания мока БД: %v", err)
    }
    defer db.Close()

    repo := &RateRepository{DB: db}
    testDate := "2024-01-16"

    // Возвращаем пустой результат
    mockRows := sqlmock.NewRows([]string{
        "Cur_ID", "Date", "Cur_Abbreviation", "Cur_Scale", "Cur_Name", "Cur_OfficialRate",
    })

    mock.ExpectQuery(`SELECT (.+) FROM exchange_rates WHERE Date = STR_TO_DATE\(\?, '%Y-%m-%d'\)`).
        WithArgs(testDate).
        WillReturnRows(mockRows)

    // Выполняем тест
    rates, err := repo.GetRatesByDate(testDate)

    // Проверяем результаты
    assert.NoError(t, err)
    assert.Empty(t, rates)
}

// Тест на получение курсов за несуществующую дату
func TestRateRepository_GetRatesByDate_NonExistentDate(t *testing.T) {
    db, mock, err := sqlmock.New()
    if err != nil {
        t.Fatalf("Ошибка создания мока БД: %v", err)
    }
    defer db.Close()

    repo := &RateRepository{DB: db}
    nonExistentDate := "2099-12-31" // Будущая дата, которой точно нет в БД

    // Ожидаем пустой результат для несуществующей даты
    mock.ExpectQuery(`SELECT (.+) FROM exchange_rates WHERE Date = STR_TO_DATE\(\?, '%Y-%m-%d'\)`).
        WithArgs(nonExistentDate).
        WillReturnRows(sqlmock.NewRows([]string{
            "Cur_ID", "Date", "Cur_Abbreviation", "Cur_Scale", "Cur_Name", "Cur_OfficialRate",
        }))

    // Выполняем тест
    rates, err := repo.GetRatesByDate(nonExistentDate)

    // Проверяем результаты
    assert.NoError(t, err, "Должны получить nil ошибку для несуществующей даты")
    assert.Empty(t, rates, "Список курсов должен быть пустым")
    assert.Len(t, rates, 0, "Длина списка курсов должна быть 0")

    // Проверяем, что все ожидания были выполнены
    if err := mock.ExpectationsWereMet(); err != nil {
        t.Errorf("Остались невыполненные ожидания: %v", err)
    }
}

// Тест на получение курсов с некорректным форматом даты
func TestRateRepository_GetRatesByDate_InvalidDateFormat(t *testing.T) {
    db, mock, err := sqlmock.New()
    if err != nil {
        t.Fatalf("Ошибка создания мока БД: %v", err)
    }
    defer db.Close()

    repo := &RateRepository{DB: db}
    invalidDate := "2024-13-45" // Неправильный формат даты

    // Ожидаем ошибку SQL из-за неправильного формата даты
    mock.ExpectQuery(`SELECT (.+) FROM exchange_rates WHERE Date = STR_TO_DATE\(\?, '%Y-%m-%d'\)`).
        WithArgs(invalidDate).
        WillReturnError(sql.ErrNoRows)

    // Выполняем тест
    rates, err := repo.GetRatesByDate(invalidDate)

    // Проверяем результаты
    assert.Error(t, err, "Должны получить ошибку для некорректной даты")
    assert.Nil(t, rates, "Список курсов должен быть nil")

    // Проверяем, что все ожидания были выполнены
    if err := mock.ExpectationsWereMet(); err != nil {
        t.Errorf("Остались невыполненные ожидания: %v", err)
    }
}
