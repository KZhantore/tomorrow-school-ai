package piscine

func AdvancedSortWordArr(a []string, f func(a, b string) int) {
	n := len(a)
	// Внешний цикл проходов
	for i := 0; i < n-1; i++ {
		// Внутренний цикл сравнения соседних строк
		for j := 0; j < n-i-1; j++ {
			// Спрашиваем у функции f, идет ли a[j] после a[j+1]
			if f(a[j], a[j+1]) > 0 {
				// Если да, меняем их местами
				a[j], a[j+1] = a[j+1], a[j]
			}
		}
	}
}
