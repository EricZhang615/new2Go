package main

func findTargetSumWays(nums []int, target int) int {
	sum := 0
	for _, num := range nums {
		sum += num
	}
	if sum < abs(target) {
		return 0
	}
	if (sum+target)%2 == 1 {
		return 0
	}
	bagSize := (sum + target) / 2
	dp := make([]int, bagSize+1)
	dp[0] = 1
	for i := 0; i < len(nums); i++ {
		for j := bagSize; j >= nums[i]; j-- {
			dp[j] += dp[j-nums[i]]
		}
	}
	return dp[bagSize]
}

func abs(n int) int {
	if n >= 0 {
		return n
	} else {
		return -n
	}
}
