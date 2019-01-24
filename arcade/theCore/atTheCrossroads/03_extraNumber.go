package atTheCrossroads

func extraNumber(a int, b int, c int) int {
	if a == b {
		return c
	} else if a == c {
		return b
	} else {
		return a
	}
}