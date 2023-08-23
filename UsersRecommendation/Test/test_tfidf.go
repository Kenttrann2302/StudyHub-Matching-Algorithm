package test

// TEST REPRESENTS TEXT DATA AS VECTOR
import (
	"fmt"
	"StudyHub-Matching-Algorithm/UsersRecommendation/VectorsRepresentation"
)

// Assert Macros for test cases
func ASSERT_TRUE(T bool) bool {
	if !T {
		return false;
	}
	return true;
}

func ASSERT_FALSE(T bool) bool {
	if T {
		return false;
	}
	return true;
}

func get_status_str(status bool) string {
	if status {
		return "PASSED";
	} 
	return "FAILED";
}

// Define the test suite
type TF_IDF_Test struct {
	test_result  [4]bool
	test_description [4]string
}

func NewTF_IDF_Test() *TF_IDF_Test{
	return &TF_IDF_Test{
		test_result: [4]bool{false, false, false, false},
		test_description: [4]string {
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
		return "";
	} 
	return testSuite.test_description[test_nums-1];
}

// function to run all the tests -> void
func (testSuite *TF_IDF_Test) runAllTests() {
	testSuite.test_result[0] = testSuite.test1();
	testSuite.test_result[1] = testSuite.test2();
	testSuite.test_result[2] = testSuite.test3();
	testSuite.test_result[3] = testSuite.test4();
}

// function to print the report
func (testSuite *TF_IDF_Test) printReport() {

}

// Test1: One less than 10 words documentation is represented in corrected vector
func (testSuite *TF_IDF_Test) test1() bool {
	// Test set up
	documentation := [2]string {
		"I want to become a software developer at Google",
		"Google is a popular search engine",
	}

	// testing the function
	tf_map := vectorsrepresentation.Tokenization(documentation[:]);
	idf_map := vectorsrepresentation.CalculateIDF(documentation[:], tf_map);
	tf_idf_map := vectorsrepresentation.CalculateTF_IDF(documentation[:], tf_map, idf_map);
	
	// checking the frequency map
	expected_tf_map := []map[string]float64 {
		{
			"I" : 1, "want" : 1, "to" : 1, "become" : 1, "a" : 1, "software" : 1, "developer" : 1, "at" : 1, "Google" : 1,
		}, 
		{
			"Google" : 1, "is" : 1, "a" : 1, "popular" : 1, "search" : 1, "engine" : 1,
		},
	}

	for i, expectedTFMap := range expected_tf_map {
		calculatedTF := tf_map[i];
		for term, expectedValue := range expectedTFMap {
			ASSERT_TRUE(expectedValue == calculatedTF[term]);
		}
	}

	// checking the inversion document frequencies
	expected_idf_map := []map[string]float64 {
		{ "I" : 0.3, "want" : 0.5, "to" : 0.7, "become" : 0.9, "a" : 0.5, "software" : 0.8, "developer" : 0.8, "at" : 0.6, "Google" : 1.0, "is" : 0.4, "popular" : 0.6, "search" : 0.7, "engine" : 0.8 },
	}

	for i, expectedValue := range expected_idf_map {
		ASSERT_TRUE(expectedValue == idf_map[i]);
	}
}

// Test2: Multiple less than 10 words documentations are represented in corrected vector
func (testSuite *TF_IDF_Test) test2() bool {

}

// Test3: One more than 50 words documentation is represented in corrected vector
func (testSuite *TF_IDF_Test) test3() bool {

}

// Test4: Multiple more than 100 words documentations are represented in corrected vector
func (testSuite *TF_IDF_Test) test4() bool {

}

