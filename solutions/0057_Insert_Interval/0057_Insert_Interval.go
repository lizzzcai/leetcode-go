package leetcode

func insert(intervals [][]int, newInterval []int) [][]int {
	n := len(intervals)
	start, end := newInterval[0], newInterval[1]

	left, right := make([][]int, 0), make([][]int, 0)

	for i := 0; i < n; i++ {
		curr := intervals[i]
		if curr[1] < start {
			left = append(left, []int{curr[0], curr[1]})
		} else if curr[0] > end {
			right = append(right, []int{curr[0], curr[1]})
		} else {
			if curr[0] < start {
				start = curr[0]
			}
			if curr[1] > end {
				end = curr[1]
			}
		}
	}

	left = append(left, []int{start, end})
	return append(left, right...)
}
