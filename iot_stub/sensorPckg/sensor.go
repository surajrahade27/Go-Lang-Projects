package sensorPckg

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB
var err error

const DNS = "root:Password@1@tcp(localhost:3306)/iot_stub?parseTime=true"

type Sensor struct {
	gorm.Model

	UnixTime string `json:"unixTime"`
	Temp     int    `json:"temp"`
	Co       int    `json:"co"`
	Co2      int    `json:"co2"`
}

func InitialMigration() {
	DB, err = gorm.Open(mysql.Open(DNS), &gorm.Config{})
	if err != nil {
		fmt.Println(err.Error())
		panic("Cannot connect to DB")
	}
	DB.AutoMigrate(&Sensor{})
}

func GetSensors(c *gin.Context) {
	var Sensors []Sensor
	DB.Find(&Sensors)
	if Sensors == nil || len(Sensors) == 0 {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.IndentedJSON(http.StatusOK, Sensors)
	}
}

func ReadJson(c *gin.Context) {

	var sensors []Sensor
	jsonFile, err := os.Open("./MOCK_DATA.json")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Successfully Opened sensors.json")
	defer jsonFile.Close()
	byteValue, _ := io.ReadAll(jsonFile)

	json.Unmarshal(byteValue, &sensors)
	// fmt.Println(sensor)

	for i, s := range sensors {
		DB.Create(&s)
		fmt.Println(i, s)
		// sh := gocron.NewScheduler(time.UTC)
		// sh.Every(3).Seconds().Do(func() {
		// 	fmt.Println("Every 3 seconds")
		// 	DB.Create(&s)
		// 	fmt.Println(i, s)
		// })
		// sh.StartAsync()
	}
	//data traverse
	// func(3) , with go call ,3 sec
	// dbCreate(sensor)

	// you can start running the scheduler in two different ways:
	// starts the scheduler asynchronously

	// starts the scheduler and blocks current execution path
	// s.StartBlocking()
	c.IndentedJSON(http.StatusOK, sensors)

}

func GetSensor(c *gin.Context) {
	var sensor Sensor
	DB.First(&sensor, c.Param("id"))
	c.IndentedJSON(http.StatusOK, sensor)
}

func CreateSensor(c *gin.Context) {

	var input Sensor
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}
	// Create book
	sensor := Sensor{UnixTime: input.UnixTime, Co: input.Co, Co2: input.Co2, Temp: input.Temp}
	DB.Create(&sensor)
	c.JSON(http.StatusOK, sensor)
}
func UpdateSensor(c *gin.Context) {

	var sensor Sensor
	DB.First(&sensor, c.Param("id"))
	if err := DB.Where("id = ?", c.Param("id")).First(&sensor).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	// Validate input
	var input Sensor
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	DB.Model(&sensor).Updates(input)

	c.JSON(http.StatusOK, gin.H{"data": sensor})

}

// Get model if exist
// func DeleteSensor(w http.ResponseWriter, r *http.Request) {
// 	w.Header().Set("Content-Type", "application/json")
// 	params :=
// 		mux.Vars(r)
// 	var sensor Sensor
// 	DB.Delete(&sensor, params["id"])
// 	DB.Save(&sensor)
// 	json.NewEncoder(w).Encode("The Sensor is Deleted Succesfully!")
// }
