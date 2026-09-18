package main

import "github.com/01-edu/z01"

type point struct {
	x int
	y int
}

func setPoint(ptr *point) {
	ptr.x = 42
	ptr.y = 21
}

func main() {
	points := &point{}
	setPoint(points)

	x1 := '0'
	for i := 0; i < points.x/10; i++ {
		x1++
	}
	x2 := '0'
	for i := 0; i < points.x%10; i++ {
		x2++
	}

	y1 := '0'
	for i := 0; i < points.y/10; i++ {
		y1++
	}
	y2 := '0'
	for i := 0; i < points.y%10; i++ {
		y2++
	}

	for _, r := range "x = AB, y = CD\n" {
		char := r
		if r == 'A' {
			char = x1
		} else if r == 'B' {
			char = x2
		} else if r == 'C' {
			char = y1
		} else if r == 'D' {
			char = y2
		}
		z01.PrintRune(char)
	}
}
