package nametbd

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_idkfunc(t *testing.T) {
	for _, tc := range []struct {
		name     string
		in       []int
		expected int
	}{
		{
			name:     "empty; expect minus one",
			in:       []int{},
			expected: -1,
		},
		{
			name:     "nil; expect minus one",
			expected: -1,
		},
		{
			name:     "w major; expect major",
			in:       []int{1, 2, 3, 2, 3, 2, 4, 2, 2, 8, 2, 2, 2},
			expected: 2,
		},
		{
			name:     "wo major; expect minus one",
			in:       []int{1, 2, 3, 4, 5, 6, 2, 2},
			expected: -1,
		},
		{
			name:     "wo major (exactly len/2); expect minus one",
			in:       []int{1, 2, 3, 4, 5, 6, 2, 2, 2, 2},
			expected: -1,
		},
		{
			name:     "w major (exactly len/2+1); expect minus one",
			in:       []int{1, 100, 3, 4, 5, 6, 100, 100, 100, 100, 100},
			expected: 100,
		},
		{
			name:     "50/50; expect minus one",
			in:       []int{1, 1, 1, 1, 2, 2, 2, 2},
			expected: -1,
		},
		{
			name:     "51/49; expect major",
			in:       []int{1, 1, 1, 1, 2, 2, 2, 2, 1},
			expected: 1,
		},
	} {
		t.Run(tc.name, func(t *testing.T) {
			require.Equal(t, tc.expected, idkfunc(tc.in))
		})
	}
}
