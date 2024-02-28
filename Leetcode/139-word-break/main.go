package main

func wordBreak(s string, wordDict []string) bool {
	dp := make([]bool, len(s)+1)
	set := make(map[string]bool)
	for _, str := range wordDict {
		set[str] = true
	}
	dp[0] = true
	for i := 1; i <= len(s); i++ {
		for j := 0; j < i; j++ {
			substr := s[j:i]
			if set[substr] && dp[j] {
				dp[i] = true
			}
		}
	}
	return dp[len(s)]
}
