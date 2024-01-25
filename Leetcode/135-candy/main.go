package main

import "fmt"

func candy(ratings []int) int {
	candyList := make([]int, len(ratings))
	candyList[0] = 1
	for i := 1; i < len(ratings); i++ {
		if ratings[i] > ratings[i-1] {
			candyList[i] = candyList[i-1] + 1
		} else {
			candyList[i] = 1
		}
	}
	res := candyList[len(candyList)-1]
	for i := len(ratings) - 2; i >= 0; i-- {
		if ratings[i] > ratings[i+1] {
			candyList[i] = max(candyList[i], candyList[i+1]+1)
			res += candyList[i]
		} else {
			res += candyList[i]
		}
	}
	return res
}

func main() {
	fmt.Println(candy([]int{1, 2, 87, 87, 87, 2, 1}))
}
