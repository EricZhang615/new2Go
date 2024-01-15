package main

func isPal(s string) bool {
	left := 0
	right := len(s) - 1
	for left < right {
		if s[left] == s[right] {
			left++
			right--
		} else {
			return false
		}
	}
	return true
}

func partition(s string) [][]string {
	var res [][]string
	var path []string
	var bc func(index int)

	dp := make([][]bool, len(s))
	for i := range dp {
		dp[i] = make([]bool, len(s))
	}

	for right := 0; right < len(s); right++ {
		for left := 0; left <= right; left++ {
			if s[left] == s[right] && (right-left <= 2 || dp[left+1][right-1]) {
				dp[left][right] = true
			}
		}
	}

	bc = func(index int) {
		if index >= len(s) {
			res = append(res, append([]string{}, path...))
			return
		}
		for i := index; i < len(s); i++ {
			if dp[index][i] {
				path = append(path, s[index:i+1])
				bc(i + 1)
				path = path[:len(path)-1]
			}
		}
	}
	bc(0)
	return res
}
