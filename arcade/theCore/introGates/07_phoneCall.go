package introGates

func phoneCall(min1 int, min2_10 int, min11 int, s int) int {
	res := 0
	switch {
	case s < min1:
		res = 0
	case s <= (min2_10*9+min1):
		s -= min1
		res += 1
		if s/min2_10 <= 9 {
			res += s/min2_10
		} else {
			res += 9
		}
	default:
		s -= (min1 + min2_10*9)
		res = 10 + s/min11
	}
	return res
}
