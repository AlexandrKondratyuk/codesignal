package arrays

import (
	"math"
	"strconv"
)

func main() {
	crypt := make([]string, 3)
	//crypt[0] = "SEND"
	//crypt[1] = "MORE"
	//crypt[2] = "MONEY"
	//solution := make([][]string, 8)
	//solution[0] = []string{"O", "0"}
	//solution[1] = []string{"M", "1"}
	//solution[2] = []string{"Y", "2"}
	//solution[3] = []string{"E", "5"}
	//solution[4] = []string{"N", "6"}
	//solution[5] = []string{"D", "7"}
	//solution[6] = []string{"R", "8"}
	//solution[7] = []string{"S", "9"}

	crypt[0] = "TEN"
	crypt[1] = "TWO"
	crypt[2] = "ONE"
	solution := make([][]string, 5)
	solution[0] = []string{"O", "1"}
	solution[1] = []string{"T", "0"}
	solution[2] = []string{"W", "9"}
	solution[3] = []string{"E", "5"}
	solution[4] = []string{"N", "4"}

	isCryptSolution(crypt, solution)
}

func isCryptSolution(crypt []string, solution [][]string) bool {
	// slice with result's value of each word
	resultsSlice := make([]int, len(crypt))
	// map with key:=letter(string) an value:=integer(int)
	cryptMap := make(map[string]int, len(solution))

	// convert STRING values into INT and add them into the MAP storage
	for _, sol := range solution {
		if val, err := strconv.Atoi(sol[1]); err == nil {
			cryptMap[sol[0]] = val
		}
	}

	// sum values each word
	for i, word := range crypt {
		var wordLength = len(word)
		var exp int

		for n, letter := range word {
			if cryptMap[string(letter)] == 0  && n == 0 && len(word) > 1{
				return false
			}

			wordLength--
			exp = int(math.Pow10(wordLength))
			resultsSlice[i] += cryptMap[string(letter)] * exp
		}
	}

	return resultsSlice[0]+resultsSlice[1] == resultsSlice[2]
}
