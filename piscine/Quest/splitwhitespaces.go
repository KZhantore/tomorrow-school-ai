package piscine

func SplitWhiteSpaces(s string) []string {
	var result []string
	var currentWord string

	// Бежим по строке символ за символом
	for i := 0; i < len(s); i++ {
		char := s[i]

		// Проверяем, является ли символ пробелом, табом или переносом строки
		if char == ' ' || char == '\t' || char == '\n' {
			// Если в текущем слове уже что-то есть, значит слово завершено
			if len(currentWord) > 0 {
				result = append(result, currentWord)
				currentWord = "" // очищаем для следующего слова
			}
		} else {
			// Если это обычный символ, добавляем его к текущему слову
			currentWord += string(char)
		}
	}

	// Помним про последнее слово: если строка закончилась не пробелом,
	// у нас в currentWord останется последнее слово. Добавляем его.
	if len(currentWord) > 0 {
		result = append(result, currentWord)
	}

	return result
}
