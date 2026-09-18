package piscine

func CollatzCountdown(start int) int {
	// Если число меньше или равно 0, возвращаем -1 согласно правилам Piscine
	if start <= 0 {
		return -1
	}
	steps := 0
	// Цикл продолжается, пока число не станет равным 1
	for start > 1 {
		if start%2 == 0 {
			start = start / 2
		} else {
			start = start*3 + 1
		}
		steps++ // Увеличиваем счетчик шагов после каждого действия
	}
	return steps
}
