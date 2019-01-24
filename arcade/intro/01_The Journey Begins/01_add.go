package _1_The_Journey_Begins

import "fmt"

func main() {
	x := "text"
	// xbytes := []byte(x)
	xbytes := []byte(x)
	xbytes[0] = 'T'

	fmt.Println(string(xbytes)) // выводит Text
}
