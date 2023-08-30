package main

import (
	"StudyHub-Matching-Algorithm/UsersRecommendation/API"
	"fmt"
)

// ----------------- MAIN -------------------- //
func main() {

	// test TF-IDF
	// testSuite := &TF_IDF_Test{}
	// testSuite.runAllTests(&testing.T{})
	// testSuite.printReport()

	// test Date-Time

	// test Geolocation
	user_key := 10234
	address_line_1 := "75 Skelton Boulevard"
	city := "Brampton"
	province := "Ontario"
	country := "Canada"
	postal_code := "L6V2S2"

	// initialize a map to store all the user's address latitude and longtitude
	user_longlat_map := make(map[int]map[string]float64)

	for i := 0; i < 10; i++ { // Replace 10 with the desired size
		user_longlat_map[i] = make(map[string]float64)
	}

	// test the function
	api.GetLongLat(
		user_key,
		address_line_1,
		city,
		province,
		country,
		postal_code,
		user_longlat_map,
	)

	// traverse through the user_longlat_map 
	fmt.Println(user_longlat_map);
}
