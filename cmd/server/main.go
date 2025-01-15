package main

import (
    "fmt"
    "log"
    "net/http"
    "testProject/config"
    "testProject/internal/handlers"
    "testProject/pkg"
    "testProject/internal/service"
)

func main() {
	// Загрузка конфигурации
	cfg := config.LoadConfig()

	// Подключение к базе данных
	database, err := db.ConnectMySQL(cfg.DatabaseDSN)
	if err != nil {
		log.Fatalf("Ошибка подключения к базе данных: %v", err)
	}
	defer database.Close()

    // Создаем и запускаем сервис обновления курсов
    rateRepo := &repository.RateRepository{DB: database}
    rateUpdater := service.NewRateUpdater(rateRepo)
    rateUpdater.StartDailyUpdate()

	// Регистрация маршрутов
	mux := http.NewServeMux()
	handlers.RegisterRoutes(mux, database)

	// Запуск сервера
	log.Printf("Сервер запущен на :%d", cfg.Port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", cfg.Port), mux))
}
