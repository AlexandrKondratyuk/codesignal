package _6_RainOfReason

import "fmt"

func main() {
	cell1 := "A1"
	cell2 := "C3"

	result := chessBoardCellColor(cell1, cell2)

	fmt.Println(result)
}

func chessBoardCellColor(cell1 string, cell2 string) bool { // A=65 '1'=49

	if (cell1[0]%2 == cell1[1]%2) == (cell2[0]%2 == cell2[1]%2) {
		return true
	}

	return false
}
