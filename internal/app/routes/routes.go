package routes

import (
	"storm-center-backend/internal/app/controllers"
	"storm-center-backend/internal/domain/models"

	"github.com/gin-gonic/gin"
)

func Routes(a *models.Application) *gin.Engine {
	baseController := controllers.NewBaseController(a)

	routes := gin.Default()
	routes.GET("/user/:id/orders", baseController.GetOrdersHandler)
	routes.POST("/order", baseController.CreateOrderHandler)
	return routes
}
