package integers

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

func Test_Add(t *testing.T) {
	t.Run("Adding two numbers", func(t *testing.T) {
		// Given
		expected := 4

		// When
		result := Add(2, 2)

		// Then
		assertCorrectMessage(t, strconv.Itoa(expected), strconv.Itoa(result))
	})
}
