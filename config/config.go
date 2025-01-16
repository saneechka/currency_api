package config

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strconv"

	"github.com/joho/godotenv"
)

type Config struct {
	DatabaseDSN string
	Port        int
}

func findRootDir() (string, error) {
	currentDir, err := os.Getwd()
	if err != nil {
		return "", fmt.Errorf("не удалось получить рабочую директорию: %v", err)
	}

	for {
		if _, err := os.Stat(filepath.Join(currentDir, ".env")); err == nil {
			return currentDir, nil
		}

		parentDir := filepath.Dir(currentDir)
		if parentDir == currentDir {
			return "", fmt.Errorf("файл .env не найден в родительских директориях")
		}
		currentDir = parentDir
	}
}

func LoadConfig() (Config, error) {
	rootDir, err := findRootDir()
	if err != nil {
		log.Printf("Предупреждение: %v", err)
	} else {
		envPath := filepath.Join(rootDir, ".env")
		if err := godotenv.Load(envPath); err != nil {
			log.Printf("Предупреждение: Не удалось загрузить файл .env из %s: %v", envPath, err)
		} else {
			log.Printf("Загружен файл .env из: %s", envPath)
		}
	}

	port, err := strconv.Atoi(getEnv("PORT", "8080"))
	if err != nil {
		return Config{}, fmt.Errorf("некорректное значение порта: %v", err)
	}

	databaseDSN := getEnv("DATABASE_DSN", "")
	if databaseDSN == "" {
		return Config{}, fmt.Errorf("DATABASE_DSN не задан в переменных окружения или файле .env")
	}

	return Config{
		DatabaseDSN: databaseDSN,
		Port:        port,
	}, nil
}

func getEnv(key, defaultValue string) string {
	value := os.Getenv(key)
	if value == "" {
		return defaultValue
	}
	return value
}
