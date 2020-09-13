package leetcode

import "sort"

func medianSlidingWindow(nums []int, k int) []float64 {
	win := make([]int, 0)
	for i := 0; i < k; i++ {
		win = append(win, nums[i])
	}

	sort.Ints(win)
	is_odd := k%2 == 1
	ans := make([]float64, 0)
	var median float64
	var idx int

	for i := k; i < len(nums); i++ {
		median = float64(win[k/2])
		if !is_odd {
			median = float64(win[k/2]+win[k/2-1]) * 0.5
		}
		ans = append(ans, median)

		// find the index of the last one
		idx = find(win, nums[i-k]) - 1
		// pop the last one out
		win = append(win[:idx], win[idx+1:]...)
		// find the insert point for current one
		idx = find(win, nums[i])
		// insert current into window
		win = append(win[:idx], append([]int{nums[i]}, win[idx:]...)...)
	}
	// add the last one
	median = float64(win[k/2])
	if !is_odd {
		median = float64(win[k/2]+win[k/2-1]) * 0.5
	}
	ans = append(ans, median)

	return ans
}

func find(arr []int, target int) int {
	l, r := 0, len(arr)-1
	for l <= r {
		mid := (l + r) / 2
		if arr[mid] > target {
			r = mid - 1
		} else if arr[mid] < target {
			l = mid + 1
		} else {
			l = mid + 1
		}
	}
	return l
}
