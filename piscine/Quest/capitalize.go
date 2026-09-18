package piscine

func Capitalize(s string) string {
	runes := []rune(s)
	newWord := true

	for i := 0; i < len(runes); i++ {
		// Проверяем, является ли символ буквой или цифрой
		if (runes[i] >= 'a' && runes[i] <= 'z') || (runes[i] >= 'A' && runes[i] <= 'Z') || (runes[i] >= '0' && runes[i] <= '9') {
			if newWord {
				// Первая буква слова -> делаем заглавной
				if runes[i] >= 'a' && runes[i] <= 'z' {
					runes[i] = runes[i] - 32
				}
				newWord = false // Слово началось, остальные символы в нем — не начало
			} else {
				// Буква внутри слова -> делаем строчной
				if runes[i] >= 'A' && runes[i] <= 'Z' {
					runes[i] = runes[i] + 32
				}
			}
		} else {
			// Встретили пробел или спецсимвол (+, !, ?, |) -> следующее слово будет новым
			newWord = true
		}
	}

	return string(runes)
}
