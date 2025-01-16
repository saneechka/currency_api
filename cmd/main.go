package main

import (
    "fmt"
    "log"
    "net/http"
    "os"
    "path/filepath"
    "strings"
    "testProject/config"
    "testProject/internal/handlers"
    "testProject/internal/middleware"
    "testProject/internal/repository"
    "testProject/internal/service"
    "testProject/pkg/logger"
    db "testProject/pkg"
)

// Функция для получения корневой директории проекта
func getRootDir() (string, error) {
    currentDir, err := os.Getwd()
    if (err != nil) {
        return "", fmt.Errorf("ошибка получения рабочей директории: %v", err)
    }

    // Проверяем, находимся ли мы в директории cmd
    if filepath.Base(currentDir) == "cmd" {
        return filepath.Dir(currentDir), nil
    }

    // Если мы не в cmd, ищем корень проекта по наличию go.mod
    dir := currentDir
    for {
        if _, err := os.Stat(filepath.Join(dir, "go.mod")); err == nil {
            return dir, nil
        }
        parent := filepath.Dir(dir)
        if parent == dir {
            return "", fmt.Errorf("не удалось найти корень проекта (go.mod не найден)")
        }
        dir = parent
    }
}

func main() {
    
    rootDir, err := getRootDir()
    if err != nil {
        log.Fatalf("Ошибка определения корневой директории проекта: %v", err)
    }

    
    logsDir := filepath.Join(rootDir, "logs")
    if err := os.MkdirAll(logsDir, 0755); err != nil {
        log.Fatalf("Ошибка создания директории логов: %v", err)
    }

    
    if !strings.HasPrefix(logsDir, rootDir) {
        log.Fatalf("Директория логов должна находиться в корне проекта")
    }


    err = logger.Init(logger.Config{
        LogLevel:   "debug",
        LogFile:    filepath.Join(logsDir, "app.log"),
        MaxSize:    10,
        MaxBackups: 5,
        MaxAge:     30,
        Console:    true,
        JSONFormat: true,
    })
    if err != nil {
        log.Fatalf("Ошибка инициализации логгера: %v", err)
    }

    cfg, err := config.LoadConfig()
    if err != nil {
        logger.Error(err, "Ошибка загрузки конфигурации", nil)
        os.Exit(1)
    }

    database, err := db.ConnectMySQL(cfg.DatabaseDSN)
    if err != nil {
        logger.Error(err, "Ошибка подключения к базе данных", nil)
        os.Exit(1)
    }
    defer database.Close()

    rateRepo := &repository.RateRepository{DB: database}
    nbrbService := service.NewNBRBService()


    rates, err := nbrbService.GetCurrentRates()
    if err != nil {
        logger.Error(err, "Ошибка получения начальных данных из API", nil)
    } else {
        for _, rate := range rates {
            if err := rateRepo.SaveRate(rate); err != nil {
                logger.Error(err, "Ошибка сохранения курса", map[string]interface{}{
                    "currency": rate.Cur_Abbreviation,
                })
            }
        }
    }

    rateUpdater := service.NewRateUpdater(rateRepo)
    rateUpdater.StartDailyUpdate()

    
    mux := http.NewServeMux()
    
    
    handler := middleware.SecurityMiddleware(
        middleware.LoggingMiddleware(
            middleware.RecoveryMiddleware(
                mux,
            ),
        ),
    )
    
    handlers.RegisterRoutes(mux, database)

    
    logger.Info("Сервер запущен", map[string]interface{}{
        "port": cfg.Port,
        "pid":  os.Getpid(),
    })

    if err := http.ListenAndServe(fmt.Sprintf(":%d", cfg.Port), handler); err != nil {
        logger.Error(err, "Ошибка запуска сервера", nil)
        os.Exit(1)
    }
}
