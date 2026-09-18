package piscine

// Atoi преобразует строку в число с учетом знаков + / - и валидацией символов.
func Atoi(s string) int {
	// Если строка пустая, конвертировать нечего
	if len(s) == 0 {
		return 0
	}

	result := 0
	sign := 1
	startIndex := 0

	// 1. Проверяем первый символ на наличие знака
	if s[0] == '-' {
		sign = -1
		startIndex = 1 // Начинаем цикл со следующего символа
	} else if s[0] == '+' {
		sign = 1
		startIndex = 1 // Начинаем цикл со следующего символа
	}

	// 2. Основной цикл конвертации
	for i := startIndex; i < len(s); i++ {
		// Если встретили не цифру — строка «грязная», возвращаем 0
		if s[i] < '0' || s[i] > '9' {
			return 0
		}

		digit := int(s[i] - '0')
		result = result*10 + digit
	}

	// 3. Умножаем результат на сохраненный знак
	return result * sign
}
