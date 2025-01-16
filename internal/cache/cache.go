package cache

import (
	"sync"
	"testProject/internal/models"
	"time"
)

type RatesCache struct {
	mutex sync.RWMutex
	data  map[string][]models.Rate
	ttl   time.Duration
}

func NewRatesCache(ttl time.Duration) *RatesCache {
	return &RatesCache{
		data: make(map[string][]models.Rate),
		ttl:  ttl,
	}
}

func (c *RatesCache) Set(date string, rates []models.Rate) {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	c.data[date] = rates

	go func() {
		time.Sleep(c.ttl)
		c.Delete(date)
	}()
}

func (c *RatesCache) Get(date string) ([]models.Rate, bool) {
	c.mutex.RLock()
	defer c.mutex.RUnlock()
	rates, exists := c.data[date]
	return rates, exists
}

func (c *RatesCache) Delete(date string) {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	delete(c.data, date)
}
