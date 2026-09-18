package piscine

func Abort(a, b, c, d, e int) int {
	// Помещаем все 5 аргументов в слайс для удобной сортировки
	arr := []int{a, b, c, d, e}
	// Сортируем массив методом пузырька (Bubble Sort)
	for i := 0; i < len(arr); i++ {
		for j := i + 1; j < len(arr); j++ {
			if arr[i] > arr[j] {
				arr[i], arr[j] = arr[j], arr[i]
			}
		}
	}
	// Возвращаем средний элемент (индекс 2 для массива из 5 элементов)
	return arr[2]
}
