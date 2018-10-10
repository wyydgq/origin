package main

func wordBreak(s string, wordDict []string) bool {
	dp := make([]int, len(s)+1)
	dp[0] = 1
	for i, _ := range s {
		if dp[i] == 0 {
			// 该位置不可到达
			continue
		}
		for _, k := range wordDict {
			t := i+len(k)
			if  t <= len(s) && s[i:t] == k { // 当前位有匹配的word
				dp[t] = dp[i] + 1
			}
		}
	}
	return dp[len(dp)-1] > 0
}