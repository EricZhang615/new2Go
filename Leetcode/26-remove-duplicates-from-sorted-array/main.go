package main

import "fmt"

func removeDuplicates(nums []int) int {
	p1, p2 := 0, 0
	l := len(nums)
	for p2 < l {
		if nums[p2] == nums[p1] {
			p2++
		} else {
			p1++
			nums[p1] = nums[p2]
			p2++
		}
	}
	return p1 + 1
}

func main() {
	nums := []int{1, 1, 2}
	fmt.Println(removeDuplicates(nums))
	fmt.Println(nums)
	nums = []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}
	fmt.Println(removeDuplicates(nums))
	fmt.Println(nums)
}
