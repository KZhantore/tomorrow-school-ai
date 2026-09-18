package piscine

func Split(s, sep string) []string {
	var result []string

	// Если строка пустая, возвращаем пустой слайс
	if len(s) == 0 {
		return result
	}

	// Если разделитель пустой, можно просто вернуть всю строку в слайсе
	if len(sep) == 0 {
		return []string{s}
	}

	start := 0 // Хранит начало текущего слова
	i := 0

	// Бежим по строке пока хватает места для проверки разделителя
	for i <= len(s)-len(sep) {
		// Проверяем, совпадает ли кусок строки с разделителем
		if s[i:i+len(sep)] == sep {
			// Добавляем слово от 'start' до текущего 'i'
			result = append(result, s[start:i])

			// Сдвигаем 'start' за пределы разделителя
			start = i + len(sep)

			// Перепрыгиваем индекс через разделитель
			i += len(sep) - 1
		}
		i++
	}

	// Добавляем хвостик строки, который остался в самом конце
	result = append(result, s[start:])

	return result
}
