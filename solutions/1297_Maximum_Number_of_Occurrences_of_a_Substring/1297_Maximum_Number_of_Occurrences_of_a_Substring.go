package leetcode

func maxFreq(s string, maxLetters int, minSize int, maxSize int) int {
	count := make(map[string]int)
	for i := 0; i < len(s)-minSize+1; i++ {
		count[s[i:i+minSize]] += 1
	}

	res := 0
	for key, value := range count {
		size := make(map[rune]bool)
		for _, c := range key {
			size[c] = true
		}
		if len(size) <= maxLetters {
			if value > res {
				res = value
			}
		}
	}
	return res
}
