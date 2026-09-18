package piscine

import "github.com/01-edu/z01"

func PrintNbrInOrder(n int) {
	if n < 0 {
		n = -n
	}
	if n == 0 {
		z01.PrintRune('0')
		return
	}

	var digits []rune
	for n > 0 {
		digits = append(digits, rune('0'+n%10))
		n /= 10
	}

	// Сортировка пузырьком
	for i := 0; i < len(digits); i++ {
		for j := i + 1; j < len(digits); j++ {
			if digits[i] > digits[j] {
				digits[i], digits[j] = digits[j], digits[i]
			}
		}
	}

	// Вывод без \n в конце
	for _, d := range digits {
		z01.PrintRune(d)
	}
}
