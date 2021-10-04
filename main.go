package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type vehicle struct {
	Make     string `json:"make"`
	Model    string `json:"model"`
	Package  string `json:"package"`
	Color    string `json:"color"`
	Year     int16  `json:"year"`
	Category string `json:"category"`
	Mileage  int16  `json:"mileage"`
	Price    int32  `json:"price"`
	Id       string `json:"id"`
}

var vehicles = []vehicle{
	{Make: "Ford", Model: "F-150"},
	{Make: "Chevy", Model: "Silverado"},
}

func main() {

	var router = gin.Default()
	router.GET("/", getVehicles)
	router.GET("/:id", getVehicleByID)
	router.Run("localhost:8080")
}

func getVehicles(context *gin.Context) {
	context.IndentedJSON(http.StatusOK, vehicles)
}

func getVehicleByID(context *gin.Context) {
	context.IndentedJSON(http.StatusOK, vehicles)
}
