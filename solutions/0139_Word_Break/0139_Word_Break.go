package leetcode

func wordBreak(s string, wordDict []string) bool {
	wordset := make(map[string]bool)
	for _, word := range wordDict {
		wordset[word] = true
	}

	dp := make([]bool, len(s)+1)
	dp[0] = true

	for i := 1; i <= len(s); i++ {
		for w, _ := range wordset {
			l := len(w)
			if i-l >= 0 && dp[i-l] && s[i-l:i] == w {
				dp[i] = dp[i-l]
				break
			}
		}
	}

	return dp[len(s)]
}
