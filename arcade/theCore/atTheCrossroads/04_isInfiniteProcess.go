package atTheCrossroads

func isInfiniteProcess(a int, b int) bool {
	if b < a || (b%2 != 0 && a%2 != 1) || (b%2 != 1 && a%2 != 0){
		return true
	}
	return false
}
