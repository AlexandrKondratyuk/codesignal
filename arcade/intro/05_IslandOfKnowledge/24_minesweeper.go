package _5_IslandOfKnowledge

import "fmt"

func main() {
	matrix := make([][]bool, 3)
	matrix[0] = []bool{true, false, false}
	matrix[1] = []bool{false, true, false}
	matrix[2] = []bool{false, false, false}

	result := minesweeper(matrix)
	fmt.Println(result)
}

func minesweeper(matrix [][]bool) [][]int {
	output := make([][]int, len(matrix))
	counter := 0

	for i := 0; i < len(matrix); i++ {
		output[i] = make([]int, len(matrix[0]))

		for j := 0; j < len(matrix[0]); j++ {

			if j != len(matrix[0])-1 && matrix[i][j+1] == true {
				counter++
			}
			if j != 0 && matrix[i][j-1] == true {
				counter++
			}
			if i != len(matrix)-1 && matrix[i+1][j] == true {
				counter++
			}
			if i != 0 && matrix[i-1][j] == true {
				counter++
			}
			if (j != len(matrix[0])-1 && i != len(matrix)-1) && matrix[i+1][j+1] == true {
				counter++
			}
			if (j != 0 && i != 0) && (matrix[i-1][j-1] == true) {
				counter++
			}
			if (j != 0 && i != len(matrix)-1) && matrix[i+1][j-1] == true {
				counter++
			}
			if (j != len(matrix[0])-1 && i != 0) && matrix[i-1][j+1] == true {
				counter++
			}

			output[i][j] = counter
			counter = 0
		}
	}

	return output
}