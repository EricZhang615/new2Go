package main

import "fmt"

func minSubArrayLen(target int, nums []int) int {
	i, j := 0, 0
	ans := len(nums) + 1
	sum := 0
	for j < len(nums) {
		sum += nums[j]
		for sum >= target {
			if i == j {
				return 1
			}
			l := j - i + 1
			if l < ans {
				ans = l
			}
			sum -= nums[i]
			i++
		}
		j++
	}
	if ans == len(nums)+1 {
		return 0
	}
	return ans
}

func main() {
	nums := []int{1, 1, 1, 1, 1, 1, 1, 1}
	target := 11
	fmt.Println(minSubArrayLen(target, nums))
}
