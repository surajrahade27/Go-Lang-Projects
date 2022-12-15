package main

import (
	"encoding/json"
	"io"
	"net/http"
	"os"
	"strconv"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

// PingHandler : handles the request when url has /ping
func PingHandler(c *gin.Context) {
	c.AbortWithStatusJSON(http.StatusOK, gin.H{
		"message": "Hi, this is from inside the server ping route",
	})
}
func CORS(c *gin.Context) {
	// First, we add the headers with need to enable CORS
	// Make sure to adjust these headers to your needs
	c.Header("Access-Control-Allow-Origin", "*")
	c.Header("Access-Control-Allow-Methods", "*")
	c.Header("Access-Control-Allow-Headers", "*")
	c.Header("Content-Type", "application/json")
	// Second, we handle the OPTIONS problem
	if c.Request.Method != "OPTIONS" {
		c.Next()
	} else {
		// Everytime we receive an OPTIONS request,
		// we just return an HTTP 200 Status Code
		// Like this, Angular can now do the real
		// request using any other method than OPTIONS
		c.AbortWithStatus(http.StatusOK)
	}
}

// JsonLoadHandler : will read json data from a file database and inject the same in the context
//
// For demo purposes we are just querying a small json file database
// Queries and filtering can follow in downstream in handlers
func JsonLoadHandler(c *gin.Context) {
	data := []Employee{}
	file, err := os.Open("MOCK_DATA.json")
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
	}
	byt, err := io.ReadAll(file)
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
	}
	if json.Unmarshal(byt, &data) != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
	}
	c.Set("jsonemployees", data)
	log.Debugf("We have about %d employees loaded", len(data))
}

// EmployeeHandler : handles single employee requests
func EmployeeHandler(c *gin.Context) {
	val, _ := c.Get("jsonemployees") // val is a type of interface{}
	data := val.([]Employee)         // type casting
	paramIDstr, _ := c.Params.Get("id")
	paramID, err := strconv.Atoi(paramIDstr)
	if err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}
	log.WithFields(log.Fields{
		"param": paramID,
	}).Debug("From inside the employee handler")
	for _, emp := range data {
		if emp.ID == paramID {
			c.AbortWithStatusJSON(http.StatusOK, gin.H{
				"email":      emp.Email,
				"first_name": emp.FirstName,
			})
			// This is a vital return statement, unless otherwise the code will fall thru and the status would be overwritten
			return
		}
	}
	c.AbortWithStatus(http.StatusNotFound)

}
