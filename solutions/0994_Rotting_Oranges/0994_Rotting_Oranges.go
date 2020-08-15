package leetcode

type item struct {
	x int
	y int
}

func orangesRotting(grid [][]int) int {
	// init slice for rotting oranges and fresh oranges
	rotting := make([]item, 0)
	fresh := make([]item, 0)
	m, n := len(grid), len(grid[0])

	// iterate over the grid and append rotting and fresh oranges to respective group
	for r, row := range grid {
		for c, val := range row {
			if val == 1 {
				fresh = append(fresh, item{r, c})
			} else if val == 2 {
				rotting = append(rotting, item{r, c})
			}
		}
	}

	steps := 0
	dirs := [][]int{[]int{0, 1}, []int{0, -1}, []int{1, 0}, []int{-1, 0}}

	for len(fresh) > 0 {

		if len(rotting) == 0 {
			return -1
		}
		new_rotting := make([]item, 0)
		// iterate over rotting orange and check if any new rotting orange from the fresh
		for _, rot := range rotting {
			for _, d := range dirs {
				nr, nc := rot.x+d[0], rot.y+d[1]
				if 0 <= nr && nr < m && 0 <= nc && nc < n && grid[nr][nc] == 1 {
					new_rotting = append(new_rotting, item{nr, nc})
					// turn the fresh into rotting
					grid[nr][nc] = 2
				}
			}
		}
		steps += 1
		// remove the new rotting
		fresh = remove(fresh, new_rotting)
		// update the new rotting
		rotting = new_rotting

	}

	return steps
}

func remove(a, b []item) []item {
	check := make(map[item]bool)
	for _, x := range b {
		check[x] = true
	}

	res := make([]item, 0)
	for _, x := range a {
		if _, ok := check[x]; !ok {
			res = append(res, x)
		}
	}
	return res
}

func orangesRotting1(grid [][]int) int {
	// init slice for rotting oranges and fresh oranges
	var total, rotted int
	rotting := make([]item, 0)
	m, n := len(grid), len(grid[0])

	// iterate over the grid and append rotting and fresh oranges to respective group
	for r, row := range grid {
		for c, val := range row {
			if val == 1 || val == 2 {
				total += 1
			}
			if val == 2 {
				rotting = append(rotting, item{r, c})
				rotted += 1
			}
		}
	}

	if total == 0 {
		return 0
	}

	steps := 0
	dirs := [][]int{{0, 1}, {0, -1}, {1, 0}, {-1, 0}}

	for len(rotting) > 0 {
		size := len(rotting)
		for i := 0; i < size; i++ {
			rot := rotting[0]
			rotting = rotting[1:]

			for _, d := range dirs {
				nr, nc := rot.x+d[0], rot.y+d[1]
				if 0 <= nr && nr < m && 0 <= nc && nc < n && grid[nr][nc] == 1 {
					rotting = append(rotting, item{nr, nc})
					// turn the fresh into rotting
					grid[nr][nc] = 2
					rotted += 1
				}
			}
		}

		steps += 1
	}

	if total == rotted {
		return steps - 1
	}

	return -1
}
