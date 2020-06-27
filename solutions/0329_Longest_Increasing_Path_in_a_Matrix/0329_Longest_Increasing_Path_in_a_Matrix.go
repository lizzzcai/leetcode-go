package leetcode

func longestIncreasingPath(matrix [][]int) int {
	R := len(matrix)
	if R == 0 {
		return 0
	}

	C := len(matrix[0])

	// memo
	memo := make([][]int, R)
	for i := 0; i < R; i++ {
		memo[i] = make([]int, C)
	}

	ans := 0
	for i := 0; i < R; i++ {
		for j := 0; j < C; j++ {
			out := dfs(i, j, R, C, matrix, memo)
			if out > ans {
				ans = out
			}
		}
	}
	return ans
}

func dfs(i, j, R, C int, matrix, memo [][]int) int {
	if memo[i][j] > 0 {
		return memo[i][j]
	}

	res := 0
	for _, d := range [][]int{[]int{1, 0}, []int{-1, 0}, []int{0, 1}, []int{0, -1}} {
		nr, nc := i+d[0], j+d[1]
		if nr >= 0 && nr < R && nc >= 0 && nc < C && matrix[i][j] < matrix[nr][nc] {
			out := dfs(nr, nc, R, C, matrix, memo)
			if out > res {
				res = out
			}
		}
	}
	memo[i][j] = res + 1
	return memo[i][j]
}
