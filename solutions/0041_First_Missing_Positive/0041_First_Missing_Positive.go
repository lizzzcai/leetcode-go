package leetcode

func firstMissingPositive(nums []int) int {
	// first missing positive must within [1, n+1]
	n := len(nums)
	for i, val := range nums {
		if val < 1 || val >= n+1 {
			nums[i] = n + 1
		}
	}

	for _, val := range nums {
		idx := Abs(val) - 1
		if idx < n {
			nums[idx] = -Abs(nums[idx])
		}
	}

	for i, val := range nums {
		if val > 0 {
			return i + 1
		}
	}
	return n + 1
}

func Abs(val int) int {
	if val < 0 {
		return -val
	}
	return val
}
