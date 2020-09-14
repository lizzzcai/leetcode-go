package leetcode

func findAllConcatenatedWordsInADict(words []string) []string {
	words_set := make(map[string]bool)
	for i := 0; i < len(words); i++ {
		words_set[words[i]] = true
	}

	var dfs func(word string) bool
	dfs = func(word string) bool {
		for i := 1; i < len(word); i++ {
			prefix := word[:i]
			suffix := word[i:]

			_, ok1 := words_set[prefix]
			_, ok2 := words_set[suffix]
			if ok1 && ok2 {
				return true
			} else if ok1 && dfs(suffix) {
				return true
			} else if ok2 && dfs(prefix) {
				return true
			}
		}
		return false
	}
	ans := make([]string, 0)
	for i := 0; i < len(words); i++ {
		if dfs(words[i]) {
			ans = append(ans, words[i])
		}
	}
	return ans
}
