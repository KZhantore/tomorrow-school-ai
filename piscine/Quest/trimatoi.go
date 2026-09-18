package piscine

func TrimAtoi(s string) int {
	sign := 1
	res := 0
	hasDigits := false // Флаг: начались ли уже цифры

	for i := 0; i < len(s); i++ {
		if s[i] >= '0' && s[i] <= '9' {
			hasDigits = true
			res = res*10 + int(s[i]-'0')
		} else if s[i] == '-' && !hasDigits {
			// Минус меняет знак, ТОЛЬКО если цифры еще не начались
			sign = -1
		} else if s[i] == '+' && !hasDigits {
			// Плюс фиксирует позицию до цифр
			sign = 1
		}
	}

	return sign * res
}
