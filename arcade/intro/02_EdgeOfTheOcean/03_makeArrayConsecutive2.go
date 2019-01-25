package main

import (
	"fmt"
)

func main(){
	res := makeArrayConsecutive2([]int{1,5})
	fmt.Println(res)
}

func makeArrayConsecutive2(statues []int) int {
	//initialize min and max values equals to statues[0]
	min, max := statues[0], statues[0]

	if len(statues) == 1 {
		return  0
	}

	// finds min & max values
	for i:= 1; i < len(statues); i++{
		if statues[i] < min {
			min = statues[i]
		} else if statues[i] > max {
			max = statues[i]
		}
	}

	//calculate and return number of missed statues
	sliceLen := len(statues)
	difBetweenMinMax := max - min
	return difBetweenMinMax - sliceLen + 1
}
