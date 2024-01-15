package main

func subsets(nums []int) [][]int {
	var res [][]int
	res = append(res, []int{})
	var path []int
	var bc func(index int)
	bc = func(index int) {
		for i := index; i < len(nums); i++ {
			path = append(path, nums[i])
			res = append(res, append([]int{}, path...))
			bc(i + 1)
			path = path[:len(path)-1]
		}
	}
	bc(0)
	return res
}
