package _1_The_Journey_Begins

func centuryFromYear(year int) int {
	if year/100 == 0 {
		return year / 100
	} else {
		return (year / 100) + 1
	}
}
