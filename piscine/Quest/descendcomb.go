package piscine

import "github.com/01-edu/z01"

func DescendComb() {
	// Loop through the first two-digit number from 99 down to 01
	for i := 99; i >= 1; i-- {
		// Loop through the second two-digit number, always smaller than i
		for j := i - 1; j >= 0; j-- {

			// Print the first number (i)
			z01.PrintRune(rune('0' + i/10))
			z01.PrintRune(rune('0' + i%10))

			// Print the space between the two numbers
			z01.PrintRune(' ')

			// Print the second number (j)
			z01.PrintRune(rune('0' + j/10))
			z01.PrintRune(rune('0' + j%10))

			// Print the comma and space separator, EXCEPT for the very last combination (01 00)
			if !(i == 1 && j == 0) {
				z01.PrintRune(',')
				z01.PrintRune(' ')
			}
		}
	}
}
