package routes

import (
	"storm-center-shop/internal/app/controllers"
	"storm-center-shop/internal/domain/models"

	"github.com/gin-gonic/gin"
)

func Routes(a *models.Application) *gin.Engine {
	baseController := controllers.NewBaseController(a)

	routes := gin.Default()
	// routes.Use()
	routes.GET("/user/:id/orders", baseController.GetOrdersHandler)
	routes.POST("/order", baseController.CreateOrderHandler)
	return routes
}
