package lanjutan_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	. "github.com/vivaldy22/go-unit-test/4-lanjutan"
)

func TestHelloWorld(t *testing.T) {
	actual := HelloWorld()
	t.Run("Hello World success", func(t *testing.T) {
		assert.Equal(t, "Hello World", actual)
	})
	t.Run("Hello World not empty", func(t *testing.T) {
		assert.NotEmpty(t, actual)
	})
}