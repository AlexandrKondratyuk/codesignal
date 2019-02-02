package _6_RainOfReason

import "fmt"

func main() {
	inputString := "crazy"
	result := alphabeticShift(inputString)
	fmt.Println(result)
}

func alphabeticShift(inputString string) string {

	outputSlice := []rune(inputString)

	for index, value := range outputSlice {
		if value == 122 {
			value = 96
		}

		value++
		outputSlice[index] = value
	}

	return string(outputSlice)
}
