package hello_world

import (
	"testing"
	"github.com/kmehran1106/learn-go-with-tests/common"
)

func Test_Hello(t *testing.T) {
	t.Run("Saying hello to people in english", func(t *testing.T) {
		// Given
		expected := "Hello, Chris!"

		// When
		result := Hello("Chris", "english")

		// Then
		common.AssertCorrectMessage(t, expected, result)
	})

	t.Run("Saying hello to people in spanish", func(t *testing.T) {
		// Given
		expected := "Hola, Chris!"

		// When
		result := Hello("Chris", "spanish")

		// Then
		common.AssertCorrectMessage(t, expected, result)
	})

	t.Run("Saying hello to people in french", func(t *testing.T) {
		// Given
		expected := "Bonjour, Chris!"

		// When
		result := Hello("Chris", "french")

		// Then
		common.AssertCorrectMessage(t, expected, result)
	})

	t.Run("Saying hello to the world in english when empty", func(t *testing.T) {
		// Given
		expected := "Hello, World!"

		// When
		result := Hello("", "")

		// Then
		common.AssertCorrectMessage(t, expected, result)
	})
}
