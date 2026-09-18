package piscine

func IterativePower(nb int, power int) int {
	// 1. Если степень отрицательная, возвращаем 0
	if power < 0 {
		return 0
	}

	// 2. Базовый случай: любое число в степени 0 равно 1
	res := 1

	// 3. Умножаем nb само на себя power раз
	for i := 0; i < power; i++ {
		res *= nb
	}

	return res
}
