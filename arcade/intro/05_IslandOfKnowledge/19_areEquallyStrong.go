package _5_IslandOfKnowledge

import "fmt"

func main() {
	yourLeft := 15
	yourRight := 10
	friendsLeft := 15
	friendsRight := 9

	result := areEquallyStrong(yourLeft,yourRight, friendsLeft, friendsRight)

	fmt.Println(result)
}

func areEquallyStrong(yourLeft int, yourRight int, friendsLeft int, friendsRight int) bool {
	if (yourLeft == friendsLeft && yourRight == friendsRight) ||
		(yourLeft == friendsRight && yourRight == friendsLeft) {
		return true
	}

	return false
}
