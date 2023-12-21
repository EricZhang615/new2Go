package main

import "fmt"

/*
func maxSlidingWindow(nums []int, k int) []int {
	queue := make([]int, 0)
	m := -10000
	i, j := -k, 0
	ans := make([]int, len(nums)-k+1)
	for j < len(nums) {
		if len(queue) == 0 || nums[j] >= m {
			queue = []int{nums[j]}
			m = nums[j]
		} else {
			queue = append(queue, nums[j])
		}
		if len(queue) >= k+1 && queue[0] == m {
			x := 1
			for y := x; y < len(queue); y++ {
				if queue[y] >= queue[x] {
					x = y
				}
			}
			m = queue[x]
			queue = queue[x:]
		}
		i++
		j++
		if i >= 0 {
			ans[i] = m
		}
	}
	return ans
}
*/

func maxSlidingWindow(nums []int, k int) []int {
	queue := make([]int, 0)
	i, j := -k, 0
	ans := make([]int, len(nums)-k+1)
	for j < len(nums) {
		for n := len(queue) - 1; n >= 0; n-- {
			if queue[n] < nums[j] {
				queue = queue[:len(queue)-1]
			} else {
				break
			}
		}
		queue = append(queue, nums[j])
		if i >= 0 && nums[i] == queue[0] {
			queue = queue[1:]
		}
		i++
		j++
		if i >= 0 {
			ans[i] = queue[0]
		}
	}
	return ans
}

func main() {
	fmt.Println(maxSlidingWindow([]int{1, 3, -1, -3, 5, 3, 6, 7}, 3))
	fmt.Println(maxSlidingWindow([]int{1}, 1))
	fmt.Println(maxSlidingWindow([]int{-6, -10, -7, -1, -9, 9, -8, -4, 10, -5, 2, 9, 0, -7, 7, 4, -2, -10, 8, 7}, 7))
}
