package leetcode

func canVisitAllRooms1(rooms [][]int) bool {
	n := len(rooms)
	visited := make([]bool, n)
	queue := make([]int, 0)

	// handle room zero
	visited[0] = true
	for _, room := range rooms[0] {
		queue = append(queue, room)
		visited[room] = true
	}

	var curr int
	for len(queue) > 0 {
		curr, queue = queue[0], queue[1:]
		for _, key := range rooms[curr] {
			if !visited[key] {
				queue = append(queue, key)
				visited[key] = true
			}
		}
	}

	for _, b := range visited {
		if !b {
			return false
		}
	}
	return true
}

func canVisitAllRooms(rooms [][]int) bool {
	n := len(rooms)
	visited := make([]bool, n)
	backtrack(0, rooms, visited)
	for _, b := range visited {
		if !b {
			return false
		}
	}
	return true
}

func backtrack(room int, rooms [][]int, visited []bool) {

	if visited[room] {
		return
	}

	visited[room] = true
	for _, key := range rooms[room] {
		backtrack(key, rooms, visited)
	}
}
