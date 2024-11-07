package inmemory

import (
	"context"
	"errors"
	"gophermart/internal/accrual/entity"
	"sync"
)

type memoryStorage struct {
	orderData map[entity.ID]entity.Order
	mutex     sync.RWMutex
}

var (
	ErrOrderNotFound = errors.New("order not found")
)

// NewMemoryStorage create new data storage in memory
func NewMemoryStorage() (*memoryStorage, error) {
	return &memoryStorage{
		orderData: make(map[entity.ID]entity.Order),
	}, nil
}

func (s *memoryStorage) CreateOrder(_ context.Context, o entity.Order) error {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	s.orderData[o.Number] = o
	return nil
}

func (s *memoryStorage) GetOrderByID(_ context.Context, id entity.ID) (entity.Order, error) {
	s.mutex.RLock()
	defer s.mutex.RUnlock()

	u, ok := s.orderData[id]
	if !ok {
		return entity.Order{}, ErrOrderNotFound
	}
	return u, nil
}

func (s *memoryStorage) Close() error {
	return nil
}
