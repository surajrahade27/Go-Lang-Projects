package main

/*
Author		: niranjan.awati@in.ncs-i.com
Date		:24-NOV-2022
This was developed to demo the capabilities of GoGIN framework and how it can easier to setup a running server with a minimal set of lines of code
*/
import (
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

func main() {
	// Setting up the logging configuration
	log.SetOutput(os.Stdout)
	log.SetReportCaller(false)
	log.SetLevel(log.DebugLevel)
	log.Info("Starting the go gin application")
	defer log.Warn("Shutting down the gin application server")

	gin.SetMode(gin.DebugMode)
	r := gin.Default()
	r.Use(CORS)
	r.GET("/ping", PingHandler)
	employees := r.Group("/employees", JsonLoadHandler)
	employees.GET("/", func(c *gin.Context) {
		val, _ := c.Get("jsonemployees") // val is a type of interface{}
		data := val.([]Employee)         // type casting
		c.AbortWithStatusJSON(http.StatusOK, gin.H{
			"total": len(data),
		})
	})
	employees.GET("/:id", EmployeeHandler)
	log.Fatal(r.Run(":8080"))
}
