package power

import "math"

func countDigits(n int) int {
	// assume that 2^n == 10^(nlog2)
	// assume that x == nlog2
	// number of digits in 10^x == x+1
	// other words number of digits in 2^n == (nlog2) + 1
	return int(float64(n)*math.Log10(2.) + 1.)
}
