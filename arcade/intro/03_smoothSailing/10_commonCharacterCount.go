package _3_smoothSailing

import (
	"fmt"
	"math"
)

func main() {
	s1 := "aabcc"
	s2 := "adcaa"
	fmt.Println(commonCharacterCount(s1, s2))
}

func commonCharacterCount(s1 string, s2 string) int {
	//create maps for transforming strings into map
	map1 := make(map[string]int)
	map2 := make(map[string]int)
	result := 0

	// add to map key:=letter & value:=quantity of repeating
	for _, v1 := range s1 {
		map1[string(v1)]++
	}
	for _, v2 := range s2 {
		map2[string(v2)]++
	}

	// check minimum quantity of repeating letters in both maps
	for key, _ := range map1 {
		v1, ok1 := map1[key]
		v2, ok2 := map2[key]
		if ok1 && ok2 {
			min := math.Min(float64(v1), float64(v2))
			result += int(min)
			fmt.Println(result)
		}
	}

	return result
}
