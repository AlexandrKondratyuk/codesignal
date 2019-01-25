package main

import (
	"fmt"
)

func main() {
	seq := make([]int, 3)
	fmt.Println(int(-1e4))
	//seq = []int{10, 1, 2, 3, 4, 5}
	seq = []int{1, 2, 1, 2}
	res := almostIncreasingSequence(seq)
	fmt.Println(res)
}

func almostIncreasingSequence(sequence []int) bool {
	counter := 0

	//default min values from -10**5 ≤ sequence[i] ≤ 10**5
	first := -100001
	second := -100001


	for i := 0; i < len(sequence); i++ {
		if sequence[i] > first {
			second, first = first, sequence[i]
		} else if sequence[i] > second {
			first = sequence[i]
			counter++
		} else {
			counter++
		}
	}
	return counter < 2
}