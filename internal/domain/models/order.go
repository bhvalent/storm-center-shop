package models

import (
	"context"
	"log"
	"storm-center-backend/internal/data"
	"storm-center-backend/pkg/api"
	"time"

	"github.com/google/uuid"
)

type Order struct {
	Id          uuid.UUID `json:"id"`
	UserId      uuid.UUID `json:"userId"`
	Items       []Item    `json:"items"`
	CreatedDate time.Time `json:"createdDate"`
	repo        data.IOrderRepository
}

type Item struct {
	Id    uuid.UUID `json:"id"`
	Name  string    `json:"name"`
	Price float64   `json:"price"`
}

func NewOrderFromCreateOrderRequest(cor api.CreateOrderRequest, or data.IOrderRepository) *Order {
	order := CreateOrderRequestToOrder(cor)
	order.repo = or
	return &order
}

func (o *Order) Create(c context.Context) (Order, error) {
	oe := orderToEntity(*o)
	newOrderEntity, err := o.repo.CreateOrder(c, oe)
	if err != nil {
		log.Fatal(err)
	}
	return OrderEntityToOrder(newOrderEntity), nil
}

func CreateOrderRequestToOrder(cor api.CreateOrderRequest) Order {
	var items []Item
	for _, requestItem := range cor.Items {
		items = append(items, RequestItemToItem(requestItem))
	}

	return Order{
		Id:          cor.Id,
		UserId:      cor.UserId,
		CreatedDate: cor.CreatedDate,
		Items:       items,
	}
}

func RequestItemToItem(ri api.RequestItem) Item {
	return Item{
		Id:    ri.Id,
		Name:  ri.Name,
		Price: ri.Price,
	}
}

func OrderEntityToOrder(oe data.OrderEntity) Order {
	var items []Item
	for _, itemEntity := range oe.Items {
		items = append(items, ItemEntityToItem(itemEntity))
	}

	return Order{
		Id:          oe.Id,
		UserId:      oe.UserId,
		CreatedDate: oe.CreatedDate,
		Items:       items,
	}
}

func ItemEntityToItem(ie data.ItemEntity) Item {
	return Item{
		Id:    ie.Id,
		Name:  ie.Name,
		Price: ie.Price,
	}
}

func orderToEntity(o Order) data.OrderEntity {
	var items []data.ItemEntity
	for _, itemEntity := range o.Items {
		items = append(items, itemToEntity(itemEntity))
	}

	return data.OrderEntity{
		Id:          o.Id,
		UserId:      o.UserId,
		CreatedDate: o.CreatedDate,
		Items:       items,
	}
}

func itemToEntity(i Item) data.ItemEntity {
	return data.ItemEntity{
		Id:    i.Id,
		Name:  i.Name,
		Price: i.Price,
	}
}
