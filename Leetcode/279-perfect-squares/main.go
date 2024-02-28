package main

import "math"

func numSquares(n int) int {
	dp := make([]int, n+1)
	dp[0] = 0
	for i := 1; i <= n; i++ {
		dp[i] = math.MaxInt
	}
	for i := 1; i <= 100; i++ {
		weight := i * i
		for j := weight; j <= n; j++ {
			if dp[j-weight] != math.MaxInt {
				dp[j] = min(dp[j], dp[j-weight]+1)
			}
		}
	}
	return dp[n]
}
