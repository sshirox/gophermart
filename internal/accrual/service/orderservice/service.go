package orderservice

import (
	"context"
	"gophermart/internal/accrual/entity"
)

type storager interface {
	CreateOrder(context.Context, entity.Order) error
	GetOrderByID(context.Context, entity.ID) (entity.Order, error)
}

type Service struct {
	storage storager
}

func NewOrderService(storage storager) *Service {
	return &Service{storage: storage}
}

func (u *Service) CreateOrder(ctx context.Context, order entity.Order) error {
	return u.storage.CreateOrder(ctx, order)
}

func (u *Service) GetOrderByID(ctx context.Context, id entity.ID) (entity.Order, error) {
	return u.storage.GetOrderByID(ctx, id)
}
