package main

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type order struct {
	Items       []item    `json:"items"`
	CreatedDate time.Time `json:"createdDate"`
}

type item struct {
	Description string  `json:"description"`
	Price       float64 `json:"price"`
}

func (app *application) getOrdersHandler(c *gin.Context) {
	orders := []order{
		{
			Items: []item{
				{
					Description: "yo this is a description",
					Price:       1.50,
				},
				{
					Description: "this is another description",
					Price:       3.00,
				},
			},
			CreatedDate: time.Now(),
		},
	}

	c.JSON(http.StatusOK, orders)
}
