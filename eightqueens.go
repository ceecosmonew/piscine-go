package piscine

import "github.com/01-edu/z01"

func EightQueens() {
	var board [8]int
	solveQueens(0, &board)
}

func solveQueens(col int, board *[8]int) {
	if col == 8 {
		printBoard(board)
		return
	}

	for row := 0; row < 8; row++ {
		if isSafe(col, row, board) {
			board[col] = row
			solveQueens(col+1, board)
		}
	}
}

func isSafe(col int, row int, board *[8]int) bool {
	for prevCol := 0; prevCol < col; prevCol++ {
		prevRow := board[prevCol]

		// Same row → attack
		if prevRow == row {
			return false
		}
		// Diagonal attack
		if prevRow-prevCol == row-col || prevRow+prevCol == row+col {
			return false
		}
	}
	return true
}

func printBoard(board *[8]int) {
	for col := 0; col < 8; col++ {
		digit := board[col] + 1   // convert 0–7 → 1–8
		r := rune(digit) + '0'    // make it a rune digit
		z01.PrintRune(r)
	}
	z01.PrintRune('\n')
}
