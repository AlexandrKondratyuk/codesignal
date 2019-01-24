package atTheCrossroads

func knapsackLight(value1 int, weight1 int, value2 int, weight2 int, maxW int) int {
	//if all <= maxW
	if weight1+weight2 <= maxW {
		return value1+value2
	}

	//if all > maxW
	if weight1>maxW && weight2>maxW  {
		return 0
	}

	if weight1<=maxW && weight2<=maxW  {
		if value1>value2 {
			return value1
		} else {
			return value2
		}
	}
	if weight1>maxW || weight2>maxW  {
		if weight1>maxW {
			return value2
		} else {
			return value1
		}
	}
	return 0
}
