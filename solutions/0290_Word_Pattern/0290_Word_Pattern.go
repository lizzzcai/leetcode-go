package leetcode

import "strings"

func wordPattern(pattern string, str string) bool {
	n := len(pattern)
	split := strings.Split(str, " ")
	if n != len(split) {
		return false
	}

	map_char := make(map[byte]string)
	map_word := make(map[string]byte)
	for i := 0; i < n; i++ {
		m1, ok1 := map_char[pattern[i]]
		_, ok2 := map_word[split[i]]

		if ok1 {
			if m1 != split[i] {
				return false
			}
		} else {
			if ok2 {
				return false
			} else {
				map_char[pattern[i]] = split[i]
				map_word[split[i]] = pattern[i]
			}
		}
	}
	return true
}
