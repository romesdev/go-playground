package functional

import (
	"testing"
)

func isEven(number int) bool {
	return number%2 == 0
}

func TestFilterIsEven(t *testing.T) {
	t.Parallel()
	inputItems := []int{1, 23, 12, 42, 21, 41}
	expectedItems := []int{12, 42}

	filteredItems := Filter(inputItems, isEven)

	if len(expectedItems) != len(filteredItems) {
		t.Fatalf("Expected length = %v, actual length = %v", len(expectedItems), len(filteredItems))
	}
	for index := range expectedItems {
		if expectedItems[index] != filteredItems[index] {
			t.Errorf("Index %v, expected = %v, actual = %v", index, expectedItems[index], filteredItems[index])
		}
	}
}

func isOdd(number int) bool {
	return number%2 != 0
}

func TestFilterIsOdd(t *testing.T) {
	t.Parallel()
	inputItems := []int{1, 23, 12, 42, 21, 41}
	expectedItems := []int{1, 23, 21, 41}

	filteredItems := Filter(inputItems, isOdd)

	if len(expectedItems) != len(filteredItems) {
		t.Fatalf("Expected length = %v, actual length = %v", len(expectedItems), len(filteredItems))
	}
	for index := range expectedItems {
		if expectedItems[index] != filteredItems[index] {
			t.Errorf("Index %v, expected = %v, actual = %v", index, expectedItems[index], filteredItems[index])
		}
	}
}
