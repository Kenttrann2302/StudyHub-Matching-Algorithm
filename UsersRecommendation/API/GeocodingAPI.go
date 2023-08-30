package api

// send GET request to get the Geocoding API to get the longtitude and latitude of the address

import (
	"encoding/json"
	"log"
	"net/http"
	"strings"
	"os"
	"io"
)

// get the API key for Geoapify
type Config struct {
	APIKey string `json:"api_key"`
}

// define the response from the API
type GeocodeResponse struct {
	Status string `json:"status"`
	Lat float64 `json:"lat"`
	Long float64 `json:"long"`
}

// function to send the GET request to get the user's address latitude and longtitude
func GetLongLat(user_key int, address_line_1 string, city string, province string, country string, postal_code string, user_longlat_map map[int]map[string]float64) {
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
	configFile, err := os.Open("config.json");
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
		log.Fatal("Request failed with status: ", response.Status);
	}

	// read the body as a string
	body, err := io.ReadAll(response.Body);
	if err != nil {
		log.Fatal("Error while reading the response body: ", err);
	}

	// Decode the JSON response
	// Declare a variable to store the deserialized response
	var json_response map[string]interface{};

	// Parse the JSON response
	marshal_json_err := json.Unmarshal(body, &json_response);
	if marshal_json_err != nil {
		log.Fatal("Error while marshal the json resonse: ", marshal_json_err);
	}

	// If nothing is wrong with parsing the JSON response -> extract required values from the response
	var geocodeResponse GeocodeResponse;
	if features, ok := json_response["features"].([]interface{}); ok && len(features) > 0 {
		feature := features[0].(map[string]interface{});
		properties := feature["properties"].(map[string]interface{});
		latitude := properties["lat"].(float64);
		longtitude := properties["lon"].(float64);
		
		// set the longtitude and latitude for the struct
		geocodeResponse.Status = "OK";
		geocodeResponse.Lat = latitude;
		geocodeResponse.Long = longtitude;

		// if nothing is wrong, then return the longtitude and the latitude of the user's address
		user_longlat_map[user_key]["Latitude"] = geocodeResponse.Lat;
		user_longlat_map[user_key]["Longtitude"] = geocodeResponse.Long;
	} else {
		log.Fatal("Cannot find the longtitude and latitude for this address!");
	}
}