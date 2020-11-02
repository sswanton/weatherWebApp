// Starter code taken from https://tutorialedge.net/golang/creating-restful-api-with-golang/

package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

// Forecast Struct for holding weather data from API
type Forecast struct {
	// nameofattribute string `json:"nameofattribute"`
	Temperature int `Json:"properties":{"periods":[0]`
}

// Array for temp holding of data
var forecastData []Forecast

func mainPage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome!")
	fmt.Println("Endpoint Hit: mainPage")
}

func handleRequests() {
	http.HandleFunc("/", mainPage)
	http.HandleFunc("/forecast", returnForecast)
	log.Fatal(http.ListenAndServe(":10000", nil))
}

func returnForecast(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint Hit: returnForecast")
	json.NewEncoder(w).Encode(forecastData)
}

// TBD integrate geolocation for query to weather API https://geocode.xyz/api#Go
// Weather API https://weather-gov.github.io/api/general-faqs
func main() {
	forecastData = []Forecast{
		Forecast{Temperature: 45},
		Forecast{Temperature: 50},
	}
	handleRequests()
}
