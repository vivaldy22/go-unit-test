package tabletest

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

//buat sebuah func yang menjumlahkan 3 buah angka

func TestAddThreeNum(t *testing.T) {
	type testCase struct {
		name     string
		num1     int
		num2     int
		num3     int
		expected int
		//num1, num2, num3 int
	}

	tests := []testCase{
		{
			name:     "1 + 2 + 3",
			num1:     1,
			num2:     2,
			num3:     3,
			expected: 6,
		},
		{
			name:     "-1 + 10 + 5",
			num1:     -1,
			num2:     10,
			num3:     4,
			expected: 14,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			hasil := AddThreeNum(tt.num1, tt.num2, tt.num3)
			assert.Equalf(t, tt.expected, hasil, "hasil tidak sama")
		})
	}
}
