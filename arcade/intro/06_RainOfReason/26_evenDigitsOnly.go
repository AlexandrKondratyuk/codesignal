package _6_RainOfReason

import (
	"fmt"
	"math"
)

func main() {
	n := 642386
	result := evenDigitsOnly(n)
	fmt.Println(result)
}

func evenDigitsOnly(n int) bool {
	for i := 0; i <= int(math.Log10(float64(n))); i++ {
		//currentNumber := int(math.Log10(float64(n / int(math.Pow10(i)))))
		currentNumber := n / int(math.Pow10(i))
		if currentNumber%2 != 0 {
			return false
		}
	}
	return true
}
