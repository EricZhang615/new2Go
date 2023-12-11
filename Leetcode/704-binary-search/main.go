package main

import "fmt"

func search(nums []int, target int) int {
	left := 0
	right := len(nums)
	middle := right / 2
	for {
		if nums[middle] == target {
			return middle
		}
		if left == right-1 {
			return -1
		}
		if nums[middle] > target {
			right = middle
			middle = (right-left)/2 + left
		} else if nums[middle] < target {
			left = middle
			middle = (right-left)/2 + left
		}
	}
}

func search2(nums []int, target int) int {
	high := len(nums)
	low := 0
	for low < high {
		mid := low + (high-low)/2
		if nums[mid] == target {
			return mid
		} else if nums[mid] > target {
			high = mid
		} else {
			low = mid + 1
		}
	}
	return -1
}

func main() {
	nums := []int{-1, 0, 3, 5, 9, 12}
	target := 2
	fmt.Println(search(nums, target))
}
