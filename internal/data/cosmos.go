package data

import (
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/data/azcosmos"
)

type Cosmos struct {
	key string
	url string
}

func NewCosmos(key string, url string) *Cosmos {
	return &Cosmos{key: key, url: url}
}

func (c *Cosmos) GetClient() *azcosmos.Client {
	cred, err := azcosmos.NewKeyCredential(c.key)
	if err != nil {
		log.Fatal("Failed to create a credential: ", err)
	}

	// Create a CosmosDB client
	client, err := azcosmos.NewClientWithKey(c.url, cred, nil)
	if err != nil {
		log.Fatal("Failed to create Azure Cosmos DB client: ", err)
	}

	return client
}
