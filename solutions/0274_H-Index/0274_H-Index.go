package leetcode

import "sort"

func hIndex(citations []int) int {
	N := len(citations)
	sort.Ints(citations)
	for i := 0; i < N; i++ {
		if citations[i] >= N-i {
			return N - i
		}
	}
	return 0
}
