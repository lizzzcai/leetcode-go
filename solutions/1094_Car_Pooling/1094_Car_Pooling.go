package leetcode

func carPooling(trips [][]int, capacity int) bool {
	if len(trips) == 0 {
		return true
	}

	timestamp := make([]int, 1001)
	for _, trip := range trips {
		timestamp[trip[1]] += trip[0]
		timestamp[trip[2]] -= trip[0]
	}

	passengers := 0
	for _, change := range timestamp {
		passengers += change
		if passengers > capacity {
			return false
		}
	}

	return true
}
