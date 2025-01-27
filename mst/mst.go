package mst

import (
	"ws/algorithms"
	"ws/graph"
	"ws/unionfind"
)

func MST(n int, edges []graph.Edge) (mst []graph.Edge, totalWeight int) {
	sort_edges := algorithms.Mergesort(edges)
	ds := unionfind.NewDisjointSet(n)
	totalWeight = 0

	for _, edge := range sort_edges {
		u, v := edge.GetUV()
		if ds.Find(u) != ds.Find(v) {
			ds.Union(u, v)
			mst = append(mst, edge)
			totalWeight += edge.GetWeight()
			if len(mst) == n-1 {
				break
			}
		}
	}
	return mst, totalWeight
}
