package _4_Exploring_the_Water

import "fmt"

func main() {
	a := []int{1,1,4}
	b := []int{1,2,3}
	res := areSimilar(a, b)
	fmt.Println(res)
}

func areSimilar(a []int, b []int) bool {
	notEqualNumSliceA := make([]int, 0)
	notEqualNumSliceB := make([]int, 0)

	// search for not equal values
	// if not equal more then 2 --> return false
	for index := range a {
		if a[index] == b[index] {
			continue
		}
		if len(notEqualNumSliceA) > 1 {
			return false
		}
		notEqualNumSliceA = append(notEqualNumSliceA, a[index])
		notEqualNumSliceB = append(notEqualNumSliceB, b[index])
	}

	// if swoped values are equal - return true
	if len(notEqualNumSliceA) > 0 {
		if notEqualNumSliceA[0] == notEqualNumSliceB[0] && notEqualNumSliceA[1] == notEqualNumSliceB[1] {
			return true
		} else if notEqualNumSliceA[0] == notEqualNumSliceB[1] && notEqualNumSliceA[1] == notEqualNumSliceB[0] {
			return true
		} else {
			return false
		}
	}

	return true
}
