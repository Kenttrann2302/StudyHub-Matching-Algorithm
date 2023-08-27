package CosineSimilarity

import "math"

/* after we got the TF-IDF vector representations of all the text documents, we will start to
calculate the cosine similarity between each of them to find how similar
between each of them.
*/

// function to calculate the magnitude of each vector
func calculateMagnitude(list_of_vectors map[int][]float64) map[int]float64 {
	vector_magnitudes_map := make(map[int]float64)
	for user_key, vector := range list_of_vectors {
		sum := 0.0
		for _, value := range vector {
			sum += value * value
		}
		vector_magnitudes_map[user_key] = math.Sqrt(sum)
	}
	return vector_magnitudes_map
}

// function to calculate the cosine similarity between each vector in the given list of users
func calculateCosineSimilarity(list_of_vectors map[int][]float64, vectors_magnitudes map[int]float64) [][]float64{
	cosineResult := make([][]float64, len(list_of_vectors));
	for i := range list_of_vectors {
		cosineResult[i] = make([]float64, len(list_of_vectors)); // create a n x n matrix for cosine similarity between 2 users
	}

	// find the dot product between 2 vectors
	dot_product := 0.0
	for row, vector := range list_of_vectors {
		for col, sec_vector := range list_of_vectors {
			if row == col {
				continue; // ignore the diagonals
			}
			for index := 0; index < len(vector); index++ {
				dot_product += vector[index] * sec_vector[index];
			}		
			// get the magnitude of each vector
			magnitude1 := vectors_magnitudes[row];
			magnitude2 := vectors_magnitudes[col];

			if magnitude1 == 0 || magnitude2 == 0 {
				cosineResult[row][col] = 0.0; 
			} else  {
				cosineResult[row][col] = dot_product / (magnitude1 * magnitude2);
			}
		}
	}

	return cosineResult
}

// function to 