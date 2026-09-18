package piscine

func RockAndRoll(n int) string {
	// 1. Проверяем на отрицательное число
	if n < 0 {
		return "error: number is negative\n"
	}
	// 2. Проверяем на делимость и на 2, и на 3 одновременно (то есть на 6)
	if n%2 == 0 && n%3 == 0 {
		return "rock and roll\n"
	}
	// 3. Проверяем только на 2
	if n%2 == 0 {
		return "rock\n"
	}
	// 4. Проверяем только на 3
	if n%3 == 0 {
		return "roll\n"
	}
	// 5. Если ни одно условие не подошло — число не делится ни на 2, ни на 3
	return "error: non divisible\n"
}
