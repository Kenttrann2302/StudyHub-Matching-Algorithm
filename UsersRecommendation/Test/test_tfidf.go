package main

// TEST REPRESENTS TEXT DATA AS VECTOR
import (
	"StudyHub-Matching-Algorithm/UsersRecommendation/VectorsRepresentation"
	"fmt"
	"math"
	"testing"
)

// Assert Macros for test cases
func ASSERT_TRUE(t *testing.T, condition bool) bool {
	if !condition {
		return false
	}
	return true
}

func ASSERT_FALSE(t *testing.T, condition bool) bool {
	if condition {
		return false
	}
	return true
}

func get_status_str(status bool) string {
	if status {
		return "PASSED"
	}
	return "FAILED"
}

// Define the test suite
type TF_IDF_Test struct {
	test_result      [5]bool
	test_description [5]string
}

func NewTF_IDF_Test() *TF_IDF_Test {
	return &TF_IDF_Test{
		test_result: [5]bool{false, false, false, false, false},
		test_description: [5]string{
			"Test1: One or Two less than 10 words documentation is represented in corrected vector",
			"Test2: Multiple less than 10 words with punctuations and special character documentations are represented in corrected vector",
			"Test3: One or Two more than 50 words documentation is represented in corrected vector",
			"Test4: Multiple more than 100 words with punctuations and special character documentations are represented in corrected vector",
			"Test5: TF-IDF Vectors represents the documentations correctly",
		},
	}
}

// function to get the test description -> string
func (testSuite *TF_IDF_Test) getTestDescription(test_nums int) string {
	if test_nums < 1 || test_nums > 5 {
		return ""
	}
	return testSuite.test_description[test_nums-1]
}

// function to run all the tests -> void
func (testSuite *TF_IDF_Test) runAllTests(t *testing.T) {
	testSuite.test_result[0] = testSuite.test1(t)
	testSuite.test_result[1] = testSuite.test2(t)
	testSuite.test_result[2] = testSuite.test3(t)
	testSuite.test_result[3] = testSuite.test4(t)
	testSuite.test_result[4] = testSuite.test5(t)
}

// function to print the report
func (testSuite *TF_IDF_Test) printReport() {
	fmt.Println("----------------- TF_IDF TEST RESULT -----------------")
	for i := 0; i < 5; i++ {
		fmt.Print(testSuite.test_description[i] + "\n" + get_status_str(testSuite.test_result[i]) + "\n")
	}
}

// function to customize comparison of 2 floats
func (testSuite *TF_IDF_Test) CompareFloats(a, b, tolerance float64) bool {
	diff := math.Abs(a - b)
	return diff <= tolerance
}

