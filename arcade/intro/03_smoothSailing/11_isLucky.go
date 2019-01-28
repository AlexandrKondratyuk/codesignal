package _3_smoothSailing

import (
	"fmt"
	"math"
)

func main() {
	n := 123042
	fmt.Println(isLucky(n))
}

func isLucky(n int) bool {
	quontityOfNumbers := int(math.Log10(float64(n))) + 1

	firstHalf := 0
	secandHalf := 0

	for i := 0; i < quontityOfNumbers; i++ {
		if i < quontityOfNumbers/2 {
			firstHalf += n % 10
		} else {
			secandHalf += n % 10
		}
		n = n / 10
	}
	return firstHalf == secandHalf
}