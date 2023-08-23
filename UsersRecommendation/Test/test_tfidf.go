package test

// TEST REPRESENTS TEXT DATA AS VECTOR
import (
	"fmt"
	b ""
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
			"Test1: One less than 10 words documentation is represented in corrected vector",
			"Test2: Multiple less than 10 words documentations are represented in corrected vector",
			"Test3: One more than 50 words documentation is represented in corrected vector",
			"Test4: Multiple more than 100 words documentations are represented in corrected vector",
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
	documentation := [1]string {
		"I want to become a software developer at Google",
	}

	// testing the function
	tf_idf_map := calculateTF_IDF()
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

