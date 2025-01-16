package main

import (
    "fmt"
    "log"
    "net/http"
    "testProject/config"
    "testProject/internal/handlers"
    "testProject/internal/repository"
    "testProject/internal/service"
    "testProject/pkg"
)

func main() {
	
	cfg := config.LoadConfig()


	database, err := db.ConnectMySQL(cfg.DatabaseDSN)
	if err != nil {
		log.Fatalf("Ошибка подключения к базе данных: %v", err)
	}
	defer database.Close()

    rateRepo := &repository.RateRepository{DB: database}
    nbrbService := service.NewNBRBService()

   
    rates, err := nbrbService.GetCurrentRates()
    if (err != nil) {
        log.Printf("Ошибка получения начальных данных из API: %v", err)
    } else {
        for _, rate := range rates {
            if err := rateRepo.SaveRate(rate); err != nil {
                log.Printf("Ошибка сохранения курса %s: %v", rate.Cur_Abbreviation, err)
            }
        }
    }

    rateUpdater := service.NewRateUpdater(rateRepo)
    rateUpdater.StartDailyUpdate()


	mux := http.NewServeMux()
	handlers.RegisterRoutes(mux, database)

	
	log.Printf("Сервер запущен на :%d", cfg.Port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", cfg.Port), mux))
}
