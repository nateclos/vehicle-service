package controllers

import (
	"dependencies/models"
	"dependencies/repositories"
	"net/http"

	"github.com/gin-gonic/gin"
)

/*
 * Controller function for retreiving all vehicles currently in the database.
 */
func GetVehicles(context *gin.Context) {

	vehicles := repositories.QueryAllVehicles()

	if (vehicles[0].VehicleError != models.Vehicle{}.VehicleError) {
		context.IndentedJSON(http.StatusBadRequest, vehicles[0].VehicleError)
	} else {
		context.IndentedJSON(http.StatusOK, vehicles)
	}

}

/*
 * Controller function for retreiving a vehicle based on the primary key which is ID.
 */
func GetVehicleByID(context *gin.Context) {

	id := context.Param("id")
	vehicle := repositories.QueryVehicleByID(string(id))

	// If the error field is empty, we can assume that no error was encountered in data access
	if (vehicle.VehicleError != models.Vehicle{}.VehicleError) {
		context.IndentedJSON(http.StatusBadRequest, vehicle.VehicleError)
	} else {
		context.IndentedJSON(http.StatusOK, vehicle)
	}

}
