package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) < 4 {
		return
	}

	chars := toInt(os.Args[2])
	files := os.Args[3:]
	hasError := false
	alreadyPrinted := false

	for _, filename := range files {
		data, err := os.ReadFile(filename)
		if err != nil {
			fmt.Printf("%s\n", err)
			hasError = true
			alreadyPrinted = true
			continue
		}

		if alreadyPrinted {
			fmt.Print("\n")
		}
		if len(files) > 1 {
			fmt.Printf("==> %s <==\n", filename)
		}

		length := len(data)
		start := 0
		if length > chars {
			start = length - chars
		}
		os.Stdout.Write(data[start:])
		if length > 0 && data[length-1] != '\n' {
			fmt.Print("\n")
		}
		alreadyPrinted = true
	}

	if hasError {
		os.Exit(1)
	}
}

func toInt(str string) int {
	res := 0
	for _, char := range str {
		res = res*10 + int(char-'0')
	}
	return res
}
