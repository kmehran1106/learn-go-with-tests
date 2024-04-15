package integers


import (
	"strconv"
	"testing"
	"github.com/kmehran1106/learn-go-with-tests/common"
)


func Test_Add(t *testing.T) {
	t.Run("Adding two numbers", func(t *testing.T) {
		// Given
		expected := 4

		// When
		result := Add(2, 2)

		// Then
		common.AssertCorrectMessage(t, strconv.Itoa(expected), strconv.Itoa(result))
	})
}
