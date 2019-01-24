package arrays

func firstDuplicate(a []int) int {
	if len(a) < 2 {
		return -1
	}

	myMap := make(map[int]int)

	for index, val := range a {
		_, ok := myMap[val]
		if ok {
			return val
		}
		myMap[val] = index
	}
	return -1
}
