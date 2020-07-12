package leetcode

import "sort"

func diagonalSort(mat [][]int) [][]int {
	R, C := len(mat), len(mat[0])
	diags := make(map[int][]int, 0)

	for r, row := range mat {
		for c, x := range row {
			diags[r-c] = append(diags[r-c], x)
		}
	}

	for _, v := range diags {
		sort.Ints(v)
	}

	for r := 0; r < R; r++ {
		for c := 0; c < C; c++ {
			mat[r][c] = diags[r-c][0]
			diags[r-c] = diags[r-c][1:]
		}
	}

	return mat
}
