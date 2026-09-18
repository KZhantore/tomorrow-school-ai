package piscine

func Compact(ptr *[]string) int {
	// Если указатель nil или слайс пустой, возвращаем 0
	if ptr == nil || len(*ptr) == 0 {
		return 0
	}
	// Создаем счетчик для элементов с непустым значением
	count := 0
	// Получаем сам слайс через разыменование
	slice := *ptr
	// Перемещаем все непустые строки в начало слайса
	for i := 0; i < len(slice); i++ {
		if slice[i] != "" {
			slice[count] = slice[i]
			count++
		}
	}
	// Обрезаем исходный слайс до количества валидных элементов
	*ptr = slice[:count]
	// Возвращаем количество заполненных элементов
	return count
}
