package vectorsrepresentation

// this file will convert date and time data into vectors representation

// a function to convert the availability date and time into vectors 
/*
	availability_date_time : {
		1 : {
			"Monday" : [2, 5, 6, 9], // this will be on a 24 hours scale: from 0 to 23
			"Tuesday" : [3, 4, 6, 8],
			...
		}
		...
	}
*/
func CalculateDateTimeVectors(availability_date_time map[int]map[string][]int) map[int][][]int {
	result := make(map[int][][]int);
	for key := range result {
		result[key] = make([][]int, 7);
		for value := range result[key] {
			result[key][value] = make([]int, 24); // the map with matrices of 7x24
		}
	}

	// calculate the date time vector for each user
	for user, available_map := range availability_date_time {
		date_index := 0; // this is the index for each day in the matrix (Monday will be 0, Tuesday will be 1, Wednesday will be 2,....)
		for _, availability := range available_map {
			for available_time := range availability {
				result[user][date_index][available_time] = 1;
			}
			date_index++;
		}
	}

	return result;
}

/* expected result: {
	12 : [[1, 0, 0, 0, 1, 1, 1], [0, 1, 1, 0, 1], ...],
	...
}
*/