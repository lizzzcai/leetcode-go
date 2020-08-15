package leetcode

func longestPalindrome(s string) int {
	count := make(map[byte]struct{})
	for i := 0; i < len(s); i++ {
		ch := s[i]
		if _, ok := count[ch]; !ok {
			count[ch] = struct{}{}
		} else {
			delete(count, ch)
		}
	}

	if len(count) == 0 {
		return len(s)
	}

	return len(s) - len(count) + 1

}
