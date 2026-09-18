package main

import (
	"os"
)

// Наш аналог Atoi для парсинга строк в int64 с проверкой на валидность
func parseNumber(s string) (int64, bool) {
	if len(s) == 0 {
		return 0, false
	}

	sign := int64(1)
	start := 0

	if s[0] == '-' {
		sign = -1
		start = 1
	} else if s[0] == '+' {
		start = 1
	}

	if start == len(s) {
		return 0, false
	}

	var res int64 = 0
	for i := start; i < len(s); i++ {
		if s[i] < '0' || s[i] > '9' {
			return 0, false
		}

		digit := int64(s[i] - '0')

		// ПРАВИЛЬНАЯ ПРОВЕРКА НА ПЕРЕПОЛНЕНИЕ:
		// Проверка на переполнение при умножении/сложении во время парсинга
		if sign == 1 {
			if res > (9223372036854775807-digit)/10 {
				return 0, false
			}
		} else {
			// Хитрый обход для MinInt64 (-9223372036854775808):
			// Сначала проверяем по границе 9223372036854775807.
			// Если digit равен 8 (последняя цифра MinInt64) и res доползет ровно до границы,
			// мы временно разрешаем это, вычитая единицу из проверки.
			limit := int64(9223372036854775807)
			if digit == 8 {
				if res > (limit-7)/10 {
					return 0, false
				}
			} else {
				if res > (limit-digit)/10 {
					return 0, false
				}
			}
		}
		res = res*10 + digit
	}

	return res * sign, true
}

// Печать строки через os.Stdout.Write
func printStr(s string) {
	os.Stdout.Write([]byte(s))
}

// Печать числа int64 без использования fmt
func printInt(n int64) {
	if n == 0 {
		printStr("0\n")
		return
	}

	var sign string
	if n < 0 {
		sign = "-"
		// Обработка крайнего случая MinInt64, так как -MinInt64 вызывает переполнение
		if n == -9223372036854775808 {
			printStr("-9223372036854775808\n")
			return
		}
		n = -n
	}

	var digits []byte
	for n > 0 {
		digits = append([]byte{byte(n%10 + '0')}, digits...)
		n /= 10
	}

	printStr(sign + string(digits) + "\n")
}

func main() {
	// 1. Проверяем количество аргументов (os.Args[0] — имя программы)
	args := os.Args[1:]
	if len(args) != 3 {
		return
	}

	// 2. Парсим первое и второе число
	val1, ok1 := parseNumber(args[0])
	val2, ok2 := parseNumber(args[2])
	if !ok1 || !ok2 {
		return
	}

	operator := args[1]

	// 3. Обработка деления и остатка на 0
	if val2 == 0 {
		if operator == "/" {
			printStr("No division by 0\n")
			return
		}
		if operator == "%" {
			printStr("No modulo by 0\n")
			return
		}
	}

	// 4. Вычисления с проверкой на переполнение (Overflow)
	var result int64

	switch operator {
	case "+":
		if (val2 > 0 && val1 > 9223372036854775807-val2) || (val2 < 0 && val1 < -9223372036854775808-val2) {
			return // Переполнение при сложении
		}
		result = val1 + val2
	case "-":
		if (val2 < 0 && val1 > 9223372036854775807+val2) || (val2 > 0 && val1 < -9223372036854775808+val2) {
			return // Переполнение при вычитании
		}
		result = val1 - val2
	case "*":
		if val1 != 0 && val2 != 0 {
			if val1 > 0 && val2 > 0 && val1 > 9223372036854775807/val2 {
				return
			}
			if val1 > 0 && val2 < 0 && val2 < -9223372036854775808/val1 {
				return
			}
			if val1 < 0 && val2 > 0 && val1 < -9223372036854775808/val2 {
				return
			}
			if val1 < 0 && val2 < 0 && val1 < 9223372036854775807/val2 {
				return
			}
		}
		result = val1 * val2
	case "/":
		result = val1 / val2
	case "%":
		result = val1 % val2
	default:
		return // Невалидный оператор
	}

	// 5. Выводим результат
	printInt(result)
}
