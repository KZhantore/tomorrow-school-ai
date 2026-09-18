package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	args := os.Args
	for i := 1; i < len(args); i++ {
		for _, letter := range args[i] {
			z01.PrintRune(letter)
		}
		z01.PrintRune('\n')
	}
}
