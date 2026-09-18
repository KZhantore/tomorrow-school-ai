package piscine

// BasicAtoi преобразует строку из цифр в целое число (int).
func BasicAtoi(s string) int {
	result := 0

	for i := 0; i < len(s); i++ {
		// Переводим символ в цифру, вычитая ASCII-код нуля
		digit := int(s[i] - '0')

		// Формируем число разряд за разрядом
		result = result*10 + digit
	}

	return result
}
