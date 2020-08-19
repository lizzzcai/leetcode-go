package leetcode

import "sort"

func merge(intervals [][]int) [][]int {

	n := len(intervals)
	if n == 0 {
		return [][]int{}
	}
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	res := make([][]int, 0)
	res = append(res, intervals[0])
	for i := 1; i < n; i++ {
		l := len(res)
		pop := res[l-1]
		// if overlap
		if pop[1] >= intervals[i][0] {
			if intervals[i][1] > pop[1] {
				res[l-1][1] = intervals[i][1]
			}

		} else {
			res = append(res, intervals[i])
		}
	}
	return res
}
