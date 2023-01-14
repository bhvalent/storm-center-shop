package main

import (
	"net/http"
	"storm-center-backend/cmd/domain"

	"github.com/gin-gonic/gin"
)

func (app *application) getOrdersHandler(c *gin.Context) {
	userId := c.Param("id")
	orders := domain.GetOrders(c, userId)
	c.JSON(http.StatusOK, gin.H{"orders": orders})
}

func (app *application) saveOrderHandler(c *gin.Context) {
	var data domain.Order
	var err error
	if err = c.BindJSON(&data); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": c.Request.Body})
		return
	}

	order, err := domain.SaveOrder(c, data)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
	}
	c.JSON(http.StatusCreated, order)
}
