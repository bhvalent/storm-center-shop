package main

import "github.com/gin-gonic/gin"

func (app *application) routes() *gin.Engine {
	routes := gin.Default()
	routes.GET("/orders", app.getOrdersHandler)
	return routes
}
