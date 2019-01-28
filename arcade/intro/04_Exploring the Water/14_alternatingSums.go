package _4_Exploring_the_Water

import "fmt"

func main() {
	a := []int{50, 60, 60, 45, 70}
	res := alternatingSums(a)
	fmt.Println(res)
}

func alternatingSums(a []int) []int {
	result := []int{0, 0}

	for index, val := range a {
		if index%2 == 0 {
			result[0] += val
		} else {
			result[1] += val
		}
	}

	return result
}
