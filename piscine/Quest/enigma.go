package piscine

func Enigma(a ***int, b *int, c *******int, d ****int) {
	// Сохраняем исходные значения типа int во временные переменные
	tempA := ***a
	tempB := *b
	tempC := *******c
	tempD := ****d
	// Перемещаем значения согласно инструкции:
	// a into c (значение из a кладем в c)
	*******c = tempA
	// c into d (значение из c кладем в d)
	****d = tempC
	// d into b (значение из d кладем в b)
	*b = tempD
	// b into a (значение из b кладем в a)
	***a = tempB
}
