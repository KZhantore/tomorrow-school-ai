package main

import (
	"os"

	"github.com/01-edu/z01"
)

// Наша собственная функция для перевода строки в число.
// Возвращает число и статус успешности (true/false).
func atoi(s string) (int, bool) {
	if len(s) == 0 {
		return 0, false
	}
	res := 0
	for i := 0; i < len(s); i++ {
		if s[i] < '0' || s[i] > '9' {
			return 0, false
		}
		res = res*10 + int(s[i]-'0')
	}
	return res, true
}

func main() {
	args := os.Args[1:]
	if len(args) == 0 {
		return
	}

	isUpper := false
	if args[0] == "--upper" {
		isUpper = true
		args = args[1:]
	}

	for _, arg := range args {
		n, success := atoi(arg)

		// Если это не число или оно выходит за рамки алфавита
		if !success || n < 1 || n > 26 {
			z01.PrintRune(' ')
			continue
		}

		var letter rune
		if isUpper {
			letter = rune('A' + (n - 1))
		} else {
			letter = rune('a' + (n - 1))
		}

		z01.PrintRune(letter)
	}

	z01.PrintRune('\n')
}
