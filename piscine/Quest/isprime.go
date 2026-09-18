package piscine

func IsPrime(nb int) bool {
	// Числа меньше или равные 1 не являются простыми
	if nb <= 1 {
		return false
	}
	// 2 и 3 — простые числа
	if nb <= 3 {
		return true
	}
	// Если число делится на 2 или 3, оно не простое
	if nb%2 == 0 || nb%3 == 0 {
		return false
	}

	// Проверяем делители от 5, пропуская четные
	// Используем i*i <= nb как эффективный предел
	for i := 5; i*i <= nb; i = i + 6 {
		if nb%i == 0 || nb%(i+2) == 0 {
			return false
		}
	}

	return true
}
