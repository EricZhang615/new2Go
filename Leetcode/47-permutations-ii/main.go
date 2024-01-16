package main

import "slices"

func permuteUnique(nums []int) [][]int {
	var res [][]int
	var path []int
	var bc func(used []bool)
	bc = func(used []bool) {
		if len(path) == len(nums) {
			res = append(res, append([]int{}, path...))
			return
		}
		for i := 0; i < len(nums); i++ {
			if i > 0 && nums[i] == nums[i-1] && used[i-1] {
				continue
			}
			if !used[i] {
				path = append(path, nums[i])
				used[i] = true
				bc(used)
				path = path[:len(path)-1]
				used[i] = false
			}
		}
	}
	slices.Sort(nums)
	bc(make([]bool, len(nums)))
	return res
}
