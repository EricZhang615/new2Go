package main

func findSubsequences(nums []int) [][]int {
	var res [][]int
	var path []int
	var bc func(index int)
	bc = func(index int) {
		set := make(map[int]bool)
		for i := index; i < len(nums); i++ {
			if (len(path) > 0 && nums[i] < path[len(path)-1]) || set[nums[i]] {
				continue
			}
			path = append(path, nums[i])
			set[nums[i]] = true
			if len(path) >= 2 {
				res = append(res, append([]int{}, path...))
			}
			bc(i + 1)
			path = path[:len(path)-1]
		}
	}
	bc(0)
	return res
}

func findSubsequences2(nums []int) [][]int {
	var res [][]int
	var path []int
	var bc func(index int)
	bc = func(index int) {
		used := make([]bool, 201)
		for i := index; i < len(nums); i++ {
			if (len(path) > 0 && nums[i] < path[len(path)-1]) || used[nums[i]+100] {
				continue
			}
			path = append(path, nums[i])
			used[nums[i]+100] = true
			if len(path) >= 2 {
				res = append(res, append([]int{}, path...))
			}
			bc(i + 1)
			path = path[:len(path)-1]
		}
	}
	bc(0)
	return res
}
