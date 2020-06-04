package pp

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_isp(t *testing.T) {
	for _, tc := range []struct {
		name     string
		in       string
		expected bool
	}{
		{
			name:     "empty; true",
			expected: true,
		},
		{
			name: "not; false",
			in:   "string",
		},
		{
			name:     "is; true",
			in:       "12344321",
			expected: true,
		},
		{
			name:     "utf is; true",
			in:       "123ø44ø321",
			expected: true,
		},
		{
			name:     "is not even; true",
			in:       "アジア",
			expected: true,
		},
	} {
		t.Run(tc.name, func(t *testing.T) {
			require.Equal(t, tc.expected, isp(tc.in))
			require.Equal(t, tc.expected, isp2(tc.in))
			require.Equal(t, tc.expected, isp3(tc.in))
		})
	}
}

var vlp = "neveroddoreven"

func Benchmark_isp(b *testing.B) {
	for i := 0; i < b.N; i++ {
		isp(vlp)
	}
}
func Benchmark_isp2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		isp2(vlp)
	}
}
func Benchmark_isp3(b *testing.B) {
	for i := 0; i < b.N; i++ {
		isp3(vlp)
	}
}
