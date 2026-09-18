package piscine

func IterativeFactorial(nb int) int {
	// 1. Обработка крайних случаев
	if nb < 0 || nb > 20 { // 20! — это предел для int64, а для int может быть и меньше
		return 0
	}
	if nb == 0 {
		return 1
	}

	res := 1
	// 2. Четкое условие завершения
	for i := 1; i <= nb; i++ {
		res *= i
	}

	return res
}
