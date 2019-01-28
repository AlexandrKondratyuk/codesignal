package _3_smoothSailing

import (
	"fmt"
	"strings"
)

func main() {
	//str := "rab"
	//str := "foo(bar)baz"
	//str := "foo(bar)baz(blim)"
	str := "foo(bar(baz))blim"
	//str := ""
	//str := "()"
	//str := "(abc)d(efg)"
	res := 	reverseInParentheses(str)
	fmt.Println(res)

}

func reverseInParentheses(inputString string) string {
	result := ""
	sliceStr := make([]string, len(inputString))
	startI := 0
	endI := 0

	for i, r := range inputString {    //copy symbols into slice of strings
		sliceStr[i] = string(r)
	}

	for i := 0; i < len(sliceStr); i++ {
		if string(sliceStr[i]) == "(" {    // looking for the 1-st "("
			startI = i
			for j := i; j < len(sliceStr); j++ { // searching nested "("
				if string(sliceStr[j]) == "(" {  // for next "("
					i = j
					continue
				} else if string(sliceStr[j]) == ")" {
					startI = i + 1
					endI = j - 1

					for ; startI < endI; startI, endI = startI+1, endI-1 {  // change letters into parentheses
						sliceStr[startI], sliceStr[endI] = sliceStr[endI], sliceStr[startI]
					}
					sliceStr[i] = "-"
					sliceStr[j] = "-"
					i = -1      // start outer loop from beginning
					break
				}
			}
		}
	}
	result = strings.Join(sliceStr,"")
	result = strings.Replace(result,"-", "",-1)

	return result
}