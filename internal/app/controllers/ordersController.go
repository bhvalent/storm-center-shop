package controllers

import (
	"net/http"
	"storm-center-backend/internal/data"
	"storm-center-backend/internal/domain/models"
	"storm-center-backend/internal/domain/services"
	"storm-center-backend/pkg/api"

	"github.com/gin-gonic/gin"
)

func (bc *BaseController) GetOrdersHandler(c *gin.Context) {
	cosmos := data.NewCosmos(bc.app.Config.CosmosDbKey, bc.app.Config.CosmosDbUrl)
	or := data.NewOrderRepository(cosmos)
	os := services.NewOrderService(or)

	userId := c.Param("id")
	orders := os.GetOrders(c, userId)
	c.JSON(http.StatusOK, gin.H{"orders": orders})
}

func (bc *BaseController) CreateOrderHandler(c *gin.Context) {
	var cor api.CreateOrderRequest
	var err error
	if err = c.BindJSON(&cor); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": c.Request.Body})
		return
	}

	cosmos := data.NewCosmos(bc.app.Config.CosmosDbKey, bc.app.Config.CosmosDbUrl)
	or := data.NewOrderRepository(cosmos)
	o := models.NewOrderFromCreateOrderRequest(cor, or)

	order, err := o.Create(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
	}

	response := orderToCreateOrderResponse(order)
	c.JSON(http.StatusCreated, response)
}

func orderToCreateOrderResponse(o models.Order) api.CreateOrderResponse {
	var items []api.RequestItem
	for _, itemEntity := range o.Items {
		items = append(items, itemToRequestItem(itemEntity))
	}

	return api.CreateOrderResponse{
		Id:          o.Id,
		UserId:      o.UserId,
		CreatedDate: o.CreatedDate,
		Items:       items,
	}
}

func itemToRequestItem(i models.Item) api.RequestItem {
	return api.RequestItem{
		Id:    i.Id,
		Name:  i.Name,
		Price: i.Price,
	}
}
