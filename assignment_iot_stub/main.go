package main

import (
	"fmt"
	"iot_stud/sensorPckg"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func initializeRouter() {
	r := mux.NewRouter()
	r.HandleFunc("/sensors", sensorPckg.GetSensors).Methods("GET")
	r.HandleFunc("/sensors/{id}", sensorPckg.GetSensor).Methods("GET")
	r.HandleFunc("/sensors/create", sensorPckg.CreateSensor).Methods("POST")
	r.HandleFunc("/sensors/{id}", sensorPckg.UpdateSensor).Methods("PUT")
	r.HandleFunc("/sensors/{id}", sensorPckg.DeleteSensor).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8080", r))
}
func main() {
	fmt.Println("Started on port :8080")
	sensorPckg.InitialMigration()
	initializeRouter()
}
