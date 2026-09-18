package piscine

func RecursiveFactorial(nb int) int {
	// 1. Обработка ошибок и ограничений (как в итеративном)
	if nb < 0 || nb > 20 {
		return 0
	}
	// 2. Базовый случай (точка остановки)
	if nb == 0 {
		return 1
	}
	// 3. Рекурсивный шаг
	return nb * RecursiveFactorial(nb-1)
}
