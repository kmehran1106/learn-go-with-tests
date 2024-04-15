package common

import (
	"testing"
)

func AssertCorrectMessage(t testing.TB, expected, result string) {
	t.Helper()
	if expected != result {
		t.Errorf("Expected value: %q; But returned %q", expected, result)
	}
}
