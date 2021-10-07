package models

// Struct representation of a vehicle.
type Vehicle struct {
	Make         string `json:"make"`
	Model        string `json:"model"`
	Package      string `json:"package"`
	Color        string `json:"color"`
	Year         int16  `json:"year"`
	Category     string `json:"category"`
	Mileage      int32  `json:"mileage"`
	Price        int32  `json:"price"`
	Id           string `json:"id"`
	VehicleError string `json:"error"`
}
