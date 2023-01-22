package math

func LCD(a int64, b int64) int64 {
	if a < b {
		a, b = b, a
	}
	if b == 0 {
		return a
	}
	return LCD(b, a%b)
}
