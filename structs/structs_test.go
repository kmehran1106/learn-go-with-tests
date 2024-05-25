package structs

import (
	"math"
	"testing"
)

func Test_CalculateWithStruct(t *testing.T) {
	assertCorrectMessage := func(t testing.TB, expected, result float64) {
		t.Helper()
		if expected != result {
			t.Errorf("Expected value: %g; But returned %g", expected, result)
		}
	}

	t.Run("Calculate perimeter of a rectangle with width 10 and height 10 using methods", func(t *testing.T) {
		// Given
		rectangle := Rectangle{10.0, 10.0}

		// When
		result := CalculatePerimeter(rectangle)
		expected := 40.0

		// Then
		assertCorrectMessage(t, expected, result)
	})

	t.Run("Calculate area of a rectangle with width 10 and height 10 using methods", func(t *testing.T) {
		// Given
		rectangle := Rectangle{10.0, 10.0}

		// When
		result := CalculateArea(rectangle)
		expected := 100.0

		// Then
		assertCorrectMessage(t, expected, result)
	})
}

func Test_CalculateWithInterface(t *testing.T) {
	assertCorrectMessage := func(t testing.TB, expected, result float64) {
		t.Helper()
		expected, result = math.Round(expected*100)/100, math.Round(result*100)/100
		if expected != result {
			t.Errorf("Expected value: %g; But returned %g", expected, result)
		}
	}

	var tests = []struct {
		shape             Shape
		expectedArea      float64
		expectedPerimeter float64
	}{
		{Rectangle{10.0, 10.0}, 100.0, 40.0},
		{Circle{10.0}, 314.0, 62.8},
	}

	for _, test := range tests {
		// Given

		// When
		resultArea := test.shape.CalculateArea()
		resultPerimeter := test.shape.CalculatePerimeter()

		// Then
		assertCorrectMessage(t, test.expectedArea, resultArea)
		assertCorrectMessage(t, test.expectedPerimeter, resultPerimeter)
	}
}
