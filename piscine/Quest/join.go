package piscine

func Join(strs []string, sep string) string {
	// Если слайс пустой, возвращаем пустую строку
	if len(strs) == 0 {
		return ""
	}
	// Инициализируем результат первой строкой из слайса
	result := strs[0]
	// Начинаем цикл со второго элемента (индекс 1)
	for i := 1; i < len(strs); i++ {
		// Перед каждым следующим элементом добавляем разделитель
		result += sep + strs[i]
	}
	return result
}
