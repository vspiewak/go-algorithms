package search

import (
	"testing"
)

type searchTest struct {
	values        []int
	key           int
	expected      int
	expectedError error
}

var searchTests = []searchTest{
	{[]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, 10, 9, nil},
	{[]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, 9, 8, nil},
	{[]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, 8, 7, nil},
	{[]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, 7, 6, nil},
	{[]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, 6, 5, nil},
	{[]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, 5, 4, nil},
	{[]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, 4, 3, nil},
	{[]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, 3, 2, nil},
	{[]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, 2, 1, nil},
	{[]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, 1, 0, nil},
	{[]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, -20, -1, ErrNotFound},
	{[]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, 20, -1, ErrNotFound},
}

func generateBenchmarkTestCase() []int {
	var testCase []int
	for i := 0; i < 10000; i++ {
		testCase = append(testCase, i)
	}
	return testCase
}

func TestBinaryRec(t *testing.T) {

	for _, test := range searchTests {
		actualValue, actualError := BinaryRec(test.values, test.key)
		if actualValue != test.expected {
			t.Errorf("test failed, searching '%d' on values '%v', expected '%d', actual '%d'", test.key, test.values, test.expected, actualValue)
		}
		if actualError != test.expectedError {
			t.Errorf("test failed, searching '%d' on values '%v', expected error '%s', actual '%s'", test.key, test.values, test.expectedError, actualError)
		}
	}
}

func TestBinaryIter(t *testing.T) {

	for _, test := range searchTests {
		actualValue, actualError := BinaryIter(test.values, test.key)
		if actualValue != test.expected {
			t.Errorf("test failed, searching '%d' on values '%v', expected '%d', actual '%d'", test.key, test.values, test.expected, actualValue)
		}
		if actualError != test.expectedError {
			t.Errorf("test failed, searching '%d' on values '%v', expected error '%s', actual '%s'", test.key, test.values, test.expectedError, actualError)
		}
	}
}

func BenchmarkBinaryRec(b *testing.B) {
	testCase := generateBenchmarkTestCase()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = BinaryRec(testCase, i)
	}
}

func BenchmarkBinaryIter(b *testing.B) {
	testCase := generateBenchmarkTestCase()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = BinaryIter(testCase, i)
	}
}
