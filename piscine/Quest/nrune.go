package piscine

func NRune(s string, n int) rune {
	i := 1
	for _, r := range s {
		if i == n {
			return r
		}
		i++
	}
	return 0
}
