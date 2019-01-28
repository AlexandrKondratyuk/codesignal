package _4_Exploring_the_Water

import (
	"fmt"
	"strings"
)

func main() {
	picture := []string{"abc", "ded"}
	//picture := []string{"a"}
	result := addBorder(picture)

	for _, v := range result {
		fmt.Println(v)
	}
}
func addBorder(picture []string) []string {
	result := []string{}
	tempString := ""

	for index, val := range picture {
		tempString = "*" + val + "*"
		if index == 0 || index == len(picture)-1 {
			if index == 0  && len(picture) == 1 {
				result = append(result, strings.Repeat("*", len(picture[0])+2), tempString, strings.Repeat("*", len(picture[0])+2))
			} else if index == 0{
				result = append(result, strings.Repeat("*", len(picture[0])+2), tempString)
			} else {
				result = append(result, tempString, strings.Repeat("*", len(picture[0])+2))
			}
		} else {
			result = append(result, tempString)
		}
	}
	return result
}
