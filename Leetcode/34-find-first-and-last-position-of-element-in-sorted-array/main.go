package main

import "fmt"

func searchRange(nums []int, target int) []int {
	left, right := 0, len(nums)
	for left < right {
		mid := left + (right-left)/2
		if nums[mid] == target {
			left, right = mid, mid
			for left > 0 {
				if nums[left-1] == target {
					left--
				} else {
					break
				}
			}
			for right < len(nums)-1 {
				if nums[right+1] == target {
					right++
				} else {
					break
				}
			}
			return []int{left, right}
		} else if nums[mid] > target {
			right = mid
		} else {
			left = mid + 1
		}
	}
	return []int{-1, -1}
}

func searchRange2(nums []int, target int) []int {
	left := getLeft(nums, target)
	right := getRight(nums, target)
	// target 比所有数都小/大
	if left == -2 || right == -2 {
		return []int{-1, -1}
	}
	// 找到了
	if right-left >= 1 {
		return []int{left, right - 1}
	}
	// 没找到
	return []int{-1, -1}
}

func getRight(nums []int, target int) int {
	left, right := 0, len(nums)
	rightBorder := -2
	for left < right {
		mid := left + (right-left)/2
		if nums[mid] > target {
			right = mid
		} else {
			left = mid + 1
			rightBorder = left
		}
	}
	return rightBorder
}

func getLeft(nums []int, target int) int {
	left, right := 0, len(nums)
	leftBorder := -2
	for left < right {
		mid := left + (right-left)/2
		if nums[mid] >= target {
			right = mid
			leftBorder = right
		} else {
			left = mid + 1
		}
	}
	return leftBorder
}

func main() {
	nums := []int{5, 7, 7, 8, 8, 10}
	target := 6
	fmt.Println(searchRange2(nums, target))
	nums = []int{1}
	target = 1
	fmt.Println(searchRange2(nums, target))
}
