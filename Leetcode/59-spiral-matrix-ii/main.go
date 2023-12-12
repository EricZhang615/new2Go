package main

import "fmt"

func generateMatrix(n int) [][]int {
	m := make([][]int, 0)
	for i := 0; i < n; i++ {
		m = append(m, make([]int, n))
	}
	round := n / 2
	mid := n % 2
	cnt := 1
	for i := 0; i < round; i++ {
		for j := i; j < n-i-1; j++ {
			m[i][j] = cnt
			cnt++
		}
		for j := i; j < n-i-1; j++ {
			m[j][n-i-1] = cnt
			cnt++
		}
		for j := n - i - 1; j > i; j-- {
			m[n-i-1][j] = cnt
			cnt++
		}
		for j := n - i - 1; j > i; j-- {
			m[j][i] = cnt
			cnt++
		}
	}
	if mid == 1 {
		m[n/2][n/2] = cnt
	}
	return m
}

func main() {
	fmt.Println(generateMatrix(3))
}
