package main

import (
	"marketeer/api"
	"marketeer/displaydataapi"
	"marketeer/itemapi"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	// System
	router.GET("/marketeer/systemcheck", api.SystemCheck)

	// User Registration and Login
	router.POST("/marketeer/registerUser", api.RegisterUser)
	router.POST("/marketeer/login", api.LogInUser)

	//item
	router.POST("/marketeer/sellitem", itemapi.SellItem)

	//display all item for sale
	router.GET("/marketeer/displayitemforsale", displaydataapi.RetrieveForSaleItems)

	router.Run(":80")
}
