package vectorsrepresentation

// this file will use One-hot-encoding technique to convert the categorical data into vectors representation for cosine similarity technique
import (

)

// function to ONE-HOT-ENCODING 
func OneHotEncoding(categories []string, user_study_pref map[int][]string, pre_calculated map[int][]int) map[int][]int {
	encoded_result := make(map[int][]int);

	// convert categories into map to achieve O(1) look up time complexity
	map_categories := make(map[string]int)
	for indx, str := range categories {
		map_categories[str] = indx; // record each string to its corresponding index in the categories array
	}
	
	// go through each users options and check 1 if that user prefer the category
	for user_key, preferences := range user_study_pref {
		encoded_result[user_key] = make([]int, len(categories));
		if pre_calculated[user_key] != nil {
			encoded_result[user_key] = pre_calculated[user_key];
			continue; // skip if the user has been calculated
		}
		for pref_indx := range preferences {
			_, exists := map_categories[preferences[pref_indx]]; // check if the preference exists in the categories
			if exists {
				encoded_result[user_key][map_categories[preferences[pref_indx]]] = 1;
			}
		}
	} 

	return encoded_result;
}