// Test1: One less than 10 words documentation is represented in corrected vector
func (testSuite *TF_IDF_Test) test1(t *testing.T) bool {
	// Test set up
	documentation := [2]string{
		"I want to become a software developer at Google",
		"Google is a popular search engine",
	}

	// Set the tolerance based on the number of decimal places you want to compare
	tolerance := 1e-3

	// testing the function
	tf_map := vectorsrepresentation.Tokenization(documentation[:])
	idf_map := vectorsrepresentation.CalculateIDF(documentation[:], tf_map)
	tf_idf_map := vectorsrepresentation.CalculateTF_IDF(documentation[:], tf_map, idf_map)

	// checking the frequency map
	expected_tf_map := []map[string]float64{
		{
			"I": float64(0.1111), "want": float64(0.1111), "to": float64(0.1111), "become": float64(0.1111), "a": float64(0.1111), "software": float64(0.1111), "developer": float64(0.1111), "at": float64(0.1111), "Google": float64(0.1111),
		},
		{
			"Google": 0.1667, "is": 0.1667, "a": 0.1667, "popular": 0.1667, "search": 0.1667, "engine": 0.1667,
		},
	}

	for i, expectedTFMap := range expected_tf_map {
		calculatedTF := tf_map[i]
		for term, expectedValue := range expectedTFMap {
			if !testSuite.CompareFloats(calculatedTF[term], expectedValue, tolerance) {
				return false
			}
		}
	}

	// checking the inversion document frequencies
	expected_idf_map := map[string]float64{
		"I": 0.6931, "want": 0.6931, "to": 0.6931, "become": 0.6931, "a": 0.0000, "software": 0.6931, "developer": 0.6931, "at": 0.6931, "Google": 0.0000, "is": 0.6931, "popular": 0.6931, "search": 0.6931, "engine": 0.6931,
	}

	for term := range expected_idf_map {
		if !testSuite.CompareFloats(expected_idf_map[term], idf_map[term], tolerance) {
			return false
		}
	}

	// checking the term frequencies - inverse document frequencies values
	expected_tf_idf_map := []map[string]float64{
		{"I": 0.0770, "want": 0.0770, "to": 0.0770, "become": 0.0770, "a": 0.0000, "software": 0.0770, "developer": 0.0770, "at": 0.0770, "Google": 0.0000},
		{"Google": 0.0000, "is": 0.1155, "a": 0.0000, "popular": 0.1155, "search": 0.1155, "engine": 0.1155},
	}

	for i, expected_tf_idf_value := range expected_tf_idf_map {
		calculatedTF_IDF := tf_idf_map[i]
		for term, expectedValue := range expected_tf_idf_value {
			if !testSuite.CompareFloats(expectedValue, calculatedTF_IDF[term], tolerance) {
				return false
			}
		}
	}

	return true
}

// "Test2: Multiple less than 10 words with punctuations and special character documentations are represented in corrected vector"
func (testSuite *TF_IDF_Test) test2(t *testing.T) bool {
	// Test set up
	documentation := [5]string{
		"I love programming, but coding is fun!",
		"Programming is fun! Really fun!",
		"I enjoy coding.",
		"Coding is interesting, isn't it?",
		"I love coding!",
	}

	// testing the function
	tf_map := vectorsrepresentation.Tokenization(documentation[:])
	idf_map := vectorsrepresentation.CalculateIDF(documentation[:], tf_map)
	tf_idf_map := vectorsrepresentation.CalculateTF_IDF(documentation[:], tf_map, idf_map)

	// Set the tolerance based on the number of decimal places you want to compare
	tolerance := 1e-3

	// checking the terms frequency map
	expected_tf_map := []map[string]float64{
		{"I": 0.1429, "love": 0.1429, "programming": 0.1429, "but": 0.1429, "coding": 0.1429, "is": 0.1429, "fun": 0.1429},
		{"programming": 0.1111, "is": 0.0556, "fun": 0.1111, "really": 0.0556},
		{"I": 0.0833, "enjoy": 0.0833, "coding": 0.0833},
		{"Coding": 0.0833, "is": 0.0833, "interesting": 0.0833, "isn't": 0.0417, "it": 0.0417},
		{"I": 0.0714, "love": 0.0714, "coding": 0.0714},
	}

	for i, expectedTFMap := range expected_tf_map {
		calculatedTF := tf_map[i]
		for term, expectedValue := range expectedTFMap {
			if !testSuite.CompareFloats(expectedValue, calculatedTF[term], tolerance) {
				return false
			}
		}
	}

	// checking the inversion document frequencies
	expected_idf_map := map[string]float64{
		"I": 0.51082, "love": 0.51082, "programming": 0.51082, "but": 0.51082, "coding": 0.51082, "is": 0.51082, "fun": 0.51082, "really": 0.51082, "enjoy": 0.51082, "interesting": 0.51082, "it": 0.51082, "isn't": 0.51082,
	}

	for term := range expected_idf_map {
		if !testSuite.CompareFloats(expected_idf_map[term], idf_map[term], tolerance) {
			return false
		}
	}

	// checking the term frequencies - inverse document frequencies values
	expected_tf_idf_map := []map[string]float64{
		{"I": 0.0319, "love": 0.0319, "programming": 0.0319, "but": 0.0319, "coding": 0.0643, "is": 0.0319, "fun": 0.0319},
		{"programming": 0.0565, "is": 0.0282, "fun": 0.0565, "really": 0.0282},
		{"I": 0.0429, "enjoy": 0.0429, "coding": 0.0429},
		{"Coding": 0.0429, "is": 0.0429, "interesting": 0.0429, "isn't": 0.0215, "it": 0.0215},
		{"I": 0.0365, "love": 0.0365, "coding": 0.0365},
	}

	for i, expected_tf_idf_value := range expected_tf_idf_map {
		calculatedTF_IDF := tf_idf_map[i]
		for term, expectedValue := range expected_tf_idf_value {
			if !testSuite.CompareFloats(calculatedTF_IDF[term], expectedValue, tolerance) {
				return false
			}
		}
	}

	return true
}

