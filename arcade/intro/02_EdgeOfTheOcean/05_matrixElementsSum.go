package main

import (
	"fmt"
)

func main() {
	m := make([][]int, 3)
	m[0] = []int{0, 1, 1, 2}
	m[1] = []int{0, 5, 0, 0}
	m[2] = []int{2, 0, 3, 3}

	res := matrixElementsSum(m)
	fmt.Println(res)
}

func matrixElementsSum(matrix [][]int) int {
	result := 0

	for c := 0; c < len(matrix[0]); c++ {
		for r := 0; r < len(matrix); r++ {
			if matrix[r][c] == 0 {
				break
			}
			result += matrix[r][c]
		}
	}
	return result
}
