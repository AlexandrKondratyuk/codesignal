package introGates

func main() {}

func metroCard(lastNumberOfDays int) []int {
	if lastNumberOfDays == 28 || lastNumberOfDays == 30 {
		return []int{31}
	} else {
		return []int{28, 30, 31}
	}
	return []int{}
}
