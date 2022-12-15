package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

// TestAPIOverHTTP : basic API testing with device ping simultion
// Despite being an IoT scenario we havent got much on the ground.
// hence we here simulate the device pinging the sever
func TestAPIOverHTTP(t *testing.T) {
	url := "http://localhost:8080/sensors/start"
	data := struct {
	}{}
	// Building the request here
	byt, err := json.Marshal(data)
	if err != nil {
		t.Error("Failed to marshal json data")
		return
	}
	fmt.Println()
	bytBody := bytes.NewReader(byt)
	req, _ := http.NewRequest("GET", url, bytBody)
	client := http.Client{
		Timeout: 3 * time.Second,
	}
	resp, err := client.Do(req)
	// fmt.Println("******RESP", resp)
	// var DB *gorm.DB
	// DB.Delete(&sensorPckg.Sensor{})
	assert.Equal(t, 200, resp.StatusCode, "Unexpected response code")
}
