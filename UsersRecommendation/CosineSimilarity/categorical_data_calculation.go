package CosineSimilarity

// this file will calculate the each user's categorical vectors cosine similarity to establish the relationship between each user
import (
	"math"
)

// function to calculate the magnitude of each categorical vector
func CalculateCategoricalMagnitude(vector_map map[int][]int, pre_calculated_vector map[int]float64) map[int]float64 {
	magnitude_result := make(map[int]float64);

	// calculate the magnitude for each user's vector
	for user_key, vector := range vector_map {
		if _, exists := pre_calculated_vector[user_key]; exists {
			magnitude_result[user_key] = pre_calculated_vector[user_key];
			continue; // skip if the user's vector magnitude has already been calculated
		}
		var sumOfSquares float64;
		for _, value := range vector {
			sumOfSquares += float64(value) * float64(value);
		}
		magnitude_result[user_key] = math.Sqrt(sumOfSquares);
	}
	return magnitude_result;
}

// function to calculate the cosine similarity of the categorical data
func CalculateCategoricalCosineSimilarity(vector_map map[int][]int, pre_calculated_vector [][]float64, vectors_magnitude map[int]float64) [][]float64 {	
	cosineResult := make([][]float64, len(vector_map));
	for i := range vector_map {
		cosineResult[i] = make([]float64, len(vector_map));
	}

	// find the cosine similarities between 2 vectors in the vector map
	for first_user_key, first_vector := range vector_map {
		for sec_user_key, sec_vector := range vector_map {
			if first_user_key == sec_user_key {
				continue; // ignore the diagonals since that will be the same users
			}
			if first_user_key < len(pre_calculated_vector) || sec_user_key < len(pre_calculated_vector[0]) {
				cosineResult[first_user_key][sec_user_key] = pre_calculated_vector[first_user_key][sec_user_key];
				continue; // skip if the relationship has been calculated before
			}

			dotProduct := 0.0;
			// calculate the dot product 
			for index := 0; index < len(first_vector); index++ {
				dotProduct += float64(first_vector[index]) * float64(sec_vector[index]);
			}

			// get the magnitude of each vector
			magnitude1 := vectors_magnitude[first_user_key];
			magnitude2 := vectors_magnitude[sec_user_key];

			if magnitude1 == 0 || magnitude2 == 0 {
				cosineResult[first_user_key][sec_user_key] = 0.0; 
			} else {
				cosineResult[first_user_key][sec_user_key] = dotProduct / (magnitude1 * magnitude2); // calculate the cosine of the angle between 2 vectors
			}
		}
	}

	return cosineResult;
}