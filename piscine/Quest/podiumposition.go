package piscine

func PodiumPosition(podium [][]string) [][]string {
	// Если подиум пустой, возвращаем его как есть
	if len(podium) == 0 {
		return podium
	}
	// Разворачиваем слайс на месте, меняя элементы с начала и конца местами
	for i, j := 0, len(podium)-1; i < j; i, j = i+1, j-1 {
		podium[i], podium[j] = podium[j], podium[i]
	}
	return podium
}
