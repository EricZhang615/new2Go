package main

func maxProfit(prices []int) int {
	buy := -1
	res := 0
	for i := 1; i < len(prices); i++ {
		if buy == -1 {
			if prices[i] > prices[i-1] {
				buy = prices[i-1]
			}
		} else {
			if prices[i] < prices[i-1] {
				res += prices[i-1] - buy
				buy = -1
			}
		}
	}
	if buy != -1 {
		res += prices[len(prices)-1] - buy
	}
	return res
}

func maxProfit2(prices []int) int {
	res := 0
	for i := 1; i < len(prices); i++ {
		res += max(prices[i]-prices[i-1], 0)
	}
	return res
}
