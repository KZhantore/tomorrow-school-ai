package piscine

// SortIntegerTable сортирует срез целых чисел по возрастанию.
func SortIntegerTable(table []int) {
	n := len(table)

	// Внешний цикл отвечает за количество проходов по массиву
	for i := 0; i < n-1; i++ {
		// Внутренний цикл сравнивает соседние элементы
		for j := 0; j < n-i-1; j++ {
			// Если левый элемент больше правого, меняем их местами
			if table[j] > table[j+1] {
				table[j], table[j+1] = table[j+1], table[j]
			}
		}
	}
}
