package piscine

func ShoppingListSort(slice []string) []string {
	n := len(slice)
	// Если слайс пустой или содержит всего 1 элемент, он уже отсортирован
	if n <= 1 {
		return slice
	}
	// Алгоритм сортировки пузырьком по длине строк
	for i := 0; i < n; i++ {
		for j := 0; j < n-i-1; j++ {
			// Сравниваем длины соседних строк
			if len(slice[j]) > len(slice[j+1]) {
				// Меняем элементы местами
				slice[j], slice[j+1] = slice[j+1], slice[j]
			}
		}
	}
	return slice
}
