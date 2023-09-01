package CosineSimilarity

// this file will calculate the cosine similarity for the availability date time between each user
import (
	"math"
)

// function to calculate the magnitude of each date time vector
func CalculateDateTimeMagnitude(date_time_vectors map[int][][]int) map[int][]float64 {
	magnitude_result := make(map[int][]float64);
	for key := range magnitude_result {
		magnitude_result[key] = make([]float64, 7);
	} 
	
	for user_key, date_vetor := range date_time_vectors {
		for date_index, vector := range date_vetor {
			magnitude := 0.0;
			for _, value := range vector {
				magnitude += float64(value * value);
			}
			magnitude_result[user_key][date_index] = math.Sqrt(magnitude);
		}
	}

	return magnitude_result;
}

/*
	Expected result for date time magnitude
	{
		1 : [1, 0, 1, ....],
		... 
	}
*/

// function to calculate the cosine similarity between each corresponding vector between each user
func CalculateDateTimeCosineSimilarity(date_time_vectors map[int][][]int, vectors_magnitude map[int][]float64) [][]float64 { // the result will be the cosine similarity after comparing all the corresponding date between 2 users
	cosineResult := make([][]float64, len(date_time_vectors));
	for i := range date_time_vectors {
		cosineResult[i] = make([]float64, len(date_time_vectors));
		// create an nxn matrix represents the cosine similarity between 2 users
	}

	// find the cosine similarities between 2 vectors in the list of vectors
	for first_user_key, first_list_vectors := range date_time_vectors {
		dot_product := 0.0;
		for sec_user_key, second_list_vectors := range date_time_vectors {
			if first_user_key == sec_user_key {
				continue; // ignore the diagonals
			}
			temp_sum_cosine_result := 0.0;
			for vector := range first_list_vectors {
				// find the dot product between 2 corresponding date vectors
				for indx := 0; indx < len(first_list_vectors[vector]); indx++ {
					dot_product += float64(first_list_vectors[vector][indx] * second_list_vectors[vector][indx]);
				}
				// get the magnitude of each vector
				magnitude_1 := vectors_magnitude[first_user_key][vector];
				magnitude_2 := vectors_magnitude[sec_user_key][vector];
				
				// find the cosine similarity between 2 vectors
				if magnitude_1 == 0 || magnitude_2 == 0 {
					temp_sum_cosine_result += 0.0;
				} else {
					temp_sum_cosine_result += (dot_product/(magnitude_1 * magnitude_2)); // calculate the sum of all the cosine similarity between 2 users between each corresponding dates
				}
			}
			// the date_time cosine similarity between 2 users will be the average cosine similarity (sum cosine result for 7 days divided by 7 days) between them
			average_cosine_similarity := temp_sum_cosine_result / float64(7.0);
			cosineResult[first_user_key][sec_user_key] = average_cosine_similarity;
		}
	}
	return cosineResult;
}

/*
	Expected result for date_time cosine result:
	[[0.2123, 0.7883, 0.321, ....]]
*/