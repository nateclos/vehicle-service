package main

import (
	"encoding/json"
	"log"
	"os"

	"github.com/gin-gonic/gin"
)

type Configuration struct {
	Host string `json:"Hostname"`
	Port string `json:"Port"`
}

func main() {

	config := decodeConfig()
	log.Println(config.Host)

	var router = gin.Default()
	router.GET("/", getVehicles)
	router.GET("/:id", getVehicleByID)
	router.Run(config.Host + config.Port)
}

func decodeConfig() Configuration {

	confFile, err := os.Open("/app/config.json")
	if err != nil {
		log.Println(err)
	}
	defer confFile.Close()

	decoder := json.NewDecoder(confFile)
	config := Configuration{}
	err = decoder.Decode(&config)
	if err != nil {
		log.Println(err)
		panic(err)
	}

	return config
}
