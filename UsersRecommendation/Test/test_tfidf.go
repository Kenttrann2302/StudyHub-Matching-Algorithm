package main

// TEST REPRESENTS TEXT DATA AS VECTOR
import (
	"StudyHub-Matching-Algorithm/UsersRecommendation/VectorsRepresentation"
	"fmt"
)

// Assert Macros for test cases
func ASSERT_TRUE(T bool) bool {
	if !T {
		return false
	}
	return true
}

func ASSERT_FALSE(T bool) bool {
	if T {
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
	test_result      [4]bool
	test_description [4]string
}

func NewTF_IDF_Test() *TF_IDF_Test {
	return &TF_IDF_Test{
		test_result: [4]bool{false, false, false, false},
		test_description: [4]string{
			"Test1: One or Two less than 10 words documentation is represented in corrected vector",
			"Test2: Multiple less than 10 words with punctuations and special character documentations are represented in corrected vector",
			"Test3: One or Two more than 50 words documentation is represented in corrected vector",
			"Test4: Multiple more than 100 words with punctuations and special character documentations are represented in corrected vector",
		},
	}
}

// function to get the test description -> string
func (testSuite *TF_IDF_Test) getTestDescription(test_nums int) string {
	if test_nums < 1 || test_nums > 4 {
		return ""
	}
	return testSuite.test_description[test_nums-1]
}

// function to run all the tests -> void
func (testSuite *TF_IDF_Test) runAllTests() {
	testSuite.test_result[0] = testSuite.test1()
	testSuite.test_result[1] = testSuite.test2()
	testSuite.test_result[2] = testSuite.test3()
	testSuite.test_result[3] = testSuite.test4()
}

// function to print the report
func (testSuite *TF_IDF_Test) printReport() {
	fmt.Println("----------------- TF_IDF TEST RESULT -----------------")
	for i := 0; i < 4; i++ {
		fmt.Print(testSuite.test_description[i] + "\n" + get_status_str(testSuite.test_result[i]) + "\n")
	}
}

// Test1: One less than 10 words documentation is represented in corrected vector
func (testSuite *TF_IDF_Test) test1() bool {
	// Test set up
	documentation := [2]string{
		"I want to become a software developer at Google",
		"Google is a popular search engine",
	}

	// testing the function
	tf_map := vectorsrepresentation.Tokenization(documentation[:])
	idf_map := vectorsrepresentation.CalculateIDF(documentation[:], tf_map)
	tf_idf_map := vectorsrepresentation.CalculateTF_IDF(documentation[:], tf_map, idf_map)

	// checking the frequency map
	expected_tf_map := []map[string]float64{
		{
			"I": 0.1111, "want": 0.1111, "to": 0.1111, "become": 0.1111, "a": 0.1111, "software": 0.1111, "developer": 0.1111, "at": 0.1111, "Google": 0.1111,
		},
		{
			"Google": 0.1667, "is": 0.1667, "a": 0.1667, "popular": 0.1667, "search": 0.1667, "engine": 0.1667,
		},
	}

	for i, expectedTFMap := range expected_tf_map {
		calculatedTF := tf_map[i]
		for term, expectedValue := range expectedTFMap {
			ASSERT_TRUE(expectedValue == calculatedTF[term])
		}
	}

	// checking the inversion document frequencies
	expected_idf_map := map[string]float64{
		"I": -0.6931, "want": -0.6931, "to": -0.6931, "become": -0.6931, "a": 0.0000, "software": -0.6931, "developer": -0.6931, "at": -0.6931, "Google": 0.0000, "is": -0.6931, "popular": -0.6931, "search": -0.6931, "engine": -0.6931,
	}

	for term := range expected_idf_map {
		ASSERT_TRUE(expected_idf_map[term] == idf_map[term])
	}

	// checking the term frequencies - inverse document frequencies values
	expected_tf_idf_map := []map[string]float64{
		{"I": -0.0770, "want": -0.0770, "to": -0.0770, "become": -0.0770, "a": 0.0000, "software": -0.6931, "developer": -0.6931, "at": -0.6931, "Google": 0.0000},
		{"Google": 0.0000, "is": -0.0770, "a": 0.0000, "popular": -0.0770, "search": -0.0770, "engine": -0.0770},
	}

	for i, expected_tf_idf_value := range expected_tf_idf_map {
		calculatedTF_IDF := tf_idf_map[i]
		for term, expectedValue := range expected_tf_idf_value {
			ASSERT_TRUE(expectedValue == calculatedTF_IDF[term])
		}
	}

	return true
}

// "Test2: Multiple less than 10 words with punctuations and special character documentations are represented in corrected vector"
func (testSuite *TF_IDF_Test) test2() bool {
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

	// checking the terms frequency map
	expected_tf_map := []map[string]float64{
		{"I": 0.0625, "love": 0.0625, "programming": 0.0625, "but": 0.0625, "coding": 0.125, "is": 0.0625, "fun": 0.0625},
		{"programming": 0.1111, "is": 0.0556, "fun": 0.1111, "really": 0.0556},
		{"I": 0.0833, "enjoy": 0.0833, "coding": 0.0833},
		{"Coding": 0.0833, "is": 0.0833, "interesting": 0.0833, "isn't": 0.0417, "it": 0.0417},
		{"I": 0.0714, "love": 0.0714, "coding": 0.0714},
	}

	for i, expectedTFMap := range expected_tf_map {
		calculatedTF := tf_map[i]
		for term, expectedValue := range expectedTFMap {
			ASSERT_TRUE(expectedValue == calculatedTF[term])
		}
	}

	// checking the inversion document frequencies
	expected_idf_map := map[string]float64{
		"I": 0.51082, "love": 0.51082, "programming": 0.51082, "but": 0.51082, "coding": 0.51082, "is": 0.51082, "fun": 0.51082, "really": 0.51082, "enjoy": 0.51082, "interesting": 0.51082, "it": 0.51082, "isn't": 0.51082,
	}

	for term := range expected_idf_map {
		ASSERT_TRUE(expected_idf_map[term] == idf_map[term])
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
			ASSERT_TRUE(expectedValue == calculatedTF_IDF[term])
		}
	}

	return true
}

