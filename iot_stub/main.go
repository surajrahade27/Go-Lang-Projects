package main

import (
	"fmt"
	"iot_stud/sensorPckg"

	"github.com/gin-gonic/gin"
)

func initializeRouter() {
	router := gin.Default()
	sensors := router.Group("sensors")
	sensors.GET("/getSensors", sensorPckg.GetSensors)
	sensors.GET("/:id", sensorPckg.GetSensor)
	sensors.POST("/createSensors", sensorPckg.CreateSensor)
	sensors.POST("/update/:id", sensorPckg.UpdateSensor)
	sensors.GET("/start", sensorPckg.ReadJson)
	router.Run("localhost:8080")
}
func main() {
	fmt.Println("Started on port :8080")
	sensorPckg.InitialMigration()
	initializeRouter()

	//read json data
}
