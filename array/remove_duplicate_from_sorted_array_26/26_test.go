package remove_duplicate_from_sorted_array_26

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test26(t *testing.T) {
	testCases := []struct {
		name         string
		input        []int
		expectedNums []int
	}{
		{
			name:         "case 1",
			input:        []int{1, 1, 2},
			expectedNums: []int{1, 2},
		},
		{
			name:         "case 2",
			input:        []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4},
			expectedNums: []int{0, 1, 2, 3, 4},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			output := removeDuplicates(tc.input)

			expectLen := len(tc.expectedNums)
			assert.Equal(t, expectLen, output)
			for i := 0; i < expectLen; i++ {
				assert.Equal(t, tc.expectedNums[i], tc.input[i])
			}
		})
	}
}
