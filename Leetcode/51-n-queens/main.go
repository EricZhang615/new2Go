package main

import "fmt"

func solveNQueens(n int) [][]string {
	var res [][]string
	var path []string
	var bc func()
	bc = func() {
		if len(path) == n {
			res = append(res, append([]string{}, path...))
			return
		}
		for i := 0; i < n; i++ {
			for j, col := range path {
				if col[i] == 'Q' {
					continue
				}
				rowLeft := i - len(path) + j
				rowRight := rowLeft + (len(path)-j)*2
				if (rowLeft >= 0 && col[rowLeft] == 'Q') || (rowRight <= n-1 && col[rowRight] == 'Q') {
					continue
				}
			}
			str := ""
			for j := 0; j < i; j++ {
				str += "."
			}
			str += "Q"
			for j := i + 1; j < n; j++ {
				str += "."
			}
			path = append(path, str)
			bc()
			path = path[:len(path)-1]
		}
		return
	}
	bc()
	return res
}

func main() {
	fmt.Println(solveNQueens(4))
}
