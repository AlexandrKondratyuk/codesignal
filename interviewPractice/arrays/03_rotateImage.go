package arrays

import (
	"fmt"
)

func main() {
	a := make([][]int, 3)
	count := 1
	for i := 0; i < 3; i++ {
		a[i] = make([]int,3)

		for j := 0; j < 3; j++ {
			a[i][j] = count
			count++
		}
	}
	fmt.Println(rotateImage(a))
}

func rotateImage(a [][]int) [][]int {
	new := make([][]int, len(a))

	for i := 0; i < len(a);i++ {
		new[i] = make([]int, len(a))
		for j := 0; j < len(a); j++ {
			new[i][j] = a[len(a)-1-j][i]
		}
	}
	return new
}
