package repositories

import (
	"database/sql"
	"dependencies/models"
	"encoding/json"
	"fmt"
	"log"
	"os"

	_ "github.com/lib/pq"
)

var VehicleDB *sql.DB

type DBConfiguration struct {
	Name     string `json:"name"`
	Host     string `json:"host"`
	Port     int16  `json:"port"`
	User     string `json:"user"`
	Password string `json:"password"`
}

func EstablishDBConnection() {

	confFile, err := os.Open("/app/db-config.json")
	if err != nil {
		log.Println(err)
	}
	defer confFile.Close()

	decoder := json.NewDecoder(confFile)
	configData := DBConfiguration{}
	err = decoder.Decode(&configData)
	if err != nil {
		log.Println(err)
		panic(err)
	}

	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		configData.Host, configData.Port, configData.User, configData.Password, configData.Name)

	db, err := sql.Open("postgres", psqlInfo)

	log.Println(err)
	if err != nil {
		panic(err)
	}

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	VehicleDB = db

}

func QueryAllVehicles() []models.Vehicle {

	var errVehicle models.Vehicle
	errArr := []models.Vehicle{errVehicle}

	sqlStatement := "SELECT * FROM vehicles;"
	rows, err := VehicleDB.Query(sqlStatement)

	if err != nil {
		errVehicle.VehicleError = err
		return errArr
	}

	defer rows.Close()

	var vehicles []models.Vehicle
	for rows.Next() {

		var vehicle models.Vehicle
		err = rows.Scan(&vehicle.Make, &vehicle.Model, &vehicle.Package, &vehicle.Color,
			&vehicle.Year, &vehicle.Category, &vehicle.Mileage, &vehicle.Price, &vehicle.Id)

		if err != nil {
			errVehicle.VehicleError = err
			return errArr
		}

		vehicles = append(vehicles, vehicle)
	}

	return vehicles
}

func QueryVehicleByID(id string) models.Vehicle {

	sqlStatement := "SELECT * FROM vehicles WHERE vehicleid=$1;"

	var vehicle models.Vehicle
	row := VehicleDB.QueryRow(sqlStatement, id)
	err := row.Scan(&vehicle.Make, &vehicle.Model, &vehicle.Package, &vehicle.Color,
		&vehicle.Year, &vehicle.Category, &vehicle.Mileage, &vehicle.Price, &vehicle.Id)

	if err != nil {
		vehicle.VehicleError = err
	}

	switch err {
	case sql.ErrNoRows:
		log.Println("No rows were returned!")
	case nil:
		log.Println(vehicle)
	default:
		panic(err)
	}

	return vehicle
}
