package piscine

import "github.com/01-edu/z01"

func PrintCombN(n int) {
	combination := make([]rune, n)
	isFirst := true // Флаг, чтобы не ставить запятую перед самой первой комбинацией

	// Передаем указатель на флаг, чтобы он менялся внутри рекурсии
	generateCombN(0, '0', n, combination, &isFirst)

	z01.PrintRune('\n')
}

func generateCombN(index int, startDigit rune, n int, combination []rune, isFirst *bool) {
	// Базовый случай: комбинация полностью собрана
	if index == n {
		// Если это НЕ первая комбинация, сначала печатаем разделитель ", "
		if !*isFirst {
			z01.PrintRune(',')
			z01.PrintRune(' ')
		}
		// Переключаем флаг: все последующие комбинации уже точно не первые
		*isFirst = false

		// Выводим саму комбинацию цифр
		for i := 0; i < n; i++ {
			z01.PrintRune(combination[i])
		}
		return
	}

	// Строго возрастающий перебор цифр
	for d := startDigit; d <= '9'; d++ {
		combination[index] = d
		// Следующая цифра должна быть как минимум на 1 больше текущей (d + 1)
		generateCombN(index+1, d+1, n, combination, isFirst)
	}
}
