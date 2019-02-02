package _6_RainOfReason

import "fmt"

func main() {
	inputArray := []int{1, 2, 1}
	elemToReplace := 1
	substitutionElem := 3

	result := arrayReplace(inputArray, elemToReplace,substitutionElem)
	fmt.Println(result)
}

func arrayReplace(inputArray []int, elemToReplace int, substitutionElem int) []int {
	outputArr := make([]int, 0,len(inputArray))

	for _, val := range inputArray {
		if val == elemToReplace {
			outputArr = append(outputArr ,substitutionElem)
			continue
		}
		outputArr = append(outputArr, val)
	}
	return outputArr
}