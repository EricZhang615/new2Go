package main

func canJump(nums []int) bool {
	tmp := 1
	for i := len(nums) - 2; i >= 0; i-- {
		if nums[i] < tmp {
			tmp++
		} else {
			tmp = 1
		}
	}
	return tmp == 1
}

func canJump2(nums []int) bool {
	cover := 0
	n := len(nums) - 1
	for i := 0; i <= cover; i++ { // 每次与覆盖值比较
		cover = max(i+nums[i], cover) //每走一步都将 cover 更新为最大值
		if cover >= n {
			return true
		}
	}
	return false
}
