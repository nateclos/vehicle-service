package controllers

import (
	"dependencies/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "your-password"
	dbname   = "calhounio_demo"
)

var vehicles = []models.Vehicle{
	{Make: "Ford", Model: "F-150"},
	{Make: "Chevy", Model: "Silverado"},
}

func GetVehicles(context *gin.Context) {
	context.IndentedJSON(http.StatusOK, vehicles)

}

func GetVehicleByID(context *gin.Context) {
	context.IndentedJSON(http.StatusOK, vehicles)
}
