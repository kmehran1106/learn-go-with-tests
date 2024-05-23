package arrays

import (
	"strconv"
	"testing"
)

func assertCorrectMessage(t testing.TB, expected, result string) {
	t.Helper()
	if expected != result {
		t.Errorf("Expected value: %q; But returned %q", expected, result)
	}
}

func Test_Sum(t *testing.T) {
	t.Run("Sum of numbers in an array all positive", func(t *testing.T) {
		// Given
		numbers := [5]int{1, 2, 3, 4, 5}

		// When
		result := Sum(numbers)
		expected := 15

		// Then
		assertCorrectMessage(t, strconv.Itoa(expected), strconv.Itoa(result))
	})

	t.Run("Sum of numbers in an array all negative", func(t *testing.T) {
		// Given
		numbers := [5]int{-1, -2, -3, -4, -5}

		// When
		result := Sum(numbers)
		expected := -15

		// Then
		assertCorrectMessage(t, strconv.Itoa(expected), strconv.Itoa(result))
	})
}
