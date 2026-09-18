package piscine

func ToLower(s string) string {
	res := []rune(s)
	for i, r := range res {
		if r >= 'A' && r <= 'Z' {
			res[i] = r + 32
		}
	}
	return string(res)
}
