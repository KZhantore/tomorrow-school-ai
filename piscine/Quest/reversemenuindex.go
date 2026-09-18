package piscine

func ReverseMenuIndex(menu []string) []string {
	n := len(menu)
	// Если меню пустое, сразу возвращаем его же
	if n == 0 {
		return menu
	}
	// Создаем новый слайс точно такой же длины без использования append
	reversed := make([]string, n)
	// Переносим элементы в обратном порядке
	for i := 0; i < n; i++ {
		reversed[i] = menu[n-1-i]
	}
	return reversed
}
