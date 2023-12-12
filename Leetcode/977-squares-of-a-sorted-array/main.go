package main

import "fmt"

func sortedSquares(nums []int) []int {
	left, right, k := 0, len(nums)-1, len(nums)-1
	s := make([]int, len(nums))
	for left < right {
		if nums[left]*nums[left] > nums[right]*nums[right] {
			s[k] = nums[left] * nums[left]
			left++
		} else {
			s[k] = nums[right] * nums[right]
			right--
		}
		k--
	}
	s[k] = nums[left] * nums[left]
	return s
}

func main() {
	nums := []int{-7, -3, 2, 3, 11}
	fmt.Println(sortedSquares(nums))
}
