package main

func combine(n int, k int) [][]int {
	var res [][]int
	var path []int
	var bc func(n int, k int, index int)
	bc = func(n int, k int, index int) {
		if len(path) == k {
			res = append(res, append([]int{}, path...))
			return
		}
		for i := index; i <= n; i++ {
			path = append(path, i)
			bc(n, k, i+1)
			path = path[:len(path)-1]
		}
	}
	bc(n, k, 1)
	return res
}