// "Test3: One or Two more than 50 words documentation is represented in corrected vector"
func (testSuite *TF_IDF_Test) test3(t *testing.T) bool {
	// Test set up
	documentation := [3]string{
		"I love programming. Programming is my passion. I enjoy coding projects. Coding empowers my creativity. I'm fascinated by algorithms. Debugging is essential for clean code.",
		"Coding is fun! I love coding projects. Solving coding challenges is satisfying. Debugging is essential for clean code. I find joy in coding. Debugging helps me learn.",
		"Programming languages offer endless possibilities. Coding empowers creativity and innovation. Software development is a dynamic field. I specialize in web development. Web technologies are constantly evolving.",
	}

	tolerance := 1e-3

	// testing the function
	tf_map := vectorsrepresentation.Tokenization(documentation[:])
	idf_map := vectorsrepresentation.CalculateIDF(documentation[:], tf_map)
	tf_idf_map := vectorsrepresentation.CalculateTF_IDF(documentation[:], tf_map, idf_map)

	// checking the terms frequency map
	expected_tf_map := []map[string]float64{
		{
			"I": 0.0435, "love": 0.0217, "programming": 0.0217, "is": 0.0217, "my": 0.0217, "passion": 0.0217, "enjoy": 0.0217, "coding": 0.0217, "projects": 0.0217, "empowers": 0.0217, "creativity": 0.0217, "fascinated": 0.0217, "by": 0.0217, "algorithms": 0.0217, "debugging": 0.0217, "essential": 0.0217, "for": 0.0217, "clean": 0.0217, "code": 0.0217,
		},
		{
			"Coding": 0.0455, "is": 0.0227, "fun": 0.0227, "I": 0.0227, "love": 0.0227, "coding": 0.0227, "projects": 0.0227, "Solving": 0.0227, "challenges": 0.0227, "satisfying": 0.0227, "Debugging": 0.0227, "essential": 0.0227, "for": 0.0227, "clean": 0.0227, "code": 0.0227, "find": 0.0227, "joy": 0.0227, "in": 0.0227, "helps": 0.0227, "me": 0.0227, "learn": 0.0227,
		},
		{
			"Programming": 0.0286, "languages": 0.0143, "offer": 0.0143, "endless": 0.0143, "possibilities": 0.0143, "Coding": 0.0143, "empowers": 0.0143, "creativity": 0.0143, "and": 0.0143, "innovation": 0.0143, "Software": 0.0143, "development": 0.0143, "is": 0.0143, "a": 0.0143, "dynamic": 0.0143, "field": 0.0143, "I": 0.0143, "specialize": 0.0143, "in": 0.0143, "web": 0.0143, "technologies": 0.0143, "are": 0.0143, "constantly": 0.0143, "evolving": 0.0143,
		},
	}

	for i, expectedTFMap := range expected_tf_map {
		calculatedTF := tf_map[i]
		for term, expectedValue := range expectedTFMap {
			if !testSuite.CompareFloats(expectedValue, calculatedTF[term], tolerance) {
				return false
			}
		}
	}

	// checking the inversion document frequencies
	expected_idf_map := map[string]float64{
		"I": 0.405465, "love": 0.405465, "programming": 0.405465, "is": 0.405465, "my": 0.405465, "passion": 0.405465, "enjoy": 0.405465, "coding": 0.405465, "projects": 0.405465, "empowers": 0.405465, "creativity": 0.405465, "fascinated": 0.405465, "by": 0.405465, "algorithms": 0.405465, "debugging": 0.405465, "essential": 0.405465, "for": 0.405465, "clean": 0.405465, "code": 0.405465, "fun": 0.405465, "Solving": 0.405465, "challenges": 0.405465, "satisfying": 0.405465, "Debugging": 0.405465, "helps": 0.405465, "me": 0.405465, "learn": 0.405465, "Programming": 0.405465, "languages": 0.405465, "offer": 0.405465, "endless": 0.405465, "possibilities": 0.405465, "Software": 0.405465, "development": 0.405465, "a": 0.405465, "dynamic": 0.405465, "field": 0.405465, "specialize": 0.405465, "in": 0.405465, "web": 0.405465, "technologies": 0.405465, "are": 0.405465, "constantly": 0.405465, "evolving": 0.405465,
	}

	for term := range expected_idf_map {
		if !testSuite.CompareFloats(expected_idf_map[term], idf_map[term], tolerance) {
			return false
		}
	}

	// checking the term frequencies - inverse document frequencies values
	expected_tf_idf_map := []map[string]float64{
		{
			"I": 0.0177, "love": 0.0088, "programming": 0.0088, "is": 0.0088, "my": 0.0088, "passion": 0.0088, "enjoy": 0.0088, "coding": 0.0088, "projects": 0.0088, "empowers": 0.0088, "creativity": 0.0088, "fascinated": 0.0088, "by": 0.0088, "algorithms": 0.0088, "debugging": 0.0088, "essential": 0.0088, "for": 0.0088, "clean": 0.0088, "code": 0.0088,
		},
		{
			"Coding": 0.0185, "is": 0.0092, "fun": 0.0092, "I": 0.0092, "love": 0.0092, "coding": 0.0092, "projects": 0.0092, "Solving": 0.0092, "challenges": 0.0092, "satisfying": 0.0092, "Debugging": 0.0092, "essential": 0.0092, "for": 0.0092, "clean": 0.0092, "code": 0.0092, "find": 0.0092, "joy": 0.0092, "in": 0.0092, "helps": 0.0092, "me": 0.0092, "learn": 0.0092,
		},
		{
			"Programming": 0.0116, "languages": 0.0058, "offer": 0.0058, "endless": 0.0058, "possibilities": 0.0058, "Coding": 0.0058, "empowers": 0.0058, "creativity": 0.0058, "and": 0.0058, "innovation": 0.0058, "Software": 0.0058, "development": 0.0058, "is": 0.0058, "a": 0.0058, "dynamic": 0.0058, "field": 0.0058, "I": 0.0058, "specialize": 0.0058, "in": 0.0058, "web": 0.0058, "technologies": 0.0058, "are": 0.0058, "constantly": 0.0058, "evolving": 0.0058,
		},
	}

	for i, expected_tf_idf_value := range expected_tf_idf_map {
		calculatedTF_IDF := tf_idf_map[i]
		for term, expectedValue := range expected_tf_idf_value {
			if !testSuite.CompareFloats(expectedValue, calculatedTF_IDF[term], tolerance) {
				return false
			}
		}
	}

	return true
}

