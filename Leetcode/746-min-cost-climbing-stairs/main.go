package main

func minCostClimbingStairs(cost []int) int {
	dp := []int{0, 0}
	for i := 2; i <= len(cost); i++ {
		minCost := min(dp[1]+cost[i-1], dp[0]+cost[i-2])
		dp[0] = dp[1]
		dp[1] = minCost
	}
	return dp[1]
}
