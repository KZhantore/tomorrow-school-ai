package piscine

import "github.com/01-edu/z01"

func EightQueens() {
	// Массив, где индекс - это колонка, а значение - строка (0-7)
	var board [8]int
	solve(&board, 0)
}

func solve(board *[8]int, col int) {
	if col == 8 {
		// Все ферзи расставлены, выводим результат
		printBoard(board)
		return
	}

	for row := 0; row < 8; row++ {
		if isSafe(board, col, row) {
			board[col] = row
			solve(board, col+1)
		}
	}
}

func isSafe(board *[8]int, col, row int) bool {
	for i := 0; i < col; i++ {
		prevRow := board[i]
		// 1. Проверка горизонтали
		if prevRow == row {
			return false
		}
		// 2. Проверка диагоналей
		// Разница строк равна разнице колонок
		diff := col - i
		if prevRow == row-diff || prevRow == row+diff {
			return false
		}
	}
	return true
}

func printBoard(board *[8]int) {
	for i := 0; i < 8; i++ {
		// +1 потому что по условию индекс строк начинается с 1
		// + '0' преобразует цифру в символ
		z01.PrintRune(rune(board[i] + 1 + '0'))
	}
	z01.PrintRune('\n')
}
