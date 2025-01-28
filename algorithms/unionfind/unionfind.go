package unionfind

type DisjointSet struct {
	parent []int
	rank   []int
	// или size []int — по желанию
}

func NewDisjointSet(n int) *DisjointSet {
	d := &DisjointSet{parent: make([]int, n), rank: make([]int, n)}
	for i := 0; i < n; i++ {
		d.parent[i] = i
		d.rank[i] = 0
	}
	return d
}

func (d *DisjointSet) Find(x int) int {
	if d.parent[x] != x {
		d.parent[x] = d.Find(d.parent[x])
	}
	return d.parent[x]
}

func (d *DisjointSet) Union(x, y int) bool {
	rx := d.Find(x)
	ry := d.Find(y)
	if rx == ry { // уже в одной компоненте
		return false
	}
	if d.rank[rx] < d.rank[ry] {
		d.parent[rx] = ry
	} else {
		d.parent[ry] = rx
		if d.rank[rx] == d.rank[ry] {
			d.rank[rx]++
		}
	}
	return true
}

func (d *DisjointSet) GetParent() []int {
	return d.parent
}

func (d *DisjointSet) GetRank() []int {
	return d.rank
}
