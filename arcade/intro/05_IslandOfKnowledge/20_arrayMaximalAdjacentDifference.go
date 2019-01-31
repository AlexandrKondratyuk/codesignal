package _5_IslandOfKnowledge

import (
	"fmt"
	"math"
)

func main() {
	arr := []int{2, 4, 1, 0}
	result := arrayMaximalAdjacentDifference(arr)
	fmt.Println(result)
}

func arrayMaximalAdjacentDifference(inputArray []int) int {
	result, delta := 0, 0

	for i := 1; i < len(inputArray); i++ {
		delta = int(math.Abs(float64(inputArray[i] - inputArray[i-1])))
		if delta > result{
			result = delta
		}
	}

	return result
}