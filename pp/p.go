package pp

import (
	"unicode/utf8"
)

func isp(s string) bool {
	if len(s) == 0 || len(s) == 1 {
		return true
	}
	var (
		i int
		j = len(s)
	)
	for {
		start, size := utf8.DecodeRuneInString(s[i:])
		end, endS := utf8.DecodeLastRuneInString(s[:j])
		if start != end {
			return false
		}
		i += size
		j -= endS
		if i >= len(s)/2 {
			break
		}

	}
	return true
}

func reverse(s string) string {
	l := len(s)
	b := make([]byte, l)
	for i := 0; i < l; {
		r, size := utf8.DecodeRuneInString(s[i:])
		i += size
		_ = utf8.EncodeRune(b[l-i:], r)
	}
	return string(b)
}

func isp2(s string) bool {
	return s == reverse(s)
}

func isp3(s string) bool {
	rr := []rune(s)
	for i := range rr {
		if rr[i] != rr[len(rr)-i-1] {
			return false
		}
	}
	return true
}
