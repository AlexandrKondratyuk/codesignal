package atTheCrossroads

func arithmeticExpression(a int, b int, c int) bool {
	if a+b==c || a-b==c || a*b==c || c*b==a  {
		return true
	}
	return false
}
