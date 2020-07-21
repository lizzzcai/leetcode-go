package leetcode

var dirs = [][]int{[]int{0, 1}, []int{0, -1}, []int{1, 0}, []int{-1, 0}}

func exist(board [][]byte, word string) bool {

	visited := make(map[int]bool)
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[0]); j++ {

			if dfs(i, j, board, 0, word, visited) {
				return true
			}
		}
	}

	return false
}

func dfs(i, j int, board [][]byte, idx int, word string, visited map[int]bool) bool {
	if idx == len(word) {
		return true
	}
	if i < 0 || i >= len(board) || j < 0 || j >= len(board[0]) || board[i][j] != word[idx] || visited[i*len(board[0])+j] == true {
		return false
	}
	result := false
	visited[i*len(board[0])+j] = true
	for _, d := range dirs {
		if dfs(i+d[0], j+d[1], board, idx+1, word, visited) == true {
			result = true
			break
		}
	}
	visited[i*len(board[0])+j] = false

	return result
}
