package controllers

import (
	"dependencies/models"
	"dependencies/repositories"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetVehicles(context *gin.Context) {

	vehicles := repositories.QueryAllVehicles()

	if (vehicles[0].VehicleError != models.Vehicle{}.VehicleError) {
		context.IndentedJSON(http.StatusBadRequest, vehicles[0].VehicleError)
	} else {
		context.IndentedJSON(http.StatusOK, vehicles)
	}

}

func GetVehicleByID(context *gin.Context) {

	id := context.Param("id")
	vehicle := repositories.QueryVehicleByID(string(id))

	if (vehicle.VehicleError != models.Vehicle{}.VehicleError) {
		context.IndentedJSON(http.StatusBadRequest, vehicle.VehicleError)
	} else {
		context.IndentedJSON(http.StatusOK, vehicle)
	}

}
