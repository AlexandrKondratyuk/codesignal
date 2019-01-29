package _4_Exploring_the_Water

import "fmt"

func main() {
	//str := "hello"
	str := "aabb"

	res := palindromeRearranging(str)
	fmt.Println(res)
}

func palindromeRearranging(inputString string) bool {
	mapUniquValues := map[string]int{} // map for keeping quantities of every letter in string value
	numberEven := 0                    // for saving quantity of EVEN number, if > 1, return false

	if len(inputString) == 1 {
		return true
	} else {
		for _, val := range inputString {
			mapUniquValues[string(val)]++
		}
	}

	for _, val := range mapUniquValues {
		if val%2 != 0 {
			numberEven++
		}
	}

	if numberEven > 1 {
		return false
	}

	return true
}

