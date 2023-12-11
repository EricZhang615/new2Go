package main

import "fmt"

func mySqrt(x int) int {
	low, high := 0, x
	if x == 1 {
		return 1
	}
	for high-low > 1 {
		mid := low + (high-low)/2
		midSqrt := mid * mid
		if midSqrt > x {
			high = mid
		} else if midSqrt < x {
			low = mid
		} else {
			return mid
		}
	}
	return low
}

func main() {
	fmt.Println(mySqrt(8))
}
