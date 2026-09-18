package piscine

func Sqrt(nb int) int {
	if nb <= 0 {
		return 0
	}
	if nb == 1 {
		return 1
	}

	// Итерируемся, пока i*i не превысит nb
	// Мы используем i <= nb/i вместо i*i <= nb, чтобы избежать переполнения (overflow)
	for i := 1; i <= nb/i; i++ {
		if i*i == nb {
			return i
		}
	}

	return 0
}
