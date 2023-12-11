package main

import "fmt"

func removeElement(nums []int, val int) int {
	p1, p2 := 0, len(nums)-1
	for p1 <= p2 {
		if nums[p1] == val {
			nums[p1] = nums[p2]
			p2--
		} else {
			p1++
		}
	}
	return p1
}

func main() {
	nums := []int{0, 1, 2, 2, 3, 0, 4, 2}
	val := 2
	fmt.Println(removeElement(nums, val))
}
