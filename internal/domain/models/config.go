package models

type Config struct {
	Port        int    `env:"PORT,required"`
	BaseUrl     string `env:"BASEURL,required"`
	Env         string `env:"ENV,required"`
	CosmosDbUrl string `env:"COSMOSDB_URL,required"`
	CosmosDbKey string `env:"COSMOSDB_KEY,required"`
}
