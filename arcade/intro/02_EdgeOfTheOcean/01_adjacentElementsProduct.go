package EdgeOfTheOcean

func adjacentElementsProduct(inputArray []int) int {
	var result int = inputArray[0]*inputArray[1]
	for i:=1; i < len(inputArray); i++ {
		if len(inputArray) <= 2 {
			return result
		}
		if inputArray[i]*inputArray[i-1] > result {
			result = inputArray[i]*inputArray[i-1]
		}
	}
	return result
}