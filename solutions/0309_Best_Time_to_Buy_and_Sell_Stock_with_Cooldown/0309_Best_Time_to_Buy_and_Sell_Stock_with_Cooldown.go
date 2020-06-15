package leetcode

func maxProfit(prices []int) int {
	if len(prices) == 0 {
		return 0
	}
	n := len(prices)
	dp := make([]int, n+1)

	min_price := prices[0]
	// dp[i] = max(dp[i-1], prices[i] - prices[j] + dp[j-2])
	for i := 1; i < n+1; i++ {
		if i >= 2 {
			min_price = Min(min_price, prices[i-1]-dp[i-2])
		}
		dp[i] = Max(dp[i-1], prices[i-1]-min_price)
	}

	return dp[n]
}

func Min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func Max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
