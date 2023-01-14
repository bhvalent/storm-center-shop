package domain

import (
	"context"
	"encoding/json"
	"log"
	"storm-center-backend/cmd/data"
	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/data/azcosmos"
	"github.com/google/uuid"
)

func GetOrders(c context.Context, userId string) []Order {
	var orders []Order

	var orderContainer = data.GetContainer("orders")

	pk := azcosmos.NewPartitionKeyString(userId)
	queryPager := orderContainer.NewQueryItemsPager("select * from orders", pk, nil)
	for queryPager.More() {
		queryResponse, err := queryPager.NextPage(c)
		if err != nil {
			log.Fatal("Error getting next page while getting orders: ", err)
		}

		for _, item := range queryResponse.Items {
			var order Order
			json.Unmarshal(item, &order)
			orders = append(orders, order)
		}
	}
	return orders
}

func SaveOrder(c context.Context, order Order) (Order, error) {
	var orderContainer = data.GetContainer("orders")

	pk := azcosmos.NewPartitionKeyString(order.UserId.String())
	order.CreatedDate = time.Now()
	order.Id = uuid.New()
	data, _ := json.Marshal(order)

	itemOptions := azcosmos.ItemOptions{
		ConsistencyLevel:             azcosmos.ConsistencyLevelSession.ToPtr(),
		EnableContentResponseOnWrite: true,
	}

	response, err := orderContainer.CreateItem(c, pk, data, &itemOptions)
	if err != nil {
		log.Print(err)
		return Order{}, err
	}

	var newOrder Order
	json.Unmarshal(response.Value, &newOrder)

	return newOrder, nil
}

type Order struct {
	Id          uuid.UUID `json:"id"`
	UserId      uuid.UUID `json:"userId"`
	Items       []Item    `json:"items"`
	CreatedDate time.Time `json:"createdDate"`
}

type Item struct {
	Id    uuid.UUID `json:"id"`
	Name  string    `json:"name"`
	Price float64   `json:"price"`
}
