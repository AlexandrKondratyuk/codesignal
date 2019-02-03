package _7_throughTheFog

import (
	"fmt"
	"math"
)

func main() {
	//a := []int{2, 4, 7}
	a := []int{1, 1, 3, 4}
	fmt.Println(absoluteValuesSumMinimization(a))
}

func absoluteValuesSumMinimization(a []int) int {
	var (
		resultValue int
		minSum      float64
		currentSum  float64
	)

	for currentIndex := 0; currentIndex < len(a); currentIndex++ {

		for i := 0; i < len(a); i++ {
			currentSum += math.Abs(float64(a[i] - a[currentIndex]))

			if currentIndex == 0 && i == len(a)-1 {
				minSum = currentSum
				resultValue = a[currentIndex]
			} else if currentIndex != 0 && i == len(a)-1 && minSum > currentSum {
				minSum = currentSum
				resultValue = a[currentIndex]
			}
		}
		currentSum = 0
	}

	return resultValue
}
