package storage

import (
	"context"
	"errors"
	"gophermart/internal/accrual/config"
	"gophermart/internal/accrual/entity"
	"gophermart/internal/accrual/storage/inmemory"
	"gophermart/internal/accrual/storage/postgres"
	"io"
)

type StorageType string

const (
	StorageTypePostgres StorageType = "postgres"
	StorageTypeInmemory StorageType = "memory"
)

type Config struct {
	StorageType StorageType
	Postgres    *config.PostgresConfig
}

type Storager interface {
	CreateOrder(ctx context.Context, o entity.Order) error
	GetOrderByID(ctx context.Context, id entity.ID) (entity.Order, error)
	io.Closer
}

func NewStorage(cfg *Config) (Storager, error) {
	switch cfg.StorageType {
	case StorageTypePostgres:
		return postgres.NewPostgresStorage(cfg.Postgres)
	case StorageTypeInmemory:
		return inmemory.NewMemoryStorage()
	default:
		return nil, errors.New("unknown storage type")
	}
}
