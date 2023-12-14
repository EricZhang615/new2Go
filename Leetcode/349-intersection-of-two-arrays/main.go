package main

import "fmt"

func intersection(nums1 []int, nums2 []int) []int {
	a := make(map[int]int)
	var ans []int
	for _, i := range nums1 {
		a[i] = 1
	}
	for _, i := range nums2 {
		if a[i] == 1 {
			ans = append(ans, i)
			a[i] = 2
		}
	}
	return ans
}

func main() {
	n1, n2 := []int{1, 2, 2, 1}, []int{2, 2}
	fmt.Println(intersection(n1, n2))
}
