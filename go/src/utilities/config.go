package utilities

import (
	"encoding/json"
	"log"
	"os"
)

type APIConfiguration struct {
	Host string `json:"hostname"`
	Port string `json:"port"`
}

/*
 * Utility function for parsing api config on startup.
 */
func ParseAPIConfig() APIConfiguration {

	confFile, err := os.Open("/app/api-config.json")
	if err != nil {
		log.Println(err)
	}
	defer confFile.Close()

	decoder := json.NewDecoder(confFile)
	configData := APIConfiguration{}
	err = decoder.Decode(&configData)
	if err != nil {
		log.Println(err)
		panic(err)
	}

	return configData
}
