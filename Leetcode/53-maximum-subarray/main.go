package main

import "math"

func maxSubArray(nums []int) int {
	res := math.MinInt
	cnt := 0
	for i := 0; i < len(nums); i++ {
		cnt += nums[i]
		res = max(res, cnt)
		if cnt <= 0 {
			cnt = 0
		}
	}
	return res
}
