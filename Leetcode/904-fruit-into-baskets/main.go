package main

import "fmt"

func totalFruit(fruits []int) (ans int) {
	cnt := map[int]int{}
	left := 0
	for right, x := range fruits {
		cnt[x]++
		for len(cnt) > 2 {
			y := fruits[left]
			cnt[y]--
			if cnt[y] == 0 {
				delete(cnt, y)
			}
			left++
		}
		ans = max(ans, right-left+1)
	}
	return
}

func max(a, b int) int {
	if b > a {
		return b
	}
	return a
}

func main() {
	fruits := []int{3, 3, 3, 1, 2, 1, 1, 2, 3, 3, 4}
	fmt.Println(totalFruit(fruits))
	fruits = []int{1, 2, 3, 2, 2}
	fmt.Println(totalFruit(fruits))
	fruits = []int{4, 7, 7, 0, 8, 3, 8, 2, 5}
	fmt.Println(totalFruit(fruits))
}
