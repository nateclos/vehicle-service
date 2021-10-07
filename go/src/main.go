package main

import (
	"dependencies/controllers"
	"dependencies/repositories"
	"dependencies/utilities"

	"github.com/gin-gonic/gin"
)

func main() {

	APIconfig := utilities.ParseAPIConfig()
	repositories.EstablishDBConnection()
	defer repositories.VehicleDB.Close()

	var router = gin.Default()
	router.GET("/", controllers.GetVehicles)
	router.GET("/:id", controllers.GetVehicleByID)
	router.Run(APIconfig.Host + ":" + APIconfig.Port)
}
