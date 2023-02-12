package services

import (
	"context"
	"storm-center-shop/internal/data"
	"storm-center-shop/internal/domain/models"
	"storm-center-shop/internal/utils"
)

type OrderService struct {
	repo data.IOrderRepository
}

func NewOrderService(orderRepository data.IOrderRepository) *OrderService {
	return &OrderService{repo: orderRepository}
}

func (os *OrderService) GetOrders(c context.Context, userId string) ([]*models.Order, error) {
	orderEntities, err := os.repo.GetOrders(c, userId)
	if err != nil {
		return nil, err
	}
	return utils.Map(orderEntities, models.OrderEntityToOrder), nil
}
