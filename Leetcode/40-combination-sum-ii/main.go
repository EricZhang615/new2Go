package main

import "slices"

func combinationSum2(candidates []int, target int) [][]int {
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
			if i > index && candidates[i] == candidates[i-1] {
				continue
			}
			path = append(path, candidates[i])
			bc(i+1, t-candidates[i])
			path = path[:len(path)-1]
		}
	}
	bc(0, target)
	return res
}