// "Test4: Multiple more than 100 words with punctuations and special character documentations are represented in corrected vector"
func (testSuite *TF_IDF_Test) test4(t *testing.T) bool {
	documentation := [8]string{
		"Exploring distant galaxies is a captivating endeavor. Telescopes reveal cosmic wonders. Astronomers decode the language of stars and galaxies.",
		"In the heart of bustling cities, life thrives. Skyscrapers pierce the sky, casting long shadows. Urban jungles merge nature and concrete in harmony.",
		"History's pages are filled with tales of courage. Warriors face battles with unwavering resolve. Their stories echo through time, inspiring generations.",
		"Across the rolling landscapes, farmers toil. Crops sway with the wind, a testament to hard work. Agriculture sustains communities and connects us to the earth.",
		"From the depths of oceans to mountaintops, life flourishes. Ecosystems interweave species in delicate balance. Biodiversity ensures our planet's resilience.",
		"Artisans craft masterpieces from raw materials. Hands mold clay, chisels shape stone. Creativity flows, giving life to sculptures, paintings, and timeless beauty.",
		"Words have power to shape thoughts and spark revolutions. Writers weave narratives that inspire change. Literature's impact is felt across cultures and time.",
		"Exploring the human mind's complexities, psychologists delve deep. Emotions, thoughts, and behaviors intertwine. Understanding ourselves brings clarity and growth.",
	}

	tolerance := 1e-3

	// testing the function
	tf_map := vectorsrepresentation.Tokenization(documentation[:])
	idf_map := vectorsrepresentation.CalculateIDF(documentation[:], tf_map)
	tf_idf_map := vectorsrepresentation.CalculateTF_IDF(documentation[:], tf_map, idf_map)

	// checking the terms frequency map
	expected_tf_map := []map[string]float64{
		{
			"Exploring": 0.0588, "distant": 0.0588, "galaxies": 0.1176, "is": 0.0588, "a": 0.0588, "captivating": 0.0588, "endeavor": 0.0588, "Telescopes": 0.0588, "reveal": 0.0588, "cosmic": 0.0588, "wonders": 0.0588, "Astronomers": 0.0588, "decode": 0.0588, "the": 0.1176, "language": 0.0588, "of": 0.0588, "stars": 0.0588, "and": 0.0588,
		},
		{
			"In": 0.0714, "the": 0.0714, "heart": 0.0714, "of": 0.0714, "bustling": 0.0714, "cities": 0.0714, "life": 0.0714, "thrives": 0.0714, "Skyscrapers": 0.0714, "pierce": 0.0714, "sky": 0.0714, "casting": 0.0714, "long": 0.0714, "shadows": 0.0714, "Urban": 0.0714, "jungles": 0.0714, "merge": 0.0714, "nature": 0.0714, "and": 0.0714, "concrete": 0.0714, "in": 0.0714, "harmony": 0.0714,
		},
		{
			"History's": 0.0526, "pages": 0.0526, "are": 0.0526, "filled": 0.0526, "with": 0.0526, "tales": 0.0526, "of": 0.0526, "courage": 0.0526, "Warriors": 0.0526, "face": 0.0526, "battles": 0.0526, "unwavering": 0.0526, "resolve": 0.0526, "Their": 0.0526, "stories": 0.0526, "echo": 0.0526, "through": 0.0526, "time": 0.0526, "inspiring": 0.0526, "generations": 0.0526,
		},
		{
			"Across": 0.0714, "the": 0.0714, "rolling": 0.0714, "landscapes": 0.0714, "farmers": 0.0714, "toil": 0.0714, "Crops": 0.0714, "sway": 0.0714, "with": 0.0714, "wind": 0.0714, "a": 0.0714, "testament": 0.0714, "to": 0.0714, "hard": 0.0714, "work": 0.0714, "Agriculture": 0.0714, "sustains": 0.0714, "communities": 0.0714, "connects": 0.0714, "us": 0.0714, "earth": 0.0714,
		},
		{
			"From": 0.0556, "the": 0.0556, "depths": 0.0556, "of": 0.0556, "oceans": 0.0556, "to": 0.0556, "mountaintops": 0.0556, "life": 0.0556, "flourishes": 0.0556, "Ecosystems": 0.0556, "interweave": 0.0556, "species": 0.0556, "in": 0.0556, "delicate": 0.0556, "balance": 0.0556, "Biodiversity": 0.0556, "ensures": 0.0556, "our": 0.0556, "planet's": 0.0556, "resilience": 0.0556,
		},
		{
			"Artisans": 0.0625, "craft": 0.0625, "masterpieces": 0.0625, "from": 0.0625, "raw": 0.0625, "materials": 0.0625, "Hands": 0.0625, "mold": 0.0625, "clay": 0.0625, "chisels": 0.0625, "shape": 0.0625, "stone": 0.0625, "Creativity": 0.0625, "flows": 0.0625, "giving": 0.0625, "life": 0.0625, "to": 0.0625, "sculptures": 0.0625, "paintings": 0.0625, "and": 0.0625, "timeless": 0.0625, "beauty": 0.0625,
		},
		{
			"Words": 0.0556, "have": 0.0556, "power": 0.0556, "to": 0.0556, "shape": 0.0556, "thoughts": 0.0556, "and": 0.1111, "spark": 0.0556, "revolutions": 0.0556, "Writers": 0.0556, "weave": 0.0556, "narratives": 0.0556, "that": 0,
		},
		{
			"Exploring": 0.0667, "the": 0.0667, "human": 0.0667, "mind's": 0.0667, "complexities": 0.0667, "psychologists": 0.0667, "delve": 0.0667, "deep": 0.0667, "Emotions": 0.0667, "thoughts": 0.0667, "and": 0.0667, "behaviors": 0.0667, "intertwine": 0.0667, "Understanding": 0.0667, "ourselves": 0.0667, "brings": 0.0667, "clarity": 0.0667, "growth": 0.0667,
		},
	}

	for i, expectedTFMap := range expected_tf_map {
		calculatedTF := tf_map[i]
		for term, expectedValue := range expectedTFMap {
			if !testSuite.CompareFloats(expectedValue, calculatedTF[term], tolerance) {
				return false
			}
			return true
		}
	}

	// checking the inversion document frequencies
	expected_idf_map := map[string]float64{
		"Exploring": 1.94591, "distant": 1.94591, "galaxies": 1.94591, "is": 1.94591, "a": 0.405465, "captivating": 1.94591, "endeavor": 1.94591, "Telescopes": 1.94591, "reveal": 1.94591, "cosmic": 1.94591, "wonders": 1.94591, "Astronomers": 1.94591, "decode": 1.94591, "the": 0.405465, "language": 1.94591, "of": 0.405465, "stars": 1.94591, "and": 0.405465, "In": 1.94591, "heart": 1.94591, "bustling": 1.94591, "cities": 1.94591, "life": 0.405465, "thrives": 1.94591, "Skyscrapers": 1.94591, "pierce": 1.94591, "sky": 1.94591, "casting": 1.94591, "long": 1.94591, "shadows": 1.94591, "Urban": 1.94591, "jungles": 1.94591, "merge": 1.94591, "nature": 1.94591, "concrete": 1.94591, "harmony": 1.94591, "History's": 1.94591, "pages": 1.94591, "are": 0.405465, "filled": 1.94591, "tales": 1.94591, "courage": 1.94591, "Warriors": 1.94591, "face": 1.94591, "battles": 1.94591, "unwavering": 1.94591, "resolve": 1.94591, "Their": 1.94591, "stories": 1.94591, "echo": 1.94591, "through": 1.94591, "time": 1.94591, "inspiring": 1.94591, "generations": 1.94591, "Across": 1.94591, "rolling": 1.94591, "landscapes": 1.94591, "farmers": 1.94591, "toil": 1.94591, "Crops": 1.94591, "sway": 1.94591, "wind": 1.94591, "testament": 1.94591, "hard": 1.94591, "work": 1.94591, "Agriculture": 1.94591, "sustains": 1.94591, "communities": 1.94591, "connects": 1.94591, "us": 1.94591, "earth": 1.94591, "From": 1.94591, "depths": 1.94591, "oceans": 1.94591, "mountaintops": 1.94591, "flourishes": 1.94591, "Ecosystems": 1.94591, "interweave": 1.94591, "species": 1.94591, "delicate": 1.94591, "balance": 1.94591, "Biodiversity": 1.94591, "ensures": 1.94591, "our": 1.94591, "planet's": 1.94591, "resilience": 1.94591, "Artisans": 1.94591, "craft": 1.94591, "masterpieces": 1.94591, "from": 1.94591, "raw": 1.94591, "materials": 1.94591, "Hands": 1.94591, "mold": 1.94591, "clay": 1.94591, "chisels": 1.94591, "stone": 1.94591, "Creativity": 1.94591, "flows": 1.94591, "giving": 1.94591, "sculptures": 1.94591, "paintings": 1.94591, "timeless": 1.94591, "beauty": 1.94591, "Words": 1.94591, "have": 1.94591, "power": 1.94591, "thoughts": 1.94591, "spark": 1.94591, "revolutions": 1.94591, "Writers": 1.94591, "weave": 1.94591, "narratives": 1.94591, "that": 1.94591, "inspire": 1.94591, "change": 1.94591, "Literature's": 1.94591, "impact": 1.94591, "felt": 1.94591, "across": 1.94591, "cultures": 1.94591, "human": 1.94591, "mind's": 1.94591, "complexities": 1.94591, "psychologists": 1.94591, "delve": 1.94591, "deep": 1.94591, "Emotions": 1.94591, "behaviors": 1.94591, "intertwine": 1.94591, "Understanding": 1.94591, "ourselves": 1.94591, "brings": 1.94591, "clarity": 1.94591, "growth": 1.94591,
	}

	for term := range expected_idf_map {
		if !testSuite.CompareFloats(expected_idf_map[term], idf_map[term], tolerance) {
			return false
		}
	}

	// checking the term frequencies - inverse document frequencies values
	expected_tf_idf_map := []map[string]float64{
		{"Across": 0.000779819, "distant": 0.000779819, "galaxies": 0.000779819, "revealing": 0.000779819, "cosmic": 0.000779819, "wonders": 0.000779819, "Telescopes": 0.000779819, "endeavor": 0.000779819, "captivating": 0.000779819},
		{"In": 0.000779819, "the": 0.000778225, "heart": 0.000779819, "bustling": 0.000779819, "cities": 0.000779819, "life": 0.000778225, "thrives": 0.000779819, "Skyscrapers": 0.000779819, "pierce": 0.000779819, "sky": 0.000779819, "casting": 0.000779819, "long": 0.000779819, "shadows": 0.000779819, "Urban": 0.000779819, "jungles": 0.000779819, "merge": 0.000779819, "nature": 0.000779819, "concrete": 0.000779819, "harmony": 0.000779819},
		{"History's": 0.000779819, "pages": 0.000779819, "filled": 0.000779819, "tales": 0.000779819, "courage": 0.000779819, "Warriors": 0.000779819, "face": 0.000779819, "battles": 0.000779819, "unwavering": 0.000779819, "resolve": 0.000779819, "Their": 0.000779819, "stories": 0.000779819, "echo": 0.000779819, "through": 0.000779819, "time": 0.000779819, "inspiring": 0.000779819, "generations": 0.000779819},
		{"Across": 0.000779819, "rolling": 0.000779819, "landscapes": 0.000779819, "farmers": 0.000779819, "toil": 0.000779819, "Crops": 0.000779819, "sway": 0.000779819, "wind": 0.000779819, "testament": 0.000779819, "hard": 0.000779819, "work": 0.000779819, "Agriculture": 0.000779819, "sustains": 0.000779819, "communities": 0.000779819, "connects": 0.000779819, "us": 0.000779819, "earth": 0.000779819},
		{"From": 0.000779819, "depths": 0.000779819, "oceans": 0.000779819, "mountaintops": 0.000779819, "flourishes": 0.000779819, "Ecosystems": 0.000779819, "interweave": 0.000779819, "species": 0.000779819, "delicate": 0.000779819, "balance": 0.000779819, "Biodiversity": 0.000779819, "ensures": 0.000779819, "our": 0.000779819, "planet's": 0.000779819, "resilience": 0.000779819},
		{"Artisans": 0.000779819, "craft": 0.000779819, "masterpieces": 0.000779819, "raw": 0.000779819, "materials": 0.000779819, "Hands": 0.000779819, "mold": 0.000779819, "clay": 0.000779819, "chisels": 0.000779819, "shape": 0.000779819, "stone": 0.000779819, "Creativity": 0.000779819, "flows": 0.000779819, "giving": 0.000779819, "sculptures": 0.000779819, "paintings": 0.000779819, "timeless": 0.000779819, "beauty": 0.000779819},
		{"Words": 0.000779819, "have": 0.000779819, "power": 0.000779819, "shape": 0.000779819, "thoughts": 0.000779819, "spark": 0.000779819, "revolutions": 0.000779819, "Writers": 0.000779819, "weave": 0.000779819, "narratives": 0.000779819, "that": 0.000779819, "inspire": 0.000779819, "change": 0.000779819, "Literature's": 0.000779819, "impact": 0.000779819, "felt": 0.000779819, "across": 0.000779819, "cultures": 0.000779819},
		{"Exploring": 0.000779819, "human": 0.000779819, "mind's": 0.000779819, "complexities": 0.000779819, "psychologists": 0.000779819, "delve": 0.000779819, "deep": 0.000779819, "Emotions": 0.000779819, "behaviors": 0.000779819, "intertwine": 0.000779819, "Understanding": 0.000779819, "ourselves": 0.000779819, "brings": 0.000779819, "clarity": 0.000779819, "growth": 0.000779819},
	}

	for i, expected_tf_idf_value := range expected_tf_idf_map {
		calculatedTF_IDF := tf_idf_map[i]
		for term, expectedValue := range expected_tf_idf_value {
			if !testSuite.CompareFloats(expectedValue, calculatedTF_IDF[term], tolerance) {
				return false
			}
		}
	}

	return true
}

