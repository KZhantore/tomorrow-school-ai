package piscine

func JumpOver(str string) string {
	// Если строка пустая или в ней меньше 3 символов,
	// возвращаем только перевод строки.
	if len(str) < 3 {
		return "\n"
	}
	result := ""
	// Начинаем с индекса 2 (это 3-й символ) и шагаем по 3
	for i := 2; i < len(str); i += 3 {
		result += string(str[i])
	}
	// Добавляем обязательный переход на новую строку
	return result + "\n"
}
