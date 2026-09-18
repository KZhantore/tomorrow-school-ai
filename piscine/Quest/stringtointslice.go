package piscine

func StringToIntSlice(str string) []int {
	// Если строка пустая, возвращаем пустую мапу/слайс
	if len(str) == 0 {
		return nil
	}
	// Сначала переводим строку в слайс рун,
	// Go сам объединит составные байты UTF-8 в правильные символы
	runes := []rune(str)
	// Создаем результирующий слайс нужной длины
	result := make([]int, len(runes))
	// Заполняем числовыми значениями рун
	for i, r := range runes {
		result[i] = int(r)
	}
	return result
}
