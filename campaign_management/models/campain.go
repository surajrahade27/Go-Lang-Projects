package models

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

const DNS = "root:Password@1@tcp(localhost:3306)/FairPrise?parseTime=true"

type Campaing struct {
	gorm.Model

	Id             int    `gorm:"AUTO_INCREMENT" json:"id"`
	Name           string `json:"name"`
	Title          string `json:"title"`
	StartDate      string `json:"startDate"`
	EndDate        string `json:"endDate"`
	Status         bool   `json:"status"`
	OrderDate      string `json:"orderDate"`
	CollectionDate string `json:"collectionDate"`
}

func InitialMigration() {
	DB, err = gorm.Open(mysql.Open(DNS), &gorm.Config{})
	if err != nil {
		fmt.Println(err.Error())
		panic("Cannot connect to DB")
	}
	DB.AutoMigrate(&Campaing{})
}

func GetCampaings(c *gin.Context) {
	var Campaings []Campaing
	DB.Find(&Campaings)
	if Campaings == nil || len(Campaings) == 0 {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.IndentedJSON(http.StatusOK, Campaings)
	}
}

func ReadJson(c *gin.Context) {

	var campaings []Campaing
	jsonFile, err := os.Open("./MOCK_DATA.json")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Successfully Opened campaings.json")
	defer jsonFile.Close()
	byteValue, _ := io.ReadAll(jsonFile)

	json.Unmarshal(byteValue, &campaings)
	// fmt.Println(campaing)

	for i, s := range campaings {
		DB.Create(&s)
		fmt.Println(i, s)
	}
	c.IndentedJSON(http.StatusOK, campaings)

}

func GetCampaing(c *gin.Context) {
	var campaing Campaing
	DB.First(&campaing, c.Param("Id"))
	c.IndentedJSON(http.StatusOK, campaing)
}
func DeleteCampaing(c *gin.Context) {
	var campaing Campaing

	DB.First(&campaing, c.Param("id"))
	if err := DB.Where("id = ?", c.Param("id")).First(&campaing).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	if campaing.Id != 0 {
		DB.Delete(&campaing)
		c.JSON(200, gin.H{"success": "campaing #" + c.Param("id") + " deleted"})
	} else {
		c.JSON(404, gin.H{"error": "User not found"})
	}

}
func CreateCampaing(c *gin.Context) {

	var input Campaing
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}
	// Create book
	campaing := Campaing{Id: input.Id, Name: input.Name, Title: input.Title, Status: input.Status, StartDate: input.StartDate, EndDate: input.EndDate, OrderDate: input.OrderDate, CollectionDate: input.CollectionDate}
	DB.Create(&campaing)
	c.JSON(http.StatusOK, campaing)
}
func UpdateCampaing(c *gin.Context) {

	var campaing Campaing
	DB.First(&campaing, c.Param("id"))
	if err := DB.Where("id = ?", c.Param("id")).First(&campaing).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	// Validate input
	var input Campaing
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	DB.Model(&campaing).Updates(input)

	c.JSON(http.StatusOK, gin.H{"data": campaing})

}
