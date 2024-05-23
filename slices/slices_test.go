package slices

import (
	"slices"
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
	t.Run("Sum of numbers in a slice of length 5", func(t *testing.T) {
		// Given
		numbers := []int{1, 2, 3, 4, 5}

		// When
		result := Sum(numbers)
		expected := 15

		// Then
		assertCorrectMessage(t, strconv.Itoa(expected), strconv.Itoa(result))
	})

	t.Run("Sum of numbers in a slice with length 3", func(t *testing.T) {
		// Given
		numbers := []int{1, 2, 3}

		// When
		result := Sum(numbers)
		expected := 6

		// Then
		assertCorrectMessage(t, strconv.Itoa(expected), strconv.Itoa(result))
	})
}

func Test_SumAll(t *testing.T) {
	t.Run("Sum of numbers in multiple slices", func(t *testing.T) {
		// Given
		numbers1 := []int{1, 2}
		numbers2 := []int{0, 9}

		// When
		result := SumAll(numbers1, numbers2)
		expected := []int{3, 9}

		// Then
		if !slices.Equal(result, expected) {
			t.Errorf("Expected value: %v; But returned %v", expected, result)
		}
	})
}
