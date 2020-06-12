package leetcode

import (
	"math"
	"strings"
)

func minStickers(stickers []string, target string) int {
	cnt := make(map[rune]int)
	for _, x := range target {
		cnt[x]++
	}
	n := len(target)
	res := math.MaxInt32
	tar := []rune(target)

	var dfs func(map[rune]int, int, int)
	dfs = func(count map[rune]int, used, i int) {
		if i == n {
			if used < res {
				res = used
			}
		} else if count[tar[i]] >= cnt[tar[i]] {
			dfs(count, used, i+1)
		} else if used+1 < res {
			for _, sticker := range stickers {
				if strings.Contains(sticker, string(tar[i])) {
					for _, x := range sticker {
						count[x]++
					}
					dfs(count, used+1, i+1)
					for _, x := range sticker {
						count[x]--
					}
				}
			}
		}

	}

	dfs(make(map[rune]int), 0, 0)
	if res == math.MaxInt32 {
		return -1
	}
	return res
}
