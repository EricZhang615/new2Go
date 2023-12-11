package main

import "fmt"

func isPerfectSquare(num int) bool {
	low, high := 0, num
	if num == 1 {
		return true
	}
	for high-low > 1 {
		mid := low + (high-low)/2
		midSqrt := mid * mid
		if midSqrt > num {
			high = mid
		} else if midSqrt < num {
			low = mid
		} else {
			return true
		}
	}
	return false
}

func main() {
	fmt.Println(isPerfectSquare(1))
}
