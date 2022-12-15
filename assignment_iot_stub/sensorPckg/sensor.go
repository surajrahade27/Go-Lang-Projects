package sensorPckg

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB
var err error

const DNS = "root:Password@1@tcp(localhost:3306)/iot_stub?parseTime=true"

type Sensor struct {
	gorm.Model

	UnixTime int64 `json:"unixTime"`
	Temp     int   `json:"temp"`
	Co       int   `json:"co"`
	Co2      int   `json:"co2"`
}

func InitialMigration() {
	DB, err = gorm.Open(mysql.Open(DNS), &gorm.Config{})
	if err != nil {
		fmt.Println(err.Error())
		panic("Cannot connect to DB")
	}
	DB.AutoMigrate(&Sensor{})
}

func GetSensors(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	var Sensors []Sensor
	DB.Find(&Sensors)
	json.NewEncoder(w).Encode(Sensors)

}

func GetSensor(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params :=
		mux.Vars(r)
	var sensor Sensor
	DB.First(&sensor, params["id"])
	json.NewEncoder(w).Encode(sensor)

}

func CreateSensor(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	var sensor Sensor
	json.NewDecoder(r.Body).Decode(&sensor)
	DB.Create(&sensor)
	json.NewEncoder(w).Encode(sensor)
}

func UpdateSensor(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params :=
		mux.Vars(r)
	var sensor Sensor
	DB.First(&sensor, params["id"])
	json.NewDecoder(r.Body).Decode(&sensor)
	DB.Save(&sensor)
	json.NewEncoder(w).Encode(sensor)
}

func DeleteSensor(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params :=
		mux.Vars(r)
	var sensor Sensor
	DB.Delete(&sensor, params["id"])
	DB.Save(&sensor)
	json.NewEncoder(w).Encode("The Sensor is Deleted Succesfully!")
}
