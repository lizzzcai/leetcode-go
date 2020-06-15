package leetcode

import "container/heap"

type Path struct {
	price int // priority
	city  int
	k     int
}

type PriorityQueue []*Path

func (pq PriorityQueue) Len() int { return len(pq) }

// to get the lowest
func (pq PriorityQueue) Less(i, j int) bool { return pq[i].price < pq[j].price }

func (pq PriorityQueue) Swap(i, j int) { pq[i], pq[j] = pq[j], pq[i] }

func (pq *PriorityQueue) Push(x interface{}) {
	item := x.(*Path)
	*pq = append(*pq, item)
}

func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	old[n-1] = nil // avoid memory leak
	*pq = old[0 : n-1]
	return item
}

func findCheapestPrice(n int, flights [][]int, src int, dst int, K int) int {

	pq := make(PriorityQueue, 0)
	heap.Init(&pq)

	graph := make(map[int]map[int]int)
	for _, f := range flights {
		mm, ok := graph[f[0]]
		if !ok {
			mm = make(map[int]int)
			graph[f[0]] = mm
		}
		mm[f[1]] = f[2]
	}

	heap.Push(&pq, &Path{0, src, K + 1})
	for pq.Len() > 0 {
		path := heap.Pop(&pq).(*Path)

		if path.city == dst {
			return path.price
		}

		if path.k > 0 {
			if nxt, ok := graph[path.city]; ok {
				for d, cost := range nxt {
					heap.Push(&pq, &Path{path.price + cost, d, path.k - 1})
				}
			}
		}

	}

	return -1
}
