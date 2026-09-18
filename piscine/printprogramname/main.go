package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	// os.Args[0] contains the path/name used to invoke the program
	args := os.Args
	if len(args) == 0 {
		return
	}

	programPath := args[0]
	nameStart := 0

	// Find the start of the actual file name by looking for the last path separator
	for i := len(programPath) - 1; i >= 0; i-- {
		if programPath[i] == '/' || programPath[i] == '\\' {
			nameStart = i + 1
			break
		}
	}

	// Extract the program name from the path
	programName := programPath[nameStart:]

	// Print the program name rune by rune
	for _, r := range programName {
		z01.PrintRune(r)
	}

	// Print a newline at the end
	z01.PrintRune('\n')
}
