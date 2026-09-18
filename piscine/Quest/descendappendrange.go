package piscine

func DescendAppendRange(max, min int) []int {
	// Если max меньше или равен min, возвращаем пустой слайс
	if max <= min {
		return []int{}
	}
	// Инициализируем пустой слайс без использования make()
	var result []int
	// Цикл идет от max вниз до min (не включая min)
	for i := max; i > min; i-- {
		result = append(result, i)
	}
	return result
}
