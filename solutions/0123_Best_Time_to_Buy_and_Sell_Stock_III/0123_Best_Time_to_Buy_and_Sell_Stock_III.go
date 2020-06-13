package leetcode

func maxProfit(prices []int) int {
	if len(prices) == 0 {
		return 0
	}
	n := len(prices)
	dp := make([][]int, 3)
	for i := 0; i < 3; i++ {
		dp[i] = make([]int, n+1)
	}

	// dp[k][i] = max(dp[k][i-1], prices[i] - prices[j] + dp[k-1][j-1])
	for k := 1; k < 3; k++ {
		min_price := prices[0]
		for i := 1; i < n+1; i++ {
			min_price = Min(min_price, prices[i-1]-dp[k-1][i-1])
			dp[k][i] = Max(dp[k][i-1], prices[i-1]-min_price)
		}
	}

	return dp[2][n]
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
