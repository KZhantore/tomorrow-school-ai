package piscine

func Index(s, substr string) int {
	rs, rsub := []rune(s), []rune(substr)
	for i := 0; i <= len(rs)-len(rsub); i++ {
		match := true
		for j := 0; j < len(rsub); j++ {
			if rs[i+j] != rsub[j] {
				match = false
				break
			}
		}
		if match {
			return i
		}
	}
	return -1
}
