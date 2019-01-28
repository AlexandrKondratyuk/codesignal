package _3_smoothSailing

import (
	"fmt"
	"sort"
)

func main() {
	arr := []int{-1, 150, 190, 170, -1, -1, 160, 180}
	fmt.Println(sortByHeight(arr))
}

func sortByHeight(a []int) []int {
	result := make([]int, 0, len(a)) // create slice of results
	tempArr := make([]int, len(a))   // create temp slice for sorting
	copy(tempArr, a)                 // copy original slice into temp slice
	sort.Ints(tempArr)               // sort our temp slice

	for _, val := range a {          //check if val == -1, then copy to result's slice
		if val == -1 {
			result = append(result, -1)
		} else {                     // else copy one by one values which sorted ascending
			for i, v := range tempArr {
				if tempArr[i] == -1 {
					continue
				} else {              // if copy value to result's slice, we remove value from temp slice
					result = append(result, v)
					tempArr = tempArr[i+1:]
					break
				}
			}
		}
	}
	return result
}