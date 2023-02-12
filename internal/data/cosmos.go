package data

import (
	"fmt"

	"github.com/Azure/azure-sdk-for-go/sdk/data/azcosmos"
)

type Cosmos struct {
	key string
	url string
}

func NewCosmos(key string, url string) *Cosmos {
	return &Cosmos{key: key, url: url}
}

func (c *Cosmos) GetClient() (*azcosmos.Client, error) {
	cred, err := azcosmos.NewKeyCredential(c.key)
	if err != nil {
		return nil, fmt.Errorf("failed to create Azure Cosmos credential")
	}

	// Create a CosmosDB client
	client, err := azcosmos.NewClientWithKey(c.url, cred, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create Azure Cosmos client")
	}

	return client, nil
}
