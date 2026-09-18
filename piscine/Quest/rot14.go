package piscine

func Rot14(s string) string {
	// Преобразуем строку в слайс рун для возможности изменения символов
	result := []rune(s)
	for i, r := range result {
		// Обработка заглавных букв
		if r >= 'A' && r <= 'Z' {
			result[i] = 'A' + (r-'A'+14)%26
		} // Важно: 'else if' должен быть на той же строке, что и '}'
		if r >= 'a' && r <= 'z' { // Можно использовать обычный if, это безопаснее для Go
			result[i] = 'a' + (r-'a'+14)%26
		}
	}
	return string(result)
}
