package leetcode

func titleToNumber(s string) int {
	ans := 0
	for _, ch := range s {
		ans = ans*26 + int(ch-'A'+1)
	}

	return ans
}
