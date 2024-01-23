package main

import "fmt"

func alternatingSubarray(nums []int) int {
	left := 0
	right := 1
	res := -1
	for right < len(nums) {
		if right-left == 1 {
			if nums[right]-nums[left] != 1 {
				left++
			} else {
				res = max(2, res)
			}
			right++
		} else {
			if nums[right]-nums[left] == (right-left)%2 {
				res = max(right-left+1, res)
				right++
			} else {
				left = right - 1
			}
		}
	}
	return res
}

func main() {
	fmt.Println(alternatingSubarray([]int{2, 3, 4, 3, 4}))
}
