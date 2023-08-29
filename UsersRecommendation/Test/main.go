package main

import (
	"StudyHub-Matching-Algorithm/UsersRecommendation/API"
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
	user_longlat_map := make(map[int]map[string]string)

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
}
