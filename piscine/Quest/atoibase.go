package piscine

func AtoiBase(s string, base string) int {
	bRunes := []rune(base)
	bLen := len(bRunes)

	// 1. Валидация базы
	if bLen < 2 {
		return 0
	}

	for i := 0; i < bLen; i++ {
		if bRunes[i] == '+' || bRunes[i] == '-' {
			return 0
		}
		for j := i + 1; j < bLen; j++ {
			if bRunes[i] == bRunes[j] {
				return 0
			}
		}
	}

	// Карта (или функция) для быстрого поиска индекса символа в базе
	// Превращаем строку s в руны
	sRunes := []rune(s)
	res := 0

	// 2. Перевод строки в десятичное число
	for i := 0; i < len(sRunes); i++ {
		// Ищем индекс символа sRunes[i] в базе base
		index := -1
		for j := 0; j < bLen; j++ {
			if sRunes[i] == bRunes[j] {
				index = j
				break
			}
		}

		// Если символ из строки не найден в базе, возвращаем 0
		if index == -1 {
			return 0
		}

		// Основная формула перевода в десятичную систему
		res = res*bLen + index
	}

	return res
}
