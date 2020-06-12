package leetcode

func sortColors(nums []int) {
	r, w, b := 0, 0, len(nums)-1
	for w <= b {
		if nums[w] == 0 {
			nums[r], nums[w] = nums[w], nums[r]
			r++
			w++
		} else if nums[w] == 1 {
			w++
		} else {
			nums[b], nums[w] = nums[w], nums[b]
			b--
		}
	}
}
