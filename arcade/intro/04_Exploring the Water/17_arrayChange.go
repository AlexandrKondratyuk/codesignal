package _4_Exploring_the_Water

import "fmt"

func main() {
	inputArray := []int{2, 3, 3, 5, 5, 5, 4, 12, 12, 10, 15}
	result := arrayChange(inputArray)
	fmt.Println(result)

}

func arrayChange(inputArray []int) int {
	currentVal := -100001
	moves := 0

	for i := 0; i < len(inputArray); {
		if inputArray[i] <= currentVal {
			inputArray[i]++
			moves++
			continue
		}
		currentVal = inputArray[i]
		i++
	}
	return moves
}