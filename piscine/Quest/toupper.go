package piscine

func ToUpper(s string) string {
	res := []rune(s)
	for i, r := range res {
		if r >= 'a' && r <= 'z' {
			res[i] = r - 32
		}
	}
	return string(res)
}
