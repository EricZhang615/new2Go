package main

func wiggleMaxLength(nums []int) int {
	res := 1
	diff := 0
	num := nums[0]
	for i := 1; i < len(nums); i++ {
		if diff == 0 {
			if nums[i]-nums[i-1] == 0 {
				continue
			}
		} else if diff > 0 {
			if nums[i]-num >= 0 {
				num = nums[i]
				continue
			}
		} else {
			if nums[i]-num <= 0 {
				num = nums[i]
				continue
			}
		}
		res++
		diff = nums[i] - num
		num = nums[i]
	}
	return res
}
