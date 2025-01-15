package config

import (
	"log"
	"os"
	"strconv"
)

type Config struct {
	DatabaseDSN string
	Port        int
}

func LoadConfig() Config {
	port, err := strconv.Atoi(getEnv("PORT", "8080"))
	if err != nil {
		log.Fatalf("Некорректный порт: %v", err)
	}

	return Config{
		DatabaseDSN: "root:admin@tcp(localhost:3306)/currency_db",
		Port:        port,
	}
}

func getEnv(key, defaultValue string) string {
	value := os.Getenv(key)
	if value == "" {
		return defaultValue
	}
	return value
}
