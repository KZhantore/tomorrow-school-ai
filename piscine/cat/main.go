package main

import (
	"io"
	"os"

	"github.com/01-edu/z01"
)

func printStr(s string) {
	for _, r := range s {
		z01.PrintRune(r)
	}
}

func main() {
	args := os.Args[1:]

	if len(args) == 0 {
		_, err := io.Copy(os.Stdout, os.Stdin)
		if err != nil {
			return
		}
		return
	}

	for _, filename := range args {
		file, err := os.Open(filename)
		if err != nil {
			printStr("ERROR: open " + filename + ": no such file or directory\n")
			os.Exit(1)
		}

		_, err = io.Copy(os.Stdout, file)
		file.Close()
		if err != nil {
			os.Exit(1)
		}
	}
}
