package tabletest_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	. "github.com/vivaldy22/go-unit-test/5-table-test"
)

func TestHelloWorld(t *testing.T) {
	tests := []struct {
		testName string
		request  string
		expected string
	}{
		{
			testName: "should return Hello enigma",
			request:  "enigma",
			expected: "Hello enigma",
		},
		{
			testName: "should return Hello all",
			request:  "all",
			expected: "Hello all",
		},
	}

	for _, tt := range tests {
		t.Run(tt.testName, func(t *testing.T) {
			actual := Hello(tt.request)
			assert.Equal(t, tt.expected, actual)
		})
	}
}
