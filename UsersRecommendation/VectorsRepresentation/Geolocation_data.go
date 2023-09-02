package vectorsrepresentation

// this file will get the longtitude and the latitude from the API and use it to represent the user's location as vectors
import (
	"math"
	"StudyHub-Matching-Algorithm/UsersRecommendation/API"
)

// a struct to represent the cartesian coordinates
type CartesianCoord struct {
	X float64
	Y float64
	Z float64
}

// helper function to send GET request to get the latitude and longtitude of the user's address
func Get_All_UsersLongLatMap(users_address map[int]map[string]string, pre_found_users_address map[int]map[string]float64) map[int]map[string]float64 {
	// send GET request to get the longtitude and latitude of the user's address
	users_long_lat_map := make(map[int]map[string]float64);
	for i := range users_address {
		if pre_found_users_address[i] != nil {
			users_long_lat_map[i] = pre_found_users_address[i];
			continue; // skip if the user's long lat has been found before
		}
		users_long_lat_map[i] = make(map[string]float64); // initialize a map for this user
		// get the long and lat for this user's address
		user_id := i;
		address_line_1 := users_address[i]["address_line_1"];
		city := users_address[i]["city"];
		province := users_address[i]["province"];
		country := users_address[i]["country"];
		postal_code := users_address[i]["postal_code"];
		api.GetLongLat(user_id, address_line_1, city, province, country, postal_code, users_long_lat_map);	
	}
	return users_long_lat_map;
}

// receive a users map with long lat -> convert into cartesian vectors
func LatLongToCartesian(users_long_lat map[int]map[string]float64, pre_calculated_long_lat map[int]CartesianCoord) map[int]CartesianCoord {
	// all user's address cartesian vectors
	cartesian_vectors_map := make(map[int]CartesianCoord);

	for i := range users_long_lat {
		// check if this user's coordinates have been calculated before
		_, exists := pre_calculated_long_lat[i];
		if exists {
			cartesian_vectors_map[i] = pre_calculated_long_lat[i];
			continue; 
		}
		// get user's key and long lat
		user_key := i;
		user_longtitude := users_long_lat[i]["Longtitude"];
		user_latitude := users_long_lat[i]["Latitude"];

		// Convert degrees to radians
		longRad := user_longtitude * (math.Pi / 180);
		latRad := user_latitude * (math.Pi / 180);

		const EARTH_RADIUS = 6371.0; // Earth's radius in kilometers

		// get the cartesian coordinates
		x_coord := EARTH_RADIUS * math.Cos(latRad) * math.Cos(longRad);
		y_coord := EARTH_RADIUS * math.Cos(latRad) * math.Sin(longRad);
		z_coord := EARTH_RADIUS * math.Sin(latRad);

		cartesian_vectors_map[user_key] = CartesianCoord{X: x_coord, Y: y_coord, Z: z_coord}; // assign cartesian coords object for a specific user
	}

	return cartesian_vectors_map;
}


