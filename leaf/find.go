package leaf

func findJudge(N int, trust [][]int) int {
	var (
		a, b = make(map[int]int), make(map[int]int)
	)
	for _, sl := range trust {
		a[sl[1]]++
		b[sl[0]]++
	}

	for i := range trust {
		if b[i] == 0 && a[i] == N-1 {
			return i
		}
	}

	return -1
}
