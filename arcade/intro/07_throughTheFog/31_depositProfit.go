package _7_throughTheFog

import "fmt"

func main() {
	deposit := 100
	rate := 20
	threshold := 170

	result := depositProfit(deposit, rate, threshold)

	fmt.Println(result)
}

func depositProfit(deposit int, rate int, threshold int) int {
	years := 0
	total := float64(deposit)

	for total < float64(threshold) {
		years++
		total += total * float64(rate) / 100
	}

	return years
}
