package functional

import (
	"strings"
	"testing"
)

func TestMap(t *testing.T) {
	inputItems := []string{"a", "b", "c"}
	expectedItems := []string{"A", "B", "C"}

	mappedItems := Map(inputItems, strings.ToUpper)

	if len(expectedItems) != len(mappedItems) {
		t.Fatalf("Expected length = %v, actual length = %v", len(expectedItems), len(mappedItems))
	}
	for index := range expectedItems {
		if expectedItems[index] != mappedItems[index] {
			t.Errorf("Index %v, expected = %v, actual = %v", index, expectedItems[index], mappedItems[index])
		}
	}
}
