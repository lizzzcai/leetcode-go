package leetcode

func wordBreak(s string, wordDict []string) []string {
	// https://leetcode.com/problems/word-break-ii/discuss/44167/My-concise-JAVA-solution-based-on-memorized-DFS
	memo := make(map[string][]string)
	memo[""] = []string{""}

	return dfs(s, wordDict, memo)

}

func dfs(s string, wordDict []string, memo map[string][]string) []string {
	res, ok := memo[s]
	if ok {
		return res
	}

	res = []string{}

	for _, word := range wordDict {
		n := len(word)
		if len(s) >= n && s[:n] == word {
			sub_res := dfs(s[n:], wordDict, memo)
			for _, sub := range sub_res {
				var tmp string
				tmp = word
				if len(sub) > 0 {
					tmp += " " + sub
				}
				res = append(res, tmp)
			}
		}
	}

	memo[s] = res
	return memo[s]
}
