package leetcode

import "sort"

func eraseOverlapIntervals(intervals [][]int) int {
	n := len(intervals)
	if n == 0 {
		return 0
	}

	sort.SliceStable(intervals, func(i, j int) bool {
		return intervals[i][1] < intervals[j][1]
	})

	curr_end := intervals[0][1]
	ans := 0

	// start from second one
	for i := 1; i < n; i++ {
		// if start >= curr_end, no overlap
		if intervals[i][0] >= curr_end {
			curr_end = intervals[i][1]
		} else {
			ans += 1
		}
	}

	return ans
}
