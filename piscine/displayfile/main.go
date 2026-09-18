package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	args := os.Args

	if len(args) < 2 {
		fmt.Println("File name missing")
		return
	}

	if len(args) > 2 {
		fmt.Println("Too many arguments")
		return
	}

	file, err := os.Open(args[1])
	if err != nil {
		return
	}
	defer file.Close()

	_, err = io.Copy(os.Stdout, file)
	if err != nil {
		return
	}
}
