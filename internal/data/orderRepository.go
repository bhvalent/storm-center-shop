package data

import (
	"context"
	"encoding/json"
	"log"
	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/data/azcosmos"
	"github.com/google/uuid"
)

type OrderRepository struct {
	db Cosmos
}

type IOrderRepository interface {
	GetOrders(c context.Context, userId string) []OrderEntity
	CreateOrder(c context.Context, o OrderEntity) (OrderEntity, error)
}

func NewOrderRepository(db *Cosmos) *OrderRepository {
	return &OrderRepository{db: *db}
}

func (or *OrderRepository) GetOrders(c context.Context, userId string) []OrderEntity {
	client := or.db.GetClient()

	orderContainer, err := client.NewContainer("stormcenter", "orders")
	if err != nil {
		log.Fatal("Failed get orders container: ", err)
	}

	var orderEntities []OrderEntity

	pk := azcosmos.NewPartitionKeyString(userId)
	queryPager := orderContainer.NewQueryItemsPager("select * from orders", pk, nil)
	for queryPager.More() {
		queryResponse, err := queryPager.NextPage(c)
		if err != nil {
			log.Fatal("Error getting next page while getting orders: ", err)
		}

		for _, item := range queryResponse.Items {
			var order OrderEntity
			json.Unmarshal(item, &order)
			orderEntities = append(orderEntities, order)
		}
	}
	return orderEntities
}

func (or *OrderRepository) CreateOrder(c context.Context, o OrderEntity) (OrderEntity, error) {
	client := or.db.GetClient()

	orderContainer, err := client.NewContainer("stormcenter", "orders")
	if err != nil {
		log.Fatal("Failed get orders container: ", err)
	}

	pk := azcosmos.NewPartitionKeyString(o.UserId.String())
	o.CreatedDate = time.Now()
	o.Id = uuid.New()
	data, _ := json.Marshal(o)

	itemOptions := azcosmos.ItemOptions{
		ConsistencyLevel:             azcosmos.ConsistencyLevelSession.ToPtr(),
		EnableContentResponseOnWrite: true,
	}

	response, err := orderContainer.CreateItem(c, pk, data, &itemOptions)
	if err != nil {
		log.Print(err)
		return OrderEntity{}, err
	}

	var newOrder OrderEntity
	json.Unmarshal(response.Value, &newOrder)

	return newOrder, nil
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
