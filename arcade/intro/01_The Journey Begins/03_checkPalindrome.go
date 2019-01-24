package _1_The_Journey_Begins

func checkPalindrome(inputString string) bool {
	myRune := []rune(inputString)
	for i, j := 0, len(myRune)-1; i < len(myRune)/2; i, j = i+1, j-1 {
		myRune[i], myRune[j] = myRune[j], myRune[i]
	}
	str := string(myRune)
	return inputString == str
}
