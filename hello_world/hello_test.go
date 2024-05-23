package hello_world

import (
	"testing"
)

func assertCorrectMessage(t testing.TB, expected, result string) {
	t.Helper()
	if expected != result {
		t.Errorf("Expected value: %q; But returned %q", expected, result)
	}
}

func Test_Hello(t *testing.T) {
	t.Run("Saying hello to people in english", func(t *testing.T) {
		// Given
		expected := "Hello, Chris!"

		// When
		result := Hello("Chris", "english")

		// Then
		assertCorrectMessage(t, expected, result)
	})

	t.Run("Saying hello to people in spanish", func(t *testing.T) {
		// Given
		expected := "Hola, Chris!"

		// When
		result := Hello("Chris", "spanish")

		// Then
		assertCorrectMessage(t, expected, result)
	})

	t.Run("Saying hello to people in french", func(t *testing.T) {
		// Given
		expected := "Bonjour, Chris!"

		// When
		result := Hello("Chris", "french")

		// Then
		assertCorrectMessage(t, expected, result)
	})

	t.Run("Saying hello to the world in english when empty", func(t *testing.T) {
		// Given
		expected := "Hello, World!"

		// When
		result := Hello("", "")

		// Then
		assertCorrectMessage(t, expected, result)
	})
}
