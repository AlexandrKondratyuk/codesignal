package _5_IslandOfKnowledge

import (
	"fmt"
	"sort"
)

func main() {
	inputArray := []int{1, 4, 10, 6, 2}
	result := avoidObstacles(inputArray)
	fmt.Println(result)
}

func avoidObstacles(inputArray []int) int {
	jumpPoints := 2
	index := 0  // index of current value from array for comparing

	sort.Ints(inputArray)

	currentVal := inputArray[index]

	for i := 0; i <= inputArray[len(inputArray)-1]; i = i + jumpPoints {
		if currentVal == i {
			jumpPoints++
			index = 0
			currentVal = inputArray[index]
			i = 0
		} else if currentVal < i {
			index++
			currentVal = inputArray[index]
			i = i - jumpPoints
			if index > len(inputArray)-1 {
				return jumpPoints
			}
			continue
		}
	}
	return jumpPoints
}
