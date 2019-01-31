package _5_IslandOfKnowledge

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	//inputString := "172.16.254.1"
	inputString := "1"
	result := isIPv4Address(inputString)
	fmt.Println(result)
}

func isIPv4Address(inputString string) bool {

	sliceString := strings.Split(inputString, ".")

	for _, val := range sliceString {
		num, ok := strconv.Atoi(val)
		if ok != nil || num > 255 || num < 0 || len(sliceString) != 4{
			return false
		}
	}

	return true
}