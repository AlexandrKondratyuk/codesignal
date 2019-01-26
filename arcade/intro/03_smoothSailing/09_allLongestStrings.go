package _3_smoothSailing

func main() {

}

func allLongestStrings(inputArray []string) []string {
	var res []string
	maxLen := 1
	for _, v := range inputArray {
		if len(v) > maxLen {
			maxLen = len(v)
		}
	}
	for _, v := range inputArray {
		if len(v) == maxLen{
			res = append(res, v)
		}
	}
	return res
}

