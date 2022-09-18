package main

import "github.com/gin-gonic/gin"

func (app *application) routes() *gin.Engine {
	routes := gin.Default()
	routes.GET("/orders", app.getOrdersHandler)
	routes.POST("/order", app.saveOrderHandler)
	return routes
}
