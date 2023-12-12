package main

import "fmt"

func spiralOrder(matrix [][]int) []int {
	m := len(matrix)
	n := len(matrix[0])
	var round, mid int
	if m > n {
		round = n / 2
		mid = n % 2
	} else {
		round = m / 2
		mid = m % 2
	}
	var ans []int
	for i := 0; i < round; i++ {
		for j := i; j < n-i-1; j++ {
			ans = append(ans, matrix[i][j])
		}
		for j := i; j < m-i-1; j++ {
			ans = append(ans, matrix[j][n-i-1])
		}
		for j := n - i - 1; j > i; j-- {
			ans = append(ans, matrix[m-i-1][j])
		}
		for j := m - i - 1; j > i; j-- {
			ans = append(ans, matrix[j][i])
		}
	}
	if mid == 1 {
		if m > n {
			for i := round; i < m-round; i++ {
				ans = append(ans, matrix[i][round])
			}
		} else {
			for i := round; i < n-round; i++ {
				ans = append(ans, matrix[round][i])
			}
		}
	}
	return ans
}

func main() {
	matrix := [][]int{{1, 2, 3, 4}, {5, 6, 7, 8}, {9, 10, 11, 12}}
	fmt.Println(spiralOrder(matrix))
	matrix = [][]int{{1, 2, 3}, {5, 6, 7}, {9, 10, 11}}
	fmt.Println(spiralOrder(matrix))
}
