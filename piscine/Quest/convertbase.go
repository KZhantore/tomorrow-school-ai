package piscine

// Функция находит индекс символа в строке базы
func getIndex(char rune, base string) int {
	for i, c := range base {
		if c == char {
			return i
		}
	}
	return -1
}

func ConvertBase(nbr, baseFrom, baseTo string) string {
	// --- Шаг 1: Переводим nbr из baseFrom в десятичное число int ---
	lenFrom := len(baseFrom)
	decimalValue := 0

	for _, char := range nbr {
		index := getIndex(char, baseFrom)
		decimalValue = decimalValue*lenFrom + index
	}

	// Крайний случай: если число ноль
	if decimalValue == 0 {
		return string(baseTo[0])
	}

	// --- Шаг 2: Переводим decimalValue из int в систему baseTo ---
	lenTo := len(baseTo)
	var resultBytes []byte

	for decimalValue > 0 {
		remainder := decimalValue % lenTo
		resultBytes = append(resultBytes, baseTo[remainder])
		decimalValue /= lenTo
	}

	// Переворачиваем срез байт, так как при делении остатки собираются с конца
	for i, j := 0, len(resultBytes)-1; i < j; i, j = i+1, j-1 {
		resultBytes[i], resultBytes[j] = resultBytes[j], resultBytes[i]
	}

	return string(resultBytes)
}
