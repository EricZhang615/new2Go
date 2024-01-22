package main

import (
	"slices"
)

func solveSudoku(board [][]byte) {
	var bc func(row, col int) bool
	var isValid func(row, col, digit int) bool
	isValid = func(row, col, digit int) bool {
		for i, num := range board[row] {
			if i != col && num == byte('0'+digit) {
				return false
			}
		}
		for i, numList := range board {
			if i != row && numList[col] == byte('0'+digit) {
				return false
			}
			if row/3*3 == i/3*3 {
				if slices.Contains(numList[col/3*3:col/3*3+3], byte('0'+digit)) {
					return false
				}
			}
		}
		return true
	}
	bc = func(row, col int) bool {
		if row == 9 {
			return true
		}
		if board[row][col] == '.' {
			for i := 1; i < 10; i++ {
				if isValid(row, col, i) {
					board[row][col] = byte('0' + i)
					var res bool
					if col == 8 {
						res = bc(row+1, 0)
					} else {
						res = bc(row, col+1)
					}
					if res {
						return true
					}
					board[row][col] = '.'
				}
			}
			return false
		}
		if col == 8 {
			return bc(row+1, 0)
		} else {
			return bc(row, col+1)
		}
	}
	bc(0, 0)
}
