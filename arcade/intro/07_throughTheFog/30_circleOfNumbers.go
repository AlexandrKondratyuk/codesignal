package _7_throughTheFog

import "fmt"

func main() {
	n := 10
	firstNumber := 2

	result := circleOfNumbers(n, firstNumber)

	fmt.Println(result)

}

func circleOfNumbers(n int, firstNumber int) int {
	if firstNumber < n/2 {
		return firstNumber + n/2
	}

	return firstNumber - n/2
}