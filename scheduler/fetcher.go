package scheduler

import (
	"log"
	"time"
)

func StartFetchJob(fetchFunc func() error) {
	go func() {
		for {
			if err := fetchFunc(); err != nil {
				log.Printf("Ошибка при сборе данных: %v", err)
			}
			time.Sleep(24 * time.Hour)
		}
	}()
}
