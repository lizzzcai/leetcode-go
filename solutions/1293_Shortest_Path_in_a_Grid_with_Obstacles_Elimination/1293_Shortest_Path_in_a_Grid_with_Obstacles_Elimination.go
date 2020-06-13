package leetcode

type Visit struct {
	r, c, remain int
}
type Move struct {
	r, c, remain, steps int
}

func shortestPath(grid [][]int, k int) int {
	R, C := len(grid), len(grid[0])
	if R == 1 && C == 1 {
		return 0
	}

	// bfs queue to find shortest step
	queue := make([]Move, 0)
	// r, c, remain, steps
	queue = append(queue, Move{0, 0, k, 0})
	// visited to store the visited cell & state
	visited := make(map[Visit]bool)
	// r, c, remain
	visited[Visit{0, 0, k}] = true

	// dirs
	dirs := [][]int{[]int{1, 0}, []int{-1, 0}, []int{0, 1}, []int{0, -1}}

	// if we have enough k, we can break to the end directly
	if k >= R+C-2 {
		return R + C - 2
	}

	// start the bfs
	for len(queue) > 0 {
		popleft := queue[0]
		r, c, remain, steps := popleft.r, popleft.c, popleft.remain, popleft.steps

		for _, d := range dirs {
			nr, nc := r+d[0], c+d[1]
			if 0 <= nr && nr < R && 0 <= nc && nc < C {
				// get the state of this step
				state := remain - grid[nr][nc]
				if _, ok := visited[Visit{nr, nc, state}]; !ok && state >= 0 {
					// append into queue and add into visited
					queue = append(queue, Move{nr, nc, state, steps + 1})
					visited[Visit{nr, nc, state}] = true
				}

				if nr == R-1 && nc == C-1 {
					return steps + 1
				}
			}
		}
		queue = queue[1:]
	}

	return -1

}
