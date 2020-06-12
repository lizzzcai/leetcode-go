package leetcode

func maximalSquare(matrix [][]byte) int {
	R := len(matrix)

	if R == 0 {
		return 0
	}
	C := len(matrix[0])

	dp := make([][]int, R+1)
	for i := 0; i < R+1; i++ {
		dp[i] = make([]int, C+1)
	}
	max_len := 0
	for i := 0; i < R; i++ {
		for j := 0; j < C; j++ {
			if matrix[i][j] == '0' {
				dp[i+1][j+1] = 0
			} else {
				dp[i+1][j+1] = Min(dp[i+1][j], dp[i][j+1], dp[i][j]) + 1
				max_len = Max(max_len, dp[i+1][j+1])
			}
		}
	}
	return max_len * max_len
}

func Min(nums ...int) int {
	res := nums[0]
	for _, x := range nums {
		if x < res {
			res = x
		}
	}
	return res
}

func Max(nums ...int) int {
	res := nums[0]
	for _, x := range nums {
		if x > res {
			res = x
		}
	}
	return res
}
