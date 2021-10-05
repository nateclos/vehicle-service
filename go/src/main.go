package main

import (
	"github.com/gin-gonic/gin"
)

func main() {

	var router = gin.Default()
	router.GET("/", getVehicles)
	router.GET("/:id", getVehicleByID)
	router.Run("localhost:8080")
}
