package tour

import (
	"testing"
)

func TestFibonacciClosure(t *testing.T) {
	expectedItems := [10]int{0, 1, 1, 2, 3, 5, 8, 13, 21, 34}

	// call fibonacci to get 10 first items
	f := Fibonacci()
	result := []int{}
	for i := 0; i < 10; i++ {
		result = append(result, f())
	}

	if len(expectedItems) != len(result) {
		t.Fatalf("Expected length = %v, actual length = %v", len(expectedItems), len(result))
	}
	for index := range expectedItems {
		if expectedItems[index] != result[index] {
			t.Errorf("Index %v, expected = %v, actual = %v", index, expectedItems[index], result[index])
		}
	}
}
