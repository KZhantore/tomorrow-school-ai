package piscine

import "github.com/01-edu/z01"

func PrintWordsTables(a []string) {
	// 1. Перебираем каждое слово в слайсе
	for i := 0; i < len(a); i++ {
		word := a[i]

		// 2. Печатаем текущее слово посимвольно
		for j := 0; j < len(word); j++ {
			z01.PrintRune(rune(word[j]))
		}

		// 3. После того как слово напечатано, делаем перенос строки
		z01.PrintRune('\n')
	}
}
