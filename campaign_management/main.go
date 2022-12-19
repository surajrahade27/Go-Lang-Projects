package main

import (
	campaingPckg "campaing_management/models"
	"fmt"

	"github.com/gin-gonic/gin"
)

func initializeRouter() {
	router := gin.Default()
	campaings := router.Group("campaings")
	campaings.GET("/getCampaings", campaingPckg.GetCampaings)
	campaings.GET("/:id", campaingPckg.GetCampaing)
	campaings.DELETE("/:id", campaingPckg.DeleteCampaing)
	campaings.POST("/createCampaing", campaingPckg.CreateCampaing)
	campaings.POST("/update/:id", campaingPckg.UpdateCampaing)
	campaings.GET("/start", campaingPckg.ReadJson)
	router.Run("localhost:8080")
}
func main() {
	fmt.Println("Started on port :8080")
	campaingPckg.InitialMigration()
	initializeRouter()

	//read json data
}
