package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
)

const (
	// TODO:  make sure you populate the url for your server here
	testURL string = ""
)

// Reading : unit sensor reading that is pushed onto the api via POST

type Reading struct {
	gorm.Model

	UnixTime string `json:"unixTime"`
	Temp     int    `json:"temp"`
	Co       int    `json:"co"`
	Co2      int    `json:"co2"`
}

// sendRequest : given a reading this will fire a request onto the api and return the http response
func sendRequest(r Reading) (*http.Response, error) {
	byt, err := json.Marshal(r)
	if err != nil {
		// t.Error("Failed to marshal json data")
		return nil, err
	}
	bytBody := bytes.NewReader(byt)
	req, _ := http.NewRequest("POST", testURL, bytBody)
	client := http.Client{
		Timeout: 5 * time.Second,
	}
	resp, err := client.Do(req)
	if err != nil {
		// t.Error("failed to send request to server: check your internet connection")
		return nil, err
	}
	return resp, nil
}

// TestAPIOverHTTP : basic API testing with device ping simultion
// Despite being an IoT scenario we havent got much on the ground.
// hence we here simulate the device pinging the sever
func TestAPIOverHTTP(t *testing.T) {
	// Setting up the data
	data200OK := []Reading{
		{Tm: time.Now().Unix(), Temp: 34.0, Co: 13.0, Co2: 600.09},
		{Tm: time.Now().Unix(), Temp: 34.4, Co: 13.0, Co2: 600.09},
		{Tm: time.Now().Unix(), Temp: 34.6, Co: 0.00, Co2: 600.09},
		{Tm: time.Now().Unix(), Temp: 0.6, Co: 13.0, Co2: 0.00},
	} // all the data that when posted will get back 200 ok from the api
	for _, d := range data200OK {
		// Building the request here
		resp, err := sendRequest(d)
		if err != nil {
			t.Errorf("Failed request: %s", err)
			continue
		}
		assert.NotNil(t, resp, "Unexpected nil response")
		assert.Equal(t, 200, resp.StatusCode, "Unexpected response code")
	}
	// The API should prevent bogus values from getting into the database
	// API is the last line of defence when it comes to data filtering
	data400Bad := []Reading{
		{Tm: 0.000, Temp: 34.0, Co: 13.0, Co2: 600.09},                                         // time cannot be zero
		{Tm: time.Now().Unix(), Temp: 500.00, Co: 13.0, Co2: 600.09},                           // temp sensor has a limit, anything more than 50oDeg is not possible for a 5V sensor
		{Tm: time.Now().Unix(), Temp: 34.6, Co: -1.00, Co2: 600.09},                            // negative carbon monoxide, unreal
		{Tm: time.Now().Unix(), Temp: 34.6, Co: 13.0, Co2: 56546456456456},                     // co2 value is unreal
		{Tm: time.Now().Add(24 * time.Hour).Unix(), Temp: 34.6, Co: 13.0, Co2: 56546456456456}, // time stamp of future
	} // all the data that when posted will get back 200 ok from the api
	for _, d := range data400Bad {
		// Building the request here
		resp, err := sendRequest(d)
		if err != nil {
			t.Errorf("Failed request: %s", err)
			continue
		}
		assert.NotNil(t, resp, "Unexpected nil response")
		assert.Equal(t, 400, resp.StatusCode, "Unexpected response code")
	}
}
