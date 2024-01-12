package main

import "fmt"

func combinationSum3(k int, n int) [][]int {
	var res [][]int
	var path []int
	var bc func(k int, n int, index int)
	bc = func(k int, n int, index int) {
		if len(path) == k {
			if n == 0 {
				res = append(res, append([]int{}, path...))
			}
			return
		}
		if n <= 0 {
			return
		}
		for i := index; i <= 9-(k-len(path))+1; i++ {
			path = append(path, i)
			bc(k, n-i, i+1)
			path = path[:len(path)-1]
		}
	}
	bc(k, n, 1)
	return res
}

func main() {
	fmt.Println(combinationSum3(3, 9))
	fmt.Println(combinationSum3(3, 9))
}
