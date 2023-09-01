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
func CalculateGeolocationCosineSimilarity(geolocation_vectors map[int]CartesianCoord, vectors_magnitude map[int]float64) [][]float64 {
	cosineResult := 
}