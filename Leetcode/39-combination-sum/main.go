package main

import (
	"fmt"
	"slices"
)

func combinationSum(candidates []int, target int) [][]int {
	var res [][]int
	var path []int
	slices.Sort(candidates)
	var bc func(index int, t int)
	bc = func(index int, t int) {
		if t == 0 {
			res = append(res, append([]int{}, path...))
			return
		}
		if t < 0 {
			return
		}
		for i := index; i < len(candidates) && t-candidates[i] >= 0; i++ {
			path = append(path, candidates[i])
			bc(i, t-candidates[i])
			path = path[:len(path)-1]
		}
	}
	bc(0, target)
	return res
}

func main() {
	fmt.Println(combinationSum([]int{2, 3, 6, 7}, 7))
}
