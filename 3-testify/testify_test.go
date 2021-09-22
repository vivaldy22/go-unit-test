package testify_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	. "github.com/vivaldy22/go-unit-test/3-testify"
)

func TestHelloWorld(t *testing.T) {
	actual := HelloWorld()
	assert.Equal(t, "Hello World", actual)
	assert.NotEqual(t, "Hello", actual)
	require.Equal(t, "Hello World", actual)
	require.NotEqual(t, "Hello", actual)
}
