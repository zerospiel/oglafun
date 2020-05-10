package leaf

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_findJudge(t *testing.T) {
	for _, tc := range []struct {
		name    string
		inN     int
		inTrust [][]int
		exp     int
	}{
		{
			name:    "found",
			inN:     5,
			inTrust: [][]int{{1, 3}, {2, 3}, {4, 3}, {4, 1}, {5, 3}, {5, 1}, {5, 4}},
			exp:     3,
		},
		{
			name:    "not found",
			inN:     5,
			inTrust: [][]int{{2, 1}, {3, 1}, {4, 2}, {4, 3}, {4, 5}, {5, 1}},
			exp:     -1,
		},
	} {
		t.Run(tc.name, func(t *testing.T) {
			actual := findJudge(tc.inN, tc.inTrust)
			require.Equal(t, tc.exp, actual)
		})
	}
}
