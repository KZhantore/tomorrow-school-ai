package piscine

import "github.com/01-edu/z01"

func PrintNbrBase(nbr int, base string) {
	bRunes := []rune(base)
	bLen := len(bRunes)

	// 1. Валидация базы
	if bLen < 2 {
		z01.PrintRune('N')
		z01.PrintRune('V')
		return
	}

	for i := 0; i < bLen; i++ {
		if bRunes[i] == '+' || bRunes[i] == '-' {
			z01.PrintRune('N')
			z01.PrintRune('V')
			return
		}
		for j := i + 1; j < bLen; j++ {
			if bRunes[i] == bRunes[j] {
				z01.PrintRune('N')
				z01.PrintRune('V')
				return
			}
		}
	}

	// 2. Обработка нуля
	if nbr == 0 {
		z01.PrintRune(bRunes[0])
		return
	}

	// 3. Обработка знака (печатаем минус, если число отрицательное)
	if nbr < 0 {
		z01.PrintRune('-')
	}

	// 4. Перевод системы счисления (работаем с n как есть, не меняя знак)
	var result []rune
	n := nbr

	for n != 0 {
		remainder := n % bLen
		// Если остаток отрицательный, делаем его положительным индексом
		if remainder < 0 {
			remainder = -remainder
		}
		result = append(result, bRunes[remainder])
		n /= bLen
	}

	// 5. Вывод результата задом наперед
	for i := len(result) - 1; i >= 0; i-- {
		z01.PrintRune(result[i])
	}
}
