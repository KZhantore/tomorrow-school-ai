package piscine

func IsSorted(f func(a, b int) int, a []int) bool {
	ascending := true
	descending := true

	// Крутим цикл до len(a) - 1, чтобы не выйти за границы слайса, проверяя a[i+1]
	for i := 0; i < len(a)-1; i++ {
		// Если текущий элемент больше следующего, это нарушает сортировку по возрастанию
		if f(a[i], a[i+1]) > 0 {
			ascending = false
		}
		// Если текущий элемент меньше следующего, это нарушает сортировку по убыванию
		if f(a[i], a[i+1]) < 0 {
			descending = false
		}
	}

	// Если слайс остался отсортированным ХОТЯ БЫ по одному из направлений, возвращаем true
	return ascending || descending
}
