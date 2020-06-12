package leetcode

type DSU struct {
	par map[int]int
	rnk map[int]int
}

func (dsu DSU) find(x int) int {
	if dsu.par[x] != x {
		dsu.par[x] = dsu.find(dsu.par[x])
	}
	return dsu.par[x]
}

func (dsu DSU) _add(x int) {
	if _, ok := dsu.par[x]; !ok {
		dsu.par[x] = x
		dsu.rnk[x] = 0
	}
}

func (dsu DSU) union(x, y int) bool {

	if _, ok := dsu.par[x]; !ok {
		dsu._add(x)
	}

	if _, ok := dsu.par[y]; !ok {
		dsu._add(y)
	}
	// get the leaders
	xr, yr := dsu.find(x), dsu.find(y)

	// if match, they are belong to same union
	if xr == yr {
		return false
	}

	// check the rank
	if dsu.rnk[xr] < dsu.rnk[yr] {
		dsu.par[xr] = yr
	} else {
		dsu.par[yr] = xr
		if dsu.rnk[xr] == dsu.rnk[yr] {
			dsu.rnk[xr] += 1
		}
	}
	return true
}

func findRedundantConnection(edges [][]int) []int {
	dsu := &DSU{make(map[int]int), make(map[int]int)}
	for _, edge := range edges {
		if !dsu.union(edge[0], edge[1]) {
			return edge
		}
	}

	return nil
}
