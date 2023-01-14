package data

import (
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/data/azcosmos"
)

const endpoint = "https://localhost:8081"
const key = "C2y6yDjf5/R+ob0N8A7Cgv30VRDJIWEHLM+4QDU5DE2nQ9nDuVTqobD4b8mGGyPMbIZnqyMsEcaGQy67XIw/Jw=="

func GetContainer(cName string) *azcosmos.ContainerClient {
	cred, err := azcosmos.NewKeyCredential(key)
	if err != nil {
		log.Fatal("Failed to create a credential: ", err)
	}

	// Create a CosmosDB client
	client, err := azcosmos.NewClientWithKey(endpoint, cred, nil)
	if err != nil {
		log.Fatal("Failed to create Azure Cosmos DB client: ", err)
	}

	// Create container client
	containerClient, err := client.NewContainer("stormcenter", cName)
	if err != nil {
		log.Fatal("Failed to create a container client:", err)
	}

	return containerClient
}
