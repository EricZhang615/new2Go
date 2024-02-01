package main

func climbStairs(n int) int {
	if n <= 2 {
		return n
	}
	dp := []int{0, 1, 2}
	for i := 3; i <= n; i++ {
		step := dp[2] + dp[1]
		dp[1] = dp[2]
		dp[2] = step
	}
	return dp[2]
}
