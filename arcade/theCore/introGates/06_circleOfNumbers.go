package introGates

func circleOfNumbers(n int, firstNumber int) int {
	if firstNumber >= n/2 {
		return firstNumber - n/2
	}
	return firstNumber + n/2
}
