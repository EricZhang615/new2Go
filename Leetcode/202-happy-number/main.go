package main

import "fmt"

func isHappy(n int) bool {
	square := []int{0, 1, 4, 9, 16, 25, 36, 49, 64, 81}
	m := make(map[int]int)
	i := n
	for i != 1 {
		m[i] = 1
		var num int
		for i != 0 {
			num += square[i%10]
			i /= 10
		}
		if m[num] == 1 {
			return false
		}
		i = num
	}
	return true
}

func main() {
	n := 19
	fmt.Println(isHappy(n))
}
