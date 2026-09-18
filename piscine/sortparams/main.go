package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	args := os.Args
	if len(args) <= 1 {
		return
	}

	subArgs := args[1:]

	for i := 0; i < len(subArgs); i++ {
		for j := 0; j < len(subArgs)-1; j++ {
			if subArgs[j] > subArgs[j+1] {
				subArgs[j], subArgs[j+1] = subArgs[j+1], subArgs[j]
			}
		}
	}
	for i := 0; i < len(subArgs); i++ {
		for _, letter := range subArgs[i] {
			z01.PrintRune(letter)
		}
		z01.PrintRune('\n')
	}
}