// "Test3: One or Two more than 50 words documentation is represented in corrected vector"
func (testSuite *TF_IDF_Test) test3() bool {
	// Test set up
	documentation := [3]string{
		"I love programming. Programming is my passion. I enjoy coding projects. Coding empowers my creativity. I'm fascinated by algorithms. Debugging is essential for clean code.",
		"Coding is fun! I love coding projects. Solving coding challenges is satisfying. Debugging is essential for clean code. I find joy in coding. Debugging helps me learn.",
		"Programming languages offer endless possibilities. Coding empowers creativity and innovation. Software development is a dynamic field. I specialize in web development. Web technologies are constantly evolving.",
	}

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
			ASSERT_TRUE(expectedValue == calculatedTF[term])
		}
	}

	// checking the inversion document frequencies
	expected_idf_map := map[string]float64{
		"I": 0.405465, "love": 0.405465, "programming": 0.405465, "is": 0.405465, "my": 0.405465, "passion": 0.405465, "enjoy": 0.405465, "coding": 0.405465, "projects": 0.405465, "empowers": 0.405465, "creativity": 0.405465, "fascinated": 0.405465, "by": 0.405465, "algorithms": 0.405465, "debugging": 0.405465, "essential": 0.405465, "for": 0.405465, "clean": 0.405465, "code": 0.405465, "fun": 0.405465, "Solving": 0.405465, "challenges": 0.405465, "satisfying": 0.405465, "Debugging": 0.405465, "helps": 0.405465, "me": 0.405465, "learn": 0.405465, "Programming": 0.405465, "languages": 0.405465, "offer": 0.405465, "endless": 0.405465, "possibilities": 0.405465, "Software": 0.405465, "development": 0.405465, "a": 0.405465, "dynamic": 0.405465, "field": 0.405465, "specialize": 0.405465, "in": 0.405465, "web": 0.405465, "technologies": 0.405465, "are": 0.405465, "constantly": 0.405465, "evolving": 0.405465,
	}

	for term := range expected_idf_map {
		ASSERT_TRUE(expected_idf_map[term] == idf_map[term])
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
			ASSERT_TRUE(expectedValue == calculatedTF_IDF[term])
		}
	}

	return true
}

// "Test4: Multiple more than 100 words with punctuations and special character documentations are represented in corrected vector"
func (testSuite *TF_IDF_Test) test4() bool {
	return false
}

// ----------------- MAIN --------------------
func main() {
	testSuite := NewTF_IDF_Test()
	testSuite.runAllTests()
	testSuite.printReport()
}
