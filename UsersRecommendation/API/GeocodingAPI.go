package api

// send GET request to get the Geocoding API to get the longtitude and latitude of the address

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"
	"os"
)

// get the API key for Geoapify
type Config struct {
	APIKey string `json:"api_key"`
}

// define the response from the API
type GeocodeResponse struct {
	Status string `json:"status"`
	Lat string `json:"lat"`
	Long string `json:"long"`
}

// function to send the GET request to get the user's address latitude and longtitude
func GetLongLat(user_key int, address_line_1 string, city string, province string, country string, postal_code string, user_longlat_map map[int]map[string]string) {
	// Split address line by comma and space
	parts1 := strings.Split(address_line_1, " ");

	// first address line
	houseNumber := parts1[0];
	streetName := "";
	for i := 1; i < len(parts1)-1; i++ {
		streetName += parts1[i] + "%20";
	}
	streetName += parts1[len(parts1)-1];

	// get the api key
	configFile, err := os.Open("./config.json");
	if err != nil {
		log.Fatal(`Error opening config file: `, err);
	}
	defer configFile.Close();

	var config Config;
	err = json.NewDecoder(configFile).Decode(&config);
	if err != nil {
		log.Fatal(`Error decoding config file: `, err);
	}

	apiKey := config.APIKey;
	
	// Construct the URL
	url := `https://api.geoapify.com/v1/geocode/search?housenumber=` + houseNumber + `&street=` + streetName + `&postcode=` + postal_code + `&city=` + city + `&country=` + country + `&apiKey=` + apiKey;

	// Send the GET request to the API
	response, err := http.Get(url);
	if err != nil {
		log.Fatal("Error sending GET request: ", err);
	}
	defer response.Body.Close();

	// Check the response status
	if response.StatusCode != http.StatusOK {
		log.Fatal("Request failed with status:", response.Status);
	}

	// Decode the JSON response 
	var geocodeResponse GeocodeResponse;
	err = json.NewDecoder(response.Body).Decode(&geocodeResponse);
	if err != nil {
		log.Fatal("Error decoding JSON response: ", err);
	}

	// Check for the status in the response
	if geocodeResponse.Status != "OK" {
		log.Fatal("Geocoding failed with status: ", geocodeResponse.Status);
	}

	// if nothing is wrong, then return the longtitude and the latitude of the user's address
	fmt.Println("Latitude: ", geocodeResponse.Lat);
	fmt.Println("Longtitude: ", geocodeResponse.Long);
	user_longlat_map[user_key]["Latitude"] = geocodeResponse.Lat;
	user_longlat_map[user_key]["Longtitude"] = geocodeResponse.Long;
}