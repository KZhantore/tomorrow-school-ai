package piscine

func MakeRange(min, max int) []int {
	// 1. Проверяем базовое условие
	if min >= max {
		return nil
	}

	// 2. Вычисляем размер и выделяем память с помощью make
	length := max - min
	result := make([]int, length)

	// 3. Заполняем слайс с помощью цикла
	// Переменная i считает индекс слайса (от 0), а min увеличивается на каждом шаге
	for i := 0; i < length; i++ {
		result[i] = min + i
	}

	// 4. Возвращаем готовый результат
	return result
}
