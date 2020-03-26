package knslice

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_merge(t *testing.T) {
	for _, tc := range []struct {
		name     string
		in       [][]int
		sliceLen int
		out      []int
	}{
		{
			name:     "one",
			in:       [][]int{{2}, {1}, {0}},
			sliceLen: 1,
			out:      []int{0, 1, 2},
		},
		{
			name:     "two",
			in:       [][]int{{2, 3}, {1, 4}, {0, 5}},
			sliceLen: 2,
			out:      []int{0, 1, 2, 3, 4, 5},
		},
		{
			name:     "three",
			in:       [][]int{{2, 7, 10}, {1, 3, 5}, {0, 4, 6}},
			sliceLen: 3,
			out:      []int{0, 1, 2, 3, 4, 5, 6, 7, 10},
		},
	} {
		t.Run(tc.name, func(t *testing.T) {
			require.Equal(t, tc.out, merge(tc.in, tc.sliceLen))
		})
	}
}

func Benchmark_merge5(b *testing.B) {
	for n := 0; n < b.N; n++ {
		merge([][]int{{2, 7, 10, 12, 15}, {1, 3, 5, 8, 9}, {0, 4, 6, 59, 60}}, 5)
	}
}

func Benchmark_merge_dumb5(b *testing.B) {
	for n := 0; n < b.N; n++ {
		merge_dumb([][]int{{2, 7, 10, 12, 15}, {1, 3, 5, 8, 9}, {0, 4, 6, 59, 60}}, 5)
	}
}
