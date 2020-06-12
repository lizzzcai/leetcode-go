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

func get_primeFactors(x int) []int {
	primes := make(map[int]bool)
	d := 2
	for d*d <= x {
		if x%d == 0 {
			for x%d == 0 {
				x /= d
			}
			primes[d] = true
		}
		d++
	}
	if len(primes) == 0 || x > 1 {
		primes[x] = true
	}

	keys := make([]int, 0, len(primes))
	for k := range primes {
		keys = append(keys, k)
	}
	return keys
}

func largestComponentSize(A []int) int {

	dsu := &DSU{make(map[int]int), make(map[int]int)}
	B := make([][]int, len(A))
	prime_to_index := make(map[int]int)
	j := 0
	for i, x := range A {
		primes := get_primeFactors(x)
		B[i] = primes
		for _, p := range primes {
			prime_to_index[p] = j
			j++
		}
	}

	for _, facs := range B {
		for _, x := range facs {
			dsu.union(prime_to_index[x], prime_to_index[facs[0]])
		}
	}

	count := make(map[int]int)
	res := 0
	for _, facs := range B {
		count[dsu.find(prime_to_index[facs[0]])] += 1
		res = Max(res, count[dsu.find(prime_to_index[facs[0]])])
	}

	return res
}

func Max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
