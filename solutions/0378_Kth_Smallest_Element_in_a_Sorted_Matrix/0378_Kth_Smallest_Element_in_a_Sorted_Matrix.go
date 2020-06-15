package leetcode

func kthSmallest(matrix [][]int, k int) int {
	R, C := len(matrix), len(matrix[0])

	l, h := matrix[0][0], matrix[R-1][C-1]

	for l <= h {
		mid := (l + h) / 2
		count := 0
		j := C - 1
		max_val := l
		for i := 0; i < R; i++ {
			for j >= 0 && matrix[i][j] > mid {
				j--
			}
			if j >= 0 {
				count += j + 1
				if matrix[i][j] > max_val {
					max_val = matrix[i][j]
				}
			}
		}

		if count == k {
			return max_val
		} else if count < k {
			l = mid + 1
		} else {
			h = mid - 1
		}
	}
	return l
}
