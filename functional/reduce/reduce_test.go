package functional

import (
	"testing"
)

func TestReduceSum(t *testing.T) {
	t.Parallel()
	inputItems := []int{1, 2, 3, 4}
	expectedResult := 10
	initValue := 0

	result := Reduce(inputItems, initValue, func(cur, next int) int {
		return cur + next
	})

	if result != expectedResult {
		t.Errorf("expected result is %v, the current result is %v", expectedResult, result)
	}
}

func TestReduceMultiply(t *testing.T) {
	t.Parallel()
	inputItems := []int{1, 2, 3, 4}
	expectedResult := 24
	initValue := 1

	result := Reduce(inputItems, initValue, func(cur, next int) int {
		return cur * next
	})

	if result != expectedResult {
		t.Errorf("expected result is %v, the current result is %v", expectedResult, result)
	}
}

func TestReduceConcatString(t *testing.T) {
	t.Parallel()
	inputItems := []string{"X", "Y", "Z"}
	expectedResult := "XYZ"
	initValue := ""

	result := Reduce(inputItems, initValue, func(cur, next string) string {
		return cur + next
	})

	if result != expectedResult {
		t.Errorf("expected result is %s, the current result is %s", expectedResult, result)
	}
}
