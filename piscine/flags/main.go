package main

import (
	"fmt"
	"os"

	"github.com/01-edu/z01"
)

func help() {
	fmt.Println("--insert\n  -i\n\t This flag inserts the string into the string passed as argument.\n--order\n  -o\n\t This flag will behave like a boolean, if it is called it will order the argument.")
}

func orderFunc(str string) {
	runes := []rune(str)
	n := len(runes)
	for i := 0; i <= n; i++ {
		for j := 0; j < n-i-1; j++ {
			if runes[j] > runes[j+1] {
				runes[j], runes[j+1] = runes[j+1], runes[j]
			}
		}
	}
	for _, r := range runes {
		z01.PrintRune(r)
	}
}

func main() {
	args := os.Args[1:]
	if len(args) == 0 {
		help()
		return
	}
	insertStr := ""
	order := false
	var str string
	foundStr := false
	for i := 0; i < len(args); i++ {
		arg := args[i]
		if arg == "-h" || arg == "--help" {
			help()
			return
		} else if len(arg) >= 3 && arg[:3] == "-i=" {
			insertStr = arg[3:]
		} else if len(arg) >= 9 && arg[:9] == "--insert=" {
			insertStr = arg[9:]
		} else if arg == "--order" || arg == "-o" {
			order = true
		} else {
			str = arg
			foundStr = true
		}
	}
	if !foundStr {
		help()
		return
	}
	result := str + insertStr
	if order {
		orderFunc(result)
	} else {
		for _, r := range result {
			z01.PrintRune(r)
		}
	}
	z01.PrintRune('\n')
}
