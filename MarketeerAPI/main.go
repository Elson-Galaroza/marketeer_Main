package main

import (
	"marketeer/api"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	// System
	router.GET("/marketeer/systemcheck", api.SystemCheck)

	// User Registration and Login
	router.POST("/marketeer/registerUser", api.RegisterUser)
	router.POST("/marketeer/login", api.LogInUser)

	router.Run(":80")
}
