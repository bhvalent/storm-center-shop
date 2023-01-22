package services

import (
	"context"
	"storm-center-backend/internal/data"
	"storm-center-backend/internal/domain/models"
	"storm-center-backend/internal/utils"
)

type OrderService struct {
	repo data.IOrderRepository
}

func NewOrderService(orderRepository data.IOrderRepository) *OrderService {
	return &OrderService{repo: orderRepository}
}

func (os *OrderService) GetOrders(c context.Context, userId string) []models.Order {
	orderEntities := os.repo.GetOrders(c, userId)
	return utils.Map(orderEntities, models.OrderEntityToOrder)
}
