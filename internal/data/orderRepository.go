package data

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/data/azcosmos"
	"github.com/google/uuid"
)

type OrderRepository struct {
	db Cosmos
}

type IOrderRepository interface {
	GetOrders(c context.Context, userId string) ([]*OrderEntity, error)
	CreateOrder(c context.Context, o OrderEntity) (*OrderEntity, error)
}

func NewOrderRepository(db *Cosmos) *OrderRepository {
	return &OrderRepository{db: *db}
}

func (or *OrderRepository) GetOrders(c context.Context, userId string) ([]*OrderEntity, error) {
	client, err := or.db.GetClient()
	if err != nil {
		return nil, fmt.Errorf("error in OrderRepository.GetOrders: %w", err)
	}

	orderContainer, err := client.NewContainer("stormcenter", "orders")
	if err != nil {
		return nil, fmt.Errorf("error while getting orders container in OrderRepository.GetOrders: %w", err)
	}

	var orderEntities []*OrderEntity

	pk := azcosmos.NewPartitionKeyString(userId)
	queryPager := orderContainer.NewQueryItemsPager("select * from orders", pk, nil)
	for queryPager.More() {
		queryResponse, err := queryPager.NextPage(c)
		if err != nil {
			return nil, fmt.Errorf("error while querying orders in OrderRepository.GetOrders: %w", err)
		}

		for _, item := range queryResponse.Items {
			var order OrderEntity
			json.Unmarshal(item, &order)
			orderEntities = append(orderEntities, &order)
		}
	}
	return orderEntities, nil
}

func (or *OrderRepository) CreateOrder(c context.Context, o OrderEntity) (*OrderEntity, error) {
	client, err := or.db.GetClient()
	if err != nil {
		return nil, fmt.Errorf("error while creating order in OrderRepository.CreateOrder: %w", err)
	}

	orderContainer, err := client.NewContainer("stormcenter", "orders")
	if err != nil {
		return nil, fmt.Errorf("error while creating order in OrderRepository.CreateOrder: %w", err)
	}

	pk := azcosmos.NewPartitionKeyString(o.UserId.String())

	o.CreatedDate = time.Now()
	o.Id = uuid.New()

	data, err := json.Marshal(o)
	if err != nil {
		return nil, fmt.Errorf("error while creating order in OrderRepository.CreateOrder: %w", err)
	}

	itemOptions := azcosmos.ItemOptions{
		ConsistencyLevel:             azcosmos.ConsistencyLevelSession.ToPtr(),
		EnableContentResponseOnWrite: true,
	}

	response, err := orderContainer.CreateItem(c, pk, data, &itemOptions)
	if err != nil {
		return nil, fmt.Errorf("error while creating order in OrderRepository.CreateOrder: %w", err)
	}

	var newOrder OrderEntity
	err = json.Unmarshal(response.Value, &newOrder)
	if err != nil {
		return nil, fmt.Errorf("error while creating order in OrderRepository.CreateOrder: %w", err)
	}

	return &newOrder, nil
}

type OrderEntity struct {
	Id          uuid.UUID    `json:"id"`
	UserId      uuid.UUID    `json:"userId"`
	Items       []ItemEntity `json:"items"`
	CreatedDate time.Time    `json:"createdDate"`
}

type ItemEntity struct {
	Id    uuid.UUID `json:"id"`
	Name  string    `json:"name"`
	Price float64   `json:"price"`
}
