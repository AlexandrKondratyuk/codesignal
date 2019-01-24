package arrays

func firstNotRepeatingCharacter(s string) string {
	m := make(map[string]int, len(s))
	for i := 0; i < len(s); i++ {
		m[string(s[i])]++
	}
	for i := 0; i < len(s); i++ {
		if m[string(s[i])] == 1 {
			return string(s[i])
		}
	}
	return "_"
}
