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

// function to calculate the cosine similarity between each vector
func calculateCosineSimilarity(list_of_vectors map[int]float64) {

}
