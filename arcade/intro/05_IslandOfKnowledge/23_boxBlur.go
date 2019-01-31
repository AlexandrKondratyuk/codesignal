package _5_IslandOfKnowledge

import "fmt"

func main() {
	//inputArr := make([][]int, 3)
	//inputArr[0] = []int{1, 1, 1}
	//inputArr[1] = []int{1, 7, 1}
	//inputArr[2] = []int{1, 1, 1}

	inputArr := make([][]int, 3)
	inputArr[0] = []int{36, 0, 18, 9}
	inputArr[1] = []int{27, 54, 9, 0}
	inputArr[2] = []int{81, 63, 72, 45}

	result := boxBlur(inputArr)
	fmt.Println(result)
}

func boxBlur(image [][]int) [][]int {
	rows, cols := len(image), len(image[0])
	outArr := make([][]int, rows-2)
	average := 0

	for r := 2; r < rows; r++ {
		outArr[r-2] = make([]int, cols-2)
		for c := 2; c < cols; c++ {

			for i := r - 2; i <= r; i++ {
				for j := c - 2; j <= c; j++ {
					average += image[i][j]
					if i == r && j == c {
						average = average / 9
						outArr[r-2][c-2] = average
						average = 0

					}
				}
			}
		}
	}
	return outArr
}