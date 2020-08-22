package leetcode

func constrainedSubsetSum(nums []int, k int) int {

	// decreasing mono queue to store max [dp(i-k), dp(i-k+1), ..., dp[i-1]]
	mq := make([]int, 0)
	ans := nums[0]
	for i := 0; i < len(nums); i++ {
		// add the max sum in dp(i-k), dp(i-k+1), ..., dp[i-1]
		if len(mq) > 0 {
			nums[i] += mq[0]
		}

		for len(mq) > 0 && mq[len(mq)-1] <= nums[i] {
			// pop out
			mq = mq[:len(mq)-1]
		}

		// if dp[i] > 0 append into m queue
		if nums[i] > 0 {
			mq = append(mq, nums[i])
		}

		// remove the dp out of range k
		for i >= k && len(mq) > 0 && mq[0] == nums[i-k] {
			mq = mq[1:len(mq)]
		}
		if nums[i] > ans {
			ans = nums[i]
		}
	}

	return ans
}