// "Test5: TF-IDF Vectors represents the documentations correctly"
func (testSuite *TF_IDF_Test) test5(t *testing.T) bool {
	// test set up
	documentation := [5]string{
		"I enjoy playing the guitar and making dinner for my family and friends.",
		"Music is my passion, and I love listening to various genres. I also like playing soccer and basketball during my leisure.",
		"I'm passionate about coding and learning new programming languages. I also like listening to Taylor Swift songs!",
		"Cooking is my hobby, and I enjoy trying new recipes. I also likes playing a variety of music instruments including guitar.",
		"I usually spend most of my day training for my basketball team. I also like to cook my own meal.",
	}

	tolerance := 1e-3

	// get the result of the function
	tf_map := vectorsrepresentation.Tokenization(documentation[:])
	idf_map := vectorsrepresentation.CalculateIDF(documentation[:], tf_map)
	tf_idf_map := vectorsrepresentation.CalculateTF_IDF(documentation[:], tf_map, idf_map)
	tf_idf_vectors := vectorsrepresentation.CalculateTFIDFVectors(tf_idf_map, documentation[:])

	// test the tf-idf vectors representation for each document
	expected_tf_idf_vectors := [][]float64{
		{0.5625, 0.5625, 1.5625, 0.125, 1.5625, 0.125, 1.5625, 1.5625, 0.125, 0.5625, 1.5625, 1.5625},
		{0.5625, 0.5625, 1.5625, 0.125, 0, 0, 0.5625, 0, 0, 0.5625, 0, 0},
		{0.5625, 0, 0, 0, 0, 0, 0.5625, 0, 0.5625, 0, 0, 0},
		{0, 0.5625, 0, 0, 0, 0, 0.5625, 0.5625, 0, 0.5625, 0, 0},
		{0.5625, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0.5625},
	}

	for i, vector := range expected_tf_idf_vectors {
		calculated_tf_idf_vector := tf_idf_vectors[i]
		for j, value := range vector {
			calc_value := calculated_tf_idf_vector[j]
			if !testSuite.CompareFloats(calc_value, value, tolerance) {
				return false
			}
		}
	}

	return true
}

// ----------------- MAIN --------------------
func main() {
	testSuite := NewTF_IDF_Test()
	testSuite.runAllTests(&testing.T{})
	testSuite.printReport()
}
