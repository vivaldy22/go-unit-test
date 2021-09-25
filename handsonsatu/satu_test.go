package handsonsatu

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAdd(t *testing.T) {
	t.Run("success 1 + 2", func(t *testing.T) {
		actual := Add(1, 2)
		expected := 3

		assert.Equal(t, expected, actual)
	})

	t.Run("success 2 + 5", func(t *testing.T) {
		actual := Add(2, 5)
		expected := 7

		assert.Equal(t, expected, actual)
	})

	t.Run("success 5 + 10", func(t *testing.T) {
		actual := Add(5, 10)
		expected := 15

		assert.Equal(t, expected, actual)
	})
}
