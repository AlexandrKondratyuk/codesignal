package main

import (
	"fmt"
)

func main() {
	fmt.Println(2<<1)
}

func shapeArea(n int) int {
	var result int = 1


	if n == 1 {
		return result
	}
	for i := 2; i <= n; i++ {

		result += 4*(i-1)
	}
	return result
}

