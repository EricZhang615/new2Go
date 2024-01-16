package main

func permute(nums []int) [][]int {
	var res [][]int
	var path []int
	var bc func(used []bool)
	bc = func(used []bool) {
		if len(path) == len(nums) {
			res = append(res, append([]int{}, path...))
			return
		}
		for i := 0; i < len(nums); i++ {
			if used[nums[i]+10] {
				continue
			}
			path = append(path, nums[i])
			used[nums[i]+10] = true
			bc(used)
			path = path[:len(path)-1]
			used[nums[i]+10] = false
		}
	}
	used := make([]bool, 21)
	bc(used)
	return res
}
