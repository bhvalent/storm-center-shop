package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (app *application) getOrdersHandler(c *gin.Context) {
	// change this to return data
	c.String(http.StatusOK, "it worked")
}
