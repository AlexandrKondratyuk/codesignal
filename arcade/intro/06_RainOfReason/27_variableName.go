package _6_RainOfReason

import "fmt"

func main() {
	//name := "var_1__Int"
	name := "qq-q"
	//name := "2w2"
	result := variableName(name)
	fmt.Println(result)
}

func variableName(name string) bool {
	for _, r := range name {
		if (name[0] >= '0' && name[0] <= '9') ||
			(r < 'a' || r > 'z') &&
				(r < 'A' || r > 'Z') &&
				(r != '_') &&
				(r < '0' || r > '9') {
			return false
		}
	}
	return true
}