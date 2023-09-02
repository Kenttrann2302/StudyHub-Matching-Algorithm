package CosineSimilarity

import (
	"math"
)

/* after we got the TF-IDF vector representations of all the text documents, we will start to
calculate the cosine similarity between each of them to find how similar
between each of them.
*/

// function to calculate the magnitude of each vector
func CalculateTFIDFMagnitude(list_of_vectors map[int][]float64, pre_calculated_magnitude map[int]float64) map[int]float64 {
	vector_magnitudes_map := make(map[int]float64)
	for user_key, vector := range list_of_vectors {
		// if the user's vector has been calculated before -> skip
		_, exists := pre_calculated_magnitude[user_key];
		if exists {
			vector_magnitudes_map[user_key] = pre_calculated_magnitude[user_key];
			continue;
		}
		sum := 0.0
		for _, value := range vector {
			sum += value * value
		}
		vector_magnitudes_map[user_key] = math.Sqrt(sum)
	}
	return vector_magnitudes_map
}

// function to calculate the cosine similarity between each vector in the given list of users
func CalculateTFIDFCosineSimilarity(list_of_vectors map[int][]float64, vectors_magnitudes map[int]float64, pre_calculated_cos_similarity [][]float64) [][]float64 {
	cosineResult := make([][]float64, len(list_of_vectors))
	for i := range list_of_vectors {
		cosineResult[i] = make([]float64, len(list_of_vectors)) // create a n x n matrix for cosine similarity between 2 users
	}

	// find the dot product between 2 vectors
	for row, vector := range list_of_vectors {
		dot_product := 0.0
		for col, sec_vector := range list_of_vectors {
			if row == col {
				continue // ignore the diagonals
			}
			if pre_calculated_cos_similarity[row][col] != math.MaxFloat64 {
				cosineResult[row][col] = pre_calculated_cos_similarity[row][col]; // if this relationship has been calculated -> skip
				continue;
			}
			for index := 0; index < len(vector); index++ {
				dot_product += vector[index] * sec_vector[index]
			}
			// get the magnitude of each vector
			magnitude1 := vectors_magnitudes[row]
			magnitude2 := vectors_magnitudes[col]

			if magnitude1 == 0 || magnitude2 == 0 {
				cosineResult[row][col] = 0.0
			} else {
				cosineResult[row][col] = dot_product / (magnitude1 * magnitude2)
			}
		}
	}

	return cosineResult;
}

/*
	Expected result for tf_idf cosine result: 
	[[0.1223, 0.5245, 0.1343, 0, 1], .....]
*/
