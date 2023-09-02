package CosineSimilarity

// this file will calculate the cosine similarity between 2 user's geolocation cartesian vectors
import (
	"math"
)

// a struct represents user's address x, y, z coordinates 
type CartesianCoord struct {
	X float64
	Y float64
	Z float64
}

// function to calculate the magnitude between each user's geolocation vector
func CalculateGeolocationMagnitude(geolocation_vectors map[int]CartesianCoord) map[int]float64 {
	magnitude_result := make(map[int]float64);
	
	// calculate the vector's magnitude for each user
	for i := range geolocation_vectors {
		user_key := i;
		vec := geolocation_vectors[i];
		magnitude := math.Sqrt(vec.X*vec.X + vec.Y*vec.Y + vec.Z*vec.Z);
		magnitude_result[user_key] = magnitude;
	}

	return magnitude_result;
}

// function to calculate the cosine similarity between each user's geolocation vector
func CalculateGeolocationCosineSimilarity(geolocation_vectors map[int]CartesianCoord, vectors_magnitude map[int]float64, pre_calculated_vectors [][]float64) [][]float64 {
	cosineResult := make([][]float64, len(geolocation_vectors));
	for i := range geolocation_vectors {
		cosineResult[i] = make([]float64, len(geolocation_vectors));
	}

	for first_user, cartesian_vector_1 := range geolocation_vectors {
		for sec_user, cartesian_vector_2 := range geolocation_vectors {
			// ignore the diagonal
			if first_user == sec_user {
				continue;
			}
			// check if the cosine similarity has been calculated before
			if first_user < len(pre_calculated_vectors) || sec_user < len(pre_calculated_vectors[0]) {
				cosineResult[first_user][sec_user] = pre_calculated_vectors[first_user][sec_user];
				continue;
			}
			// find the dotProduct betweeen 2 users
			dotProduct := cartesian_vector_1.X * cartesian_vector_2.X + cartesian_vector_1.Y * cartesian_vector_2.Y + cartesian_vector_1.Z * cartesian_vector_2.Z;

			// get 2 vectors magnitudes
			magnitudeVec1 := vectors_magnitude[first_user];
			magnitudeVec2 := vectors_magnitude[sec_user];

			cosineResult[first_user][sec_user] = dotProduct / (magnitudeVec1 * magnitudeVec2);
		}
	}
	return cosineResult;
}