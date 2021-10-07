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

/*
 * Establishes database connection.
 */
func EstablishDBConnection() {

	// Opening db config file.
	confFile, err := os.Open("/app/db-config.json")
	if err != nil {
		log.Println(err)
	}
	defer confFile.Close()

	// Scanning config file into config struct.
	decoder := json.NewDecoder(confFile)
	configData := DBConfiguration{}
	err = decoder.Decode(&configData)
	if err != nil {
		log.Println(err)
		panic(err)
	}

	// Creating connection string.
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		configData.Host, configData.Port, configData.User, configData.Password, configData.Name)

	db, err := sql.Open("postgres", psqlInfo)

	if err != nil {
		log.Println(err)
		panic(err)
	}

	// Sets the global db instance pointer.
	VehicleDB = db

}

/*
 * Data access function to get all vehicles currently in the database.
 */
func QueryAllVehicles() []models.Vehicle {

	// Dummy response structs for if an error is encountered.
	var errVehicle models.Vehicle
	errArr := []models.Vehicle{errVehicle}

	// Check for if the database connection is active.
	dbUp := testDBConnection()
	if !dbUp {
		errArr[0].VehicleError = "ERROR: Database connection not active!"
		return errArr
	}

	// Querying for all rows.
	sqlStatement := "SELECT * FROM vehicles;"
	rows, err := VehicleDB.Query(sqlStatement)

	// If an error is encountered when querying.
	if err != nil {
		errVehicle.VehicleError = err.Error()
		return errArr
	}

	// Defer closing the data from the DB until after execution is complete.
	defer rows.Close()

	// Iterate over rows and scan rows into vehicle structs.
	var vehicles []models.Vehicle
	for rows.Next() {

		var vehicle models.Vehicle
		err = rows.Scan(&vehicle.Make, &vehicle.Model, &vehicle.Package, &vehicle.Color,
			&vehicle.Year, &vehicle.Category, &vehicle.Mileage, &vehicle.Price, &vehicle.Id)

		if err != nil {
			errVehicle.VehicleError = err.Error()
		}

		vehicles = append(vehicles, vehicle)
	}

	return vehicles
}

/*
 * Data access function to a specific vehicle from the database given an ID.
 */
func QueryVehicleByID(id string) models.Vehicle {

	// Our return vehicle struct.
	var vehicle models.Vehicle

	// Check for if the database connection is active.
	dbUp := testDBConnection()
	if !dbUp {
		vehicle.VehicleError = "ERROR: Database connection not active!"
		return vehicle
	}

	sqlStatement := "SELECT * FROM vehicles WHERE vehicleid=$1;"

	// Query for a single row and attempt to scan it into our vehicle struct.
	row := VehicleDB.QueryRow(sqlStatement, id)
	err := row.Scan(&vehicle.Make, &vehicle.Model, &vehicle.Package, &vehicle.Color,
		&vehicle.Year, &vehicle.Category, &vehicle.Mileage, &vehicle.Price, &vehicle.Id)

	// If an error is encountered store it.
	if err != nil {
		vehicle.VehicleError = string(err.Error())
	}

	switch err {
	case sql.ErrNoRows:
		log.Println("No rows were returned!")
	case nil:
		log.Println(vehicle)
	default:
		log.Println(err)
	}

	return vehicle
}

// Reuable function to test if the DB connection is active, true if it is, false otherwise.
func testDBConnection() bool {

	err := VehicleDB.Ping()
	if err != nil {
		log.Println(err)
		return false
	}
	return true
}
