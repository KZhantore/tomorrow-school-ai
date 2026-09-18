package piscine

func AppendRange(min, max int) []int {
	// 1. Проверяем условие, если min больше или равен max
	if min >= max {
		return nil
	}

	// 2. Создаем пустой слайс (так как make использовать нельзя)
	var result []int

	// 3. Запускаем цикл от min до max (исключая max)
	for i := min; i < max; i++ {
		result = append(result, i) // добавляем число в слайс
	}

	// 4. Возвращаем готовый слайс
	return result
}
