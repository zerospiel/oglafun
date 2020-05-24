package power

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_countDigits(t *testing.T) {
	for _, tc := range []struct {
		name     string
		in       int
		expected int
	}{
		{
			name:     "0; expect 1",
			expected: 1,
		},
		{
			name:     "1; expect 1",
			in:       1,
			expected: 1,
		},
		{
			name:     "10; expect 4",
			in:       10,
			expected: 4,
		},
		{
			name:     "4; expect 2",
			in:       4,
			expected: 2,
		},
		{
			name:     "16; expect 5",
			in:       16,
			expected: 5,
		},
		{
			name:     "17; expect 6",
			in:       17,
			expected: 6,
		},
		{
			name: "-1; expect 0 due to int cast",
			in:   -1,
		},
	} {
		t.Run(tc.name, func(t *testing.T) {
			require.Equal(t, tc.expected, countDigits(tc.in))
		})
	}
}
